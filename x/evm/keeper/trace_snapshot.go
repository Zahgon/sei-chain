package keeper

import (
	"sync"

	sctypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

// TraceSnapshotStore holds bounded in-memory SC snapshots keyed by block height.
// Each entry is a Committer.Copy() sharing COW nodes with the live memiavl tree.
//
// Operational signals to watch when this store is in use:
//   - memiavl gauges MemNodeTotalSize / NumOfMemNode: rise if held snapshots
//     pin too many COW nodes (bigger window or higher TPS amplifies this).
//   - trace baker counters (BakedCount / DroppedCount / FailedCount): if
//     DroppedCount climbs or BakedCount lags the chain tip, the baker is
//     falling behind and trace cache hit rate will drop.
type TraceSnapshotStore struct {
	mu        sync.Mutex
	snapshots map[int64]sctypes.Committer
	window    int64
	latest    int64
}

type snapshotRefReleaser interface {
	ReleaseSnapshotRefs() error
}

func NewTraceSnapshotStore(window int64) *TraceSnapshotStore { _ = "STUB: not implemented"; return nil }

// Put records a snapshot and evicts entries older than (latest - window).
// In-flight traces use Lease, so evicted map entries can be closed explicitly.
func (s *TraceSnapshotStore) Put(height int64, snap sctypes.Committer) {
	_ = "STUB: not implemented"
	return
}

// Lease returns an owned snapshot copy and a release function for trace state.
func (s *TraceSnapshotStore) Lease(height int64) (sctypes.Committer, func()) {
	_ = "STUB: not implemented"
	return *new(sctypes.Committer), nil
}

func (s *TraceSnapshotStore) Get(height int64) sctypes.Committer {
	_ = "STUB: not implemented"
	return *new(sctypes.Committer)
}

// Close releases all retained snapshots.
func (s *TraceSnapshotStore) Close() { _ = "STUB: not implemented"; return }

func releaseSnapshotRefs(snap sctypes.Committer) { _ = "STUB: not implemented"; return }
