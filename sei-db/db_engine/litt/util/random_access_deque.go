package util

// The minimum initial capacity of a RandomAccessDeque.
const minimumInitialCapacity = 32

// A double-ended queue (deque) that supports O(1) lookup by index.
//
// - Insertion time: O(1) average, O(n) worst-case (when resizing is needed)
// - Deletion time: O(1) average, array space is not reclaimed
// - Lookup time by index: O(1)
// - Iteration: O(1) to build iterator, O(1) per step
//
// This data structure is not thread safe.
type RandomAccessDeque[T any] struct {
	// The current number of elements in the deque.
	size uint64
	// Underlying data storage
	data []T
	// The index in data that corresponds to the logical start of the deque.
	startIndex uint64
	// The index in data that corresponds to the logical end of the deque (one past the last element).
	endIndex uint64
	// The initial capacity of the deque. Used when calling Clear().
	initialCapacity uint64
}

// Create a new RandomAccessDeque with the specified initial capacity. Queue can grow beyond this capacity if needed.
func NewRandomAccessDeque[T any](initialCapacity uint64) *RandomAccessDeque[T] {
	_ = "STUB: not implemented"
	return nil
}

// Get the number of elements in the deque.
//
// O(1)
func (s *RandomAccessDeque[T]) Size() uint64 {
	_ = "STUB: not implemented"

	// Syntactic sugar for Size() == 0
	//
	// O(1)
	return 0
}

func (s *RandomAccessDeque[T]) IsEmpty() bool {
	_ = "STUB: not implemented"

	// Insert a value at the front of the deque. This value will have index 0 after insertion, and all other values will
	// have their indices increased by 1.
	//
	// O(1) average, O(n) worst-case (when resizing is needed)
	return false
}

func (s *RandomAccessDeque[T]) PushFront(value T) { _ = "STUB: not implemented"; return }

// wrap around

// Return the value at the front of the deque without removing it. Panics if the deque is empty.
//
// O(1)
func (s *RandomAccessDeque[T]) PeekFront() T { _ = "STUB: not implemented"; return *new(T) }

// Return the value at the front of the deque without removing it. If the deque is empty, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TryPeekFront() (value T, ok bool) {
	_ = "STUB: not implemented"
	return *

	// Remove and return the value at the front of the deque. Panics if the deque is empty.
	//
	// O(1)
	new(T), false
}

func (s *RandomAccessDeque[T]) PopFront() T { _ = "STUB: not implemented"; return *new(T) }

// Remove and return the value at the front of the deque. If the deque is empty, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TryPopFront() (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

//nolint:gosec // slice length non-negative
// wrap around

// Insert a value at the back of the deque. This value will have index Size()-1 after insertion.
//
// O(1) average, O(n) worst-case (when resizing is needed)
func (s *RandomAccessDeque[T]) PushBack(value T) { _ = "STUB: not implemented"; return }

//nolint:gosec // slice length non-negative
// wrap around

// Return the value at the back of the deque without removing it. Panics if the deque is empty.
//
// O(1)
func (s *RandomAccessDeque[T]) PeekBack() T { _ = "STUB: not implemented"; return *new(T) }

// Return the value at the back of the deque without removing it. If the deque is empty, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TryPeekBack() (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Remove and return the value at the back of the deque. Panics if the deque is empty.
//
// O(1)
func (s *RandomAccessDeque[T]) PopBack() T { _ = "STUB: not implemented"; return *new(T) }

// Remove and return the value at the back of the deque. If the deque is empty, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TryPopBack() (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Get the value at the specified index. Panics if the index is out of bounds.
//
// O(1)
func (s *RandomAccessDeque[T]) Get(index uint64) T { _ = "STUB: not implemented"; return *new(T) }

