package mempool

import (
	"container/list"
	"context"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// TxCache defines an interface for raw transaction caching in a mempool.
// Currently, a TxCache does not allow direct reading or getting of transaction
// values. A TxCache is used primarily to push transactions and removing
// transactions. Pushing via Push returns a boolean telling the caller if the
// transaction already exists in the cache or not.
type TxCache interface {
	// Reset resets the cache to an empty state.
	Reset()

	// Push adds the given transaction key to the cache and returns true if it was
	// newly added. Otherwise, it returns false.
	Push(tx types.TxHash) bool

	// Remove removes the given transaction key from the cache.
	Remove(tx types.TxHash)

	// Size returns the current size of the cache
	Size() int
}

var _ TxCache = (*LRUTxCache)(nil)

// LRUTxCache maintains a thread-safe LRU cache of raw transactions. The cache
// only stores the hash of the raw transaction.
type LRUTxCache struct {
	mtx       sync.Mutex
	size      int
	cacheMap  map[cacheKey]*list.Element
	list      *list.List
	maxKeyLen int
}

type cacheKey = string

// NewLRUTxCache creates an LRU (Least Recently Used) cache that stores
// transactions by key. Keys are derived from the transaction key and trimmed to
// at most maxKeyLen bytes for predictable and efficient storage. If maxKeyLen is
// zero or negative, keys are not trimmed. When the cache exceeds cacheSize, the
// least recently used entry is evicted.
//
// Note that maxKeyLen should be set with care. While a smaller value saves
// memory, it increases the risk of key collisions, which can lead to false
// positives in cache lookups. A larger value reduces collision risk but uses
// more memory. A common choice is to use the full length of a cryptographic hash
// (e.g., 32 bytes for SHA-256) to balance memory usage and collision risk.
func NewLRUTxCache(cacheSize int, maxKeyLen int) *LRUTxCache { _ = "STUB: not implemented"; return nil }

func (c *LRUTxCache) Reset() { _ = "STUB: not implemented"; return }

func (c *LRUTxCache) Push(txHash types.TxHash) bool { _ = "STUB: not implemented"; return false }

func (c *LRUTxCache) Remove(txHash types.TxHash) { _ = "STUB: not implemented"; return }

func (c *LRUTxCache) Size() int { _ = "STUB: not implemented"; return 0 }

func (c *LRUTxCache) toCacheKey(key types.TxHash) cacheKey {
	_ = "STUB: not implemented"
	return *new(cacheKey)
}

// NopTxCache defines a no-op raw transaction cache.
type NopTxCache struct{}

var _ TxCache = (*NopTxCache)(nil)

func (NopTxCache) Reset()                 { _ = "STUB: not implemented"; return }
func (NopTxCache) Push(types.TxHash) bool { _ = "STUB: not implemented"; return false }
func (NopTxCache) Remove(types.TxHash)    { _ = "STUB: not implemented"; return }
func (NopTxCache) Size() int              { _ = "STUB: not implemented"; return 0 }

// DuplicateTxCache implements TxCacheWithTTL using go-cache
type DuplicateTxCache struct {
	maxSize   int
	cache     *cache.Cache
	maxKeyLen int
}

// NewDuplicateTxCache creates a new cache with TTL for transaction keys at a
// given max size. Keys are derived from the transaction key and trimmed to at
// most maxKeyLen bytes for predictable and efficient storage. If maxKeyLen is
// zero or negative, keys are not trimmed. When the cache exceeds cacheSize, the
// least recently used entry is evicted.
//
// Note that maxKeyLen should be set with care. While a smaller value saves
// memory, it increases the risk of key collisions, which can lead to false
// positives in cache lookups. A larger value reduces collision risk but uses
// more memory. A common choice is to use the full length of a cryptographic hash
// (e.g., 32 bytes for SHA-256) to balance memory usage and collision risk.
func NewDuplicateTxCache(maxSize int, defaultExpiration time.Duration, maxKeyLen int) *DuplicateTxCache {
	_ = "STUB: not implemented"
	return nil
}

// Force cleanup interval to 0 - otherwise go-cache leaks a goroutine.
// TODO: replace with a more reasonable implementation of cache, which doesn't do such things.

func (t *DuplicateTxCache) Run(ctx context.Context, cleanupInterval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Periodically delete the expired items.

// Set adds a transaction to the cache with TTL
func (t *DuplicateTxCache) Set(txHash types.TxHash, counter int) { _ = "STUB: not implemented"; return }

// Get retrieves the counter for a transaction key
func (t *DuplicateTxCache) Get(txHash types.TxHash) (counter int, found bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Increment increments the counter for a transaction key, extending TTL
func (t *DuplicateTxCache) Increment(txHash types.TxHash) { _ = "STUB: not implemented"; return }

// Only set a new key if the cache is not full

// Reset clears the cache
func (t *DuplicateTxCache) Reset() {
	_ = "STUB: not implemented"

	// Stop stops the cache and cleans up background goroutines
	return
}

func (t *DuplicateTxCache) Stop() {
	_ = "STUB: not implemented"
	// go-cache doesn't have a Stop method, but we can flush it
	// The janitor goroutine will be cleaned up by the garbage collector
	// when the cache object is no longer referenced
	return
}

func (t *DuplicateTxCache) GetForMetrics() (int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// txHashToString converts a TxHash (byte array) to a stable string key.
func (t *DuplicateTxCache) toCacheKey(key types.TxHash) cacheKey {
	_ = "STUB: not implemented"
	return *new(cacheKey)
}

func trimToSize(key types.TxHash, maxKeyLen int) []byte { _ = "STUB: not implemented"; return nil }
