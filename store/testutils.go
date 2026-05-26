package store

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func NewTestKVStore() types.KVStore { _ = "STUB: not implemented"; return *new(types.KVStore) }

func NewTestCacheMultiStore(stores map[types.StoreKey]types.CacheWrapper) types.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(types.CacheMultiStore)
}
