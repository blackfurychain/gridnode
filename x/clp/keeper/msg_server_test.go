package keeper_test

import (
	"errors"
	"testing"

	gridapp "github.com/Gridironchain/gridnode/app"
	admintest "github.com/Gridironchain/gridnode/x/admin/test"
	admintypes "github.com/Gridironchain/gridnode/x/admin/types"
	clpkeeper "github.com/Gridironchain/gridnode/x/clp/keeper"
	tokenregistrytypes "github.com/Gridironchain/gridnode/x/tokenregistry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	"github.com/Gridironchain/gridnode/x/clp/test"
	"github.com/Gridironchain/gridnode/x/clp/types"
)

func TestMsgServer_DecommissionPool(t *testing.T) {
	testcases := []struct {
		name                string
		createBalance       bool
		createPool          bool
		createLPs           bool
		poolAsset           string
		address             string
		nativeBalance       sdk.Int
		externalBalance     sdk.Int
		nativeAssetAmount   sdk.Uint
		externalAssetAmount sdk.Uint
		poolUnits           sdk.Uint
		msg                 *types.MsgDecommissionPool
		err                 error
		errString           error
	}{
		{
			name:          "pool does not exist",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			msg: &types.MsgDecommissionPool{
				Signer: "xxx",
				Symbol: "xxx",
			},
			errString: errors.New("pool does not exist"),
		},
		{
			name:                "wrong address",
			createBalance:       false,
			createPool:          true,
			createLPs:           false,
			poolAsset:           "eth",
			address:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:       sdk.NewInt(10000),
			externalBalance:     sdk.NewInt(10000),
			nativeAssetAmount:   sdk.NewUint(1000),
			externalAssetAmount: sdk.NewUint(1000),
			poolUnits:           sdk.NewUint(1000),
			msg: &types.MsgDecommissionPool{
				Signer: "xxx",
				Symbol: "eth",
			},
			errString: errors.New("decoding bech32 failed: invalid bech32 string length 3"),
		},
		{
			name:                "invalid user",
			createBalance:       false,
			createPool:          true,
			createLPs:           false,
			poolAsset:           "eth",
			address:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:       sdk.NewInt(10000),
			externalBalance:     sdk.NewInt(10000),
			nativeAssetAmount:   sdk.NewUint(1000),
			externalAssetAmount: sdk.NewUint(1000),
			poolUnits:           sdk.NewUint(1000),
			msg: &types.MsgDecommissionPool{
				Signer: "did:fury:g1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8wtmssn",
				Symbol: "eth",
			},
			errString: errors.New("user does not have permission to decommission pool: invalid"),
		},
		{
			name:                "balance too high",
			createBalance:       true,
			createPool:          true,
			createLPs:           true,
			poolAsset:           "eth",
			address:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:       sdk.NewInt(10000),
			externalBalance:     sdk.NewInt(10000),
			nativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
			externalAssetAmount: sdk.NewUint(1000),
			poolUnits:           sdk.NewUint(1000),
			msg: &types.MsgDecommissionPool{
				Signer: "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				Symbol: "eth",
			},
			errString: errors.New("Pool Balance too high to be decommissioned"),
		},
		{
			name:                "liquidity provider does not exist",
			createBalance:       true,
			createPool:          true,
			createLPs:           false,
			poolAsset:           "eth",
			address:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:       sdk.NewInt(10000),
			externalBalance:     sdk.NewInt(10000),
			nativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
			externalAssetAmount: sdk.NewUint(1000),
			poolUnits:           sdk.NewUint(1000),
			msg: &types.MsgDecommissionPool{
				Signer: "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				Symbol: "eth",
			},
			errString: errors.New("Pool Balance too high to be decommissioned"),
		},
		{
			name:                "insufficient funds",
			createBalance:       true,
			createPool:          true,
			createLPs:           true,
			poolAsset:           "eth",
			address:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:       sdk.NewInt(10000),
			externalBalance:     sdk.NewInt(10000),
			nativeAssetAmount:   sdk.NewUint(1000),
			externalAssetAmount: sdk.NewUint(1000),
			poolUnits:           sdk.NewUint(1000),
			msg: &types.MsgDecommissionPool{
				Signer: "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				Symbol: "eth",
			},
			errString: errors.New("0eth is smaller than 1000eth: insufficient funds: unable to add balance: Unable to add liquidity provider"),
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *gridapp.GridironchainApp, genesisState gridapp.GenesisState) gridapp.GenesisState {
				trGs := &tokenregistrytypes.GenesisState{
					Registry: &tokenregistrytypes.Registry{
						Entries: []*tokenregistrytypes.RegistryEntry{
							{Denom: tc.poolAsset, BaseDenom: tc.poolAsset, Decimals: 18, Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP}},
							{Denom: "fury", BaseDenom: "fury", Decimals: 18, Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP}},
						},
					},
				}
				bz, _ := app.AppCodec().MarshalJSON(trGs)
				genesisState["tokenregistry"] = bz

				if tc.createBalance {
					balances := []banktypes.Balance{
						{
							Address: tc.address,
							Coins: sdk.Coins{
								sdk.NewCoin(tc.poolAsset, tc.externalBalance),
								sdk.NewCoin("fury", tc.nativeBalance),
							},
						},
					}

					bankGs := banktypes.DefaultGenesisState()
					bankGs.Balances = append(bankGs.Balances, balances...)
					bz, _ = app.AppCodec().MarshalJSON(bankGs)
					genesisState["bank"] = bz
				}

				if tc.createPool {
					pools := []*types.Pool{
						{
							ExternalAsset:        &types.Asset{Symbol: tc.poolAsset},
							NativeAssetBalance:   tc.nativeAssetAmount,
							ExternalAssetBalance: tc.externalAssetAmount,
							PoolUnits:            tc.poolUnits,
						},
					}
					clpGs := types.DefaultGenesisState()
					if tc.createLPs {
						lps := []*types.LiquidityProvider{
							{
								Asset:                    &types.Asset{Symbol: tc.poolAsset},
								LiquidityProviderAddress: tc.address,
								LiquidityProviderUnits:   tc.nativeAssetAmount,
							},
						}
						clpGs.LiquidityProviders = append(clpGs.LiquidityProviders, lps...)
					}
					clpGs.Params = types.Params{
						MinCreatePoolThreshold: 100,
					}
					clpGs.AddressWhitelist = append(clpGs.AddressWhitelist, tc.address)
					clpGs.PoolList = append(clpGs.PoolList, pools...)
					bz, _ = app.AppCodec().MarshalJSON(clpGs)
					genesisState["clp"] = bz
				}

				return genesisState
			})

			msgServer := clpkeeper.NewMsgServerImpl(app.ClpKeeper)

			_, err := msgServer.DecommissionPool(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.errString != nil {
				require.EqualError(t, err, tc.errString.Error())
				return
			}
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgServer_Swap(t *testing.T) {
	testcases := []struct {
		name                            string
		createBalance                   bool
		createPool                      bool
		createLPs                       bool
		poolAsset                       string
		decimals                        int64
		address                         string
		maxFuryLiquidityThresholdAsset string
		swapFeeParams                   types.SwapFeeParams
		nativeBalance                   sdk.Int
		externalBalance                 sdk.Int
		nativeAssetAmount               sdk.Uint
		externalAssetAmount             sdk.Uint
		poolUnits                       sdk.Uint
		currentFuryLiquidityThreshold  sdk.Uint
		expectedRunningThresholdEnd     sdk.Uint
		maxFuryLiquidityThreshold      sdk.Uint
		nativeBalanceEnd                sdk.Int
		externalBalanceEnd              sdk.Int
		poolAssetPermissions            []tokenregistrytypes.Permission
		nativeAssetPermissions          []tokenregistrytypes.Permission
		msg                             *types.MsgSwap
		err                             error
		errString                       error
	}{
		{
			name:                            "sent asset token not supported",
			createBalance:                   false,
			createPool:                      false,
			createLPs:                       false,
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			maxFuryLiquidityThresholdAsset: "fury",
			msg: &types.MsgSwap{
				Signer:             "xxx",
				SentAsset:          &types.Asset{Symbol: "xxx"},
				ReceivedAsset:      &types.Asset{Symbol: "xxx"},
				SentAmount:         sdk.NewUint(1),
				MinReceivingAmount: sdk.NewUint(1),
			},
			errString: errors.New("Token not supported by gridchain"),
		},
		{
			name:                            "received asset token not supported",
			createBalance:                   false,
			createPool:                      false,
			createLPs:                       false,
			poolAsset:                       "eth",
			decimals:                        18,
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			maxFuryLiquidityThresholdAsset: "fury",
			msg: &types.MsgSwap{
				Signer:             "xxx",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "xxx"},
				SentAmount:         sdk.NewUint(1),
				MinReceivingAmount: sdk.NewUint(1),
			},
			errString: errors.New("Token not supported by gridchain"),
		},
		{
			name:                            "external asset permission denied",
			createBalance:                   false,
			createPool:                      false,
			createLPs:                       false,
			poolAsset:                       "eth",
			decimals:                        18,
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			maxFuryLiquidityThresholdAsset: "fury",
			msg: &types.MsgSwap{
				Signer:             "xxx",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(1),
				MinReceivingAmount: sdk.NewUint(1),
			},
			errString: errors.New("permission denied for denom"),
		},
		{
			name:                            "native asset permission denied",
			createBalance:                   false,
			createPool:                      false,
			createLPs:                       false,
			poolAsset:                       "eth",
			decimals:                        18,
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			maxFuryLiquidityThresholdAsset: "fury",
			msg: &types.MsgSwap{
				Signer:             "xxx",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(1),
				MinReceivingAmount: sdk.NewUint(1),
			},
			errString: errors.New("permission denied for denom"),
		},
		{
			name:                            "received amount below expected",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			decimals:                        18,
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(1000),
			externalAssetAmount:             sdk.NewUint(1000),
			poolUnits:                       sdk.NewUint(1000),
			nativeBalanceEnd:                sdk.NewInt(10000),
			externalBalanceEnd:              sdk.NewInt(10000),
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			maxFuryLiquidityThresholdAsset: "fury",
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(1),
				MinReceivingAmount: sdk.NewUint(1),
			},
			errString: errors.New("Unable to swap, received amount is below expected"),
		},
		{
			name:                            "success",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			decimals:                        18,
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(1000),
			externalAssetAmount:             sdk.NewUint(1000),
			poolUnits:                       sdk.NewUint(1000),
			nativeBalanceEnd:                sdk.NewInt(10090),
			externalBalanceEnd:              sdk.NewInt(9900),
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			expectedRunningThresholdEnd:     sdk.NewUint(1090),
			maxFuryLiquidityThresholdAsset: "fury",
			maxFuryLiquidityThreshold:      sdk.NewUint(2000),
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(100),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
		{
			name:                            "sell native token over threshold",
			poolAsset:                       "eth",
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			maxFuryLiquidityThresholdAsset: "fury",
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "fury"},
				ReceivedAsset:      &types.Asset{Symbol: "eth"},
				SentAmount:         sdk.NewUintFromString("10000000000000000000000"),
				MinReceivingAmount: sdk.NewUint(1),
			},
			errString: errors.New("Unable to swap, reached maximum fury liquidity threshold"),
		},
		{
			name:                            "sell native token over threshold - zero threshold",
			poolAsset:                       "eth",
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(0),
			maxFuryLiquidityThresholdAsset: "fury",
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "fury"},
				ReceivedAsset:      &types.Asset{Symbol: "eth"},
				SentAmount:         sdk.NewUintFromString("10000000000000000000000"),
				MinReceivingAmount: sdk.NewUint(1),
			},
			errString: errors.New("Unable to swap, reached maximum fury liquidity threshold"),
		},
		{
			name:                            "sell native token just over threshold",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(100000),
			externalAssetAmount:             sdk.NewUint(200000),
			poolUnits:                       sdk.NewUint(100000),
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			maxFuryLiquidityThresholdAsset: "eth",
			externalBalanceEnd:              sdk.NewInt(10500),
			nativeBalanceEnd:                sdk.NewInt(9749),
			expectedRunningThresholdEnd:     sdk.NewUint(498),
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "fury"},
				ReceivedAsset:      &types.Asset{Symbol: "eth"},
				SentAmount:         sdk.NewUint(251),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
		{
			name:                            "sell native token just below threshold",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(100000),
			externalAssetAmount:             sdk.NewUint(200000),
			poolUnits:                       sdk.NewUint(100000),
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(4000),
			expectedRunningThresholdEnd:     sdk.NewUint(3500),
			externalBalanceEnd:              sdk.NewInt(10498),
			nativeBalanceEnd:                sdk.NewInt(9750),
			maxFuryLiquidityThresholdAsset: "eth",
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "fury"},
				ReceivedAsset:      &types.Asset{Symbol: "eth"},
				SentAmount:         sdk.NewUint(250),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
		{
			name:                            "buy fury, threshold in eth",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(100000),
			externalAssetAmount:             sdk.NewUint(200000),
			poolUnits:                       sdk.NewUint(100000),
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(200),
			expectedRunningThresholdEnd:     sdk.NewUint(298),
			maxFuryLiquidityThresholdAsset: "eth",
			maxFuryLiquidityThreshold:      sdk.NewUint(1000),
			externalBalanceEnd:              sdk.NewInt(9900),
			nativeBalanceEnd:                sdk.NewInt(10049),
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(100),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
		{
			name:                            "buy fury, threshold in eth, low max liquidity threshold",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(100000),
			externalAssetAmount:             sdk.NewUint(200000),
			poolUnits:                       sdk.NewUint(100000),
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(200),
			expectedRunningThresholdEnd:     sdk.NewUint(250),
			maxFuryLiquidityThresholdAsset: "eth",
			maxFuryLiquidityThreshold:      sdk.NewUint(250),
			externalBalanceEnd:              sdk.NewInt(9900),
			nativeBalanceEnd:                sdk.NewInt(10049),
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(100),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
		{
			name:                            "buy fury, threshold in eth, maximum max liquidity threshold",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(100000),
			externalAssetAmount:             sdk.NewUint(200000),
			poolUnits:                       sdk.NewUint(100000),
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUintFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935"),
			expectedRunningThresholdEnd:     sdk.NewUintFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935"),
			maxFuryLiquidityThresholdAsset: "eth",
			maxFuryLiquidityThreshold:      sdk.NewUintFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935"),
			externalBalanceEnd:              sdk.NewInt(9900),
			nativeBalanceEnd:                sdk.NewInt(10049),
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(100),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
		{
			name:                            "correct amount when external asset has eighteen decimals",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "decvar",
			decimals:                        18,
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:                   sdk.NewInt(10),
			externalBalance:                 sdk.NewInt(10),
			nativeAssetAmount:               sdk.NewUint(400),
			externalAssetAmount:             sdk.NewUint(100),
			poolUnits:                       sdk.NewUint(100),
			nativeBalanceEnd:                sdk.NewInt(13),
			externalBalanceEnd:              sdk.NewInt(9),
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			expectedRunningThresholdEnd:     sdk.NewUint(1003),
			maxFuryLiquidityThresholdAsset: "fury",
			maxFuryLiquidityThreshold:      sdk.NewUint(2000),
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "decvar"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(1),
				MinReceivingAmount: sdk.NewUint(0),
			},
		},
		{
			name:                            "correct amount when external asset has nineteen decimals",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "decvar",
			decimals:                        19,
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:                   sdk.NewInt(10),
			externalBalance:                 sdk.NewInt(10),
			nativeAssetAmount:               sdk.NewUint(400),
			externalAssetAmount:             sdk.NewUint(100),
			poolUnits:                       sdk.NewUint(100),
			nativeBalanceEnd:                sdk.NewInt(13),
			externalBalanceEnd:              sdk.NewInt(9),
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			expectedRunningThresholdEnd:     sdk.NewUint(1003),
			maxFuryLiquidityThresholdAsset: "fury",
			maxFuryLiquidityThreshold:      sdk.NewUint(2000),
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "decvar"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(1),
				MinReceivingAmount: sdk.NewUint(0),
			},
		},
		{
			name:                            "eth:fury - swap fee 5%",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			decimals:                        18,
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(1000),
			externalAssetAmount:             sdk.NewUint(1000),
			poolUnits:                       sdk.NewUint(1000),
			nativeBalanceEnd:                sdk.NewInt(10086),
			externalBalanceEnd:              sdk.NewInt(9900),
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			expectedRunningThresholdEnd:     sdk.NewUint(1086),
			maxFuryLiquidityThresholdAsset: "fury",
			maxFuryLiquidityThreshold:      sdk.NewUint(2000),
			swapFeeParams: types.SwapFeeParams{
				DefaultSwapFeeRate: sdk.NewDecWithPrec(3, 3),
				TokenParams: []*types.SwapFeeTokenParams{
					{
						Asset:       "fury",
						SwapFeeRate: sdk.NewDecWithPrec(1, 1), //10%
					},
					{
						Asset:       "eth",
						SwapFeeRate: sdk.NewDecWithPrec(5, 2), //5%
					},
					{
						Asset:       "usdc",
						SwapFeeRate: sdk.NewDecWithPrec(2, 2), //2%
					},
				},
			},
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "eth"},
				ReceivedAsset:      &types.Asset{Symbol: "fury"},
				SentAmount:         sdk.NewUint(100),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
		{
			name:                            "fury:eth - swap fee 10%",
			createBalance:                   true,
			createPool:                      true,
			createLPs:                       true,
			poolAsset:                       "eth",
			decimals:                        18,
			address:                         "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:                   sdk.NewInt(10000),
			externalBalance:                 sdk.NewInt(10000),
			nativeAssetAmount:               sdk.NewUint(1000),
			externalAssetAmount:             sdk.NewUint(1000),
			poolUnits:                       sdk.NewUint(1000),
			nativeBalanceEnd:                sdk.NewInt(9900),
			externalBalanceEnd:              sdk.NewInt(10081),
			poolAssetPermissions:            []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			currentFuryLiquidityThreshold:  sdk.NewUint(1000),
			expectedRunningThresholdEnd:     sdk.NewUint(910),
			maxFuryLiquidityThresholdAsset: "fury",
			maxFuryLiquidityThreshold:      sdk.NewUint(2000),
			swapFeeParams: types.SwapFeeParams{
				DefaultSwapFeeRate: sdk.NewDecWithPrec(3, 3),
				TokenParams: []*types.SwapFeeTokenParams{
					{
						Asset:       "fury",
						SwapFeeRate: sdk.NewDecWithPrec(1, 1), //10%
					},
					{
						Asset:       "eth",
						SwapFeeRate: sdk.NewDecWithPrec(5, 2), //5%
					},
					{
						Asset:       "usdc",
						SwapFeeRate: sdk.NewDecWithPrec(2, 2), //2%
					},
				},
			},
			msg: &types.MsgSwap{
				Signer:             "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				SentAsset:          &types.Asset{Symbol: "fury"},
				ReceivedAsset:      &types.Asset{Symbol: "eth"},
				SentAmount:         sdk.NewUint(100),
				MinReceivingAmount: sdk.NewUint(1),
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *gridapp.GridironchainApp, genesisState gridapp.GenesisState) gridapp.GenesisState {

				trGs := &tokenregistrytypes.GenesisState{
					Registry: &tokenregistrytypes.Registry{
						Entries: []*tokenregistrytypes.RegistryEntry{
							{Denom: tc.poolAsset, BaseDenom: tc.poolAsset, Decimals: tc.decimals, Permissions: tc.poolAssetPermissions},
							{Denom: "fury", BaseDenom: "fury", Decimals: 18, Permissions: tc.nativeAssetPermissions},
						},
					},
				}
				bz, _ := app.AppCodec().MarshalJSON(trGs)
				genesisState["tokenregistry"] = bz

				if tc.createBalance {
					externalCLP, _ := sdk.NewIntFromString(tc.externalAssetAmount.String())
					nativeCLP, _ := sdk.NewIntFromString(tc.nativeAssetAmount.String())
					clpAddrs := app.AccountKeeper.GetModuleAddress("clp").String()

					balances := []banktypes.Balance{
						{
							Address: tc.address,
							Coins: sdk.Coins{
								sdk.NewCoin(tc.poolAsset, tc.externalBalance),
								sdk.NewCoin("fury", tc.nativeBalance),
							},
						},
						{
							Address: clpAddrs,
							Coins: sdk.Coins{
								sdk.NewCoin(tc.poolAsset, externalCLP),
								sdk.NewCoin("fury", nativeCLP),
							},
						},
					}

					bankGs := banktypes.DefaultGenesisState()
					bankGs.Balances = append(bankGs.Balances, balances...)
					bz, _ = app.AppCodec().MarshalJSON(bankGs)
					genesisState["bank"] = bz
				}

				if tc.createPool {
					pools := []*types.Pool{
						{
							ExternalAsset:        &types.Asset{Symbol: tc.poolAsset},
							NativeAssetBalance:   tc.nativeAssetAmount,
							ExternalAssetBalance: tc.externalAssetAmount,
							PoolUnits:            tc.poolUnits,
						},
					}
					clpGs := types.DefaultGenesisState()
					if tc.createLPs {
						lps := []*types.LiquidityProvider{
							{
								Asset:                    &types.Asset{Symbol: tc.poolAsset},
								LiquidityProviderAddress: tc.address,
								LiquidityProviderUnits:   tc.poolUnits,
							},
						}
						clpGs.LiquidityProviders = append(clpGs.LiquidityProviders, lps...)
					}
					clpGs.Params = types.Params{
						MinCreatePoolThreshold: 100,
					}
					clpGs.AddressWhitelist = append(clpGs.AddressWhitelist, tc.address)
					clpGs.PoolList = append(clpGs.PoolList, pools...)
					bz, _ = app.AppCodec().MarshalJSON(clpGs)
					genesisState["clp"] = bz
				}

				return genesisState
			})

			app.ClpKeeper.SetPmtpCurrentRunningRate(ctx, sdk.NewDec(0))
			app.ClpKeeper.SetSwapFeeParams(ctx, &tc.swapFeeParams)

			liquidityProtectionParam := app.ClpKeeper.GetLiquidityProtectionParams(ctx)
			liquidityProtectionParam.MaxFuryLiquidityThresholdAsset = tc.maxFuryLiquidityThresholdAsset
			liquidityProtectionParam.MaxFuryLiquidityThreshold = tc.maxFuryLiquidityThreshold
			liquidityProtectionParam.IsActive = true
			app.ClpKeeper.SetLiquidityProtectionParams(ctx, liquidityProtectionParam)
			app.ClpKeeper.SetLiquidityProtectionCurrentFuryLiquidityThreshold(ctx, tc.currentFuryLiquidityThreshold)

			msgServer := clpkeeper.NewMsgServerImpl(app.ClpKeeper)

			_, err := msgServer.Swap(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.errString != nil {
				require.EqualError(t, err, tc.errString.Error())
				return
			}
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			externalAssetBalanceRequest := banktypes.QueryBalanceRequest{
				Address: tc.address,
				Denom:   tc.poolAsset,
			}

			externalAssetBalanceResponse, err := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &externalAssetBalanceRequest)
			if err != nil {
				t.Fatalf("err: %v", err)
			}

			externalAssetBalance := externalAssetBalanceResponse.Balance.Amount

			require.Equal(t, tc.externalBalanceEnd.String(), externalAssetBalance.String())

			nativeAssetBalanceRequest := banktypes.QueryBalanceRequest{
				Address: tc.address,
				Denom:   "fury",
			}

			nativeAssetBalanceResponse, err := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &nativeAssetBalanceRequest)
			if err != nil {
				t.Fatalf("err: %v", err)
			}

			nativeAssetBalance := nativeAssetBalanceResponse.Balance.Amount

			require.Equal(t, tc.nativeBalanceEnd.String(), nativeAssetBalance.String())

			runningThreshold := app.ClpKeeper.GetLiquidityProtectionRateParams(ctx).CurrentFuryLiquidityThreshold
			require.Equal(t, tc.expectedRunningThresholdEnd.String(), runningThreshold.String())

		})
	}
}

