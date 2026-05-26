package kv

import (
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

type Store struct {
	storetypes.KVStore

	writeWhitelist map[string]struct{}
}

func NewStore(parent storetypes.KVStore, writeWhitelistKeys []string) storetypes.KVStore {
	_ = "STUB: not implemented"
	return *new(storetypes.KVStore)
}

func (store *Store) Set(key []byte, value []byte) { _ = "STUB: not implemented"; return }

func (store *Store) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (store *Store) validateKeyForWrite(key []byte) { _ = "STUB: not implemented"; return }

// Panic since the Store interface does not return error on Set. Can be
// intercepted by calling routine through `err := recover()`
