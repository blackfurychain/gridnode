package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// This file contains sample values which can be used to speed up the process of writing test cases
// Any or all values can be replaced in the the individual functions where these are used
func getDefaultArgs() Args {
	amount, ok := sdk.NewIntFromString("100000000000000000000000")
	if !ok {
		panic("Cannot parse amount")
	}

	senderName := "grid"
	path := hd.CreateHDPath(118, 0, 0).String()
	toAddr, err := sdk.AccAddressFromBech32("did:fury:g1l7hypmqk2yc334vc6vmdwzp5sdefygj23y4thn")
	if err != nil {
		panic(toAddr)
	}

	kr, err := keyring.New("gridironchain", "test", os.TempDir(), nil)
	if err != nil {
		panic(err)
	}
	mnemonic := "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow"

	accInfo, err := kr.NewAccount(senderName, mnemonic, "", path, hd.Secp256k1)
	if err != nil {
		accInfo, err = kr.Key(senderName)
		if err != nil {
			panic(err)
		}
	}

	return Args{
		ChainID:          "gridironchain-devnet-1",
		GasPrice:         "",
		GasAdjustment:    0,
		Keybase:          kr,
		ChannelID:        "",
		Sender:           accInfo.GetAddress(),
		GridironchainReceiver: toAddr,
		CosmosReceiver:   "",
		Amount:           sdk.NewCoins(sdk.NewCoin("fury", amount)),
		TimeoutTimestamp: 0,
		Fees:             "1000000fury",
		Network:          Devnet,
		SenderName:       senderName,
	}
}

func defaultAssert(res *sdk.TxResponse, testName string) *sdk.TxResponse {
	// Works only in block
	if res.Code != 0 {
		panic("Transaction Failed")
	}
	return res
}

func setNetwork(args *Args, network Network) {
	args.Network = network
	switch args.Network {
	case Devnet:
		args.ChainID = "gridironchain-devnet"
	case TestNet:
		args.ChainID = "gridironchain-testnet"
	case MainNet:
		args.ChainID = "gridironchain"
	case LocalNet:
		args.ChainID = "localnet"
	default:
		panic("Network is a required arg")
	}
}
