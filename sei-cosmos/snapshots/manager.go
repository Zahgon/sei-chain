package snapshots

import (
	"io"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/snapshots/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "snapshots")

const (
	opNone     operation = ""
	opSnapshot operation = "snapshot"
	opPrune    operation = "prune"
	opRestore  operation = "restore"

	chunkBufferSize = 4

	snapshotMaxItemSize = int(64e6) // SDK has no key/value size limit, so we set an arbitrary limit
)

// operation represents a Manager operation. Only one operation can be in progress at a time.
type operation string

// restoreDone represents the result of a restore operation.
type restoreDone struct {
	complete bool  // if true, restore completed successfully (not prematurely)
	err      error // if non-nil, restore errored
}

// Manager manages snapshot and restore operations for an app, making sure only a single
// long-running operation is in progress at any given time, and provides convenience methods
// mirroring the ABCI interface.
//
// Although the ABCI interface (and this manager) passes chunks as byte slices, the internal
// snapshot/restore APIs use IO streams (i.e. chan io.ReadCloser), for two reasons:
//
//  1. In the future, ABCI should support streaming. Consider e.g. InitChain during chain
//     upgrades, which currently passes the entire chain state as an in-memory byte slice.
//     https://github.com/tendermint/tendermint/issues/5184
//
//  2. io.ReadCloser streams automatically propagate IO errors, and can pass arbitrary
//     errors via io.Pipe.CloseWithError().
type Manager struct {
	store      *Store
	multistore types.Snapshotter
	extensions map[string]types.ExtensionSnapshotter

	mtx                sync.Mutex
	operation          operation
	chRestore          chan<- io.ReadCloser
	chRestoreDone      <-chan restoreDone
	restoreChunkHashes [][]byte
	restoreChunkIndex  uint32
}

// NewManager creates a new manager.
func NewManager(store *Store, multistore types.Snapshotter) *Manager {
	_ = "STUB: not implemented"
	return nil
}

// NewManagerWithExtensions creates a new manager.
func NewManagerWithExtensions(store *Store, multistore types.Snapshotter, extensions map[string]types.ExtensionSnapshotter) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) SetMultiStore(s types.Snapshotter) { _ = "STUB: not implemented"; return }

func (m *Manager) Close() error { _ = "STUB: not implemented"; return nil }

// RegisterExtensions register extension snapshotters to manager
func (m *Manager) RegisterExtensions(extensions ...types.ExtensionSnapshotter) error {
	_ = "STUB: not implemented"
	return nil
}

// begin starts an operation, or errors if one is in progress. It manages the mutex itself.
func (m *Manager) begin(op operation) error { _ = "STUB: not implemented"; return nil }

// beginLocked begins an operation while already holding the mutex.
func (m *Manager) beginLocked(op operation) error { _ = "STUB: not implemented"; return nil }

// end ends the current operation.
func (m *Manager) end() { _ = "STUB: not implemented"; return }

// endLocked ends the current operation while already holding the mutex.
func (m *Manager) endLocked() { _ = "STUB: not implemented"; return }

// sortedExtensionNames sort extension names for deterministic iteration.
func (m *Manager) sortedExtensionNames() []string { _ = "STUB: not implemented"; return nil }

// Create creates a snapshot and returns its metadata.
func (m *Manager) Create(height uint64) (*types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Spawn goroutine to generate snapshot chunks and pass their io.ReadClosers through a channel

// createSnapshot do the heavy work of snapshotting after the validations of request are done
// the produced chunks are written to the channel.
func (m *Manager) createSnapshot(height uint64, ch chan<- io.ReadCloser) {
	_ = "STUB: not implemented"
	return
}

// write extension metadata

// List lists snapshots, mirroring ABCI ListSnapshots. It can be concurrent with other operations.
func (m *Manager) List() ([]*types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil,

		// LoadChunk loads a chunk into a byte slice, mirroring ABCI LoadChunk. It can be called
		// concurrently with other operations. If the chunk does not exist, nil is returned.
		nil
}

func (m *Manager) LoadChunk(height uint64, format uint32, chunk uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prune prunes snapshots, if no other operations are in progress.
func (m *Manager) Prune(retain uint32) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Restore begins an async snapshot restoration, mirroring ABCI OfferSnapshot. Chunks must be fed
// via RestoreChunk() until the restore is complete or a chunk fails.
func (m *Manager) Restore(snapshot types.Snapshot) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // ChunkHashes length is bounded by Chunks which is uint32

// check multistore supported format preemptive

// Start an asynchronous snapshot restoration, passing chunks and completion status via channels.

// restoreSnapshot do the heavy work of snapshot restoration after preliminary checks on request have passed.
func (m *Manager) restoreSnapshot(snapshot types.Snapshot, chChunks <-chan io.ReadCloser) error {
	_ = "STUB: not implemented"
	return nil
}

// RestoreChunk adds a chunk to an active snapshot restoration, mirroring ABCI ApplySnapshotChunk.
// Chunks must be given until the restore is complete, returning true, or a chunk errors.
func (m *Manager) RestoreChunk(chunk []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check if any errors have occurred yet.

// Verify the chunk hash.

// Pass the chunk to the restore, and wait for completion if it was the final one.

// IsFormatSupported returns if the snapshotter supports restoration from given format.
func IsFormatSupported(snapshotter types.ExtensionSnapshotter, format uint32) bool {
	_ = "STUB: not implemented"
	return false
}
