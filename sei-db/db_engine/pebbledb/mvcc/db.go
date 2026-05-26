package mvcc

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/cockroachdb/pebble/v2"

	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/wal"
)

const (
	VersionSize = 8

	PrefixStore        = "s/k:"
	LenPrefixStore     = 4
	StorePrefixTpl     = "s/k:%s/" // s/k:<storeKey>
	latestVersionKey   = "s/_latest"
	earliestVersionKey = "s/_earliest"
	// descendingMVCCMarkerKey flags that the DB was initialized with the
	// descending-version MVCC encoding. Its absence on a populated DB means
	// the data was written by the legacy ascending-version build and is not
	// safe to read with this code.
	descendingMVCCMarkerKey = "s/_mvcc_descending"
	tombstoneVal            = "TOMBSTONE"

	// TODO: Make configurable
	ImportCommitBatchSize = 10000
	PruneCommitBatchSize  = 50
	DeleteCommitBatchSize = 50
	MinWALEntriesToKeep   = 1000
)

var (
	_ types.StateStore = (*Database)(nil)

	defaultWriteOpts = pebble.NoSync
)

type Database struct {
	storage      *pebble.DB
	asyncWriteWG sync.WaitGroup
	config       config.StateStoreConfig
	// Earliest version for db after pruning
	earliestVersion atomic.Int64
	// Latest version for db
	latestVersion atomic.Int64
	// descending indicates whether this DB uses descending-version MVCC
	// encoding (fresh DBs created by this build) or the legacy
	// ascending-version encoding (DBs created by the previous build). The
	// mode is detected on open and is immutable for the lifetime of the
	// Database.
	descending bool

	// Map of module to when each was last updated
	// Used in pruning to skip over stores that have not been updated recently
	storeKeyDirty sync.Map

	// Changelog used to support async write
	streamHandler wal.ChangelogWAL

	// Pending changes to be written to the DB
	pendingChanges chan VersionedChangesets

	// Cancel function for background metrics collection
	metricsCancel context.CancelFunc
}

type VersionedChangesets struct {
	Version    int64
	Changesets []*proto.NamedChangeSet
	Done       chan struct{} // non-nil for barrier: closed when this entry is processed
}

func OpenDB(dataDir string, config config.StateStoreConfig) (types.StateStore, error) {
	_ = "STUB: not implemented"
	return *new(types.StateStore), nil
}

// Select comparer based on config. Note: UseDefaultComparer is NOT backwards compatible
// with existing databases created with MVCCComparer - Pebble will refuse to open due to
// comparer name mismatch. Only use UseDefaultComparer for NEW databases.

// TODO: Delete once we remove support

// FormatMajorVersion is pinned to a specific version to prevent accidental
// breaking changes when updating the pebble dependency. Using FormatNewest
// would cause the on-disk format to silently upgrade when pebble is updated,
// making the database incompatible with older software versions.
// When upgrading this version, ensure it's an intentional, documented change.

// 64 MB

// Configure L0 with explicit settings
// 32 KB
// 256 KB

// Configure L1+ levels, inheriting from previous level

// 32 KB
// 256 KB

// Disable bloom filter at bottommost level (L6) - bloom filters are less useful
// at the bottom level since most data lives there and false positive rate is low

//TODO: add a new config and check if readonly = true to support readonly mode

// Initialize earliest version

// Initialize latest version

// Start background metrics collection

func (db *Database) Close() error {
	_ = "STUB: not implemented"
	// Stop background metrics collection
	return nil
}

// First, stop accepting new pending changes and drain the worker

// Wait for the async writes to finish

// Now close the WAL stream

// Make Close idempotent: Pebble panics if Close is called twice.

