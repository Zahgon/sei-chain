package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding channel type.
func NewDecodeStore(cdc codec.BinaryCodec, kvA, kvB kv.Pair) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
