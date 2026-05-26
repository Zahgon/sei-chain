package disktable

import (
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/disktable/keymap"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/disktable/segment"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// controlLoop runs a goroutine that handles control messages for the disk table.
type controlLoop struct {
	logger *slog.Logger

	// diskTable is the disk table that this control loop is associated with.
	diskTable *DiskTable

	// errorMonitor is used to react to fatal errors anywhere in the disk table.
	errorMonitor *util.ErrorMonitor

	// controllerChannel is the channel for messages sent to the control loop.
	controllerChannel chan any

	// The index of the lowest numbered segment. After initial creation, only the garbage collection
	// thread is permitted to read/write this value for the sake of thread safety.
	lowestSegmentIndex uint32

	// The index of the highest numbered segment. All writes are applied to this segment.
	highestSegmentIndex uint32

	// This value mirrors highestSegmentIndex, but is thread safe to read from external goroutines.
	// There are several unit tests that read this value, and so there needs to be a threadsafe way
	// to access it. Since new segments are added on an infrequent basis and this is never read in
	// production, maintaining this atomic variable has negligible overhead.
	threadsafeHighestSegmentIndex atomic.Uint32

	// segmentLock protects access to the variables segments and highestSegmentIndex.
	// Does not protect the segments themselves.
	segmentLock sync.RWMutex

	// All segments currently in use. Only the control loop modifies this map, but other threads may read from it.
	// The control loop does not need to hold a lock when doing read operations on this map, since no other thread
	// will modify it. The control loop does need to hold a lock when modifying this map, though, and other threads
	// must hold a lock when reading from it.
	segments map[uint32]*segment.Segment

	// The number of bytes contained within the immutable segments. This tracks the number of bytes that are
	// on disk, not bytes in memory. For thread safety, this variable may only be read/written in the constructor
	// and in the control loop.
	immutableSegmentSize uint64

	// The target size for value files.
	targetFileSize uint32

	// The maximum number of keys in a segment.
	maxKeyCount uint32

	// The target size for key files.
	targetKeyFileSize uint64

	// The size of the disk table is stored here.
	size *atomic.Uint64

	// The number of keys in the table.
	keyCount *atomic.Int64

	// clock is the time source used by the disk table.
	clock func() time.Time

	// The locations where segment files are stored.
	segmentPaths []*segment.SegmentPath

	// Controls if snapshotting is enabled or not.
	snapshottingEnabled bool

	// The table's metadata.
	metadata *tableMetadata

	// whether fsync mode is enabled.
	fsync bool

	// If true, then the control loop has been stopped.
	stopped atomic.Bool

	// Encapsulates metrics for the database.
	metrics *metrics.LittDBMetrics

	// The table's name.
	name string

	// The maximum number of keys that can be garbage collected in a single batch.
	gcBatchSize uint64

	// The keymap used to store key-to-address mappings.
	keymap keymap.Keymap

	// The goroutine responsible for blocking on flush operations.
	flushLoop *flushLoop

	// garbageCollectionPeriod is the period at which garbage collection is run.
	garbageCollectionPeriod time.Duration
}

// enqueue enqueues a request to the control loop. Returns an error if the request could not be sent due to the
// database being in a panicked state. Only types defined in control_loop_messages.go are permitted to be sent
// to the control loop.
func (c *controlLoop) enqueue(request controlLoopMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// run runs the control loop for the disk table. It has sole responsibility for scheduling all operations that
// mutate the data in the disk table.
func (c *controlLoop) run() { _ = "STUB: not implemented"; return }

// doGarbageCollection performs garbage collection on all segments, deleting old ones as necessary.
func (c *controlLoop) doGarbageCollection() { _ = "STUB: not implemented"; return }

// No TTL set, so nothing to do.

// We can't delete an unsealed segment.

// Segment is not old enough to be deleted.

// Segment is old enough to be deleted.

// Deletion of segment files will happen when the segment is released by all reservation holders.

// getReservedSegment returns the segment with the given index. Segment is reserved, and it is the caller's
// responsibility to release the reservation when done. Returns true if the segment was found and reserved,
// and false if the segment could not be found or could not be reserved.
func (c *controlLoop) getReservedSegment(index uint32) (*segment.Segment, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// segmented was deleted out from under us

// getSegments returns the segments of the disk table. It is only legal to call this after the control loop has been
// stopped.
func (c *controlLoop) getSegments() (map[uint32]*segment.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// updateCurrentSize updates the size of the table.
func (c *controlLoop) updateCurrentSize() { _ = "STUB: not implemented"; return }

// handleWriteRequest handles a controlLoopWriteRequest control message.
func (c *controlLoop) handleWriteRequest(req *controlLoopWriteRequest) {
	_ = "STUB: not implemented"
	return
}

// Do the write.

// Check to see if the write caused the mutable segment to become full.

// Mutable segment is full. Before continuing, we need to expand the segments.

// expandSegments seals the latest segment and creates a new mutable segment.
func (c *controlLoop) expandSegments() error {
	_ = "STUB: not implemented"

	// Seal the previous segment.
	return nil
}

// Unfortunately, it is necessary to block until the sealing has been completed. Although this may result
// in a brief interruption in new write work being sent to the segment, expanding the number of segments is
// infrequent, even for very high throughput workloads.

// Record the size of the segment.

// Create a new segment.

// handleFlushRequest handles the part of the flush that is performed on the control loop.
// The control loop is responsible for enqueuing the flush request in the segment's work queue (thus
// ensuring a serial ordering with respect to other operations on the control loop), but not for
// waiting for the segment to finish the flush.
func (c *controlLoop) handleFlushRequest(req *controlLoopFlushRequest) {
	_ = "STUB: not implemented"
	// This method will enqueue a flush operation within the segment. Once that is done,
	// it becomes the responsibility of the flush loop to wait for the flush to complete.
	return
}

// The flush loop is responsible for the remaining parts of the flush.

// handleControlLoopSetShardingFactorRequest updates the sharding factor of the disk table. If the requested
// sharding factor is the same as before, no action is taken. If it is different, the sharding factor is updated,
// the current mutable segment is sealed, and a new mutable segment is created.
func (c *controlLoop) handleControlLoopSetShardingFactorRequest(req *controlLoopSetShardingFactorRequest) {
	_ = "STUB: not implemented"
	return
}

// No action necessary.

// This seals the current mutable segment and creates a new one. The new segment will have the new sharding factor.

// handleShutdownRequest performs tasks necessary to cleanly shut down the disk table.
func (c *controlLoop) handleShutdownRequest(req *controlLoopShutdownRequest) {
	_ = "STUB: not implemented"
	// Instruct the flush loop to stop.
	return
}

// Seal the mutable segment

// Flush the keys that are now durable in the segment.

// Stop the keymap
