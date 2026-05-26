package dbcache

import (
	"time"
)

// CacheConfig defines configuration for a sharded LRU read-through cache.
type CacheConfig struct {
	// The number of shards in the cache. Must be a power of two and greater than 0.
	ShardCount uint64
	// The maximum size of the cache, in bytes. 0 disables the cache.
	MaxSize uint64
	// The estimated overhead per entry, in bytes. Used to calculate effective cache
	// capacity. Derive experimentally; may differ between builds and architectures.
	EstimatedOverheadPerEntry uint64
	// Name used as the "cache" attribute on OTel metrics. Empty string disables metrics.
	MetricsName string
	// How often to scrape cache size for metrics. Ignored if MetricsName is empty.
	MetricsScrapeInterval time.Duration
}

// DefaultCacheConfig returns a CacheConfig with sensible defaults.
func DefaultCacheConfig() CacheConfig { _ = "STUB: not implemented"; return *new(CacheConfig) }

// Validate checks that the configuration is sane and returns an error if it is not.
func (c *CacheConfig) Validate() error { _ = "STUB: not implemented"; return nil }
