package simulation

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding ibc type.
func NewDecodeStore(k keeper.Keeper) func(kvA, kvB kv.Pair) string {
	_ = "STUB: not implemented"
	return nil
}
