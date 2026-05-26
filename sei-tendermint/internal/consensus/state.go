package consensus

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.opentelemetry.io/otel/sdk/trace"
	otrace "go.opentelemetry.io/otel/trace"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	cstypes "github.com/sei-protocol/sei-chain/sei-tendermint/internal/consensus/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	tmtime "github.com/sei-protocol/sei-chain/sei-tendermint/libs/time"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// Consensus sentinel errors
var (
	ErrInvalidProposalSignature     = errors.New("error invalid proposal signature")
	ErrInvalidProposer              = errors.New("error invalid proposer")
	ErrInvalidHeaderProposer        = errors.New("error invalid header proposer")
	ErrAddingVote                   = errors.New("error adding vote")
	ErrSignatureFoundInPastBlocks   = errors.New("found signature from the same key")
	ErrInvalidProposalPartSetHeader = errors.New("error invalid proposal part set header")

	errPubKeyIsNotSet = errors.New("pubkey is not set. Look for \"Can't get private validator pubkey\" errors")
)

var msgQueueSize = 1000
var heartbeatInterval = 10 * time.Second

// msgs from the reactor which may update the state
type msgInfo struct {
	Msg         Message
	PeerID      types.NodeID
	ReceiveTime time.Time
}

// internally generated messages which may update the state
type timeoutInfo struct {
	Duration time.Duration         `json:"duration,string"`
	Height   int64                 `json:"height,string"`
	Round    int32                 `json:"round"`
	Step     cstypes.RoundStepType `json:"step"`
}

func (ti *timeoutInfo) Less(b *timeoutInfo) bool {
	_ = "STUB: not implemented"
	// sort by height, then round, then step
	return false
}

// This is copy-pasted logic, supposedly allowing for updating the timeout with Step 0 without incrementing the step.
// Note that because of this Less is NOT a strict order.
// TODO(gprusak): Figure out why we special case step 0 and fix it.

func (ti *timeoutInfo) String() string { _ = "STUB: not implemented"; return "" }

// interface to the evidence pool
type evidencePool interface {
	// reports conflicting votes to the evidence pool to be processed into evidence
	ReportConflictingVotes(voteA, voteB *types.Vote)
}

// State handles execution of the consensus algorithm.
// It processes votes and proposals, and upon reaching agreement,
// commits blocks to the chain and executes them against the application.
// The internal state machine receives input from peers, the internal validator, and from a timer.
type State struct {
	// config details
	config        *config.ConsensusConfig
	privValidator utils.Option[types.PrivValidator] // for signing votes
	// privValidator pubkey, memoized for the duration of one block
	// to avoid extra requests to HSM
	privValidatorPubKey utils.Option[crypto.PubKey]

	// store blocks and commits
	blockStore sm.BlockStore

	stateStore sm.Store

	// create and execute blocks
	blockExec *sm.BlockExecutor

	// notify us if txs are available
	txMempool *mempool.TxMempool

	// add evidence to the pool
	// when it's detected
	evpool evidencePool

	// internal state
	mtx        sync.RWMutex
	roundState cstypes.SafeRoundState
	state      sm.State // State until height-1.

	// state changes may be triggered by: msgs from peers,
	// msgs from ourself, or by timeouts
	peerMsgQueue     chan msgInfo
	internalMsgQueue chan msgInfo
	timeoutTicker    TimeoutTicker

	// we use eventBus to trigger msg broadcasts in the reactor,
	// and to notify external subscribers, eg. through a websocket
	eventBus *eventbus.EventBus

	// a Write-Ahead Log ensures we can recover from any kind of crash
	// and helps us avoid signing conflicting votes
	wal          *WAL
	replayMode   bool // so we don't log signing errors during replay
	doWALCatchup bool // determines if we even try to do the catchup

	// for tests where we want to limit the number of transitions the state makes
	nSteps int

	// some functions can be overwritten for testing
	doPrevote   func(ctx context.Context, height int64, round int32)
	setProposal func(proposal *types.Proposal, t time.Time) error

	// synchronous pubsub between consensus state and reactor.
	// eventValidBlock is emitting a copy of round state, in which the
	// block parts will be collected, so it should not be treated as immutable.
	eventValidBlock   utils.AtomicSend[utils.Option[*cstypes.RoundState]]
	eventNewRoundStep func(state *cstypes.RoundState)
	eventVote         func(vote *types.Vote)
	eventMsg          func(msgInfo)

	// for reporting metrics
	metrics *Metrics

	tracer            otrace.Tracer
	heightSpan        otrace.Span
	heightBeingTraced int64
	tracingCtx        context.Context
}

