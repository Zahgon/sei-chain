package consensus

import (
	"context"
	"sync/atomic"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	cstypes "github.com/sei-protocol/sei-chain/sei-tendermint/internal/consensus/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	tmcons "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/consensus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var (
	_ service.Service = (*Reactor)(nil)
)

type desc = p2p.ChannelDescriptor[*tmcons.Message]

func GetStateChannelDescriptor() desc { _ = "STUB: not implemented"; return *new(desc) }

func GetDataChannelDescriptor() desc {
	_ = "STUB: not implemented"

	// TODO: Consider a split between gossiping current block and catchup
	// stuff. Once we gossip the whole block there is nothing left to send
	// until next height or round.
	return *new(desc)
}

func GetVoteChannelDescriptor() desc { _ = "STUB: not implemented"; return *new(desc) }

func GetVoteSetChannelDescriptor() desc { _ = "STUB: not implemented"; return *new(desc) }

const (
	StateChannel       = p2p.ChannelID(0x20)
	DataChannel        = p2p.ChannelID(0x21)
	VoteChannel        = p2p.ChannelID(0x22)
	VoteSetBitsChannel = p2p.ChannelID(0x23)

	maxMsgSize = 4194304 // 4MB; NOTE: keep larger than types.PartSet sizes.
)

// Reactor defines a reactor for the consensus service.
type Reactor struct {
	service.BaseService
	cfg *config.Config

	state    *State
	router   *p2p.Router
	channels channelBundle
	eventBus *eventbus.EventBus
	Metrics  *Metrics

	peers       utils.RWMutex[map[types.NodeID]*PeerState]
	roundState  atomic.Pointer[cstypes.RoundState]
	readySignal utils.AtomicSend[bool]
}

// NewReactor returns a reference to a new consensus reactor, which implements
// the service.Service interface. It accepts a logger, consensus state, references
// to relevant p2p Channels and a channel to listen for peer updates on. The
// reactor will close all p2p Channels when stopping.
func NewReactor(
	cs *State,
	router *p2p.Router,
	eventBus *eventbus.EventBus,
	waitSync bool,
	metrics *Metrics,
	cfg *config.Config,
) (*Reactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type channelBundle struct {
	state  *p2p.Channel[*tmcons.Message]
	data   *p2p.Channel[*tmcons.Message]
	vote   *p2p.Channel[*tmcons.Message]
	votSet *p2p.Channel[*tmcons.Message]
}

// OnStart starts separate go routines for each p2p Channel and listens for
// envelopes on each. In addition, it also listens for peer updates and handles
// messages on that p2p channel accordingly. The caller must be sure to execute
// OnStop to ensure the outbound p2p Channels are closed.
func (r *Reactor) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// All of the channel processing routines are noops until readySignal,
// but they are spawned immediately so that they keep the channels empty until then.

// OnStop stops the reactor by signaling to all spawned goroutines to exit and
// blocking until they all exit, as well as unsubscribing from events and stopping
// state.
func (r *Reactor) OnStop() {
	_ = "STUB: not implemented"

	// WaitSync returns whether the consensus reactor is waiting for state/block sync.
	return
}

func (r *Reactor) WaitSync() bool { _ = "STUB: not implemented"; return false }

// SwitchToConsensus switches from block-sync mode to consensus mode. It resets
// the state, turns off block-sync, and starts the consensus state-machine.
func (r *Reactor) SwitchToConsensus(state sm.State, skipWAL bool) {
	_ = "STUB: not implemented"
	return
}

// we have no votes, so reconstruct LastCommit from SeenCommit

// String returns a string representation of the Reactor.
func (r *Reactor) String() string { _ = "STUB: not implemented"; return "" }

// GetPeerState returns PeerState for a given NodeID.
func (r *Reactor) GetPeerState(peerID types.NodeID) (*PeerState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *Reactor) broadcastNewRoundStepMessage(rs *cstypes.RoundState) {
	_ = "STUB: not implemented"
	return
}

// Broadcasts NewValidBlockMessage whenever new valid block is reported.
// It rebroadcasts the NewValidBlockMessage periodically to ensure that peers know which parts
// we are missing. It is critical in case we have small number of peers (for example just 1),
// and they are overloaded (they drop messages a lot).
func (r *Reactor) broadcastValidBlockRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Rebroadcasting is a fallback mechanism, no need to expose the frequency as
	// a config parameter.
	return nil
}

// Block parts bit array might be updated between iterations,
// so we need to reconstruct the message each time.

func (r *Reactor) broadcastHasVoteMessage(vote *types.Vote) { _ = "STUB: not implemented"; return }

func makeRoundStepMessage(rs *cstypes.RoundState) *NewRoundStepMessage {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) sendNewRoundStepMessage(peerID types.NodeID) { _ = "STUB: not implemented"; return }

func (r *Reactor) updateRoundStateRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) gossipDataForCatchup(rs *cstypes.RoundState, prs *cstypes.PeerRoundState, ps *PeerState) {
	_ = "STUB: not implemented"
	return
}

