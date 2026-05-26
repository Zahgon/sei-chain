package state

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
)

func cachingStateFetcher(store Store) func() (State, error) { _ = "STUB: not implemented"; return nil }

// TxConstraintsFetcherFromStore returns the precomputed consensus-derived mempool limits for the
// current persisted state.
func TxConstraintsFetcherFromStore(store Store) mempool.TxConstraintsFetcher {
	_ = "STUB: not implemented"
	return *new(mempool.TxConstraintsFetcher)
}

func TxConstraintsFetcherForState(state State) mempool.TxConstraintsFetcher {
	_ = "STUB: not implemented"
	return *new(mempool.TxConstraintsFetcher)
}
