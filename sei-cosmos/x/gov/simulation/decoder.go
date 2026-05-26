package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding gov type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	_ = "STUB: not implemented"
	return nil
}
