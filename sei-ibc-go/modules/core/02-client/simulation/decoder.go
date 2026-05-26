package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/keeper"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ ClientUnmarshaler = (*keeper.Keeper)(nil)

// ClientUnmarshaler defines an interface for unmarshaling ICS02 interfaces.
type ClientUnmarshaler interface {
	MustUnmarshalClientState([]byte) exported.ClientState
	MustUnmarshalConsensusState([]byte) exported.ConsensusState
}

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding client type.
func NewDecodeStore(cdc ClientUnmarshaler, kvA, kvB kv.Pair) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
