package avail

import (
	"context"
	"errors"
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/consensus/persist"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/data"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// ErrBadLane .
var ErrBadLane = errors.New("bad lane")

const BlocksPerLane = 3 * BlocksPerLanePerCommit
const BlocksPerLanePerCommit = 10

// State represents the Data Availability Plane and Ordered Event Log.
// Although it resides in a sub-package, it serves as the "source of truth" for:
// - Block data: storing and disseminating raw transaction payloads (lanes).
// - Finality tracking: acting as a persistent buffer for CommitQCs and AppQCs.
// - Pruning: managing memory by deleting data once enough execution proofs (AppVotes) are seen.
//
// NOTE: This component is more than an observer; it actively aggregates AppVotes
// to trigger internal pruning, which allows it to manage memory independently
// of the main consensus loop.
type State struct {
	key   types.SecretKey
	data  *data.State
	inner utils.Watch[*inner]

	// persisters groups all disk persistence components.
	// Always initialized: real when stateDir is set, no-op otherwise.
	persisters persisters
}

// persisters holds all disk persistence components. Either all are present
// (real I/O) or all are no-op (testing). It is a pure I/O struct — all inner
// state access goes through State methods.
type persisters struct {
	pruneAnchor persist.Persister[*pb.PersistedAvailPruneAnchor]
	blocks      *persist.BlockPersister
	commitQCs   *persist.CommitQCPersister
}

// innerFile is the A/B file prefix for avail inner state persistence.
const innerFile = "avail_inner"

// PruneAnchor is the decoded form of the persisted prune anchor
// (AppQC + matching CommitQC pair). It serves as the crash-recovery
// pruning watermark.
type PruneAnchor struct {
	AppQC    *types.AppQC
	CommitQC *types.CommitQC
}

// PruneAnchorConv converts between PruneAnchor and its protobuf representation.
var PruneAnchorConv = protoutils.Conv[*PruneAnchor, *pb.PersistedAvailPruneAnchor]{
	Encode: func(a *PruneAnchor) *pb.PersistedAvailPruneAnchor {
		return &pb.PersistedAvailPruneAnchor{
			AppQc:    types.AppQCConv.Encode(a.AppQC),
			CommitQc: types.CommitQCConv.Encode(a.CommitQC),
		}
	},
	Decode: func(p *pb.PersistedAvailPruneAnchor) (*PruneAnchor, error) {
		if p.AppQc == nil || p.CommitQc == nil {
			return nil, fmt.Errorf("incomplete prune anchor: AppQC=%v CommitQC=%v", p.AppQc != nil, p.CommitQc != nil)
		}
		appQC, err := types.AppQCConv.Decode(p.AppQc)
		if err != nil {
			return nil, fmt.Errorf("decode AppQC: %w", err)
		}
		commitQC, err := types.CommitQCConv.Decode(p.CommitQc)
		if err != nil {
			return nil, fmt.Errorf("decode CommitQC: %w", err)
		}
		return &PruneAnchor{AppQC: appQC, CommitQC: commitQC}, nil
	},
}

// loadPersistedState creates persisters for the given directory option and loads
// any existing state from disk. When dir is None, all persisters are no-op
// and no state is loaded. When a prune anchor is present, stale commitQCs and
// blocks below the anchor are filtered out before returning.
func loadPersistedState(dir utils.Option[string]) (utils.Option[*loadedAvailState], persisters, error) {
	_ = "STUB: not implemented"
	return nil, *new(persisters), nil
}

// NewState constructs a new availability state.
// stateDir is None when persistence is disabled (testing only); a no-op
// persist goroutine still runs to bump cursors without disk I/O.
func NewState(key types.SecretKey, data *data.State, stateDir utils.Option[string]) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Truncate WAL entries below the prune anchor that were filtered out by
// loadPersistedState.

func (s *State) FirstCommitQC() types.RoadIndex {
	_ = "STUB: not implemented"
	return *new(types.RoadIndex)
}

// Data returns the data state.
func (s *State) Data() *data.State {
	_ = "STUB: not implemented"

	// LastCommitQC returns receiver of the LastCommitQC.
	return nil
}

func (s *State) LastCommitQC() utils.AtomicRecv[utils.Option[*types.CommitQC]] {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) waitForCommitQC(ctx context.Context, idx types.RoadIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// LastAppQC returns the latest observed AppQC.
func (s *State) LastAppQC() utils.Option[*types.AppQC] { _ = "STUB: not implemented"; return nil }

// WaitForAppQC waits until there is an AppQC for the given index or higher.
// Returns this AppQC and the corresponding CommitQC.
// Together they provide enough information to prune the availability state.
func (s *State) WaitForAppQC(ctx context.Context, idx types.RoadIndex) (*types.AppQC, *types.CommitQC, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CommitQC returns the CommitQC for the given index.
func (s *State) CommitQC(ctx context.Context, idx types.RoadIndex) (*types.CommitQC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushCommitQC pushes a CommitQC to the state.
// Waits until all previous CommitQCs are pushed.
func (s *State) PushCommitQC(ctx context.Context, qc *types.CommitQC) error {
	_ = "STUB: not implemented"
	return nil
}

// The persist goroutine publishes latestCommitQC after writing to disk
// (or immediately for no-op persisters), so consensus won't advance
// until the CommitQC is durable.

// PushAppVote pushes an AppVote to the state.
func (s *State) PushAppVote(ctx context.Context, v *types.Signed[*types.AppVote]) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for the corresponding commitQC.

// Early exit if not useful (we collect <=1 AppQC per road index).

// Verify the vote against the CommitQC.

// Push the vote.

// PushAppQC pushes an AppQC to the state. It requires a corresponding CommitQC
// as a justification.
func (s *State) PushAppQC(appQC *types.AppQC, commitQC *types.CommitQC) error {
	_ = "STUB: not implemented"
	// Check whether it is needed before verifying.
	return nil
}

// Defense-in-depth check, it should never happen that >f validators sign
// a proposal which does not match the commitQC's global range.

// NextBlock returns the index of the next missing block in local storage for the given lane.
func (s *State) NextBlock(lane types.LaneID) types.BlockNumber {
	_ = "STUB: not implemented"
	return *new(types.BlockNumber)
}

// Block returns block n of the given lane.
// Waits until the block is available.
// Returns ErrPruned if the block has been already pruned.
func (s *State) Block(ctx context.Context, lane types.LaneID, n types.BlockNumber) (*types.Signed[*types.LaneProposal], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushBlock pushes a block to the state.
// Waits until all previous blocks are available.
func (s *State) PushBlock(ctx context.Context, p *types.Signed[*types.LaneProposal]) error {
	_ = "STUB: not implemented"
	return nil
}

// not needed any more

// Verify parent hash chain to prevent a malicious producer from
// breaking the block chain, which would deadlock header reconstruction.
// A mismatch means the producer equivocated (produced a different
// chain than we already have). We log it to aid debugging stalled
// lanes but do not return an error — the caller should not tear
// down the peer connection over an equivocating producer.
// NOTE: after pruning (q.first >= q.next), we cannot verify the parent
// hash because the previous block is gone. This is safe because
// headers() never follows the first block's parentHash in a LaneRange.

// PushVote pushes a LaneVote to the state.
// Waits until the lane has enough capacity for the new vote.
// It does NOT wait for the previous votes.
func (s *State) PushVote(ctx context.Context, vote *types.Signed[*types.LaneVote]) error {
	_ = "STUB: not implemented"
	return nil
}

// headers collects headers for the given range.
func (s *State) headers(ctx context.Context, lr *types.LaneRange) ([]*types.BlockHeader, error) {
	_ = "STUB: not implemented"
	// Empty range is always available.
	return nil, nil
}

//nolint:gosec // i is bounded by len(headers) which is a small block range; no overflow risk

// If pruned, then give up.

// Check if we have the header.

// Otherwise, wait.

func (s *State) fullCommitQC(ctx context.Context, n types.RoadIndex) (*types.FullCommitQC, error) {
	_ = "STUB: not implemented"
	// Collect the CommitQC.
	return nil, nil
}

// Collect the headers from the votes.

// WaitForCapacity waits until the given lane has enough capacity for a new block.
func (s *State) WaitForCapacity(ctx context.Context, lane types.LaneID) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForLaneQCs waits until there is at least 1 LaneQC with a block not finalized by prev.
func (s *State) WaitForLaneQCs(
	ctx context.Context, prev utils.Option[*types.CommitQC],
) (map[types.LaneID]*types.LaneQC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProduceBlock appends a new block to the producers lane.
// Blocks until the lane has enough capacity for the new block.
func (s *State) ProduceBlock(ctx context.Context, payload *types.Payload) (*types.Signed[*types.LaneProposal], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: produceBlock is a separate function for testing - consider improving the tests to use ProduceBlock only.
func (s *State) produceBlock(ctx context.Context, key types.SecretKey, payload *types.Payload) (*types.Signed[*types.LaneProposal], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run runs the background tasks of the state.
//
// Goroutines: this method spawns long-lived goroutines via scope.SpawnNamed
// (the persist loop and the FullCommitQC→data-state pusher). Inside
// runPersist, scope.Parallel spawns short-lived goroutines for concurrent
// per-lane block and commit-QC persistence. The persist package itself does
// not spawn goroutines.
func (s *State) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Task inserting FullCommitQCs and local blocks to data state.

// Collect the blocks we have locally.

// We are not expected to have all the blocks locally - only the available ones.

// We don't need to check the blocks against the headers,
// as bad blocks will be filtered out by PushQC anyway.

// runPersist is the main loop for the persist goroutine.
// Write order:
//  1. Prune anchor (AppQC + CommitQC pair) — the crash-recovery watermark (sequential).
//  2. commitQCs.MaybePruneAndPersist and each lane's blocks.MaybePruneAndPersistLane run
//     concurrently via scope.Parallel (separate WALs, no early cancellation; first error
//     is returned after all tasks finish).
//     Each path publishes (markCommitQCsPersisted / markBlockPersisted) per entry so voting
//     unblocks ASAP.
//
// The prune anchor is a pruning watermark: on restart we resume from it.
//
// TODO: use a single WAL for anchor and CommitQCs to make
// this atomic rather than relying on write order.
func (s *State) runPersist(ctx context.Context, pers persisters) error {
	_ = "STUB: not implemented"
	return nil
}

// Prune CommitQC anchor: same Option drives commit-QC WAL and per-lane block WAL
// (truncate-then-append below this QC).

// 1. Persist prune anchor first — establishes the crash-recovery watermark.

// 2. Persist commit-QCs and per-lane blocks in parallel.
// Callees handle empty inputs gracefully (no-op when nothing to write/truncate).

// persistBatch holds the data collected under lock for one persist iteration.
type persistBatch struct {
	blocks      []*types.Signed[*types.LaneProposal]
	commitQCs   []*types.CommitQC
	pruneAnchor utils.Option[*PruneAnchor]
}

// advancePersistedBlockStart updates the per-lane block admission watermark
// after durably writing the prune anchor. This unblocks PushBlock/ProduceBlock
// waiters that are gated on persistedBlockStart + BlocksPerLane.
func (s *State) advancePersistedBlockStart(commitQC *types.CommitQC) {
	_ = "STUB: not implemented"
	return
}

// markBlockPersisted advances the per-lane block persistence cursor.
// Called after each block is persisted so that RecvBatch (and therefore
// voting) can unblock as soon as the block is durable. Safe for concurrent
// callers (acquires s.inner lock internally).
func (s *State) markBlockPersisted(lane types.LaneID, next types.BlockNumber) {
	_ = "STUB: not implemented"
	return
}

// markCommitQCsPersisted publishes the latest persisted CommitQC,
// gating consensus from advancing until the QC is durable.
func (s *State) markCommitQCsPersisted(qc *types.CommitQC) { _ = "STUB: not implemented"; return }

// collectPersistBatch waits for new blocks or commitQCs and collects them under lock.
func (s *State) collectPersistBatch(ctx context.Context, lastPersistedAppQCNext types.RoadIndex) (persistBatch, error) {
	_ = "STUB: not implemented"
	return *new(persistBatch), nil
}

// Derive the CommitQC persist cursor from latestCommitQC. This is
// safe because latestCommitQC is only advanced by markCommitQCsPersisted
// (after disk write) and on startup (from disk). prune() does NOT
// update latestCommitQC, so this always reflects persistence state.
// The max clamp with commitQCs.first handles the case where prune()
// fast-forwarded the queue past the cursor.
