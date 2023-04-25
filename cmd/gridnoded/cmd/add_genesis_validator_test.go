package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/Gridironchain/gridnode/app"
	//gridnodedcmd "github.com/Gridironchain/gridnode/cmd/gridnoded/cmd"
	"github.com/Gridironchain/gridnode/x/oracle"
)

func TestAddGenesisValidatorCmd(t *testing.T) {
	homeDir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer func(path string) {
		err := os.RemoveAll(path)
		require.NoError(t, err)
	}(homeDir)
	initCmd, _ := NewRootCmd()
	initBuf := new(bytes.Buffer)
	initCmd.SetOut(initBuf)
	initCmd.SetErr(initBuf)

	initCmd.SetArgs([]string{"init", "test", "--home=" + homeDir})
	app.SetConfig(false)
	expectedValidatorBech32 := "did:fury:gvaloper1rwqp4q88ue83ag3kgnmxxypq0td59df4r3mcwj"
	expectedValidator, err := sdk.ValAddressFromBech32(expectedValidatorBech32)
	require.NoError(t, err)
	addValCmd, _ := NewRootCmd()
	addValBuf := new(bytes.Buffer)
	addValCmd.SetOut(addValBuf)
	addValCmd.SetErr(addValBuf)
	addValCmd.SetArgs([]string{"add-genesis-validators", expectedValidatorBech32, "--home=" + homeDir})
	// Run init
	err = svrcmd.Execute(initCmd, homeDir)
	require.NoError(t, err)
	// Run add-genesis-validators

	err = svrcmd.Execute(addValCmd, homeDir)
	require.NoError(t, err)
	// Load genesis state from temp home dir and parse JSON
	serverCtx := server.GetServerContextFromCmd(addValCmd)
	genFile := serverCtx.Config.GenesisFile()
	appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
	require.NoError(t, err)
	// Setup app to get oracle keeper and ctx.
	gridapp := app.Setup(false)
	ctx := gridapp.BaseApp.NewContext(false, tmproto.Header{})
	// Run loaded genesis through InitGenesis on oracle module
	mm := module.NewManager(
		oracle.NewAppModule(gridapp.OracleKeeper),
	)
	_ = mm.InitGenesis(ctx, gridapp.AppCodec(), appState)
	// Assert validator
	validators := gridapp.OracleKeeper.GetOracleWhiteList(ctx)
	require.Equal(t, []sdk.ValAddress{expectedValidator}, validators)
}
