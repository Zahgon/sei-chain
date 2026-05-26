package segment

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

const (

	// MetadataFileExtension is the file extension for the metadata file.
	MetadataFileExtension = ".metadata"

	// MetadataSwapExtension is the file extension for the metadata swap file. This file is used to atomically update
	// the metadata file by doing an atomic rename of the swap file to the metadata file. If this file is ever
	// present when the database first starts, it is an artifact of a crash during a metadata update, and should be
	// deleted.
	MetadataSwapExtension = MetadataFileExtension + util.SwapFileExtension

	// V3MetadataSize is the size of the metadata file at LatestSegmentVersion (ShardedAddressSegmentVersion).
	// Layout:
	//   - 4 bytes for version
	//   - 1 byte for the sharding factor
	//   - 8 bytes for lastValueTimestamp
	//   - 4 bytes for keyCount
	//   - 1 byte for sealed
	V3MetadataSize = 18
)

// metadataFile contains metadata about a segment. This file contains metadata about the data segment, such as
// serialization version and the lastValueTimestamp when the file was sealed.
type metadataFile struct {
	// The segment index. This value is encoded in the file name.
	index uint32

	// The serialization version for this segment, used to permit smooth data migrations.
	// This value is encoded in the file.
	segmentVersion SegmentVersion

	// The sharding factor for this segment. This value is encoded in the file.
	shardingFactor uint8

	// The time when the last value was written into the segment, in nanoseconds since the epoch. A segment can
	// only be deleted when all values within it are expired, and so we only need to keep track of the
	// lastValueTimestamp of the last value (which always expires last). This value is irrelevant if the segment is
	// not yet sealed. This value is encoded in the file.
	lastValueTimestamp uint64

	// The number of keys in the segment. This value is undefined if the segment is not yet sealed.
	// This value is encoded in the file.
	keyCount uint32

	// If true, the segment is sealed and no more data can be written to it. If false, then data can still be written
	// to this segment. This value is encoded in the file.
	sealed bool

	// Path data for the segment file. This information is not serialized in the metadata file.
	segmentPath *SegmentPath

	// If true, then use fsync to make metadata updates atomic. Should always be true in production, but can be
	// set to false in tests to speed up unit tests. Not serialized to the file.
	fsync bool
}

// createMetadataFile creates a new metadata file. When this method returns, the metadata file will
// be durably written to disk.
func createMetadataFile(
	index uint32,
	shardingFactor uint8,
	path *SegmentPath,
	fsync bool,
) (*metadataFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadMetadataFile loads the metadata file from disk, looking in the given parent directories until it finds the file.
// If the file is not found, it returns an error.
func loadMetadataFile(index uint32, segmentPaths []*SegmentPath, fsync bool) (*metadataFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // path within segment directory

// MetadataFileExtension is the file extension for the metadata file. Metadata file names have the form "X.metadata",
// where X is the segment index.
func getMetadataFileIndex(fileName string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gosec // segment index fits uint32

// Size returns the size of the metadata file in bytes.
func (m *metadataFile) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Name returns the file name for this metadata file.
func (m *metadataFile) name() string { _ = "STUB: not implemented"; return "" }

// Path returns the full path to this metadata file.
func (m *metadataFile) path() string { _ = "STUB: not implemented"; return "" }

// Seal seals the segment. This action will atomically write the metadata file to disk one final time,
// and should only be performed when all data that will be written to the key/value files has been made durable.
func (m *metadataFile) seal(now time.Time, keyCount uint32) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // wall-clock nanos non-negative

// serialize serializes the metadata file to a byte array.
func (m *metadataFile) serialize() []byte { _ = "STUB: not implemented"; return nil }

// deserialize deserializes the metadata file from a byte array.
func (m *metadataFile) deserialize(data []byte) error { _ = "STUB: not implemented"; return nil }

// write atomically writes the metadata file to disk.
func (m *metadataFile) write() error { _ = "STUB: not implemented"; return nil }

// snapshot creates a hard link to the file in the snapshot directory, and a soft link to the hard linked file in the
// soft link directory. Requires that the file is sealed and that snapshotting is enabled.
func (m *metadataFile) snapshot() error { _ = "STUB: not implemented"; return nil }

// delete deletes the metadata file from disk. If the file is a snapshot (i.e., a symlink), this method will also
// delete the actual file that the symlink points to.
func (m *metadataFile) delete() error { _ = "STUB: not implemented"; return nil }
