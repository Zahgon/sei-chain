package util

import (
	"time"

	"go.opentelemetry.io/otel/metric"
)

const cacheMeterName = "litt"

// CacheMetrics is a struct that holds OTel metrics for a cache. A nil
// CacheMetrics instance acts as a no-op for all report* methods.
//
// Multiple CacheMetrics instances may be created for the same process; each
// receives references to the same underlying instruments because OTel
// instrument registration is idempotent. The "cache" attribute (set at
// construction time) distinguishes series in the exporter
// (e.g. litt_chunk_cache_keys_added{cache="chunk_read"}).
type CacheMetrics struct {
	// Pre-computed attribute option reused on every recording to avoid
	// per-call allocations on the hot path.
	attrs metric.MeasurementOption

	keyCount        metric.Int64Gauge
	weight          metric.Int64Gauge
	keysAdded       metric.Int64Counter
	weightAdded     metric.Int64Counter
	evictionLatency metric.Float64Histogram
}

// NewCacheMetrics creates a new CacheMetrics that records via the global OTel
// MeterProvider. The cacheName is attached as the "cache" attribute on every
// observation, allowing multiple cache instances to be distinguished in the
// exporter (for example "chunk_read" vs "chunk_write").
//
// The caller must have configured a MeterProvider before calling this (e.g.
// commonmetrics.SetupOtelPrometheus).
func NewCacheMetrics(cacheName string) *CacheMetrics { _ = "STUB: not implemented"; return nil }

// reportInsertion is used to report an entry being inserted into the cache.
func (m *CacheMetrics) reportInsertion(weight uint64) { _ = "STUB: not implemented"; return }

//nolint:gosec // weight fits int64

// reportEviction is used to report an entry being evicted from the cache.
func (m *CacheMetrics) reportEviction(age time.Duration) { _ = "STUB: not implemented"; return }

// reportCurrentSize is used to report the current size/weight of the cache.
func (m *CacheMetrics) reportCurrentSize(size int, weight uint64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // size fits int64
//nolint:gosec // weight fits int64