// ensure that the peer's PartSetHeader is correct

// not our height, so it does not matter.
// not our height, so it does not matter

func (r *Reactor) gossipDataRoutine(ctx context.Context, ps *PeerState) error {
	_ = "STUB: not implemented"
	return nil
}

// Send proposal Block parts?

// this tells peer that this part applies to us
// this tells peer that this part applies to us

// if the peer is on a previous height that we have, help catch up

// If we never received the commit message from the peer, the block parts
// will not be initialized.

// Continue the loop since prs is a copy and not effected by this
// initialization.

// if height and round don't match, sleep

// By here, height and round match.
// Proposal block parts were already matched and sent if any were wanted.
// (These can match on hash so the round doesn't matter)
// Now consider sending other things, like the Proposal itself.

// Send Proposal && ProposalPOL BitArray?

// Proposal: share the proposal metadata with peer.

// NOTE: A peer might have received a different proposal message, so
// this Proposal msg will be rejected!

// ProposalPOL: lets peer know which POL votes we have so far. The peer
// must receive ProposalMessage first. Note, rs.Proposal was validated,
// so rs.Proposal.POLRound <= rs.Round, so we definitely have
// rs.Votes.Prevotes(rs.Proposal.POLRound).

// pickSendVote picks a vote and sends it to the peer. It will return true if
// there is a vote to send and false otherwise.
func (r *Reactor) pickSendVote(ps *PeerState, votes types.VoteSetReader) bool {
	_ = "STUB: not implemented"
	return false
}

// expensive, so we only want to call if debug is on

// SetHasVote can fail because ps state (in particular the bitarrays)
// is not verified and it depends what peer has sent us.

func (r *Reactor) gossipVotesForHeight(
	rs *cstypes.RoundState,
	prs *cstypes.PeerRoundState,
	ps *PeerState,
) bool {
	_ = "STUB: not implemented"
	return false
}

// if there are lastCommits to send...

// if there are POL prevotes to send...

// if there are prevotes to send...

// if there are precommits to send...

// if there are prevotes to send...(which are needed because of validBlock mechanism)

// if there are POLPrevotes to send...

func (r *Reactor) gossipVotesRoutine(ctx context.Context, ps *PeerState) error {
	_ = "STUB: not implemented"
	return nil
}

// if height matches, then send LastCommit, Prevotes, and Precommits

// special catchup logic -- if peer is lagging by height 1, send LastCommit

// catchup logic -- if peer is lagging by more than 1, send Commit

// Load the block's extended commit for prs.Height, which contains precommit
// signatures for prs.Height.

// NOTE: `queryMaj23Routine` has a simple crude design since it only comes
// into play for liveness when there's a signature DDoS attack happening.
func (r *Reactor) queryMaj23Routine(ctx context.Context, ps *PeerState) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO create more reliable copies of these
// structures so the following go routines don't race

