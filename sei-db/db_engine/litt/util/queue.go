package util

// A standard generic queue.
//
// This struct is not thread safe.
type Queue[T any] struct {
	// The underlying data
	data *RandomAccessDeque[T]
}

// Creates a new Queue with the given initial capacity.
func NewQueue[T any](initialCapacity uint64) *Queue[T] { _ = "STUB: not implemented"; return nil }

// Push an item onto the queue.
func (q *Queue[T]) Push(item T) { _ = "STUB: not implemented"; return }

// Pop an item off the queue. Panics if the queue is empty.
func (q *Queue[T]) Pop() T {
	_ = "STUB: not implemented"
	return *

	// TryPop tries to pop an item off the queue. Returns the item and true if successful, or the zero value
	// and false if the queue is empty.
	new(T)
}

func (q *Queue[T]) TryPop() (item T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

// Peek at the item at the front of the queue without removing it. Panics if the queue is empty.
func (q *Queue[T]) Peek() T {
	_ = "STUB: not implemented"
	return *

	// TryPeek tries to peek at the item at the front of the queue without removing it. Returns the item and true
	// if successful, or the zero value and false if the queue is empty.
	new(T)
}

func (q *Queue[T]) TryPeek() (item T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

// Returns the number of items in the queue.
func (q *Queue[T]) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Clears all items from the queue.
func (q *Queue[T]) Clear() {
	_ = "STUB: not implemented"

	// Get an iterator over the elements in the queue.
	return
}

func (q *Queue[T]) Iterator() func(yield func(uint64, T) bool) {
	_ = "STUB: not implemented"
	return nil

	// Get an item at the given index in the queue. Panics if the index is out of bounds.
}

func (q *Queue[T]) Get(index uint64) T {
	_ = "STUB: not implemented"
	return *

	// Set the item at the given index in the queue. Panics if the index is out of bounds.
	new(T)
}

func (q *Queue[T]) Set(index uint64, value T) (previousValue T) {
	_ = "STUB: not implemented"
	return *new(T)
}
