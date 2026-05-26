package mvcc

import (
	"sync"

	"github.com/cockroachdb/pebble/v2"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// This file contains the ascending-version MVCC iterator used for legacy DBs
// that were written by the pre-descending build. It is a verbatim port of the
// iterator implementation from main and is intentionally kept isolated from
// the descending fast-path iterator to avoid subtle interactions between the
// two encoding schemes.
//
// Archive nodes that cannot migrate will continue to use this path.

var _ types.DBIterator = (*ascendingIterator)(nil)

// ascendingIterator is the legacy iterator. Versions of a logical key sort
// oldest-first on disk, so finding the visible version for a target height
// requires a SeekLT(version+1) dance rather than a cheap First().
type ascendingIterator struct {
	source             *pebble.Iterator
	prefix, start, end []byte
	version            int64
	valid              bool
	reverse            bool
	iterationCount     int64
	storeKey           string

	closeSync sync.Once
}

func newAscendingIterator(src *pebble.Iterator, prefix, mvccStart, mvccEnd []byte, version int64, earliestVersion int64, reverse bool, storeKey string) *ascendingIterator {
	_ = "STUB: not implemented"
	// Return invalid iterator if requested iterator height is lower than earliest version after pruning
	return nil
}

// move the underlying PebbleDB iterator to the first key

// XXX: This should not happen as that would indicate we have a malformed MVCC key.

// We need to check whether initial key iterator visits has a version <= requested version
// If larger version, call next to find another key which does

// If version is less, seek to the largest version of that key <= requested iterator version
// It is guaranteed this won't move the iterator to a key that is invalid since
// curKeyVersionDecoded <= requested iterator version, so there exists at least one version of currKey SeekLT may move to

// Make sure we skip to the next key if the current one is tombstone
// Only check if iterator is still valid after the seek/next operations above

// Domain returns the domain of the iterator. The caller must not modify the
// return values.
func (itr *ascendingIterator) Domain() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (itr *ascendingIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

func (itr *ascendingIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC value.

func (itr *ascendingIterator) nextForward() { _ = "STUB: not implemented"; return }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

// First move the iterator to the next prefix, which may not correspond to the
// desired version for that key, e.g. if the key was written at a later version,
// so we seek back to the latest desired version, s.t. the version is <= itr.version.

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

// the next key must have itr.prefix as the prefix

// Move the iterator to the closest version to the desired version, so we
// append the current iterator key to the prefix and seek to that key.

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

// There exists cases where the SeekLT() call moved us back to the same key
// we started at, so we must move to next key, i.e. two keys forward.

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

// We need to verify that every Next call either moves the iterator to a key whose version
// is less than or equal to requested iterator version, or exhausts the iterator

// If iterator is at a entry whose version is higher than requested version, call nextForward again

// The cursor might now be pointing at a key/value pair that is tombstoned.
// If so, we must move the cursor.

func (itr *ascendingIterator) nextReverse() { _ = "STUB: not implemented"; return }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

// First move the iterator to the next prefix, which may not correspond to the
// desired version for that key, e.g. if the key was written at a later version,
// so we seek back to the latest desired version, s.t. the version is <= itr.version.

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

// the next key must have itr.prefix as the prefix

// Move the iterator to the closest version to the desired version, so we
// append the current iterator key to the prefix and seek to that key.

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

// We need to verify that every Next call either moves the iterator to a key whose version
// is less than or equal to requested iterator version, or exhausts the iterator

// If iterator is at a entry whose version is higher than requested version, call nextReverse again

// The cursor might now be pointing at a key/value pair that is tombstoned.
// If so, we must move the cursor.

func (itr *ascendingIterator) Next() { _ = "STUB: not implemented"; return }

func (itr *ascendingIterator) Valid() bool {
	_ = "STUB: not implemented"
	// once invalid, forever invalid
	return false
}

// if source has error, consider it invalid

// if key is at the end or past it, consider it invalid

func (itr *ascendingIterator) Error() error { _ = "STUB: not implemented"; return nil }

func (itr *ascendingIterator) Close() error { _ = "STUB: not implemented"; return nil }

// Record the number of iterations performed by this iterator

func (itr *ascendingIterator) assertIsValid() { _ = "STUB: not implemented"; return }

// cursorTombstoned checks if the current cursor is pointing at a key/value pair
// that is tombstoned. If the cursor is tombstoned, <true> is returned, otherwise
// <false> is returned. In the case where the iterator is valid but the key/value
// pair is tombstoned, the caller should call Next(). Note, this method assumes
// the caller assures the iterator is valid first!
func (itr *ascendingIterator) cursorTombstoned() bool { _ = "STUB: not implemented"; return false }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC value.

// If the tombstone suffix is empty, we consider this a zero value and thus it
// is not tombstoned.

// If the tombstone suffix is non-empty and greater than the target version,
// the value is not tombstoned.
