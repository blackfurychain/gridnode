package key

import (
	"testing"

	"github.com/tyler-smith/go-bip39"
)

var (
	name             = "cool-frost"
	address          = "did:fury:g1hu2lxzusgf4ezp9etyue34zusx822209fhvzts"
	validatorAddress = "did:fury:gvaloper1hu2lxzusgf4ezp9etyue34zusx822209g4fzhk"
	consensusAddress = "did:fury:gvalcons1hu2lxzusgf4ezp9etyue34zusx822209ux67mh"
	random           = "qY3XtZc4a16jTnQWsJUwDvxfC2giHmSB"
	mnemonic         = "flock toss tip service element interest leisure bright subway critic copy lazy zero limb unveil reveal ecology slab detail wrong smooth fun pond choice"
)

func TestGenerateMnemonic(t *testing.T) {
	k := NewKey(name, random)
	k.GenerateMnemonic()

	if !bip39.IsMnemonicValid(k.Mnemonic) {
		t.Error("mnemonic is invalid")
	}
}

func TestRecoverFromMnemonic(t *testing.T) {
	k := NewKey(name, random)
	if err := k.RecoverFromMnemonic(mnemonic); err != nil {
		t.Error(err)
	}

	if k.Address != address {
		t.Errorf("expected %s, got %s", address, k.Address)
	}

	if k.ValidatorAddress != validatorAddress {
		t.Errorf("expected %s, got %s", validatorAddress, k.ValidatorAddress)
	}

	if k.ConsensusAddress != consensusAddress {
		t.Errorf("expected %s, got %s", consensusAddress, k.ConsensusAddress)
	}
}
