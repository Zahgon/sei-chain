package benchmark

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/benchmark/config"
)

// metrics is a struct that holds various performance metrics for the benchmark. If configured, periodically
// writes a summary to the log. The intention is to expose data about the benchmark's performance even if
// prometheus is not available or configured.
type metrics struct {
	ctx    context.Context
	logger *slog.Logger

	// The configuration for the benchmark.
	config *config.BenchmarkConfig

	// The time when the benchmark started.
	startTime time.Time

	// The number of bytes written since the benchmark started.
	bytesWritten atomic.Uint64

	// The number of bytes read since the benchmark started.
	bytesRead atomic.Uint64

	// The number of write operations performed since the benchmark started.
	writeCount atomic.Uint64

	// The number of read operations performed since the benchmark started.
	readCount atomic.Uint64

	// The number of flush operations performed since the benchmark started.
	flushCount atomic.Uint64

	// The amount of time spent writing data.
	nanosecondsSpentWriting atomic.Uint64

	// The amount of time spent reading data.
	nanosecondsSpentReading atomic.Uint64

	// The amount of time spent flushing data.
	nanosecondsSpentFlushing atomic.Uint64

	// Longest write duration observed.
	longestWriteDuration atomic.Uint64

	// Longest read duration observed.
	longestReadDuration atomic.Uint64

	// Longest flush duration observed.
	longestFlushDuration atomic.Uint64
}

// newMetrics initializes a new metrics object.
func newMetrics(
	ctx context.Context,
	logger *slog.Logger,
	config *config.BenchmarkConfig,
) *metrics {
	_ = "STUB: not implemented"
	return nil
}

// reportWrite records a write operation.
func (m *metrics) reportWrite(writeDuration time.Duration, bytesWritten uint64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // duration non-negative

// Update the longest write duration if this one is longer.

//nolint:gosec // durations comfortably fit
//nolint:gosec // duration non-negative

// reportRead records a read operation.
func (m *metrics) reportRead(readDuration time.Duration, bytesRead uint64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // duration non-negative

// Update the longest read duration if this one is longer.

//nolint:gosec // durations comfortably fit
//nolint:gosec // duration non-negative

// reportFlush records a flush operation.
func (m *metrics) reportFlush(flushDuration time.Duration) { _ = "STUB: not implemented"; return }

//nolint:gosec // duration non-negative

// Update the longest flush duration if this one is longer.

//nolint:gosec // durations comfortably fit
//nolint:gosec // duration non-negative

// reportGenerator runs in a goroutine and periodically logs the metrics to the console.
func (m *metrics) reportGenerator() { _ = "STUB: not implemented"; return }

// Metrics logging is disabled.

// Context cancelled, stop reporting.

// logMetrics logs the current metrics to the console.
func (m *metrics) logMetrics() { _ = "STUB: not implemented"; return }

//nolint:gosec // duration non-negative

//nolint:gosec // duration non-negative

//nolint:gosec // duration non-negative

//nolint:gosec // duration non-negative
