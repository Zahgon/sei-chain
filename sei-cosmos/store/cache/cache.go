package cache

import (
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"

	lru "github.com/hashicorp/golang-lru/v2"
)

var (
	_ types.CommitKVStore             = (*CommitKVStoreCache)(nil)
	_ types.MultiStorePersistentCache = (*CommitKVStoreCacheManager)(nil)

	// DefaultCommitKVStoreCacheSize defines the persistent ARC cache size for a
	// CommitKVStoreCache.
	DefaultCommitKVStoreCacheSize uint = 100000
)

type (
	// CommitKVStoreCache implements an inter-block (persistent) cache that wraps a
	// CommitKVStore. Reads first hit the internal ARC (Adaptive Replacement Cache).
	// During a cache miss, the read is delegated to the underlying CommitKVStore
	// and cached. Deletes and writes always happen to both the cache and the
	// CommitKVStore in a write-through manner. Caching performed in the
	// CommitKVStore and below is completely irrelevant to this layer.
	CommitKVStoreCache struct {
		types.CommitKVStore
		cache       *lru.TwoQueueCache[string, []byte]
		cacheKVSize int

		// the same CommitKVStoreCache may be accessed concurrently by multiple
		// goroutines due to transaction parallelization
		mtx sync.RWMutex
	}

	// CommitKVStoreCacheManager maintains a mapping from a StoreKey to a
	// CommitKVStoreCache. Each CommitKVStore, per StoreKey, is meant to be used
	// in an inter-block (persistent) manner and typically provided by a
	// CommitMultiStore.
	CommitKVStoreCacheManager struct {
		cacheSize   uint
		caches      map[string]types.CommitKVStore
		cacheKVSize int
	}
)

func NewCommitKVStoreCache(store types.CommitKVStore, size uint, cacheKVSize int) *CommitKVStoreCache {
	_ = "STUB: not implemented"
	return nil
}

//#nosec G115 -- bounds checked above

func NewCommitKVStoreCacheManager(size uint, cacheKVSize int) *CommitKVStoreCacheManager {
	_ = "STUB: not implemented"
	return nil
}

// GetStoreCache returns a Cache from the CommitStoreCacheManager for a given
// StoreKey. If no Cache exists for the StoreKey, then one is created and set.
// The returned Cache is meant to be used in a persistent manner.
func (cmgr *CommitKVStoreCacheManager) GetStoreCache(key types.StoreKey, store types.CommitKVStore) types.CommitKVStore {
	_ = "STUB: not implemented"
	return *new(types.CommitKVStore)
}

// Unwrap returns the underlying CommitKVStore for a given StoreKey.
func (cmgr *CommitKVStoreCacheManager) Unwrap(key types.StoreKey) types.CommitKVStore {
	_ = "STUB: not implemented"
	return *new(types.CommitKVStore)
}

// Reset resets in the internal caches.
func (cmgr *CommitKVStoreCacheManager) Reset() { _ = "STUB: not implemented"; return }

// not deleting CommitKVStoreCache themselves from the manager to prevent
// Unwrap returning nil

// CacheWrap implements the CacheWrapper interface
func (ckv *CommitKVStoreCache) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// getFromCache queries the write-through cache for a value by key.
func (ckv *CommitKVStoreCache) getFromCache(key []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// getAndWriteToCache queries the underlying CommitKVStore and writes the result
func (ckv *CommitKVStoreCache) getAndWriteToCache(key []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves a value by key. It will first look in the write-through cache.
// If the value doesn't exist in the write-through cache, the query is delegated
// to the underlying CommitKVStore.
func (ckv *CommitKVStoreCache) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// if not found in the cache, query the underlying CommitKVStore and init cache value

// Set inserts a key/value pair into both the write-through cache and the
// underlying CommitKVStore.
func (ckv *CommitKVStoreCache) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// Delete removes a key/value pair from both the write-through cache and the
// underlying CommitKVStore.
func (ckv *CommitKVStoreCache) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (ckv *CommitKVStoreCache) Reset() { _ = "STUB: not implemented"; return }
