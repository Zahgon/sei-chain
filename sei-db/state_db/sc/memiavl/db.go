package memiavl

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/alitto/pond"

	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/wal"
)

const LockFileName = "LOCK"

var errReadOnly = errors.New("db is read-only")

// DB implements DB-like functionalities on top of MultiTree:
// - async snapshot rewriting
// - Write-ahead-log
//
// The memiavl.db directory looks like this:
// ```
// > current -> snapshot-N
// > snapshot-N
// >  bank
// >    kvs
// >    nodes
// >    metadata
// >  acc
// >  ... other stores
// > rlog
// ```
type DB struct {
	*MultiTree
	dir      string
	fileLock FileLock
	readOnly bool
	opts     Options

	// streamHandler is the changelog WAL owned by MemIAVL.
	// It is opened during OpenDB (if present / allowed) and closed in DB.Close().
	streamHandler wal.ChangelogWAL
	// pendingLogEntry accumulates changes (changesets + upgrades) to be written
	// into the changelog WAL on the next Commit().
	pendingLogEntry proto.ChangelogEntry

	// result channel of snapshot rewrite goroutine
	snapshotRewriteChan chan snapshotResult
	// context cancel function to cancel the snapshot rewrite goroutine
	snapshotRewriteCancelFunc context.CancelFunc
	// the number of old snapshots to keep (excluding the latest one)
	snapshotKeepRecent uint32
	// block interval to take a new snapshot
	snapshotInterval uint32
	// minimum time interval between snapshots
	// Protected by db.mtx (only accessed in Commit call chain)
	snapshotMinTimeInterval time.Duration
	// timestamp of the last successful snapshot creation
	// Protected by db.mtx (only accessed in Commit call chain)
	lastSnapshotTime time.Time

	// pruneSnapshotLock guards concurrent prune operations; use TryLock in pruneSnapshots
	pruneSnapshotLock sync.Mutex
	// closed guards against double Close(), protected by db.mtx
	closed bool

	// walIndexDelta is the difference: version - walIndex for any entry.
	// Since both WAL indices and versions are strictly contiguous, this delta is constant.
	// Computed once when opening the DB from the first WAL entry.
	walIndexDelta int64

	// The assumptions to concurrency:
	// - The methods on DB are protected by a mutex
	// - Each call of OpenDB loads a separate instance, in query scenarios,
	//   it should be immutable, the cache stores will handle the temporary writes.
	// - The DB for the state machine will handle writes through the Commit call,
	//   this method is the sole entry point for tree modifications, and there's no concurrency internally
	//   (the background snapshot rewrite is handled separately), so we don't need locks in the Tree.
	mtx sync.Mutex
	// worker goroutine IdleTimeout = 5s
	snapshotWriterPool *pond.WorkerPool
}

const (
	SnapshotPrefix = "snapshot-"
	SnapshotDirLen = len(SnapshotPrefix) + 20
)

// getSnapshotModTime returns the modification time of the current snapshot directory.
// It reads the "current" symlink to get the actual snapshot directory.
// If the directory doesn't exist or there's an error, returns current time as fallback.
// This is safe for first startup (no snapshots yet) and error recovery scenarios.
//
// NOTE: Relies on filesystem ModTime which may be inaccurate after backup/restore operations.
// TODO: Consider storing timestamp in MultiTreeMetadata for reliability.
func getSnapshotModTime(dir string) time.Time {
	_ = "STUB: not implemented"
	// Read the "current" symlink to get the actual snapshot directory
	return *new(time.Time)
}

// First startup: no snapshot exists yet, use current time
// This is expected and normal for new nodes

// Unexpected error reading symlink

// Clean the path and validate it's within the expected parent directory

// Snapshot directory was deleted but symlink exists (inconsistent state)

// Other stat errors (permission, I/O error, etc.)

