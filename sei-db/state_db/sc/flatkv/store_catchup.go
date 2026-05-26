package flatkv

// walOffsetForVersion returns the WAL offset whose entry has the given version.
// Returns 0 if the WAL is empty or the version predates all WAL entries.
//
// Strategy: try the arithmetic shortcut (O(1) reads) first -- it works when
// each version maps 1:1 to a sequential offset. On mismatch, fall back to
// binary search (O(log N) reads) which handles gaps and batched versions.
func (s *CommitStore) walOffsetForVersion(version int64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Fast path: O(1) arithmetic guess.
//nolint:gosec // version >= firstVer checked above

// Slow path: binary search over [firstOff, lastOff].

// walVersionAtOffset reads a single WAL entry and returns its version.
func (s *CommitStore) walVersionAtOffset(off uint64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// catchup replays WAL entries from the current committedVersion up to (and
// including) targetVersion. If targetVersion <= 0, replay continues to the
// end of the WAL.
//
// Each replayed entry runs through ApplyChangeSets (which updates
// workingLtHash) and commitBatches (which persists to the per-DB PebbleDBs).
// After all entries are replayed, global metadata is flushed once.
//
// catchup enforces WAL continuity from committedVersion+1: if the WAL has been
// truncated past committedVersion (e.g. a tooling clone raced with a live
// snapshot that triggered tryTruncateWAL), it returns an error rather than
// silently skipping the missing versions and corrupting the LtHash.
func (s *CommitStore) catchup(targetVersion int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Replayed blocks are reported regardless of catchup outcome: even on a
// later error, the blocks that did replay are real progress that moved
// committedVersion forward.
//
// CurrentVersion is intentionally NOT recorded here: the write-mode
// callers (LoadVersion, Rollback) record it on success themselves, and
// recording it from catchup would pollute the gauge when openReadOnly
// invokes catchup during a read-only historical load.

// Resolve the WAL boundary versions once so we can both pick the right
// start offset AND detect a truncated-past-snapshot gap.

// Nothing past committedVersion in the WAL: clean no-op.

// Gap detection: WAL must cover committedVersion+1, otherwise we
// would jump straight to walFirstVer and silently skip the
// versions in between. Return loud error rather than corrupt
// committedLtHash.

// Bound end offset to avoid deserializing entries past the target:
// O(target - snapshot) instead of O(WAL_size).

// Defensive continuity check: even with the boundary check
// above, a corrupted WAL could have an internal hole. Refuse to
// keep applying past such a hole.

// During catchup with Sync=false, per-entry batch commits can leave
// data only in OS/page cache. Flush once before advancing global
// metadata so global watermark won't get ahead of data durability.
