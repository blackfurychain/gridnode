package clp

import (
	"fmt"
	"strconv"
	"time"

	kpr "github.com/Gridironchain/gridnode/x/clp/keeper"
	"github.com/Gridironchain/gridnode/x/clp/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, keeper kpr.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if keeper.IsDistributionBlock(ctx) {
		keeper.ProviderDistributionPolicyRun(ctx)
	}

	params := keeper.GetRewardsParams(ctx)
	pools := keeper.GetPools(ctx)
	currentPeriod := keeper.GetCurrentRewardPeriod(ctx, params)
	if currentPeriod != nil && !currentPeriod.RewardPeriodAllocation.IsZero() {

		isDistributionBlock := kpr.IsDistributionBlockPure(ctx.BlockHeight(), currentPeriod.RewardPeriodStartBlock, currentPeriod.RewardPeriodMod)

		currentBlockDistribution := kpr.CalcBlockDistribution(currentPeriod)
		blockDistributionAccu := keeper.GetBlockDistributionAccu(ctx)
		blockDistribution := blockDistributionAccu.Add(currentBlockDistribution)
		if isDistributionBlock {
			err := keeper.DistributeDepthRewards(ctx, blockDistribution, currentPeriod, pools)
			keeper.SetBlockDistributionAccu(ctx, sdk.ZeroUint())
			if err != nil {
				keeper.Logger(ctx).Error(fmt.Sprintf("Rewards policy run error %s", err.Error()))
			}
		} else {
			keeper.SetBlockDistributionAccu(ctx, blockDistribution)
		}
	}

	// res, stop := keeper.BalanceModuleAccountCheck()(ctx)
	// if stop {
	// 	// replace panic with an error log
	// 	// panic(res)
	// 	keeper.Logger(ctx).Error(res)
	// }

	// res, stop = keeper.UnitsCheck()(ctx)
	// if stop {
	// 	keeper.Logger(ctx).Error(res)
	// }

	return []abci.ValidatorUpdate{}
}

