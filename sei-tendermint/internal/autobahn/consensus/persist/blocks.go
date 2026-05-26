// TODO: add Prometheus metrics for blocks written and truncated.
package persist

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "internal", "autobahn", "consensus", "persist")

const blocksDir = "blocks"

// LoadedBlock is a block loaded from disk during state restoration.
type LoadedBlock struct {
	Number   types.BlockNumber
	Proposal *types.Signed[*types.LaneProposal]
}

// laneWALState is the mutable state of a lane WAL, protected by laneWAL's
// mutex. Block numbers within a lane are contiguous, so the first block
// number is derived: nextBlockNum - Count().
type laneWALState struct {
	*indexedWAL[*types.Signed[*types.LaneProposal]]
	nextBlockNum types.BlockNumber
}

func (s *laneWALState) firstBlockNum() utils.Option[types.BlockNumber] {
	_ = "STUB: not implemented"
	return nil
}

// persistBlock writes a proposal to the WAL and advances nextBlockNum.
// Caller must hold the per-lane lock.
func (s *laneWALState) persistBlock(proposal *types.Signed[*types.LaneProposal]) error {
	_ = "STUB: not implemented"
	return nil
}

// truncateForAnchor truncates the WAL so that `first` becomes the oldest
// retained block number. Caller must hold the per-lane lock.
func (s *laneWALState) truncateForAnchor(lane types.LaneID, first types.BlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// WAL is empty; nothing to truncate but advance the cursor so
// the next persistBlock expects the right block number.

// loadAll reads all entries from the lane WAL and returns the loaded blocks.
// Also restores nextBlockNum from the last entry.
func (s *laneWALState) loadAll() ([]LoadedBlock, error) { _ = "STUB: not implemented"; return nil, nil }

// laneWAL wraps a laneWALState with a mutex that serializes all writes and
// truncations on a single lane.
type laneWAL struct {
	state utils.Mutex[*laneWALState]
}

func (lw *laneWAL) maybePruneAndPersist(
	lane types.LaneID,
	anchor utils.Option[*types.CommitQC],
	proposals []*types.Signed[*types.LaneProposal],
	afterEach utils.Option[func(*types.Signed[*types.LaneProposal])],
) error {
	_ = "STUB: not implemented"
	return nil
}

func (lw *laneWAL) close() error { _ = "STUB: not implemented"; return nil }

// BlockPersister manages block persistence using one WAL per lane.
// Each lane gets its own WAL in a subdirectory named by hex-encoded lane ID,
// so truncation is independent per lane. A single shared WAL would be simpler
// but a lane whose blocks are never included in a committed block (e.g. the
// validator is removed from the committee) would prevent truncation of all
// other lanes' entries that follow it.
// When dir is None, all disk I/O is skipped (no-op mode).
//
// All public methods are safe for concurrent use. The lanes map is protected
// by an RWMutex; each laneWAL has its own Mutex for write serialization.
// MaybePruneAndPersistLane holds the per-lane lock for the entire
// truncate-then-append sequence, so concurrent calls on the same lane
// serialize correctly. Different lanes are fully parallel.
//
// NOTE: MaybePruneAndPersistLane releases the map RLock before acquiring
// the per-lane lock. This is safe because lanes are only added, never
// removed. If lane deletion is added in the future, the map RLock must be
// held through the WAL write.
type BlockPersister struct {
	dir   utils.Option[string] // immutable after construction
	lanes utils.RWMutex[map[types.LaneID]*laneWAL]
}

func laneDir(lane types.LaneID) string { _ = "STUB: not implemented"; return "" }

func newLaneWALState(dir string) (*laneWALState, error) { _ = "STUB: not implemented"; return nil, nil }

// NewBlockPersister opens (or creates) per-lane WALs in subdirectories of
// blocks/ and replays all persisted entries. Returns the persister and loaded
// blocks grouped by lane (sorted by block number). Corrupt tail entries are
// auto-truncated by the WAL library.
// When stateDir is None, returns a no-op persister.
func NewBlockPersister(stateDir utils.Option[string]) (*BlockPersister, map[types.LaneID][]LoadedBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// getOrCreateLane returns the laneWAL for the given lane, creating it if
// necessary. Uses double-checked locking: fast path reads under RLock;
// slow path (lane creation) promotes to a write Lock.
// The returned pointer is safe to use after the lock is released because
// lanes are only ever added, never removed (see BlockPersister doc).
// Returns an error if called on a no-op persister (caller should check first).
func (bp *BlockPersister) getOrCreateLane(lane types.LaneID) (*laneWAL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fast path: read-only check.

// Slow path: create under write lock (double-checked).

// MaybePruneAndPersistLane optionally truncates the lane's WAL and/or appends
// new proposals, depending on which arguments are present:
//
//   - anchor set, proposals non-empty: truncate WAL below anchor, then append (runtime path).
//   - anchor set, proposals empty:     truncate only, no appends (startup prune path).
//   - anchor empty, proposals non-empty: append only, no truncation.
//   - anchor empty, proposals empty:     no-op.
//
// afterEach, when present, is called after each successful append. It is
// invoked while the per-lane lock is held, so it must not re-enter the
// persister.
// No-op persister (dir=None): skips disk I/O but still invokes afterEach.
// Does not spawn goroutines — the caller schedules parallelism per lane.
//
// The per-lane lock is held for the entire truncate-then-append sequence,
// so concurrent calls on the same lane serialize correctly.
func (bp *BlockPersister) MaybePruneAndPersistLane(
	lane types.LaneID,
	anchor utils.Option[*types.CommitQC],
	proposals []*types.Signed[*types.LaneProposal],
	afterEach utils.Option[func(*types.Signed[*types.LaneProposal])],
) error {
	_ = "STUB: not implemented"
	return nil
}

// close shuts down all per-lane WALs. Internal: only used by tests and
// NewBlockPersister (error cleanup). Production code does not close WALs
// at shutdown — the OS reclaims resources on process exit.
// Safe for concurrent use.
func (bp *BlockPersister) close() error { _ = "STUB: not implemented"; return nil }

// no-op persister (persistence disabled)
