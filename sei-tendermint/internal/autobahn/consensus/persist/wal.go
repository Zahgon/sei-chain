package persist

import (
	dbwal "github.com/sei-protocol/sei-chain/sei-db/wal"
)

// codec is the marshal/unmarshal pair needed to store T in a WAL.
// protoutils.Conv[T, P] satisfies this interface automatically.
type codec[T any] interface {
	Marshal(T) []byte
	Unmarshal([]byte) (T, error)
}

// indexedWAL wraps a WAL with monotonic index tracking and typed entries.
// Callers map domain-specific indices (BlockNumber, RoadIndex) to WAL
// indices via Count() and firstIdx. Not safe for concurrent use.
//
// INVARIANT: firstIdx <= nextIdx (enforced by openIndexedWAL, Write,
// TruncateBefore, and TruncateAll). Count() relies on this for unsigned
// subtraction safety.
type indexedWAL[T any] struct {
	wal      *dbwal.WAL[T]
	firstIdx uint64 // WAL index of the oldest entry; == nextIdx when empty
	nextIdx  uint64 // WAL index that the next Write will be assigned
}

// openIndexedWAL creates (or opens) a WAL in dir with synchronous, unbatched
// writes and fsync. The prune anchor (persisted via A/B files) is the
// crash-recovery watermark, but fsync on the WAL provides additional
// durability for entries not yet covered by the anchor.
// Initializes index tracking from the WAL's stored offsets so the caller can
// immediately Write, ReadAll, or TruncateBefore.
func openIndexedWAL[T any](dir string, codec codec[T]) (*indexedWAL[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// synchronous writes
// no batching

// CRITICAL: AllowEmpty must be true in the config above. With AllowEmpty,
// an empty log reports first > last (e.g. first=1, last=0 for a brand-new
// log, or first=N+1, last=N after TruncateAll). Without AllowEmpty, an
// empty log returns (0, 0), and the formula below would give Count() == 1.
// A non-empty log always has first <= last with first >= 1. In both cases,
// setting firstIdx = first and nextIdx = last + 1 yields Count() == 0
// when empty.

// Write appends entry to the WAL, advancing nextIdx.
func (w *indexedWAL[T]) Write(entry T) error { _ = "STUB: not implemented"; return nil }

// TruncateBefore reads the entry at walIdx, passes it to verify, and — if
// verify returns nil — removes all entries before walIdx. The verify callback
// lets callers assert that the WAL index maps to the expected domain object
// before a destructive operation.
//
// TruncateBefore assumes contiguous appending and removal. Callers that need
// to fast-forward past gaps should use TruncateAll directly.
// walIdx == nextIdx removes all entries (no entry to verify).
// walIdx > nextIdx is an error (gap in contiguous sequence).
func (w *indexedWAL[T]) TruncateBefore(walIdx uint64, verify func(T) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify the surviving entry when one exists.

// TruncateWhile scans entries from the front and removes all leading entries
// for which the predicate returns true. Stops at the first entry where the
// predicate returns false (that entry and all after it are kept).
// If the predicate returns true for all entries, the WAL is emptied.
func (w *indexedWAL[T]) TruncateWhile(pred func(T) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadAll returns all entries in the WAL. Returns nil if empty.
func (w *indexedWAL[T]) ReadAll() ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

// FirstIdx returns the WAL index of the oldest entry.
// Equal to nextIdx when the WAL is empty.
func (w *indexedWAL[T]) FirstIdx() uint64 {
	_ = "STUB: not implemented"

	// Count returns the number of entries in the WAL.
	// Empty when firstIdx == nextIdx (both after fresh open and after TruncateAll).
	return 0
}

func (w *indexedWAL[T]) Count() uint64 { _ = "STUB: not implemented"; return 0 }

// TruncateAll removes all entries from the WAL, leaving it empty for new writes.
// The underlying index counter is preserved (next Write continues from where
// it left off); firstIdx is advanced to nextIdx so Count() == 0.
// Used when all entries are stale (e.g. the prune anchor advanced past
// everything persisted).
func (w *indexedWAL[T]) TruncateAll() error { _ = "STUB: not implemented"; return nil }

// TruncateAfter removes all entries >= walIdx (keeps entries < walIdx).
// If walIdx <= firstIdx, all entries are removed.
// If walIdx >= nextIdx, this is a no-op.
func (w *indexedWAL[T]) TruncateAfter(walIdx uint64) error { _ = "STUB: not implemented"; return nil }

// Close shuts down the underlying WAL.
func (w *indexedWAL[T]) Close() error { _ = "STUB: not implemented"; return nil }
