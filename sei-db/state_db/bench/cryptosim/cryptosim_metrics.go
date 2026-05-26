package cryptosim

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/metric"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
)

const cryptosimMeterName = "cryptosim"

var receiptWriteLatencyBuckets = []float64{
	0.001, 0.0025, 0.005, 0.0075, 0.01,
	0.015, 0.02, 0.03, 0.05, 0.075,
	0.1, 0.25, 0.5, 0.75, 1, 2.5, 5,
}

var receiptReadLatencyBuckets = []float64{
	0.00001, 0.00005, 0.0001, 0.00025, 0.0005,
	0.001, 0.0025, 0.005, 0.01, 0.025,
	0.05, 0.1, 0.25, 0.5, 1,
}

var receiptLogFilterLatencyBuckets = []float64{
	0.00001, 0.00005, 0.0001, 0.00025, 0.0005,
	0.001, 0.0025, 0.005, 0.01, 0.025,
	0.05, 0.075, 0.1, 0.15, 0.25,
	0.5, 0.75, 1, 1.5, 2,
	2.5, 3, 4, 5, 7.5, 10,
}

// CryptosimMetrics holds OpenTelemetry metrics for the cryptosim benchmark.
// Metrics are exported via whatever exporter is configured on the global OTel
// MeterProvider (e.g., Prometheus, OTLP). This package does not import Prometheus.
type CryptosimMetrics struct {
	ctx context.Context

	blocksFinalizedTotal       metric.Int64Counter
	transactionsProcessedTotal metric.Int64Counter
	totalAccounts              metric.Int64Gauge
	hotAccounts                metric.Int64Gauge
	coldAccounts               metric.Int64Gauge
	dormantAccounts            metric.Int64Gauge
	totalErc20Contracts        metric.Int64Gauge
	dbCommitsTotal             metric.Int64Counter
	dataDirSizeBytes           metric.Int64Gauge
	dataDirAvailableBytes      metric.Int64Gauge
	logDirSizeBytes            metric.Int64Gauge
	processReadBytesTotal      metric.Int64Counter
	processWriteBytesTotal     metric.Int64Counter
	processReadCountTotal      metric.Int64Counter
	processWriteCountTotal     metric.Int64Counter
	uptimeSeconds              metric.Float64Gauge

	// Receipt metrics
	receiptBlockWriteDuration      metric.Float64Histogram
	receiptChannelDepth            metric.Int64Gauge
	receiptsWrittenTotal           metric.Int64Counter
	receiptErrorsTotal             metric.Int64Counter
	receiptReadDuration            metric.Float64Histogram
	receiptReadsTotal              metric.Int64Counter
	receiptCacheHitsTotal          metric.Int64Counter
	receiptCacheMissesTotal        metric.Int64Counter
	receiptReadsFoundTotal         metric.Int64Counter
	receiptReadsNotFoundTotal      metric.Int64Counter
	receiptLogFilterDuration       metric.Float64Histogram
	receiptLogFilterCacheHitsTotal metric.Int64Counter
	receiptLogFilterCacheMissTotal metric.Int64Counter
	receiptLogFilterLogsReturned   metric.Int64Histogram
	cacheFilterScanDuration        metric.Float64Histogram
	cacheGetDuration               metric.Float64Histogram

	mainThreadPhase              *metrics.PhaseTimer
	transactionPhaseTimerFactory *metrics.PhaseTimerFactory
}

// NewCryptosimMetrics creates metrics for the cryptosim benchmark using the
// global OTel MeterProvider. The caller (e.g., main) must configure the
// MeterProvider with a Prometheus or other exporter before calling this.
// When ctx is cancelled, background sampling goroutines exit.
// Data directory size sampling is started automatically when
// BackgroundMetricsScrapeInterval > 0.
//
// Unit convention: Use WithUnit values from the UCUM standard (see
// https://ucum.org/ucum). Durations use "s" (seconds). Bytes use "By".
// Counts use curly-brace annotations, e.g. "{count}" for generic counts or
// more specific "{block}", "{transaction}", "{account}" to match what is measured.
func NewCryptosimMetrics(
	ctx context.Context,
	dbPhaseTimer *metrics.PhaseTimer,
	config *CryptoSimConfig,
) *CryptosimMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (m *CryptosimMetrics) startUptimeSampling(startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// startProcessIOSampling starts a goroutine that periodically samples process
// I/O counters (read/write bytes and operation counts) via gopsutil and adds
// deltas to OTel counters. Use rate() on these counters for throughput and IOPS.
// Skipped on darwin: gopsutil does not implement process.IOCounters on macOS.
func (m *CryptosimMetrics) startProcessIOSampling(intervalSeconds int) {
	_ = "STUB: not implemented"
	return
}

// startPeriodicSampling runs sampleFn immediately and then every interval
// seconds in a background goroutine until m.ctx is cancelled.
func (m *CryptosimMetrics) startPeriodicSampling(intervalSeconds int, sampleFn func()) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) startDirSizeSampling(dir string, gauge metric.Int64Gauge, intervalSeconds int) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) startAvailableDiskSpaceSampling(dir string, intervalSeconds int) {
	_ = "STUB: not implemented"
	return
}

// uint64ToInt64Clamped converts a uint64 to int64, clamping to math.MaxInt64 to avoid overflow.
func uint64ToInt64Clamped(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

func measureDataDirAvailableBytes(dataDir string) int64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec

func measureDataDirSize(dataDir string) int64 { _ = "STUB: not implemented"; return 0 }

func (m *CryptosimMetrics) ReportBlockFinalized(transactionCount int64) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) ReportDBCommit() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) SetTotalNumberOfAccounts(total int64, hot int64, cold int64) {
	_ = "STUB: not implemented"
	return
}

// IncrementTotalNumberOfAccounts updates the account gauges after adding one account.
// Pass the new totals: total, hot, cold. Dormant is derived as total - hot - cold.
func (m *CryptosimMetrics) IncrementTotalNumberOfAccounts(total int64, hot int64, cold int64) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) SetTotalNumberOfERC20Contracts(total int64) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) GetTransactionPhaseTimerInstance() *metrics.PhaseTimer {
	_ = "STUB: not implemented"
	return nil
}

func (m *CryptosimMetrics) SetMainThreadPhase(phase string) { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) RecordReceiptBlockWriteDuration(latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) ReportReceiptsWritten(count int64) { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) ReportReceiptError() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) RecordReceiptReadDuration(seconds float64) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) ReportReceiptRead() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) ReportReceiptCacheHit() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) ReportReceiptCacheMiss() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) ReportReceiptReadFound() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) ReportReceiptReadNotFound() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) RecordReceiptLogFilterDuration(seconds float64) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) ReportLogFilterCacheHit() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) ReportLogFilterCacheMiss() { _ = "STUB: not implemented"; return }

func (m *CryptosimMetrics) RecordLogFilterLogsReturned(count int64) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) RecordCacheFilterScanDuration(seconds float64) {
	_ = "STUB: not implemented"
	return
}

func (m *CryptosimMetrics) RecordCacheGetDuration(seconds float64) {
	_ = "STUB: not implemented"
	return
}

// startReceiptChannelDepthSampling periodically records the depth of the receipt channel.
func (m *CryptosimMetrics) startReceiptChannelDepthSampling(ch <-chan *block, intervalSeconds int) {
	_ = "STUB: not implemented"
	return
}
