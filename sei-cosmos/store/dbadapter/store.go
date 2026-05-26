package dbadapter

import (
	"io"

	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

// Wrapper type for dbm.Db with implementation of KVStore
type Store struct {
	dbm.DB
}

func (dsa Store) GetWorkingHash() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Get wraps the underlying DB's Get method panicing on error.
		nil
}

func (dsa Store) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Has wraps the underlying DB's Has method panicing on error.
func (dsa Store) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

// Set wraps the underlying DB's Set method panicing on error.
func (dsa Store) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// Delete wraps the underlying DB's Delete method panicing on error.
func (dsa Store) Delete(key []byte) { _ = "STUB: not implemented"; return }

// Iterator wraps the underlying DB's Iterator method panicing on error.
func (dsa Store) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// ReverseIterator wraps the underlying DB's ReverseIterator method panicing on error.
func (dsa Store) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// GetStoreType returns the type of the store.
func (Store) GetStoreType() types.StoreType {
	_ = "STUB: not implemented"
	return *

	// CacheWrap branches the underlying store.
	new(types.StoreType)
}

func (dsa Store) CacheWrap(storeKey types.StoreKey) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

// CacheWrapWithTrace implements KVStore.
func (dsa Store) CacheWrapWithTrace(storeKey types.StoreKey, w io.Writer, tc types.TraceContext) types.CacheWrap {
	_ = "STUB: not implemented"
	return *new(types.CacheWrap)
}

func (dsa Store) VersionExists(version int64) bool { _ = "STUB: not implemented"; return false }

func (dsa Store) DeleteAll(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (dsa Store) GetAllKeyStrsInRange(start, end []byte) (res []string) {
	_ = "STUB: not implemented"
	return nil
}

// dbm.DB implements KVStore so we can CacheKVStore it.
var _ types.KVStore = Store{}
