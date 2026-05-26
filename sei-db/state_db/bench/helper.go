package bench

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/snapshots"
	commonevm "github.com/sei-protocol/sei-chain/sei-db/common/keys"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/bench/wrappers"
	sctypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
)

const (
	// EVMStoreName simulates the EVM module store
	EVMStoreName = commonevm.EVMStoreKey

	// KeySize EVM storage key: 0x03 prefix + 20-byte address + 32-byte slot = 53 bytes
	KeySize   = 53
	ValueSize = 32
)

// TestScenario bundles benchmark parameters and distribution.
type TestScenario struct {
	Name           string
	TotalKeys      int64
	NumBlocks      int64
	DuplicateRatio float64 // 0.0 = all inserts, 1.0 = all updates
	// The database backend to use for the benchmark.
	Backend      wrappers.DBType
	Distribution KeyDistribution

	// SnapshotPath, when set, points to a state sync snapshot chunks directory
	// (e.g. "<node_home>/data/snapshots/<height>/<format>/") containing numbered
	// chunk files (0, 1, 2, ...). Before the benchmark begins, the snapshot is
	// imported into the database via the native Committer.Importer path as a
	// preparation stage.
	SnapshotPath string
}

// KeyDistribution defines how many keys to generate per block.
type KeyDistribution func(numBlocks, totalKeys, block int64) int64

// EvenDistribution generates same number of keys on each block.
func EvenDistribution(numBlocks, totalKeys, _ int64) int64 { _ = "STUB: not implemented"; return 0 }

// BurstyDistribution emits periodic bursts with optional jitter.
// Example: base=100 keys/block, burstEvery=5, burstMultiplier=3 =>
// blocks 0,5,10... emit 300 keys; other blocks emit 100 keys (then +/- jitter).
func BurstyDistribution(seed int64, burstEvery, burstMultiplier, maxJitter int64) KeyDistribution {
	_ = "STUB: not implemented"
	return *new(KeyDistribution)
}

// NormalDistribution samples keys per block from a normal distribution.
// Example: totalKeys=1000, numBlocks=10, stddevFactor=0.2 =>
// mean=100 keys, stddev=20 keys
func NormalDistribution(seed int64, stddevFactor float64) KeyDistribution {
	_ = "STUB: not implemented"
	return *new(KeyDistribution)
}

// RampDistribution linearly ramps keysPerBlock by a factor over the run.
// Example: totalKeys=1000, numBlocks=10, startFactor=0.5, endFactor=1.5 =>
// per-block base=100; block 0 ~50 keys, block 9 ~150 keys (linearly interpolated).
func RampDistribution(startFactor, endFactor float64) KeyDistribution {
	_ = "STUB: not implemented"
	return *new(KeyDistribution)
}

// ProgressReporter reports benchmark progress periodically.
type ProgressReporter struct {
	totalKeys   int64
	totalBlocks int64
	keysWritten atomic.Int64
	startTime   time.Time
	done        chan struct{}
	interval    time.Duration
}

// NewProgressReporter creates a new progress reporter.
func NewProgressReporter(totalKeys, totalBlocks int64, interval time.Duration) *ProgressReporter {
	_ = "STUB: not implemented"
	return nil
}

// Start begins periodic progress reporting in a background goroutine.
func (p *ProgressReporter) Start() { _ = "STUB: not implemented"; return }

// Stop stops the progress reporter and prints final stats.
func (p *ProgressReporter) Stop() { _ = "STUB: not implemented"; return }

// Add records that keys were written.
func (p *ProgressReporter) Add(keys int) { _ = "STUB: not implemented"; return }

func (p *ProgressReporter) report() { _ = "STUB: not implemented"; return }

// startChangesetGenerator streams per-block changesets based on the scenario distribution.
func startChangesetGenerator(scenario TestScenario) <-chan *proto.NamedChangeSet {
	_ = "STUB: not implemented"
	return nil
}

func keyFromIndex(index int64) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec // index validated non-negative above

//nolint:gosec

// parseSnapshotHeight extracts the block height from a state sync snapshot
// chunks directory path. The expected layout is <snapshots>/<height>/<format>/,
// so the height is the parent of the format directory.
func parseSnapshotHeight(chunksDir string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// openSnapshotStream opens the numbered chunk files in chunksDir and returns a
// StreamReader that decompresses and demuxes the protobuf item stream.
func openSnapshotStream(chunksDir string) (*snapshots.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// importSnapshot reads a state sync snapshot from chunksDir and feeds every
// item through the given Importer (AddModule / AddNode). This is the same
// import path used by the real state sync restore logic.
// Returns the total number of leaf keys imported.
func importSnapshot(chunksDir string, importer sctypes.Importer) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

// runBenchmark runs the benchmark with optional progress reporting.
// If withProgress is true, reports keys/sec every 5 seconds to stdout.
func runBenchmark(b *testing.B, scenario TestScenario, withProgress bool) {
	_ = "STUB: not implemented"
	return
}

// Load snapshot if available

// close to make sure all data got flushed
