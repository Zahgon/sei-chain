package statesync

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/statesync"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const (
	// chunkTimeout is the timeout while waiting for the next chunk from the chunk queue.
	chunkTimeout = 2 * time.Minute

	// minimumDiscoveryTime is the lowest allowable time for a
	// SyncAny discovery time.
	minimumDiscoveryTime = 5 * time.Second
)

var (
	// errAbort is returned by Sync() when snapshot restoration is aborted.
	errAbort = errors.New("state sync aborted")
	// errRetrySnapshot is returned by Sync() when the snapshot should be retried.
	errRetrySnapshot = errors.New("retry snapshot")
	// errRejectSnapshot is returned by Sync() when the snapshot is rejected.
	errRejectSnapshot = errors.New("snapshot was rejected")
	// errRejectFormat is returned by Sync() when the snapshot format is rejected.
	errRejectFormat = errors.New("snapshot format was rejected")
	// errRejectSender is returned by Sync() when the snapshot sender is rejected.
	errRejectSender = errors.New("snapshot sender was rejected")
	// errVerifyFailed is returned by Sync() when app hash or last height
	// verification fails.
	errVerifyFailed = errors.New("verification with app failed")
	// errTimeout is returned by Sync() when we've waited too long to receive a chunk.
	errTimeout = errors.New("timed out waiting for chunk")
	// errNoSnapshots is returned by SyncAny() if no snapshots are found and discovery is disabled.
	errNoSnapshots = errors.New("no suitable snapshots found")
)

// syncer runs a state sync against an ABCI app. Use either SyncAny() to automatically attempt to
// sync all snapshots in the pool (pausing to discover new ones), or Sync() to sync a specific
// snapshot. Snapshots and chunks are fed via AddSnapshot() and AddChunk() as appropriate.
type syncer struct {
	stateProvider StateProvider
	conn          *proxy.Proxy
	snapshots     *snapshotPool
	snapshotCh    *p2p.Channel[*pb.Message]
	chunkCh       *p2p.Channel[*pb.Message]
	tempDir       string
	fetchers      int32
	retryTimeout  time.Duration

	mtx     sync.RWMutex
	chunks  *chunkQueue
	metrics *Metrics

	avgChunkTime             int64
	lastSyncedSnapshotHeight int64
	processingSnapshot       *snapshot
	useLocalSnapshot         bool
}

// AddChunk adds a chunk to the chunk queue, if any. It returns false if the chunk has already
// been added to the queue, or an error if there's no sync in progress.
func (s *syncer) AddChunk(chunk *chunk) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// AddSnapshot adds a snapshot to the snapshot pool. It returns true if a new, previously unseen
// snapshot was accepted and added.
func (s *syncer) AddSnapshot(peerID types.NodeID, snapshot *snapshot) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AddPeer adds a peer to the pool. For now we just keep it simple and send a
// single request to discover snapshots, later we may want to do retries and stuff.
func (s *syncer) AddPeer(peerID types.NodeID) { _ = "STUB: not implemented"; return }

// RemovePeer removes a peer from the pool.
func (s *syncer) RemovePeer(peerID types.NodeID) { _ = "STUB: not implemented"; return }

// SyncAny tries to sync any of the snapshots in the snapshot pool, waiting to discover further
// snapshots if none were found and discoveryTime > 0. It returns the latest state and block commit
// which the caller must use to bootstrap the node.
func (s *syncer) SyncAny(
	ctx context.Context,
	discoveryTime time.Duration,
	requestSnapshots func() error,
) (sm.State, *types.Commit, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil, nil
}

// The app may ask us to retry a snapshot restoration, in which case we need to reuse
// the snapshot and chunk queue from the previous loop iteration.

// Ensure chunks is always closed on function exit

// If not nil, we're going to retry restoration of the same snapshot.

//nolint:gosec // snapshot.Height is a valid block height

// Discard snapshot and chunks for next iteration

// Sync executes a sync for a specific snapshot, returning the latest state and block commit which
// the caller must use to bootstrap the node.
func (s *syncer) Sync(ctx context.Context, snapshot *snapshot, chunks *chunkQueue) (sm.State, *types.Commit, error) {
	_ = "STUB: not implemented"
	return *new(sm.State), nil, nil
}

// Fetch the app hash corresponding to the snapshot

// check if the main context was triggered

// catch the case where all the light client providers have been exhausted

// Offer snapshot to ABCI app.

// Spawn chunk fetchers. They will terminate when the chunk queue is closed or context canceled.

// Optimistically build new state, so we don't discover any light client failures at the end.

// check if the main context was triggered

// check if the provider context exceeded the 10 seconds deadline

// Restore snapshot

// Verify app and app version

// Done! 🎉

// offerSnapshot offers a snapshot to the app. It returns various errors depending on the app's
// response, or nil if the snapshot was accepted.
func (s *syncer) offerSnapshot(ctx context.Context, snapshot *snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// applyChunks applies chunks to the app. It returns various errors depending on the app's
// response, or nil once the snapshot is fully restored.
func (s *syncer) applyChunks(ctx context.Context, chunks *chunkQueue, start time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// Discard and refetch any chunks as requested by the app

// Reject any senders as requested by the app

func (s *syncer) fetchLocalChunks(ctx context.Context, snapshot *snapshot, chunks *chunkQueue) {
	_ = "STUB: not implemented"
	return
}

// Keep checking until the context is canceled (restore is done), in case any
// chunks need to be refetched.

// fetchChunks requests chunks from peers, receiving allocations from the chunk queue. Chunks
// will be received from the reactor via syncer.AddChunks() to chunkQueue.Add().
func (s *syncer) fetchChunks(ctx context.Context, snapshot *snapshot, chunks *chunkQueue) {
	_ = "STUB: not implemented"
	return
}

// Keep checking until the context is canceled (restore is done), in case any
// chunks need to be refetched.

// requestChunk requests a chunk from a peer.
//
// returns nil if there are no peers for the given snapshot or the
// request is successfully made and an error if the request cannot be
// completed
func (s *syncer) requestChunk(snapshot *snapshot, chunk uint32) { _ = "STUB: not implemented"; return }

// verifyApp verifies the sync, checking the app hash, last block height and app version
func (s *syncer) verifyApp(ctx context.Context, snapshot *snapshot, appVersion uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// sanity check that the app version in the block matches the application's own record
// of its version

// An error here most likely means that the app hasn't inplemented state sync
// or the Info call correctly

//nolint:gosec // LastBlockHeight is a non-negative block height
