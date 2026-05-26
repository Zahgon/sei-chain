package flatkv

import (
	seidbtypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

var _ Iterator = (*sequentialIterator)(nil)

// sequentialIterator iterates through a slice of DBs one at a time.
// It fully drains the current DB before moving to the next.
type sequentialIterator struct {
	dbs   []seidbtypes.KeyValueDB
	dbIdx int // index into dbs for the current DB
	iter  seidbtypes.KeyValueDBIterator
	err   error
}

// openCurrent opens an iterator on dbs[dbIdx]. Returns false if no more DBs.
func (s *sequentialIterator) openCurrent() bool { _ = "STUB: not implemented"; return false }

// advanceDB closes the current iterator and moves to the next DB,
// positioning at the first non-meta key. Returns true if positioned.
// If the current iterator has an error, it is captured and iteration stops.
func (s *sequentialIterator) advanceDB() bool { _ = "STUB: not implemented"; return false }

func skipMeta(it seidbtypes.KeyValueDBIterator) { _ = "STUB: not implemented"; return }

func (s *sequentialIterator) Domain() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (s *sequentialIterator) Valid() bool { _ = "STUB: not implemented"; return false }

func (s *sequentialIterator) Error() error { _ = "STUB: not implemented"; return nil }

func (s *sequentialIterator) Close() error { _ = "STUB: not implemented"; return nil }

func (s *sequentialIterator) First() bool { _ = "STUB: not implemented"; return false }

func (s *sequentialIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (s *sequentialIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (s *sequentialIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// Unsupported positioning methods — not needed for forward-only scanning.

func (s *sequentialIterator) Last() bool         { _ = "STUB: not implemented"; return false }
func (s *sequentialIterator) SeekGE([]byte) bool { _ = "STUB: not implemented"; return false }
func (s *sequentialIterator) SeekLT([]byte) bool { _ = "STUB: not implemented"; return false }
func (s *sequentialIterator) Prev() bool         { _ = "STUB: not implemented"; return false }

// RawGlobalIterator returns an iterator that walks each data DB sequentially
// in fixed order (account → code → storage → legacy). Within each DB the
// keys are returned in PebbleDB's natural order. Per-DB _meta/* keys are
// skipped. Pending writes are not visible. metadataDB is not included.
func (s *CommitStore) RawGlobalIterator() Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}