func OpenDB(targetVersion int64, opts Options) (database *DB, _err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cleanup any temporary directories left by interrupted snapshot rewrite

// find the biggest snapshot version that's less than or equal to the target version

// Snapshot mmap files are loaded with MADV_RANDOM in OpenSnapshot().

// MemIAVL owns changelog lifecycle: always open the WAL here.
// Even in read-only mode we may need WAL replay to reconstruct non-snapshot versions.

// Compute WAL index delta (only needed once per DB open)

// If WAL is empty, set delta so first WAL entry aligns with NextVersion().

// Replay WAL to catch up to target version (if WAL has entries)

// downgrade `"current"` link first

// truncate the rlog file (if WAL is provided and has entries)

// Use O(1) conversion: walIndex = version - delta

// prune snapshots that's larger than the target version

// create worker pool. recv tasks to write snapshot

// Initialize lastSnapshotTime from the current snapshot directory's modification time
// This ensures accurate time tracking even after restarts
// Read the "current" symlink to get the actual snapshot directory's ModTime

// Apply initial stores on a fresh DB (version 0) so they get persisted to WAL.
// This creates the trees and populates pendingLogEntry, which will be written
// to WAL on the first Commit().
// ApplyUpgrades is idempotent (skips existing trees), so this is safe.

// GetWAL returns the WAL handler for changelog operations.
func (db *DB) GetWAL() wal.ChangelogWAL {
	_ = "STUB: not implemented"
	return *

	// GetWALIndexDelta returns the precomputed delta between version and WAL index.
	// This allows O(1) conversion: version = walIndex + delta, walIndex = version - delta
	new(wal.ChangelogWAL)
}

func (db *DB) GetWALIndexDelta() int64 { _ = "STUB: not implemented"; return 0 }

func removeTmpDirs(rootDir string) error { _ = "STUB: not implemented"; return nil }

// ReadOnly returns whether the DB is opened in read-only mode.
func (db *DB) ReadOnly() bool {
	_ = "STUB: not implemented"

	// SetInitialVersion wraps `MultiTree.SetInitialVersion`.
	// it will do a snapshot rewrite, because we can't use rlog to record this change,
	// we need it to convert versions to rlog index in the first place.
	return false
}

func (db *DB) SetInitialVersion(initialVersion int64) error { _ = "STUB: not implemented"; return nil }

// ApplyUpgrades wraps MultiTree.ApplyUpgrades to add a lock.
func (db *DB) ApplyUpgrades(upgrades []*proto.TreeNameUpgrade) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyChangeSets wraps MultiTree.ApplyChangeSets to add a lock.
func (db *DB) ApplyChangeSets(changeSets []*proto.NamedChangeSet) (_err error) {
	_ = "STUB: not implemented"
	return nil
}

// ApplyChangeSet wraps MultiTree.ApplyChangeSet to add a lock.
// It merges the changeset into any existing pending entry for the same store,
// rather than blindly appending, to ensure at most one WAL entry per store per
// block. Without this, WAL replay (Catchup) would treat duplicate entries as
// separate versions, causing a state divergence.
func (db *DB) ApplyChangeSet(name string, changeSet proto.ChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

// mergePendingChangesets merges new changesets into pendingLogEntry.Changesets.
// If a store already has a pending changeset, the new pairs are appended to it
// rather than creating a duplicate entry. This ensures the WAL entry has at most
// one changeset per store, which is required for correct replay (Catchup calls
// SaveVersion once per changeset, so duplicates would incorrectly bump tree versions).
func (db *DB) mergePendingChangesets(changeSets []*proto.NamedChangeSet) {
	_ = "STUB: not implemented"
	return
}

// checkAsyncTasks checks the status of background tasks non-blocking-ly and process the result
func (db *DB) checkAsyncTasks() error { _ = "STUB: not implemented"; return nil }

// CommittedVersion returns the current version of the MultiTree.
func (db *DB) CommittedVersion() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// checkBackgroundSnapshotRewrite check the result of background snapshot rewrite, cleans up the old snapshots and switches to a new multitree
func (db *DB) checkBackgroundSnapshotRewrite() error {
	_ = "STUB: not implemented"
	// check the completeness of background snapshot rewriting
	return nil
}

// channel was closed without sending a result
// Still prune old snapshots to prevent accumulation

// background snapshot rewrite failed

// Still prune old snapshots to prevent accumulation

// wait for potential pending writes to finish, to make sure we catch up to latest state.
// in real world, block execution should be slower than tree updates, so this should not block for long.

// catchup the remaining entries in rlog

// do the switch

// reset memnode counter

// pruneSnapshots prunes old snapshots, keeping only snapshotKeepRecent recent ones.
// Note: WAL truncation is now handled by CommitStore after each commit.
func (db *DB) pruneSnapshots() { _ = "STUB: not implemented"; return }

// ignore any newer snapshot directories, there could be ongoning snapshot rewrite.

// computeWALIndexDelta computes the constant delta between version and WAL index.
// Since both are strictly contiguous, we only need to read one entry.
// Returns (delta, hasEntries, error). hasEntries is false if WAL is empty.
func computeWALIndexDelta(stream wal.ChangelogWAL) (int64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// empty WAL

// Read just the first entry to compute delta

// delta = version - index, so for any entry: version = index + delta
// #nosec G115 -- WAL indices are always much smaller than MaxInt64 in practice

// versionToWALIndex converts a version to its corresponding WAL index using the precomputed delta.
// Returns 0 if the version would result in an invalid (negative or zero) index.
func (db *DB) versionToWALIndex(version int64) uint64 { _ = "STUB: not implemented"; return 0 }

// #nosec G115 -- index is guaranteed positive by the check above

// walIndexToVersion converts a WAL index to its corresponding version using the precomputed delta.
func (db *DB) walIndexToVersion(index uint64) int64 {
	_ = "STUB: not implemented"
	// #nosec G115 -- WAL indices are always much smaller than MaxInt64 in practice
	return 0
}

// Commit wraps SaveVersion to bump the version and finalize the tree state.
// MemIAVL owns the changelog: it writes the pending changelog entry before committing the tree.
func (db *DB) Commit() (version int64, _err error) { _ = "STUB: not implemented"; return 0, nil }

// Commit the in-memory tree state FIRST.
// MemIAVL is purely in-memory; SaveVersion() doesn't persist anything.
// The changelog WAL is our persistence layer.

// Write to WAL AFTER successful SaveVersion.
// Rationale: If SaveVersion fails but we already wrote to WAL, we'd have
// a WAL entry for a version that was never committed. On replay, this would
// corrupt state. By writing WAL after SaveVersion succeeds, we ensure WAL
// only contains valid committed versions. If WAL write fails after SaveVersion,
// we lose this version on crash (rollback to prior state), but remain consistent.
//
// Note: Write() automatically checks for any previous async write errors.

// Rewrite tree snapshot if applicable

// tryTruncateWAL best-effort truncates old WAL entries that are older than the earliest snapshot.
func (db *DB) tryTruncateWAL() { _ = "STUB: not implemented"; return }

func (db *DB) Copy() *DB { _ = "STUB: not implemented"; return nil }

func (db *DB) copy() *DB { _ = "STUB: not implemented"; return nil }

func (db *DB) ReleaseSnapshotRefs() error { _ = "STUB: not implemented"; return nil }

// RewriteSnapshot writes the current version of memiavl into a snapshot, and update the `current` symlink.
func (db *DB) RewriteSnapshot(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Check if snapshot already exists

// targetPath exists but is not a directory - this is unexpected

// Rename temporary directory to final location

// Clean up temporary directory on rename failure

func (db *DB) Reload() error { _ = "STUB: not implemented"; return nil }

func (db *DB) reload() error { _ = "STUB: not implemented"; return nil }

func (db *DB) reloadMultiTree(mtree *MultiTree) error {
	_ = "STUB: not implemented"
	// The caller is responsible for ensuring mtree is caught up to the latest state
	// (either via Catchup from WAL or by loading a current snapshot).
	return nil
}

// rewriteIfApplicable execute the snapshot rewrite strategy according to current height
func (db *DB) rewriteIfApplicable(height int64) { _ = "STUB: not implemented"; return }

// Create snapshot when all conditions are met:
// 1. Block height interval is reached (height - last snapshot >= interval)
// 2. Minimum time interval has elapsed (prevents excessive snapshots during catch-up)
// 3. Block height % snapshot interval == 0

type snapshotResult struct {
	mtree *MultiTree
	err   error
}

// RewriteSnapshotBackground rewrite snapshot in a background goroutine,
// `Commit` will check the complete status, and switch to the new snapshot.
func (db *DB) RewriteSnapshotBackground() error { _ = "STUB: not implemented"; return nil }

func (db *DB) rewriteSnapshotBackground() error { _ = "STUB: not implemented"; return nil }

// Use buffered channel to avoid blocking the goroutine when sending result

// Update snapshot timestamp at start (not at completion) for accurate interval calculation

// Release per-tree snapshot refs; don't call cloned.Close() which
// would also tear down the live db's writer pool and stream handler.

// Disable prefetch when loading newly created snapshot in background.
// Profiling shows: with prefetch = 35 min, without = 15 min (20 min difference!)
// The snapshot was just written, so some data is still in page cache, but mincore()
// checks mmap pages (not file cache) and reports low residency because mmap hasn't
// been accessed yet. Prefetch causes unnecessary I/O that competes with ongoing
// commits and evicts hot pages from the active snapshot still being used by main chain.
// Use cloned.opts instead of db.opts to avoid race condition with Close()

// Snapshot mmap files are loaded with MADV_RANDOM in OpenSnapshot().

// do a best effort catch-up, will do another final catch-up in main thread.

func (db *DB) Close() error { _ = "STUB: not implemented"; return nil }

// Wait for any ongoing prune to finish, then block new prunes

// Close rewrite channel first - must wait for background goroutine before closing WAL

// Wait for goroutine to finish and send result

// Close the returned mtree to avoid resource leak

// Close WAL after snapshot rewrite goroutine has fully exited.

// Stop the snapshot writer pool

// Close file lock

// TreeByName wraps MultiTree.TreeByName to add a lock.
func (db *DB) TreeByName(name string) *Tree { _ = "STUB: not implemented"; return nil }

// Version wraps MultiTree.Version to add a lock.
func (db *DB) Version() int64 { _ = "STUB: not implemented"; return 0 }

// LastCommitInfo returns the last commit info.
func (db *DB) LastCommitInfo() *proto.CommitInfo { _ = "STUB: not implemented"; return nil }

func (db *DB) SaveVersion(updateCommitInfo bool) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (db *DB) WorkingCommitInfo() *proto.CommitInfo { _ = "STUB: not implemented"; return nil }

// UpdateCommitInfo wraps MultiTree.UpdateCommitInfo to add a lock.
func (db *DB) UpdateCommitInfo() { _ = "STUB: not implemented"; return }

// WriteSnapshot wraps MultiTree.WriteSnapshot to add a lock.
func (db *DB) WriteSnapshot(dir string) error { _ = "STUB: not implemented"; return nil }

func snapshotName(version int64) string { _ = "STUB: not implemented"; return "" }

func currentPath(root string) string { _ = "STUB: not implemented"; return "" }

func currentTmpPath(root string) string { _ = "STUB: not implemented"; return "" }

func currentVersion(root string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseVersion(name string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// seekSnapshot find the biggest snapshot version that's smaller than or equal to the target version,
// returns 0 if not found.
func seekSnapshot(root string, targetVersion int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetEarliestVersion returns the earliest snapshot name in the db
func GetEarliestVersion(root string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// init a empty memiavl db
//
// ```
// snapshot-0
//
//	commit_info
//
// current -> snapshot-0
// ```
func initEmptyDB(dir string, initialVersion uint32) error { _ = "STUB: not implemented"; return nil }

// create tmp worker pool

// updateCurrentSymlink creates or replace the current symbolic link atomically.
// it could fail under concurrent usage for tmp file conflicts.
func updateCurrentSymlink(dir, snapshot string) error { _ = "STUB: not implemented"; return nil }

// assuming file renaming operation is atomic

// traverseSnapshots traverse the snapshot list in specified order.
func traverseSnapshots(dir string, ascending bool, callback func(int64) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// atomicRemoveDir is equavalent to `mv snapshot snapshot-tmp && rm -r snapshot-tmp`
func atomicRemoveDir(path string) error { _ = "STUB: not implemented"; return nil }

// createDBIfNotExist detects if db does not exist and try to initialize an empty one.
func createDBIfNotExist(dir string, initialVersion uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func isSnapshotName(name string) bool { _ = "STUB: not implemented"; return false }

// GetLatestVersion finds the latest version number without loading the whole db,
// it's needed for upgrade module to check store upgrades,
// it returns 0 if db doesn't exist or is empty.
func GetLatestVersion(dir string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
