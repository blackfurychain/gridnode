package keeper

import (
	"github.com/Gridironchain/gridnode/x/admin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) InitialMigration(ctx sdk.Context) error {
	var accounts []*types.AdminAccount
	if ctx.ChainID() == "gridchain-1" {
		accounts = types.ProdAdminAccounts()
	} else {
		accounts = types.InitialAdminAccounts()
	}
	for _, account := range accounts {
		m.keeper.SetAdminAccount(ctx, account)
	}

	return nil
}
