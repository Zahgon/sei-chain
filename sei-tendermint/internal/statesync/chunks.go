package statesync

import (
	"errors"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// errDone is returned by chunkQueue.Next() when all chunks have been returned.
var errDone = errors.New("chunk queue has completed")

// chunk contains data for a chunk.
type chunk struct {
	Height uint64
	Format uint32
	Index  uint32
	Chunk  []byte
	Sender types.NodeID
}

// chunkQueue manages chunks for a state sync process, ordering them if requested. It acts as an
// iterator over all chunks, but callers can request chunks to be retried, optionally after
// refetching.
type chunkQueue struct {
	sync.Mutex
	snapshot       *snapshot                  // if this is nil, the queue has been closed
	dir            string                     // temp dir for on-disk chunk storage
	chunkFiles     map[uint32]string          // path to temporary chunk file
	chunkSenders   map[uint32]types.NodeID    // the peer who sent the given chunk
	chunkAllocated map[uint32]bool            // chunks that have been allocated via Allocate()
	chunkReturned  map[uint32]bool            // chunks returned via Next()
	waiters        map[uint32][]chan<- uint32 // signals WaitFor() waiters about chunk arrival
}

// newChunkQueue creates a new chunk queue for a snapshot, using a temp dir for storage.
// Callers must call Close() when done.
func newChunkQueue(snapshot *snapshot, tempDir string) (*chunkQueue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds a chunk to the queue. It ignores chunks that already exist, returning false.
func (q *chunkQueue) Add(chunk *chunk) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// queue is closed

// Signal any waiters that the chunk has arrived.

// Allocate allocates a chunk to the caller, making it responsible for fetching it. Returns
// errDone once no chunks are left or the queue is closed.
func (q *chunkQueue) Allocate() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // map length bounded by snapshot.Chunks which is uint32

// Close closes the chunk queue, cleaning up all temporary files.
func (q *chunkQueue) Close() error { _ = "STUB: not implemented"; return nil }

// Discard discards a chunk. It will be removed from the queue, available for allocation, and can
// be added and returned via Next() again. If the chunk is not already in the queue this does
// nothing, to avoid it being allocated to multiple fetchers.
func (q *chunkQueue) Discard(index uint32) error { _ = "STUB: not implemented"; return nil }

// discard discards a chunk, scheduling it for refetching. The caller must hold the mutex lock.
func (q *chunkQueue) discard(index uint32) error { _ = "STUB: not implemented"; return nil }

// DiscardSender discards all *unreturned* chunks from a given sender. If the caller wants to
// discard already returned chunks, this can be done via Discard().
func (q *chunkQueue) DiscardSender(peerID types.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSender returns the sender of the chunk with the given index, or empty if
// not found.
func (q *chunkQueue) GetSender(index uint32) types.NodeID {
	_ = "STUB: not implemented"
	return *new(types.NodeID)
}

// Has checks whether a chunk exists in the queue.
func (q *chunkQueue) Has(index uint32) bool { _ = "STUB: not implemented"; return false }

// load loads a chunk from disk, or nil if the chunk is not in the queue. The caller must hold the
// mutex lock.
func (q *chunkQueue) load(index uint32) (*chunk, error) { _ = "STUB: not implemented"; return nil, nil }

// Next returns the next chunk from the queue, or errDone if all chunks have been returned. It
// blocks until the chunk is available. Concurrent Next() calls may return the same chunk.
func (q *chunkQueue) Next() (*chunk, error) { _ = "STUB: not implemented"; return nil, nil }

// queue closed

// nextUp returns the next chunk to be returned, or errDone if all chunks have been returned. The
// caller must hold the mutex lock.
func (q *chunkQueue) nextUp() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// Retry schedules a chunk to be retried, without refetching it.
func (q *chunkQueue) Retry(index uint32) { _ = "STUB: not implemented"; return }

// RetryAll schedules all chunks to be retried, without refetching them.
func (q *chunkQueue) RetryAll() { _ = "STUB: not implemented"; return }

// Size returns the total number of chunks for the snapshot and queue, or 0 when closed.
func (q *chunkQueue) Size() uint32 { _ = "STUB: not implemented"; return 0 }

// WaitFor returns a channel that receives a chunk index when it arrives in the queue, or
// immediately if it has already arrived. The channel is closed without a value if the queue is
// closed or if the chunk index is not valid.
func (q *chunkQueue) WaitFor(index uint32) <-chan uint32 { _ = "STUB: not implemented"; return nil }

func (q *chunkQueue) numChunksReturned() int { _ = "STUB: not implemented"; return 0 }
