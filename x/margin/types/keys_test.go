package types_test

import (
	"reflect"
	"testing"

	"github.com/Gridironchain/gridnode/x/margin/types"
)

func TestTypes_GetMTPKey(t *testing.T) {
	got := types.GetMTPKey("xxx", 1)
	want := []byte{1, 120, 120, 120, 0, 0, 0, 0, 0, 0, 0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
