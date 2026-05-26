package cachekv

import (
	"io"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
	dbm "github.com/tendermint/tm-db"
)

// Store wraps an in-memory cache around an underlying types.KVStore.
type Store struct {
	mtx           sync.RWMutex
	cache         *sync.Map
	deleted       *sync.Map
	unsortedCache *sync.Map
	sortedCache   *dbm.MemDB // always ascending sorted
	parent        types.KVStore
	storeKey      types.StoreKey
	cacheSize     int
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

//----------------------------------------
// Iteration

// Iterator implements types.KVStore.
func (store *Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// ReverseIterator implements types.KVStore.
func (store *Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

func (store *Store) getOrInitSortedCache() *dbm.MemDB { _ = "STUB: not implemented"; return nil }

func (store *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// TODO: (occ) Note that for iterators, we'll need to have special handling (discussed in RFC) to ensure proper validation

// close out parent iterator, then reraise panic

func (store *Store) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

func findStartIndex(strL []string, startQ string) int {
	_ = "STUB: not implemented"
	// Modified binary search to find the very first element in >=startQ.
	return 0
}

// Handle condition where there might be multiple values equal to startQ.
// We are looking for the very first value < midStL, that i+1 will be the first
// element >= midStr.

// midStrL > startQ

func findEndIndex(strL []string, endQ string) int { _ = "STUB: not implemented"; return 0 }

// Modified binary search to find the very first element <endQ.

// Handle condition where there might be multiple values equal to startQ.
// We are looking for the very first value < midStL, that i+1 will be the first
// element >= midStr.

// midStrL > startQ

// Binary search failed, now let's find a value less than endQ.

type sortState int

const (
	stateUnsorted sortState = iota
	stateAlreadySorted
)

// Constructs a slice of dirty items, to use w/ memIterator.
func (store *Store) dirtyItems(start, end []byte) { _ = "STUB: not implemented"; return }

// Nothing to do here.

// If the unsortedCache is too big, its costs too much to determine
// what's in the subset we are concerned about.
// If you are interleaving iterator calls with writes, this can easily become an
// O(N^2) overhead.
// Even without that, too many range checks eventually becomes more expensive
// than just not having the cache.
// store.emitUnsortedCacheSizeMetric()

func (store *Store) clearUnsortedCacheSubset(unsorted []*kv.Pair, sortState sortState) {
	_ = "STUB: not implemented"
	return
}

// deleted element, tracked by store.deleted
// setting arbitrary value

func (store *Store) deleteKeysFromUnsortedCache(unsorted []*kv.Pair) {
	_ = "STUB: not implemented"
	return
}

//----------------------------------------
// etc

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
