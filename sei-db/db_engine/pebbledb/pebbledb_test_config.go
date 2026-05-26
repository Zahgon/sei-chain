package pebbledb

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/dbcache"
)

// DefaultTestConfig returns a PebbleDBConfig suitable for testing.
// Allocates a smaller block cache and disables metrics.
func DefaultTestConfig(t *testing.T) PebbleDBConfig {
	_ = "STUB: not implemented"
	return *new(PebbleDBConfig)
}

// DefaultTestCacheConfig returns a CacheConfig suitable for testing.
func DefaultTestCacheConfig() dbcache.CacheConfig {
	_ = "STUB: not implemented"
	return *new(dbcache.CacheConfig)
}
