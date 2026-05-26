package main

import (
	"log/slog"
	"sync/atomic"

	"github.com/urfave/cli/v2"
)

// rebaseCommand is the command to rebase a LittDB database.
func rebaseCommand(ctx *cli.Context) error { _ = "STUB: not implemented"; return nil }

// rebase moves LittDB database files from one location to another (locally). This function is idempotent. If it
// crashes part of the way through, just run it again and it will continue where it left off.
func rebase(
	logger *slog.Logger,
	sources []string,
	destinations []string,
	preserveOriginal bool,
	fsync bool,
	verbose bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore non-existent source paths. They could have been deleted by a prior run of this command.

// Don't immediately take a lock on the source directories. Each source directory will be locked individually
// before its data is transferred. Because source directories are deleted after their data is transferred,
// it is inconvenient to hold the locks in this outer scope (since we need to release the lock to
// delete the directory).

// Acquire locks on all destination directories.

// Figure out which directories are going away. We will need to transfer their data to new locations.

// If the source directory is not in the destination set, it is going away.

// If any of the segment files are symlinks, that means that we are dealing with a snapshot.

// For each directory that is going away, transfer its data to the new destination.

// Get a count of the segment files in the source directories.
// Also checks whether any of the segment files are symlinks.
func countSegmentFiles(sources []string) (count int64, symlinkFound bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// Walk the file tree to find all files ending with .metadata, .keys, or .values.

// Skip directories

// Ignore "table.metadata" files, as they are not segment files.

// Check if the file is a segment file.

// transfers all data in a directory to the specified destinations.
func transferDataInDirectory(
	logger *slog.Logger,
	source string,
	destinations []string,
	preserveOriginal bool,
	fsync bool,
	verbose bool,
	totalSegmentFileCount int64,
	segmentFileCount *atomic.Int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire a lock on the source directory.

// double release is a no-op

// Transfer each table stored in this directory.

// Release the lock so we can delete the directory.

// Delete the directory.

func transferDataInTable(
	logger *slog.Logger,
	source string,
	tableName string,
	destinations []string,
	preserveOriginal bool,
	fsync bool,
	verbose bool,
	totalSegmentFileCount int64,
	segmentFileCount *atomic.Int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Once all data in a table is transferred, delete the table directory.

// deleteBoundaryFiles deletes the boundary files for a table. Only will be present if the source
// directory contains symlink snapshots.
func deleteBoundaryFiles(logger *slog.Logger, source string, tableName string, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

// delete the old snapshot directory for a table. This will be reconstructed the next time the DB is loaded.
func deleteSnapshotDirectory(source string, tableName string) error {
	_ = "STUB: not implemented"
	return nil
}

// In the destination directories, create directories for the tables (if they don't exist).
func createDestinationTableDirectories(destinations []string, tableName string, fsync bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer the keymap (if it is present in the source directory).
func transferKeymap(
	source string,
	tableName string,
	destinations []string,
	preserveOriginal bool,
	fsync bool,
	verbose bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// transfers data in the segments/ directory
func transferSegmentData(
	source string,
	tableName string,
	destinations []string,
	preserveOriginal bool,
	fsync bool,
	verbose bool,
	totalSegmentFileCount int64,
	segmentFileCount *atomic.Int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Now that we've copied the segment files, we can delete the original directory.

// Transfer a single segment file (i.e. *.metadata, *.keys, *.values).
func transferSegmentFile(
	segmentName string,
	segmentFilePath string,
	tableName string,
	destinations []string,
	preserveOriginal bool,
	fsync bool,
	verbose bool,
	totalSegmentFileCount int64,
	segmentFileCount *atomic.Int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// transfers the table metadata file, if it is present.
func transferTableMetadata(
	source string,
	tableName string,
	destinations []string,
	preserveOriginal bool,
	fsync bool,
	verbose bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Determines the location where a file should be transferred given a list of options.
// This function is deterministic. This is important! If a rebase is interrupted, the
// second attempt should always transfer the file to the same location as the first attempt.
func determineDestination(source string, destinations []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
