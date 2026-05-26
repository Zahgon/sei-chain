package prefix

import (
	"io"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

var _ types.KVStore = Store{}

// Store is similar with tendermint/tendermint/libs/db/prefix_db
// both gives access only to the limited subset of the store
// for convinience or safety
type Store struct {
	parent types.KVStore
	prefix []byte
}

func NewStore(parent types.KVStore, prefix []byte) Store {
	_ = "STUB: not implemented"
	return *new(Store)
}

func cloneAppend(bz []byte, tail []byte) (res []byte) { _ = "STUB: not implemented"; return nil }

func (s Store) key(key []byte) (res []byte) { _ = "STUB: not implemented"; return nil }

// Implements Store
func (s Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// Parent returns the underlying KVStore (without the key prefix). Used when unwrapping
// to a root store that supports ABCI proofs (e.g. eth_getProof).
func (s Store) Parent() types.KVStore {
	_ = "STUB: not implemented"

	// Implements CacheWrap
	return *new(types.KVStore)
}

func (s Store) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements the KVStore interface.
func (s Store) CacheWrapWithTrace(storeKey types.StoreKey, w io.Writer, tc types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

func (s Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Implements KVStore
func (s Store) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Implements KVStore
func (s Store) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

// Implements KVStore
func (s Store) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// Implements KVStore
func (s Store) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (s Store) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (s Store) GetAllKeyStrsInRange(start, end []byte) []string {
	_ = "STUB: not implemented"
	return nil
}

// Implements KVStore
// Check https://github.com/tendermint/tendermint/blob/master/libs/db/prefix_db.go#L106
func (s Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// ReverseIterator implements KVStore
// Check https://github.com/tendermint/tendermint/blob/master/libs/db/prefix_db.go#L129
func (s Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

func (s Store) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

var _ types.Iterator = (*prefixIterator)(nil)

type prefixIterator struct {
	prefix []byte
	start  []byte
	end    []byte
	iter   types.Iterator
	valid  bool
}

func newPrefixIterator(prefix, start, end []byte, parent types.Iterator) *prefixIterator {
	_ = "STUB: not implemented"
	return nil
}

// Implements Iterator
func (pi *prefixIterator) Domain() ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil,

		// Implements Iterator
		nil
}

func (pi *prefixIterator) Valid() bool { _ = "STUB: not implemented"; return false }

// Implements Iterator
func (pi *prefixIterator) Next() { _ = "STUB: not implemented"; return }

// TODO: shouldn't pi be set to nil instead?

// Implements Iterator
func (pi *prefixIterator) Key() (key []byte) { _ = "STUB: not implemented"; return nil }

// Implements Iterator
func (pi *prefixIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// Implements Iterator
func (pi *prefixIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Error returns an error if the prefixIterator is invalid defined by the Valid
// method.
func (pi *prefixIterator) Error() error { _ = "STUB: not implemented"; return nil }

// copied from github.com/tendermint/tendermint/libs/db/prefix_db.go
func stripPrefix(key []byte, prefix []byte) []byte { _ = "STUB: not implemented"; return nil }

// wrapping types.PrefixEndBytes
func cpIncr(bz []byte) []byte { _ = "STUB: not implemented"; return nil }
