package data

import (
	"context"
	"errors"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/consensus/persist"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

const blocksCacheSize = 4000

// ErrNotFound is returned when the resource is not found.
var ErrNotFound = errors.New("not found")

// ErrPruned is returned when the resource has been is pruned.
var ErrPruned = errors.New("pruned")

// Config is the config for the data State.
type Config struct {
	// Committee.
	Committee *types.Committee
	// PruneAfter is the duration after which the state prunes executed blocks.
	PruneAfter utils.Option[time.Duration]
}

// StateAPI is the interface of the State for consuming global blocks
// and reporting AppHashes.
type StateAPI interface {
	prometheus.Collector
	GlobalBlock(ctx context.Context, n types.GlobalBlockNumber) (*types.GlobalBlock, error)
	// PushAppHash blocks until block n and its QC are durably persisted,
	// ensuring AppVotes are only issued for data that survives a crash.
	PushAppHash(ctx context.Context, n types.GlobalBlockNumber, hash types.AppHash) error
}

var _ StateAPI = (*State)(nil)

// DataWAL groups the WALs used by State for crash recovery.
// Both fields are always non-nil; when stateDir is None they are no-op persisters.
// This is temporary and will go away once we switch to a proper storage solution.
type DataWAL struct {
	Blocks    *persist.GlobalBlockPersister
	CommitQCs *persist.FullCommitQCPersister
}

// Close shuts down both WALs.
func (dw *DataWAL) Close() error { _ = "STUB: not implemented"; return nil }

// TruncateBefore removes entries fully before n from both WALs in parallel.
// A crash between the two calls just leaves stale entries in one WAL,
// which are harmless on reload.
func (dw *DataWAL) TruncateBefore(n types.GlobalBlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// reconcile fixes cursor inconsistencies between the two WALs that can
// result from crashes during parallel persistence or pruning.
//
// The possible WAL states on startup and how they are handled:
//
//	Case  Blocks    QCs       Scenario                          Action
//	1     empty     empty     Fresh start                       no-op
//	2     [a,b]     empty     QCs lost (corruption)             error (statesync needed)
//	3     empty     [X,Y)     Blocks lost (crash)               Prefix: fast-forward blocks.next to X
//	4     [a,b]     [X,Y)     Normal (a=X, b<Y)                 no-op
//	5     [a,b]     [X,Y)     Prune crash: blocks ahead (a>X)   Prefix: truncate QCs to a
//	6     [a,b]     [X,Y)     Prune crash: QCs ahead (a<X)      Prefix: truncate blocks to X
//	7     [a,b]     [X,Y)     Persist crash: blocks past (b>=Y) Tail: truncate blocks to Y
//	8     [a,b]     [X,Y)     QCs ahead (normal, b<Y)           Tail: no-op (blocks catch up)
func (dw *DataWAL) reconcile(committee *types.Committee) error {
	_ = "STUB: not implemented"
	return nil

	// Fix tail: remove blocks past QC range.
}

// Blocks exist but QCs WAL is empty — data is corrupted.

// Fix prefix: align both WALs to the later start.

// NewDataWAL constructs both global-block and global-commitqc WALs.
// When stateDir is None, the returned persisters are no-ops.
func NewDataWAL(stateDir utils.Option[string], committee *types.Committee) (*DataWAL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reconcile cursor inconsistency: a crash between the two parallel
// TruncateBefore calls can leave one WAL truncated while the other
// still has stale entries. Advance both to the max starting point.

type appProposalWithTimestamp struct {
	proposal  *types.AppProposal
	timestamp time.Time
}

type inner struct {
	qcs          map[types.GlobalBlockNumber]*types.FullCommitQC      // [first,nextQC)
	blocks       map[types.GlobalBlockNumber]*types.Block             // [first,nextBlock) + subset of [nextBlock,nextQC)
	appProposals map[types.GlobalBlockNumber]appProposalWithTimestamp // [first,nextAppProposal)

	// blockHashes is a hash → height index mirroring blocks. Maintained
	// in lockstep with blocks via insertBlock / pruneFirst, so it covers
	// exactly the same retain window without a separate prune cursor or
	// startup warmup. Powers BlockByHash.
	//
	// TODO(autobahn): remove once a writer is wired into block execution
	// that populates sei-db/ledger_db/block.BlockDB. BlockDB has a built-in
	// hash index that survives process restart and lives outside Autobahn's
	// RetainHeight pruning, making this in-memory index obsolete.
	blockHashes map[types.BlockHeaderHash]types.GlobalBlockNumber

	// first <= nextAppProposal <= nextBlockToPersist <= nextBlock <= nextQC
	//
	// This invariant guarantees no race between pruning and persisting:
	// blocks are not eligible for pruning until they have an AppProposal
	// (first <= nextAppProposal), which requires persistence
	// (nextAppProposal <= nextBlockToPersist).
	first              types.GlobalBlockNumber
	nextAppProposal    types.GlobalBlockNumber
	nextBlockToPersist types.GlobalBlockNumber
	nextBlock          types.GlobalBlockNumber
	nextQC             types.GlobalBlockNumber
}

func newInner(committee *types.Committee) *inner { _ = "STUB: not implemented"; return nil }

// skipTo advances all cursors to n, discarding everything before it.
// Used on recovery when the first loaded QC starts past committee.FirstBlock()
// (i.e. data before n was pruned in a previous run).
func (i *inner) skipTo(n types.GlobalBlockNumber) { _ = "STUB: not implemented"; return }

// insertQC verifies and inserts a FullCommitQC into the inner state.
// Accepts QCs whose range starts at or before nextQC (partially pruned
// prefix is silently skipped). Rejects gaps where gr.First > nextQC.
func (i *inner) insertQC(committee *types.Committee, qc *types.FullCommitQC) error {
	_ = "STUB: not implemented"
	return nil
}

// fully behind, skip

// insertBlock inserts a pre-verified block into the inner state.
// Requires a QC to already be present for block n. Callers must verify
// the block signature before calling.
//
// insertBlock does NOT advance nextBlock — callers should call
// updateNextBlock after inserting one or more blocks. This separation
// allows batch insertion (e.g. PushQC inserts multiple blocks, then
// advances nextBlock once).
func (i *inner) insertBlock(committee *types.Committee, n types.GlobalBlockNumber, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// outside QC range

// already have it

func (i *inner) updateNextBlock(m *dataMetrics) { _ = "STUB: not implemented"; return }

func (i *inner) pruneFirst(now time.Time, m *dataMetrics) { _ = "STUB: not implemented"; return }

// State of the chain.
// Contains blocks in global order and proofs of their finality.
type State struct {
	cfg     *Config
	metrics *dataMetrics
	inner   utils.Watch[*inner]
	dataWAL *DataWAL
}

// NewState constructs a new data State.
// dataWAL persists blocks and QCs to WALs for crash recovery and provides
// preloaded data from the previous run. Use NewDataWAL to construct it.
func NewState(cfg *Config, dataWAL *DataWAL) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Fast-forward cursors to where data starts. Use blocks as golden:
	// per-block pruning may split a QC range, so blocks determine where
	// useful data starts. QCs before that are kept for verification but
	// don't set inner.first.
}

// Restore QCs. insertQC handles partially pruned QCs (range starts
// before inner.first) by skipping the pruned prefix.

// Restore blocks. Verify contiguity as defense in depth.

// outside QC range (stale from reconcile)

// Advance nextBlock through contiguous blocks. Don't use
// updateNextBlock: stale timestamps would skew metrics.

// Data loaded from WALs was already persisted in the previous run.

// WAL cursor consistency was resolved by DataWAL.reconcile at construction.
// Verify the blocks persister cursor is not behind inner.nextBlock.

// Committee returns the committee.
func (s *State) Committee() *types.Committee { _ = "STUB: not implemented"; return nil }

// PushQC pushes FullCommitQC and a subset of blocks that were finalized by it.
// Pushing the qc and blocks is atomic, so that no unnecessary GetBlock RPCs are issued.
// Even if the qc was already pushed earlier, the blocks are pushed anyway.
func (s *State) PushQC(ctx context.Context, qc *types.FullCommitQC, blocks []*types.Block) error {
	_ = "STUB: not implemented"
	// Wait until QC is needed.
	return nil
}

// Verify data.

// Atomically insert QC and blocks.

// Match blocks against stored (already verified) QC headers.

// QC returns the FullCommitQC proving finality of the block n.
func (s *State) QC(ctx context.Context, n types.GlobalBlockNumber) (*types.FullCommitQC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushBlock pushes block to the state.
// Waits until the block header is available.
func (s *State) PushBlock(ctx context.Context, n types.GlobalBlockNumber, block *types.Block) error {
	_ = "STUB: not implemented"
	// Verify outside the lock to avoid holding it during expensive crypto.
	return nil
}

// NextBlock returns the index of the next block to be pushed.
func (s *State) NextBlock() types.GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *new(types.GlobalBlockNumber)
}

// GlobalBlockByHash returns the finalized GlobalBlock whose stored header
// hashes to the given value, or None if no such block is currently in the
// retained range. The lookup-and-construct happens under a single lock so
// the returned block matches the looked-up hash atomically — pruning can't
// change which height a hash maps to between the index check and the block
// construction. Tracks the same retain window as Block / GlobalBlock since
// the hash index is maintained in lockstep by insertBlock / pruneFirst.
//
// Returns an error in the signature for forward-compat with the eventual
// switch to sei-db/ledger_db/block.BlockDB.GetBlockByHash. Today's
// in-memory implementation never errors.
//
// TODO(autobahn): when BlockDB is wired, take a ctx parameter and narrow
// the error contract — db-internal errors should surface by shutting down
// the persistence background task (matching how persistence handles errors
// today), so the query path's error stays bounded to context.Canceled.
func (s *State) GlobalBlockByHash(hash types.BlockHeaderHash) (utils.Option[*types.GlobalBlock], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Block returns the block with the given global number.
// This function is used for syncing - GlobalBlock can be derived from Block and FullCommitQC,
// which have to be fetched upfront anyway.
func (s *State) Block(ctx context.Context, n types.GlobalBlockNumber) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TryBlock returns the block with the given global number.
// Returns ErrPruned if the block has already been pruned.
// Returns ErrNotFound if the block is not available yet.
func (s *State) TryBlock(n types.GlobalBlockNumber) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// globalBlockAt assembles the GlobalBlock at height n from inner state.
// Caller must have verified n is in [inner.first, inner.nextBlock); n
// outside that range nil-derefs on inner.blocks[n] / inner.qcs[n].
func (i *inner) globalBlockAt(c *types.Committee, n types.GlobalBlockNumber) *types.GlobalBlock {
	_ = "STUB: not implemented"
	return nil
}

// GlobalBlock returns the block with the given global number.
// Returns ErrPruned if the block has already been pruned.
func (s *State) GlobalBlock(ctx context.Context, n types.GlobalBlockNumber) (*types.GlobalBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushAppHash marks blocks up to n as executed. Hash is the execution result.
// Waits for the block to be durably persisted before proceeding.
func (s *State) PushAppHash(ctx context.Context, n types.GlobalBlockNumber, hash types.AppHash) error {
	_ = "STUB: not implemented"
	return nil
}

// AppProposal returns the lowest AppProposal containing the block n.
// WARNING: currently we do not enforce all blocks to have AppProposal, therefore
// an AppProposal for a later block might be returned instead.
func (s *State) AppProposal(ctx context.Context, n types.GlobalBlockNumber) (*types.AppProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PruneBefore removes blocks, QCs, and AppProposals before retainFrom.
// Blocks at retainFrom and above are kept. Per-block pruning may split
// a QC range; this is handled on recovery (NewState skips partial QC prefixes).
func (s *State) PruneBefore(retainFrom types.GlobalBlockNumber) error {
	_ = "STUB: not implemented"
	return nil
}

// Can only prune executed blocks (those with AppProposals).

// Keep at least one entry so WALs are never empty on restart.

// Truncate WALs outside the lock to avoid holding it during disk I/O.

// runPersist is a background goroutine that persists blocks and QCs to WALs.
// It waits for in-memory data to advance past the persistence cursors, then
// writes QCs and blocks in parallel. QCs are persisted up to nextQC (eagerly),
// blocks up to nextBlock. nextBlockToPersist advances to min(persistedQC, persistedBlock)
// to unblock PushAppHash only when both are durable.
// Errors propagate vertically (kill the component).
func (s *State) runPersist(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Initialize from nextBlockToPersist, not nextQC/nextBlock. PushQC may
	// have been called before Run() (race between p2p startup and Run),
	// advancing nextQC/nextBlock beyond what's been persisted. Starting
	// from nextBlockToPersist ensures we don't skip unpersisted data.
	return nil
}

// Wait for unpersisted data and snapshot what needs writing.

// Collect deduplicated QCs for [persistedQC, nextQC).

// Collect blocks for [persistedBlock, nextBlock).

// Persist QCs and blocks in parallel.

// Advance nextBlockToPersist to where both QCs and blocks are durable.

func (s *State) runPruning(ctx context.Context, after time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Prune blocks old enough. Keep at least one entry.
// Per-block pruning may split QC ranges; handled on recovery.
// TODO: a proper fix would not prune until AppQC exists.

// Wait for at least 2 entries before retrying. Without +1,
// the loop would spin when only one entry remains (kept by
// the +1 guard above).

// Truncate WALs outside the lock to avoid holding it during disk I/O.

// Wait until the next pruning time.

// Run runs the background tasks of the data State.
// TODO(gprusak): add support for starting execution from non-zero commit QC.
func (s *State) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