// NewState returns a new State.
func NewState(
	cfg *config.ConsensusConfig,
	wal *WAL,
	store sm.Store,
	blockExec *sm.BlockExecutor,
	blockStore sm.BlockStore,
	txMempool *mempool.TxMempool,
	evpool evidencePool,
	eventBus *eventbus.EventBus,
	traceProviderOps []trace.TracerProviderOption,
	metrics *Metrics,
) *State {
	_ = "STUB: not implemented"
	return nil
}

// set function defaults (may be overwritten before calling Start)

func (cs *State) updateStateFromStore() error { _ = "STUB: not implemented"; return nil }

// if the new state is equivalent to the old state, we should not trigger a state update.

// We have no votes, so reconstruct LastCommit from SeenCommit.

// String returns a string.
func (cs *State) String() string {
	_ = "STUB: not implemented"
	// better not to access shared variables
	return ""
}

// GetState returns a copy of the chain state.
// TESTONLY.
func (cs *State) GetState() sm.State { _ = "STUB: not implemented"; return *new(sm.State) }

// GetLastHeight returns the last height committed.
// If there were no blocks, returns 0.
func (cs *State) GetLastHeight() int64 { _ = "STUB: not implemented"; return 0 }

// GetRoundState returns a shallow copy of the internal consensus state.
func (cs *State) GetRoundState() *cstypes.RoundState { _ = "STUB: not implemented"; return nil }

// GetRoundStateJSON returns a json of RoundState. UNSTABLE.
func (cs *State) GetRoundStateJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetRoundStateSimpleJSON returns a json of RoundStateSimple. UNSTABLE.
func (cs *State) GetRoundStateSimpleJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetValidators returns a copy of the current validators.
func (cs *State) GetValidators() (int64, []*types.Validator) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SetPrivValidator sets the private validator account for signing votes. It
// immediately requests pubkey and caches it.
func (cs *State) SetPrivValidator(ctx context.Context, priv utils.Option[types.PrivValidator]) {
	_ = "STUB: not implemented"
	return
}

// SetTimeoutTicker sets the local timer. It may be useful to overwrite for
// testing.
func (cs *State) SetTimeoutTicker(timeoutTicker TimeoutTicker) { _ = "STUB: not implemented"; return }

// LoadCommit loads the commit for a given height.
func (cs *State) LoadCommit(height int64) *types.Commit { _ = "STUB: not implemented"; return nil }

// NOTE: Retrieving the height of the most recent block and retrieving
// the most recent commit does not currently occur as an atomic
// operation. We check the height and commit here in case a more recent
// commit has arrived since retrieving the latest height.

// Run loads the latest state via the WAL, and starts the timeout and
// receive routines.
func (cs *State) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// we need the timeoutRoutine for replay so
// we don't block on the tick chan.

// We may have lost some votes if the process crashed reload from consensus
// log to catchup.

// Double Signing Risk Reduction

// now start the receiveRoutine

// schedule the first round!
// use GetRoundState so we don't race the receiveRoutine for access

//------------------------------------------------------------
// Public interface for passing messages into the consensus state, possibly causing a state transition.
// If peerID == "", the msg is considered internal.
// Messages are added to the appropriate queue (peer or internal).
// If the queue is full, the function may block.
// TODO: should these return anything or let callers just use events?

