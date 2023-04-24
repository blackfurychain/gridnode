package keeper

import (
	"encoding/json"
	"strconv"

	"github.com/Gridironchain/gridnode/x/clp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PoolFuryMap map[*types.Pool]sdk.Uint
type LpFuryMap map[string]sdk.Uint
type LpPoolMap map[string][]LPPool

func (k Keeper) ProviderDistributionPolicyRun(ctx sdk.Context) {
	a, b, c := k.doProviderDistribution(ctx)
	k.TransferProviderDistribution(ctx, a, b, c)
}

func (k Keeper) doProviderDistribution(ctx sdk.Context) (PoolFuryMap, LpFuryMap, LpPoolMap) {
	blockHeight := ctx.BlockHeight()
	params := k.GetProviderDistributionParams(ctx)
	if params == nil {
		return make(PoolFuryMap), make(LpFuryMap), make(LpPoolMap)
	}

	period := FindProviderDistributionPeriod(blockHeight, params.DistributionPeriods)
	if period == nil {
		return make(PoolFuryMap), make(LpFuryMap), make(LpPoolMap)
	}

	allPools := k.GetPools(ctx)
	return k.CollectProviderDistributions(ctx, allPools, period.DistributionPeriodBlockRate)
}

func (k Keeper) TransferProviderDistribution(ctx sdk.Context, poolFuryMap PoolFuryMap, lpFuryMap LpFuryMap, lpPoolMap LpPoolMap) {
	k.TransferProviderDistributionGeneric(ctx, poolFuryMap, lpFuryMap, lpPoolMap, "lppd/liquidity_provider_payout_error", "lppd/distribution")

	for pool, sub := range poolFuryMap {
		// will never fail
		k.RemoveFuryFromPool(ctx, pool, sub) // nolint:errcheck
	}
}

func (k Keeper) TransferProviderDistributionGeneric(ctx sdk.Context, poolFuryMap PoolFuryMap, lpFuryMap LpFuryMap, lpPoolMap LpPoolMap, typeStr string, successEventType string) {
	for lpAddress, totalFury := range lpFuryMap {
		addr, _ := sdk.AccAddressFromBech32(lpAddress) // We know this can't fail as we previously filtered out invalid strings
		coin := sdk.NewCoin(types.NativeSymbol, sdk.NewIntFromBigInt(totalFury.BigInt()))

		//TransferCoinsFromPool(pool, provider_fury, provider_address)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin))
		if err != nil {
			fireLPPayoutErrorEvent(ctx, addr, typeStr, err)

			for _, lpPool := range lpPoolMap[lpAddress] {
				poolFuryMap[lpPool.Pool] = poolFuryMap[lpPool.Pool].Sub(lpPool.Amount)
			}
		} else {
			fireDistributeSuccessEvent(ctx, lpAddress, lpPoolMap[lpAddress], totalFury, successEventType)
		}
	}
}

func fireDistributeSuccessEvent(ctx sdk.Context, lpAddress string, pools []LPPool, totalDistributed sdk.Uint, typeStr string) {
	data := PrintPools(pools)
	successEvent := sdk.NewEvent(
		typeStr,
		sdk.NewAttribute("recipient", lpAddress),
		sdk.NewAttribute("total_amount", totalDistributed.String()),
		sdk.NewAttribute("amounts", data),
	)

	ctx.EventManager().EmitEvents(sdk.Events{successEvent})
}

type FormattedPool struct {
	Pool   string   `json:"pool"`
	Amount sdk.Uint `json:"amount"`
}

func PrintPools(pools []LPPool) string {
	var formattedPools = make([]FormattedPool, len(pools))

	for i, pool := range pools {
		formattedPools[i] = FormattedPool{Pool: pool.Pool.ExternalAsset.Symbol, Amount: pool.Amount}
	}

	data, _ := json.Marshal(formattedPools) // as used, this should never return an error
	return string(data)
}

func fireLPPayoutErrorEvent(ctx sdk.Context, address sdk.AccAddress, typeStr string, err error) {
	failureEvent := sdk.NewEvent(
		typeStr,
		sdk.NewAttribute("liquidity_provider", address.String()),
		sdk.NewAttribute(types.AttributeKeyError, err.Error()),
		sdk.NewAttribute(types.AttributeKeyHeight, strconv.FormatInt(ctx.BlockHeight(), 10)),
	)

	ctx.EventManager().EmitEvents(sdk.Events{failureEvent})
}

//nolint
func fireDistributionEvent(ctx sdk.Context, amount sdk.Uint, to sdk.Address) {
	coin := sdk.NewCoin(types.NativeSymbol, sdk.NewIntFromBigInt(amount.BigInt()))
	distribtionEvent := sdk.NewEvent(
		types.EventTypeProviderDistributionDistribution,
		sdk.NewAttribute(types.AttributeProbiverDistributionAmount, coin.String()),
		sdk.NewAttribute(types.AttributeProbiverDistributionReceiver, to.String()),
		sdk.NewAttribute(types.AttributeKeyHeight, strconv.FormatInt(ctx.BlockHeight(), 10)),
	)

	ctx.EventManager().EmitEvents(sdk.Events{distribtionEvent})
}

func FindProviderDistributionPeriod(currentHeight int64, periods []*types.ProviderDistributionPeriod) *types.ProviderDistributionPeriod {
	for _, period := range periods {
		if isActivePeriod(currentHeight, period.DistributionPeriodStartBlock, period.DistributionPeriodEndBlock) {
			return period
		}
	}

	return nil
}

func isActivePeriod(current int64, start, end uint64) bool {
	return current >= int64(start) && current <= int64(end)
}

func (k Keeper) CollectProviderDistributions(ctx sdk.Context, pools []*types.Pool, blockRate sdk.Dec) (PoolFuryMap, LpFuryMap, LpPoolMap) {
	poolFuryMap := make(PoolFuryMap, len(pools))
	lpMap := make(LpFuryMap, 0)
	lpPoolMap := make(LpPoolMap, 0)

	partitions, err := k.GetAllLiquidityProvidersPartitions(ctx)
	if err != nil {
		fireLPPGetLPsErrorEvent(ctx, err)
	}

	for _, pool := range pools {
		lps, exists := partitions[*pool.ExternalAsset]
		if !exists { // TODO: fire event
			continue
		}
		lpsFiltered := FilterValidLiquidityProviders(ctx, lps)
		furyToDistribute := CollectProviderDistribution(ctx, pool, sdk.NewDecFromBigInt(pool.NativeAssetBalance.BigInt()),
			blockRate, pool.PoolUnits, lpsFiltered, lpMap, lpPoolMap)
		poolFuryMap[pool] = furyToDistribute
	}

	return poolFuryMap, lpMap, lpPoolMap
}

type ValidLiquidityProvider struct {
	Address sdk.AccAddress
	LP      *types.LiquidityProvider
}

type LPPool struct {
	Pool   *types.Pool
	Amount sdk.Uint
}

func PoolFuryMapToLPPools(poolFuryMap PoolFuryMap) []LPPool {
	arr := make([]LPPool, 0, len(poolFuryMap))

	for pool, coins := range poolFuryMap {
		arr = append(arr, LPPool{Pool: pool, Amount: coins})
	}

	return arr
}

func FilterValidLiquidityProviders(ctx sdk.Context, lps []*types.LiquidityProvider) []ValidLiquidityProvider {
	var valid []ValidLiquidityProvider //nolint

	for _, lp := range lps {
		address, err := sdk.AccAddressFromBech32(lp.LiquidityProviderAddress)
		if err != nil {
			//k.Logger(ctx).Error(fmt.Sprintf("Liquidity provider address %s error %s", lp.LiquidityProviderAddress, err.Error()))
			fireLPAddressErrorEvent(ctx, lp.LiquidityProviderAddress, err)
			continue
		}

		valid = append(valid, ValidLiquidityProvider{Address: address, LP: lp})
	}

	return valid
}

func fireLPPGetLPsErrorEvent(ctx sdk.Context, err error) {
	failureEvent := sdk.NewEvent(
		"lppd/get_liquidity_providers_error",
		sdk.NewAttribute(types.AttributeKeyError, err.Error()),
		sdk.NewAttribute(types.AttributeKeyHeight, strconv.FormatInt(ctx.BlockHeight(), 10)),
	)

	ctx.EventManager().EmitEvents(sdk.Events{failureEvent})
}

func CollectProviderDistribution(ctx sdk.Context, pool *types.Pool, poolDepthFury, blockRate sdk.Dec, poolUnits sdk.Uint, lps []ValidLiquidityProvider, globalLpFuryMap LpFuryMap, globalLpPoolMap LpPoolMap) sdk.Uint {
	totalFuryDistribute := sdk.ZeroUint()

	//	fury_provider_distribution = r_block * pool_depth_fury
	furyPd := blockRate.Mul(poolDepthFury)
	furyPdUint := sdk.NewUintFromBigInt(furyPd.RoundInt().BigInt())
	for _, lp := range lps {
		providerFury := CalcProviderDistributionAmount(furyPd, poolUnits, lp.LP.LiquidityProviderUnits)
		totalFuryDistribute = totalFuryDistribute.Add(providerFury)

		// TODO: find a proper solution
		if totalFuryDistribute.GT(furyPdUint) {
			providerFury = furyPdUint.Sub(totalFuryDistribute.Sub(providerFury))
			totalFuryDistribute = furyPdUint
		}

		addr := lp.Address.String()

		globalLpFury := globalLpFuryMap[addr]
		if globalLpFury == (sdk.Uint{}) {
			globalLpFury = sdk.ZeroUint()
		}
		globalLpFuryMap[addr] = globalLpFury.Add(providerFury)

		elem := LPPool{Pool: pool, Amount: providerFury}
		globalLpPool := globalLpPoolMap[addr]
		if globalLpPool == nil {
			arr := []LPPool{elem}
			globalLpPoolMap[addr] = arr
		} else {
			globalLpPool = append(globalLpPool, elem)
			globalLpPoolMap[addr] = globalLpPool
		}
	}

	return totalFuryDistribute
}

func fireLPAddressErrorEvent(ctx sdk.Context, address string, err error) {
	failureEvent := sdk.NewEvent(
		"lppd/liquidity_provider_address_error",
		sdk.NewAttribute("liquidity_provider", address),
		sdk.NewAttribute(types.AttributeKeyError, err.Error()),
		sdk.NewAttribute(types.AttributeKeyHeight, strconv.FormatInt(ctx.BlockHeight(), 10)),
	)

	ctx.EventManager().EmitEvents(sdk.Events{failureEvent})
}

func CalcProviderDistributionAmount(furyProviderDistribution sdk.Dec, totalPoolUnits, providerPoolUnits sdk.Uint) sdk.Uint {
	//provider_percentage = provider_units / total_pool_units
	providerPercentage := sdk.NewDecFromBigInt(providerPoolUnits.BigInt()).Quo(sdk.NewDecFromBigInt(totalPoolUnits.BigInt()))

	//provider_fury = provider_percentage * fury_provider_distribution
	providerFury := providerPercentage.Mul(furyProviderDistribution)

	return sdk.Uint(providerFury.RoundInt())
}

func (k Keeper) SetProviderDistributionParams(ctx sdk.Context, params *types.ProviderDistributionParams) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ProviderDistributionParamsPrefix, k.cdc.MustMarshal(params))
}

func (k Keeper) GetProviderDistributionParams(ctx sdk.Context) *types.ProviderDistributionParams {
	params := types.ProviderDistributionParams{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ProviderDistributionParamsPrefix)
	k.cdc.MustUnmarshal(bz, &params)

	return &params
}

func (k Keeper) IsDistributionBlock(ctx sdk.Context) bool {
	blockHeight := ctx.BlockHeight()
	params := k.GetProviderDistributionParams(ctx)
	period := FindProviderDistributionPeriod(blockHeight, params.DistributionPeriods)
	if period == nil {
		return false
	}

	startHeight := period.DistributionPeriodStartBlock
	mod := period.DistributionPeriodMod

	return IsDistributionBlockPure(blockHeight, startHeight, mod)
}

// do the thing every mod blocks starting at startHeight
func IsDistributionBlockPure(blockHeight int64, startHeight, mod uint64) bool {
	return (blockHeight-int64(startHeight))%int64(mod) == 0
}
