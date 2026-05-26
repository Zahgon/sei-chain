package memiavl

import (
	dbm "github.com/tendermint/tm-db"
)

var _ dbm.Iterator = (*Iterator)(nil)

type Iterator struct {
	// domain of iteration, end is exclusive
	start, end []byte
	ascending  bool
	zeroCopy   bool

	// cache the next key-value pair
	key, value []byte

	valid bool

	stack []Node
}

func NewIterator(start, end []byte, ascending bool, root Node, zeroCopy bool) *Iterator {
	_ = "STUB: not implemented"
	return nil
}

// cache the first key-value

func (iter *Iterator) Domain() ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil,

		// Valid implements dbm.Iterator.
		nil
}

func (iter *Iterator) Valid() bool {
	_ = "STUB: not implemented"

	// Error implements dbm.Iterator
	return false
}

func (iter *Iterator) Error() error {
	_ = "STUB: not implemented"

	// Key implements dbm.Iterator
	return nil
}

func (iter *Iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// Value implements dbm.Iterator
func (iter *Iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// Next implements dbm.Iterator
func (iter *Iterator) Next() { _ = "STUB: not implemented"; return }

// pop node

// push children to stack

// Close implements dbm.Iterator
func (iter *Iterator) Close() error { _ = "STUB: not implemented"; return nil }
