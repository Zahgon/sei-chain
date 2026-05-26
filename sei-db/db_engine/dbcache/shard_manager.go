package dbcache

import (
	"errors"
	"hash/maphash"
	"sync"
)

var ErrNumShardsNotPowerOfTwo = errors.New("numShards must be a power of two and > 0")

// A utility for assigning keys to shard indices.
type shardManager struct {
	// A random seed that makes it hard for an attacker to predict the shard index and to skew the distribution.
	seed maphash.Seed
	// Used to perform a quick modulo operation to get the shard index (since numShards is a power of two)
	mask uint64
	// reusable Hash objects to avoid allocs
	pool sync.Pool
}

// Creates a new Sharder. Number of shards must be a power of two and greater than 0.
func newShardManager(numShards uint64) (*shardManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// secret, randomized

// Shard returns a shard index in [0, numShards).
// addr should be the raw address bytes (e.g., 20-byte ETH address).
func (s *shardManager) Shard(addr []byte) uint64 { _ = "STUB: not implemented"; return 0 }
