package keeper_test

import (
	"testing"

	"github.com/Gridironchain/gridnode/x/margin/keeper"
	"github.com/Gridironchain/gridnode/x/margin/test"
	margintypes "github.com/Gridironchain/gridnode/x/margin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_NewQueryServer(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)

	addMTPKey(t, ctx, app, app.MarginKeeper, "ceth", "fury", "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92", margintypes.Position_LONG, 1, sdk.NewDec(20))

	queryServer := keeper.NewQueryServer(app.MarginKeeper)

	res, err := queryServer.GetPositionsForAddress(sdk.WrapSDKContext(ctx), &margintypes.PositionsForAddressRequest{
		Address:    "did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92",
		Pagination: nil,
	})
	require.NoError(t, err)
	require.Len(t, res.Mtps, 1)
}
