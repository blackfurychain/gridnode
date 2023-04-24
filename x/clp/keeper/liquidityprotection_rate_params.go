package keeper

import (
	"github.com/Gridironchain/gridnode/x/clp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLiquidityProtectionRateParams(ctx sdk.Context, params types.LiquidityProtectionRateParams) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LiquidityProtectionRateParamsPrefix, k.cdc.MustMarshal(&params))
}

func (k Keeper) GetLiquidityProtectionRateParams(ctx sdk.Context) types.LiquidityProtectionRateParams {
	params := types.LiquidityProtectionRateParams{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LiquidityProtectionRateParamsPrefix)
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

func (k Keeper) SetLiquidityProtectionCurrentFuryLiquidityThreshold(ctx sdk.Context, currentFuryLiquidityThreshold sdk.Uint) {
	currentParams := k.GetLiquidityProtectionRateParams(ctx)
	currentParams.CurrentFuryLiquidityThreshold = currentFuryLiquidityThreshold
	k.SetLiquidityProtectionRateParams(ctx, currentParams)
}
