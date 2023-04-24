package keeper_test

import (
	"testing"

	"github.com/Gridironchain/gridnode/x/ethbridge/test"
	"github.com/Gridironchain/gridnode/x/ibctransfer/keeper"
	scibctransfertypes "github.com/Gridironchain/gridnode/x/ibctransfer/types"
	scibctransfermocks "github.com/Gridironchain/gridnode/x/ibctransfer/types/mocks"
	tokenregistrytest "github.com/Gridironchain/gridnode/x/tokenregistry/test"
	tokenregistrytypes "github.com/Gridironchain/gridnode/x/tokenregistry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktransfertypes "github.com/cosmos/ibc-go/v2/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

/* Test that when a conversion is needed the right amounts are converted before sending to underlying SDK Transfer. */
func TestMsgServer_Transfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	bankKeeper := scibctransfermocks.NewMockBankKeeper(ctrl)
	msgSrv := scibctransfermocks.NewMockMsgServer(ctrl)
	app, ctx, _ := tokenregistrytest.CreateTestApp(false)
	addrs, _ := test.CreateTestAddrs(2)
	app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
		Denom:                "fury",
		Decimals:             18,
		IbcCounterpartyDenom: "xfury",
		Permissions:          []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCEXPORT},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
		Denom:       "xfury",
		Decimals:    10,
		UnitDenom:   "fury",
		Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCEXPORT},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
		Denom:       "misconfigured",
		Decimals:    18,
		Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_IBCEXPORT},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &tokenregistrytypes.RegistryEntry{
		Denom:                "ceth",
		Decimals:             18,
		IbcCounterpartyDenom: "ceth",
		Permissions:          []tokenregistrytypes.Permission{},
	})
	furySmallest, ok := sdk.NewIntFromString("183456789")
	require.True(t, ok)
	tt := []struct {
		name                 string
		err                  error
		bankKeeper           scibctransfertypes.BankKeeper
		msgSrv               scibctransfertypes.MsgServer
		msg                  *sdktransfertypes.MsgTransfer
		setupMsgServerCalls  func()
		setupBankKeeperCalls func()
	}{
		{
			name:       "transfer smallest fury without rounding",
			bankKeeper: bankKeeper,
			msgSrv:     msgSrv,
			msg: sdktransfertypes.NewMsgTransfer(
				"transfer",
				"channel-0",
				sdk.NewCoin("fury", furySmallest),
				addrs[0].String(),
				addrs[1].String(),
				clienttypes.NewHeight(0, 0),
				0,
			),
			setupMsgServerCalls: func() {
				msgSrv.EXPECT().Transfer(gomock.Any(), &sdktransfertypes.MsgTransfer{
					SourcePort:       "transfer",
					SourceChannel:    "channel-0",
					Token:            sdk.NewCoin("fury", furySmallest),
					Sender:           addrs[0].String(),
					Receiver:         addrs[1].String(),
					TimeoutHeight:    clienttypes.NewHeight(0, 0),
					TimeoutTimestamp: 0,
				})
			},
			setupBankKeeperCalls: func() {},
		},
		{
			name:       "transfer denom without ibc export permission",
			err:        tokenregistrytypes.ErrPermissionDenied,
			bankKeeper: bankKeeper,
			msgSrv:     msgSrv,
			msg: sdktransfertypes.NewMsgTransfer(
				"transfer",
				"channel-0",
				sdk.NewCoin("ceth", sdk.NewInt(1)),
				addrs[0].String(),
				addrs[1].String(),
				clienttypes.NewHeight(0, 0),
				0,
			),
			setupBankKeeperCalls: func() {},
			setupMsgServerCalls:  func() {},
		},
		{
			name:       "transfer denom is not whitelisted",
			err:        tokenregistrytypes.ErrPermissionDenied,
			bankKeeper: bankKeeper,
			msgSrv:     msgSrv,
			msg: sdktransfertypes.NewMsgTransfer(
				"transfer",
				"channel-0",
				sdk.NewCoin("caave", sdk.NewInt(1)),
				addrs[0].String(),
				addrs[1].String(),
				clienttypes.NewHeight(0, 0),
				0,
			),
			setupBankKeeperCalls: func() {},
			setupMsgServerCalls:  func() {},
		},
		{
			name:       "transfer denom alias with unit denom set in registry",
			err:        tokenregistrytypes.ErrPermissionDenied,
			bankKeeper: bankKeeper,
			msgSrv:     msgSrv,
			msg: sdktransfertypes.NewMsgTransfer(
				"transfer",
				"channel-0",
				sdk.NewCoin("xfury", sdk.NewInt(1)),
				addrs[0].String(),
				addrs[1].String(),
				clienttypes.NewHeight(0, 0),
				0,
			),
			setupBankKeeperCalls: func() {},
			setupMsgServerCalls:  func() {},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMsgServerCalls()
			tc.setupBankKeeperCalls()
			srv := keeper.NewMsgServerImpl(tc.msgSrv, tc.bankKeeper, app.TokenRegistryKeeper)
			_, err := srv.Transfer(sdk.WrapSDKContext(ctx), tc.msg)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
