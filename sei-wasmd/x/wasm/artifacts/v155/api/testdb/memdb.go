package testdb

import (
	"sync"

	"github.com/google/btree"
)

const (
	// The approximate number of items and children per B-tree node. Tuned with benchmarks.
	bTreeDegree = 32
)

// item is a btree.Item with byte slices as keys and values
type item struct {
	key   []byte
	value []byte
}

// Less implements btree.Item.
func (i *item) Less(other btree.Item) bool {
	_ = "STUB: not implemented"
	// this considers nil == []byte{}, but that's ok since we handle nil endpoints
	// in iterators specially anyway
	return false
}

// newKey creates a new key item.
func newKey(key []byte) *item { _ = "STUB: not implemented"; return nil }

// newPair creates a new pair item.
func newPair(key, value []byte) *item { _ = "STUB: not implemented"; return nil }

// MemDB is an in-memory database backend using a B-tree for storage.
//
// For performance reasons, all given and returned keys and values are pointers to the in-memory
// database, so modifying them will cause the stored values to be modified as well. All DB methods
// already specify that keys and values should be considered read-only, but this is especially
// important with MemDB.
type MemDB struct {
	mtx   sync.RWMutex
	btree *btree.BTree
}

// NewMemDB creates a new in-memory database.
func NewMemDB() *MemDB { _ = "STUB: not implemented"; return nil }

// Get implements DB.
func (db *MemDB) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Has implements DB.
func (db *MemDB) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Set implements DB.
func (db *MemDB) Set(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

// set sets a value without locking the mutex.
func (db *MemDB) set(key []byte, value []byte) { _ = "STUB: not implemented"; return }

// SetSync implements DB.
func (db *MemDB) SetSync(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

// Delete implements DB.
func (db *MemDB) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// delete deletes a key without locking the mutex.
func (db *MemDB) delete(key []byte) { _ = "STUB: not implemented"; return }

// DeleteSync implements DB.
func (db *MemDB) DeleteSync(key []byte) error { _ = "STUB: not implemented"; return nil }

// Close implements DB.
func (db *MemDB) Close() error {
	_ = "STUB: not implemented"
	// Close is a noop since for an in-memory database, we don't have a destination to flush
	// contents to nor do we want any data loss on invoking Close().
	// See the discussion in https://github.com/tendermint/tendermint/libs/pull/56
	return nil
}

// Print implements DB.
func (db *MemDB) Print() error { _ = "STUB: not implemented"; return nil }

// Stats implements DB.
func (db *MemDB) Stats() map[string]string { _ = "STUB: not implemented"; return nil }

// Iterator implements DB.
// Takes out a read-lock on the database until the iterator is closed.
func (db *MemDB) Iterator(start, end []byte) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}

// ReverseIterator implements DB.
// Takes out a read-lock on the database until the iterator is closed.
func (db *MemDB) ReverseIterator(start, end []byte) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}

// IteratorNoMtx makes an iterator with no mutex.
func (db *MemDB) IteratorNoMtx(start, end []byte) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}

// ReverseIteratorNoMtx makes an iterator with no mutex.
func (db *MemDB) ReverseIteratorNoMtx(start, end []byte) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}
