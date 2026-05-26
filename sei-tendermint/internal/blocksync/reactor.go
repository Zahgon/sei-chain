package blocksync

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/consensus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/store"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/blocksync"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var _ service.Service = (*Reactor)(nil)

const (
	// BlockSyncChannel is a channel for blocks and status updates
	BlockSyncChannel = p2p.ChannelID(0x40)

	trySyncIntervalMS = 10

	// ask for best height every 10s
	statusUpdateInterval = 10 * time.Second

	// check if we should switch to consensus reactor
	switchToConsensusIntervalSeconds = 1

	// switch to consensus after this duration of inactivity
	syncTimeout = 180 * time.Second
)

// TODO(gprusak): that's not sufficient - parsing proto requires checking nils everywhere.
func wrap[T *pb.BlockRequest | *pb.NoBlockResponse | *pb.BlockResponse | *pb.StatusRequest | *pb.StatusResponse](msg T) *pb.Message {
	_ = "STUB: not implemented"
	return nil
}

func GetChannelDescriptor() p2p.ChannelDescriptor[*pb.Message] {
	_ = "STUB: not implemented"
	return nil
}

type consensusReactor interface {
	// For when we switch from block sync reactor to the consensus
	// machine.
	SwitchToConsensus(state sm.State, skipWAL bool)
}

type peerError struct {
	err    error
	peerID types.NodeID
}

func (e peerError) Error() string { _ = "STUB: not implemented"; return "" }

type blocksyncResult struct{ stateSynced bool }

// Reactor handles long-term catchup syncing.
type Reactor struct {
	service.BaseService

	// immutable
	initialState sm.State
	// store
	stateStore sm.Store

	blockExec             *sm.BlockExecutor
	store                 sm.BlockStore
	pool                  *BlockPool
	consReactor           consensusReactor
	blockSync             *atomicBool
	previousMaxPeerHeight int64

	// blocksyncReady fires when blocksync should start processing blocks —
	// either at OnStart (if blockSync was initially set) or via
	// SwitchToBlockSync. Pre-spawned requestRoutine and poolRoutine wait on
	// it before doing any work.
	blocksyncReady utils.AtomicSend[utils.Option[blocksyncResult]]
	// consensusReady fires once the blocksync->consensus handoff has
	// happened. The pre-spawned autoRestartIfBehind monitor gates on this
	// signal.
	consensusReady utils.AtomicSend[bool]

	router  *p2p.Router
	channel *p2p.Channel[*pb.Message]

	requestsCh <-chan BlockRequest
	errorsCh   <-chan peerError

	metrics  *consensus.Metrics
	eventBus *eventbus.EventBus

	syncStartTime time.Time

	restartEvent              func()
	lastRestartTime           time.Time
	blocksBehindThreshold     uint64
	blocksBehindCheckInterval time.Duration
	restartCooldownSeconds    uint64
}

