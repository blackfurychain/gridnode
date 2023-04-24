package keeper_test

import (
	"github.com/Gridironchain/gridnode/x/dispensation/test"
	"github.com/Gridironchain/gridnode/x/dispensation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeeper_IsLastBlock(t *testing.T) {
	app, ctx := test.CreateTestApp(false)
	assert.False(t, app.DispensationKeeper.IsLastBlock(ctx))
	maxMintAmount, ok := sdk.NewIntFromString(types.MaxMintAmount)
	assert.True(t, ok)
	counterCoin := sdk.NewCoin("fury", maxMintAmount.Sub(sdk.OneInt()))
	app.DispensationKeeper.SetMintController(ctx,
		types.MintController{TotalCounter: counterCoin})
	assert.True(t, app.DispensationKeeper.IsLastBlock(ctx))
}

func TestKeeper_TokensCanBeMinted(t *testing.T) {
	app, ctx := test.CreateTestApp(false)
	assert.True(t, app.DispensationKeeper.TokensCanBeMinted(ctx))
	maxMintAmount, ok := sdk.NewIntFromString(types.MaxMintAmount)
	assert.True(t, ok)
	counterCoin := sdk.NewCoin("fury", maxMintAmount)
	app.DispensationKeeper.SetMintController(ctx,
		types.MintController{TotalCounter: counterCoin})
	assert.False(t, app.DispensationKeeper.TokensCanBeMinted(ctx))
}
