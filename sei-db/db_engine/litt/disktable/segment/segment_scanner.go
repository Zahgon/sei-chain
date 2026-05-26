package segment

import (
	"log/slog"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
)

// scanDirectories scans directories for segment files and returns a map of metadata, key, and value files.
// Also returns a list of garbage files that should be deleted. Does not do anything to files with unrecognized
// extensions.
func scanDirectories(logger *slog.Logger, segmentPaths []*SegmentPath) (
	metadataFiles map[uint32]string,
	keyFiles map[uint32]string,
	valueFiles map[uint32][]string,
	garbageFiles []string,
	highestSegmentIndex uint32,
	lowestSegmentIndex uint32,
	err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, 0, 0, nil
}

// key is the file's segment index, value is the file's path

// No segments found, fix the index.

// diagnoseMissingFile decides what to do with specific missing files. If the segment is either the segment
// with the lowest index or the segment with the highest index, it is possible for files to be missing due to
// non-catastrophic reasons (i.e. a crash during cleanup). If the segment is neither the lowest nor highest segment,
// then missing files signal non-recoverable DB corruption, and an error is returned.
func diagnoseMissingFile(
	logger *slog.Logger,
	index uint32,
	lowestFileIndex uint32,
	highestFileIndex uint32,
	fileType string,
	damagedSegments map[uint32]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// This can happen if we crash while creating a new segment. Recoverable.

// This can happen when deleting the oldest segment. Recoverable.

// Database is missing internal files. Catastrophic failure.

// lookForMissingFiles ensures that all files that should be present are actually present. Returns an error
// if files are missing in a way that cannot be recovered. If recoverable, returns a list of orphaned files.
// An "orphaned file" is defined as a file on disk for a segment that is missing one or more of its files.
// For example, if a segment has a metadata file but is missing its key file, the metadata file is considered orphaned.
func lookForMissingFiles(
	logger *slog.Logger,
	lowestSegmentIndex uint32,
	highestSegmentIndex uint32,
	metadataFiles map[uint32]string,
	keyFiles map[uint32]string,
	valueFiles map[uint32][]string,
	fsync bool,
) (orphanedFiles []string, damagedSegments map[uint32]struct{}, error error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Special case, only happens when starting a table from scratch.
// Files aren't actually missing, so no need to log anything.

// Check for missing metadata file.

// Check for missing key file.

// Check for missing value files (there should be exactly one value file per shard).

// If the metadata file is missing but we haven't yet returned an error, all of the value files
// are automatically considered orphaned.

// We need to know the sharding factor to check for missing value files.

// Catalogue the shards we do have.

// Check that we have each shard.

// If we are missing a file in this segment, all other files in the segment are considered orphaned.

// deleteOrphanedFiles deletes any files that are in the orphan set.
func deleteOrphanedFiles(logger *slog.Logger, orphanedFiles []string) error {
	_ = "STUB: not implemented"
	return nil
}

// linkSegments links together adjacent segments via SetNextSegment().
func linkSegments(lowestSegmentIndex uint32, highestSegmentIndex uint32, segments map[uint32]*Segment) error {
	_ = "STUB: not implemented"
	return nil
}

// Only one segment, nothing to link. This is checked explicitly to avoid 0-1 underflow.

// GatherSegmentFiles scans a directory for segment files and loads them into memory.
func GatherSegmentFiles(
	logger *slog.Logger,
	errorMonitor *util.ErrorMonitor,
	segmentPaths []*SegmentPath,
	snapshottingEnabled bool,
	now time.Time,
	cleanOrphans bool,
	fsync bool,
) (lowestSegmentIndex uint32, highestSegmentIndex uint32, segments map[uint32]*Segment, err error) {
	_ = "STUB: not implemented"

	// Scan the root directories for segment files.
	return 0, 0, nil, nil
}

// Delete any garbage files. Ignore files with unrecognized extensions.

// Check for missing files.

// Clean up any orphaned segment files.

// Adjust the segment range to exclude orphaned segments.

// Load all healthy segments.

// Stitch together the segments.
