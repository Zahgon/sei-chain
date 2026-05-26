package avail

// queue is a collection of objects of type T, indexed by type I in range [first,next).
// Supports pushing new items to the back and popping items from the front.
type queue[I ~uint64, T any] struct {
	q     map[I]T
	first I
	next  I
}

func newQueue[I ~uint64, T any]() *queue[I, T] { _ = "STUB: not implemented"; return nil }

func (q *queue[I, T]) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (q *queue[I, T]) pushBack(t T) { _ = "STUB: not implemented"; return }

func (q *queue[I, T]) prune(newFirst I) { _ = "STUB: not implemented"; return }
