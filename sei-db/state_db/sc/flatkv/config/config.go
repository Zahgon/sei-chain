package config

import (
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/dbcache"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/pebbledb"
)

const (
	DefaultSnapshotInterval   uint32 = 10000
	DefaultSnapshotKeepRecent uint32 = 2
)

// Config defines configuration for the FlatKV (EVM) commit store.
type Config struct {
	// DataDir is the root directory for the FlatKV data files.
	// Must be set before calling Validate().
	DataDir string

	// Fsync controls whether PebbleDB writes (data DBs + metadataDB) use fsync.
	// WAL always uses NoSync (matching memiavl); crash recovery relies on
	// WAL catchup, which is idempotent.
	// Default: false
	Fsync bool `mapstructure:"fsync"`

	// AsyncWriteBuffer defines the size of the async write buffer for data DBs.
	// Set <= 0 for synchronous writes.
	// Default: 0 (synchronous)
	AsyncWriteBuffer int `mapstructure:"async-write-buffer"`

	// SnapshotInterval defines how often (in blocks) a PebbleDB checkpoint
	// snapshot is taken. 0 disables auto-snapshots.
	// Without periodic snapshots the WAL grows unbounded and every restart
	// replays the entire history from snapshot-0.
	// Default: 10000
	SnapshotInterval uint32 `mapstructure:"snapshot-interval"`

	// SnapshotKeepRecent defines how many old snapshots to keep besides the
	// latest one. 0 means keep only the current snapshot (no old snapshots).
	// Default: 2
	SnapshotKeepRecent uint32 `mapstructure:"snapshot-keep-recent"`

	// EnablePebbleMetrics defines if the Pebble metrics should be enabled.
	// Default: true
	EnablePebbleMetrics bool `mapstructure:"enable-pebble-metrics"`

	// AccountDBConfig defines the PebbleDB configuration for the account database.
	AccountDBConfig pebbledb.PebbleDBConfig

	// AccountCacheConfig defines the cache configuration for the account database.
	AccountCacheConfig dbcache.CacheConfig

	// CodeDBConfig defines the PebbleDB configuration for the code database.
	CodeDBConfig pebbledb.PebbleDBConfig

	// CodeCacheConfig defines the cache configuration for the code database.
	CodeCacheConfig dbcache.CacheConfig

	// StorageDBConfig defines the PebbleDB configuration for the storage database.
	StorageDBConfig pebbledb.PebbleDBConfig

	// StorageCacheConfig defines the cache configuration for the storage database.
	StorageCacheConfig dbcache.CacheConfig

	// LegacyDBConfig defines the PebbleDB configuration for the legacy database.
	LegacyDBConfig pebbledb.PebbleDBConfig

	// LegacyCacheConfig defines the cache configuration for the legacy database.
	LegacyCacheConfig dbcache.CacheConfig

	// MetadataDBConfig defines the PebbleDB configuration for the metadata database.
	MetadataDBConfig pebbledb.PebbleDBConfig

	// MetadataCacheConfig defines the cache configuration for the metadata database.
	MetadataCacheConfig dbcache.CacheConfig

	// Controls the number of goroutines in the DB read pool. The number of threads in this pool is equal to
	// ReaderThreadsPerCore * runtime.NumCPU() + ReaderConstantThreadCount.
	ReaderThreadsPerCore float64

	// Controls the number of goroutines in the DB read pool. The number of threads in this pool is equal to
	// ReaderThreadsPerCore * runtime.NumCPU() + ReaderConstantThreadCount.
	ReaderConstantThreadCount int

	// Controls the size of the queue for work sent to the read pool.
	ReaderPoolQueueSize int

	// Controls the number of goroutines pre-allocated in the thread pool for miscellaneous operations.
	// The number of threads in this pool is equal to MiscThreadsPerCore * runtime.NumCPU() + MiscConstantThreadCount.
	MiscPoolThreadsPerCore float64

	// Controls the number of goroutines pre-allocated in the thread pool for miscellaneous operations.
	// The number of threads in this pool is equal to MiscThreadsPerCore * runtime.NumCPU() + MiscConstantThreadCount.
	MiscConstantThreadCount int
}

// DefaultConfig returns Config with safe default values.
func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// Copy returns a deep copy of the Config.
func (c *Config) Copy() *Config {
	_ = "STUB: not implemented"
	// The nested PebbleDB configs are value types, so a shallow struct copy is sufficient.
	return nil
}

// Validate checks that the configuration is sane and returns an error if it is not.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
