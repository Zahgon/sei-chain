package store

import (
	sdkkv "github.com/sei-protocol/sei-chain/sei-cosmos/types/kv"
)

// Gets the first item.
func First(st KVStore, start, end []byte) (kv sdkkv.Pair, ok bool) {
	_ = "STUB: not implemented"
	return *new(sdkkv.Pair), false
}

// Gets the last item.  `end` is exclusive.
func Last(st KVStore, start, end []byte) (kv sdkkv.Pair, ok bool) {
	_ = "STUB: not implemented"
	return *new(sdkkv.Pair), false
}

// Skip this one, end is exclusive.
