package consensus

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "internal", "consensus")

// Functionality to replay blocks and messages on recovery from a crash.
// There are two general failure scenarios:
//
//  1. failure during consensus
//  2. failure while applying the block
//
// The former is handled by the WAL, the latter by the proxyApp Handshake on
// restart, which ultimately hands off the work to the WAL.

//-----------------------------------------
// 1. Recover from failure during consensus
// (by replaying messages from the WAL)
//-----------------------------------------

// Unmarshal and apply a single message to the consensus state as if it were
// received in receiveRoutine.
// NOTE: receiveRoutine should not be running.
func (cs *State) readReplayMessage(ctx context.Context, msg WALMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip meta messages which exist for demarcating boundaries.

// Replay only those messages since the last block.  `timeoutRoutine` should
// run concurrently to read off tickChan.
func (cs *State) catchupReplay(ctx context.Context, csHeight int64) error {
	_ = "STUB: not implemented"
	// Set replayMode to true so we don't log signing errors.
	return nil
}

// This is expected in case of state/block sync - we have not participated in
// the recent heights at all.

// NOTE: since the priv key is set when the msgs are received
// it will attempt to eg double sign but we can just ignore it
// since the votes will be replayed and we'll get to the next step

//---------------------------------------------------
// 2. Recover from failure while applying the block.
// (by handshaking with the app to figure out where
// we were last, and using the WAL to recover there.)
//---------------------------------------------------

type Handshaker struct {
	stateStore      sm.Store
	initialState    sm.State
	store           sm.BlockStore
	eventBus        *eventbus.EventBus
	genDoc          *types.GenesisDoc
	consensusPolicy types.ConsensusPolicy

	nBlocks int // number of blocks applied to the state
}

func NewHandshaker(
	stateStore sm.Store,
	state sm.State,
	store sm.BlockStore,
	eventBus *eventbus.EventBus,
	genDoc *types.GenesisDoc,
	consensusPolicy types.ConsensusPolicy,
) *Handshaker {
	_ = "STUB: not implemented"
	return nil
}

func newReplayTxMempool(app *proxy.Proxy) *mempool.TxMempool { _ = "STUB: not implemented"; return nil }

// NBlocks returns the number of blocks applied to the state.
func (h *Handshaker) NBlocks() int {
	_ = "STUB: not implemented"

	// TODO: retry the handshake/replay if it fails ?
	return 0
}

func (h *Handshaker) Handshake(ctx context.Context, app *proxy.Proxy) error {
	_ = "STUB: not implemented"
	return nil
}

// Only set the version if there is no existing state.

// Replay blocks up to the latest in the blockstore.

// TODO: (on restart) replay mempool

// ReplayBlocks replays all blocks since appBlockHeight and ensures the result
// matches the current state.
// Returns the final AppHash or an error.
func (h *Handshaker) ReplayBlocks(
	ctx context.Context,
	state sm.State,
	appHash []byte,
	appBlockHeight int64,
	app *proxy.Proxy,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If appBlockHeight == 0 it means that we are at genesis and hence should send InitChain.

// The validator set from genesis is expected to match the output of InitChain.

// we only update state when we are in initial state
// If the app did not return an app hash, we keep the one set from the genesis doc in
// the state. We don't set appHash since we don't want the genesis doc app hash
// recorded in the genesis block. We should probably just remove GenesisDoc.AppHash.

// If the app returned validators, update the state.

// If validator set is not set in genesis and still empty after InitChain, exit.

// We update the last results hash with the empty hash, to conform with RFC-6962.

// First handle edge cases and constraints on the storeBlockHeight and storeBlockBase.

// the app has no state, and the block store is truncated above the initial height

// the app is too far behind truncated store (can be 1 behind since we replay the next)

// the app should never be ahead of the store (but this is under app's control)

// the state should never be ahead of the store (this is under tendermint's control)

// store should be at most one ahead of the state (this is under tendermint's control)

// Now either store is equal to state, or one ahead.
// For each, consider all cases of where the app could be, given app <= store

// Tendermint ran Commit and saved the state.
// Either the app is asking for replay, or we're all synced up.

// the app is behind, so replay blocks, but no need to go through WAL (state is already synced to store)

// We're good! But we need to reindex events

// We saved the block in the store but haven't updated the state,
// so we'll need to replay a block using the WAL.

// the app is further behind than it should be, so replay blocks
// but leave the last block to go through the WAL

// We haven't run Commit (both the state and app are one block behind),
// so replayBlock with the real app.
// NOTE: We could instead use the cs.WAL on cs.Start,
// but we'd have to allow the WAL to replay a block that wrote it's #ENDHEIGHT

// We ran Commit, but didn't save the state, so replayBlock with mock app.

func (h *Handshaker) replayBlocks(
	ctx context.Context,
	state sm.State,
	app *proxy.Proxy,
	appBlockHeight,
	storeBlockHeight int64,
	mutateState bool,
) ([]byte, error) {
	_ = "STUB: not implemented"
	// App is further behind than it should be, so we need to replay blocks.
	// We replay all blocks from appBlockHeight+1.
	//
	// Note that we don't have an old version of the state,
	// so we by-pass state validation/mutation using sm.ExecCommitBlock.
	// This also means we won't be saving validator sets if they change during this period.
	// TODO: Load the historical information to fix this and just use state.ApplyBlock
	//
	// If mutateState == true, the final block is replayed with h.replayBlock()
	return nil, nil
}

// Extra check to ensure the app was not changed in a way it shouldn't have.

// We emit events for the index services at the final block due to the sync issue when
// the node shutdown during the block committing status.

// sync the final block

// ApplyBlock on the proxyApp with the last block.
func (h *Handshaker) replayBlock(
	ctx context.Context,
	state sm.State,
	height int64,
	app *proxy.Proxy,
) (sm.State, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil
}

// Use stubs for both mempool and evidence pool since no transactions nor
// evidence are needed here - block already exists.

// replayEvents will be called during restart to avoid tx missing to be indexed
func (h *Handshaker) replayEvents(height int64) error { _ = "STUB: not implemented"; return nil }

func checkAppHashEqualsOneFromBlock(appHash []byte, block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func checkAppHashEqualsOneFromState(appHash []byte, state sm.State) error {
	_ = "STUB: not implemented"
	return nil
}
