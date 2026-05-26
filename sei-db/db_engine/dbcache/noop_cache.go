package dbcache

import (
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

var _ Cache = (*noOpCache)(nil)

// noOpCache is a Cache that performs no caching. Every Get falls through
// to the provided Reader. Set, Delete, and BatchSet are no-ops.
// Useful for testing the storage layer without cache interference, or for
// workloads where caching is not beneficial.
type noOpCache struct{}

// NewNoOpCache creates a Cache that always reads via the provided Reader and never caches.
func NewNoOpCache() Cache { _ = "STUB: not implemented"; return *new(Cache) }

func (c *noOpCache) Get(read Reader, key []byte, _ bool) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *noOpCache) BatchGet(read Reader, keys map[string]types.BatchGetResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *noOpCache) Set([]byte, []byte) {
	_ = "STUB: not implemented"
	// intentional no-op
	return
}

func (c *noOpCache) Delete([]byte) {
	_ = "STUB: not implemented"
	// intentional no-op
	return
}

func (c *noOpCache) BatchSet([]CacheUpdate) error {
	_ = "STUB: not implemented"
	// intentional no-op
	return nil
}
