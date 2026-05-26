package flatkv

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

// isClosed reports whether the store's DB handles have been released.
func (s *CommitStore) isClosed() bool { _ = "STUB: not implemented"; return false }

// closeDBsOnly closes all database handles and the WAL but retains the
// file lock, preventing a race window during Rollback or LoadVersion.
func (s *CommitStore) closeDBsOnly() error { _ = "STUB: not implemented"; return nil }

// Close drains thread pools, closes all database instances, cancels the
// store's context to stop background goroutines (caches, metrics), and
// releases the file lock.
func (s *CommitStore) Close() error { _ = "STUB: not implemented"; return nil }

// CleanupOrphanedReadOnlyDirs acquires the writer lock and removes readonly-*
// working directories left behind by a previous process crash. It is a
// startup-only API and must be called before any read-only instances are
// created in the current process. The acquired writer lock is retained for
// subsequent LoadVersion(..., false) calls.
func (s *CommitStore) CleanupOrphanedReadOnlyDirs() error { _ = "STUB: not implemented"; return nil }

// Exporter creates an exporter for the given version by opening a read-only
// clone and performing a full scan of all DBs. The returned exporter must be
// closed when done (which also closes the read-only clone).
func (s *CommitStore) Exporter(version int64) (types.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(types.Exporter), nil
}
