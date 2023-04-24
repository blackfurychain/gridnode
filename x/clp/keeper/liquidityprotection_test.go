package keeper_test

import (
	"testing"

	gridapp "github.com/Gridironchain/gridnode/app"
	"github.com/Gridironchain/gridnode/x/clp/test"
	"github.com/Gridironchain/gridnode/x/clp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_GetNativePrice(t *testing.T) {
	testcases := []struct {
		name                     string
		pricingAsset             string
		createPool               bool
		poolNativeAssetBalance   sdk.Uint
		poolExternalAssetBalance sdk.Uint
		pmtpCurrentRunningRate   sdk.Dec
		expectedPrice            sdk.Dec
		expectedError            error
	}{
		{
			name:          "success",
			pricingAsset:  types.NativeSymbol,
			expectedPrice: sdk.NewDec(1),
		},
		{
			name:          "fail pool does not exist",
			pricingAsset:  "usdc",
			expectedError: types.ErrMaxFuryLiquidityThresholdAssetPoolDoesNotExist,
		},
		{
			name:                     "success non fury pricing asset",
			pricingAsset:             "usdc",
			createPool:               true,
			poolNativeAssetBalance:   sdk.NewUint(100000),
			poolExternalAssetBalance: sdk.NewUint(1000),
			pmtpCurrentRunningRate:   sdk.OneDec(),
			expectedPrice:            sdk.MustNewDecFromStr("0.02"),
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *gridapp.GridironchainApp, genesisState gridapp.GenesisState) gridapp.GenesisState {

				if tc.createPool {
					pools := []*types.Pool{
						{
							ExternalAsset:        &types.Asset{Symbol: tc.pricingAsset},
							NativeAssetBalance:   tc.poolNativeAssetBalance,
							ExternalAssetBalance: tc.poolExternalAssetBalance,
						},
					}
					clpGs := types.DefaultGenesisState()

					clpGs.Params = types.Params{
						MinCreatePoolThreshold: 1,
					}
					clpGs.PoolList = append(clpGs.PoolList, pools...)
					bz, _ := app.AppCodec().MarshalJSON(clpGs)
					genesisState["clp"] = bz
				}

				return genesisState
			})

			liquidityProtectionParams := app.ClpKeeper.GetLiquidityProtectionParams(ctx)
			liquidityProtectionParams.MaxFuryLiquidityThresholdAsset = tc.pricingAsset
			app.ClpKeeper.SetLiquidityProtectionParams(ctx, liquidityProtectionParams)
			app.ClpKeeper.SetPmtpCurrentRunningRate(ctx, tc.pmtpCurrentRunningRate)

			price, err := app.ClpKeeper.GetNativePrice(ctx)

			if tc.expectedError != nil {
				require.EqualError(t, err, tc.expectedError.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedPrice.String(), price.String()) // compare strings so that the expected amounts can be read from the failure message
		})
	}
}

func TestKeeper_IsBlockedByLiquidityProtection(t *testing.T) {
	testcases := []struct {
		name                           string
		currentFuryLiquidityThreshold sdk.Uint
		nativeAmount                   sdk.Uint
		nativePrice                    sdk.Dec
		expectedIsBlocked              bool
	}{
		{
			name:                           "not blocked",
			currentFuryLiquidityThreshold: sdk.NewUint(180),
			nativeAmount:                   sdk.NewUint(900),
			nativePrice:                    sdk.MustNewDecFromStr("0.2"),
			expectedIsBlocked:              false,
		},
		{
			name:                           "blocked",
			currentFuryLiquidityThreshold: sdk.NewUint(179),
			nativeAmount:                   sdk.NewUint(900),
			nativePrice:                    sdk.MustNewDecFromStr("0.2"),
			expectedIsBlocked:              true,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := test.CreateTestApp(false)

			liquidityProtectionRateParams := app.ClpKeeper.GetLiquidityProtectionRateParams(ctx)
			liquidityProtectionRateParams.CurrentFuryLiquidityThreshold = tc.currentFuryLiquidityThreshold
			app.ClpKeeper.SetLiquidityProtectionRateParams(ctx, liquidityProtectionRateParams)

			isBlocked := app.ClpKeeper.IsBlockedByLiquidityProtection(ctx, tc.nativeAmount, tc.nativePrice)

			require.Equal(t, tc.expectedIsBlocked, isBlocked)
		})
	}
}

