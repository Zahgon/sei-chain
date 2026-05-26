package operations

import (
	"errors"

	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv"
	"github.com/sei-protocol/sei-chain/sei-db/wal"
)

const (
	flatkvSnapshotPrefix = "snapshot-"
	flatkvSnapshotDirLen = len(flatkvSnapshotPrefix) + 20

	// maxCloneRetries bounds the number of retries when the source snapshot
	// is pruned mid-clone by a live writer (atomicRemoveDir race) or when the
	// live writer truncates the WAL past our snapshot between snapshot and
	// changelog clone steps.
	maxCloneRetries = 3
)

// errSourceChurning marks transient races where the source FlatKV directory
// mutates (snapshot pruned, WAL truncated) between our reads. It is the
// sentinel that prepareFlatKVToolingCloneWith uses to decide whether to
// retry instead of bailing out.
var errSourceChurning = errors.New("flatkv source kept churning during clone")

// openedFlatKV wraps a temp-cloned FlatKV store used by tooling.
//
// The tools intentionally operate on a temp clone of the selected snapshot +
// WAL so they do not compete with a live node for the FlatKV writer lock.
type openedFlatKV struct {
	*flatkv.CommitStore
	tempDir string
}

func (o *openedFlatKV) Close() error { _ = "STUB: not implemented"; return nil }

// openFlatKVReadOnly opens FlatKV tooling state at the given height.
//
// Instead of opening the source directory directly (which would contend for
// FlatKV's writer lock on a live node), this clones the relevant snapshot and
// changelog into a temp directory and opens that isolated clone.
//
// Consistency on a live node:
//   - snapshot-N/ directories are immutable after creation (Pebble
//     Checkpoint + atomic Rename). Their contents never change; only
//     wholesale pruning via atomicRemoveDir can remove them.
//   - Snapshot files are hard-linked. A hardlink preserves the inode even if
//     the live node prunes the source snapshot mid-operation, so the tool sees
//     a stable snapshot until it releases its temp dir.
//   - Changelog files are byte-copied, not linked, because WAL recovery can
//     truncate a corrupted tail when the cloned store opens.
//   - If the whole snapshot directory is renamed to "-removing" between our
//     os.ReadDir and os.Link calls, we surface ENOENT, re-select the
//     snapshot, and retry up to maxCloneRetries times.
//
// height=0 means latest version.
func openFlatKVReadOnly(dbDir string, height int64) (*openedFlatKV, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareFlatKVToolingClone(dbDir string, height int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func prepareFlatKVToolingCloneWith(dbDir string, height int64, tryClone func(string, int64) (string, error)) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// isCloneRetryableError reports whether err indicates a transient race with
// the live writer that we should retry: either the snapshot or a WAL segment
// vanished mid-read (ENOENT), or our post-clone validation observed the WAL
// being truncated past our snapshot version.
func isCloneRetryableError(err error) bool { _ = "STUB: not implemented"; return false }

func tryPrepareFlatKVToolingClone(dbDir string, height int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Place the temp clone inside dbDir so it is on the exact same mounted
// filesystem as the source snapshots. A sibling directory is not enough:
// dbDir itself is often a mount point on dedicated data volumes.

// Detect the snapshot/WAL race: a live writer can roll a new
// snapshot between our snapshot clone and our changelog copy and
// then truncateWAL up to that newer snapshot's version. If that
// happened, the cloned WAL no longer covers snapshotVersion+1,
// and a downstream catchup would silently jump over missing
// versions. Surface it as a retryable error so the outer loop
// re-selects the snapshot and tries again.

// verifyClonedWALCovers opens the cloned WAL just long enough to ensure it
// either is empty, ends at or before snapshotVersion (no replay needed), or
// starts at or before snapshotVersion+1 (catchup can resume cleanly).
func verifyClonedWALCovers(dstChangelogDir string, snapshotVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

func readWALEntryVersion(walLog wal.ChangelogWAL, off uint64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func selectFlatKVSnapshot(dbDir string, height int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func listFlatKVSnapshots(dir string) []int64 { _ = "STUB: not implemented"; return nil }

func isFlatKVSnapshotName(name string) bool { _ = "STUB: not implemented"; return false }

// cloneDirRecursive clones an immutable snapshot directory into dst by
// hardlinking every regular file. EXDEV is treated as a fatal configuration
// error: snapshots can be many GB, and the previous behavior of falling back
// to a byte-copy on tmpfs (the historical $TMPDIR default) routinely OOM'd
// nodes and exhausted /tmp. Callers must ensure the tool clone dir lives on
// the same filesystem as the source FlatKV directory.
//
// Hardlinking is safe because:
//   - snapshot-N files are immutable after Pebble Checkpoint + Rename.
//
// It also lets the tool survive a concurrent atomicRemoveDir on the source:
// once we have hardlinks, the inodes persist until we release the temp dir,
// even if the live node prunes the source snapshot mid-operation.
func cloneDirRecursive(src, dst string) error { _ = "STUB: not implemented"; return nil }

// copyDirRecursive clones a mutable directory by byte-copying every regular
// file. Changelog files must not share inodes with live WAL segments because
// WAL open/recovery may truncate a corrupted tail in the cloned store.
func copyDirRecursive(src, dst string) error { _ = "STUB: not implemented"; return nil }

func cloneDirRecursiveWith(src, dst string, cloneFile func(string, string) error) error {
	_ = "STUB: not implemented"
	return nil
}

// linkOnly hardlinks src to dst and refuses to silently byte-copy if the
// hardlink fails because of a cross-device boundary. Snapshots can be many
// GB; falling back to a copy is unsafe (slow, RAM-intensive, fills tmpfs)
// and is the bug this guard exists to surface. Any other error (including
// ENOENT from a mid-clone prune) is returned as-is so callers can retry.
func linkOnly(src, dst string) error { _ = "STUB: not implemented"; return nil }

func isCrossDeviceLinkError(err error) bool { _ = "STUB: not implemented"; return false }

func copyFile(src, dst string) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // src is selected from a FlatKV snapshot/changelog clone tree.

//nolint:gosec // dst is allocated inside the tool's temporary clone directory.
