package util

import "sync"

var _ Cache[string, string] = &threadSafeCache[string, string]{}

// threadSafeCache is a thread-safe wrapper around a Cache.
type threadSafeCache[K comparable, V any] struct {
	cache Cache[K, V]
	lock  sync.RWMutex
}

// NewThreadSafeCache wraps a Cache in a thread-safe wrapper.
func NewThreadSafeCache[K comparable, V any](cache Cache[K, V]) Cache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (t *threadSafeCache[K, V]) Get(key K) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (t *threadSafeCache[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (t *threadSafeCache[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (t *threadSafeCache[K, V]) Weight() uint64 { _ = "STUB: not implemented"; return 0 }

func (t *threadSafeCache[K, V]) SetMaxWeight(capacity uint64) { _ = "STUB: not implemented"; return }
