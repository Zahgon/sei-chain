package dbcache

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-db/common/threading"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

var _ Cache = (*cache)(nil)

// A standard implementation of a flatcache.
type cache struct {
	ctx context.Context

	// A utility for assigning keys to shard indices.
	shardManager *shardManager

	// The shards in the cache.
	shards []*shard

	// A pool for asynchronous reads.
	readPool threading.Pool

	// A pool for miscellaneous operations that are neither computationally intensive nor IO bound.
	miscPool threading.Pool
}

// Creates a new Cache. If cfg.MetricsName is non-empty, OTel metrics are enabled and the
// background size scrape runs every cfg.MetricsScrapeInterval.
func NewStandardCache(
	ctx context.Context,
	cfg *CacheConfig,
	readPool threading.Pool,
	miscPool threading.Pool,
) (Cache, error) {
	_ = "STUB: not implemented"
	return *new(Cache), nil
}

func (c *cache) getCacheSizeInfo() (bytes uint64, entries uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (c *cache) BatchSet(updates []CacheUpdate) error {
	_ = "STUB: not implemented"
	// Sort entries by shard index so each shard is locked only once.
	return nil
}

func (c *cache) BatchGet(read Reader, keys map[string]types.BatchGetResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cache) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (c *cache) Get(read Reader, key []byte, updateLru bool) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *cache) Set(key []byte, value []byte) { _ = "STUB: not implemented"; return }
