package types

import (
	"sync"
	"time"
)

// Per-tx tracer caps. Bound memory and wire size so a pathological tx (huge
// iterator scans, many opened iterators) can't make the profile response
// unbounded. Values picked to comfortably cover normal EVM txs while
// capping worst-case overhead at a few MB per module.
const (
	maxStoreTraceIterators    = 16
	maxStoreTraceIteratorKeys = 64
)

// StoreTracer collects every KVStore access (Get/Has/Set/Delete/iterator)
// performed under a debug_traceTransactionProfile call, grouped by module.
// One tracer is attached to the sdk.Context for the duration of the traced
// tx; callers must treat it as single-tx-scoped. All methods are safe for
// concurrent use.
type StoreTracer struct {
	Modules        map[string]*ModuleTrace
	nextIteratorID int
	mu             *sync.Mutex
}

// ModuleTrace holds every access event for a single module within a trace,
// plus a per-iterator roll-up.
type ModuleTrace struct {
	Accesses        []Access
	Iterators       []*IteratorTrace
	iteratorIndexBy map[int]int
}

// IteratorTrace aggregates one opened iterator: its bounds, direction, the
// keys it surfaced to the tx (capped at maxStoreTraceIteratorKeys; Truncated
// flags overflow), and cumulative Next() count + time.
type IteratorTrace struct {
	Start         []byte
	End           []byte
	Ascending     bool
	Keys          [][]byte
	NextCount     int
	DurationNanos int64
	Truncated     bool
}

// Access is a single access event in a module's chronological log. Value is
// only populated for Get/Set/IteratorValue; other ops leave it nil.
type Access struct {
	Op            OpType
	Key           []byte
	Value         []byte
	DurationNanos int64
}

// OpType tags an Access with the operation the tx performed.
type OpType int

const (
	Get OpType = iota
	Has
	Set
	Delete
	IteratorOpen
	IteratorNext
	IteratorValue
)

func (o OpType) String() string { _ = "STUB: not implemented"; return "" }

// NewStoreTracer returns an empty StoreTracer ready to record per-module
// access events for a single debug_traceTransactionProfile call.
func NewStoreTracer() *StoreTracer { _ = "STUB: not implemented"; return nil }

func (st *StoreTracer) Get(key []byte, value []byte, module string, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (st *StoreTracer) Set(key []byte, value []byte, module string, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (st *StoreTracer) Has(key []byte, module string, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (st *StoreTracer) Delete(key []byte, module string, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// StartIterator records the opening of an iterator over [start, end) and
// allocates a tracer-scoped iteratorID the caller uses to tag subsequent
// Next/Value events. Past maxStoreTraceIterators the IteratorTrace record is
// dropped (the access-log event is still recorded) and the returned ID lets
// later calls no-op gracefully.
func (st *StoreTracer) StartIterator(start, end []byte, ascending bool, module string, duration time.Duration) int {
	_ = "STUB: not implemented"
	return 0
}

// RecordIteratorValue records that the tx read the current key/value from
// the iterator identified by iteratorID. Beyond maxStoreTraceIteratorKeys the
// iterator is flagged Truncated and further keys are dropped from the
// per-iterator sample (the access-log event is still recorded).
func (st *StoreTracer) RecordIteratorValue(iteratorID int, key []byte, value []byte, module string) {
	_ = "STUB: not implemented"
	return
}

// RecordIteratorNext records a Next() advance on the iterator identified by
// iteratorID, adding to its cumulative step count and stepping time.
func (st *StoreTracer) RecordIteratorNext(iteratorID int, module string, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (st *StoreTracer) getOrSetModuleTrace(module string) (mt *ModuleTrace) {
	_ = "STUB: not implemented"
	return nil
}

func (st *StoreTracer) recordAccess(module string, access Access) {
	_ = "STUB: not implemented"
	return
}

// Clear resets the tracer to its empty state so a single StoreTracer can be
// reused across successive trace requests on the same connection.
func (st *StoreTracer) Clear() { _ = "STUB: not implemented"; return }

type OperationSummary struct {
	Count      int   `json:"count"`
	TotalNanos int64 `json:"totalNanos"`
}

type StoreTraceDump struct {
	Modules map[string]ModuleTraceDump  `json:"modules"`
	Stats   map[string]OperationSummary `json:"stats,omitempty"`
}

type ModuleTraceDump struct {
	Reads     map[string]string           `json:"reads"`
	Has       []string                    `json:"has"`
	Stats     map[string]OperationSummary `json:"stats,omitempty"`
	Iterators []IteratorTraceDump         `json:"iterators,omitempty"`
}

type IteratorTraceDump struct {
	Start      string   `json:"start,omitempty"`
	End        string   `json:"end,omitempty"`
	Ascending  bool     `json:"ascending"`
	Keys       []string `json:"keys,omitempty"`
	NextCount  int      `json:"nextCount"`
	TotalNanos int64    `json:"totalNanos"`
	Truncated  bool     `json:"truncated,omitempty"`
}

// Dump materializes the tracer's accumulated per-module accesses into a
// wire-shaped StoreTraceDump. Reads that were later overwritten by a Set or
// Delete during the same tx are excluded so the Reads map reflects the
// pre-state the tx observed.
func (st *StoreTracer) Dump() StoreTraceDump {
	_ = "STUB: not implemented"
	return *new(StoreTraceDump)
}

func (st *StoreTracer) dumpLocked() StoreTraceDump {
	_ = "STUB: not implemented"
	return *new(StoreTraceDump)
}

// any read for key XYZ after a Set/Delete to XYZ is discarded
// because the result doesn't represent prestate.

func updateSummary(stats map[string]OperationSummary, op OpType, durationNanos int64) {
	_ = "STUB: not implemented"
	return
}

// DerivePrestateToJson returns a JSON encoding of the current trace state,
// used by debug_traceTransaction to attach AppState to the response.
func (st *StoreTracer) DerivePrestateToJson() []byte { _ = "STUB: not implemented"; return nil }
