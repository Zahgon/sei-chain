package state

import (
	"io"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	seidbtypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

const StoreTypeSSStore = 100

var (
	_ types.KVStore   = (*Store)(nil)
	_ types.Queryable = (*Store)(nil)
)

// Store wraps a SS store and implements a cosmos KVStore
type Store struct {
	store    seidbtypes.StateStore
	storeKey types.StoreKey
	version  int64
}

func NewStore(store seidbtypes.StateStore, storeKey types.StoreKey, version int64) *Store {
	_ = "STUB: not implemented"
	return nil
}

func (st *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

func (st *Store) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

func (st *Store) CacheWrapWithTrace(storeKey types.StoreKey, w io.Writer, tc types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

func (st *Store) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (st *Store) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

func (st *Store) Set(_, _ []byte) { _ = "STUB: not implemented"; return }

func (st *Store) Delete(_ []byte) { _ = "STUB: not implemented"; return }

func (st *Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

func (st *Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

func (st *Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (st *Store) Query(req abci.RequestQuery) (res abci.ResponseQuery) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseQuery)
}

// get by key
// data holds the key bytes

func (st *Store) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

func (st *Store) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (st *Store) GetAllKeyStrsInRange(start, end []byte) (res []string) {
	_ = "STUB: not implemented"
	return nil
}