// mvccEncode encodes a key with the MVCC version encoding matching this
// Database's on-disk mode.
func (db *Database) mvccEncode(key []byte, version int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// decodeVersion decodes an on-disk MVCC version using the encoding matching
// this Database's mode.
func (db *Database) decodeVersion(vBz []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// PebbleMetrics returns the underlying Pebble DB metrics for observability (e.g. compaction/flush counts).
// Returns nil if the database is closed.
func (db *Database) PebbleMetrics() *pebble.Metrics { _ = "STUB: not implemented"; return nil }

func (db *Database) SetLatestVersion(version int64) error { _ = "STUB: not implemented"; return nil }

func (db *Database) GetLatestVersion() int64 { _ = "STUB: not implemented"; return 0 }

// detectMVCCMode inspects the DB to determine which MVCC encoding to use.
//
//   - If the descendingMVCCMarkerKey sentinel is present, the DB was created
//     by this build and is in descending mode.
//   - If the marker is absent but latestVersionKey is present, the DB was
//     populated by the legacy ascending-version build. We open it in
//     ascending mode without writing the marker (legacy DBs stay unmarked
//     forever).
//   - If both markers are absent the DB is fresh; we write the descending
//     marker and return descending mode.
func detectMVCCMode(db *pebble.DB) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Legacy DB: no marker, has data. Open in ascending mode.

// Fresh DB: mark it and use descending mode.

func retrieveLatestVersion(db *pebble.DB) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (db *Database) SetEarliestVersion(version int64, ignoreVersion bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *Database) GetEarliestVersion() int64 { _ = "STUB: not implemented"; return 0 }

// Retrieves earliest version from db, if not found, return 0
func retrieveEarliestVersion(db *pebble.DB) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// retrieveVersionKey reads a little-endian uint64 version from the given
// metadata key. Returns 0 when the key is absent (fresh DB).
func retrieveVersionKey(db *pebble.DB, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Has dispatches between descending- and ascending-mode implementations
// depending on the on-disk encoding detected at open time.
func (db *Database) Has(storeKey string, version int64, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Get dispatches between descending- and ascending-mode implementations
// depending on the on-disk encoding detected at open time.
func (db *Database) Get(storeKey string, targetVersion int64, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *Database) ApplyChangesetSync(version int64, changeset []*proto.NamedChangeSet) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

// Check if version is 0 and change it to 1
// We do this specifically since keys written as part of genesis state come in as version 0
// But pebbledb treats version 0 as special, so apply the changeset at version 1 instead

// Create batch and persist latest version in the batch

// Mark the store as updated

// Update latest version after all writes succeed (only if higher to avoid lowering it when writing out of order)

func (db *Database) ApplyChangesetAsync(version int64, changesets []*proto.NamedChangeSet) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

// Record pending queue depth

// Write to WAL

// Add to pending changes first

func (db *Database) writeAsyncInBackground() { _ = "STUB: not implemented"; return }

// WaitForPendingWrites waits for all pending writes to be processed
func (db *Database) WaitForPendingWrites() { _ = "STUB: not implemented"; return }

// Prune dispatches between descending- and ascending-mode implementations
// depending on the on-disk encoding detected at open time.
func (db *Database) Prune(version int64) error { _ = "STUB: not implemented"; return nil }

// Iterator dispatches between descending- and ascending-mode implementations
// depending on the on-disk encoding detected at open time.
func (db *Database) Iterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

// ReverseIterator dispatches between descending- and ascending-mode
// implementations depending on the on-disk encoding detected at open time.
func (db *Database) ReverseIterator(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

// ---------------------------------------------------------------------------
// Descending-mode implementation (the fast path used by DBs created by this
// build). Versions of a logical key sort newest-first on disk, so Pebble's
// First() / SeekGE lands directly on the latest visible version without
// iterating older ones. The ascending-mode counterparts live in
// db_ascending.go for legacy DBs.
// ---------------------------------------------------------------------------

func (db *Database) hasDescending(storeKey string, version int64, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (db *Database) getDescending(storeKey string, targetVersion int64, key []byte) (_ []byte, _err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pruneDescending attempts to prune all versions up to and including the current version
// Get the range of keys, manually iterate over them and delete them
// We add a heuristic to skip over a module's keys during pruning if it hasn't been updated
// since the last time pruning occurred.
// NOTE: There is a rare case when a module's keys are skipped during pruning even though
// it has been updated. This occurs when that module's keys are updated in between pruning runs, the node after is restarted.
// This is not a large issue given the next time that module is updated, it will be properly pruned thereafter.
func (db *Database) pruneDescending(version int64) (_err error) {
	_ = "STUB: not implemented"
	// Defensive check: ensure database is not closed
	return nil
}

// we increment by 1 to include the provided version

// Ignore metadata entries during pruning

// Store current key and version

// XXX: This should never happen given we skip the metadata keys.

// For every new module visited, check to see last time it was updated

// Skip a store's keys if version it was last updated is less than last prune height

// Reset per-logical-key state when the logical key changes.

// Fast path: under descending encoding, versions of a key are stored
// newest-first. When the newest real version is above the prune
// height, seek directly to the first version <= prune height for
// this key instead of iterating through every above-prune version.

// Descending iteration: for a given logical key we see newest→oldest.
// Versions > prune height are always kept. For versions <= prune
// height, keep only the newest one when KeepLastVersion is true;
// delete every other such version.

// Commit any leftover delete ops in batch

func (db *Database) iteratorDescending(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (db *Database) reverseIteratorDescending(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func getMVCCSliceDescending(db *pebble.DB, storeKey string, key []byte, version int64) (_ []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeMVCCEntryDescending validates that the iterator's current entry
// belongs to prefixedKey at a version <= target and returns a safe copy of the
// value. Assumes descending version encoding.
func decodeMVCCEntryDescending(rawIterKey, rawIterValue, prefixedKey []byte, version int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func visibleValueAtVersionDescending(prefixedVal []byte, targetVersion int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func iteratorUpperBoundForStoreDescending(storeKey string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func iteratorUpperBoundForLogicalKeyDescending(key []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Taken from pebbledb prefix upper bound
// Returns smallest key strictly greater than the prefix
func prefixEnd(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// Import loads the initial version of the state in parallel with numWorkers goroutines
// TODO: Potentially add retries instead of panics
func (db *Database) Import(version int64, ch <-chan types.SnapshotNode) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

// RawIterate iterates over all keys and values for a store
func (db *Database) RawIterate(storeKey string, fn func(key []byte, value []byte, version int64) bool) (bool, error) {
	_ = "STUB: not implemented"
	// Iterate through all keys and values for a store
	return false, nil
}

// Ignore metadata entries

// Store current key and version

// Only iterate through module

// Parse prefix out of the key

// Decode the value

// Call callback fn

func (db *Database) DeleteKeysAtVersion(module string, version int64) error {
	_ = "STUB: not implemented"
	return nil
}

// stop iteration on error

// Commit any remaining deletions.

func isMetadataKey(key []byte) bool { _ = "STUB: not implemented"; return false }

func storePrefix(storeKey string) []byte { _ = "STUB: not implemented"; return nil }

func prependStoreKey(storeKey string, key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Parses store from key with format "s/k:{store}/..."
func parseStoreKey(key []byte) (string, error) {
	_ = "STUB: not implemented"
	// Convert byte slice to string only once
	return "", nil
}

// Find the first occurrence of "/" after the prefix

// Return the substring between the prefix and the first "/"

func valTombstoned(value []byte) bool { _ = "STUB: not implemented"; return false }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC value.

// If the tombstone suffix is empty, we consider this a zero value and thus it
// is not tombstoned.

// collectMetricsInBackground periodically collects PebbleDB internal metrics
func (db *Database) collectMetricsInBackground(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Collect metrics every 10 seconds

// collectAndRecordMetrics collects PebbleDB internal metrics and records them
func (db *Database) collectAndRecordMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

// Compaction metrics - report raw counts

// Flush metrics - report raw counts

// Storage metrics per level with level as attribute

//nolint:gosec
//nolint:gosec

// Memtable metrics

//nolint:gosec

// WAL metrics
//nolint:gosec

// Cache metrics - report raw counts
