package commitment

import (
	"io"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	seidbproto "github.com/sei-protocol/sei-chain/sei-db/proto"
	sctypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

var (
	_ types.CommitKVStore = (*Store)(nil)
	_ types.Queryable     = (*Store)(nil)
)

// Store Implements types.KVStore and CommitKVStore.
type Store struct {
	tree      sctypes.CommitKVStore
	changeSet seidbproto.ChangeSet
}

func NewStore(tree sctypes.CommitKVStore) *Store { _ = "STUB: not implemented"; return nil }

func (st *Store) Commit(_ bool) types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

func (st *Store) LastCommitID() types.CommitID {
	_ = "STUB: not implemented"
	return *new(types.CommitID)
}

// SetPruning panics as pruning options should be provided at initialization
// since IAVl accepts pruning options directly.
func (st *Store) SetPruning(_ types.PruningOptions) { _ = "STUB: not implemented"; return }

// SetPruning panics as pruning options should be provided at initialization
// since IAVl accepts pruning options directly.
func (st *Store) GetPruning() types.PruningOptions {
	_ = "STUB: not implemented"
	return *new(types.PruningOptions)
}

func (st *Store) GetWorkingHash() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Implements Store.
		nil
}

func (st *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

func (st *Store) CacheWrap(k types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements the Store interface.
func (st *Store) CacheWrapWithTrace(k types.StoreKey, w io.Writer, tc types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// Implements types.KVStore.
//
// we assume Set is only called in `Commit`, so the written state is only visible after commit.
func (st *Store) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// Implements types.KVStore.
func (st *Store) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Implements types.KVStore.
func (st *Store) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

// Implements types.KVStore.
//
// we assume Delete is only called in `Commit`, so the written state is only visible after commit.
func (st *Store) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (st *Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

func (st *Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// SetInitialVersion sets the initial version of the IAVL tree. It is used when
// starting a new chain at an arbitrary height.
// implements interface StoreWithInitialVersion
func (st *Store) SetInitialVersion(_ int64) { _ = "STUB: not implemented"; return }

// PopChangeSet returns the change set and clear it
func (st *Store) PopChangeSet() seidbproto.ChangeSet {
	_ = "STUB: not implemented"
	return *new(seidbproto.ChangeSet)
}

func (st *Store) Query(req abci.RequestQuery) (res abci.ResponseQuery) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// get by key
// data holds the key bytes

// get proof from tree and convert to merkle.Proof before adding to result

func (st *Store) VersionExists(version int64) bool {
	_ = "STUB: not implemented"
	// one version per SC tree
	return false
}

func (st *Store) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (st *Store) GetAllKeyStrsInRange(start, end []byte) (res []string) {
	_ = "STUB: not implemented"
	return nil
}

func (st *Store) GetChangedPairs(prefix []byte) (res []*seidbproto.KVPair) {
	_ = "STUB: not implemented"
	// not sure if we can assume pairs are sorted or not, so be conservative
	// here and iterate through everything
	return nil
}
