package blocksim

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/block"
	"golang.org/x/time/rate"
)

// The benchmark runner for the blocksim benchmark.
type BlockSim struct {
	ctx    context.Context
	cancel context.CancelFunc

	config *BlocksimConfig

	db        block.BlockDB
	generator *BlockGenerator
	metrics   *BlocksimMetrics

	// Console reporting state.
	consoleUpdatePeriod         time.Duration
	lastConsoleUpdateTime       time.Time
	lastConsoleUpdateBlockCount int64
	startTimestamp              time.Time
	totalBlocksWritten          int64
	totalTransactionsWritten    int64
	totalBytesWritten           int64
	highestBlockHeight          uint64

	// A message is sent on this channel when the benchmark is fully stopped.
	closeChan chan struct{}

	// Suspend/resume toggle channel.
	suspendChan chan bool

	// Enforces a maximum block write rate (if enabled).
	rateLimiter *rate.Limiter
}

// Creates a new blocksim benchmark runner.
func NewBlockSim(
	ctx context.Context,
	config *BlocksimConfig,
	metrics *BlocksimMetrics,
) (*BlockSim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

// The main loop of the benchmark.
func (b *BlockSim) run() { _ = "STUB: not implemented"; return }

func (b *BlockSim) maybeThrottle() { _ = "STUB: not implemented"; return }

func (b *BlockSim) handleNextBlock(blk *block.BinaryBlock) { _ = "STUB: not implemented"; return }

// Periodic flush.
//nolint:gosec

// Periodic prune.

func (b *BlockSim) suspend() {
	_ = "STUB: not implemented"
	// Flush before suspending so state is durable.
	return
}

// Reset console metrics on resume.

func (b *BlockSim) teardown() { _ = "STUB: not implemented"; return }

func (b *BlockSim) generateConsoleReport(force bool) { _ = "STUB: not implemented"; return }

//nolint:gosec

// Blocks until the benchmark has halted.
func (b *BlockSim) BlockUntilHalted() { _ = "STUB: not implemented"; return }

// Close shuts down the benchmark and releases resources.
func (b *BlockSim) Close() error { _ = "STUB: not implemented"; return nil }

// Suspend the benchmark. Call Resume() to continue.
func (b *BlockSim) Suspend() { _ = "STUB: not implemented"; return }

// Resume the benchmark after a Suspend().
func (b *BlockSim) Resume() { _ = "STUB: not implemented"; return }

// openBlockDB creates a BlockDB for the given backend name.
func openBlockDB(backend string, dataDir string) (block.BlockDB, error) {
	_ = "STUB: not implemented"
	return *new(block.BlockDB), nil
}

// removeContents deletes all entries inside dir without removing dir itself.
func removeContents(dir string) error { _ = "STUB: not implemented"; return nil }