func TestMsgServer_RemoveLiquidity(t *testing.T) {
	testcases := []struct {
		name                   string
		createBalance          bool
		createPool             bool
		createLPs              bool
		poolAsset              string
		address                string
		nativeBalance          sdk.Int
		externalBalance        sdk.Int
		nativeAssetAmount      sdk.Uint
		externalAssetAmount    sdk.Uint
		poolUnits              sdk.Uint
		poolAssetPermissions   []tokenregistrytypes.Permission
		nativeAssetPermissions []tokenregistrytypes.Permission
		msg                    *types.MsgRemoveLiquidity
		err                    error
		errString              error
	}{
		{
			name:          "sent asset token not supported",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			msg: &types.MsgRemoveLiquidity{
				Signer:        "xxx",
				ExternalAsset: &types.Asset{Symbol: "xxx"},
				WBasisPoints:  sdk.NewInt(1),
				Asymmetry:     sdk.NewInt(1),
			},
			errString: errors.New("Token not supported by gridchain"),
		},
		{
			name:          "received asset token not supported",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			poolAsset:     "eth",
			msg: &types.MsgRemoveLiquidity{
				Signer:        "xxx",
				ExternalAsset: &types.Asset{Symbol: "xxx"},
				WBasisPoints:  sdk.NewInt(1),
				Asymmetry:     sdk.NewInt(1),
			},
			errString: errors.New("Token not supported by gridchain"),
		},
		{
			name:          "external asset permission denied",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			poolAsset:     "eth",
			msg: &types.MsgRemoveLiquidity{
				Signer:        "xxx",
				ExternalAsset: &types.Asset{Symbol: "eth"},
				WBasisPoints:  sdk.NewInt(1),
				Asymmetry:     sdk.NewInt(1),
			},
			errString: errors.New("permission denied for denom"),
		},
		{
			name:                 "pool does not exist",
			createBalance:        false,
			createPool:           false,
			createLPs:            false,
			poolAsset:            "xxx",
			poolAssetPermissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgRemoveLiquidity{
				Signer:        "xxx",
				ExternalAsset: &types.Asset{Symbol: "xxx"},
				WBasisPoints:  sdk.NewInt(1),
				Asymmetry:     sdk.NewInt(1),
			},
			errString: errors.New("pool does not exist"),
		},
		{
			name:                 "no lp",
			createBalance:        true,
			createPool:           true,
			createLPs:            false,
			poolAsset:            "eth",
			address:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:        sdk.NewInt(10000),
			externalBalance:      sdk.NewInt(10000),
			nativeAssetAmount:    sdk.NewUint(1000),
			externalAssetAmount:  sdk.NewUint(1000),
			poolUnits:            sdk.NewUint(1000),
			poolAssetPermissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgRemoveLiquidity{
				Signer:        "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset: &types.Asset{Symbol: "eth"},
				WBasisPoints:  sdk.NewInt(1),
				Asymmetry:     sdk.NewInt(1),
			},
			errString: errors.New("liquidity Provider does not exist"),
		},
		{
			name:                   "received amount below expected",
			createBalance:          true,
			createPool:             true,
			createLPs:              true,
			poolAsset:              "eth",
			address:                "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(1000),
			externalAssetAmount:    sdk.NewUint(1000),
			poolUnits:              sdk.NewUint(1000),
			poolAssetPermissions:   []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgRemoveLiquidity{
				Signer:        "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset: &types.Asset{Symbol: "eth"},
				WBasisPoints:  sdk.NewInt(1),
				Asymmetry:     sdk.NewInt(0),
			},
		},
		{
			name:                   "received amount below expected",
			createBalance:          true,
			createPool:             true,
			createLPs:              true,
			poolAsset:              "eth",
			address:                "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(1000),
			externalAssetAmount:    sdk.NewUint(1000),
			poolUnits:              sdk.NewUint(1000),
			poolAssetPermissions:   []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgRemoveLiquidity{
				Signer:        "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset: &types.Asset{Symbol: "eth"},
				WBasisPoints:  sdk.NewInt(1),
				Asymmetry:     sdk.NewInt(1),
			},
			err: types.ErrAsymmetricRemove,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *gridapp.GridironchainApp, genesisState gridapp.GenesisState) gridapp.GenesisState {
				trGs := &tokenregistrytypes.GenesisState{
					Registry: &tokenregistrytypes.Registry{
						Entries: []*tokenregistrytypes.RegistryEntry{
							{Denom: tc.poolAsset, BaseDenom: tc.poolAsset, Decimals: 18, Permissions: tc.poolAssetPermissions},
							{Denom: "fury", BaseDenom: "fury", Decimals: 18, Permissions: tc.nativeAssetPermissions},
						},
					},
				}
				bz, _ := app.AppCodec().MarshalJSON(trGs)
				genesisState["tokenregistry"] = bz

				if tc.createBalance {
					balances := []banktypes.Balance{
						{
							Address: "did:fury:g1pjm228rsgwqf23arkx7lm9ypkyma7mzrad6f3n",
							Coins: sdk.Coins{
								sdk.NewCoin(tc.poolAsset, sdk.Int(tc.externalAssetAmount)),
								sdk.NewCoin("fury", sdk.Int(tc.nativeAssetAmount)),
							},
						},
						{
							Address: tc.address,
							Coins: sdk.Coins{
								sdk.NewCoin(tc.poolAsset, tc.externalBalance),
								sdk.NewCoin("fury", tc.nativeBalance),
							},
						},
					}

					bankGs := banktypes.DefaultGenesisState()
					bankGs.Balances = append(bankGs.Balances, balances...)
					bz, _ = app.AppCodec().MarshalJSON(bankGs)
					genesisState["bank"] = bz
				}

				if tc.createPool {
					pools := []*types.Pool{
						{
							ExternalAsset:        &types.Asset{Symbol: tc.poolAsset},
							NativeAssetBalance:   tc.nativeAssetAmount,
							ExternalAssetBalance: tc.externalAssetAmount,
							PoolUnits:            tc.poolUnits,
						},
					}
					clpGs := types.DefaultGenesisState()
					if tc.createLPs {
						lps := []*types.LiquidityProvider{
							{
								Asset:                    &types.Asset{Symbol: tc.poolAsset},
								LiquidityProviderAddress: tc.address,
								LiquidityProviderUnits:   tc.nativeAssetAmount,
							},
						}
						clpGs.LiquidityProviders = append(clpGs.LiquidityProviders, lps...)
					}
					clpGs.Params = types.Params{
						MinCreatePoolThreshold: 100,
					}
					clpGs.AddressWhitelist = append(clpGs.AddressWhitelist, tc.address)
					clpGs.PoolList = append(clpGs.PoolList, pools...)
					bz, _ = app.AppCodec().MarshalJSON(clpGs)
					genesisState["clp"] = bz
				}

				return genesisState
			})

			app.ClpKeeper.SetPmtpCurrentRunningRate(ctx, sdk.NewDec(1))

			msgServer := clpkeeper.NewMsgServerImpl(app.ClpKeeper)

			_, err := msgServer.RemoveLiquidity(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.errString != nil {
				require.EqualError(t, err, tc.errString.Error())
				return
			}
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgServer_CreatePool(t *testing.T) {
	testcases := []struct {
		name                   string
		createBalance          bool
		createPool             bool
		createLPs              bool
		poolAsset              string
		address                string
		nativeBalance          sdk.Int
		externalBalance        sdk.Int
		nativeAssetAmount      sdk.Uint
		externalAssetAmount    sdk.Uint
		poolUnits              sdk.Uint
		poolAssetPermissions   []tokenregistrytypes.Permission
		nativeAssetPermissions []tokenregistrytypes.Permission
		msg                    *types.MsgCreatePool
		err                    error
		errString              error
	}{
		{
			name:          "total amount too low",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			msg: &types.MsgCreatePool{
				Signer:              "xxx",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUint(1000),
				ExternalAssetAmount: sdk.NewUint(1000),
			},
			errString: errors.New("total amount is less than minimum threshold"),
		},
		{
			name:          "external asset token not supported",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			msg: &types.MsgCreatePool{
				Signer:              "xxx",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("Token not supported by gridchain"),
		},
		{
			name:          "external asset permission denied",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			poolAsset:     "eth",
			msg: &types.MsgCreatePool{
				Signer:              "xxx",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("permission denied for denom"),
		},
		{
			name:                 "pool already exists",
			createBalance:        true,
			createPool:           true,
			createLPs:            true,
			poolAsset:            "eth",
			address:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:        sdk.NewInt(10000),
			externalBalance:      sdk.NewInt(10000),
			nativeAssetAmount:    sdk.NewUint(1000),
			externalAssetAmount:  sdk.NewUint(1000),
			poolUnits:            sdk.NewUint(1000),
			poolAssetPermissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgCreatePool{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("Unable to create pool"),
		},
		{
			name:                 "user does have enough balance of required coin",
			createBalance:        true,
			createPool:           false,
			createLPs:            false,
			poolAsset:            "eth",
			address:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:        sdk.NewInt(10000),
			externalBalance:      sdk.NewInt(10000),
			nativeAssetAmount:    sdk.NewUint(1000),
			externalAssetAmount:  sdk.NewUint(1000),
			poolUnits:            sdk.NewUint(1000),
			poolAssetPermissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgCreatePool{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("user does not have enough balance of the required coin: Unable to set pool"),
		},
		{
			name:                 "successful",
			createBalance:        true,
			createPool:           false,
			createLPs:            false,
			poolAsset:            "eth",
			address:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			nativeBalance:        sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			externalBalance:      sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			nativeAssetAmount:    sdk.NewUint(1000),
			externalAssetAmount:  sdk.NewUint(1000),
			poolUnits:            sdk.NewUint(1000),
			poolAssetPermissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgCreatePool{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *gridapp.GridironchainApp, genesisState gridapp.GenesisState) gridapp.GenesisState {
				trGs := &tokenregistrytypes.GenesisState{
					Registry: &tokenregistrytypes.Registry{
						Entries: []*tokenregistrytypes.RegistryEntry{
							{Denom: tc.poolAsset, BaseDenom: tc.poolAsset, Decimals: 18, Permissions: tc.poolAssetPermissions},
							{Denom: "fury", BaseDenom: "fury", Decimals: 18, Permissions: tc.nativeAssetPermissions},
						},
					},
				}
				bz, _ := app.AppCodec().MarshalJSON(trGs)
				genesisState["tokenregistry"] = bz

				if tc.createBalance {
					balances := []banktypes.Balance{
						{
							Address: tc.address,
							Coins: sdk.Coins{
								sdk.NewCoin(tc.poolAsset, tc.externalBalance),
								sdk.NewCoin("fury", tc.nativeBalance),
							},
						},
					}

					bankGs := banktypes.DefaultGenesisState()
					bankGs.Balances = append(bankGs.Balances, balances...)
					bz, _ = app.AppCodec().MarshalJSON(bankGs)
					genesisState["bank"] = bz
				}

				if tc.createPool {
					pools := []*types.Pool{
						{
							ExternalAsset:        &types.Asset{Symbol: tc.poolAsset},
							NativeAssetBalance:   tc.nativeAssetAmount,
							ExternalAssetBalance: tc.externalAssetAmount,
							PoolUnits:            tc.poolUnits,
						},
					}
					clpGs := types.DefaultGenesisState()
					if tc.createLPs {
						lps := []*types.LiquidityProvider{
							{
								Asset:                    &types.Asset{Symbol: tc.poolAsset},
								LiquidityProviderAddress: tc.address,
								LiquidityProviderUnits:   tc.nativeAssetAmount,
							},
						}
						clpGs.LiquidityProviders = append(clpGs.LiquidityProviders, lps...)
					}
					clpGs.Params = types.Params{
						MinCreatePoolThreshold: 100,
					}
					clpGs.AddressWhitelist = append(clpGs.AddressWhitelist, tc.address)
					clpGs.PoolList = append(clpGs.PoolList, pools...)
					bz, _ = app.AppCodec().MarshalJSON(clpGs)
					genesisState["clp"] = bz
				}

				return genesisState
			})

			app.ClpKeeper.SetPmtpCurrentRunningRate(ctx, sdk.NewDec(1))

			msgServer := clpkeeper.NewMsgServerImpl(app.ClpKeeper)

			_, err := msgServer.CreatePool(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.errString != nil {
				require.EqualError(t, err, tc.errString.Error())
				return
			}
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgServer_AddLiquidity(t *testing.T) {
	testcases := []struct {
		name                                   string
		createBalance                          bool
		createPool                             bool
		createLPs                              bool
		poolAsset                              string
		address                                string
		userNativeAssetBalance                 sdk.Int
		userExternalAssetBalance               sdk.Int
		poolNativeAssetBalance                 sdk.Uint
		poolExternalAssetBalance               sdk.Uint
		poolNativeLiabilities                  sdk.Uint
		poolExternalLiabilities                sdk.Uint
		poolUnits                              sdk.Uint
		poolAssetPermissions                   []tokenregistrytypes.Permission
		nativeAssetPermissions                 []tokenregistrytypes.Permission
		msg                                    *types.MsgAddLiquidity
		liquidityProtectionActive              bool
		maxFuryLiquidityThreshold             sdk.Uint
		currentFuryLiquidityThreshold         sdk.Uint
		expectedPoolUnits                      sdk.Uint
		expectedLPUnits                        sdk.Uint
		err                                    error
		errString                              error
		expectedUpdatedFuryLiquidityThreshold sdk.Uint
	}{
		{
			name:          "external asset token not supported",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			msg: &types.MsgAddLiquidity{
				Signer:              "xxx",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("Token not supported by gridchain"),
		},
		{
			name:          "external asset permission denied",
			createBalance: false,
			createPool:    false,
			createLPs:     false,
			poolAsset:     "eth",
			msg: &types.MsgAddLiquidity{
				Signer:              "xxx",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("permission denied for denom"),
		},
		{
			name:                     "pool does not exist",
			createBalance:            true,
			createPool:               false,
			createLPs:                false,
			poolAsset:                "eth",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.NewInt(10000),
			userExternalAssetBalance: sdk.NewInt(10000),
			poolNativeAssetBalance:   sdk.NewUint(1000),
			poolExternalAssetBalance: sdk.NewUint(1000),
			poolUnits:                sdk.NewUint(1000),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("pool does not exist"),
		},
		{
			name:                     "user does have enough balance of required coin",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "eth",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.NewInt(10000),
			userExternalAssetBalance: sdk.NewInt(10000),
			poolNativeAssetBalance:   sdk.NewUint(1000),
			poolExternalAssetBalance: sdk.NewUint(1000),
			poolUnits:                sdk.NewUint(1000),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			errString: errors.New("user does not have enough balance of the required coin: Unable to add liquidity"),
		},
		{
			name:                     "one side of pool empty - zero amount of external asset added",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "eth",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			userExternalAssetBalance: sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			poolNativeAssetBalance:   sdk.ZeroUint(),
			poolExternalAssetBalance: sdk.NewUint(123),
			poolUnits:                sdk.NewUint(1000),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUint(100),
				ExternalAssetAmount: sdk.ZeroUint(),
			},
			errString: errors.New("amount is invalid"),
		},
		{
			name:                     "one side of pool empty - external and native asset added",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "eth",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			userExternalAssetBalance: sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			poolNativeAssetBalance:   sdk.ZeroUint(),
			poolExternalAssetBalance: sdk.NewUint(123),
			poolUnits:                sdk.NewUint(1000),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUint(178),
				ExternalAssetAmount: sdk.NewUint(156),
			},
			liquidityProtectionActive:              false,
			expectedUpdatedFuryLiquidityThreshold: sdk.ZeroUint(),
			expectedPoolUnits:                      sdk.NewUint(178),
			expectedLPUnits:                        sdk.NewUint(178),
		},
		{
			name:                     "success - symmetric",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "eth",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			userExternalAssetBalance: sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			poolNativeAssetBalance:   sdk.NewUint(1000),
			poolExternalAssetBalance: sdk.NewUint(1000),
			poolUnits:                sdk.NewUint(1000),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			liquidityProtectionActive:              false,
			expectedUpdatedFuryLiquidityThreshold: sdk.ZeroUint(),
			expectedPoolUnits:                      sdk.NewUintFromString("1000000000000001000"),
			expectedLPUnits:                        sdk.NewUintFromString("1000000000000000000"),
		},
		{
			name:                     "success - nearly symmetric",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(4000000000000000000),
				ExternalAssetAmount: sdk.NewUint(68140),
			},
			liquidityProtectionActive:              false,
			expectedUpdatedFuryLiquidityThreshold: sdk.ZeroUint(),
			expectedPoolUnits:                      sdk.NewUintFromString("23662661153298862513590992"),
			expectedLPUnits:                        sdk.NewUintFromString("602841478820653038"),
		},
		{
			name:                     "success - swap external",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(0),
				ExternalAssetAmount: sdk.NewUint(68140),
			},
			liquidityProtectionActive:              false,
			expectedUpdatedFuryLiquidityThreshold: sdk.ZeroUint(),
			expectedPoolUnits:                      sdk.NewUintFromString("23662660751003435747009552"),
			expectedLPUnits:                        sdk.NewUintFromString("200546052054071598"),
		},
		{
			name:                     "success - swap native",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(4000000000000000000),
				ExternalAssetAmount: sdk.ZeroUint(),
			},
			liquidityProtectionActive:              false,
			expectedUpdatedFuryLiquidityThreshold: sdk.ZeroUint(),
			expectedPoolUnits:                      sdk.NewUintFromString("23662660951949037742990437"),
			expectedLPUnits:                        sdk.NewUintFromString("401491654050052483"),
		},
		{
			name:                     "success - swap native - with liabilities",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.ZeroUint(),
			poolExternalAssetBalance: sdk.ZeroUint(),
			poolNativeLiabilities:    sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalLiabilities:  sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(4000000000000000000),
				ExternalAssetAmount: sdk.ZeroUint(),
			},
			liquidityProtectionActive:              false,
			expectedUpdatedFuryLiquidityThreshold: sdk.ZeroUint(),
			expectedPoolUnits:                      sdk.NewUintFromString("23662660951949037742990437"),
			expectedLPUnits:                        sdk.NewUintFromString("401491654050052483"),
		},
		{
			name:                     "success - symmetric - liquidity protection enabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "eth",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			userExternalAssetBalance: sdk.Int(sdk.NewUintFromString(types.PoolThrehold)),
			poolNativeAssetBalance:   sdk.NewUint(1000),
			poolExternalAssetBalance: sdk.NewUint(1000),
			poolUnits:                sdk.NewUint(1000),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "eth"},
				NativeAssetAmount:   sdk.NewUintFromString(types.PoolThrehold),
				ExternalAssetAmount: sdk.NewUintFromString(types.PoolThrehold),
			},
			liquidityProtectionActive:              true,
			maxFuryLiquidityThreshold:             sdk.NewUint(1336005328924242545),
			currentFuryLiquidityThreshold:         sdk.NewUint(10),
			expectedUpdatedFuryLiquidityThreshold: sdk.NewUint(10),
			expectedPoolUnits:                      sdk.NewUintFromString("1000000000000001000"),
			expectedLPUnits:                        sdk.NewUintFromString("1000000000000000000"),
		},
		{
			name:                     "success - swap external - liquidity protection enabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(0),
				ExternalAssetAmount: sdk.NewUint(68140),
			},
			liquidityProtectionActive:              true,
			maxFuryLiquidityThreshold:             sdk.NewUint(13360053289242425450),
			currentFuryLiquidityThreshold:         sdk.NewUint(10),
			expectedUpdatedFuryLiquidityThreshold: sdk.NewUint(1330659558593215210),
			expectedPoolUnits:                      sdk.NewUintFromString("23662660751003435747009552"),
			expectedLPUnits:                        sdk.NewUintFromString("200546052054071598"),
		},
		{
			name:                     "success - swap native- liquidity protection enabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(4000000000000000000),
				ExternalAssetAmount: sdk.ZeroUint(),
			},
			liquidityProtectionActive:              true,
			maxFuryLiquidityThreshold:             sdk.NewUint(1336005328924242545),
			currentFuryLiquidityThreshold:         sdk.NewUint(1336005328924242544),
			expectedUpdatedFuryLiquidityThreshold: sdk.NewUint(4008015986772728),
			expectedPoolUnits:                      sdk.NewUintFromString("23662660951949037742990437"),
			expectedLPUnits:                        sdk.NewUintFromString("401491654050052483"),
		},
		{
			name:                     "failure - swap native - liquidity protection enabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(4000000000000000000),
				ExternalAssetAmount: sdk.ZeroUint(),
			},
			liquidityProtectionActive:      true,
			maxFuryLiquidityThreshold:     sdk.NewUint(1336005328924242545),
			currentFuryLiquidityThreshold: sdk.NewUint(1336005328924242543),
			errString:                      types.ErrReachedMaxFuryLiquidityThreshold,
		},
		{
			name:                     "failure - swap external - sell external disabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP, tokenregistrytypes.Permission_DISABLE_SELL},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(0),
				ExternalAssetAmount: sdk.NewUint(68140),
			},
			liquidityProtectionActive: false,
			errString:                 tokenregistrytypes.ErrNotAllowedToSellAsset,
		},
		{
			name:                     "failure - swap external - buy native disabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:   []tokenregistrytypes.Permission{tokenregistrytypes.Permission_DISABLE_BUY},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(0),
				ExternalAssetAmount: sdk.NewUint(68140),
			},
			liquidityProtectionActive: false,
			errString:                 tokenregistrytypes.ErrNotAllowedToBuyAsset,
		},
		{
			name:                     "failure - swap native - sell native disabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP},
			nativeAssetPermissions:   []tokenregistrytypes.Permission{tokenregistrytypes.Permission_DISABLE_SELL},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(4000000000000000000),
				ExternalAssetAmount: sdk.ZeroUint(),
			},
			liquidityProtectionActive: false,
			errString:                 tokenregistrytypes.ErrNotAllowedToSellAsset,
		},
		{
			name:                     "failure - swap native - buy external disabled",
			createBalance:            true,
			createPool:               true,
			createLPs:                true,
			poolAsset:                "cusdc",
			address:                  "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
			userNativeAssetBalance:   sdk.Int(sdk.NewUint(4000000000000000000)),
			userExternalAssetBalance: sdk.Int(sdk.NewUint(68140)),
			poolNativeAssetBalance:   sdk.NewUintFromString("157007500498726220240179086"),
			poolExternalAssetBalance: sdk.NewUint(2674623482959),
			poolUnits:                sdk.NewUintFromString("23662660550457383692937954"),
			poolAssetPermissions:     []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP, tokenregistrytypes.Permission_DISABLE_BUY},
			msg: &types.MsgAddLiquidity{
				Signer:              "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
				ExternalAsset:       &types.Asset{Symbol: "cusdc"},
				NativeAssetAmount:   sdk.NewUint(4000000000000000000),
				ExternalAssetAmount: sdk.ZeroUint(),
			},
			liquidityProtectionActive: false,
			errString:                 tokenregistrytypes.ErrNotAllowedToBuyAsset,
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *gridapp.GridironchainApp, genesisState gridapp.GenesisState) gridapp.GenesisState {
				trGs := &tokenregistrytypes.GenesisState{
					Registry: &tokenregistrytypes.Registry{
						Entries: []*tokenregistrytypes.RegistryEntry{
							{Denom: tc.poolAsset, BaseDenom: tc.poolAsset, Decimals: 18, Permissions: tc.poolAssetPermissions},
							{Denom: "fury", BaseDenom: "fury", Decimals: 18, Permissions: tc.nativeAssetPermissions},
						},
					},
				}
				bz, _ := app.AppCodec().MarshalJSON(trGs)
				genesisState["tokenregistry"] = bz

				if tc.createBalance {
					balances := []banktypes.Balance{
						{
							Address: tc.address,
							Coins: sdk.Coins{
								sdk.NewCoin(tc.poolAsset, tc.userExternalAssetBalance),
								sdk.NewCoin("fury", tc.userNativeAssetBalance),
							},
						},
					}

					bankGs := banktypes.DefaultGenesisState()
					bankGs.Balances = append(bankGs.Balances, balances...)
					bz, _ = app.AppCodec().MarshalJSON(bankGs)
					genesisState["bank"] = bz
				}

				if tc.createPool {
					pools := []*types.Pool{
						{
							ExternalAsset:        &types.Asset{Symbol: tc.poolAsset},
							NativeAssetBalance:   tc.poolNativeAssetBalance,
							ExternalAssetBalance: tc.poolExternalAssetBalance,
							PoolUnits:            tc.poolUnits,
							NativeLiabilities:    tc.poolNativeLiabilities,
							ExternalLiabilities:  tc.poolExternalLiabilities,
						},
					}
					clpGs := types.DefaultGenesisState()
					if tc.createLPs {
						lps := []*types.LiquidityProvider{
							{
								Asset:                    &types.Asset{Symbol: tc.poolAsset},
								LiquidityProviderAddress: tc.address,
								LiquidityProviderUnits:   sdk.NewUint(0),
							},
						}
						clpGs.LiquidityProviders = append(clpGs.LiquidityProviders, lps...)
					}
					clpGs.Params = types.Params{
						MinCreatePoolThreshold: 100,
					}
					clpGs.AddressWhitelist = append(clpGs.AddressWhitelist, tc.address)
					clpGs.PoolList = append(clpGs.PoolList, pools...)
					bz, _ = app.AppCodec().MarshalJSON(clpGs)
					genesisState["clp"] = bz
				}

				return genesisState
			})

			app.ClpKeeper.SetPmtpCurrentRunningRate(ctx, sdk.NewDec(1))

			liqProParams := app.ClpKeeper.GetLiquidityProtectionParams(ctx)
			liqProParams.IsActive = tc.liquidityProtectionActive
			liqProParams.MaxFuryLiquidityThreshold = tc.maxFuryLiquidityThreshold
			liqProParams.MaxFuryLiquidityThresholdAsset = types.NativeSymbol
			app.ClpKeeper.SetLiquidityProtectionParams(ctx, liqProParams)

			app.ClpKeeper.SetLiquidityProtectionCurrentFuryLiquidityThreshold(ctx, tc.currentFuryLiquidityThreshold)

			msgServer := clpkeeper.NewMsgServerImpl(app.ClpKeeper)

			_, err := msgServer.AddLiquidity(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.errString != nil {
				require.EqualError(t, err, tc.errString.Error())
				return
			}
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			lp, _ := app.ClpKeeper.GetLiquidityProvider(ctx, tc.poolAsset, tc.address)
			pool, _ := app.ClpKeeper.GetPool(ctx, tc.poolAsset)

			require.Equal(t, tc.expectedPoolUnits.String(), pool.PoolUnits.String()) // compare strings so that the expected amounts can be read from the failure message
			require.Equal(t, tc.expectedLPUnits.String(), lp.LiquidityProviderUnits.String())

			updatedThreshold := app.ClpKeeper.GetLiquidityProtectionRateParams(ctx).CurrentFuryLiquidityThreshold
			require.Equal(t, tc.expectedUpdatedFuryLiquidityThreshold.String(), updatedThreshold.String())
		})
	}
}

