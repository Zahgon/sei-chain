//go:build rocksdbBackend
// +build rocksdbBackend

package mvcc

import (
	"sync"
	"sync/atomic"

	"github.com/linxGnu/grocksdb"

	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/wal"
)

const (
	TimestampSize = 8

	StorePrefixTpl     = "s/k:%s/"
	latestVersionKey   = "s/latest"
	earliestVersionKey = "s/earliest"

	// TODO: Make configurable
	ImportCommitBatchSize = 10000
	MinWALEntriesToKeep   = 1000
)

var (
	_ types.StateStore = (*Database)(nil)

	defaultWriteOpts = grocksdb.NewDefaultWriteOptions()
	defaultReadOpts  = grocksdb.NewDefaultReadOptions()
)

type VersionedChangesets struct {
	Version    int64
	Changesets []*proto.NamedChangeSet
}

type Database struct {
	storage  *grocksdb.DB
	config   config.StateStoreConfig
	cfHandle *grocksdb.ColumnFamilyHandle

	// tsLow reflects the full_history_ts_low CF value. Since pruning is done in
	// a lazy manner, we use this value to prevent reads for versions that will
	// be purged in the next compaction.
	tsLow int64

	// Earliest version for db after pruning
	earliestVersion int64
	// Latest version for db
	latestVersion atomic.Int64

	asyncWriteWG sync.WaitGroup

	// Changelog used to support async write
	streamHandler wal.ChangelogWAL

	// Pending changes to be written to the DB
	pendingChanges chan VersionedChangesets
}

func OpenDB(dataDir string, config config.StateStoreConfig) (*Database, error) {
	_ = "STUB: not implemented"
	//TODO: add a new config and check if readonly = true to support readonly mode
	return nil, nil
}

// Initialize earliest version

func (db *Database) getSlice(storeKey string, version int64, key []byte) (*grocksdb.Slice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Database) SetLatestVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (db *Database) GetLatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

// retrieveLatestVersion retrieves the latest version from the database, if not found, return 0.
func retrieveLatestVersion(storage *grocksdb.DB) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (db *Database) SetEarliestVersion(version int64, ignoreVersion bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *Database) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }

// retrieveEarliestVersion retrieves the earliest version from the database, if not found, return 0.
func retrieveEarliestVersion(storage *grocksdb.DB) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (db *Database) Has(storeKey string, version int64, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (db *Database) Get(storeKey string, version int64, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyChangesetSync apply all changesets for a single version in blocking way
func (db *Database) ApplyChangesetSync(version int64, changeset []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	// Check if version is 0 and change it to 1
	// We do this specifically since keys written as part of genesis state come in as version 0
	// But pebbledb treats version 0 as special, so apply the changeset at version 1 instead
	// Port this over to rocksdb for consistency
	return nil
}

// Update latest version in batch

// Update latest version once all writes succeed

func (db *Database) ApplyChangesetAsync(version int64, changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	// Add to pending changes
	return nil
}

// Write to WAL

func (db *Database) writeAsyncInBackground() { _ = "STUB: not implemented"; return }

// Prune attempts to prune all versions up to and including the provided version.
// This is done internally by updating the full_history_ts_low RocksDB value on
// the column families, s.t. all versions less than full_history_ts_low will be
// dropped.
//
// Note, this does NOT incur an immediate full compaction, i.e. this performs a
// lazy prune. Future compactions will honor the increased full_history_ts_low
// and trim history when possible.
func (db *Database) Prune(version int64) error {
	_ = "STUB: not implemented"
	// Defensive check: ensure database is not closed
	return nil
}

// we increment by 1 to include the provided version

func (db *Database) Iterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (db *Database) ReverseIterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

// Import loads the initial version of the state in parallel with numWorkers goroutines
// TODO: Potentially add retries instead of panics
func (db *Database) Import(version int64, ch <-chan types.SnapshotNode) error {
	_ = "STUB: not implemented"
	return nil
}

// RawIterate iterates over all keys and values for a store
// TODO: Accept list of storeKeys to export
func (db *Database) RawIterate(storeKey string, fn func(key []byte, value []byte, version int64) bool) (bool, error) {
	_ = "STUB: not implemented"
	// If store key provided, only iterate over keys with prefix
	return false, nil
}

// Set timestamp lower and upper bound to iterate over all keys in db

// Call callback fn

// newTSReadOptions returns ReadOptions used in the RocksDB column family read.
func newTSReadOptions(version int64) *grocksdb.ReadOptions { _ = "STUB: not implemented"; return nil }

func storePrefix(storeKey string) []byte { _ = "STUB: not implemented"; return nil }

func prependStoreKey(storeKey string, key []byte) []byte { _ = "STUB: not implemented"; return nil }

// copyAndFreeSlice will copy a given RocksDB slice and free it. If the slice does
// not exist, <nil> will be returned.
func copyAndFreeSlice(s *grocksdb.Slice) []byte { _ = "STUB: not implemented"; return nil }

func readOnlySlice(s *grocksdb.Slice) []byte { _ = "STUB: not implemented"; return nil }

func cloneAppend(bz []byte, tail []byte) (res []byte) { _ = "STUB: not implemented"; return nil }

func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

// Close the pending changes channel to signal the background goroutine to stop

// Wait for the async writes to finish processing all buffered items

// Close the changelog stream first

// Only set to nil after background goroutine has finished