// NewReactor returns new reactor instance.
func NewReactor(
	stateStore sm.Store,
	blockExec *sm.BlockExecutor,
	store *store.BlockStore,
	consReactor consensusReactor,
	router *p2p.Router,
	blockSync bool,
	metrics *consensus.Metrics,
	eventBus *eventbus.EventBus,
	restartEvent func(), // should be idempotent and non-blocking
	selfRemediationConfig *config.SelfRemediationConfig,
) (*Reactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // validated in config.ValidateBasic against MaxInt64

// OnStart starts separate go routines for each p2p Channel and listens for
// envelopes on each. In addition, it also listens for peer updates and handles
// messages on that p2p channel accordingly. The caller must be sure to execute
// OnStop to ensure the outbound p2p Channels are closed.
//
// If blockSync is enabled, we also start the pool and the pool processing
// goroutine. If the pool fails to start, an error is returned.
func (r *Reactor) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// NOTE: The capacity should be larger than the peer count.

// Pre-spawn all long-running routines so their lifetime is bound to the
// BaseService WaitGroup. Conditional routines gate on AtomicSend[bool]
// signals so SwitchToBlockSync (and the in-poolRoutine consensus handoff
// for autoRestartIfBehind) can wake them later without spawning fresh
// goroutines from outside OnStart.

// OnStop stops the BlockPool. The reactor's own long-running goroutines were
// registered with the BaseService WaitGroup via Spawn in OnStart, so the
// BaseService blocks Stop() on their exit before this method returns.
func (r *Reactor) OnStop() { _ = "STUB: not implemented"; return }

// respondToPeer loads a block and sends it to the requesting peer, if we have it.
// Otherwise, we'll respond saying we do not have it.
func (r *Reactor) respondToPeer(msg *pb.BlockRequest, peerID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// handleMessage handles an Envelope sent from a peer on a specific p2p Channel.
// It will handle errors and any possible panics gracefully. A caller can handle
// any error returned by sending a PeerError on the respective channel.
func (r *Reactor) handleMessage(m p2p.RecvMsg[*pb.Message]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// processBlockSyncCh initiates a blocking process where we listen for and handle
// envelopes on the BlockSyncChannel and blockSyncOutBridgeCh. Any error encountered during
// message execution will result in a PeerError being sent on the BlockSyncChannel.
// When the reactor is stopped, we will catch the signal and close the p2p Channel
// gracefully.
func (r *Reactor) processBlockSyncCh(ctx context.Context) { _ = "STUB: not implemented"; return }

// autoRestartIfBehind will check if the node is behind the max peer height by
// a certain threshold. If it is, the node will attempt to restart itself
// TODO(gprusak): this should be a sub task of the consensus reactor instead.
func (r *Reactor) autoRestartIfBehind(ctx context.Context) { _ = "STUB: not implemented"; return }

//nolint:gosec // validated in config.ValidateBasic against MaxInt64

// We do not restart if we are not lagging behind, or we are already in block sync mode

// Check if we have met cooldown time

// Send signal to restart the node

// processPeerUpdate processes a PeerUpdate.
func (r *Reactor) processPeerUpdate(peerUpdate p2p.PeerUpdate) { _ = "STUB: not implemented"; return }

// send a status update the newly added peer

// processPeerUpdates initiates a blocking process where we listen for and handle
// PeerUpdate messages. When the reactor is stopped, we will catch the signal and
// close the p2p PeerUpdatesCh gracefully.
func (r *Reactor) processPeerUpdates(ctx context.Context) { _ = "STUB: not implemented"; return }

// SwitchToBlockSync is called by the state sync reactor when switching to fast
// sync.
func (r *Reactor) SwitchToBlockSync(ctx context.Context, state sm.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reactor) requestRoutine(ctx context.Context) { _ = "STUB: not implemented"; return }

// poolRoutine handles messages from the poolReactor telling the reactor what to
// do.
//
// NOTE: Don't sleep in the FOR_LOOP or otherwise slow it down!
func (r *Reactor) poolRoutine(ctx context.Context, stateSynced bool) {
	_ = "STUB: not implemented"
	return
}

// Use the node-scoped context: SwitchToConsensus is a handoff
// to a peer reactor whose lifecycle is not tied to blocksync.

// Wake the pre-spawned auto-restart monitor.

// NOTE: It is a subtle mistake to process more than a single block at a
// time (e.g. 10) here, because we only send one BlockRequest per loop
// iteration. The ratio mismatch can result in starving of blocks, i.e. a
// sudden burst of requests and responses, and repeat. Consequently, it is
// better to split these routines rather than coupling them as it is
// written here.
//
// TODO: Uncouple from request routine.

// see if there are any blocks to sync

// we need to have fetched two consecutive blocks in order to perform blocksync verification

// try again quickly next loop

// Finally, verify the first block using the second's commit.
//
// NOTE: We can probably make this more efficient, but note that calling
// first.Hash() doesn't verify the tx contents, so MakePartSet() is
// currently necessary.
// TODO(sergio): Should we also validate against the extended commit?

// validate the block before we persist it

// If either of the checks failed we log the error and request for a new block
// at that height

// NOTE: We've already removed the peer's request, but we still need
// to clean up the rest.

// We use LastCommit here instead of extCommit. extCommit is not
// guaranteed to be populated by the peer if extensions are not enabled.
// Currently, the peer should provide an extCommit even if the vote extension data are absent
// but this may change so using second.LastCommit is safer.

// TODO: Same thing for app - but we would need a way to get the hash
// without persisting the state.

func (r *Reactor) GetMaxPeerBlockHeight() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reactor) GetTotalSyncedTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *Reactor) GetRemainingSyncTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *Reactor) PublishStatus(event types.EventDataBlockSyncStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// atomicBool is an atomic Boolean, safe for concurrent use by multiple
// goroutines.
type atomicBool int32

// newAtomicBool creates an atomicBool with given initial value.
func newAtomicBool(ok bool) *atomicBool { _ = "STUB: not implemented"; return nil }

// Set sets the Boolean to true.
func (ab *atomicBool) Set() { _ = "STUB: not implemented"; return }

// UnSet sets the Boolean to false.
func (ab *atomicBool) UnSet() { _ = "STUB: not implemented"; return }

// IsSet returns whether the Boolean is true.
func (ab *atomicBool) IsSet() bool { _ = "STUB: not implemented"; return false }
