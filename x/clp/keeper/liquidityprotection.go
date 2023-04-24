package keeper

import (
	"errors"

	"github.com/Gridironchain/gridnode/x/clp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLiquidityProtectionParams(ctx sdk.Context, params *types.LiquidityProtectionParams) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LiquidityProtectionParamsPrefix, k.cdc.MustMarshal(params))
}

func (k Keeper) GetLiquidityProtectionParams(ctx sdk.Context) *types.LiquidityProtectionParams {
	params := types.LiquidityProtectionParams{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LiquidityProtectionParamsPrefix)
	k.cdc.MustUnmarshal(bz, &params)
	return &params
}

// This method should only be called if buying or selling native asset.
// If sellNative is false then this method assumes that buyNative is true.
// The nativePrice should be in MaxFuryLiquidityThresholdAsset
// NOTE: this method panics if sellNative is true and the value of the sell amount
// is greater than the value of currentFuryLiquidityThreshold. Call IsBlockedByLiquidityProtection
// before if unsure.
func (k Keeper) MustUpdateLiquidityProtectionThreshold(ctx sdk.Context, sellNative bool, nativeAmount sdk.Uint, nativePrice sdk.Dec) {
	liquidityProtectionParams := k.GetLiquidityProtectionParams(ctx)
	maxFuryLiquidityThreshold := liquidityProtectionParams.MaxFuryLiquidityThreshold
	currentFuryLiquidityThreshold := k.GetLiquidityProtectionRateParams(ctx).CurrentFuryLiquidityThreshold

	if liquidityProtectionParams.IsActive {
		nativeValue := CalcFuryValue(nativeAmount, nativePrice)

		var updatedFuryLiquidityThreshold sdk.Uint
		if sellNative {
			if currentFuryLiquidityThreshold.LT(nativeValue) {
				panic(errors.New("expect sell native value to be less than currentFuryLiquidityThreshold"))
			} else {
				updatedFuryLiquidityThreshold = currentFuryLiquidityThreshold.Sub(nativeValue)
			}
		} else {
			// This is equivalent to currentFuryLiquidityThreshold := sdk.MinUint(currentFuryLiquidityThreshold.Add(nativeValue), maxFuryLiquidityThreshold)
			// except it prevents any overflows when adding the nativeValue
			// Assume that maxFuryLiquidityThreshold >= currentFuryLiquidityThreshold
			if maxFuryLiquidityThreshold.Sub(currentFuryLiquidityThreshold).LT(nativeValue) {
				updatedFuryLiquidityThreshold = maxFuryLiquidityThreshold
			} else {
				updatedFuryLiquidityThreshold = currentFuryLiquidityThreshold.Add(nativeValue)
			}
		}

		k.SetLiquidityProtectionCurrentFuryLiquidityThreshold(ctx, updatedFuryLiquidityThreshold)
	}
}

// Currently this calculates the native price on the fly
// Calculates the price of the native token in MaxFuryLiquidityThresholdAsset
func (k Keeper) GetNativePrice(ctx sdk.Context) (sdk.Dec, error) {
	liquidityProtectionParams := k.GetLiquidityProtectionParams(ctx)
	maxFuryLiquidityThresholdAsset := liquidityProtectionParams.MaxFuryLiquidityThresholdAsset

	if types.StringCompare(maxFuryLiquidityThresholdAsset, types.NativeSymbol) {
		return sdk.OneDec(), nil
	}
	pool, err := k.GetPool(ctx, maxFuryLiquidityThresholdAsset)
	if err != nil {
		return sdk.Dec{}, types.ErrMaxFuryLiquidityThresholdAssetPoolDoesNotExist
	}

	return CalcFurySpotPrice(&pool, k.GetPmtpRateParams(ctx).PmtpCurrentRunningRate)

}

// The nativePrice should be in MaxFuryLiquidityThresholdAsset
func (k Keeper) IsBlockedByLiquidityProtection(ctx sdk.Context, nativeAmount sdk.Uint, nativePrice sdk.Dec) bool {
	value := CalcFuryValue(nativeAmount, nativePrice)
	currentFuryLiquidityThreshold := k.GetLiquidityProtectionRateParams(ctx).CurrentFuryLiquidityThreshold
	return currentFuryLiquidityThreshold.LT(value)
}
