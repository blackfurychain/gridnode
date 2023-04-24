package ethbridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/Gridironchain/gridnode/x/ethbridge/keeper"
	"github.com/Gridironchain/gridnode/x/ethbridge/types"
)

func DefaultGenesis() *types.GenesisState {
	return &types.GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) (res []abci.ValidatorUpdate) {
	// SetCethReceiverAccount
	if data.CethReceiveAccount != "" {
		receiveAccount, err := sdk.AccAddressFromBech32(data.CethReceiveAccount)
		if err != nil {
			panic(err)
		}
		keeper.SetCethReceiverAccount(ctx, receiveAccount)
	}

	// AddPeggyTokens
	if data.PeggyTokens != nil {
		for _, tokenStr := range data.PeggyTokens {
			keeper.AddPeggyToken(ctx, tokenStr)
		}
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	peggyTokens := keeper.GetPeggyToken(ctx)
	receiveAccount := keeper.GetCethReceiverAccount(ctx)

	return &types.GenesisState{
		PeggyTokens:        peggyTokens.Tokens,
		CethReceiveAccount: receiveAccount.String(),
	}
}

func ValidateGenesis(data types.GenesisState) error {
	return nil
}