// AddVote inputs a vote.
func (cs *State) AddVote(ctx context.Context, vote *types.Vote, peerID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: wait for event?!

// SetProposal inputs a proposal.
func (cs *State) SetProposal(ctx context.Context, proposal *types.Proposal, peerID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: wait for event?!

// AddProposalBlockPart inputs a part of the proposal block.
func (cs *State) AddProposalBlockPart(ctx context.Context, height int64, round int32, part *types.Part, peerID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: wait for event?!

// SetProposalAndBlock inputs the proposal and all block parts.
func (cs *State) SetProposalAndBlock(
	ctx context.Context,
	proposal *types.Proposal,
	block *types.Block,
	parts *types.PartSet,
	peerID types.NodeID,
) error {
	_ = "STUB: not implemented"
	return nil
}

//------------------------------------------------------------
// internal functions for managing the state

func (cs *State) updateHeight(height int64) { _ = "STUB: not implemented"; return }

func (cs *State) updateRoundStep(round int32, step cstypes.RoundStepType) {
	_ = "STUB: not implemented"
	return
}

// enterNewRound(height, 0) at cs.StartTime.
func (cs *State) scheduleRound0(rs *cstypes.RoundState) { _ = "STUB: not implemented"; return }

// Attempt to schedule a timeout (by sending timeoutInfo on the tickChan)
func (cs *State) scheduleTimeout(duration time.Duration, height int64, round int32, step cstypes.RoundStepType) {
	_ = "STUB: not implemented"
	return
}

// send a msg into the receiveRoutine regarding our own proposal, block part, or vote
func (cs *State) sendInternalMessage(ctx context.Context, mi msgInfo) {
	_ = "STUB: not implemented"
	return
}

// NOTE: using the go-routine means our votes can
// be processed out of order.
// TODO: use CList here for strict determinism and
// attempt push to internalMsgQueue in receiveRoutine

// Reconstruct the LastCommit from the SeenCommit. SeenCommit
// is saved along with the block.
func (cs *State) reconstructLastCommit(state sm.State) { _ = "STUB: not implemented"; return }

func (cs *State) votesFromSeenCommit(state sm.State) (*types.VoteSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates State and increments height to match that of state.
// The round becomes 0 and cs.Step becomes cstypes.RoundStepNewHeight.
func (cs *State) updateToState(state sm.State) { _ = "STUB: not implemented"; return }

// This might happen when someone else is mutating cs.state.
// Someone forgot to pass in state.Copy() somewhere?!

// If state isn't further out than cs.state, just ignore.
// This happens when SwitchToConsensus() is called in the reactor.
// We don't want to reset e.g. the Votes, but we still want to
// signal the new round step, because other services (eg. txMempool)
// depend on having an up-to-date peer state!

// Reset fields based on state.

// Very first commit should be empty.

// Otherwise, use cs.Votes

// NOTE: when Tendermint starts, it has no votes. reconstructLastCommit
// must be called to reconstruct LastCommit from SeenCommit.

// Next desired block height

// RoundState fields

// "Now" makes it easier to sync up dev nodes.
// We add timeoutCommit to allow transactions
// to be gathered for the first block.
// And alternative solution that relies on clocks:
// cs.StartTime = state.LastBlockTime.Add(timeoutCommit)

// Reset the valid block message, since we no longer need block parts
// from the previous height. This is just for clarity - it wouldn't hurt
// to just keep the value from the previous height.

// Finally, broadcast RoundState

func (cs *State) newStep() { _ = "STUB: not implemented"; return }

// newStep is called by updateToState in NewState before the eventBus is set!

func (cs *State) heartbeater(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//-----------------------------------------
// the main go routines

// receiveRoutine handles messages which may cause state transitions.
// it's argument (n) is the number of messages to process before exiting - use 0 to run forever
// It keeps the RoundState and is the only thing that updates it.
// Updates (state transitions) happen on timeouts, complete proposals, and 2/3 majorities.
// State must be locked before any internal state is updated.
func (cs *State) receiveRoutine(ctx context.Context, maxSteps int) error {
	_ = "STUB: not implemented"
	return nil
}

// There are a couple of cases where the we
// panic with an error from deeper within the
// state machine and in these cases, typically
// during a normal shutdown, we can continue
// with normal shutdown with safety. These
// cases are:

// don't re-panic if the panic is just an
// error and we're already trying to shut down

// Re-panic to ensure the node terminates.
//

// Channel signaling that transactions are available.
// nil (blocks forever) if waiting for transactions is disabled.

// handles proposals, block parts, votes
// may generate internal events (votes, complete proposals, 2/3 majorities)

// handles proposals, block parts, votes

// tockChan:

// if the timeout is relevant to the rs
// go to the next step

// TODO should we handle context cancels here?

func (cs *State) fsyncAndCompleteProposal(ctx context.Context, fsyncUponCompletion bool, height int64, span otrace.Span, onPropose bool) {
	_ = "STUB: not implemented"
	return
}

// fsync

// state transitions on complete-proposal, 2/3-any, 2/3-one
func (cs *State) handleMsg(ctx context.Context, mi msgInfo, fsyncUponCompletion bool) {
	_ = "STUB: not implemented"
	return
}

// will not cause transition.
// once proposal is set, we can receive block parts

// If we have already created block parts, we can exit early if block part matches

// Check hash proof matches. If so, we can return

// if the proposal is complete, we'll enterPrevote or tryFinalizeCommit

// We unlock here to yield to any routines that need to read the the RoundState.
// Previously, this code held the lock from the point at which the final block
// part was received until the block executed against the application.
// This prevented the reactor from being able to retrieve the most updated
// version of the RoundState. The reactor needs the updated RoundState to
// gossip the now completed block.
//
// This code can be further improved by either always operating on a copy
// of RoundState and only locking when switching out State's copy of
// RoundState with the updated copy or by emitting RoundState events in
// more places for routines depending on it to listen for.

// attempt to add the vote and dupeout the validator if its a duplicate signature
// if the vote gives us a 2/3-any or 2/3-one, we transition

// TODO: punish peer
// We probably don't want to stop the peer here. The vote does not
// necessarily comes from a malicious peer but can be just broadcasted by
// a typical peer.
// https://github.com/tendermint/tendermint/issues/1281

// NOTE: the vote is broadcast to peers by the reactor listening
// for vote events

// TODO: If rs.Height == vote.Height && rs.Round < vote.Round,
// the peer is sending us CatchupCommit precommits.
// We could make note of this and help filter in broadcastHasVoteMessage().

func (cs *State) handleTimeout(
	ctx context.Context,
	ti timeoutInfo,
	rs cstypes.RoundState,
) {
	_ = "STUB: not implemented"
	return
}

// timeouts must be for current height, round, step

// the timeout will now cause a state transition

// NewRound event fired from enterNewRound.
// XXX: should we fire timeout here (for timeout commit)?

func (cs *State) handleTxsAvailable(ctx context.Context) { _ = "STUB: not implemented"; return }

// We only need to do this for round 0.

// timeoutCommit phase

// enterPropose will be called by enterNewRound

// +1ms to ensure RoundStepNewRound timeout always happens after RoundStepNewHeight

// after timeoutCommit

func (cs *State) getTracingCtx(defaultCtx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

//-----------------------------------------------------------------------------
// State functions
// Used internally by handleTimeout and handleMsg to make state transitions

// Enter: `timeoutNewHeight` by startTime (commitTime+timeoutCommit),
//
//	or, if SkipTimeoutCommit==true, after receiving all precommits from (height,round-1)
//
// Enter: `timeoutPrecommits` after any +2/3 precommits from (height,round-1)
// Enter: +2/3 precommits for nil at (height,round-1)
// Enter: +2/3 prevotes any or +2/3 precommits for block or any from (height, round)
// NOTE: cs.StartTime was already set for height.
func (cs *State) enterNewRound(ctx context.Context, height int64, round int32, entryLabel string) {
	_ = "STUB: not implemented"
	return
}

// TODO: remove panics in this function and return an error

// increment validators if necessary

// Setup new round
// we don't fire newStep for this step,
// but we fire an event, so update the round step first

// We've already reset these upon new height,
// and meanwhile we might have received a proposal
// for round 0.

// also track next round (round+1) to allow round-skipping

// Wait for txs to be available in the mempool
// before we enterPropose in round 0. If the last block changed the app hash,
// we may need an empty "proof" block, and enterPropose immediately.

// needProofBlock returns true on the first height (so the genesis app hash is signed right away)
// and where the last block (height-1) caused the app hash to change
func (cs *State) needProofBlock(height int64) bool { _ = "STUB: not implemented"; return false }

// Enter (CreateEmptyBlocks): from enterNewRound(height,round)
// Enter (CreateEmptyBlocks, CreateEmptyBlocksInterval > 0 ):
//
//	after enterNewRound(height,round), after timeout of CreateEmptyBlocksInterval
//
// Enter (!CreateEmptyBlocks) : after enterNewRound(height,round), once txs are in the mempool
func (cs *State) enterPropose(ctx context.Context, height int64, round int32, entryLabel string) {
	_ = "STUB: not implemented"
	return
}

// If this validator is the proposer of this round, and the previous block time is later than
// our local clock time, wait to propose until our local clock time has passed the block time.

// Done enterPropose:

// If we have the whole proposal + POL, then goto Prevote now.
// else, we'll enterPrevote when the rest of the proposal is received (in AddProposalBlockPart),
// or else after timeoutPropose

// Do not count enterPrevote latency into enterPropose latency

// If we don't get the proposal and all block parts quick enough, enterPrevote

// Nothing more to do if we're not a validator

// If this node is a validator & proposer in the current round, it will
// miss the opportunity to create a block.

// if not a validator, we're done

func (cs *State) decideProposal(ctx context.Context, height int64, round int32, privValidator types.PrivValidator, privValidatorPubKey crypto.PubKey) {
	_ = "STUB: not implemented"
	return
}

// Decide on block

// If there is valid block, choose that.

// Create a new proposal block from state/txs from the mempool.

// Flush the WAL. Otherwise, we may not recompute the same proposal to sign,
// and the privValidator will refuse to sign anything.

// Make proposal

// wait the max amount we would wait for a proposal

// send proposal and block parts on internal msg queue

// Returns true if the proposal block is complete &&
// (if POLRound was proposed, we have +2/3 prevotes from there).
func (cs *State) isProposalComplete() bool { _ = "STUB: not implemented"; return false }

// we have the proposal. if there's a POLRound,
// make sure we have the prevotes from it too

// if this is false the proposer is lying or we haven't received the POL yet

// Create the next block to propose and return it. Returns nil block upon error.
//
// We really only need to return the parts, but the block is returned for
// convenience so we can log the proposal block.
//
// NOTE: keep it side-effect free for clarity.
// CONTRACT: cs.privValidator is not nil.
func (cs *State) createProposalBlock(ctx context.Context) (block *types.Block, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert panic to error

// TODO(sergio): wouldn't it be easier if CreateProposalBlock accepted cs.LastCommit directly?

// We're creating a proposal for the first block.
// The commit is empty, but not nil.

// Make the commit from LastCommit

// This shouldn't happen.

// If this node is a validator & proposer in the current round, it will
// miss the opportunity to create a block.

// Instead of panicking, return the error which will be caught by our defer recovery

// Enter: `timeoutPropose` after entering Propose.
// Enter: proposal block and POL is ready.
// If we received a valid proposal within this round and we are not locked on a block,
// we will prevote for block.
// Otherwise, if we receive a valid proposal that matches the block we are
// locked on or matches a block that received a POL in a round later than our
// locked round, prevote for the proposal, otherwise vote nil.
func (cs *State) enterPrevote(ctx context.Context, height int64, round int32, entryLabel string) {
	_ = "STUB: not implemented"
	return
}

// Done enterPrevote:

// Sign and broadcast vote as necessary

// Once `addVote` hits any +2/3 prevotes, we will go to PrevoteWait
// (so we have more time to try and collect +2/3 prevotes for a single block)

func (cs *State) proposalIsTimely() bool { _ = "STUB: not implemented"; return false }

func (cs *State) defaultDoPrevote(ctx context.Context, height int64, round int32) {
	_ = "STUB: not implemented"
	return
}

// Check that a proposed block was not received within this round (and thus executing this from a timeout).

// Attempt to reconstruct block, in case more transactions have arrived to mempool.

// Validate proposal block, from Tendermint's perspective

// ProposalBlock is invalid, prevote nil.

/*
	The block has now passed Tendermint's validation rules.
	Before prevoting the block received from the proposer for the current round and height,
	we request the Application, via the ProcessProposal, ABCI call to confirm that the block is
	valid. If the Application does not accept the block, Tendermint prevotes nil.

	WARNING: misuse of block rejection by the Application can seriously compromise Tendermint's
	liveness properties. Please see PrepareProposal-ProcessProposal coherence and determinism
	properties in the ABCI++ specification.
*/

// Vote nil if the Application rejected the block

/*
	22: upon <PROPOSAL, h_p, round_p, v, −1> from proposer(h_p, round_p) while step_p = propose do
	23: if valid(v) && (lockedRound_p = −1 || lockedValue_p = v) then
	24: broadcast <PREVOTE, h_p, round_p, id(v)>

	Here, cs.Proposal.POLRound corresponds to the -1 in the above algorithm rule.
	This means that the proposer is producing a new proposal that has not previously
	seen a 2/3 majority by the network.

	If we have already locked on a different value that is different from the proposed value,
	we prevote nil since we are locked on a different value. Otherwise, if we're not locked on a block
	or the proposal matches our locked block, we prevote the proposal.
*/

/*
	28: upon <PROPOSAL, h_p, round_p, v, v_r> from proposer(h_p, round_p) AND 2f + 1 <PREVOTE, h_p, v_r, id(v)> while
	step_p = propose && (v_r ≥ 0 && v_r < round_p) do
	29: if valid(v) && (lockedRound_p ≤ v_r || lockedValue_p = v) then
	30: broadcast <PREVOTE, h_p, round_p, id(v)>

	This rule is a bit confusing but breaks down as follows:

	If we see a proposal in the current round for value 'v' that lists its valid round as 'v_r'
	AND this validator saw a 2/3 majority of the voting power prevote 'v' in round 'v_r', then we will
	issue a prevote for 'v' in this round if 'v' is valid and either matches our locked value OR
	'v_r' is a round greater than or equal to our current locked round.

	'v_r' can be a round greater than to our current locked round if a 2/3 majority of
	the network prevoted a value in round 'v_r' but we did not lock on it, possibly because we
	missed the proposal in round 'v_r'.
*/

// Enter: any +2/3 prevotes at next round.
func (cs *State) enterPrevoteWait(height int64, round int32) { _ = "STUB: not implemented"; return }

// Done enterPrevoteWait:

// Wait for some more prevotes; enterPrecommit

// Enter: `timeoutPrevote` after any +2/3 prevotes.
// Enter: `timeoutPrecommit` after any +2/3 precommits.
// Enter: +2/3 precomits for block or nil.
// Lock & precommit the ProposalBlock if we have enough prevotes for it (a POL in this round)
// else, precommit nil otherwise.
func (cs *State) enterPrecommit(ctx context.Context, height int64, round int32, entryLabel string) {
	_ = "STUB: not implemented"
	return
}

// Done enterPrecommit:

// check for a polka

// If we don't have a polka, we must precommit nil.

// At this point +2/3 prevoted for a particular block or nil.

// the latest POLRound should be this round.

// +2/3 prevoted nil. Precommit nil.

// At this point, +2/3 prevoted for a particular block.

// If we never received a proposal for this block, we must precommit nil

// If the proposal time does not match the block time, precommit nil.

// If we're already locked on that block, precommit it, and update the LockedRound

// If greater than 2/3 of the voting power on the network prevoted for
// the proposed block, update our locked block to this block and issue a
// precommit vote for it.

// Validate the block.

// There was a polka in this round for a block we don't have.
// Fetch that block, and precommit nil.

// Enter: any +2/3 precommits for next round.
func (cs *State) enterPrecommitWait(height int64, round int32) { _ = "STUB: not implemented"; return }

// Done enterPrecommitWait:

// wait for some more precommits; enterNewRound

// Enter: +2/3 precommits for block
func (cs *State) enterCommit(ctx context.Context, height int64, commitRound int32, entryLabel string) {
	_ = "STUB: not implemented"
	return
}

// Done enterCommit:
// keep cs.Round the same, commitRound points to the right Precommits set.

// Maybe finalize immediately.

// The Locked* fields no longer matter.
// Move them over to ProposalBlock if they match the commit hash,
// otherwise they'll be cleared in updateToState.

// If we don't have the block being committed, set up to get it.

// We're getting the wrong block.
// Set up ProposalBlockParts, clear ProposalBlock and keep waiting for the parts.

// If we have the block AND +2/3 commits for it, finalize.
func (cs *State) tryFinalizeCommit(ctx context.Context, height int64) {
	_ = "STUB: not implemented"
	return
}

// TODO: this happens every time if we're not a validator (ugly logs)
// TODO: ^^ wait, why does it matter that we're a validator?

// Increment height and goto cstypes.RoundStepNewHeight
func (cs *State) finalizeCommit(ctx context.Context, height int64) {
	_ = "STUB: not implemented"
	return
}

// Save to blockStore.

// NOTE: the seenCommit is local justification to commit this block,
// but may differ from the LastCommit included in the next block

// Calculate consensus time

// Happens during replay if we already saved the block but didn't commit

// Write EndHeightMessage{} for this height, implying that the blockstore
// has saved the block.
//
// If we crash before writing this EndHeightMessage{}, we will recover by
// running ApplyBlock during the ABCI handshake when we restart.  If we
// didn't save the block to the blockstore before writing
// EndHeightMessage{}, we'd have to change WAL replay -- currently it
// complains about replaying for heights where an #ENDHEIGHT entry already
// exists.
//
// Either way, the State should not be resumed until we
// successfully call ApplyBlock (ie. later here, or in Handshake after
// restart).

// Create a copy of the state for staging and an event cache for txs.

// Execute and commit the block, update and save the state, and update the mempool.
// NOTE The block.AppHash won't reflect these txs until the next block.

// must be called before we update state

// NewHeightStep!

// Private validator might have changed it's key pair => refetch pubkey.

// cs.StartTime is already set.
// Schedule Round0 to start soon.

// By here,
// * cs.Height has been increment to height+1
// * cs.Step is now cstypes.RoundStepNewHeight
// * cs.StartTime is set to when we will start round0.

func (cs *State) RecordMetrics(height int64, block *types.Block) { _ = "STUB: not implemented"; return }

// height=0 -> MissingValidators and MissingValidatorsPower are both 0.
// Remember that the first LastCommit is intentionally empty, so it's not
// fair to increment missing validators number.

// Sanity check that commit size matches validator set size - only applies
// after first block.

// Metrics won't be updated, but it's not critical.

// NOTE: byzantine validators power and count is only for consensus evidence i.e. duplicate vote

// Block Interval metric

// Latency metric for prevote delay

//nolint:gosec // ValidRound is a small consensus round number

//-----------------------------------------------------------------------------

func (cs *State) defaultSetProposal(proposal *types.Proposal, recvTime time.Time) error {
	_ = "STUB: not implemented"
	// Already have one
	// TODO: possibly catch double proposals
	return nil
}

// Preemptively re-verify the proposal.

// Does not apply

// If we already know the commit block for this height, ignore proposals that don't match it.

// Verify signature

// We don't update cs.ProposalBlockParts if it is already set.
// This happens if we're already in cstypes.RoundStepCommit or if there is a valid block in the current round.
// TODO: We can check if Proposal is for a different block as this is a sign of misbehavior!

// apply the same check as in SetHasProposal

// NOTE: block is not necessarily valid.
// Asynchronously triggers either enterPrevote (before we timeout of propose) or tryFinalizeCommit,
// once we have the full block.
func (cs *State) addProposalBlockPart(
	msg *BlockPartMessage,
	peerID types.NodeID,
) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Blocks might be reused, so round mismatch is OK

// We're not expecting a block part.

// NOTE: this can happen when we've gone to a higher round and
// then receive parts from the previous round - not necessarily a bad peer.

// NOTE: it's possible to receive complete proposal blocks for future rounds without having the proposal

func (cs *State) getBlockFromBlockParts() (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *State) tryCreateProposalBlock(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Block already constructed.

// NOTE: it's possible to receive complete proposal blocks for future rounds without having the proposal

// If we just have all the parts, reconstruct the block.

// This can happen if the BlockParts header is broken.

// Attempt to reconstruct from the Proposal.TxHashes.

// For some reason we only attempt that on non-proposing validators.

// Constructed block needs to match the expected parts.
// This check is optimistic, because proposer may provide mismatching PartSetHeader.

// Construct block and block parts.

// Now check if parts were actually expected.

// Build a proposal block from mempool txs. If cs.config.GossipTransactionKeyOnly=true
// proposals only contain txHashes so we rebuild the block using mempool txs
func (cs *State) buildProposalBlock(proposal *types.Proposal) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

func (cs *State) handleCompleteProposal(ctx context.Context, height int64, handleBlockPartSpan otrace.Span) {
	_ = "STUB: not implemented"
	// Update Valid* if we can.
	return
}

// TODO: In case there is +2/3 majority in Prevotes set for some
// block and cs.ProposalBlock contains different block, either
// proposer is faulty or voting power of faulty processes is more
// than 1/3. We should trigger in the future accountability
// procedure at this point.

// Do not count prevote/precommit/commit into handleBlockPartMsg's span

// Move onto the next step

// this is optimisation as this will be triggered when prevote is added

// If we're waiting on the proposal block...

// Attempt to add the vote. if its a duplicate signature, dupeout the validator
func (cs *State) tryAddVote(ctx context.Context, vote *types.Vote, peerID types.NodeID, handleVoteMsgSpan otrace.Span) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If the vote height is off, we'll just ignore it,
// But if it's a conflicting sig, add it to the cs.evpool.
// If it's otherwise invalid, punish peer.
//nolint: gocritic

// report conflicting votes to the evidence pool

// Either
// 1) bad peer OR
// 2) not a bad peer? this can also err sometimes with "Unexpected step" OR
// 3) tmkms use with multiple validators connecting to a single tmkms instance
//		(https://github.com/tendermint/tendermint/issues/3839).

func (cs *State) addVote(
	ctx context.Context,
	vote *types.Vote,
	peerID types.NodeID,
	handleVoteMsgSpan otrace.Span,
) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// A precommit for the previous height?
// These come in while we wait timeoutCommit

// Late precommit at prior height is ignored

// if we can skip timeoutCommit and have all the votes now,

// go straight to new round (skip timeout commit)
// cs.scheduleTimeout(time.Duration(0), cs.Height, 0, cstypes.RoundStepNewHeight)

// Height mismatch is ignored.
// Not necessarily a bad peer, but not favorable behavior.

// Either duplicate, or error upon cs.Votes.AddByIndex()

// Check to see if >2/3 of the voting power on the network voted for any non-nil block.

// Greater than 2/3 of the voting power on the network voted for some
// non-nil block

// Update Valid* if we can.

// we're getting the wrong block

// If +2/3 prevotes for *anything* for future round:

// Round-skip if there is any 2/3+ of votes ahead of us

// current round

// If the proposal is now complete, enter prevote of cs.Round.

// Executed as TwoThirdsMajority could be from a higher round

// CONTRACT: cs.privValidator is not nil.
func (cs *State) signVote(
	ctx context.Context,
	privValidator types.PrivValidator,
	msgType tmproto.SignedMsgType,
	hash []byte,
	header types.PartSetHeader,
) (*types.Vote, error) {
	_ = "STUB: not implemented"
	// Flush the WAL. Otherwise, we may not recompute the same vote to sign,
	// and the privValidator will refuse to sign anything.
	return nil, nil
}

// If the signedMessageType is for precommit,
// use our local precommit Timeout as the max wait time for getting a singed commit. The same goes for prevote.

// sign the vote and publish on internalMsgQueue
func (cs *State) signAddVote(
	ctx context.Context,
	msgType tmproto.SignedMsgType,
	hash []byte,
	header types.PartSetHeader,
) *types.Vote {
	_ = "STUB: not implemented"
	return nil
}

// the node does not have a key

// Vote won't be signed, but it's not critical.

// If the node not in the validator set, do nothing.

// TODO: pass pubKey to signVote

// updatePrivValidatorPubKey get's the private validator public key and
// memoizes it. This func returns an error if the private validator is not
// responding or responds with an error.
func (cs *State) updatePrivValidatorPubKey(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// set context timeout depending on the configuration and the State step,
// this helps in avoiding blocking of the remote signer connection.

// look back to check existence of the node's consensus votes before joining consensus
func (cs *State) checkDoubleSigningRisk(height int64) error { _ = "STUB: not implemented"; return nil }

func (cs *State) calculatePrevoteMessageDelayMetrics() { _ = "STUB: not implemented"; return }

//---------------------------------------------------------

func (cs *State) proposeTimeout(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (cs *State) voteTimeout(round int32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (cs *State) commitTime(t time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (cs *State) bypassCommitTimeout() bool { _ = "STUB: not implemented"; return false }

func (cs *State) calculateProposalTimestampDifferenceMetric() { _ = "STUB: not implemented"; return }

// proposerWaitTime determines how long the proposer should wait to propose its next block.
// If the result is zero, a block can be proposed immediately.
//
// Block times must be monotonically increasing, so if the block time of the previous
// block is larger than the proposer's current time, then the proposer will sleep
// until its local clock exceeds the previous block time.
func proposerWaitTime(lt tmtime.Source, bt time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
