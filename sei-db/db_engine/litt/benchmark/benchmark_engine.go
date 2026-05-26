package benchmark

import (
	"context"
	"log/slog"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/benchmark/config"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// BenchmarkEngine is a tool for benchmarking LittDB performance.
type BenchmarkEngine struct {
	ctx    context.Context
	cancel context.CancelFunc
	logger *slog.Logger

	// The configuration for the benchmark.
	config *config.BenchmarkConfig

	// The database to be benchmarked.
	db litt.DB

	// The table in the database where data is stored.
	table litt.Table

	// Keeps track of data to read and write.
	dataTracker *DataTracker

	// The maximum write throughput in bytes per second for each worker thread.
	writeBytesPerSecondPerThread uint64

	// The maximum read throughput in bytes per second for each worker thread.
	readBytesPerSecondPerThread uint64

	// The burst size for write rate limiting.
	writeBurstSize uint64

	// The burst size for read rate limiting.
	readBurstSize uint64

	// Records benchmark metrics.
	metrics *metrics

	// errorMonitor is used to handle fatal errors in the benchmark engine.
	errorMonitor *util.ErrorMonitor
}

// NewBenchmarkEngine creates a new BenchmarkEngine with the given configuration.
func NewBenchmarkEngine(configPath string) (*BenchmarkEngine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // bounded above by MaxShardingFactor

//nolint:gosec // parallelism positive

// If we set the write burst size smaller than an individual value, then the rate limiter will never
// permit any writes. Ideally, we'd just set the burst size to 0 since we don't want bursty/volatile writes,
// but since we are using the rate.Limiter utility, we are required to set a burst size, and a burst size
// smaller than an individual value will cause the rate limiter to never permit writes.

//nolint:gosec // parallelism positive

// If we set the read burst size smaller than an individual value we need to read, then the rate limiter will
// never permit us to read that value.

// Logger returns the logger used by the benchmark engine.
func (b *BenchmarkEngine) Logger() *slog.Logger {
	_ = "STUB: not implemented"

	// Run executes the benchmark. This method blocks forever, or until the benchmark is stopped via control-C or
	// encounters an error.
	return nil
}

func (b *BenchmarkEngine) Run() error { _ = "STUB: not implemented"; return nil }

// If a time limit is set, create a timer to cancel the context after the specified duration

// multiply by 2 to make configured value the average

// Sleep a short time to prevent all goroutines from starting in lockstep.

// Sleep a short time to prevent all goroutines from starting in lockstep.

// Create a channel to listen for OS signals

// Wait for signal

// Cancel the context when signal is received

// writer runs on a goroutine and writes data to the database.
func (b *BenchmarkEngine) writer() { _ = "STUB: not implemented"; return }

//nolint:gosec // burst sized by config

// verifyValue checks if the actual value read from the database matches the expected value.
func (b *BenchmarkEngine) verifyValue(expected *ReadInfo, actual []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// reader runs on a goroutine and reads data from the database.
func (b *BenchmarkEngine) reader() { _ = "STUB: not implemented"; return }

//nolint:gosec // burst sized by config

// This can happen when the context gets cancelled.
