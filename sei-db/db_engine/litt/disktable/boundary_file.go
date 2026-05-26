package disktable

// The name of the file that defines the lower bound of a LittDB snapshot directory.
const LowerBoundFileName = "lower-bound.txt"

// The name of the file that defines the upper bound of a LittDB snapshot directory.
const UpperBoundFileName = "upper-bound.txt"

// BoundaryType is an enum that describes the type of boundary file.
type BoundaryType bool

const (
	// A boundary file that defines the lowest valid segment index in a snapshot directory.
	LowerBound BoundaryType = true
	// A boundary file that defines the highest valid segment index in a snapshot directory.
	UpperBound BoundaryType = false
)

type BoundaryFile struct {
	// The type of this boundary file.
	boundaryType BoundaryType

	// The parent directory where this file is stored.
	parentDirectory string

	// If true, then the boundary is defined, otherwise it is undefined.
	// If undefined, the boundary index should be considered invalid.
	defined bool

	// The segment index of the boundary. Describes a lower/upper segment index. If this is a lower bound file,
	// it describes the lowest segment index that is valid within the snapshot directory (inclusive). If this is
	// an upper bound file, it describes the highest segment index that is valid within the snapshot directory
	// (also inclusive).
	boundaryIndex uint32
}

// LoadBoundaryFile loads a boundary file from the specified parent directory. If the boundary file does not exist,
// then this method returns an object that can be used to create a new boundary file at the specified path (i.e. by
// calling Write() or Update()).
func LoadBoundaryFile(boundaryType BoundaryType, parentDirectory string) (*BoundaryFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Atomically update the value of the boundary file.
func (b *BoundaryFile) Update(newBoundary uint32) error { _ = "STUB: not implemented"; return nil }

// Get the file name of the boundary file.
func (b *BoundaryFile) Name() string { _ = "STUB: not implemented"; return "" }

// Get the full path where the boundary file is stored.
func (b *BoundaryFile) Path() string { _ = "STUB: not implemented"; return "" }

// Serialize the boundary file to a byte slice.
func (b *BoundaryFile) serialize() []byte { _ = "STUB: not implemented"; return nil }

// Serialize the boundary file to a byte slice. Since end users may interact with this file,
// serialize in a human-readable format.

func (b *BoundaryFile) deserialize(data []byte) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // boundary index fits uint32

// Write the boundary file to disk.
func (b *BoundaryFile) Write() error { _ = "STUB: not implemented"; return nil }

// fsync is not necessary, in an advent of a crash the boundary files get repaired

// Returns true if this boundary file is defined. If undefined, it means that the boundary index is invalid
// and should not be used.
func (b *BoundaryFile) IsDefined() bool { _ = "STUB: not implemented"; return false }

// Get the boundary index described by this file.
//
// If this is a lower bound, then it describes the highest segment index in a snapshot directory that has been garbage
// collected. As a result, LittDB will not snapshot any segments with this index or lower.
//
// If this is an upper bound, then it describes the highest segment index that LittDB has fully taken a snapshot of.
// External processes using the snapshot should ignore any segment with an index greater than this.
func (b *BoundaryFile) BoundaryIndex() uint32 { _ = "STUB: not implemented"; return 0 }