func TestKeeper_MustUpdateLiquidityProtectionThreshold(t *testing.T) {
	testcases := []struct {
		name                           string
		maxFuryLiquidityThreshold     sdk.Uint
		currentFuryLiquidityThreshold sdk.Uint
		isActive                       bool
		nativeAmount                   sdk.Uint
		nativePrice                    sdk.Dec
		sellNative                     bool
		expectedUpdatedThreshold       sdk.Uint
		expectedPanicError             string
	}{
		{
			name:                           "sell native",
			maxFuryLiquidityThreshold:     sdk.NewUint(100000000),
			currentFuryLiquidityThreshold: sdk.NewUint(180),
			isActive:                       true,
			nativeAmount:                   sdk.NewUint(900),
			nativePrice:                    sdk.MustNewDecFromStr("0.2"),
			sellNative:                     true,
			expectedUpdatedThreshold:       sdk.ZeroUint(),
		},
		{
			name:                           "buy native",
			maxFuryLiquidityThreshold:     sdk.NewUint(100000000),
			currentFuryLiquidityThreshold: sdk.NewUint(180),
			isActive:                       true,
			nativeAmount:                   sdk.NewUint(900),
			nativePrice:                    sdk.MustNewDecFromStr("0.2"),
			sellNative:                     false,
			expectedUpdatedThreshold:       sdk.NewUint(360),
		},
		{
			name:                           "liquidity protection disabled",
			maxFuryLiquidityThreshold:     sdk.NewUint(100000000),
			currentFuryLiquidityThreshold: sdk.NewUint(180),
			isActive:                       false,
			expectedUpdatedThreshold:       sdk.NewUint(180),
		},
		{
			name:                           "panics if sell native value greater than current threshold",
			currentFuryLiquidityThreshold: sdk.NewUint(180),
			isActive:                       true,
			nativeAmount:                   sdk.NewUint(900),
			nativePrice:                    sdk.MustNewDecFromStr("1"),
			sellNative:                     true,
			expectedPanicError:             "expect sell native value to be less than currentFuryLiquidityThreshold",
		},
		{
			name:                           "does not exceed max",
			maxFuryLiquidityThreshold:     sdk.NewUint(150),
			currentFuryLiquidityThreshold: sdk.NewUint(100),
			isActive:                       true,
			nativeAmount:                   sdk.NewUint(80),
			nativePrice:                    sdk.MustNewDecFromStr("1"),
			sellNative:                     false,
			expectedUpdatedThreshold:       sdk.NewUint(150),
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := test.CreateTestApp(false)

			liquidityProtectionParams := app.ClpKeeper.GetLiquidityProtectionParams(ctx)
			liquidityProtectionParams.IsActive = tc.isActive
			liquidityProtectionParams.MaxFuryLiquidityThreshold = tc.maxFuryLiquidityThreshold
			app.ClpKeeper.SetLiquidityProtectionParams(ctx, liquidityProtectionParams)

			liquidityProtectionRateParams := app.ClpKeeper.GetLiquidityProtectionRateParams(ctx)
			liquidityProtectionRateParams.CurrentFuryLiquidityThreshold = tc.currentFuryLiquidityThreshold
			app.ClpKeeper.SetLiquidityProtectionRateParams(ctx, liquidityProtectionRateParams)

			if tc.expectedPanicError != "" {
				require.PanicsWithError(t, tc.expectedPanicError, func() {
					app.ClpKeeper.MustUpdateLiquidityProtectionThreshold(ctx, tc.sellNative, tc.nativeAmount, tc.nativePrice)
				})
				return
			}

			app.ClpKeeper.MustUpdateLiquidityProtectionThreshold(ctx, tc.sellNative, tc.nativeAmount, tc.nativePrice)

			liquidityProtectionRateParams = app.ClpKeeper.GetLiquidityProtectionRateParams(ctx)

			require.Equal(t, tc.expectedUpdatedThreshold.String(), liquidityProtectionRateParams.CurrentFuryLiquidityThreshold.String())
		})
	}
}
