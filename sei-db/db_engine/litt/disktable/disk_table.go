package disktable

import (
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/disktable/keymap"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/disktable/segment"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/types"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

var _ litt.ManagedTable = (*DiskTable)(nil)

// keymapReloadBatchSize is the size of the batch used for reloading keys from segments into the keymap.
const keymapReloadBatchSize = 1024

const tableFlushChannelCapacity = 8

// DiskTable manages a table's Segments.
type DiskTable struct {
	// The logger for the disk table.
	logger *slog.Logger

	// errorMonitor is a struct that permits the DB to "panic". There are many goroutines that function under the
	// hood, and many of these threads could, in theory, encounter errors which are unrecoverable. In such situations,
	// the desirable outcome is for the DB to report the error and then refuse to do additional work. If the DB is in a
	// broken state, it is much better to refuse to do work than to continue to do work and potentially corrupt data.
	errorMonitor *util.ErrorMonitor

	// The root directories for the disk table. Each of these directories' name matches the name of the table.
	roots []string

	// Configures the location where segment data is stored.
	segmentPaths []*segment.SegmentPath

	// The table's name.
	name string

	// The table's metadata.
	metadata *tableMetadata

	// A map of keys to their addresses.
	keymap keymap.Keymap

	// The path to the keymap directory.
	keymapPath string

	// The type file for the keymap.
	keymapTypeFile *keymap.KeymapTypeFile

	// unflushedDataCache is a map of keys to their values that may not have been flushed to disk yet. This is used as a
	// lookup table when data is requested from the table before it has been flushed to disk.
	unflushedDataCache sync.Map

	// clock is the time source used by the disk table.
	clock func() time.Time

	// The number of bytes contained within all segments, including the mutable segment. This tracks the number of
	// bytes that are on disk, not bytes in memory.
	size atomic.Uint64

	// The number of keys in the table.
	keyCount atomic.Int64

	// The control loop is a goroutine responsible for scheduling operations that mutate the table.
	controlLoop *controlLoop

	// The flush loop is a goroutine responsible for blocking on flush operations.
	flushLoop *flushLoop

	// Encapsulates metrics for the database.
	metrics *metrics.LittDBMetrics

	// Set to true when the table is closed. This is used to prevent double closing.
	closed atomic.Bool

	// Set to true when the table is destroyed. This is used to prevent double destroying.
	destroyed atomic.Bool

	// If true then ensure file operations are synced to disk.
	fsync bool

	// Manages flush requests and flush request batching. This is a performance optimization.
	flushCoordinator *flushCoordinator
}

// NewDiskTable creates a new DiskTable.
func NewDiskTable(
	config *litt.Config,
	name string,
	keymap keymap.Keymap,
	keymapPath string,
	keymapTypeFile *keymap.KeymapTypeFile,
	roots []string,
	reloadKeymap bool,
	metrics *metrics.LittDBMetrics) (litt.ManagedTable, error) {
	_ = "STUB: not implemented"
	return *new(litt.ManagedTable), nil
}

// For each root directory, create a segment directory if it doesn't exist.

// Delete any orphaned swap files:

// Find the table metadata file or create a new one.

// We've found an existing metadata file. Use it.

// No metadata file exists yet. Create a new one in the first root.

// Metadata file exists, so we need to load it.

// Load segments.

// Create the mutable segment

// Initialize snapshot files if snapshotting is enabled.

// Start the flush loop.

// Start the control loop.

func (d *DiskTable) KeyCount() uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // key count non-negative

func (d *DiskTable) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// repairSnapshot is responsible for making any required repairs to the snapshot directories. This is needed
// if there is a crash, resulting in a segment not being fully snapshotted. It is also needed if LittDB has
// been rebased (which breaks symlinks) or manually modified (e.g. by the LittDB cli). Returns the new upper bound
// file for the repaired snapshot.
func (d *DiskTable) repairSnapshot(
	symlinkDirectory string,
	lowestSegmentIndex uint32,
	highestSegmentIndex uint32,
	segments map[uint32]*segment.Segment) (*BoundaryFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prevent other processes from messing with the symlink table directory while we are working on it.

// Delete all data from the previous snapshot. This directory will contain a bunch of symlinks. It's a lot
// simpler to just rebuild this from scratch than it is to try to figure out which symlinks are valid
// and which are not. Building this is super fast, so this is not a performance concern.

// There is only the mutable segment, nothing else to do.

// The lower bound file contains the index of the highest segment that has been GC'd by an external process.
// We should ignore the segment at this index, and all segments with lower indices.

// Skip iterating over the highest segment index (i.e. don't do i <= highestSegmentIndex). The highest segment
// index is mutable and cannot be snapshotted until it has been sealed.

// Signal that the segment files are now fully snapshotted and safe to use.
// The highest segment index is the mutable segment, which is not snapshotted.

// reloadKeymap reloads the keymap from the segments. This is necessary when the keymap is lost, the keymap doesn't
// save its data on disk, or we are migrating from one keymap type to another.
func (d *DiskTable) reloadKeymap(
	segments map[uint32]*segment.Segment,
	lowestSegmentIndex uint32,
	highestSegmentIndex uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// ignore unsealed segment, this will have been created in the current session and will not
// yet contain any data.

// Now that the keymap is loaded, write the marker file that indicates that the keymap is fully loaded.
// If we crash prior to writing this file, the keymap will reload from the segments again.

//nolint:gosec // path within keymap directory

func (d *DiskTable) Name() string {
	_ = "STUB: not implemented"

	// Close stops the disk table. Flushes all data out to disk.
	return ""
}

func (d *DiskTable) Close() error { _ = "STUB: not implemented"; return nil }

// Destroy stops the disk table and delete all files.
func (d *DiskTable) Destroy() error { _ = "STUB: not implemented"; return nil }

// already destroyed

// release all segments

// wait for segments to delete themselves

// delete all segment directories (ignore snapshots -- this is the responsibility of an outside process to clean)

// delete the snapshot hardlink directory

// destroy the keymap

// delete the metadata file

// delete the root directories for the table

// SetTTL sets the TTL for the disk table. If set to 0, no TTL is enforced. This setting affects both new
// data and data already written.
func (d *DiskTable) SetTTL(ttl time.Duration) error { _ = "STUB: not implemented"; return nil }

func (d *DiskTable) SetShardingFactor(shardingFactor uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiskTable) Get(key []byte) (value []byte, exists bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// First, check if the key is in the unflushed data map.
// If so, return it from there.

// Look up the address of the data.

// Reserve the segment that contains the data.

// Read the data from disk.

func (d *DiskTable) CacheAwareGet(
	key []byte,
	onlyReadFromCache bool,
) (value []byte, exists bool, hot bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, false, nil
}

// First, check if the key is in the unflushed data map. If so, return it from there.
// Performance wise, this has equivalent semantics to reading the value from
// a cache, so we'd might as well count it as a cache hit.

// Look up the address of the data.

// The value exists but we are not allowed to read it from disk.

// Reserve the segment that contains the data.

// This can happen if there is a race between this thread and the GC thread, i.e.
// if we start reading a value just as the garbage collector decides to delete it.

// Read the data from disk.

func (d *DiskTable) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

func (d *DiskTable) PutBatch(batch []*types.KVPair) error { _ = "STUB: not implemented"; return nil }

func (d *DiskTable) Exists(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Flush flushes all data to disk. Blocks until all data previously submitted to Put has been written to disk.
func (d *DiskTable) Flush() error {
	_ = "STUB: not implemented"
	// The flush coordinator batches flush requests together to improve performance if
	// flushes are being requested very frequently.
	return nil
}

// actually flushes the internal DB
func (d *DiskTable) flushInternal() error { _ = "STUB: not implemented"; return nil }

func (d *DiskTable) SetWriteCacheSize(size uint64) error { _ = "STUB: not implemented"; return nil }

// this implementation does not provide a cache, if a cache is needed then it must be provided by a wrapper

func (d *DiskTable) SetReadCacheSize(size uint64) error { _ = "STUB: not implemented"; return nil }

// this implementation does not provide a cache, if a cache is needed then it must be provided by a wrapper

func (d *DiskTable) RunGC() error { _ = "STUB: not implemented"; return nil }

// writeKeysToKeymap flushes all keys to the keymap. Once they are flushed, it also removes the keys from the
// unflushedDataCache.
func (d *DiskTable) writeKeysToKeymap(keys []*types.ScopedKey) error {
	_ = "STUB: not implemented"
	return nil

	// Nothing to flush.
}

// Keys are now durably written to both the segment and the keymap. It is therefore safe to remove them from the
// unflushed data cache.
