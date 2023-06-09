package symbol_translator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	gridchainDenomFeedface = "ibc/FEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACE"
	ethereumSymbolFeeface = "Face"
)

func TestNewSymbolTranslatorFromJsonBytes(t *testing.T) {
	_, err := NewSymbolTranslatorFromJSONBytes([]byte("foo"))
	assert.Error(t, err)

	q := ` {"ibc/FEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACE": "Face"} `
	x, err := NewSymbolTranslatorFromJSONBytes([]byte(q))
	assert.NoError(t, err)
	assert.NotNil(t, x)
	assert.Equal(t, x.GridironchainToEthereum(gridchainDenomFeedface), ethereumSymbolFeeface)
	assert.Equal(t, x.EthereumToGridironchain(ethereumSymbolFeeface), gridchainDenomFeedface)
	assert.Equal(t, x.GridironchainToEthereum("verbatim"), "verbatim")
	assert.Equal(t, x.EthereumToGridironchain("verbatim"), "verbatim")
}

func TestNewSymbolTranslator(t *testing.T) {
	s := NewSymbolTranslator()
	assert.Equal(t, s.GridironchainToEthereum("something"), "something")
}
