package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
)

// TransferUnmarshaler defines the expected encoding store functions.
type TransferUnmarshaler interface {
	MustUnmarshalDenomTrace([]byte) types.DenomTrace
}

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding DenomTrace type.
func NewDecodeStore(cdc TransferUnmarshaler) func(kvA, kvB kv.Pair) string {
	_ = "STUB: not implemented"
	return nil
}
