package mvcc

import (
	"github.com/cockroachdb/pebble/v2"
)

type Batch struct {
	storage    *pebble.DB
	version    int64
	ops        []batchOp
	descending bool
}

type batchOp struct {
	key    []byte
	value  []byte
	delete bool
}

// NewBatch creates a new Batch using the supplied MVCC encoding mode.
func NewBatch(storage *pebble.DB, version int64, descending bool) (*Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Batch) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *Batch) Reset() { _ = "STUB: not implemented"; return }

func (b *Batch) set(storeKey string, tombstone int64, key, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Batch) Set(storeKey string, key, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Batch) Delete(storeKey string, key []byte) error { _ = "STUB: not implemented"; return nil }

func (b *Batch) Write() error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // block heights are non-negative and fit in int64

// For writing kv pairs in any order of version
type RawBatch struct {
	storage    *pebble.DB
	ops        []batchOp
	descending bool
}

// NewRawBatch creates a new RawBatch using the supplied MVCC encoding mode.
func NewRawBatch(storage *pebble.DB, descending bool) (*RawBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *RawBatch) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *RawBatch) Reset() { _ = "STUB: not implemented"; return }

func (b *RawBatch) set(storeKey string, tombstone int64, key, value []byte, version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *RawBatch) Set(storeKey string, key, value []byte, version int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *RawBatch) Delete(storeKey string, key []byte, version int64) error {
	_ = "STUB: not implemented"
	return nil
}

// HardDelete physically removes the key by encoding it with the batch’s version
// and calling the underlying pebble.Batch.Delete.
func (b *Batch) HardDelete(storeKey string, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *RawBatch) Write() error { _ = "STUB: not implemented"; return nil }

// writeBatchOps applies ops to a new pebble batch in sorted order, records
// otel metrics, and commits. The optional beforeCommit hook runs on the
// pebble batch right before commit (used by Batch.Write to stamp the
// latest-version metadata key).
func writeBatchOps(storage *pebble.DB, ops []batchOp, beforeCommit func(*pebble.Batch) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func sortBatchOps(ops []batchOp) { _ = "STUB: not implemented"; return }
