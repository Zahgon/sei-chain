package multi

import (
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

type Store struct {
	storetypes.MultiStore

	storeKeyToWriteWhitelist map[string][]string
}

func NewStore(parent storetypes.MultiStore, storeKeyToWriteWhitelist map[string][]string) storetypes.MultiStore {
	_ = "STUB: not implemented"
	return *new(storetypes.MultiStore)
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
