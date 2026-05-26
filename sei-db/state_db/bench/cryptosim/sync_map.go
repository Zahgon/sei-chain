package cryptosim

import (
	"iter"
	"sync"
)

// A thread safe map-like data structure. Unlike sync.Map, supports generics.
type SyncMap[K comparable, V any] struct {
	base sync.Map
}

// NewSyncMap returns a new empty SyncMap.
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] { _ = "STUB: not implemented"; return nil }

// Put stores the key-value pair in the map.
func (m *SyncMap[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

// Clear removes all key-value pairs from the map.
func (m *SyncMap[K, V]) Clear() {
	_ = "STUB: not implemented"

	// Get returns the value for key and true if present, or the zero value of V and false otherwise.
	return
}

func (m *SyncMap[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// All returns an iterator over the map's key-value pairs for use with range.
func (m *SyncMap[K, V]) Iterator() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }
