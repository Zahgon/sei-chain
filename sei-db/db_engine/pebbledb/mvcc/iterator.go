package mvcc

import (
	"sync"

	"github.com/cockroachdb/pebble/v2"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

var _ types.DBIterator = (*iterator)(nil)

// iterator implements the Iterator interface. It wraps a PebbleDB iterator
// with added MVCC key handling logic. The iterator will iterate over the key space
// in the provided domain for a given version. If a key has been written at the
// provided version, that key/value pair will be iterated over. Otherwise, the
// latest version for that key/value pair will be iterated over s.t. it's less
// than the provided version. The start key must not be empty.
type iterator struct {
	source             *pebble.Iterator
	prefix, start, end []byte
	version            int64
	valid              bool
	reverse            bool
	useDefaultComparer bool
	iterationCount     int64
	storeKey           string

	closeSync sync.Once
}

func newPebbleDBIterator(src *pebble.Iterator, prefix, mvccStart, mvccEnd []byte, version int64, earliestVersion int64, reverse bool, useDefaultComparer bool, storeKey string) *iterator {
	_ = "STUB: not implemented"
	// Return invalid iterator if requested iterator height is lower than earliest version after pruning
	return nil
}

// move the underlying PebbleDB iterator to the first key

// XXX: This should not happen as that would indicate we have a malformed MVCC key.

// Make sure we skip to the next key if the current one is tombstone
// Only check if iterator is still valid after the seek/next operations above

func (itr *iterator) seekVisibleVersionForKey(targetKey []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (itr *iterator) nextLogicalKey(currKey []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (itr *iterator) nextLogicalKeyByScan(currKey []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (itr *iterator) prevLogicalKey(currKey []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (itr *iterator) positionAtOrAfterKey(startKey []byte) { _ = "STUB: not implemented"; return }

func (itr *iterator) positionAtOrBeforeKey(startKey []byte) { _ = "STUB: not implemented"; return }

// Domain returns the domain of the iterator. The caller must not modify the
// return values.
func (itr *iterator) Domain() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (itr *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

func (itr *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC value.

func (itr *iterator) nextForward() { _ = "STUB: not implemented"; return }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

func (itr *iterator) nextReverse() { _ = "STUB: not implemented"; return }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC key.

func (itr *iterator) Next() { _ = "STUB: not implemented"; return }

func (itr *iterator) Valid() bool {
	_ = "STUB: not implemented"
	// once invalid, forever invalid
	return false
}

// if source has error, consider it invalid

// if key is at the end or past it, consider it invalid

func (itr *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (itr *iterator) Close() error { _ = "STUB: not implemented"; return nil }

// Record the number of iterations performed by this iterator

func (itr *iterator) assertIsValid() { _ = "STUB: not implemented"; return }

// cursorTombstoned checks if the current cursor is pointing at a key/value pair
// that is tombstoned. If the cursor is tombstoned, <true> is returned, otherwise
// <false> is returned. In the case where the iterator is valid but the key/value
// pair is tombstoned, the caller should call Next(). Note, this method assumes
// the caller assures the iterator is valid first!
func (itr *iterator) cursorTombstoned() bool { _ = "STUB: not implemented"; return false }

// XXX: This should not happen as that would indicate we have a malformed
// MVCC value.

// If the tombstone suffix is empty, we consider this a zero value and thus it
// is not tombstoned.

// If the tombstone suffix is non-empty and greater than the target version,
// the value is not tombstoned.

func (itr *iterator) DebugRawIterate() { _ = "STUB: not implemented"; return }

// The first key may not represent the desired target version, so move the
// cursor to the correct location.
