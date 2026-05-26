package gaskv

import (
	"io"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

type IStoreTracer interface {
	Get([]byte, []byte, string, time.Duration)
	Has([]byte, string, time.Duration)
	Set([]byte, []byte, string, time.Duration)
	Delete([]byte, string, time.Duration)
	StartIterator([]byte, []byte, bool, string, time.Duration) int
	RecordIteratorValue(int, []byte, []byte, string)
	RecordIteratorNext(int, string, time.Duration)
	DerivePrestateToJson() []byte
	Clear()
}

// traceStart returns time.Now() when tracer is non-nil, or the zero Time
// otherwise. Keeps time.Now off the gaskv hot path when no tracer is attached.
func traceStart(tracer IStoreTracer) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

var _ types.KVStore = &Store{}

// Store applies gas tracking to an underlying KVStore. It implements the
// KVStore interface.
type Store struct {
	gasMeter   types.GasMeter
	gasConfig  types.GasConfig
	parent     types.KVStore
	moduleName string
	tracer     IStoreTracer
}

// NewStore returns a reference to a new GasKVStore.
func NewStore(parent types.KVStore, gasMeter types.GasMeter, gasConfig types.GasConfig, moduleName string, tracer IStoreTracer) *Store {
	_ = "STUB: not implemented"
	return nil
}

// Implements Store.
func (gs *Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *new(types.StoreType)
}

func (gs *Store) GetWorkingHash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Implements KVStore.
func (gs *Store) Get(key []byte) (value []byte) { _ = "STUB: not implemented"; return nil }

// TODO overflow-safe math?

// Implements KVStore.
func (gs *Store) Set(key []byte, value []byte) { _ = "STUB: not implemented"; return }

// TODO overflow-safe math?

// Implements KVStore.
func (gs *Store) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

// Implements KVStore.
func (gs *Store) Delete(key []byte) { _ = "STUB: not implemented"; return }

// charge gas to prevent certain attack vectors even though space is being freed

// Iterator implements the KVStore interface. It returns an iterator which
// incurs a flat gas cost for seeking to the first key/value pair and a variable
// gas cost based on the current value's length if the iterator is valid.
func (gs *Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// ReverseIterator implements the KVStore interface. It returns a reverse
// iterator which incurs a flat gas cost for seeking to the first key/value pair
// and a variable gas cost based on the current value's length if the iterator
// is valid.
func (gs *Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// Implements KVStore.
func (gs *Store) CacheWrap(_ types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements the KVStore interface.
func (gs *Store) CacheWrapWithTrace(_ types.StoreKey, _ io.Writer, _ types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

func (gs *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// if there is a panic, we close the iterator then reraise

func (gs *Store) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

func (gs *Store) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (gs *Store) GetAllKeyStrsInRange(start, end []byte) (res []string) {
	_ = "STUB: not implemented"
	return nil
}

type gasIterator struct {
	gasMeter   types.GasMeter
	gasConfig  types.GasConfig
	parent     types.Iterator
	moduleName string
	tracer     IStoreTracer
	iteratorID int
}

func newGasIterator(gasMeter types.GasMeter, gasConfig types.GasConfig, parent types.Iterator, moduleName string, tracer IStoreTracer) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// Implements Iterator.
func (gi *gasIterator) Domain() (start []byte, end []byte) {
	_ = "STUB: not implemented"
	return nil,

		// Implements Iterator.
		nil
}

func (gi *gasIterator) Valid() bool { _ = "STUB: not implemented"; return false }

// Next implements the Iterator interface. It seeks to the next key/value pair
// in the iterator. It incurs a flat gas cost for seeking and a variable gas
// cost based on the current value's length if the iterator is valid.
func (gi *gasIterator) Next() { _ = "STUB: not implemented"; return }

// Key implements the Iterator interface. It returns the current key and it does
// not incur any gas cost.
func (gi *gasIterator) Key() (key []byte) { _ = "STUB: not implemented"; return nil }

// Value implements the Iterator interface. It returns the current value and it
// does not incur any gas cost.
func (gi *gasIterator) Value() (value []byte) { _ = "STUB: not implemented"; return nil }

// Implements Iterator.
func (gi *gasIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Error delegates the Error call to the parent iterator.
func (gi *gasIterator) Error() error { _ = "STUB: not implemented"; return nil }

// consumeSeekGas consumes on each iteration step a flat gas cost and a variable gas cost
// based on the current value's length.
func (gi *gasIterator) consumeSeekGas() { _ = "STUB: not implemented"; return }
