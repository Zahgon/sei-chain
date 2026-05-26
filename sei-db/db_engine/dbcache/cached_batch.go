package dbcache

import (
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// cachedBatch wraps a types.Batch and applies pending mutations to the cache
// after a successful commit.
type cachedBatch struct {
	inner   types.Batch
	cache   Cache
	pending []CacheUpdate
}

var _ types.Batch = (*cachedBatch)(nil)

func newCachedBatch(inner types.Batch, cache Cache) *cachedBatch {
	_ = "STUB: not implemented"
	return nil
}

func (cb *cachedBatch) Set(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (cb *cachedBatch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (cb *cachedBatch) Commit(opts types.WriteOptions) error { _ = "STUB: not implemented"; return nil }

// A cache write can only fail during a shutdown when the cache's context is cancelled,
// or when the cache's work pools have their contexts cancelled. Continuing to use the
// cache after shutdown is not permissible, and so this method must return an error.

func (cb *cachedBatch) Len() int { _ = "STUB: not implemented"; return 0 }

func (cb *cachedBatch) Reset() { _ = "STUB: not implemented"; return }

func (cb *cachedBatch) Close() error { _ = "STUB: not implemented"; return nil }
