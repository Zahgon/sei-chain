package pex

import (
	"context"
	"errors"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
	"golang.org/x/time/rate"
)

var (
	logger = seilog.NewLogger("tendermint", "internal", "p2p", "pex")

	_ service.Service = (*Reactor)(nil)
)

var maxPeerRecvRate = rate.Every(100 * time.Millisecond)

const (
	// PexChannel is a channel for PEX messages
	PexChannel = 0x00

	// over-estimate of max NetAddress size
	// hexID (40) + IP (16) + Port (2) + Name (100) ...
	// NOTE: dont use massive DNS name ..
	maxAddressSize = 256

	// NOTE: amplification factor!
	// small request results in up to maxMsgSize response
	maxMsgSize = 1000 + maxAddressSize*p2p.MaxPexAddrs

	// the minimum time one peer can send another request to the same peer
	maxPeerRecvBurst    = 10
	DefaultSendInterval = 10 * time.Second
)

var ErrNoPeersAvailable = errors.New("no available peers to send a PEX request to (retrying)")

// TODO: We should decide whether we want channel descriptors to be housed
// within each reactor (as they are now) or, considering that the reactor doesn't
// really need to care about the channel descriptors, if they should be housed
// in the node module.
func ChannelDescriptor() p2p.ChannelDescriptor[*pb.PexMessage] {
	_ = "STUB: not implemented"
	return nil
}

// The peer exchange or PEX reactor supports the peer manager by sending
// requests to other peers for addresses that can be given to the peer manager
// and at the same time advertises addresses to peers that need more.
type Reactor struct {
	service.BaseService
	sendInterval time.Duration
	router       *p2p.Router
	// peerLimiters limits the number of messages received from peers.
	peerLimiters utils.Mutex[map[types.NodeID]*rate.Limiter]
	channel      *p2p.Channel[*pb.PexMessage]
}

// NewReactor returns a reference to a new reactor.
func NewReactor(
	router *p2p.Router,
	sendInterval time.Duration,
) (*Reactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reactor) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnStart implements service.Service.
func (r *Reactor) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnStop implements service.Service.
func (r *Reactor) OnStop() { _ = "STUB: not implemented"; return }

func wrap[T *pb.PexRequest | *pb.PexResponse](msg T) *pb.PexMessage {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) sendRoutine(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processPexCh implements a blocking event loop where we listen for p2p
// Envelope messages from the pexCh.
func (r *Reactor) processPexCh(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processPeerUpdates processes peer updates.
func (r *Reactor) processPeerUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(gprusak): peer updates and pexChannel are not synchronized, therefore we
// might have race conditions where we receive messages from peers which are already marked as down.

// handlePexMessage handles envelopes sent from peers on the PexChannel.
// If an update was received, a new polling interval is returned; otherwise the
// duration is 0.
func (r *Reactor) handlePexMessage(m p2p.RecvMsg[*pb.PexMessage]) error {
	_ = "STUB: not implemented"
	return nil
}

// Fetch peers from the peer manager, convert NodeAddresses into URL
// strings, and send them back to the caller.

// Verify that the response does not exceed the safety limit.
