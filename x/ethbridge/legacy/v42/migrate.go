package v42

import (
	v039ethbridge "github.com/Gridironchain/gridnode/x/ethbridge/legacy/v39"
	"github.com/Gridironchain/gridnode/x/ethbridge/types"
)

func Migrate(state v039ethbridge.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		CethReceiveAccount: state.CethReceiverAccount.String(),
		PeggyTokens:        state.PeggyTokens,
	}
}
