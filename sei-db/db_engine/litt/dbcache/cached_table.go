package dbcache

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/types"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

var _ litt.ManagedTable = &cachedTable{}

// cachedTable wraps a table and adds caching functionality.
type cachedTable struct {
	// The base table to wrap.
	base litt.ManagedTable
	// This cache holds values that were recently written to the table.
	writeCache util.Cache[string, []byte]
	// This cache holds values that were recently read from the base table.
	readCache util.Cache[string, []byte]
	// Metrics for the table.
	metrics *metrics.LittDBMetrics
}

// NewCachedTable creates wrapper around a table that caches recently written and read values.
func NewCachedTable(
	base litt.ManagedTable,
	writeCache util.Cache[string, []byte],
	readCache util.Cache[string, []byte],
	metrics *metrics.LittDBMetrics,
) litt.ManagedTable {
	_ = "STUB: not implemented"
	return *new(litt.ManagedTable)
}

func (c *cachedTable) KeyCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *cachedTable) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *cachedTable) Name() string { _ = "STUB: not implemented"; return "" }

func (c *cachedTable) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) PutBatch(batch []*types.KVPair) error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) Get(key []byte) (value []byte, exists bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// In theory, there is a race condition here where call to CacheAwareGet() made concurrently with a call to Put()
// might find the data to exist but not to be hot. This is not a problem though, since it will be hard to trigger and
// since it is not a violation of the consistency/correctness guarantees made by LittDB. Caching is inherently a
// "best effort" optimization, and so it's not worth adding extra locking in order to prevent this edge case.
//
// Scenario:
// - Thread A calls Put() on key K, and Put() does not return right away.
// - Thread B calls CacheAwareGet() on key K with onlyReadFromCache set to true.
// - Thread B checks the cache, and finds that the value is not there.
// - LittDB flushes the value out to disk before thread A's Put() returns, specifically before thread A inserts
//   the value into the write cache. The timing of this is exceptionally unlikely, but not impossible.
// - Thread B gets to the part of CacheAwareGet() where it checks the base table for the value. Since the
//   base table has flushed the value out to disk, it says that the value exists but does not fetch it since
//   onlyReadFromCache is true.
// - Thread A finishes calling Put(), and key K is now in the cache.
//
//   |                     Thread A                                               Thread B
//  Time                      |                                                      |
//   |             Put(key K, ...) starts                                            |
//   v                        |                                                      |
//                            |                                 CacheAwareGet(key K, ...) -> value not present
//                            |                                                      |
//      K is inserted into the unflushed data map                                    |
//                            |                                                      |
//                            |                                 CacheAwareGet(key K, ...) -> present and hot
//                            |                                                      |
//     K is flushed to disk and removed from the unflushed data map                  |
//         (highly irregular but not impossible timing)                              |
//                            |                                                      |
//                            |                                 CacheAwareGet(key K, ...) -> present and cold
//                            |                                                      |
//           K is inserted into the write cache                                      |
//                            |                                                      |
//                            |                                 CacheAwareGet(key K, ...) -> present and hot
//                            |                                                      |
//                  Put (key K, ...) returns                                         |

func (c *cachedTable) CacheAwareGet(
	key []byte,
	onlyReadFromCache bool,
) (value []byte, exists bool, hot bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, false, nil
}

// The value was recently written

// The value was recently read

func (c *cachedTable) Exists(key []byte) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *cachedTable) Flush() error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) SetTTL(ttl time.Duration) error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) SetWriteCacheSize(size uint64) error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) SetReadCacheSize(size uint64) error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) Close() error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) Destroy() error { _ = "STUB: not implemented"; return nil }

func (c *cachedTable) SetShardingFactor(shardingFactor uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cachedTable) RunGC() error { _ = "STUB: not implemented"; return nil }
