package im

import (
	"cmp"
	"hash/maphash"
	"iter"

	"github.com/benbjohnson/immutable"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type Map[K comparable, V any] struct{ m *immutable.Map[K, V] }

type hasher[K comparable] struct{ seed maphash.Seed }

func (h hasher[K]) Hash(key K) uint32 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // intentional truncation; only lower 32 bits needed for hash bucketing

func (h hasher[K]) Equal(a, b K) bool { _ = "STUB: not implemented"; return false }

func NewMap[K comparable, V any]() Map[K, V] { _ = "STUB: not implemented"; return nil }

func (m Map[K, V]) Get(key K) (V, bool) {
	_ = "STUB: not implemented"
	return *

	// GetOpt returns the value under key, or None, if key is missing.
	new(V), false
}

func (m Map[K, V]) GetOpt(key K) utils.Option[V] { _ = "STUB: not implemented"; return nil }

func (m Map[K, V]) Set(key K, value V) Map[K, V] { _ = "STUB: not implemented"; return nil }

// SetOpt sets key to the given value, or deletes the key if mvalue is None.
func (m Map[K, V]) SetOpt(key K, mvalue utils.Option[V]) Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m Map[K, V]) Delete(key K) Map[K, V] { _ = "STUB: not implemented"; return nil }

func (m Map[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (m Map[K, V]) All() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

type comparer[K any] func(K, K) int

func (c comparer[K]) Compare(a, b K) int { _ = "STUB: not implemented"; return 0 }

type SortedMap[K, V any] struct{ m *immutable.SortedMap[K, V] }

func NewSortedMap[K, V any](cmp func(K, K) int) SortedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func NewOrderedMap[K cmp.Ordered, V any]() SortedMap[K, V] { _ = "STUB: not implemented"; return nil }

func (m SortedMap[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (m SortedMap[K, V]) Set(key K, value V) SortedMap[K, V] { _ = "STUB: not implemented"; return nil }

func (m SortedMap[K, V]) Delete(key K) SortedMap[K, V] { _ = "STUB: not implemented"; return nil }

func (m SortedMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }
