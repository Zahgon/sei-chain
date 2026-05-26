package dbcache

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/metric"
)

const cacheMeterName = "seidb_pebblecache"

// CacheMetrics records OTel metrics for a pebblecache instance.
// All report methods are nil-safe: if the receiver is nil, they are no-ops,
// allowing the cache to call them unconditionally regardless of whether metrics
// are enabled.
//
// The cacheName is used as the "cache" attribute on all recorded metrics,
// enabling multiple cache instances to be distinguished in dashboards.
type CacheMetrics struct {
	// Pre-computed attribute option reused on every recording to avoid
	// per-call allocations on the hot path.
	attrs metric.MeasurementOption

	sizeBytes   metric.Int64Gauge
	sizeEntries metric.Int64Gauge
	hits        metric.Int64Counter
	misses      metric.Int64Counter
	missLatency metric.Float64Histogram
}

// newCacheMetrics creates a CacheMetrics that records cache statistics via OTel.
// A background goroutine scrapes cache size every scrapeInterval until ctx is
// cancelled. The cacheName is attached as the "cache" attribute to all recorded
// metrics, enabling multiple cache instances to be distinguished in dashboards.
//
// Multiple instances are safe: OTel instrument registration is idempotent, so each
// call receives references to the same underlying instruments. The "cache" attribute
// distinguishes series (e.g. pebblecache_hits{cache="state"}).
func newCacheMetrics(
	ctx context.Context,
	cacheName string,
	scrapeInterval time.Duration,
	getSize func() (bytes uint64, entries uint64),
) *CacheMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (cm *CacheMetrics) reportCacheHits(count int64) { _ = "STUB: not implemented"; return }

func (cm *CacheMetrics) reportCacheMisses(count int64) { _ = "STUB: not implemented"; return }

func (cm *CacheMetrics) reportCacheMissLatency(latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// collectLoop periodically scrapes cache size from the provided function
// and records it as gauge values. It exits when ctx is cancelled.
func (cm *CacheMetrics) collectLoop(
	ctx context.Context,
	interval time.Duration,
	getSize func() (bytes uint64, entries uint64),
) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // G115: safe, cache size fits int64
//nolint:gosec // G115: safe, entry count fits int64
