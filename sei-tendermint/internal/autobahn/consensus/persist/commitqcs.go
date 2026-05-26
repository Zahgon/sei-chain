// TODO: add Prometheus metrics for commitQCs written and truncated.
package persist

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

const commitqcsDir = "commitqcs"

// LoadedCommitQC is a CommitQC loaded from disk during state restoration.
type LoadedCommitQC struct {
	Index types.RoadIndex
	QC    *types.CommitQC
}

// commitQCState is the mutable state protected by CommitQCPersister's mutex.
type commitQCState struct {
	iw   utils.Option[*indexedWAL[*types.CommitQC]]
	next types.RoadIndex
}

// persistCommitQC writes a CommitQC to the WAL. Caller must hold the lock.
// Duplicates (idx < next) are silently ignored for idempotent startup.
// Gaps (idx > next) return an error (breaks linear index mapping).
func (s *commitQCState) persistCommitQC(qc *types.CommitQC) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteBefore truncates WAL entries below the anchor's index, then
// re-persists the anchor for crash recovery. Caller must hold the lock.
func (s *commitQCState) deleteBefore(anchor *types.CommitQC) error {
	_ = "STUB: not implemented"
	return nil
}

// CommitQCPersister manages CommitQC persistence using a WAL.
// Entries are appended in order; each entry is self-describing (the serialized
// CommitQC contains its RoadIndex). The WAL index is append order, not
// RoadIndex — the indexedWAL tracks first/next indices to enable truncation.
// When iw is None, all disk I/O is skipped but cursor tracking still works.
type CommitQCPersister struct {
	state utils.Mutex[*commitQCState]
}

// NewCommitQCPersister opens (or creates) a WAL in the commitqcs/ subdirectory
// and replays all persisted entries. Returns the persister and a sorted slice of
// loaded CommitQCs. Corrupt tail entries are auto-truncated by the WAL library.
// When stateDir is None, returns a no-op persister.
//
// After crash recovery with an empty WAL (e.g. TruncateAll completed but no
// new write followed), LoadNext() returns 0. The caller MUST use
// MaybePruneAndPersist with the prune CommitQC in Anchor to re-establish the
// cursor and re-persist the anchor's CommitQC before appending more QCs.
func NewCommitQCPersister(stateDir utils.Option[string]) (*CommitQCPersister, []LoadedCommitQC, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// LoadNext returns the road index of the first CommitQC that has not been
// persisted (exclusive upper bound of what's on disk).
func (cp *CommitQCPersister) LoadNext() types.RoadIndex {
	_ = "STUB: not implemented"
	return *new(types.RoadIndex)
}

// MaybePruneAndPersist optionally truncates the WAL and/or appends new
// CommitQCs, depending on which arguments are present:
//
//   - anchor set, commitQCs non-empty: truncate WAL below anchor, re-persist
//     the anchor QC for crash recovery, then append new QCs (runtime path).
//   - anchor set, commitQCs empty:     truncate and re-persist anchor only
//     (startup prune path).
//   - anchor empty, commitQCs non-empty: append only, no truncation.
//   - anchor empty, commitQCs empty:     no-op.
//
// The lock is held for the entire truncate-then-append sequence, so callers
// need not coordinate ordering.
// afterEach, when present, is called after each successful append. It is
// invoked while the lock is held, so it must not re-enter the persister.
func (cp *CommitQCPersister) MaybePruneAndPersist(
	anchor utils.Option[*types.CommitQC],
	commitQCs []*types.CommitQC,
	afterEach utils.Option[func(*types.CommitQC)],
) error {
	_ = "STUB: not implemented"
	return nil
}

// Close shuts down the WAL. Safe to call multiple times (idempotent).
func (cp *CommitQCPersister) Close() error { _ = "STUB: not implemented"; return nil }

// no-op persister or already closed

func loadAllCommitQCs(s *commitQCState) ([]LoadedCommitQC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no-op persister (persistence disabled)
