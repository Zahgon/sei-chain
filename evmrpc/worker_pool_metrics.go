package evmrpc

import (
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// Environment variable to enable debug metrics printing to stdout
// Set EVM_DEBUG_METRICS=true to enable periodic metrics printing
const EVMDebugMetricsEnvVar = "EVM_DEBUG_METRICS"

// IsDebugMetricsEnabled checks if debug metrics printing is enabled via environment variable
func IsDebugMetricsEnabled() bool { _ = "STUB: not implemented"; return false }

// Error type constants for categorization
const (
	ErrTypeRangeTooLarge = "range_too_large"
	ErrTypeTimeout       = "timeout"
	ErrTypeRateLimited   = "rate_limited"
	ErrTypeBackpressure  = "backpressure"
	ErrTypeIOSaturated   = "io_saturated"
	ErrTypeBlockNotFound = "block_not_found"
	ErrTypeQueueFull     = "queue_full"
	ErrTypeOther         = "other"
)

// WorkerPoolMetrics tracks worker pool performance metrics
type WorkerPoolMetrics struct {
	// Worker pool stats
	TotalWorkers    atomic.Int32
	ActiveWorkers   atomic.Int32
	QueueCapacity   atomic.Int32
	QueueDepth      atomic.Int32
	PeakQueueDepth  atomic.Int32
	TasksSubmitted  atomic.Int64
	TasksCompleted  atomic.Int64
	TasksRejected   atomic.Int64 // Queue full rejections
	TasksPanicked   atomic.Int64
	TotalWaitTimeNs atomic.Int64 // Total time tasks spent waiting in queue
	TotalExecTimeNs atomic.Int64 // Total task execution time

	// DB Semaphore stats
	DBSemaphoreCapacity   atomic.Int32
	DBSemaphoreAcquired   atomic.Int32
	DBSemaphoreWaitTimeNs atomic.Int64
	DBSemaphoreWaitCount  atomic.Int64

	// eth_getLogs specific stats
	GetLogsRequests      atomic.Int64
	GetLogsErrors        atomic.Int64
	GetLogsSuccess       atomic.Int64 // Successful requests
	GetLogsBlockRangeSum atomic.Int64 // Sum of block ranges for average calculation
	GetLogsLatencySumNs  atomic.Int64 // Sum of latencies for average calculation
	GetLogsPeakRange     atomic.Int64
	GetLogsMaxLatencyNs  atomic.Int64 // Max latency observed

	// Error type breakdown
	ErrRangeTooLarge atomic.Int64
	ErrTimeout       atomic.Int64
	ErrRateLimited   atomic.Int64
	ErrBackpressure  atomic.Int64
	ErrIOSaturated   atomic.Int64
	ErrBlockNotFound atomic.Int64
	ErrQueueFull     atomic.Int64
	ErrOther         atomic.Int64

	// Block range distribution buckets (total requests)
	RangeBucket1to10      atomic.Int64 // 1-10 blocks
	RangeBucket11to100    atomic.Int64 // 11-100 blocks
	RangeBucket101to500   atomic.Int64 // 101-500 blocks
	RangeBucket501to1000  atomic.Int64 // 501-1000 blocks
	RangeBucket1001to2000 atomic.Int64 // 1001-2000 blocks
	RangeBucketOver2000   atomic.Int64 // >2000 blocks

	// Block range success counts (for calculating success rate per bucket)
	RangeBucket1to10Success      atomic.Int64
	RangeBucket11to100Success    atomic.Int64
	RangeBucket101to500Success   atomic.Int64
	RangeBucket501to1000Success  atomic.Int64
	RangeBucket1001to2000Success atomic.Int64
	RangeBucketOver2000Success   atomic.Int64

	// Subscription stats
	ActiveSubscriptions atomic.Int32
	SubscriptionErrors  atomic.Int64

	// Time window for TPS calculation
	windowStart    time.Time
	windowRequests atomic.Int64
	mu             sync.RWMutex
}

var (
	metricsPrinterOnce sync.Once
	metricsStopOnce    sync.Once
	metricsStopChan    chan struct{}
)

var (
	meter = otel.Meter("evmrpc_workerpool")

	otelMetrics = struct {
		workersTotal          metric.Int64Gauge
		workersActive         metric.Int64Gauge
		workersIdle           metric.Int64Gauge
		queueCapacity         metric.Int64Gauge
		queueDepth            metric.Int64Gauge
		queuePeak             metric.Int64Gauge
		queueUtilizationPct   metric.Float64Gauge
		tasksSubmittedTotal   metric.Int64Gauge
		tasksCompletedTotal   metric.Int64Gauge
		tasksRejectedTotal    metric.Int64Gauge
		tasksPanickedTotal    metric.Int64Gauge
		dbSemaphoreCapacity   metric.Int64Gauge
		dbSemaphoreInUse      metric.Int64Gauge
		dbSemaphoreAvailable  metric.Int64Gauge
		dbSemaphoreWaitCount  metric.Int64Gauge
		subscriptionsActive   metric.Int64Gauge
		getLogsRequestsTotal  metric.Int64Gauge
		getLogsSuccessTotal   metric.Int64Gauge
		getLogsErrorsTotal    metric.Int64Gauge
		getLogsTPS            metric.Float64Gauge
		getLogsAvgBlockRange  metric.Float64Gauge
		getLogsPeakBlockRange metric.Int64Gauge
		getLogsAvgLatencyMs   metric.Float64Gauge
		getLogsMaxLatencyMs   metric.Float64Gauge
		errRangeTooLarge      metric.Int64Gauge
		errRateLimited        metric.Int64Gauge
		errBackpressure       metric.Int64Gauge
		avgQueueWaitMs        metric.Float64Gauge
		avgExecTimeMs         metric.Float64Gauge
		avgDBWaitMs           metric.Float64Gauge
	}{
		workersTotal: must(meter.Int64Gauge(
			"evmrpc_workerpool_workers_total",
			metric.WithDescription("Total worker count"),
			metric.WithUnit("{count}"),
		)),
		workersActive: must(meter.Int64Gauge(
			"evmrpc_workerpool_workers_active",
			metric.WithDescription("Active worker count"),
			metric.WithUnit("{count}"),
		)),
		workersIdle: must(meter.Int64Gauge(
			"evmrpc_workerpool_workers_idle",
			metric.WithDescription("Idle worker count"),
			metric.WithUnit("{count}"),
		)),
		queueCapacity: must(meter.Int64Gauge(
			"evmrpc_workerpool_queue_capacity",
			metric.WithDescription("Task queue capacity"),
			metric.WithUnit("{count}"),
		)),
		queueDepth: must(meter.Int64Gauge(
			"evmrpc_workerpool_queue_depth",
			metric.WithDescription("Current task queue depth"),
			metric.WithUnit("{count}"),
		)),
		queuePeak: must(meter.Int64Gauge(
			"evmrpc_workerpool_queue_peak",
			metric.WithDescription("Peak queue depth observed"),
			metric.WithUnit("{count}"),
		)),
		queueUtilizationPct: must(meter.Float64Gauge(
			"evmrpc_workerpool_queue_utilization_pct",
			metric.WithDescription("Queue utilization percentage"),
			metric.WithUnit("1"),
		)),
		tasksSubmittedTotal: must(meter.Int64Gauge(
			"evmrpc_workerpool_tasks_submitted_total",
			metric.WithDescription("Tasks submitted"),
			metric.WithUnit("{count}"),
		)),
		tasksCompletedTotal: must(meter.Int64Gauge(
			"evmrpc_workerpool_tasks_completed_total",
			metric.WithDescription("Tasks completed"),
			metric.WithUnit("{count}"),
		)),
		tasksRejectedTotal: must(meter.Int64Gauge(
			"evmrpc_workerpool_tasks_rejected_total",
			metric.WithDescription("Tasks rejected due to full queue"),
			metric.WithUnit("{count}"),
		)),
		tasksPanickedTotal: must(meter.Int64Gauge(
			"evmrpc_workerpool_tasks_panicked_total",
			metric.WithDescription("Tasks that panicked"),
			metric.WithUnit("{count}"),
		)),
		dbSemaphoreCapacity: must(meter.Int64Gauge(
			"evmrpc_db_semaphore_capacity",
			metric.WithDescription("DB semaphore capacity"),
			metric.WithUnit("{count}"),
		)),
		dbSemaphoreInUse: must(meter.Int64Gauge(
			"evmrpc_db_semaphore_inuse",
			metric.WithDescription("DB semaphore currently acquired"),
			metric.WithUnit("{count}"),
		)),
		dbSemaphoreAvailable: must(meter.Int64Gauge(
			"evmrpc_db_semaphore_available",
			metric.WithDescription("DB semaphore available slots"),
			metric.WithUnit("{count}"),
		)),
		dbSemaphoreWaitCount: must(meter.Int64Gauge(
			"evmrpc_db_semaphore_wait_count",
			metric.WithDescription("DB semaphore wait count"),
			metric.WithUnit("{count}"),
		)),
		subscriptionsActive: must(meter.Int64Gauge(
			"evmrpc_subscriptions_active",
			metric.WithDescription("Active subscriptions"),
			metric.WithUnit("{count}"),
		)),
		getLogsRequestsTotal: must(meter.Int64Gauge(
			"evmrpc_getlogs_requests_total",
			metric.WithDescription("Total eth_getLogs requests"),
			metric.WithUnit("{count}"),
		)),
		getLogsSuccessTotal: must(meter.Int64Gauge(
			"evmrpc_getlogs_success_total",
			metric.WithDescription("Successful eth_getLogs requests"),
			metric.WithUnit("{count}"),
		)),
		getLogsErrorsTotal: must(meter.Int64Gauge(
			"evmrpc_getlogs_errors_total",
			metric.WithDescription("Errored eth_getLogs requests"),
			metric.WithUnit("{count}"),
		)),
		getLogsTPS: must(meter.Float64Gauge(
			"evmrpc_getlogs_tps",
			metric.WithDescription("eth_getLogs throughput (req/s)"),
			metric.WithUnit("1/s"),
		)),
		getLogsAvgBlockRange: must(meter.Float64Gauge(
			"evmrpc_getlogs_avg_blockrange",
			metric.WithDescription("Average block range for eth_getLogs"),
			metric.WithUnit("{blocks}"),
		)),
		getLogsPeakBlockRange: must(meter.Int64Gauge(
			"evmrpc_getlogs_peak_blockrange",
			metric.WithDescription("Peak block range for eth_getLogs"),
			metric.WithUnit("{blocks}"),
		)),
		getLogsAvgLatencyMs: must(meter.Float64Gauge(
			"evmrpc_getlogs_avg_latency_ms",
			metric.WithDescription("Average eth_getLogs latency (ms)"),
			metric.WithUnit("ms"),
		)),
		getLogsMaxLatencyMs: must(meter.Float64Gauge(
			"evmrpc_getlogs_max_latency_ms",
			metric.WithDescription("Max eth_getLogs latency (ms)"),
			metric.WithUnit("ms"),
		)),
		errRangeTooLarge: must(meter.Int64Gauge(
			"evmrpc_getlogs_errors_range_too_large",
			metric.WithDescription("Errors due to block range too large"),
			metric.WithUnit("{count}"),
		)),
		errRateLimited: must(meter.Int64Gauge(
			"evmrpc_getlogs_errors_rate_limited",
			metric.WithDescription("Errors due to rate limiting"),
			metric.WithUnit("{count}"),
		)),
		errBackpressure: must(meter.Int64Gauge(
			"evmrpc_getlogs_errors_backpressure",
			metric.WithDescription("Errors due to backpressure"),
			metric.WithUnit("{count}"),
		)),
		avgQueueWaitMs: must(meter.Float64Gauge(
			"evmrpc_workerpool_avg_queue_wait_ms",
			metric.WithDescription("Average queue wait time (ms)"),
			metric.WithUnit("ms"),
		)),
		avgExecTimeMs: must(meter.Float64Gauge(
			"evmrpc_workerpool_avg_exec_time_ms",
			metric.WithDescription("Average execution time (ms)"),
			metric.WithUnit("ms"),
		)),
		avgDBWaitMs: must(meter.Float64Gauge(
			"evmrpc_db_semaphore_avg_wait_ms",
			metric.WithDescription("Average DB semaphore wait (ms)"),
			metric.WithUnit("ms"),
		)),
	}
)

// GetGlobalMetrics returns the metrics from the global worker pool
// This is a convenience function for accessing metrics without importing worker pool
func GetGlobalMetrics() *WorkerPoolMetrics { _ = "STUB: not implemented"; return nil }

// StartMetricsPrinter starts a background goroutine that prints metrics every interval
// This is idempotent - only the first call will start the printer
// Note: Printing to stdout is controlled by the EVM_DEBUG_METRICS environment variable
// Set EVM_DEBUG_METRICS=true to enable debug output
func StartMetricsPrinter(interval time.Duration) { _ = "STUB: not implemented"; return }

// Export to Prometheus (gauges need periodic update)

// Print to stdout only if debug is enabled

// StopMetricsPrinter stops the metrics printer, idempotent.
func StopMetricsPrinter() { _ = "STUB: not implemented"; return }

// RecordTaskSubmitted records a task submission
// Note: Prometheus export is done in batch via ExportPrometheusMetrics()
func (m *WorkerPoolMetrics) RecordTaskSubmitted() { _ = "STUB: not implemented"; return }

// Update peak if needed

// RecordTaskStarted records when a task starts executing
// Note: Prometheus export is done in batch via ExportPrometheusMetrics()
func (m *WorkerPoolMetrics) RecordTaskStarted(queuedAt time.Time) {
	_ = "STUB: not implemented"
	return
}

// RecordTaskCompleted records a task completion
// Note: Prometheus export is done in batch via ExportPrometheusMetrics()
func (m *WorkerPoolMetrics) RecordTaskCompleted(startedAt time.Time) {
	_ = "STUB: not implemented"
	return
}

// RecordTaskRejected records a task rejection (queue full)
// Note: Prometheus export is done in batch via ExportPrometheusMetrics()
func (m *WorkerPoolMetrics) RecordTaskRejected() { _ = "STUB: not implemented"; return }

// RecordTaskPanicked records a task panic
// Note: Prometheus export is done in batch via ExportPrometheusMetrics()
func (m *WorkerPoolMetrics) RecordTaskPanicked() { _ = "STUB: not implemented"; return }

// RecordDBSemaphoreAcquire records acquiring the DB semaphore
func (m *WorkerPoolMetrics) RecordDBSemaphoreAcquire() { _ = "STUB: not implemented"; return }

// RecordDBSemaphoreRelease records releasing the DB semaphore
func (m *WorkerPoolMetrics) RecordDBSemaphoreRelease() { _ = "STUB: not implemented"; return }

// RecordDBSemaphoreWait records time spent waiting for DB semaphore
// Note: Prometheus export is done in batch via ExportPrometheusMetrics()
func (m *WorkerPoolMetrics) RecordDBSemaphoreWait(waitTime time.Duration) {
	_ = "STUB: not implemented"
	return
}

// RecordGetLogsRequest records an eth_getLogs request with detailed error categorization
func (m *WorkerPoolMetrics) RecordGetLogsRequest(blockRange int64, latency time.Duration, startTime time.Time, err error) {
	_ = "STUB: not implemented"
	return
}

// Update max latency

// Update peak range

// Record block range distribution

// Categorize errors and record success per bucket

// Note: Prometheus export is done in batch via ExportPrometheusMetrics()

// recordBlockRangeBucket records the block range into the appropriate bucket
func (m *WorkerPoolMetrics) recordBlockRangeBucket(blockRange int64) {
	_ = "STUB: not implemented"
	return
}

// recordBlockRangeBucketSuccess records a successful request in the appropriate bucket
func (m *WorkerPoolMetrics) recordBlockRangeBucketSuccess(blockRange int64) {
	_ = "STUB: not implemented"
	return
}

// categorizeError categorizes an error into specific types
func (m *WorkerPoolMetrics) categorizeError(err error) { _ = "STUB: not implemented"; return }

// contains is a helper function for case-insensitive substring matching
func contains(s, substr string) bool { _ = "STUB: not implemented"; return false }

func toLower(s string) string { _ = "STUB: not implemented"; return "" }

func containsLower(s, substr string) bool { _ = "STUB: not implemented"; return false }

// RecordSubscriptionStart records a new subscription
func (m *WorkerPoolMetrics) RecordSubscriptionStart() { _ = "STUB: not implemented"; return }

// RecordSubscriptionEnd records subscription end
func (m *WorkerPoolMetrics) RecordSubscriptionEnd() { _ = "STUB: not implemented"; return }

// RecordSubscriptionError records a subscription error
func (m *WorkerPoolMetrics) RecordSubscriptionError() { _ = "STUB: not implemented"; return }

// Export to Prometheus

// GetTPS calculates the current TPS based on time window
func (m *WorkerPoolMetrics) GetTPS() float64 { _ = "STUB: not implemented"; return 0 }

// ResetTPSWindow resets the TPS calculation window
func (m *WorkerPoolMetrics) ResetTPSWindow() { _ = "STUB: not implemented"; return }

// GetAverageQueueWaitTime returns average time tasks spend waiting in queue
func (m *WorkerPoolMetrics) GetAverageQueueWaitTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetAverageExecTime returns average task execution time
func (m *WorkerPoolMetrics) GetAverageExecTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetAverageDBWaitTime returns average DB semaphore wait time
func (m *WorkerPoolMetrics) GetAverageDBWaitTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetAverageBlockRange returns average block range for eth_getLogs
func (m *WorkerPoolMetrics) GetAverageBlockRange() float64 { _ = "STUB: not implemented"; return 0 }

// GetAverageLatency returns average eth_getLogs latency
func (m *WorkerPoolMetrics) GetAverageLatency() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetSnapshot returns a snapshot of current metrics
func (m *WorkerPoolMetrics) GetSnapshot() MetricsSnapshot {
	_ = "STUB: not implemented"
	return *new(MetricsSnapshot)
}

// Worker pool

// DB Semaphore

// eth_getLogs

// Error type breakdown

// Block range distribution

// Subscriptions

// MetricsSnapshot represents a point-in-time snapshot of metrics
type MetricsSnapshot struct {
	Timestamp time.Time

	// Worker pool
	TotalWorkers     int32
	ActiveWorkers    int32
	IdleWorkers      int32
	QueueCapacity    int32
	QueueDepth       int32
	QueueUtilization float64
	PeakQueueDepth   int32
	TasksSubmitted   int64
	TasksCompleted   int64
	TasksRejected    int64
	TasksPending     int64
	AvgQueueWaitTime time.Duration
	AvgExecTime      time.Duration

	// DB Semaphore
	DBSemaphoreCapacity int32
	DBSemaphoreInUse    int32
	DBSemaphoreAvail    int32
	AvgDBWaitTime       time.Duration

	// eth_getLogs
	GetLogsTPS       float64
	GetLogsTotal     int64
	GetLogsSuccess   int64
	GetLogsErrors    int64
	GetLogsErrorRate float64
	AvgBlockRange    float64
	PeakBlockRange   int64
	AvgLatency       time.Duration
	MaxLatency       time.Duration

	// Error type breakdown
	ErrRangeTooLarge int64
	ErrTimeout       int64
	ErrRateLimited   int64
	ErrBackpressure  int64
	ErrIOSaturated   int64
	ErrBlockNotFound int64
	ErrQueueFull     int64
	ErrOther         int64

	// Block range distribution (total and success for calculating success rate)
	RangeBucket1to10             int64
	RangeBucket1to10Success      int64
	RangeBucket11to100           int64
	RangeBucket11to100Success    int64
	RangeBucket101to500          int64
	RangeBucket101to500Success   int64
	RangeBucket501to1000         int64
	RangeBucket501to1000Success  int64
	RangeBucket1001to2000        int64
	RangeBucket1001to2000Success int64
	RangeBucketOver2000          int64
	RangeBucketOver2000Success   int64

	// Subscriptions
	ActiveSubscriptions int32
	SubscriptionErrors  int64
}

// PrintMetrics prints current metrics to stdout
func (m *WorkerPoolMetrics) PrintMetrics() { _ = "STUB: not implemented"; return }

// Worker Pool Section

// DB Semaphore Section

// eth_getLogs Section

// Error Breakdown Section

// Block Range Distribution Section with success rate

// Subscriptions Section

// Alert conditions

// printRangeBucket prints a single range bucket with success rate
func printRangeBucket(name string, total, success, totalRequests int64) {
	_ = "STUB: not implemented"
	return
}

func repeatStr(s string, count int) string { _ = "STUB: not implemented"; return "" }

// ResetMetrics resets all metrics (useful for testing)
func (m *WorkerPoolMetrics) ResetMetrics() { _ = "STUB: not implemented"; return }

// Reset error breakdown

// Reset block range buckets

// Reset block range success buckets

// Reset subscriptions

// ========================================
// Prometheus Metrics Export Functions
// ========================================

// ExportPrometheusMetrics exports all metrics to OTel
// This should be called periodically (e.g., every 5 seconds)
// All metrics are exported as gauges for efficiency (batch export instead of per-operation)
func (m *WorkerPoolMetrics) ExportPrometheusMetrics() { _ = "STUB: not implemented"; return }

// Worker Pool Gauges

// Task counters (exported as gauges for batch efficiency)

// DB Semaphore Gauges

// Subscriptions Gauge

// eth_getLogs specific gauges

// Error breakdown

// Average timings
