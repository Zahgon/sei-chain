package statesync

import (
	"context"
	"errors"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/statesync"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var (
	ErrNoConnectedPeers    = errors.New("no available peers to dispatch request to")
	ErrUnsolicitedResponse = errors.New("unsolicited light block response")
	ErrPeerAlreadyBusy     = errors.New("peer is already processing a request")
	ErrDisconnected        = errors.New("dispatcher disconnected")
)

// A Dispatcher multiplexes concurrent requests by multiple peers for light blocks.
// Only one request per peer can be sent at a time. Subsequent concurrent requests will
// report an error from the LightBlock method.
// NOTE: It is not the responsibility of the dispatcher to verify the light blocks.
type Dispatcher struct {
	// the channel with which to send light block requests on
	requestCh *p2p.Channel[*pb.Message]

	// all pending calls that have been dispatched and are awaiting an answer
	calls utils.Mutex[map[types.NodeID]chan *types.LightBlock]
}

func NewDispatcher(requestChannel *p2p.Channel[*pb.Message]) *Dispatcher {
	_ = "STUB: not implemented"
	return nil
}

// LightBlock uses the request channel to fetch a light block from a given peer
// tracking, the call and waiting for the reactor to pass back the response. A nil
// LightBlock response is used to signal that the peer doesn't have the requested LightBlock.
func (d *Dispatcher) LightBlock(ctx context.Context, height int64, peer types.NodeID) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	// dispatch the request to the peer
	return nil, nil
}

// clean up the call after a response is returned

// wait for a response, cancel or timeout

// dispatch takes a peer and allocates it a channel so long as it's not already
// busy and the receiving channel is still running. It then dispatches the message
func (d *Dispatcher) dispatch(ctx context.Context, peer types.NodeID, height int64) (chan *types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if a request for the same peer has already been made

// send request
//nolint:gosec // height is a validated positive block height

// Respond allows the underlying process which receives requests on the
// requestCh to respond with the respective light block. A nil response is used to
// represent that the receiver of the request does not have a light block at that height.
func (d *Dispatcher) Respond(ctx context.Context, lb *tmproto.LightBlock, peer types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// check that the response came from a request

// this can also happen if the response came in after the timeout

// If lb is nil we take that to mean that the peer didn't have the requested light
// block and thus pass on the nil to the caller.

// Close shuts down the dispatcher and cancels any pending calls awaiting responses.
// Peers awaiting responses that have not arrived are delivered a nil block.
func (d *Dispatcher) Close() { _ = "STUB: not implemented"; return }

// don't close the channel here as it's closed in
// other handlers, and would otherwise get garbage
// collected.

//----------------------------------------------------------------

// BlockProvider is a p2p based light provider which uses a dispatcher connected
// to the state sync reactor to serve light blocks to the light client
//
// TODO: This should probably be moved over to the light package but as we're
// not yet officially supporting p2p light clients we'll leave this here for now.
//
// NOTE: BlockProvider will return an error with concurrent calls. However, we don't
// need a mutex because a light client (and the backfill process) will never call a
// method more than once at the same time
type BlockProvider struct {
	peer       types.NodeID
	chainID    string
	dispatcher *Dispatcher
}

// Creates a block provider which implements the light client Provider interface.
func NewBlockProvider(peer types.NodeID, chainID string, dispatcher *Dispatcher) *BlockProvider {
	_ = "STUB: not implemented"
	return nil
}

// LightBlock fetches a light block from the peer at a specified height returning either a
// light block or an appropriate error.
func (p *BlockProvider) LightBlock(ctx context.Context, height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check that the height requested is the same one returned

// perform basic validation

// ReportEvidence should allow for the light client to report any light client
// attacks. This is a no op as there currently isn't a way to wire this up to
// the evidence reactor (we should endeavor to do this in the future but for now
// it's not critical for backwards verification)
func (p *BlockProvider) ReportEvidence(ctx context.Context, ev types.Evidence) error {
	_ = "STUB: not implemented"

	// String implements stringer interface
	return nil
}

func (p *BlockProvider) String() string { _ = "STUB: not implemented"; return "" }

// Returns the ID address of the provider (NodeID of peer)
func (p *BlockProvider) ID() string { _ = "STUB: not implemented"; return "" }

//----------------------------------------------------------------

type peerListInner struct {
	peers []types.NodeID
}

// peerList is a rolling list of peers. This is used to distribute the load of
// retrieving blocks over all the peers the reactor is connected to
type PeerList struct {
	inner utils.Watch[*peerListInner]
}

func NewPeerList() *PeerList { _ = "STUB: not implemented"; return nil }

func (l *PeerList) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *PeerList) Pop(ctx context.Context) types.NodeID {
	_ = "STUB: not implemented"
	return *new(types.NodeID)
}

func (l *PeerList) Append(peer types.NodeID) { _ = "STUB: not implemented"; return }

func (l *PeerList) Remove(peer types.NodeID) { _ = "STUB: not implemented"; return }

func (l *PeerList) All() []types.NodeID { _ = "STUB: not implemented"; return nil }

func (l *PeerList) WaitUntilContains(ctx context.Context, id types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}
