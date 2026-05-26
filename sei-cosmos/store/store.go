package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	_ = "STUB: not implemented"
	return *new(types.CommitMultiStore)
}

func NewCommitMultiStoreWithArchival(db dbm.DB, archivalDb dbm.DB, archivalVersion int64) types.CommitMultiStore {
	_ = "STUB: not implemented"
	return *new(types.CommitMultiStore)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	_ = "STUB: not implemented"
	return *new(types.MultiStorePersistentCache)
}
