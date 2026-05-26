//go:build rocksdbBackend
// +build rocksdbBackend

package mvcc

import (
	"github.com/linxGnu/grocksdb"
)

type Batch struct {
	version  int64
	ts       [TimestampSize]byte
	storage  *grocksdb.DB
	cfHandle *grocksdb.ColumnFamilyHandle
	batch    *grocksdb.WriteBatch
}

// NewBatch creates a new versioned batch used for batch writes. The caller
// must ensure to call Write() on the returned batch to commit the changes and to
// destroy the batch when done.
func NewBatch(db *Database, version int64) Batch { _ = "STUB: not implemented"; return *new(Batch) }

func (b Batch) Size() int { _ = "STUB: not implemented"; return 0 }

func (b Batch) Reset() { _ = "STUB: not implemented"; return }

func (b Batch) Set(storeKey string, key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (b Batch) Delete(storeKey string, key []byte) error { _ = "STUB: not implemented"; return nil }

func (b Batch) Write() error { _ = "STUB: not implemented"; return nil }
