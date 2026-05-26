package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

type AuthUnmarshaler interface {
	UnmarshalAccount([]byte) (types.AccountI, error)
	GetCodec() codec.BinaryCodec
}

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding auth type.
func NewDecodeStore(ak AuthUnmarshaler) func(kvA, kvB kv.Pair) string {
	_ = "STUB: not implemented"
	return nil
}
