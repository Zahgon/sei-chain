package utils

import (
	"iter"
)

// RingBuf is a ring buffer.
// NOT thread-safe.
type RingBuf[T any] struct {
	first int
	len   int
	buf   []T
}

// NewRingBuf creates a new ring buffer with the given capacity.
func NewRingBuf[T any](capacity int) RingBuf[T] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the ring buffer.
func (r *RingBuf[T]) Len() int {
	_ = "STUB: not implemented"

	// Full returns true if the ring buffer is full.
	return 0
}

func (r *RingBuf[T]) Full() bool { _ = "STUB: not implemented"; return false }

// Get returns the i-th element of the ring buffer.
// Panics if i is out of range.
func (r *RingBuf[T]) Get(i int) T { _ = "STUB: not implemented"; return *new(T) }

// TryGet returns the i-th element of the ring buffer.
func (r *RingBuf[T]) TryGet(i int) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Last returns the last element of the ring buffer.
func (r *RingBuf[T]) Last() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// PushBack adds an element to the back of the ring buffer.
// Panics if the ring buffer is full.
func (r *RingBuf[T]) PushBack(x T) { _ = "STUB: not implemented"; return }

// PopFront removes and returns the first element of the ring buffer.
// Panics if the ring buffer is empty.
func (r *RingBuf[T]) PopFront() T { _ = "STUB: not implemented"; return *new(T) }

// All iterates over all the elements in the ring buffer.
func (r *RingBuf[T]) All() iter.Seq[T] { _ = "STUB: not implemented"; return nil }
