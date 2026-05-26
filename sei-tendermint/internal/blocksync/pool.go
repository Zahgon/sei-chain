package blocksync

import (
	"context"
	"sync"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/libs/flowrate"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

/*
eg, L = latency = 0.1s
	P = num peers = 10
	FN = num full nodes
	BS = 1kB block size
	CB = 1 Mbit/s = 128 kB/s
	CB/P = 12.8 kB
	B/S = CB/P/BS = 12.8 blocks/s

	12.8 * 0.1 = 1.28 blocks on conn
*/

var logger = seilog.NewLogger("tendermint", "internal", "blocksync")

const (
	requestInterval           = 100 * time.Millisecond
	maxTotalRequesters        = 50
	maxPeerErrBuffer          = 1000
	maxPendingRequests        = maxTotalRequesters
	maxPendingRequestsPerPeer = 20

	// Minimum recv rate to ensure we're receiving blocks from a peer fast
	// enough. If a peer is not sending us data at at least that rate, we
	// consider them to have timedout and we disconnect.
	//
	// Assuming a DSL connection (not a good choice) 128 Kbps (upload) ~ 15 KB/s,
	// sending data across atlantic ~ 7.5 KB/s.
	minRecvRate = 7680

	// Maximum difference between current and new block's height.
	maxDiffBetweenCurrentAndReceivedBlockHeight = 100

	// Used to indicate the reason of the redo
	PeerRemoved RetryReason = "PeerRemoved"
	BadBlock    RetryReason = "BadBlock"

	peerTimeout = 2 * time.Second
)

// Interface abstracting p2p.Router for tests.
type router interface {
	IsBlockSyncPeer(types.NodeID) bool
	Evict(id types.NodeID, err error)
	Connected(types.NodeID) bool
}

/*
	Peers self report their heights when we join the block pool.
	Starting from our latest pool.height, we request blocks
	in sequence from peers that reported higher heights than ours.
	Every so often we ask peers what height they're on so we can keep going.

	Requests are continuously made for blocks of higher heights until
	the limit is reached. If most of the requests have no available peers, and we
	are not at peer limits, we can probably switch to consensus reactor
*/

// BlockRequest stores a block request identified by the block Height and the
// PeerID responsible for delivering the block.
type BlockRequest struct {
	Height int64
	PeerID types.NodeID
}

// BlockPool keeps track of the block sync peers, block requests and block responses.
type BlockPool struct {
	service.BaseService

	lastAdvance time.Time

	mtx sync.RWMutex
	// block requests
	requesters map[int64]*bpRequester
	height     int64 // the lowest key in requesters.
	// peers
	peers         map[types.NodeID]*bpPeer
	router        router
	maxPeerHeight int64 // the biggest reported height

	// atomic
	numPending int32 // number of requests pending assignment or block response

	requestsCh chan<- BlockRequest
	errorsCh   chan<- peerError

	startHeight               int64
	lastHundredBlockTimeStamp time.Time
	lastSyncRate              float64
	cancels                   []context.CancelFunc
}

// NewBlockPool returns a new BlockPool with the height equal to start. Block
// requests and errors will be sent to requestsCh and errorsCh accordingly.
func NewBlockPool(
	start int64,
	requestsCh chan<- BlockRequest,
	errorsCh chan<- peerError,
	router router,
) *BlockPool {
	_ = "STUB: not implemented"
	return nil
}

// OnStart implements service.Service by spawning requesters routine and recording
// pool's start time.
func (pool *BlockPool) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (pool *BlockPool) OnStop() {
	_ = "STUB: not implemented"
	// Requester shutdown must not block behind a full requestsCh; Stop cancels ctx
	// and waits for the Spawn-managed requester goroutine to exit.
	return
}

// Stop requesters outside pool.mtx; their shutdown path may observe pool state.

// spawns requesters as needed
func (pool *BlockPool) makeRequestersRoutine(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// This is preferable to using a timer because the request interval
// is so small. Larger request intervals may necessitate using a
// timer/ticker.

// request for more blocks.

func (pool *BlockPool) removeTimedoutPeers() { _ = "STUB: not implemented"; return }

// check if peer timed out

// curRate can be 0 on start

// GetStatus returns pool's height, numPending requests and the number of
// requesters.
func (pool *BlockPool) GetStatus() (height int64, numPending int32, lenRequesters int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// IsCaughtUp returns true if this node is caught up, false - otherwise.
func (pool *BlockPool) IsCaughtUp() bool { _ = "STUB: not implemented"; return false }

// Need at least 2 peers to be considered caught up.

// NOTE: we use maxPeerHeight - 1 because to sync block H requires block H+1
// to verify the LastCommit.

// PeekTwoBlocks returns blocks at pool.height and pool.height+1. We need to
// see the second block's Commit to validate the first block. So we peek two
// blocks at a time. We return an extended commit, containing vote extensions
// and their associated signatures, as this is critical to consensus in ABCI++
// as we switch from block sync to consensus mode.
//
// The caller will verify the commit.
func (pool *BlockPool) PeekTwoBlocks() (first, second *types.Block) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PopRequest pops the first block at pool.height.
// It must have been validated by the second Commit from PeekTwoBlocks.
func (pool *BlockPool) PopRequest() { _ = "STUB: not implemented"; return }

// the lastSyncRate will be updated every 100 blocks, it uses the adaptive filter
// to smooth the block sync rate and the unit represents the number of blocks per second.

// RedoRequest invalidates the block at pool.height,
// Remove the peer and redo request from others.
// Returns the ID of the removed peer.
func (pool *BlockPool) RedoRequest(height int64) types.NodeID {
	_ = "STUB: not implemented"
	return *new(types.NodeID)
}

// Redo all requesters associated with this peer.

// AddBlock validates that the block comes from the peer it was expected from
// and calls the requester to store it.
//
// This requires an extended commit at the same height as the supplied block -
// the block contains the last commit, but we need the latest commit in case we
// need to switch over from block sync to consensus at this height. If the
// height of the extended commit and the height of the block do not match, we
// do not add the block and return an error.
// TODO: ensure that blocks come in order for each peer.
func (pool *BlockPool) AddBlock(peerID types.NodeID, block *types.Block, blockSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// MaxPeerHeight returns the highest reported height.
func (pool *BlockPool) MaxPeerHeight() int64 { _ = "STUB: not implemented"; return 0 }

// LastAdvance returns the time when the last block was processed (or start
// time if no blocks were processed).
func (pool *BlockPool) LastAdvance() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// SetPeerRange sets the peer's alleged blockchain base and height.
func (pool *BlockPool) SetPeerRange(peerID types.NodeID, base int64, height int64) {
	_ = "STUB: not implemented"
	return
}

// RemovePeer will redo all requesters associated with this peer.

// RemovePeer removes the peer with peerID from the pool. If there's no peer
// with peerID, function is a no-op.
func (pool *BlockPool) RemovePeer(peerID types.NodeID) { _ = "STUB: not implemented"; return }

func (pool *BlockPool) removePeer(peerID types.NodeID, redo bool) {
	_ = "STUB: not implemented"
	return
}

// Find a new peer with the biggest height and update maxPeerHeight if the
// peer's height was the biggest.

// If no peers are left, maxPeerHeight is set to 0.
func (pool *BlockPool) updateMaxPeerHeight() { _ = "STUB: not implemented"; return }

// Pick an available peer with the given height available.
// If no peers are available, returns nil.
func (pool *BlockPool) pickIncrAvailablePeer(height int64) *bpPeer {
	_ = "STUB: not implemented"
	return nil
}

// Remove peers with 0 score and shuffle list

// We only want to work with peers that are ready & connected (not dialing)

// randomly pick one with weak entropy.

func (pool *BlockPool) makeNextRequester(ctx context.Context) { _ = "STUB: not implemented"; return }

func (pool *BlockPool) requestersLen() int64 { _ = "STUB: not implemented"; return 0 }

func (pool *BlockPool) sendRequest(ctx context.Context, height int64, peerID types.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

func (pool *BlockPool) sendError(err error, peerID types.NodeID) { _ = "STUB: not implemented"; return }

func (pool *BlockPool) targetSyncBlocks() int64 { _ = "STUB: not implemented"; return 0 }

func (pool *BlockPool) getLastSyncRate() float64 { _ = "STUB: not implemented"; return 0 }

//-------------------------------------

type bpPeer struct {
	didTimeout  bool
	numPending  int32
	height      int64
	base        int64
	pool        *BlockPool
	id          types.NodeID
	recvMonitor *flowrate.Monitor

	timeout *time.Timer
	startAt time.Time
}

func (peer *bpPeer) resetMonitor() { _ = "STUB: not implemented"; return }

func (peer *bpPeer) resetTimeout() { _ = "STUB: not implemented"; return }

func (peer *bpPeer) incrPending() { _ = "STUB: not implemented"; return }

func (peer *bpPeer) decrPending(recvSize int) { _ = "STUB: not implemented"; return }

func (peer *bpPeer) onTimeout() { _ = "STUB: not implemented"; return }

//-------------------------------------

type bpRequester struct {
	service.BaseService
	pool          *BlockPool
	height        int64
	gotBlockCh    chan struct{}
	redoCh        chan RedoOp // redo may send multitime, add peerId to identify repeat
	timeoutTicker *time.Ticker
	mtx           sync.Mutex
	peerID        types.NodeID
	block         *types.Block
}

type RetryReason string

type RedoOp struct {
	PeerId types.NodeID
	Reason RetryReason
}

func newBPRequester(pool *BlockPool, height int64) *bpRequester {
	_ = "STUB: not implemented"
	return nil
}

func (bpr *bpRequester) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*bpRequester) OnStop() {
	_ = "STUB: not implemented"

	// Returns 0 if block doesn't already exist.
	// Returns -1 if peer doesn't match.
	// Return 1 if block exist and peer matches.
	return
}

func (bpr *bpRequester) setBlock(block *types.Block, peerID types.NodeID) int {
	_ = "STUB: not implemented"
	return 0
}

func (bpr *bpRequester) getBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (bpr *bpRequester) getPeerID() types.NodeID {
	_ = "STUB: not implemented"
	return *new(types.NodeID)
}

// This is called from the requestRoutine, upon redo().
func (bpr *bpRequester) reset(force bool) bool { _ = "STUB: not implemented"; return false }

// Do not reset if we already have a block

// Tells bpRequester to pick another peer and try again.
// NOTE: Nonblocking, and does nothing if another redo
// was already requested.
func (bpr *bpRequester) redo(peerID types.NodeID, retryReason RetryReason) {
	_ = "STUB: not implemented"
	return
}

// Responsible for making more requests as necessary
// Returns only when a block is found (e.g. AddBlock() is called)
func (bpr *bpRequester) requestRoutine(ctx context.Context) { _ = "STUB: not implemented"; return }

// Pick a peer to send request to.

// This is preferable to using a timer because the request
// interval is so small. Larger request intervals may
// necessitate using a timer/ticker.

// Send request and wait.

// if we don't have an existing block or this is a bad block
// we should reset the previous block

// We got a block!
