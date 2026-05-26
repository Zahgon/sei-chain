package metrics

import (
	"time"

	"go.opentelemetry.io/otel/metric"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// Metrics to possibly add in the future:
//  - total disk used, broken down by root
//  - disk available on each root
//  - control loop idle fraction
//    - main control loop
//    - flush loop
//    - shard control loops
//    - keyfile control loop
//  - total number of segments
//  - average segment span (i.e. difference in time between first and last values written to a segment)
//  - segment creation rate
//  - used/unused segment space (useful for detecting shard assignment issues)

const littMeterName = "litt"

// LittDBMetrics encapsulates metrics for a LittDB. Metrics are exported via
// whatever exporter is configured on the global OTel MeterProvider (e.g.
// Prometheus, OTLP). The caller is responsible for setting up the provider
// before calling NewLittDBMetrics (see commonmetrics.SetupOtelPrometheus).
//
// Per-table observations are tagged with a "table" attribute. A nil
// LittDBMetrics acts as a no-op for all Report* methods.
type LittDBMetrics struct {
	// The size of individual tables in the database.
	tableSizeInBytes metric.Int64Gauge

	// The number of keys in individual tables in the database.
	tableKeyCount metric.Int64Gauge

	// The number of bytes read from disk since startup.
	bytesReadCounter metric.Int64Counter

	// The number of keys read from disk since startup.
	keysReadCounter metric.Int64Counter

	// The number of cache hits since startup.
	cacheHitCounter metric.Int64Counter

	// The number of cache misses since startup.
	cacheMissCounter metric.Int64Counter

	// Reports on the read latency of the database. This metric includes both cache hits and cache misses.
	readLatency metric.Float64Histogram

	// Reports on the write latency of the database, but only measures the time to read a value when a
	// cache miss occurs.
	cacheMissLatency metric.Float64Histogram

	// The number of bytes written to disk since startup. Only includes values, not metadata.
	bytesWrittenCounter metric.Int64Counter

	// The number of keys written to disk since startup.
	keysWrittenCounter metric.Int64Counter

	// Reports on the write latency of the database.
	writeLatency metric.Float64Histogram

	// The number of times a flush operation has been performed.
	flushCount metric.Int64Counter

	// Reports on the latency of a flush operation.
	flushLatency metric.Float64Histogram

	// Reports on the latency of a flushing segment files. This is a subset of the time spent during a flush operation.
	segmentFlushLatency metric.Float64Histogram

	// Reports on the latency of a keymap flush operation. This is a subset of the time spent during a flush operation.
	keymapFlushLatency metric.Float64Histogram

	// The latency of garbage collection operations.
	garbageCollectionLatency metric.Float64Histogram

	// Metrics for the write cache.
	writeCacheMetrics *util.CacheMetrics

	// Metrics for the read cache.
	readCacheMetrics *util.CacheMetrics
}

// NewLittDBMetrics creates a new LittDBMetrics instance backed by the global
// OTel MeterProvider. The caller must configure a MeterProvider with a
// Prometheus or other exporter before calling this (e.g. via
// commonmetrics.SetupOtelPrometheus).
func NewLittDBMetrics() *LittDBMetrics { _ = "STUB: not implemented"; return nil }

// tableAttr returns the OTel measurement option that tags an observation with
// the given table name. Allocated per call rather than cached because callers
// pass arbitrary table names; for hot-path call sites consider caching upstream.
func tableAttr(tableName string) metric.MeasurementOption {
	_ = "STUB: not implemented"
	return *new(metric.MeasurementOption)
}

// CollectPeriodicMetrics is a method that is periodically called to collect metrics. Tables are not permitted to be
// added or dropped while this method is running.
func (m *LittDBMetrics) CollectPeriodicMetrics(tables map[string]litt.ManagedTable) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // table size fits int64

//nolint:gosec // key count fits int64

// ReportReadOperation reports the results of a read operation.
func (m *LittDBMetrics) ReportReadOperation(
	tableName string,
	latency time.Duration,
	dataSize uint64,
	cacheHit bool) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // data size fits int64

// ReportWriteOperation reports the results of a write operation.
func (m *LittDBMetrics) ReportWriteOperation(
	tableName string,
	latency time.Duration,
	batchSize uint64,
	dataSize uint64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // data size fits int64
//nolint:gosec // batch size fits int64

// ReportFlushOperation reports the results of a flush operation.
func (m *LittDBMetrics) ReportFlushOperation(tableName string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ReportSegmentFlushLatency reports the amount of time taken to flush value files.
func (m *LittDBMetrics) ReportSegmentFlushLatency(tableName string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ReportKeymapFlushLatency reports the amount of time taken to flush the keymap.
func (m *LittDBMetrics) ReportKeymapFlushLatency(tableName string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ReportGarbageCollectionLatency reports the latency of a garbage collection operation.
func (m *LittDBMetrics) ReportGarbageCollectionLatency(tableName string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *LittDBMetrics) GetWriteCacheMetrics() *util.CacheMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (m *LittDBMetrics) GetReadCacheMetrics() *util.CacheMetrics {
	_ = "STUB: not implemented"
	return nil
}
