package flatkv

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

var _ types.Exporter = (*KVExporter)(nil)

// KVExporter exports all committed data from a read-only FlatKV store as raw
// physical key/value pairs. It uses RawGlobalIterator to walk every data DB
// (account → code → storage → legacy) and emits each row as a single
// SnapshotNode without any parsing or conversion.
//
// All emitted SnapshotNodes carry the export version and Height=0 (leaf).
//
// The caller must Close the exporter when done.
type KVExporter struct {
	store   *CommitStore
	version int64
	iter    Iterator
}

func NewKVExporter(store *CommitStore, version int64) *KVExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *KVExporter) Next() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *KVExporter) Close() error { _ = "STUB: not implemented"; return nil }
