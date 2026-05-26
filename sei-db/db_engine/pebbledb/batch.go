package pebbledb

import (
	"github.com/cockroachdb/pebble/v2"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// pebbleBatch wraps a Pebble batch for atomic writes.
// Important: Callers must call Close() after Commit() to release batch resources,
// even if Commit() succeeds. Failure to Close() will leak memory.
type pebbleBatch struct {
	b *pebble.Batch
}

var _ types.Batch = (*pebbleBatch)(nil)

func (p *pebbleDB) NewBatch() types.Batch { _ = "STUB: not implemented"; return *new(types.Batch) }

func (pb *pebbleBatch) Set(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (pb *pebbleBatch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (pb *pebbleBatch) Commit(opts types.WriteOptions) error { _ = "STUB: not implemented"; return nil }

func (pb *pebbleBatch) Len() int { _ = "STUB: not implemented"; return 0 }

func (pb *pebbleBatch) Reset() { _ = "STUB: not implemented"; return }

func (pb *pebbleBatch) Close() error { _ = "STUB: not implemented"; return nil }
