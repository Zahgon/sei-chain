package disktable

import (
	"log/slog"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// flushLoop is a struct that runs a goroutine that is responsible for blocking on flush operations.
type flushLoop struct {
	logger *slog.Logger

	// the parent disk table
	diskTable *DiskTable

	// Responsible for handling fatal DB errors.
	errorMonitor *util.ErrorMonitor

	// flushChannel is a channel used to enqueue work on the flush loop.
	flushChannel chan any

	// metrics encapsulates metrics for the DB.
	metrics *metrics.LittDBMetrics

	// provides the current time
	clock func() time.Time

	// the name of the table
	name string

	// This file stores the highest segment index that is fully snapshot. Written as segments are sealed
	// and copied to the snapshot directory, read by the external snapshot consumer.
	upperBoundSnapshotFile *BoundaryFile
}

// enqueue sends work to be handled on the flush loop. Will return an error if the DB is panicking.
func (f *flushLoop) enqueue(request flushLoopMessage) error { _ = "STUB: not implemented"; return nil }

// run is responsible for handling operations that flush data (i.e. calls to Flush() and when the mutable segment
// is sealed). In theory, this work could be done on the main control loop, but doing so would block new writes while
// a flush is in progress. In order to keep the writing threads busy, it is critical that flush do not block the
// control loop.
func (f *flushLoop) run() { _ = "STUB: not implemented"; return }

// handleSealRequest handles the part of the seal operation that is performed on the flush loop.
// We don't want to send a flush request to a segment that has already been sealed. By performing the sealing
// on the flush loop, we ensure that this can never happen. Any previously scheduled flush requests against the
// segment that is being sealed will be processed prior to this request being processed due to the FIFO nature
// of the flush loop channel. When a sealing operation begins, the control loop blocks, and does not unblock until
// the seal is finished and a new mutable segment has been created. This means that no future flush requests will be
// sent to the segment that is being sealed, since only the control loop can schedule work for the flush loop.
func (f *flushLoop) handleSealRequest(req *flushLoopSealRequest) { _ = "STUB: not implemented"; return }

// Flush the keys that are now durable in the segment.

// Snapshotting can wait until after we have sent a response. No need for the Flush() caller to wait for
// snapshotting. Flush() only cares about the data's crash durability, and is completely independent of
// snapshotting.

// Update the boundary file. The consumer of the snapshot uses this information to determine when segments
// are fully copied to the snapshot directory.

// handleFlushRequest handles the part of the flush that is performed on the flush loop.
func (f *flushLoop) handleFlushRequest(req *flushLoopFlushRequest) {
	_ = "STUB: not implemented"
	return
}