func BeginBlocker(ctx sdk.Context, k kpr.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	MeasureBlockTime(ctx, k)

	// get current block height
	currentHeight := ctx.BlockHeight()

	/*
		Liquidity protection current fury liquidity threshold update
	*/
	liquidityProtectionParams := k.GetLiquidityProtectionParams(ctx)
	maxFuryLiquidityThreshold := liquidityProtectionParams.MaxFuryLiquidityThreshold
	maxFuryLiquidityThresholdAsset := liquidityProtectionParams.MaxFuryLiquidityThresholdAsset
	if liquidityProtectionParams.IsActive {
		currentFuryLiquidityThreshold := k.GetLiquidityProtectionRateParams(ctx).CurrentFuryLiquidityThreshold
		// Validation check ensures that Epoch length =/= zero
		replenishmentAmount := maxFuryLiquidityThreshold.QuoUint64(liquidityProtectionParams.EpochLength)

		// This is equivalent to:
		//    proposedThreshold := currentFuryLiquidityThreshold.Add(replenishmentAmount)
		//    currentFuryLiquidityThreshold = sdk.MinUint(proposedThreshold, maxFuryLiquidityThreshold)
		// except it prevents any overflows when adding the replenishmentAmount
		if maxFuryLiquidityThreshold.Sub(currentFuryLiquidityThreshold).LT(replenishmentAmount) {
			currentFuryLiquidityThreshold = maxFuryLiquidityThreshold
		} else {
			currentFuryLiquidityThreshold = currentFuryLiquidityThreshold.Add(replenishmentAmount)
		}

		k.SetLiquidityProtectionCurrentFuryLiquidityThreshold(ctx, currentFuryLiquidityThreshold)
		k.Logger(ctx).Info(fmt.Sprintf("Liquidity Protection | maxFuryLiquidityThreshold: %s | asset: %s | currentFuryLiquidityThreshold: %s | maxPerBlock: %s", maxFuryLiquidityThreshold, maxFuryLiquidityThresholdAsset, k.GetLiquidityProtectionRateParams(ctx).CurrentFuryLiquidityThreshold, replenishmentAmount))
	}

	// get PMTP period params
	pmtpPeriodStartBlock := k.GetPmtpParams(ctx).PmtpPeriodStartBlock
	pmtpPeriodEndBlock := k.GetPmtpParams(ctx).PmtpPeriodEndBlock
	// Start Policy
	if currentHeight == pmtpPeriodStartBlock &&
		k.GetPmtpEpoch(ctx).EpochCounter == 0 &&
		k.GetPmtpEpoch(ctx).BlockCounter == 0 {
		k.PolicyStart(ctx)
		_ = ctx.EventManager().EmitTypedEvent(&types.EventPolicy{
			EventType:            "policy_start",
			PmtpPeriodStartBlock: strconv.Itoa(int(pmtpPeriodStartBlock)),
			PmtpPeriodEndBlock:   strconv.Itoa(int(pmtpPeriodEndBlock)),
		})
		k.Logger(ctx).Info(fmt.Sprintf("Starting new policy | Start Height : %d | End Height : %d", pmtpPeriodStartBlock, pmtpPeriodEndBlock))
	}
	// default to current pmtp current running rate value
	pmtpCurrentRunningRate := k.GetPmtpRateParams(ctx).PmtpCurrentRunningRate

	/*
		Epoch counters are used to keep track of policy execution
		EpochCounter tracks the number of Epochs left (Decrementing counter )
		BlockCounter tracks the number of Blocks left (Decrementing counter ) in an Epoch
	*/

	// Manage Block Counter
	if currentHeight >= pmtpPeriodStartBlock &&
		currentHeight <= pmtpPeriodEndBlock &&
		// TODO : Check this condition might not be needed
		k.GetPmtpEpoch(ctx).EpochCounter > 0 {
		// Calculate R running for policy params
		pmtpCurrentRunningRate = k.PolicyCalculations(ctx)
		k.DecrementPmtpBlockCounter(ctx)
	}
	// Manage Epoch Counter
	if k.GetPmtpEpoch(ctx).BlockCounter == 0 &&
		currentHeight < pmtpPeriodEndBlock &&
		currentHeight >= pmtpPeriodStartBlock {
		k.DecrementPmtpEpochCounter(ctx)
		k.SetPmtpBlockCounter(ctx, k.GetPmtpParams(ctx).PmtpPeriodEpochLength)
	}

	if currentHeight == pmtpPeriodEndBlock {
		// Setting it to zero to check when we start policy
		k.SetPmtpEpoch(ctx, types.PmtpEpoch{
			EpochCounter: 0,
			BlockCounter: 0,
		})
		// Set inter policy rate to running rate
		k.SetPmtpInterPolicyRate(ctx, pmtpCurrentRunningRate)
		_ = ctx.EventManager().EmitTypedEvent(&types.EventPolicy{
			EventType:            "policy_end",
			PmtpPeriodStartBlock: strconv.Itoa(int(pmtpPeriodStartBlock)),
			PmtpPeriodEndBlock:   strconv.Itoa(int(pmtpPeriodEndBlock)),
		})
		k.Logger(ctx).Info(fmt.Sprintf("Ending Policy | Start Height : %d | End Height : %d", pmtpPeriodStartBlock, pmtpPeriodEndBlock))
	}

	err := k.PolicyRun(ctx, pmtpCurrentRunningRate)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("error in running policy | Error Message : %s ", err.Error()))
	}

}

var blockTime *time.Time

func MeasureBlockTime(ctx sdk.Context, k kpr.Keeper) {
	now := time.Now()
	if blockTime == nil {
		blockTime = &now
		return
	}

	elapsed := now.Sub(*blockTime)
	blockTime = &now
	k.Logger(ctx).Info(fmt.Sprint("Block took ", elapsed.Seconds(), "s to execute"))
}
