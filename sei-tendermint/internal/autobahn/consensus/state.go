package consensus

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/avail"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/consensus/persist"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/data"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// ViewTimeoutFunc is a function that specifies the timeout for the given view.
// - constant for production
// - custom for tests.
type ViewTimeoutFunc = func(types.View) time.Duration

// Config holds the configuration for the consensus state.
type Config struct {
	Key         types.SecretKey
	ViewTimeout ViewTimeoutFunc
	// PersistentStateDir is the directory where the consensus state is persisted.
	// If None, persistence is disabled - DANGEROUS, may lead to SLASHING on restart.
	PersistentStateDir utils.Option[string]
}

// State represents the high-level Consensus Control Plane.
// It is responsible for:
// - View management: tracking rounds and leader election.
// - Voting: aggregating signatures for Prepare, Commit, and Timeout phases.
// - Proposals: constructing and verifying block proposals.
//
// NOTE: While this is the "brain", it relies on the "avail" package as its
// primary data store and synchronization sequencer.
type State struct {
	cfg   *Config
	avail *avail.State
	// metrics *Metrics
	inner     utils.Mutex[*utils.AtomicSend[inner]]
	innerRecv utils.AtomicRecv[inner]

	// persister writes inner's persistedInner to disk when PersistentStateDir is set; None when disabled.
	persister utils.Option[persist.Persister[*pb.PersistedInner]]

	timeoutVotes utils.Mutex[*timeoutVotes]
	prepareVotes utils.Mutex[*prepareVotes]
	commitVotes  utils.Mutex[*commitVotes]

	myView        utils.AtomicSend[types.ViewSpec]
	myProposal    utils.AtomicSend[utils.Option[*types.FullProposal]]
	myPrepareVote utils.AtomicSend[utils.Option[*types.ConsensusReqPrepareVote]]
	myCommitVote  utils.AtomicSend[utils.Option[*types.ConsensusReqCommitVote]]
	myTimeoutVote utils.AtomicSend[utils.Option[*types.FullTimeoutVote]]
	myTimeoutQC   utils.AtomicSend[utils.Option[*types.TimeoutQC]]
}

// TODO: replace with a single ConsensusMsg stream.
func (s *State) SubscribeProposal() utils.AtomicRecv[utils.Option[*types.FullProposal]] {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) SubscribePrepareVote() utils.AtomicRecv[utils.Option[*types.ConsensusReqPrepareVote]] {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) SubscribeCommitVote() utils.AtomicRecv[utils.Option[*types.ConsensusReqCommitVote]] {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) SubscribeTimeoutVote() utils.AtomicRecv[utils.Option[*types.FullTimeoutVote]] {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) SubscribeTimeoutQC() utils.AtomicRecv[utils.Option[*types.TimeoutQC]] {
	_ = "STUB: not implemented"
	return nil
}

// NewState constructs a new state.
func NewState(cfg *Config, data *data.State) (*State, error) {
	_ = "STUB: not implemented"
	// Create persister first so newInner can receive the loaded data
	// instead of reading the files directly.
	return nil, nil
}

// newState is the internal constructor exposed for tests that need to inject
// a custom persister (e.g. a failing mock). Production code should use NewState.
func newState(
	cfg *Config,
	data *data.State,
	pers utils.Option[persist.Persister[*pb.PersistedInner]],
	persistedData utils.Option[*pb.PersistedInner],
) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// metrics: NewMetrics(),

func (s *State) timeoutQC() utils.AtomicRecv[utils.Option[*types.TimeoutQC]] {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) prepareQC() utils.AtomicRecv[utils.Option[*types.PrepareQC]] {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) commitQC() utils.AtomicRecv[utils.Option[*types.CommitQC]] {
	_ = "STUB: not implemented"
	return nil
}

// WaitForCapacity waits until a new block can be produced by this node.
func (s *State) WaitForCapacity(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// ProduceBlock produces a new block with the given payload.
// Returns ErrNoCapacity if there is currently no capacity for the next block.
// Run WaitForCapacity before calling ProduceBlock.
func (s *State) ProduceBlock(ctx context.Context, payload *types.Payload) (*types.Signed[*types.LaneProposal], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushProposal processes an unverified FullProposal message.
func (s *State) PushProposal(ctx context.Context, proposal *types.FullProposal) error {
	_ = "STUB: not implemented"
	return nil
}

// PushTimeoutQC processes an unverified TimeoutQC message.
func (s *State) PushTimeoutQC(ctx context.Context, qc *types.TimeoutQC) error {
	_ = "STUB: not implemented"
	return nil
}

// PushPrepareVote processes an unverified Prepare vote message.
func (s *State) PushPrepareVote(vote *types.Signed[*types.PrepareVote]) error {
	_ = "STUB: not implemented"
	return nil
}

// PushCommitVote processes an unverified CommitVote message.
func (s *State) PushCommitVote(vote *types.Signed[*types.CommitVote]) error {
	_ = "STUB: not implemented"
	return nil
}

// PushTimeoutVote processes an unverified FullTimeoutVote message.
func (s *State) PushTimeoutVote(vote *types.FullTimeoutVote) error {
	_ = "STUB: not implemented"
	return nil
}

// Data is the underlying data state.
func (s *State) Data() *data.State { _ = "STUB: not implemented"; return nil }
func (s *State) Avail() *avail.State {
	_ = "STUB: not implemented"

	// Constructs new proposals.
	return nil
}

func (s *State) runPropose(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// not the leader.

// Try repropose.

// Wait for laneQCs.

// Construct a full proposal.

func updateOutput[T types.ConsensusReq](w *utils.AtomicSend[utils.Option[T]], v T) {
	_ = "STUB: not implemented"
	return
}

// Updates the outputs based on the inner state.
// Persists state to disk before broadcasting votes to ensure votes are durable
// before dissemination (prevents double-voting on crash).
// myView update is safe before persist — it only triggers proposing and timeout
// timers, neither of which constitutes a vote.
func (s *State) runOutputs(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Persist to disk before broadcasting votes to the network.

// Run runs the background processes of the consensus state.
func (s *State) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// s.metrics.ObserveCommitQC(qc)
// We push the locally generated CommitQC into "avail" to act as a
// sequencer and to trigger data pruning.

// We pull the CommitQC back from "avail" for dissemination. This ensures
// that we only push CommitQCs that have been successfully "logged" and
// sequenced by the availability layer.
