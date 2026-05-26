package mvcc

import (
	"github.com/cockroachdb/pebble/v2"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// This file contains the ascending-version MVCC implementation used to read
// and write legacy DBs that were created before the descending-version fast
// path was introduced. It is a verbatim port of main's Get/Has/Iterator/
// ReverseIterator/Prune path, adjusted only to use the *Ascending encoding
// helpers and the ascendingIterator type. Archive nodes that cannot migrate
// will continue to hit this path.

func (db *Database) hasAscending(storeKey string, version int64, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (db *Database) getAscending(storeKey string, targetVersion int64, key []byte) (_ []byte, _err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A tombstone of zero or a target version that is less than the tombstone
// version means the key is not deleted at the target version.

// A tombstone of zero or a target version that is less than the tombstone
// version means the key is not deleted at the target version.

// the value is considered deleted

func (db *Database) pruneAscending(version int64) (_err error) {
	_ = "STUB: not implemented"
	// Defensive check: ensure database is not closed
	return nil
}

// we increment by 1 to include the provided version

// Ignore metadata entries during pruning

// Store current key and version

// XXX: This should never happen given we skip the metadata keys.

// For every new module visited, check to see last time it was updated

// Skip a store's keys if version it was last updated is less than last prune height

// Seek to next key if we are at a version which is higher than prune height
// Do not seek to next key if KeepLastVersion is false and we need to delete the previous key in pruning

// Delete a key if another entry for that key exists at a larger version than original but leq to the prune height
// Also delete a key if it has been tombstoned and its version is leq to the prune height
// Also delete a key if KeepLastVersion is false and version is leq to the prune height

// Update prevKey and prevVersion for next iteration

// Commit any leftover delete ops in batch

func (db *Database) iteratorAscending(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func (db *Database) reverseIteratorAscending(storeKey string, version int64, start, end []byte) (types.DBIterator, error) {
	_ = "STUB: not implemented"
	return *new(types.DBIterator), nil
}

func getMVCCSliceAscending(db *pebble.DB, storeKey string, key []byte, version int64) ([]byte, error) {
	_ = "STUB: not implemented"
	// end domain is exclusive, so we need to increment the version by 1
	return nil, nil
}
