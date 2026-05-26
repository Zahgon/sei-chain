package config

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/dbcache"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/pebbledb"
)

func smallTestPebbleConfig() pebbledb.PebbleDBConfig {
	_ = "STUB: not implemented"
	return *new(pebbledb.PebbleDBConfig)
}

func smallTestCacheConfig() dbcache.CacheConfig {
	_ = "STUB: not implemented"
	return *new(dbcache.CacheConfig)
}

// DefaultTestConfig returns a Config suitable for unit tests. It uses
// t.TempDir() as the DataDir root, small cache sizes, and disables metrics.
func DefaultTestConfig(t *testing.T) *Config { _ = "STUB: not implemented"; return nil }
