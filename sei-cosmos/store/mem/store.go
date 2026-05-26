package mem

import (
	"io"

	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/dbadapter"
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

var (
	_ types.KVStore   = (*Store)(nil)
	_ types.Committer = (*Store)(nil)
)

// Store implements an in-memory only KVStore. Entries are persisted between
// commits and thus between blocks. State in Memory store is not committed as part of app state but maintained privately by each node
type Store struct {
	dbadapter.Store
}

func NewStore() *Store { _ = "STUB: not implemented"; return nil }

func NewStoreWithDB(db *dbm.MemDB) *Store { _ = "STUB: not implemented"; return nil }

// GetStoreType returns the Store's type.
func (s Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// CacheWrap branches the underlying store.
func (s Store) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements KVStore.
func (s Store) CacheWrapWithTrace(storeKey types.StoreKey, w io.Writer, tc types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// Commit performs a no-op as entries are persistent between commitments.
func (s *Store) Commit(_ bool) (id types.CommitID) {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

func (s *Store) SetPruning(pruning types.PruningOptions) {
	_ = "STUB: not implemented"

	// GetPruning is a no-op as pruning options cannot be directly set on this store.
	// They must be set on the root commit multi-store.
	return
}

func (s *Store) GetPruning() types.PruningOptions {
	_ = "STUB: not implemented"
	return *new(types.PruningOptions)
}

func (s Store) LastCommitID() (id types.CommitID) {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}
