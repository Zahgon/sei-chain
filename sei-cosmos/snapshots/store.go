package snapshots

import (
	"io"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/snapshots/types"
	db "github.com/tendermint/tm-db"
)

const (
	// keyPrefixSnapshot is the prefix for snapshot database keys
	keyPrefixSnapshot byte = 0x01
)

// Store is a snapshot store, containing snapshot metadata and binary chunks.
type Store struct {
	db  db.DB
	dir string

	mtx    sync.Mutex
	saving map[uint64]bool // heights currently being saved
}

// NewStore creates a new snapshot store.
func NewStore(db db.DB, dir string) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

// Delete deletes a snapshot.
func (s *Store) Delete(height uint64, format uint32) error { _ = "STUB: not implemented"; return nil }

// Get fetches snapshot info from the database.
func (s *Store) Get(height uint64, format uint32) (*types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get fetches the latest snapshot from the database, if any.
func (s *Store) GetLatest() (*types.Snapshot, error) { _ = "STUB: not implemented"; return nil, nil }

// List lists snapshots, in reverse order (newest first).
func (s *Store) List() ([]*types.Snapshot, error) { _ = "STUB: not implemented"; return nil, nil }

// Load loads a snapshot (both metadata and binary chunks). The chunks must be consumed and closed.
// Returns nil if the snapshot does not exist.
func (s *Store) Load(height uint64, format uint32) (*types.Snapshot, <-chan io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// LoadChunk loads a chunk from disk, or returns nil if it does not exist. The caller must call
// Close() on it when done.
func (s *Store) LoadChunk(height uint64, format uint32, chunk uint32) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// loadChunkFile loads a chunk from disk, and errors if it does not exist.
func (s *Store) loadChunkFile(height uint64, format uint32, chunk uint32) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Prune removes old snapshots. The given number of most recent heights (regardless of format) are retained.
func (s *Store) Prune(retain uint32) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // len(skip) is bounded by number of distinct snapshot heights, practically small

// Since Delete() deletes a specific format, while we want to prune a height, we clean up
// the height directory as well

// Save saves a snapshot to disk, returning it.
func (s *Store) Save(
	height uint64, format uint32, chunks <-chan io.ReadCloser,
) (*types.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// saveSnapshot saves snapshot metadata to the database.
func (s *Store) saveSnapshot(snapshot *types.Snapshot) error { _ = "STUB: not implemented"; return nil }

// pathHeight generates the path to a height, containing multiple snapshot formats.
func (s *Store) pathHeight(height uint64) string { _ = "STUB: not implemented"; return "" }

// pathSnapshot generates a snapshot path, as a specific format under a height.
func (s *Store) pathSnapshot(height uint64, format uint32) string {
	_ = "STUB: not implemented"
	return ""
}

// pathChunk generates a snapshot chunk path.
func (s *Store) pathChunk(height uint64, format uint32, chunk uint32) string {
	_ = "STUB: not implemented"
	return ""
}

// decodeKey decodes a snapshot key.
func decodeKey(k []byte) (uint64, uint32, error) { _ = "STUB: not implemented"; return 0, 0, nil }

// encodeKey encodes a snapshot key.
func encodeKey(height uint64, format uint32) []byte { _ = "STUB: not implemented"; return nil }
