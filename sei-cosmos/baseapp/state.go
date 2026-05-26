package baseapp

import (
	"sync"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type state struct {
	ms  sdk.CacheMultiStore
	ctx sdk.Context
	mtx *sync.RWMutex
}

// CacheMultiStore calls and returns a CacheMultiStore on the state's underling
// CacheMultiStore.
func (st *state) CacheMultiStore() sdk.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(sdk.CacheMultiStore)
}

func (st *state) MultiStore() sdk.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(sdk.CacheMultiStore)
}

func (st *state) SetMultiStore(ms sdk.CacheMultiStore) *state {
	_ = "STUB: not implemented"
	return nil
}

// Context returns the Context of the state.
func (st *state) Context() sdk.Context { _ = "STUB: not implemented"; return *new(sdk.Context) }

func (st *state) SetContext(ctx sdk.Context) *state { _ = "STUB: not implemented"; return nil }
