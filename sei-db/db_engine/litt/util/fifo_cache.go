package util

import "time"

var _ Cache[string, string] = &FIFOCache[string, string]{}

// FIFOCache is a cache that evicts the least recently added item when the cache is full. Useful for situations
// where time of addition is a better predictor of future access than time of most recent access.
type FIFOCache[K comparable, V any] struct {
	weightCalculator WeightCalculator[K, V]

	currentWeight uint64
	maxWeight     uint64
	data          map[K]V
	evictionQueue *Queue[*insertionRecord]
	metrics       *CacheMetrics
}

// insertionRecord is a record of when a key was inserted into the cache, and is used to decide when it should be
// evicted.
type insertionRecord struct {
	// The key that was added to the cache.
	key any
	// The time at which the key was added to the cache.
	timestamp time.Time
}

// NewFIFOCache creates a new FIFOCache. If the calculator is nil, the weight of each key-value pair will be 1.
func NewFIFOCache[K comparable, V any](
	maxWeight uint64,
	calculator WeightCalculator[K, V],
	metrics *CacheMetrics) Cache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (f *FIFOCache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (f *FIFOCache[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

// this item won't fit in the cache no matter what we evict

func (f *FIFOCache[K, V]) evict() { _ = "STUB: not implemented"; return }

func (f *FIFOCache[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (f *FIFOCache[K, V]) Weight() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *FIFOCache[K, V]) SetMaxWeight(capacity uint64) { _ = "STUB: not implemented"; return }
