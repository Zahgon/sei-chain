package datastructures

import (
	"sync"

	"golang.org/x/exp/constraints"
)

// A map-like data structure that is guaranteed to be data race free during write
// operations. It is a typed wrapper over the builtin typeless `sync.Map`. The
// CRUD interface is exactly the same as those of `sync.Map`.
type TypedSyncMap[K constraints.Ordered, V any] struct {
	internal *sync.Map
}

func NewTypedSyncMap[K constraints.Ordered, V any]() *TypedSyncMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypedSyncMap[K, V]) Load(key K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *TypedSyncMap[K, V]) Store(key K, value V) { _ = "STUB: not implemented"; return }

func (m *TypedSyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *TypedSyncMap[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

func (m *TypedSyncMap[K, V]) Range(f func(K, V) bool) {
	_ = "STUB: not implemented"
	// All map iterations should be deterministic, so we apply f in sorted order to avoid nondeterminism
	return
}

func (m *TypedSyncMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *TypedSyncMap[K, V]) DeepCopy(copier func(V) V) *TypedSyncMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypedSyncMap[K, V]) DeepApply(toApply func(V)) { _ = "STUB: not implemented"; return }

// A nested map data structure that is guaranteed to be data race free during write
// operations. It is the synchronous equivalent of type map[K1]map[K2]V. Besides
// `sync.Map`'s existing interfaces, it also provides convenient methods to read/write
// nested values directly. For example, to set value `v` for outer key `k1` and inner
// key `k2`, one can simply call StoreNested(k1, k2, v), without worrying about creating
// the inner map if it doesn't exist.
type TypedNestedSyncMap[K1 constraints.Ordered, K2 constraints.Ordered, V any] struct {
	*TypedSyncMap[K1, *TypedSyncMap[K2, V]]
	mu *sync.Mutex // XXXNested methods have write operations outside sync.Map
}

func NewTypedNestedSyncMap[K1 constraints.Ordered, K2 constraints.Ordered, V any]() *TypedNestedSyncMap[K1, K2, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypedNestedSyncMap[K1, K2, V]) LoadNested(key1 K1, key2 K2) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *TypedNestedSyncMap[K1, K2, V]) StoreNested(key1 K1, key2 K2, value V) {
	_ = "STUB: not implemented"
	return
}

func (m *TypedNestedSyncMap[K1, K2, V]) LoadOrStoreNested(key1 K1, key2 K2, value V) (actual V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *TypedNestedSyncMap[K1, K2, V]) DeleteNested(key1 K1, key2 K2) {
	_ = "STUB: not implemented"
	return
}

func (m *TypedNestedSyncMap[K1, K2, V]) DeepCopy(copier func(V) V) *TypedNestedSyncMap[K1, K2, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypedNestedSyncMap[K1, K2, V]) DeepApply(toApply func(V)) {
	_ = "STUB: not implemented"
	return
}
