package multiversion

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// tracked iterator is a wrapper around an existing iterator to track the iterator progress and monitor which keys are iterated.
type trackedIterator struct {
	types.Iterator

	iterateset *iterationTracker
}

func NewTrackedIterator(iter types.Iterator, iterationTracker *iterationTracker) *trackedIterator {
	_ = "STUB: not implemented"
	return nil
}

func (ti *trackedIterator) Valid() bool { _ = "STUB: not implemented"; return false }

// if no longer valid, remove the early stop key since we reached end of range

// Key calls the iterator.Key() and adds the key to the iterateset, then returns the key from the iterator
func (ti *trackedIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// add key to the tracker

// Value calls the iterator.Key() and adds the key to the iterateset, then returns the value from the iterator
func (ti *trackedIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// add key to the tracker

func (ti *trackedIterator) Next() {
	_ = "STUB: not implemented"
	// add current key to the tracker
	return
}

// call next
