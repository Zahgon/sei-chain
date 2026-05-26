//go:build rocksdbBackend
// +build rocksdbBackend

package mvcc

import (
	"github.com/linxGnu/grocksdb"
)

const (
	// CFNameStateStorage defines the RocksDB column family name for versioned state
	// storage.
	CFNameStateStorage = "state_storage"

	// CFNameDefault defines the RocksDB column family name for the default column.
	CFNameDefault = "default"
)

// NewRocksDBOpts returns the options used for the RocksDB column family for use
// in state storage.
//
// FIXME: We do not enable dict compression for SSTFileWriter, because otherwise
// the file writer won't report correct file size.
// Ref: https://github.com/facebook/rocksdb/issues/11146
func NewRocksDBOpts(sstFileWriter bool) *grocksdb.Options { _ = "STUB: not implemented"; return nil }

// block based table options

// 1G block cache

// Improve sst file creation speed: compaction or sst file writer.

// compression options at bottommost level

// 110k

// OpenRocksDB opens a RocksDB database connection for versioned reading and writing.
// It also returns a column family handle for versioning using user-defined timestamps.
// The default column family is used for metadata, specifically key/value pairs
// that are stored on another column family named with "state_storage", which has
// user-defined timestamp enabled.
func OpenRocksDB(dataDir string) (*grocksdb.DB, *grocksdb.ColumnFamilyHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// OpenRocksDBAndTrimHistory opens a RocksDB handle similar to `OpenRocksDB`,
// but it also trims the versions newer than target one, such that it can be used
// for rollback.
func OpenRocksDBAndTrimHistory(dataDir string, version int64) (*grocksdb.DB, *grocksdb.ColumnFamilyHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
