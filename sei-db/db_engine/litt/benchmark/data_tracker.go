package benchmark

import (
	"context"
	"math/rand"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/benchmark/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// WriteInfo contains information needed to perform a write operation.
type WriteInfo struct {
	// The index of the key to write.
	KeyIndex uint64
	// The key to write.
	Key []byte
	// The value to write.
	Value []byte
}

// ReadInfo contains information needed to perform a read operation.
type ReadInfo struct {
	// The key to read.
	Key []byte
	// The value we expect to read.
	Value []byte
}

// DataTracker is responsible for tracking key-value pairs that have been written to the database, and for generating
// new key-value pairs to be written.
type DataTracker struct {
	ctx    context.Context
	cancel context.CancelFunc

	// A source of randomness.
	rand *rand.Rand

	// The configuration for the benchmark.
	config *config.BenchmarkConfig

	// The directory where cohort files are stored.
	cohortDirectory string

	// A map from cohort index to information about the cohort.
	cohorts map[uint64]*Cohort

	// The cohort that is currently being used to generate keys for writing.
	activeCohort *Cohort

	// A set of cohorts that have been completely written to the database (i.e. cohorts that are safe to read).
	completeCohortSet map[uint64]struct{}

	// A set of keys passed to ReportWrite() that have not yet been fully processed.
	writtenKeysSet map[uint64]struct{}

	// The index of the oldest cohort being tracked.
	lowestCohortIndex uint64

	// The index of the newest cohort being tracked.
	highestCohortIndex uint64

	// Consider all key indices that have been generated this session (i.e. ignore keys indices generated prior to the
	// most recent restart). We want to find the highest key index that has been written to the database AND
	// where all lower key indices have also been written as well.
	highestWrittenKeyIndex int64

	// Consider all cohorts that have been generated this session (i.e. ignore cohorts generated prior to the most
	// recent restart). We want to find the highest cohort index that has been fully written to the database AND
	// where all cohorts with lower indices have also been written as well.
	highestWrittenCohortIndex int64

	// A channel containing keys-value pairs that are ready to be written.
	writeInfoChan chan *WriteInfo

	// A channel containing keys that are ready to be read.
	readInfoChan chan *ReadInfo

	// A channel containing information about keys that have been written to the database.
	writtenKeyIndicesChan chan uint64

	// Responsible for producing "random" data for key-value pairs.
	generator *DataGenerator

	// The TTL minus a safety margin. Cohorts are considered to be expired if keys in them are older than this.
	safeTTL time.Duration

	// The size of the values in bytes for new cohorts.
	valueSize uint64

	// This channel has capacity one and initially has one value in it. This value is drained when the DataTracker is
	// fully stopped. Other threads can use this to block until the DataTracker is fully stopped.
	closedChan chan struct{}

	// Used to handle fatal errors in the DataTracker.
	errorMonitor *util.ErrorMonitor
}

// NewDataTracker creates a new DataTracker instance, loading all relevant cohorts from disk.
func NewDataTracker(
	ctx context.Context,
	config *config.BenchmarkConfig,
	errorMonitor *util.ErrorMonitor,
) (*DataTracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the cohort directory if it doesn't exist.

// Gather the set of complete cohorts. These are the cohorts we can read from.

// Create an initial active cohort.

// Starting fresh, create a new cohort starting from key index 0.

// Will be drained when the DataTracker is closed.

//nolint:gosec // indices fit int64
//nolint:gosec // indices fit int64

// gatherCohorts loads cohorts from files on disk. The lowest/highest cohort indices are valid if and only if the
// cohorts map is not empty. If no cohorts are found, the lowest and highest cohort indices will be 0.
func gatherCohorts(cohortDirPath string) (
	lowestCohortIndex uint64,
	highestCohortIndex uint64,
	cohorts map[uint64]*Cohort,
	err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil

	// walk over files in path
	// for each file, check if it is a cohort file
	// if it is, load the cohort and add it to the map
	// if it is not, ignore it
}

// Delete any swap files discovered

// Special case, no cohorts found.

// LargestReadableValueSize returns the size of the largest value possible to read from the database,
// given current configuration. Considers both values previously written and stored
// (possibly with different configurations), and values that may be written in the future with the
// current configuration.
func (t *DataTracker) LargestReadableValueSize() uint64 { _ = "STUB: not implemented"; return 0 }

// GetWriteInfo returns information required to perform a write operation. It returns the key index (which is needed to
// call MarkHighestIndexWritten()), the key, and the value. Data is generated on background goroutines in order to
// make this method very fast. Will not block as long as data can be generated in the background fast enough.
// May return nil if the context is cancelled.
func (t *DataTracker) GetWriteInfo() *WriteInfo { _ = "STUB: not implemented"; return nil }

// ReportWrite is called when a key has been written to the database. This means that the key is now safe to be read.
func (t *DataTracker) ReportWrite(index uint64) { _ = "STUB: not implemented"; return }

// GetReadInfo returns information required to perform a read operation. Blocks until there is data eligible to be read.
func (t *DataTracker) GetReadInfo() *ReadInfo { _ = "STUB: not implemented"; return nil }

// GetReadInfoWithTimeout returns information required to perform a read operation. Waits the specified timeout for
// data to be eligible to be read. If no data is available within the time limit, returns nil.
func (t *DataTracker) GetReadInfoWithTimeout(timeout time.Duration) *ReadInfo {
	_ = "STUB: not implemented"
	return nil
}

// Close stops the key manager's background tasks.
func (t *DataTracker) Close() { _ = "STUB: not implemented"; return }

// dataGenerator is responsible for generating data in the background.
func (t *DataTracker) dataGenerator() { _ = "STUB: not implemented"; return }

// Edge case: when stared up for the first time, there won't be any values eligible to be read.
// We have to handle this in a special manner to prevent nil values from being inserted into
// the readInfoChan.

// track keys that have been written so that we can read them in the future

// prepare a value to be eventually written

// perform garbage collection on cohorts

// Standard case.

// track keys that have been written so that we can read them in the future

// prepare a value to be eventually written

// prepare a value to be eventually read

// perform garbage collection on cohorts

// handleWrittenKey handles a key that has been written to the database.
func (t *DataTracker) handleWrittenKey(keyIndex uint64) {
	_ = "STUB: not implemented"
	// Add key index to the set of written keys we are tracking.
	return
}

// Determine the highest key index written so far that also has all lower key indices written.

//nolint:gosec // index non-negative

// The next key has been written, mark it as such.
//nolint:gosec // index fits int64

// Once we find the first key that has not been written, we can stop checking.
// We want t.highestWrittenKeyIndex to be the highest key index that has been written
// without any gaps in the sequence.

// Determine the highest cohort index written so far that also has all lower cohorts written.

//nolint:gosec // index non-negative

// Don't ever mark the active cohort as complete.

//nolint:gosec // index fits int64
// We've found a cohort that has all keys written.
//nolint:gosec // index fits int64

// Once we find the first cohort that does not have all keys written, we can stop checking.

// generateNextWriteInfo generates the next write info to be placed into the writeInfoChan.
func (t *DataTracker) generateNextWriteInfo() *WriteInfo { _ = "STUB: not implemented"; return nil }

// generateNextReadInfo generates the next read info to be placed into the readInfoChan.
func (t *DataTracker) generateNextReadInfo() *ReadInfo { _ = "STUB: not implemented"; return nil }

// No cohorts are complete, so we can't read anything.

// map iteration is random in golang, so this will yield a random complete cohort.

// DoCohortGC performs garbage collection on the cohorts, removing cohorts with entries that are nearing expiration.
func (t *DataTracker) DoCohortGC() {
	_ = "STUB: not implemented"

	// Check all cohorts except for the active cohort (i.e. the one with index t.highestCohortIndex).
	return
}

// Stop once we find the first cohort that is not eligible for deletion.

// Edge case: we've been writing data slow enough that the active cohort has expired.
// Create a new active cohort.
