package tracekv

import (
	"io"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

const (
	writeOp     operation = "write"
	readOp      operation = "read"
	deleteOp    operation = "delete"
	iterKeyOp   operation = "iterKey"
	iterValueOp operation = "iterValue"
)

type (
	// Store implements the KVStore interface with tracing enabled.
	// Operations are traced on each core KVStore call and written to the
	// underlying io.writer.
	//
	// TODO: Should we use a buffered writer and implement Commit on
	// Store?
	Store struct {
		parent  types.KVStore
		writer  io.Writer
		context types.TraceContext
	}

	// operation represents an IO operation
	operation string

	// traceOperation implements a traced KVStore operation
	traceOperation struct {
		Operation operation              `json:"operation"`
		Key       string                 `json:"key"`
		Value     string                 `json:"value"`
		Metadata  map[string]interface{} `json:"metadata"`
	}
)

// NewStore returns a reference to a new traceKVStore given a parent
// KVStore implementation and a buffered writer.
func NewStore(parent types.KVStore, writer io.Writer, tc types.TraceContext) *Store {
	_ = "STUB: not implemented"
	return nil
}

func (tkv *Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Get implements the KVStore interface. It traces a read operation and
// delegates a Get call to the parent KVStore.
func (tkv *Store) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Set implements the KVStore interface. It traces a write operation and
// delegates the Set call to the parent KVStore.
func (tkv *Store) Set(key []byte, value []byte) { _ = "STUB: not implemented"; return }

// Delete implements the KVStore interface. It traces a write operation and
// delegates the Delete call to the parent KVStore.
func (tkv *Store) Delete(key []byte) { _ = "STUB: not implemented"; return }

// Has implements the KVStore interface. It delegates the Has call to the
// parent KVStore.
func (tkv *Store) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

// Iterator implements the KVStore interface. It delegates the Iterator call
// the to the parent KVStore.
func (tkv *Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// ReverseIterator implements the KVStore interface. It delegates the
// ReverseIterator call the to the parent KVStore.
func (tkv *Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// iterator facilitates iteration over a KVStore. It delegates the necessary
// calls to it's parent KVStore.
func (tkv *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

type traceIterator struct {
	parent  types.Iterator
	writer  io.Writer
	context types.TraceContext
}

func newTraceIterator(w io.Writer, parent types.Iterator, tc types.TraceContext) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// Domain implements the Iterator interface.
func (ti *traceIterator) Domain() (start []byte, end []byte) {
	_ = "STUB: not implemented"
	return nil,

		// Valid implements the Iterator interface.
		nil
}

func (ti *traceIterator) Valid() bool { _ = "STUB: not implemented"; return false }

// Next implements the Iterator interface.
func (ti *traceIterator) Next() {
	_ = "STUB: not implemented"

	// Key implements the Iterator interface.
	return
}

func (ti *traceIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// Value implements the Iterator interface.
func (ti *traceIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// Close implements the Iterator interface.
func (ti *traceIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Error delegates the Error call to the parent iterator.
func (ti *traceIterator) Error() error { _ = "STUB: not implemented"; return nil }

// GetStoreType implements the KVStore interface. It returns the underlying
// KVStore type.
func (tkv *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

// Parent returns the wrapped KVStore (without tracing). Used to reach proof-capable stores.
func (tkv *Store) Parent() types.KVStore {
	_ = "STUB: not implemented"

	// CacheWrap implements the KVStore interface. It panics because a Store
	// cannot be branched.
	return *new(types.KVStore)
}

func (tkv *Store) CacheWrap(_ types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements the KVStore interface. It panics as a
// Store cannot be branched.
func (tkv *Store) CacheWrapWithTrace(_ types.StoreKey, _ io.Writer, _ types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

func (tkv *Store) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

func (tkv *Store) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (tkv *Store) GetAllKeyStrsInRange(start, end []byte) []string {
	_ = "STUB: not implemented"
	return nil
}

// writeOperation writes a KVStore operation to the underlying io.Writer as
// JSON-encoded data where the key/value pair is base64 encoded.
func writeOperation(w io.Writer, op operation, tc types.TraceContext, key, value []byte) {
	_ = "STUB: not implemented"
	return
}
