package pebbledb

import (
	"context"

	"github.com/cockroachdb/pebble/v2"

	"github.com/sei-protocol/sei-chain/sei-db/common/threading"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/dbcache"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// pebbleDB implements the db_engine.DB interface using PebbleDB.
type pebbleDB struct {
	db            *pebble.DB
	metricsCancel context.CancelFunc
}

var _ types.KeyValueDB = (*pebbleDB)(nil)

// Open opens (or creates) a Pebble-backed DB at path, returning a KeyValueDB
func Open(
	ctx context.Context,
	config *PebbleDBConfig,
) (_ types.KeyValueDB, err error) {
	_ = "STUB: not implemented"
	return *new(types.KeyValueDB), nil
}

// FormatMajorVersion is pinned to a specific version to prevent accidental
// breaking changes when updating the pebble dependency. Using FormatNewest
// would cause the on-disk format to silently upgrade when pebble is updated,
// making the database incompatible with older software versions.
// When upgrading this version, ensure it's an intentional, documented change.

// 64 MB

// Configure L0 with explicit settings
// 32 KB
// 256 KB

// Configure L1+ levels, inheriting from previous level

// 32 KB
// 256 KB

// Disable bloom filter at bottommost level (L6) - bloom filters are less useful
// at the bottom level since most data lives there and false positive rate is low

// OpenWithCache opens a Pebble-backed DB and wraps it with a read-through cache.
// When cacheConfig.MaxSize is 0 a no-op (passthrough) cache is used.
func OpenWithCache(
	ctx context.Context,
	config *PebbleDBConfig,
	cacheConfig *dbcache.CacheConfig,
	readPool threading.Pool,
	miscPool threading.Pool,
) (types.KeyValueDB, error) {
	_ = "STUB: not implemented"
	return *new(types.KeyValueDB), nil
}

func (p *pebbleDB) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pebbleDB) BatchGet(keys map[string]types.BatchGetResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pebbleDB) Set(key, value []byte, opts types.WriteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pebbleDB) Delete(key []byte, opts types.WriteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pebbleDB) NewIter(opts *types.IterOptions) (types.KeyValueDBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.KeyValueDBIterator), nil
}

func (p *pebbleDB) Flush() error { _ = "STUB: not implemented"; return nil }

func (p *pebbleDB) Checkpoint(destDir string) error { _ = "STUB: not implemented"; return nil }

var _ types.Checkpointable = (*pebbleDB)(nil)

func (p *pebbleDB) Close() error {
	_ = "STUB: not implemented"
	// Make Close idempotent: Pebble panics if Close is called twice.
	return nil
}

func toPebbleWriteOpts(opts types.WriteOptions) *pebble.WriteOptions {
	_ = "STUB: not implemented"
	return nil
}
