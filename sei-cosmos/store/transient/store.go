package transient

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/dbadapter"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

var _ types.Committer = (*Store)(nil)
var _ types.KVStore = (*Store)(nil)

// Store is a wrapper for a MemDB with Commiter implementation
type Store struct {
	dbadapter.Store
}

// Constructs new MemDB adapter
func NewStore() *Store { _ = "STUB: not implemented"; return nil }

// Implements CommitStore
// Commit cleans up Store.
func (ts *Store) Commit(_ bool) (id types.CommitID) {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

func (ts *Store) SetPruning(_ types.PruningOptions) {
	_ = "STUB: not implemented"

	// GetPruning is a no-op as pruning options cannot be directly set on this store.
	// They must be set on the root commit multi-store.
	return
}

func (ts *Store) GetPruning() types.PruningOptions {
	_ = "STUB: not implemented"
	return *new(types.PruningOptions)
}

// Implements CommitStore
func (ts *Store) LastCommitID() (id types.CommitID) {
	_ = "STUB: not implemented"

	// Implements Store.
	return *new(types.CommitID)
}

func (ts *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}