func TestMsgServer_AddProviderDistribution(t *testing.T) {
	admin := "did:fury:g1gy2ne7m62uer4h5s4e7xlfq7aeem5zpw2nrxn8"
	nonAdmin := "did:fury:g1gy2ne7m62uer4h5s4e7xlfq7aeem5zpwx6nu9r"
	ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *gridapp.GridironchainApp, genesisState gridapp.GenesisState) gridapp.GenesisState {
		adminGs := &admintypes.GenesisState{
			AdminAccounts: admintest.GetAdmins(admin),
		}
		bz, _ := app.AppCodec().MarshalJSON(adminGs)
		genesisState["admin"] = bz
		trGs := &tokenregistrytypes.GenesisState{
			Registry: nil,
		}
		bz, _ = app.AppCodec().MarshalJSON(trGs)
		genesisState["tokenregistry"] = bz

		return genesisState
	})
	msgServer := clpkeeper.NewMsgServerImpl(app.ClpKeeper)

	_, err := msgServer.AddProviderDistributionPeriod(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)

	var periods []*types.ProviderDistributionPeriod
	validPeriod := types.ProviderDistributionPeriod{DistributionPeriodStartBlock: 10, DistributionPeriodEndBlock: 10, DistributionPeriodBlockRate: sdk.NewDecWithPrec(1, 2), DistributionPeriodMod: 1}
	wrongPeriod := types.ProviderDistributionPeriod{DistributionPeriodStartBlock: 10, DistributionPeriodEndBlock: 9, DistributionPeriodBlockRate: sdk.NewDecWithPrec(1, 2), DistributionPeriodMod: 1}

	periods = append(periods, &wrongPeriod)
	msg := types.MsgAddProviderDistributionPeriodRequest{Signer: admin, DistributionPeriods: periods}
	_, err = msgServer.AddProviderDistributionPeriod(sdk.WrapSDKContext(ctx), &msg)
	require.Error(t, err)
	// check events didn't fire
	require.Equal(t, len(ctx.EventManager().Events()), 0)

	periods[0] = &validPeriod
	msg = types.MsgAddProviderDistributionPeriodRequest{Signer: admin, DistributionPeriods: periods}
	_, err = msgServer.AddProviderDistributionPeriod(sdk.WrapSDKContext(ctx), &msg)
	require.NoError(t, err)
	// check events fired
	require.Equal(t, len(ctx.EventManager().Events()), 2)

	// non admin acc
	msg = types.MsgAddProviderDistributionPeriodRequest{Signer: nonAdmin, DistributionPeriods: periods}
	_, err = msgServer.AddProviderDistributionPeriod(sdk.WrapSDKContext(ctx), &msg)
	require.Error(t, err)
	// check no additional events fired
	require.Equal(t, len(ctx.EventManager().Events()), 2)

	cbp := app.ClpKeeper.GetProviderDistributionParams(ctx)
	require.NotNil(t, cbp)
	require.Equal(t, 1, len(cbp.DistributionPeriods))
	require.Equal(t, *cbp.DistributionPeriods[0], validPeriod)
}
