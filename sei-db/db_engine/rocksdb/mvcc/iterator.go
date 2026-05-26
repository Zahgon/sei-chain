//go:build rocksdbBackend
// +build rocksdbBackend

package mvcc

import (
	"sync"

	"github.com/linxGnu/grocksdb"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

var _ types.DBIterator = (*iterator)(nil)

type iterator struct {
	source             *grocksdb.Iterator
	readOpts           *grocksdb.ReadOptions
	prefix, start, end []byte
	version            int64
	reverse            bool
	invalid            bool
	closeOnce          sync.Once
}

func NewRocksDBIterator(source *grocksdb.Iterator, readOpts *grocksdb.ReadOptions, prefix, start, end []byte, version int64, earliestVersion int64, reverse bool) *iterator {
	_ = "STUB: not implemented"
	// Return invalid iterator if requested iterator height is lower than earliest version after pruning
	return nil
}

// end or after key

// Domain returns the domain of the iterator. The caller must not modify the
// return values.
func (itr *iterator) Domain() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (itr *iterator) Valid() bool {
	_ = "STUB: not implemented"
	// once invalid, forever invalid
	return false
}

// if source has error, consider it invalid

// if source is invalid, consider it invalid

// if key is at the end or past it, consider it invalid

func (itr *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (itr *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (itr *iterator) Next() { _ = "STUB: not implemented"; return }

func (itr *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (itr *iterator) Close() error { _ = "STUB: not implemented"; return nil }

func (itr *iterator) assertIsValid() { _ = "STUB: not implemented"; return }
