package segment

// The name of the directory where segment files are stored. The segment directory is created at
// "$STORAGE_PATH/$TABLE_NAME/segments". Each table has at least one segment directory. Tables may
// have multiple segment directories if more than one path is provided to Litt.Config.Paths.
const SegmentDirectory = "segments"

// The name of the directory where hard links to segment files are stored for snapshotting (if enabled).
// The hard link directory is created at "$STORAGE_PATH/$TABLE_NAME/snapshot".
const HardLinkDirectory = "snapshot"

// SegmentPath encapsulates various file paths utilized by segment files.
type SegmentPath struct {
	// The directory where the segment file is stored.
	segmentDirectory string
	// If snapshotting is enabled, the directory where a Snapshot will put a hard link to the segment file.
	// An empty string if snapshotting is not enabled.
	hardlinkPath string
	// If snapshotting is enabled, the directory where a Snapshot will put a soft link to the hard link of a
	// segment file. An empty string if snapshotting is not enabled.
	softlinkPath string
}

// NewSegmentPath creates a new SegmentPath. Each segment file's location on disk is determined by a SegmentPath object.
//
// The storageRoot is a location where LittDB is storing data, i.e. one of the paths from Litt.Config.Paths.
//
// softlinkRoot will be an empty string if snapshotting is not enabled, or a path to the root directory where
// Snapshot soft links will be created. The presence (or absence) of this path is used by LittDB to
// determine if snapshotting is enabled.
//
// The tableName is the name of the table that owns the segment file.
func NewSegmentPath(
	storageRoot string,
	softlinkRoot string,
	tableName string,
) (*SegmentPath, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildSegmentPaths creates a list of SegmentPath objects for each storage root provided.
func BuildSegmentPaths(
	storageRoots []string,
	softlinkRoot string,
	tableName string,
) ([]*SegmentPath, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SegmentDirectory returns the parent directory where segment files are stored.
func (p *SegmentPath) SegmentDirectory() string { _ = "STUB: not implemented"; return "" }

// HardlinkPath returns the path where hard links to segment files will be created for snapshotting.
func (p *SegmentPath) HardlinkPath() string { _ = "STUB: not implemented"; return "" }

// SoftlinkPath returns the path where soft links to hard links of segment files will be created for snapshotting.
func (p *SegmentPath) SoftlinkPath() string { _ = "STUB: not implemented"; return "" }

// snapshottingEnabled checks if snapshotting is enabled.
func (p *SegmentPath) snapshottingEnabled() bool { _ = "STUB: not implemented"; return false }

// MakeDirectories creates the necessary directories described by the SegmentPath if they do not already exist.
func (p *SegmentPath) MakeDirectories(fsync bool) error { _ = "STUB: not implemented"; return nil }

// Snapshot creates a hard link to the file in the Snapshot directory, and a symlink to that hard link in the soft link
// directory. The fileName should just be the name of the file, not its full path. The file is expected to be in the
// segmentDirectory.
func (p *SegmentPath) Snapshot(fileName string) error { _ = "STUB: not implemented"; return nil }
