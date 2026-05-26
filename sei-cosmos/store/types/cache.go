package types

import (
	"sync"
)

const DefaultCacheSizeLimit = 4000000 // TODO: revert back to 1000000 after paritioning r/w caches

// If value is nil but deleted is false, it means the parent doesn't have the
// key.  (No need to delete upon Write())
type CValue struct {
	value []byte
	dirty bool
}

func NewCValue(value []byte, dirty bool) *CValue { _ = "STUB: not implemented"; return nil }

func (v *CValue) Value() []byte { _ = "STUB: not implemented"; return nil }

func (v *CValue) Dirty() bool { _ = "STUB: not implemented"; return false }

type CacheBackend interface {
	Get(string) (*CValue, bool)
	Set(string, *CValue)
	Len() int
	Delete(string)
	Range(func(string, *CValue) bool)
}

// This struct is solely for the purpose of preventing the process from crashing because of
// OOM. It is not intended for usage at limit during normal operation. The node operator
// should minimize the time of running at cache limit by switching to a machine with larger
// RAM and bump up cache limit in app config, once the old limit is seen to be reached.
type BoundedCache struct {
	CacheBackend
	limit int

	mu         *sync.Mutex
	metricName []string
}

func NewBoundedCache(backend CacheBackend, limit int) *BoundedCache {
	_ = "STUB: not implemented"
	return nil
}

// cosmos_bounded_cache

func (c *BoundedCache) emitKeysEvictedMetrics(keysToEvict int) { _ = "STUB: not implemented"; return }

func (c *BoundedCache) Set(key string, val *CValue) { _ = "STUB: not implemented"; return }

func (c *BoundedCache) Delete(key string) { _ = "STUB: not implemented"; return }

func (c *BoundedCache) DeleteAll() { _ = "STUB: not implemented"; return }

func (c *BoundedCache) Range(f func(string, *CValue) bool) { _ = "STUB: not implemented"; return }
