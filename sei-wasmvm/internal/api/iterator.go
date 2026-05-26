package api

import (
	"sync"

	"github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

// frame stores all Iterators for one contract call
type frame []types.Iterator

// iteratorFrames contains one frame for each contract call, indexed by contract call ID.
var (
	iteratorFrames      = make(map[uint64]frame)
	iteratorFramesMutex sync.Mutex
)

// this is a global counter for creating call IDs
var (
	latestCallID      uint64
	latestCallIDMutex sync.Mutex
)

// startCall is called at the beginning of a contract call to create a new frame in iteratorFrames.
// It updates latestCallID for generating a new call ID.
func startCall() uint64 { _ = "STUB: not implemented"; return 0 }

// removeFrame removes the frame with for the given call ID.
// The result can be nil when the frame is not initialized,
// i.e. when startCall() is called but no iterator is stored.
func removeFrame(callID uint64) frame { _ = "STUB: not implemented"; return *new(frame) }

// endCall is called at the end of a contract call to remove one item the iteratorFrames
func endCall(callID uint64) {
	_ = "STUB: not implemented"
	// we pull removeFrame in another function so we don't hold the mutex while cleaning up the removed frame
	return
}

// free all iterators in the frame when we release it

// storeIterator will add this to the end of the frame for the given ID and return a reference to it.
// We start counting with 1, so the 0 value is flagged as an error. This means we must
// remember to do idx-1 when retrieving
func storeIterator(callID uint64, it types.Iterator, frameLenLimit int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// store at array position `oldFrameLen`

// #nosec G115 -- newIndex is always positive (oldFrameLen >= 0, so newIndex >= 1)

// retrieveIterator will recover an iterator based on index. This ensures it will not be garbage collected.
// We start counting with 1, in storeIterator so the 0 value is flagged as an error. This means we must
// remember to do idx-1 when retrieving
func retrieveIterator(callID uint64, index uint64) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// Check bounds before converting to int to avoid integer overflow
