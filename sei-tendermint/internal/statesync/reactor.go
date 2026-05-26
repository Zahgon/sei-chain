package statesync

import (
	"context"
	"sync"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/store"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/statesync"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var (
	_ service.Service = (*Reactor)(nil)
)

type isPBMessage interface {
	*pb.SnapshotsRequest |
		*pb.SnapshotsResponse |
		*pb.ChunkRequest |
		*pb.ChunkResponse |
		*pb.LightBlockRequest |
		*pb.LightBlockResponse |
		*pb.ParamsRequest |
		*pb.ParamsResponse
}

func wrap[T isPBMessage](msg T) *pb.Message { _ = "STUB: not implemented"; return nil }

const (
	// SnapshotChannel exchanges snapshot metadata
	SnapshotChannel = p2p.ChannelID(0x60)

	// ChunkChannel exchanges chunk contents
	ChunkChannel = p2p.ChannelID(0x61)

	// LightBlockChannel exchanges light blocks
	LightBlockChannel = p2p.ChannelID(0x62)

	// ParamsChannel exchanges consensus params
	ParamsChannel = p2p.ChannelID(0x63)

	// recentSnapshots is the number of recent snapshots to send and receive per peer.
	recentSnapshots = 10

	// snapshotMsgSize is the maximum size of a snapshotResponseMessage
	snapshotMsgSize = int(4e6) // ~4MB

	// chunkMsgSize is the maximum size of a chunkResponseMessage
	chunkMsgSize = int(16e6) // ~16MB

	// lightBlockMsgSize is the maximum size of a lightBlockResponseMessage
	lightBlockMsgSize = int(1e7) // ~1MB

	// paramMsgSize is the maximum size of a paramsResponseMessage
	paramMsgSize = int(1e5) // ~100kb

	// maxLightBlockRequestRetries is the amount of retries acceptable before
	// the backfill process aborts
	maxLightBlockRequestRetries = 40
)

func GetSnapshotChannelDescriptor() p2p.ChannelDescriptor[*pb.Message] {
	_ = "STUB: not implemented"
	return nil
}

func GetChunkChannelDescriptor() p2p.ChannelDescriptor[*pb.Message] {
	_ = "STUB: not implemented"
	return nil
}

func GetLightBlockChannelDescriptor() p2p.ChannelDescriptor[*pb.Message] {
	_ = "STUB: not implemented"
	return nil
}

func GetParamsChannelDescriptor() p2p.ChannelDescriptor[*pb.Message] {
	_ = "STUB: not implemented"
	return nil
}

// Metricer defines an interface used for the rpc sync info query, please see statesync.metrics
// for the details.
type Metricer interface {
	TotalSnapshots() int64
	ChunkProcessAvgTime() time.Duration
	SnapshotHeight() int64
	SnapshotChunksCount() int64
	SnapshotChunksTotal() int64
	BackFilledBlocks() int64
	BackFillBlocksTotal() int64
}

// Reactor handles state sync, both restoring snapshots for the local node and
// serving snapshots for other nodes.
type Reactor struct {
	service.BaseService

	chainID       string
	initialHeight int64
	cfg           config.StateSyncConfig
	stateStore    sm.Store
	blockStore    *store.BlockStore

	conn         *proxy.Proxy
	tempDir      string
	router       *p2p.Router
	evict        func(types.NodeID, error)
	postSyncHook func(context.Context, sm.State) error

	// when true, the reactor will, during startup perform a
	// statesync for this node, and otherwise just provide
	// snapshots to other nodes.
	needsStateSync bool

	// Dispatcher is used to multiplex light block requests and responses over multiple
	// peers used by the p2p state provider and in reverse sync.
	dispatcher *Dispatcher
	peers      *PeerList

	// These will only be set when a state sync is in progress. It is used to feed
	// received snapshots and chunks into the syncer and manage incoming and outgoing
	// providers.
	mtx            sync.RWMutex
	initSyncer     func() *syncer
	requestSnaphot func() error
	syncer         *syncer
	providers      map[types.NodeID]*BlockProvider
	stateProvider  StateProvider

	eventBus           *eventbus.EventBus
	metrics            *Metrics
	backfillBlockTotal int64
	backfilledBlocks   int64

	// For some reason channels below used to be processed synchronously.
	// Now each of these has their own processing loop, but to simulate the previous
	// behavior we use a mutex to ensure only one message is processed at a time across all channels.
	// TODO(gprusak): verify that the message handlers can be executed concurrenty and remove this mutex.
	processChGuard    sync.Mutex
	snapshotChannel   *p2p.Channel[*pb.Message]
	chunkChannel      *p2p.Channel[*pb.Message]
	lightBlockChannel *p2p.Channel[*pb.Message]
	paramsChannel     *p2p.Channel[*pb.Message]

	// keep track of the last time we saw no available peers, so we can restart if it's been too long
	lastNoAvailablePeers time.Time

	// Used to signal a restart the node on the application level
	restartEvent                  func()
	restartNoAvailablePeersWindow time.Duration
}

// NewReactor returns a reference to a new state sync reactor, which implements
// the service.Service interface. It accepts a logger, connections for snapshots
// and querying, a router used to open the required p2p channels, and a channel
// to listen for peer updates on. Note, the reactor will close all p2p Channels
// when stopping.
func NewReactor(
	chainID string,
	initialHeight int64,
	cfg config.StateSyncConfig,
	conn *proxy.Proxy,
	router *p2p.Router,
	stateStore sm.Store,
	blockStore *store.BlockStore,
	tempDir string,
	ssMetrics *Metrics,
	eventBus *eventbus.EventBus,
	postSyncHook func(context.Context, sm.State) error,
	needsStateSync bool,
	restartEvent func(),
	selfRemediationConfig *config.SelfRemediationConfig,
) (*Reactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // validated in config.ValidateBasic against MaxInt64

func (r *Reactor) initStateProvider(ctx context.Context, chainID string, initialHeight int64) error {
	_ = "STUB: not implemented"
	return nil
}

// OnStart starts separate go routines for each p2p Channel and listens for
// ms on each. In addition, it also listens for peer updates and handles
// messages on that p2p channel accordingly. Note, we do not launch a go-routine to
// handle individual ms as to not have to deal with bounding workers or pools.
// The caller must be sure to execute OnStop to ensure the outbound p2p Channels are
// closed. No error is returned.
func (r *Reactor) OnStart(ctx context.Context) error {
	_ = "STUB: not implemented"
	// define constructor and helper functions, that hold
	// references to these channels for use later. This is not
	// ideal.
	return nil
}

// request snapshots from all currently connected peers

// OnStop stops the reactor by signaling to all spawned goroutines to exit and
// blocking until they all exit.
func (r *Reactor) OnStop() {
	_ = "STUB: not implemented"
	// tell the dispatcher to stop sending any more requests
	return
}

// Sync runs a state sync, fetching snapshots and providing chunks to the
// application. At the close of the operation, Sync will bootstrap the state
// store and persist the commit at that height so that either consensus or
// blocksync can commence. It will then proceed to backfill the necessary amount
// of historical blocks before participating in consensus
func (r *Reactor) Sync(ctx context.Context) (sm.State, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil
}

// We need at least two peers (for cross-referencing of light blocks) before we can
// begin state sync

// reset syncing objects at the close of Sync

// Backfill sequentially fetches, verifies and stores light blocks in reverse
// order. It does not stop verifying blocks until reaching a block with a height
// and time that is less or equal to the stopHeight and stopTime. The
// trustedBlockID should be of the header at startHeight.
func (r *Reactor) Backfill(ctx context.Context, state sm.State) error {
	_ = "STUB: not implemented"
	return nil
}

// ensure that stop height doesn't go below the initial height

// this essentially makes stop time a void criteria for termination

func (r *Reactor) backfill(
	ctx context.Context,
	chainID string,
	startHeight, stopHeight, initialHeight int64,
	trustedBlockID types.BlockID,
	stopTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// fetch light blocks across four workers. The aim with deploying concurrent
// workers is to equate the network messaging time with the verification
// time. Ideally we want the verification process to never have to be
// waiting on blocks. If it takes 4s to retrieve a block and 1s to verify
// it, then steady state involves four workers.

// pop the next peer of the list to send a request to

// request the light block with a timeout

// once the peer has returned a value, add it back to the peer list to be used again

// we don't punish the peer as it might just have not responded in time

// As we are fetching blocks backwards, if this node doesn't have the block it likely doesn't
// have any prior ones, thus we remove it from the peer list.

// run a validate basic. This checks the validator set and commit
// hashes line up

// add block to queue to be verified

// verify all light blocks

// validate the header hash. We take the last block id of the
// previous header (i.e. one height above) as the trusted hash which
// we equate to. ValidatorsHash and CommitHash have already been
// checked in the `ValidateBasic`

// save the signed headers

// check if there has been a change in the validator set

// save all the heights that the last validator set was the same

// update the lastChangeHeight

// The block height might be less than the stopHeight because of the stopTime condition
// hasn't been fulfilled.

// save the final batch of validators

// handleSnapshotMessage handles ms sent from peers on the
// SnapshotChannel. It returns an error only if the Envelope.Message is unknown
// for this channel. This should never be called outside of handleMessage.
func (r *Reactor) handleSnapshotMessage(ctx context.Context, m p2p.RecvMsg[*pb.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// handleChunkMessage handles ms sent from peers on the ChunkChannel.
// It returns an error only if the Envelope.Message is unknown for this channel.
// This should never be called outside of handleMessage.
func (r *Reactor) handleChunkMessage(ctx context.Context, m p2p.RecvMsg[*pb.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) handleLightBlockMessage(ctx context.Context, m p2p.RecvMsg[*pb.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: If we don't have the light block we will send a nil light block
// back to the requested node, indicating that we don't have it.

func (r *Reactor) handleParamsMessage(ctx context.Context, m p2p.RecvMsg[*pb.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // height from peer is validated above

// It is not peers fault that we cannot send it consensus params. Just log the received error.

func (r *Reactor) recoverToErr(err *error) { _ = "STUB: not implemented"; return }

func (r *Reactor) processSnapshotCh(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *Reactor) processChunkCh(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *Reactor) processLightBlockCh(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *Reactor) processParamsCh(ctx context.Context) { _ = "STUB: not implemented"; return }

// processPeerUpdate processes a PeerUpdate, returning an error upon failing to
// handle the PeerUpdate or if a panic is recovered.
func (r *Reactor) processPeerUpdate(peerUpdate p2p.PeerUpdate) { _ = "STUB: not implemented"; return }

// Reset

// we do this in a separate routine to not block whilst waiting for the light client to finish
// whatever call it's currently executing

// processPeerUpdates initiates a blocking process where we listen for and handle
// PeerUpdate messages. When the reactor is stopped, we will catch the signal and
// close the p2p PeerUpdatesCh gracefully.
func (r *Reactor) processPeerUpdates(ctx context.Context) { _ = "STUB: not implemented"; return }

// recentSnapshots fetches the n most recent snapshots from the app
func (r *Reactor) recentSnapshots(ctx context.Context, n uint32) ([]*snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetchLightBlock works out whether the node has a light block at a particular
// height and if so returns it so it can be gossiped to peers
func (r *Reactor) fetchLightBlock(height uint64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	//nolint:gosec // height validated by Message.Validate() upstream
	return nil, nil
}

func (r *Reactor) waitForEnoughPeers(ctx context.Context, numPeers int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) TotalSnapshots() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reactor) ChunkProcessAvgTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *Reactor) SnapshotHeight() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reactor) SnapshotChunksCount() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reactor) SnapshotChunksTotal() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reactor) BackFilledBlocks() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reactor) BackFillBlocksTotal() int64 { _ = "STUB: not implemented"; return 0 }
