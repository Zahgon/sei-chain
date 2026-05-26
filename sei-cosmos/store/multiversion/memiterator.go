package multiversion

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	occtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/occ"
)

// Iterates over iterKVCache items.
// if key is nil, means it was deleted.
// Implements Iterator.
type memIterator struct {
	types.Iterator
	mvkv *VersionIndexedStore
}

func (store *VersionIndexedStore) newMemIterator(
	start, end []byte,
	items *dbm.MemDB,
	ascending bool,
) *memIterator {
	_ = "STUB: not implemented"
	return nil
}

// try to get value from the writeset, otherwise try to get from multiversion store, otherwise try to get from parent
func (mi *memIterator) Value() []byte {
	_ = "STUB: not implemented"

	// TODO: verify that this is correct
	return nil
}

type validationIterator struct {
	types.Iterator

	mvStore      MultiVersionStore
	writeset     WriteSet
	index        int
	abortChannel chan occtypes.Abort

	// this ensure that we serve consistent values throughout the lifecycle of the validationIterator - this should prevent race conditions causing an iterator to become invalid while being used
	readCache map[string][]byte
}

func (store *Store) newMVSValidationIterator(
	index int,
	start, end []byte,
	items *dbm.MemDB,
	ascending bool,
	writeset WriteSet,
	abortChannel chan occtypes.Abort,
) *validationIterator {
	_ = "STUB: not implemented"
	return nil
}

// try to get value from the writeset, otherwise try to get from multiversion store, otherwise try to get from parent iterator
func (vi *validationIterator) Value() []byte {
	_ = "STUB: not implemented"

	// try fetch from writeset - return if exists
	return nil
}

// serve value from readcache (means it has previously been accessed by this iterator so we want consistent behavior here)

// get the value from the multiversion store

// if we have an estimate, write to abort channel

// if we have a deleted value, return nil
