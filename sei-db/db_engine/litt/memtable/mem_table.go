package memtable

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/types"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

var _ litt.ManagedTable = &memTable{}

// expirationRecord is a record of when a key was inserted into the table.
type expirationRecord struct {
	// The time at which the key was inserted into the table.
	creationTime time.Time
	// A stringified version of the key.
	key string
}

// memTable is a simple implementation of a Table that stores its data in memory.
type memTable struct {
	// A function that returns the current time.
	clock func() time.Time

	// The name of the table.
	name string

	// The time-to-live for data in this table.
	ttl time.Duration

	// The actual data store.
	data map[string][]byte

	// Keeps track of when data should be deleted.
	expirationQueue *util.Queue[*expirationRecord]

	// Protects access to data and expirationQueue.
	//
	// This implementation could be made with smaller granularity locks to improve multithreaded performance,
	// at the cost of code complexity. But since this implementation is primary intended for use in tests,
	// such optimization is not necessary.
	lock sync.RWMutex

	shutdown atomic.Bool
}

// NewMemTable creates a new in-memory table.
func NewMemTable(config *litt.Config, name string) litt.ManagedTable {
	_ = "STUB: not implemented"
	return *new(litt.ManagedTable)
}

// this is a class designed for use in testing, not worth properly handling errors

func (m *memTable) Size() uint64 {
	_ = "STUB: not implemented"
	// Technically speaking, this table stores zero bytes on disk, and this method
	// is contractually obligated to return only the size of the data on disk.
	return 0
}

func (m *memTable) Name() string { _ = "STUB: not implemented"; return "" }

func (m *memTable) KeyCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (m *memTable) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

func (m *memTable) PutBatch(batch []*types.KVPair) error { _ = "STUB: not implemented"; return nil }

func (m *memTable) Get(key []byte) (value []byte, exists bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *memTable) CacheAwareGet(key []byte, _ bool) (value []byte, exists bool, hot bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, false, nil
}

func (m *memTable) Exists(key []byte) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *memTable) Flush() error {
	_ = "STUB: not implemented"
	// This is a no-op for a memory table. Memory tables are ephemeral by nature.
	return nil
}

func (m *memTable) SetTTL(ttl time.Duration) error { _ = "STUB: not implemented"; return nil }

func (m *memTable) Destroy() error { _ = "STUB: not implemented"; return nil }

func (m *memTable) Close() error { _ = "STUB: not implemented"; return nil }

func (m *memTable) SetWriteCacheSize(size uint64) error { _ = "STUB: not implemented"; return nil }

func (m *memTable) SetReadCacheSize(size uint64) error { _ = "STUB: not implemented"; return nil }

func (m *memTable) SetShardingFactor(shardingFactor uint8) error {
	_ = "STUB: not implemented"
	// the memory table has no concept of sharding
	return nil
}

func (m *memTable) RunGC() error { _ = "STUB: not implemented"; return nil }
