package flatkv

// On-disk layout under <home>/flatkv/:
//
//	flatkv/
//	  current -> snapshot-NNNNN              (symlink to active snapshot)
//	  snapshot-NNNNN/                        (immutable checkpoint)
//	    account/                             (PebbleDB: addr → AccountValue)
//	    code/                                (PebbleDB: addr → bytecode)
//	    storage/                             (PebbleDB: addr||slot → value)
//	    legacy/                              (PebbleDB: full key → value)
//	    metadata/                            (PebbleDB: version + LtHash)
//	  working/                               (mutable clone of active snapshot)
//	    account/, code/, storage/, legacy/, metadata/
//	    SNAPSHOT_BASE                        (records source snapshot name)
//	  changelog/                             (WAL, shared across snapshots)
const (
	// snapshotPrefix is the directory name prefix for versioned snapshots.
	snapshotPrefix = "snapshot-"
	// snapshotDirLen is the full directory name length: "snapshot-" + 20-digit zero-padded version.
	snapshotDirLen = len(snapshotPrefix) + 20

	// currentLink is the symlink name pointing to the active snapshot directory.
	currentLink = "current"
	// currentTmpLink is a temporary symlink used during atomic swap of currentLink.
	currentTmpLink = "current-tmp"

	// workingDirName is cloned from the baseline snapshot on each open().
	// Mutable DB operations go here, keeping snapshot dirs immutable.
	workingDirName = "working"

	// snapshotBaseFile records which snapshot the working dir was cloned from.
	// When the current symlink still points at the same snapshot, we skip
	// the expensive RemoveAll+re-clone on restart because WAL catchup is
	// idempotent and will bring the working dir up to date.
	snapshotBaseFile = "SNAPSHOT_BASE"
)

func snapshotName(version int64) string { _ = "STUB: not implemented"; return "" }

func isSnapshotName(name string) bool { _ = "STUB: not implemented"; return false }

