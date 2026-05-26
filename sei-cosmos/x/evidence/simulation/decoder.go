package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/exported"
)

type EvidenceUnmarshaler interface {
	UnmarshalEvidence([]byte) (exported.Evidence, error)
}

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding evidence type.
func NewDecodeStore(cdc EvidenceUnmarshaler) func(kvA, kvB kv.Pair) string {
	_ = "STUB: not implemented"
	return nil
}
