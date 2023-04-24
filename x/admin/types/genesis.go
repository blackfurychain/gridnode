package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

func UnmarshalGenesis(marshaler codec.JSONCodec, state json.RawMessage) GenesisState {
	var genesisState GenesisState
	if state != nil {
		err := marshaler.UnmarshalJSON(state, &genesisState)
		if err != nil {
			panic(fmt.Sprintf("Failed to get genesis state from app state: %s", err.Error()))
		}
	}
	return genesisState
}

func ProdAdminAccounts() []*AdminAccount {
	return []*AdminAccount{
		{
			AdminType:    AdminType_ADMIN,
			AdminAddress: "grid144w8cpva2xkly74xrms8djg69y3mljzplx3fjt",
		},
		{
			AdminType:    AdminType_PMTPREWARDS,
			AdminAddress: "grid144w8cpva2xkly74xrms8djg69y3mljzplx3fjt",
		},
		{
			AdminType:    AdminType_CLPDEX,
			AdminAddress: "grid144w8cpva2xkly74xrms8djg69y3mljzplx3fjt",
		},
		{
			AdminType:    AdminType_TOKENREGISTRY,
			AdminAddress: "grid144w8cpva2xkly74xrms8djg69y3mljzplx3fjt",
		},
		{
			AdminType:    AdminType_ETHBRIDGE,
			AdminAddress: "grid144w8cpva2xkly74xrms8djg69y3mljzplx3fjt",
		},
		{
			AdminType:    AdminType_TOKENREGISTRY,
			AdminAddress: "grid1npzemsc4s5gxpv2qt3na97tna03cj2h5xxe3cw",
		},
		{
			AdminType:    AdminType_ETHBRIDGE,
			AdminAddress: "grid10wgwh7g3jktemd4d8jnswnf0zyk3hsq3uk3tff",
		},
	}
}

func InitialAdminAccounts() []*AdminAccount {
	return []*AdminAccount{
		{
			AdminType:    AdminType_ADMIN,
			AdminAddress: "grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd",
		},
		{
			AdminType:    AdminType_CLPDEX,
			AdminAddress: "grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd",
		},
		{
			AdminType:    AdminType_CLPDEX,
			AdminAddress: "grid1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5",
		},
		{
			AdminType:    AdminType_PMTPREWARDS,
			AdminAddress: "grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd",
		},
		{
			AdminType:    AdminType_PMTPREWARDS,
			AdminAddress: "grid144w8cpva2xkly74xrms8djg69y3mljzplx3fjt",
		},
		{
			AdminType:    AdminType_PMTPREWARDS,
			AdminAddress: "grid1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5",
		},
		{
			AdminType:    AdminType_ETHBRIDGE,
			AdminAddress: "grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd",
		},
		{
			AdminType:    AdminType_TOKENREGISTRY,
			AdminAddress: "grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd",
		},
	}
}
