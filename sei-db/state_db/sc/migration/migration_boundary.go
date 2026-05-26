package migration

//nolint:godot
/*
                        <-- Lexographically ordered keys -->

      keys in this region migrated           keys in this region not migrated
    |--------------------------------------|--------------------------------------|
    ^                                      ^                                      ^
    |                                      |                                      |
  first key                       migration boundary                          last key
                                --> moves each block -->
*/

// headerSize is the size of the fixed-length prefix of an in-progress
// serialized boundary: 1 status byte + 4-byte big-endian module-name length.
const headerSize = 5

// Defines the boundary between migrated and unmigrated keys.
type MigrationBoundary struct {
	// The status of the migration.
	status MigrationStatus
	// The module name of the highest migrated key.
	moduleName string
	// The highest migrated key.
	key []byte
}

var (
	// A boundary for a migration that has not yet started. No key is considered migrated.
	MigrationBoundaryNotStarted = MigrationBoundary{
		status: MigrationNotStarted,
	}
	// A boundary for a migration that has completed. All keys are considered migrated.
	MigrationBoundaryComplete = MigrationBoundary{
		status: MigrationComplete,
	}
)

// Create a new migration boundary with the given key.
//
// The key slice is stored by reference; callers must not mutate it after
// this call.
func NewMigrationBoundary(
	// The module name of the highest migrated key.
	moduleName string,
	// The highest migrated key.
	key []byte,
) MigrationBoundary {
	_ = "STUB: not implemented"
	return *new(MigrationBoundary)
}

// Equals returns true if two boundaries represent the same migration state.
func (mb *MigrationBoundary) Equals(other MigrationBoundary) bool {
	_ = "STUB: not implemented"
	return false
}

// Status returns the lifecycle status of the migration.
func (mb *MigrationBoundary) Status() MigrationStatus {
	_ = "STUB: not implemented"

	// ModuleName returns the module name of the highest migrated key.
	return *new(MigrationStatus)
}

func (mb *MigrationBoundary) ModuleName() string { _ = "STUB: not implemented"; return "" }

// Key returns the highest migrated key. The returned slice aliases the
// boundary's internal state and must not be mutated.
func (mb *MigrationBoundary) Key() []byte {
	_ = "STUB: not implemented"

	// String returns a human-readable representation of the boundary.
	// For in-progress boundaries the key is hex-encoded so that arbitrary binary
	// keys round-trip unambiguously.
	//
	// Format: "MigrationBoundary{status=<status>, module=<name>, key=<hex>}"
	return nil
}

func (mb *MigrationBoundary) String() string { _ = "STUB: not implemented"; return "" }

// Checks to see if a key has been migrated yet. Compares key against boundary, returning true if key is to the left
// (or equal to) the boundary. Returns false if the key is to the right of the boundary.
func (mb *MigrationBoundary) IsMigrated(moduleName string, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Serialize encodes the boundary as a byte slice. The returned slice is
// freshly allocated and independent of the boundary.
//
// For notStarted/complete: [status byte]
// For inProgress: [status byte] [4-byte BE moduleName length] [moduleName] [key]
func (mb *MigrationBoundary) Serialize() []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// DeserializeMigrationBoundary decodes a byte slice produced by Serialize.
// The returned boundary owns its key and is independent of the input slice.
func DeserializeMigrationBoundary(data []byte) (MigrationBoundary, error) {
	_ = "STUB: not implemented"
	return *new(MigrationBoundary), nil
}
