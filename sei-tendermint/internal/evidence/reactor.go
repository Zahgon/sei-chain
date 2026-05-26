package evidence

import (
	"context"
	"sync"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var _ service.Service = (*Reactor)(nil)

const (
	EvidenceChannel = p2p.ChannelID(0x38)

	maxMsgSize = 4194304 // 4MB TODO make it configurable

	// broadcast all uncommitted evidence this often. This sets when the reactor
	// goes back to the start of the list and begins sending the evidence again.
	// Most evidence should be committed in the very next block that is why we wait
	// just over the block production rate before sending evidence again.
	broadcastEvidenceInterval = 60 * time.Second
)

// GetChannelDescriptor produces an instance of a descriptor for this
// package's required channels.
func GetChannelDescriptor() p2p.ChannelDescriptor[*pb.Evidence] {
	_ = "STUB: not implemented"
	return nil
}

// Reactor handles evpool evidence broadcasting amongst peers.
type Reactor struct {
	service.BaseService

	evpool *Pool
	router *p2p.Router

	mtx sync.Mutex

	peerRoutines map[types.NodeID]context.CancelFunc
	channel      *p2p.Channel[*pb.Evidence]
}

// NewReactor returns a reference to a new evidence reactor, which implements the
// service.Service interface. It accepts a p2p Channel dedicated for handling
// envelopes with EvidenceList messages.
func NewReactor(
	router *p2p.Router,
	evpool *Pool,
) (*Reactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OnStart starts separate go routines for each p2p Channel and listens for
// envelopes on each. In addition, it also listens for peer updates and handles
// messages on that p2p channel accordingly. The caller must be sure to execute
// OnStop to ensure the outbound p2p Channels are closed. No error is returned.
func (r *Reactor) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnStop stops the reactor by signaling to all spawned goroutines to exit and
// blocking until they all exit.
func (r *Reactor) OnStop() { _ = "STUB: not implemented"; return }

// handleEvidenceMessage handles envelopes sent from peers on the EvidenceChannel.
// It returns an error only if the Envelope.Message is unknown for this channel
// or if the given evidence is invalid. This should never be called outside of
// handleMessage.
func (r *Reactor) handleEvidenceMessage(ctx context.Context, m p2p.RecvMsg[*pb.Evidence]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Process the evidence received from a peer
// Evidence is sent and received one by one

// If we're given invalid evidence by the peer, notify the router that
// we should remove this peer by returning an error.

// processEvidenceCh implements a blocking event loop where we listen for p2p
// Envelope messages from the evidenceCh.
func (r *Reactor) processEvidenceCh(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// processPeerUpdate processes a PeerUpdate. For new or live peers it will check
// if an evidence broadcasting goroutine needs to be started. For down or
// removed peers, it will check if an evidence broadcasting goroutine
// exists and signal that it should exit.
//
// FIXME: The peer may be behind in which case it would simply ignore the
// evidence and treat it as invalid. This would cause the peer to disconnect.
// The peer may also receive the same piece of evidence multiple times if it
// connects/disconnects frequently from the broadcasting peer(s).
//
// REF: https://github.com/tendermint/tendermint/issues/4727
func (r *Reactor) processPeerUpdate(ctx context.Context, peerUpdate p2p.PeerUpdate) {
	_ = "STUB: not implemented"
	return
}

// Do not allow starting new evidence broadcast loops after reactor shutdown
// has been initiated. This can happen after we've manually closed all
// peer broadcast loops, but the router still sends in-flight peer updates.

// Check if we've already started a goroutine for this peer, if not we create
// a new done channel so we can explicitly close the goroutine if the peer
// is later removed, we increment the waitgroup so the reactor can stop
// safely, and finally start the goroutine to broadcast evidence to that peer.

// Check if we've started an evidence broadcasting goroutine for this peer.
// If we have, we signal to terminate the goroutine via the channel's closure.
// This will internally decrement the peer waitgroup and remove the peer
// from the map of peer evidence broadcasting goroutines.

// processPeerUpdates initiates a blocking process where we listen for and handle
// PeerUpdate messages. When the reactor is stopped, we will catch the signal and
// close the p2p PeerUpdatesCh gracefully.
func (r *Reactor) processPeerUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// broadcastEvidenceLoop starts a blocking process that continuously reads pieces
// of evidence off of a linked-list and sends the evidence in a p2p Envelope to
// the given peer by ID. This should be invoked in a goroutine per unique peer
// ID via an appropriate PeerUpdate. The goroutine can be signaled to gracefully
// exit by either explicitly closing the provided doneCh or by the reactor
// signaling to stop.
//
// TODO: This should be refactored so that we do not blindly gossip evidence
// that the peer has already received or may not be ready for.
//
// REF: https://github.com/tendermint/tendermint/issues/4727
func (r *Reactor) broadcastEvidenceLoop(ctx context.Context, peerID types.NodeID, evidenceCh *p2p.Channel[*pb.Evidence]) {
	_ = "STUB: not implemented"
	return
}

// This happens because the CElement we were looking at got garbage
// collected (removed). That is, NextWait() returned nil. So we can go
// ahead and start from the beginning.

// This should never happen, but for some historical reasons there is a recover() deferred
// in this function. Current behavior is that evidence will stop being broadcasted to a peer
// if this happens, but the node with not crash.
// TODO(gprusak): reevaluate if this broadcastEvidenceLoop may panic even if there is no bug.
//   If it cannot, remove recover().

// Send the evidence to the corresponding peer. Note, the peer may be behind
// and thus would not be able to process the evidence correctly. Also, the
// peer may receive this piece of evidence multiple times if it added and
// removed frequently from the broadcasting peer.

// In case the observed element has been removed or we waited too long for the next element,
// we start sending from the beginning.
