package pebbledb

import (
	"github.com/cockroachdb/pebble/v2"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// pebbleIterator implements db_engine.Iterator using PebbleDB.
// Key/Value follow Pebble's zero-copy semantics; see db_engine.Iterator contract.
type pebbleIterator struct {
	it *pebble.Iterator
}

var _ types.KeyValueDBIterator = (*pebbleIterator)(nil)

func (pi *pebbleIterator) First() bool          { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) Last() bool           { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) Valid() bool          { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) SeekGE(k []byte) bool { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) SeekLT(k []byte) bool { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) Next() bool           { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) NextPrefix() bool     { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) Prev() bool           { _ = "STUB: not implemented"; return false }
func (pi *pebbleIterator) Key() []byte          { _ = "STUB: not implemented"; return nil }
func (pi *pebbleIterator) Value() []byte        { _ = "STUB: not implemented"; return nil }
func (pi *pebbleIterator) Error() error         { _ = "STUB: not implemented"; return nil }
func (pi *pebbleIterator) Close() error         { _ = "STUB: not implemented"; return nil }
