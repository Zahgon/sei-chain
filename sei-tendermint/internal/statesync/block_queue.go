package statesync

import (
	"sync"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type lightBlockResponse struct {
	block *types.LightBlock
	peer  types.NodeID
}

// a block queue is used for asynchronously fetching and verifying light blocks
type blockQueue struct {
	mtx sync.Mutex

	// cursors to keep track of which heights need to be fetched and verified
	fetchHeight  int64
	verifyHeight int64

	// termination conditions
	initialHeight int64
	stopHeight    int64
	stopTime      time.Time
	terminal      *types.LightBlock

	// track failed heights so we know what blocks to try fetch again
	failed *maxIntHeap
	// also count retries to know when to give up
	retries    int
	maxRetries int

	// store inbound blocks and serve them to a verifying thread via a channel
	pending  map[int64]lightBlockResponse
	verifyCh chan lightBlockResponse

	// waiters are workers on idle until a height is required
	waiters []chan int64

	// this channel is closed once the verification process is complete
	doneCh chan struct{}
}

func newBlockQueue(
	startHeight, stopHeight, initialHeight int64,
	stopTime time.Time,
	maxRetries int,
) *blockQueue {
	_ = "STUB: not implemented"
	return nil
}

// Add adds a block to the queue to be verified and stored
// CONTRACT: light blocks should have passed basic validation
func (q *blockQueue) add(l lightBlockResponse) { _ = "STUB: not implemented"; return }

// return early if the process has already finished

// sometimes more blocks are fetched then what is necessary. If we already
// have what we need then ignore this

// if the block that was returned is at the verify height then the verifier
// is already waiting for this block so we send it directly to them

// else we add it in the pending bucket

// Lastly, if the incoming block is past the stop time and stop height or
// is equal to the initial height then we mark it as the terminal block.

// NextHeight returns the next height that needs to be retrieved.
// We assume that for every height allocated that the peer will eventually add
// the block or signal that it needs to be retried
func (q *blockQueue) nextHeight() <-chan int64 { _ = "STUB: not implemented"; return nil }

// if a previous process failed then we pick up this one

// return and decrement the fetch height

// at this point there is no height that we know we need so we create a
// waiter to hold out for either an outgoing request to fail or a block to
// fail verification

// Finished returns true when the block queue has has all light blocks retrieved,
// verified and stored. There is no more work left to be done
func (q *blockQueue) done() <-chan struct{} {
	_ = "STUB: not implemented"

	// VerifyNext pulls the next block off the pending queue and adds it to a
	// channel if it's already there or creates a waiter to add it to the
	// channel once it comes in. NOTE: This is assumed to
	// be a single thread as light blocks need to be sequentially verified.
	return nil
}

func (q *blockQueue) verifyNext() <-chan lightBlockResponse { _ = "STUB: not implemented"; return nil }

// Retry is called when a dispatcher failed to fetch a light block or the
// fetched light block failed verification. It signals to the queue to add the
// height back to the request queue
func (q *blockQueue) retry(height int64) { _ = "STUB: not implemented"; return }

// we don't need to retry if this is below the terminal height

// Success is called when a light block has been successfully verified and
// processed
func (q *blockQueue) success() { _ = "STUB: not implemented"; return }

func (q *blockQueue) error() error { _ = "STUB: not implemented"; return nil }

// close the queue and respective channels
func (q *blockQueue) close() { _ = "STUB: not implemented"; return }

// CONTRACT: must have a write lock. Use close instead
func (q *blockQueue) _closeChannels() {
	_ = "STUB: not implemented"

	// wait for the channel to be drained
	return
}

// A max-heap of ints.
type maxIntHeap []int64

func (h maxIntHeap) Len() int           { _ = "STUB: not implemented"; return 0 }
func (h maxIntHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (h maxIntHeap) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (h *maxIntHeap) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (h *maxIntHeap) Pop() interface{} { _ = "STUB: not implemented"; return nil }
