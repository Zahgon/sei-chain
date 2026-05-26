package datastructures

import (
	"sync"
)

// A set-like data structure that is guaranteed to be data race free during write
// operations. It can return internal data as a slice with a comparator provided,
// so that the resulting slice has a deterministic ordering.
type SyncSet[T comparable] struct {
	dict map[T]bool
	mu   *sync.Mutex
}

func NewSyncSet[T comparable](initial []T) SyncSet[T] { _ = "STUB: not implemented"; return nil }

func (s *SyncSet[T]) Add(val T) { _ = "STUB: not implemented"; return }

func (s *SyncSet[T]) AddAll(vals []T) { _ = "STUB: not implemented"; return }

func (s *SyncSet[T]) Remove(val T) { _ = "STUB: not implemented"; return }

func (s *SyncSet[T]) RemoveAll(vals []T) { _ = "STUB: not implemented"; return }

func (s *SyncSet[T]) Contains(val T) bool { _ = "STUB: not implemented"; return false }

func (s *SyncSet[T]) ToOrderedSlice(comparator func(T, T) bool) []T {
	_ = "STUB: not implemented"
	return nil
}

func (s *SyncSet[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func StringComparator(s1 string, s2 string) bool { _ = "STUB: not implemented"; return false }
