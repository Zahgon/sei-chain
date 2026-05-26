package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
)

// Iterator over all the keys with a certain prefix in ascending order
func KVStorePrefixIterator(kvs KVStore, prefix []byte) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

// Iterator over all the keys with a certain prefix in descending order.
func KVStoreReversePrefixIterator(kvs KVStore, prefix []byte) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

// DiffKVStores compares two KVstores and returns all the key/value pairs
// that differ from one another. It also skips value comparison for a set of provided prefixes.
func DiffKVStores(a KVStore, b KVStore, prefixesToSkip [][]byte) (kvAs, kvBs []kv.Pair) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip value comparison if we matched a prefix

// PrefixEndBytes returns the []byte that would end a
// range query for all []byte with a certain prefix
// Deals with last byte of prefix being FF without overflowing
func PrefixEndBytes(prefix []byte) []byte { _ = "STUB: not implemented"; return nil }

// InclusiveEndBytes returns the []byte that would end a
// range query such that the input would be included
func InclusiveEndBytes(inclusiveBytes []byte) []byte { _ = "STUB: not implemented"; return nil }
