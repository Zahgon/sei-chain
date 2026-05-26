package migration

// MockMigrationIterator is a MigrationIterator backed by an in-memory map.
// Useful as a test double and as a reference implementation for validating
// test logic independently of any real DB.
//
// The underlying Data map may be mutated between NextBatch calls. If
// autoRebuild is true, the iterator automatically re-flattens and repositions
// before each NextBatch call. Otherwise, call Rebuild manually after mutating.
//
// The reserved MigrationStore module is always excluded from the iteration
// output: it holds migration metadata owned by MigrationManager and is not
// eligible for migration.
type MockMigrationIterator struct {
	Data        map[string]map[string][]byte
	autoRebuild bool
	entries     []ValueToMigrate
	position    int
	boundary    MigrationBoundary
}

var _ MigrationIterator = (*MockMigrationIterator)(nil)

// NewMockMigrationIterator creates an iterator from the given data,
// positioned at the start (boundary defaults to MigrationBoundaryNotStarted).
// If autoRebuild is true, the iterator re-reads from Data before every
// NextBatch call, so external mutations are picked up automatically.
func NewMockMigrationIterator(data map[string]map[string][]byte, autoRebuild bool) *MockMigrationIterator {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockMigrationIterator) SetBoundary(boundary MigrationBoundary) {
	_ = "STUB: not implemented"
	return
}

// Rebuild re-flattens and re-sorts the Data map, then repositions the
// iterator so that the next NextBatch call resumes just past the current
// boundary. Call this after adding or removing entries from Data.
func (m *MockMigrationIterator) Rebuild() { _ = "STUB: not implemented"; return }

func (m *MockMigrationIterator) NextBatch(size int) ([]ValueToMigrate, MigrationBoundary, error) {
	_ = "STUB: not implemented"
	return nil, *new(MigrationBoundary), nil
}

// This batch drained the iterator; report Complete eagerly so
// the caller can finalize in the same step.

// flattenAndSort converts a nested map into a sorted slice of ValueToMigrate,
// ordered lexicographically by (ModuleName, Key). The MigrationStore module
// is skipped: its contents are migration metadata, not payload data.
func flattenAndSort(data map[string]map[string][]byte) []ValueToMigrate {
	_ = "STUB: not implemented"
	return nil
}

// computeStartPosition returns the index of the first entry that has not yet
// been migrated according to the given boundary.
func computeStartPosition(entries []ValueToMigrate, boundary MigrationBoundary) int {
	_ = "STUB: not implemented"
	return 0
}