func parseSnapshotVersion(name string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func currentPath(root string) string { _ = "STUB: not implemented"; return "" }

// currentSnapshotDir reads the current symlink and returns the full path
// and parsed version. Returns os.ErrNotExist if the symlink does not exist.
func currentSnapshotDir(root string) (dir string, version int64, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// seekSnapshot finds the highest snapshot version <= targetVersion.
// Returns 0 and an error if no qualifying snapshot exists.
func seekSnapshot(root string, targetVersion int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// traverseSnapshots iterates snapshot directories in the given order.
// ascending=true  -> lowest version first
// ascending=false -> highest version first
// The callback returns (stop, err). Traversal halts on stop=true or err!=nil.
func traverseSnapshots(dir string, ascending bool, fn func(int64) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// updateCurrentSymlink atomically updates the current symlink to point at snapshotDir.
// snapshotDir should be the bare directory name (e.g. "snapshot-00000000000000000100"),
// not a full path.
func updateCurrentSymlink(root, snapshotDir string) error { _ = "STUB: not implemented"; return nil }

// snapshotDBDirs lists the DB subdirectory names included in a snapshot.
var snapshotDBDirs = []string{accountDBDir, codeDBDir, storageDBDir, legacyDBDir, metadataDir}

// removeTmpDirs removes any directories ending in "-tmp" or "-removing"
// left over from interrupted snapshot writes or deletes.
func removeTmpDirs(dir string) error { _ = "STUB: not implemented"; return nil }

// createWorkingDir ensures a mutable working directory exists, cloned from
// snapDir. If the working dir already exists and was cloned from the same
// snapshot (recorded in SNAPSHOT_BASE), the expensive re-clone is skipped
// because WAL catchup is idempotent and will bring data up to date.
func createWorkingDir(snapDir, workDir string) error { _ = "STUB: not implemented"; return nil }

// reuseWorkingDir returns true if workDir exists and was cloned from the
// same snapshot, meaning a full re-clone can be skipped.
func reuseWorkingDir(workDir, snapBase string) bool { _ = "STUB: not implemented"; return false }

//nolint:gosec // path built from internal working dir layout

func writeSnapshotBase(workDir, snapBase string) error { _ = "STUB: not implemented"; return nil }

// cloneDir copies a single PebbleDB directory. Immutable .sst files are
// hard-linked; everything else is byte-copied. LOCK files are skipped.
func cloneDir(src, dst string) error { _ = "STUB: not implemented"; return nil }

// Fall back to copy if hardlink fails (e.g. cross-device).

func copyFile(src, dst string) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // path built from internal snapshot layout

//nolint:gosec // path built from internal snapshot layout

// atomicRemoveDir renames the directory to a trash name then removes it,
// preventing half-deleted snapshots on crash.
func atomicRemoveDir(path string) error { _ = "STUB: not implemented"; return nil }

// resolveSnapshotDir returns the full path to the active snapshot directory.
// It handles four cases: (1) current symlink exists, (2) migration from
// pre-snapshot flat layout, (3) recovery from a partial migration crash,
// or (4) initialization of a fresh empty snapshot.
func (s *CommitStore) resolveSnapshotDir(flatkvDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// No flat dirs. Check for an orphaned snapshot directory — this happens
// when a previous migration moved all dirs but crashed before creating
// the current symlink.

// migrateFlatLayout moves the existing flat DB directories
// (account/, code/, storage/, metadata/) into a snapshot directory and
// creates the current symlink.
//
// The function is idempotent: directories that were already moved by a
// previous partial attempt are skipped, so recovery from a mid-migration
// crash completes the remaining moves.
func (s *CommitStore) migrateFlatLayout(flatkvDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Determine version for the snapshot name. The metadata DB might still
// be at the flat location or might have been moved in a prior attempt.

//nolint:gosec // block height, always < MaxInt64

// Metadata already moved — look for the snapshot dir from a prior attempt.

// WriteSnapshot creates a PebbleDB checkpoint of the committed state.
// The snapshot is written into a versioned subdirectory under the flatkv root
// (e.g. flatkv/snapshot-00000000000000000100) and the current symlink is updated.
// The dir parameter is ignored; snapshots are always stored alongside the live data.
func (s *CommitStore) WriteSnapshot(_ string) (err error) { _ = "STUB: not implemented"; return nil }

// Deterministic order (slice, not map) for reproducibility.

// idempotent: stale final may exist

// Keep SNAPSHOT_BASE in sync so the next restart reuses the working dir
// instead of re-cloning from the snapshot and replaying the full WAL gap.

// pruneSnapshots removes old snapshots beyond SnapshotKeepRecent, keeping
// the latest snapshot (currentVersion) plus the N most recent older ones.
// Best-effort: errors are logged but do not fail the snapshot operation.
func (s *CommitStore) pruneSnapshots(dir string, currentVersion int64) int {
	_ = "STUB: not implemented"
	return 0
}

// Rollback restores state to targetVersion by rewinding to the highest
// snapshot <= targetVersion, replaying WAL to reach the target, and
// truncating all WAL entries and snapshots beyond that point.
//
// Crash safety: the WAL is truncated BEFORE catchup writes any data to
// PebbleDB. If the process crashes after truncation but before catchup
// completes, the next restart will simply re-run catchup against the
// already-truncated WAL, converging to targetVersion.
func (s *CommitStore) Rollback(targetVersion int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Force a fresh working dir clone from the rollback snapshot: the
// current working dir may contain data beyond targetVersion.

// Truncate WAL beyond targetVersion BEFORE catchup (crash safety).

// Target predates all WAL entries; clear the entire WAL to
// prevent re-application. tidwall/wal cannot truncate to empty,
// so we close, delete, and reopen.

// verifyWALTail checks that the last WAL entry has the expected version.
func (s *CommitStore) verifyWALTail(expectedVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

// tryTruncateWAL is a best-effort truncation of WAL entries that are older
// than the earliest snapshot. This prevents unbounded WAL growth while
// keeping enough entries for rollback to any retained snapshot.
func (s *CommitStore) tryTruncateWAL() { _ = "STUB: not implemented"; return }

// Find the earliest (lowest-version) snapshot — we must keep WAL entries
// from that point onward so rollback to it is possible.
