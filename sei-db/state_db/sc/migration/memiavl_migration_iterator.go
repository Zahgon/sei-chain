package migration

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/memiavl"
)

// MemiavlMigrationIterator is a MigrationIterator that walks a memiavl.DB.
//
// The set of trees to migrate is fixed at construction time. Trees added
// to the DB after construction are not migrated; a tree that was present
// at construction but later disappears causes NextBatch to return an
// error.
//
// Callers can restrict the iterator to a subset of stores via the
// storesToMigrate argument; stores present in the DB but not in that
// whitelist are left untouched and never appear in a returned batch.
//
// The reserved MigrationStore tree is always excluded, even if listed
// in storesToMigrate: it holds migration metadata owned by
// MigrationManager and is not eligible for migration.
type MemiavlMigrationIterator struct {
	db        *memiavl.DB
	treeNames []string
	treeIdx   int
	boundary  MigrationBoundary
}

var _ MigrationIterator = (*MemiavlMigrationIterator)(nil)

// NewMemiavlMigrationIterator creates a MemiavlMigrationIterator positioned at
// the start of the given DB (boundary defaults to MigrationBoundaryNotStarted).
//
// storesToMigrate restricts the set of trees the iterator will walk:
//   - nil or empty: every tree in the DB is migrated.
//   - non-empty: only the listed trees are migrated; unlisted trees are
//     skipped entirely.
//
// Entries in storesToMigrate that do not correspond to an existing tree
// in the DB are silently ignored, so callers can pass a stable store
// list without worrying about trees that happened to be absent at open.
// The reserved MigrationStore tree is always filtered out even if listed.
func NewMemiavlMigrationIterator(
	// The DB to iterate.
	db *memiavl.DB,
	// The stores to iterate+migrate. If empty, all stores will be migrated.
	storesToMigrate []string,
) *MemiavlMigrationIterator {
	_ = "STUB: not implemented"
	return nil
}

func (m *MemiavlMigrationIterator) SetBoundary(boundary MigrationBoundary) {
	_ = "STUB: not implemented"
	return
}

func (m *MemiavlMigrationIterator) NextBatch(size int) ([]ValueToMigrate, MigrationBoundary, error) {
	_ = "STUB: not implemented"
	return nil, *new(MigrationBoundary), nil
}

// tree.Iterator's start bound is inclusive, so append a 0x00
// byte to get a start key strictly greater than boundary.Key().

// All trees fully drained; this was the final batch. Report
// Complete eagerly so the caller can finalize in the same step.

// computeStartTreeIndex returns the index of the first tree that may contain
// unmigrated keys according to the given boundary.
func computeStartTreeIndex(treeNames []string, boundary MigrationBoundary) int {
	_ = "STUB: not implemented"
	return 0
}

// copyBytes returns a newly allocated copy of b, or nil if b is nil.
func copyBytes(b []byte) []byte { _ = "STUB: not implemented"; return nil }