// maybe send Height/Round/Prevotes

// maybe send Height/Round/ProposalPOL

// maybe send Height/Round/Precommits

// Little point sending LastCommitRound/LastCommit, these are fleeting and
// non-blocking.

// maybe send Height/CatchupCommitRound/CatchupCommit

// handleStateMessage handles envelopes sent from peers on the StateChannel.
// An error is returned if the message is unrecognized or if validation fails.
// If we fail to find the peer state for the envelope sender, we perform a no-op
// and return. This can happen when we process the envelope after the peer is
// removed.
func (r *Reactor) handleStateMessage(m p2p.RecvMsg[*tmcons.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// peer claims to have a maj23 for some BlockID at <H,R,S>

// Respond with a VoteSetBitsMessage showing which votes we have and
// consequently shows which we don't have.

// handleDataMessage handles envelopes sent from peers on the DataChannel. If we
// fail to find the peer state for the envelope sender, we perform a no-op and
// return. This can happen when we process the envelope after the peer is
// removed.
func (r *Reactor) handleDataMessage(ctx context.Context, m p2p.RecvMsg[*tmcons.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// handleVoteMessage handles envelopes sent from peers on the VoteChannel. If we
// fail to find the peer state for the envelope sender, we perform a no-op and
// return. This can happen when we process the envelope after the peer is
// removed.
func (r *Reactor) handleVoteMessage(ctx context.Context, m p2p.RecvMsg[*tmcons.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// handleVoteSetBitsMessage handles envelopes sent from peers on the
// VoteSetBitsChannel. If we fail to find the peer state for the envelope sender,
// we perform a no-op and return. This can happen when we process the envelope
// after the peer is removed.
func (r *Reactor) handleVoteSetBitsMessage(m p2p.RecvMsg[*tmcons.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) recoverToErr(err *error) { _ = "STUB: not implemented"; return }

// processStateCh initiates a blocking process where we listen for and handle
// envelopes on the StateChannel. Any error encountered during message
// execution will result in a PeerError being sent on the StateChannel. When
// the reactor is stopped, we will catch the signal and close the p2p Channel
// gracefully.
func (r *Reactor) processStateCh(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processDataCh initiates a blocking process where we listen for and handle
// envelopes on the DataChannel. Any error encountered during message
// execution will result in a PeerError being sent on the DataChannel. When
// the reactor is stopped, we will catch the signal and close the p2p Channel
// gracefully.
func (r *Reactor) processDataCh(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processVoteCh initiates a blocking process where we listen for and handle
// envelopes on the VoteChannel. Any error encountered during message
// execution will result in a PeerError being sent on the VoteChannel. When
// the reactor is stopped, we will catch the signal and close the p2p Channel
// gracefully.
func (r *Reactor) processVoteCh(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processVoteCh initiates a blocking process where we listen for and handle
// envelopes on the VoteSetBitsChannel. Any error encountered during message
// execution will result in a PeerError being sent on the VoteSetBitsChannel.
// When the reactor is stopped, we will catch the signal and close the p2p
// Channel gracefully.
func (r *Reactor) processVoteSetBitsCh(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// processPeerUpdates initiates a blocking process where we listen for and handle
// PeerUpdate messages. When the reactor is stopped, we will catch the signal and
// close the p2p PeerUpdatesCh gracefully.
func (r *Reactor) processPeerUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Only ctx.Canceled is expected and only once peerCtx is done.

func (r *Reactor) runPeer(ctx context.Context, ps *PeerState) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) recordPeerMsg(msg msgInfo) { _ = "STUB: not implemented"; return }

func (r *Reactor) SetStateSyncingMetrics(v float64) { _ = "STUB: not implemented"; return }

func (r *Reactor) SetBlockSyncingMetrics(v float64) { _ = "STUB: not implemented"; return }
