package reactor

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var (
	logger = seilog.NewLogger("tendermint", "internal", "mempool")

	_ service.Service = (*Reactor)(nil)
)

const MempoolChannel p2p.ChannelID = 0x30

// Reactor implements a service that contains mempool of txs that are broadcasted
// amongst peers. It maintains a map from peer ID to counter, to prevent gossiping
// txs to the peers you received it from.
type Reactor struct {
	service.BaseService

	cfg     *config.MempoolConfig
	mempool *mempool.TxMempool
	ids     *IDs

	router *p2p.Router

	failedCheckTxCounts utils.Mutex[map[types.NodeID]int]

	channel      *p2p.Channel[*pb.Message]
	readyToStart chan struct{}
}

// NewReactor returns a reference to a new reactor.
func NewReactor(cfg *config.MempoolConfig, txmp *mempool.TxMempool, router *p2p.Router) (*Reactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reactor) MarkReadyToStart() { _ = "STUB: not implemented"; return }

// GetChannelDescriptor produces an instance of a descriptor for this package's
// required channels.
func GetChannelDescriptor(cfg *config.MempoolConfig) p2p.ChannelDescriptor[*pb.Message] {
	_ = "STUB: not implemented"
	return nil
}

// OnStart starts separate goroutines for each p2p channel and listens for
// envelopes on each. In addition, it also listens for peer updates and handles
// messages on that p2p channel accordingly. The caller must be sure to execute
// OnStop to ensure the outbound p2p channels are closed.
func (r *Reactor) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnStop stops the reactor by signaling to all spawned goroutines to exit and
// blocking until they all exit.
func (r *Reactor) OnStop() {
	_ = "STUB: not implemented"

	// handleMempoolMessage handles envelopes sent from peers on the MempoolChannel.
	// For every tx in the message, we execute CheckTx. It returns an error if an
	// empty set of txs are sent in an envelope or if we receive an unexpected
	// message type.
	return
}

func (r *Reactor) handleMempoolMessage(ctx context.Context, m p2p.RecvMsg[*pb.Message]) error {
	_ = "STUB: not implemented"
	return nil
}

// If the tx is in the cache, then we've been gossiped a tx
// that we've already got. Gossip should be smarter, but it's
// not a problem.

// Do not propagate context cancellation errors, but do not
// continue to check transactions from this message if we are
// shutting down.

func (r *Reactor) accountFailedCheckTx(nodeID types.NodeID, err error) {
	_ = "STUB: not implemented"
	return
}

// handleMessage handles an envelope sent from a peer on a specific p2p channel.
// It will handle errors and any possible panics gracefully. A caller can handle
// any error returned by sending a PeerError on the respective channel.
func (r *Reactor) handleMessage(ctx context.Context, m p2p.RecvMsg[*pb.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// processMempoolCh implements a blocking event loop where we listen for p2p
// envelope messages from the mempool channel.
func (r *Reactor) processMempoolCh(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// processPeerUpdates initiates a blocking process where we listen for and
// handle PeerUpdate messages. When the reactor is stopped, we will catch the
// signal and close the p2p PeerUpdatesCh gracefully.
func (r *Reactor) processPeerUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// We keep peer management even when broadcasting is disabled,
// so that failedCheckTxCounts WAI.

func (r *Reactor) broadcastTxRoutine(ctx context.Context, peerID types.NodeID) {
	_ = "STUB: not implemented"
	return
}
