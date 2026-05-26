package store

import (
	"io"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// Store wraps an in-memory cache around an underlying types.KVStore.
type Store struct {
	mtx       sync.RWMutex
	cache     *sync.Map
	deleted   *sync.Map
	parent    types.KVStore
	storeKey  types.StoreKey
	cacheSize int
}

var _ types.CacheKVStore = (*Store)(nil)

// NewStore creates a new Store object
func NewStore(parent types.KVStore, storeKey types.StoreKey, cacheSize int) *Store {
	_ = "STUB: not implemented"
	return nil
}

func (store *Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetStoreType implements Store.
func (store *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// getFromCache queries the write-through cache for a value by key.
func (store *Store) getFromCache(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Get implements types.KVStore.
func (store *Store) Get(key []byte) (value []byte) { _ = "STUB: not implemented"; return nil }

// Set implements types.KVStore.
func (store *Store) Set(key []byte, value []byte) { _ = "STUB: not implemented"; return }

// Has implements types.KVStore.
func (store *Store) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

// Delete implements types.KVStore.
func (store *Store) Delete(key []byte) { _ = "STUB: not implemented"; return }

// Implements Cachetypes.KVStore.
func (store *Store) Write() { _ = "STUB: not implemented"; return }

// We need a copy of all of the keys.
// Not the best, but probably not a bottleneck depending.

// TODO: Consider allowing usage of Batch, which would allow the write to
// at least happen atomically.

// We use []byte(key) instead of conv.UnsafeStrToBytes because we cannot
// be sure if the underlying store might do a save with the byteslice or
// not. Once we get confirmation that .Delete is guaranteed not to
// save the byteslice, then we can assume only a read-only copy is sufficient.

// It already exists in the parent, hence delete it.

// CacheWrap implements CacheWrapper.
func (store *Store) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements the CacheWrapper interface.
func (store *Store) CacheWrapWithTrace(storeKey types.StoreKey, w io.Writer, tc types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

func (store *Store) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

// Only entrypoint to mutate store.cache.
func (store *Store) setCacheValue(key, value []byte, deleted bool, dirty bool) {
	_ = "STUB: not implemented"
	return
}

func (store *Store) isDeleted(key string) bool { _ = "STUB: not implemented"; return false }

func (store *Store) GetParent() types.KVStore {
	_ = "STUB: not implemented"
	return *new(types.KVStore)
}

func (store *Store) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (store *Store) GetAllKeyStrsInRange(start, end []byte) (res []string) {
	_ = "STUB: not implemented"
	return nil
}

// we don't want to break out of the iteration since cache isn't sorted

func (store *Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// ReverseIterator implements types.KVStore.
func (store *Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}