// Get the value at the specified index. If the index is out of bounds, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TryGet(index uint64) (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Get an element indexed from the last thing in the deque. Equivalent to Get(Size() - 1 - index).
// Panics if the index is out of bounds.
//
// O(1)
func (s *RandomAccessDeque[T]) GetFromBack(index uint64) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Get an element indexed from the last thing in the deque. Equivalent to TryGet(Size() - 1 - index).
// If the index is out of bounds, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TryGetFromBack(index uint64) (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Set the value at the specified index, replacing the existing value, which is returned.
// Panics if the index is out of bounds.
//
// O(1)
func (s *RandomAccessDeque[T]) Set(index uint64, value T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Set the value at the specified index, replacing the existing value, which is returned.
// If the index is out of bounds, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TrySet(index uint64, value T) (previousValue T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Set an element indexed from the last thing in the deque, replacing the existing value, which is returned.
// Equivalent to Set(Size() - 1 - index, value).
// Panics if the index is out of bounds.
//
// O(1)
func (s *RandomAccessDeque[T]) SetFromBack(index uint64, value T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Set an element indexed from the last thing in the deque, replacing the existing value, which is returned.
// Equivalent to TrySet(Size() - 1 - index, value).
// If the index is out of bounds, returns ok==false.
//
// O(1)
func (s *RandomAccessDeque[T]) TrySetFromBack(index uint64, value T) (previousValue T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Clear all elements from the deque. Reclaims space in the underlying array.
//
// O(1)
func (s *RandomAccessDeque[T]) Clear() { _ = "STUB: not implemented"; return }

// Reset the underlying array to allow garbage collection of contained elements.

// Get an iterator over the elements in the deque, from front to back. It is not safe to get an iterator,
// modify the deque, and then use the iterator again.
//
// O(1) to call this method, O(1) per iteration step.
func (s *RandomAccessDeque[T]) Iterator() func(yield func(uint64, T) bool) {
	_ = "STUB: not implemented"
	return nil
}

// no-op

// Get an iterator over the elements in the deque, from the specified index to back. It is not safe to get an iterator,
// modify the deque, and then use the iterator again.
// Panics if the index is out of bounds.
//
// O(1) to call this method, O(1) per iteration step.
func (s *RandomAccessDeque[T]) IteratorFrom(index uint64) func(yield func(uint64, T) bool) {
	_ = "STUB: not implemented"
	return nil
}

// Get an iterator over the elements in the deque, from the specified index to back. It is not safe to get an iterator,
// modify the deque, and then use the iterator again.
// If the index is out of bounds, returns ok==false.
//
// O(1) to call this method, O(1) per iteration step.
func (s *RandomAccessDeque[T]) TryIteratorFrom(index uint64) (func(yield func(uint64, T) bool), bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Get an iterator over the elements in the deque, from back to front. It is not safe to get an iterator,
// modify the deque, and then use the iterator again.
//
// O(1) to call this method, O(1) per iteration step.
func (s *RandomAccessDeque[T]) ReverseIterator() func(yield func(uint64, T) bool) {
	_ = "STUB: not implemented"
	return nil
}

// no-op

// Get an iterator over the elements in the deque, from the specified index to front. It is not safe to get an iterator,
// modify the deque, and then use the iterator again.
// Panics if the index is out of bounds.
//
// O(1) to call this method, O(1) per iteration step.
func (s *RandomAccessDeque[T]) ReverseIteratorFrom(index uint64) func(yield func(uint64, T) bool) {
	_ = "STUB: not implemented"
	return nil
}

// Get an iterator over the elements in the deque, from the specified index to front. It is not safe to get an iterator,
// modify the deque, and then use the iterator again.
// If the index is out of bounds, returns ok==false.
//
// O(1) to call this method, O(1) per iteration step.
func (s *RandomAccessDeque[T]) TryReverseIteratorFrom(index uint64) (func(yield func(uint64, T) bool), bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Resize the underlying array to accommodate at least one more insertion. Preserves existing elements.
// If no resizing is needed, this is a no-op.
func (s *RandomAccessDeque[T]) resizeForInsertion() { _ = "STUB: not implemented"; return }

// Perform a binary search in the deque for an element matching the compare function. Assumes that
// the deque is sorted according to the same compare function. If an exact match can't be found,
// returns the index of the location where the value would be inserted if it were inserted in the proper location.
//
// The compare function `compare(a V, b T) int` should return:
//   - negative value if a < b
//   - zero if a == b
//   - positive value if a > b
//
// If the deque is not sorted or if the ordering is not a total ordering, the return value is undefined. This function
// is not defined as a method on RandomAccessDeque due to this fact. Not all RandomAccessDeque instances will be sorted,
// and so this function is not always valid to call.
func BinarySearchInOrderedDeque[V any, T any](
	deque *RandomAccessDeque[T],
	value V,
	compare func(a V, b T) int) (index uint64, exact bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Index is the external index in the deque, from 0 to size-1, not indices as they
// appear in the underlying array.

// We've found an exact match.

// value < target, search left half
//
//      value is here
//  |-----------------------|-----------------------|
// left                   target                  right

// value > target, search right half
//
//                               value is here
//  |-----------------------|-----------------------|
// left                   target                  right

// We've found an exact match.

// value < element, so missing value should go to the left of it

// value > element, so missing value should go to the right of it
