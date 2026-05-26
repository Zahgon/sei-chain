package cachemulti

import (
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// Since `CacheMultiStore` has a method with the same name, we have to
// type alias here or otherwise we won't be able to inherit or implement
// `CacheMultiStore` the method.
type sdkCacheMultiStore = storetypes.CacheMultiStore

type Store struct {
	sdkCacheMultiStore

	storeKeyToWriteWhitelist map[string][]string
}

func NewStore(parent storetypes.CacheMultiStore, storeKeyToWriteWhitelist map[string][]string) storetypes.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(storetypes.CacheMultiStore)
}

func (cms Store) CacheMultiStore() storetypes.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(storetypes.CacheMultiStore)
}

func (cms Store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	_ = "STUB: not implemented"
	return *new(storetypes.KVStore)
}

// whitelist nothing
