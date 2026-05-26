package testdb

import (
	"context"
)

const (
	// Size of the channel buffer between traversal goroutine and iterator. Using an unbuffered
	// channel causes two context switches per item sent, while buffering allows more work per
	// context switch. Tuned with benchmarks.
	chBufferSize = 64
)

// memDBIterator is a memDB iterator.
type memDBIterator struct {
	ch     <-chan *item
	cancel context.CancelFunc
	item   *item
	start  []byte
	end    []byte
	useMtx bool
}

var _ Iterator = (*memDBIterator)(nil)

// newMemDBIterator creates a new memDBIterator.
func newMemDBIterator(db *MemDB, start []byte, end []byte, reverse bool) *memDBIterator {
	_ = "STUB: not implemented"
	return nil
}

func newMemDBIteratorMtxChoice(db *MemDB, start []byte, end []byte, reverse bool, useMtx bool) *memDBIterator {
	_ = "STUB: not implemented"
	return nil
}

// Because we use [start, end) for reverse ranges, while btree uses (start, end], we need
// the following variables to handle some reverse iteration conditions ourselves.

// must handle this specially, since nil is considered less than anything else

// abort after start, since we use [start, end) while btree uses (start, end]

// skip end and abort after start, since we use [start, end) while btree uses (start, end]

// prime the iterator with the first value, if any

// Close implements Iterator.
func (i *memDBIterator) Close() error { _ = "STUB: not implemented"; return nil }

// drain channel

// Domain implements Iterator.
func (i *memDBIterator) Domain() ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil,

		// Valid implements Iterator.
		nil
}

func (i *memDBIterator) Valid() bool { _ = "STUB: not implemented"; return false }

// Next implements Iterator.
func (i *memDBIterator) Next() { _ = "STUB: not implemented"; return }

// Error implements Iterator.
func (i *memDBIterator) Error() error {
	_ = "STUB: not implemented"
	// famous last words
	return nil
}

// Key implements Iterator.
func (i *memDBIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// Value implements Iterator.
func (i *memDBIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (i *memDBIterator) assertIsValid() { _ = "STUB: not implemented"; return }
