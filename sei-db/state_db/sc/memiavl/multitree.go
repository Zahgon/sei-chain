package memiavl

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/alitto/pond"
	"golang.org/x/time/rate"

	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/wal"
)

const (
	MetadataFileName = "__metadata"

	// Load time threshold for performance warning (300 seconds)
	slowLoadThreshold = 5 * time.Minute
)

type NamedTree struct {
	*Tree
	Name string
}

// MultiTree manages multiple memiavl tree together,
// all the trees share the same latest version, the snapshots are always created at the same version.
//
// The snapshot structure is like this:
// ```
// > snapshot-V
// >  metadata
// >  bank
// >   kvs
// >   nodes
// >   metadata
// >  acc
// >  other stores...
// ```
type MultiTree struct {
	// if the tree is start from genesis, it's the initial version of the chain,
	// if the tree is imported from snapshot, it's the imported version plus one,
	// it always corresponds to the rlog entry with index 1.
	// Use atomic for concurrent read/write safety
	initialVersion atomic.Uint32

	zeroCopy bool

	trees          []NamedTree    // always ordered by tree name
	treesByName    map[string]int // index of the trees by name
	lastCommitInfo proto.CommitInfo

	// the initial metadata loaded from disk snapshot
	metadata proto.MultiTreeMetadata
}

func NewEmptyMultiTree(initialVersion uint32) *MultiTree { _ = "STUB: not implemented"; return nil }

func LoadMultiTree(ctx context.Context, dir string, opts Options) (*MultiTree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check for cancellation

// initial version is necessary for rlog index conversion

// TreeByName returns the tree by name, returns nil if not found
func (t *MultiTree) TreeByName(name string) *Tree { _ = "STUB: not implemented"; return nil }

// Trees returns all the trees together with the name, ordered by name.
func (t *MultiTree) Trees() []NamedTree { _ = "STUB: not implemented"; return nil }

func (t *MultiTree) SetInitialVersion(initialVersion int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *MultiTree) setInitialVersion(initialVersion int64) { _ = "STUB: not implemented"; return }

func (t *MultiTree) SetZeroCopy(zeroCopy bool) { _ = "STUB: not implemented"; return }

// Copy returns a snapshot of the tree which won't be corrupted by further modifications on the main tree.
func (t *MultiTree) Copy() *MultiTree { _ = "STUB: not implemented"; return nil }

func (t *MultiTree) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (t *MultiTree) SnapshotVersion() int64 { _ = "STUB: not implemented"; return 0 }

func (t *MultiTree) LastCommitInfo() *proto.CommitInfo { _ = "STUB: not implemented"; return nil }

//lint:ignore U1000 lifecycle method retained for completeness
func (t *MultiTree) apply(entry proto.ChangelogEntry) error { _ = "STUB: not implemented"; return nil }

// ApplyUpgrades store name upgrades
func (t *MultiTree) ApplyUpgrades(upgrades []*proto.TreeNameUpgrade) error {
	_ = "STUB: not implemented"
	return nil
}

// rebuild in the end

// swap deletion

// rename tree

// add tree

// ApplyChangeSet applies change set for a single tree.
func (t *MultiTree) ApplyChangeSet(name string, changeSet proto.ChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyChangeSets applies change sets for multiple trees.
func (t *MultiTree) ApplyChangeSets(changeSets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

// WorkingCommitInfo returns the commit info for the working tree
func (t *MultiTree) WorkingCommitInfo() *proto.CommitInfo { _ = "STUB: not implemented"; return nil }

// SaveVersion bumps the versions of all the stores and optionally returns the new app hash
func (t *MultiTree) SaveVersion(updateCommitInfo bool) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// clear the dirty informaton

func (t *MultiTree) buildCommitInfo(version int64) *proto.CommitInfo {
	_ = "STUB: not implemented"
	return nil
}

// UpdateCommitInfo update lastCommitInfo based on current status of trees.
// it's needed if `updateCommitInfo` is set to `false` in `ApplyChangeSet`.
func (t *MultiTree) UpdateCommitInfo() { _ = "STUB: not implemented"; return }

// Catchup replays WAL entries to catch up the tree to the target or latest version.
// delta is the difference between version and WAL index (version = walIndex + delta).
// endVersion specifies the target version (0 means catch up to latest).
func (t *MultiTree) Catchup(ctx context.Context, stream wal.ChangelogWAL, delta int64, endVersion int64) error {
	_ = "STUB: not implemented"
	return nil

	// Get actual WAL index range
}

// Empty WAL - nothing to replay

// Calculate start index: walIndex = version - delta
// We want to start from currentVersion + 1

// Ensure startIndex is within valid range (handle negative case before uint64 conversion)

// Nothing to replay - tree is already caught up

// Check for cancellation

// Safety check: skip entries we already have (should not happen with correct startIndex)

// If endVersion is specified, stop at that version

func (t *MultiTree) WriteSnapshot(ctx context.Context, dir string, wp *pond.WorkerPool) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteSnapshotWithRateLimit writes snapshot with optional rate limiting.
// rateMBps is the rate limit in MB/s. 0 means unlimited.
// A single global limiter is shared across ALL trees and files to ensure
// the total write rate is capped at the configured value.
func (t *MultiTree) WriteSnapshotWithRateLimit(ctx context.Context, dir string, wp *pond.WorkerPool, rateMBps int) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

// Create a single global limiter shared by all trees and files
// This ensures total write rate is capped regardless of parallelism

// Write EVM first to avoid disk I/O contention, then parallel

// writeSnapshotPriorityEVM writes EVM tree first, then others in parallel
// Best strategy: reduces disk I/O contention for the largest tree
// limiter is a shared rate limiter. nil means unlimited.
func (t *MultiTree) writeSnapshotPriorityEVM(ctx context.Context, dir string, wp *pond.WorkerPool, limiter *rate.Limiter) error {
	_ = "STUB: not implemented"
	return nil

	// Phase 1: Write EVM tree first (if it exists)
}

// Phase 2: Write all other trees in parallel

// NOTE: We use explicit WaitGroup instead of pond.GroupContext because
// GroupContext.Wait() returns immediately on context cancellation without
// waiting for workers to finish. This causes data races when DB.Close()
// destroys mmap resources that background writers are still accessing.
// The WaitGroup ensures all goroutines fully exit before we return.

// Capture loop variables for goroutine closure to avoid data race
// Create new variable for closure capture

// write commit info

// WriteFileSync calls `f.Sync` after before closing the file
func WriteFileSync(name string, data []byte) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func (t *MultiTree) Close() error { _ = "STUB: not implemented"; return nil }

func (t *MultiTree) ReplaceWith(other *MultiTree) error { _ = "STUB: not implemented"; return nil }

func readMetadata(dir string) (*proto.MultiTreeMetadata, error) {
	_ = "STUB: not implemented"
	// load commit info
	return nil, nil
}
