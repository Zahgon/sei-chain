package main

import (
	"log/slog"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/disktable/segment"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/util"
	"github.com/urfave/cli/v2"
)

func pushCommand(ctx *cli.Context) error { _ = "STUB: not implemented"; return nil }

// push uses rsync to transfer LittDB data to the remote location(s)
func push(
	logger *slog.Logger,
	sources []string,
	destinations []string,
	user string,
	host string,
	port uint64,
	keyPath string,
	knownHosts string,
	deleteAfterTransfer bool,
	fsync bool,
	threads uint64,
	throttleMB float64,
	verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

// split bandwidth between workers

// Lock source files. It would be nice to also lock the remote directories, but that's tricky given that
// we are interacting with the remote machine via SSH and rsync.

// Create an SSH session to the remote host.

// Figure out which files are already present at the destination(s). Although these files may be partial, we always
// want to preserve any pre-existing arrangements of files at the destination(s).
//
// The returned map is a map from file name (e.g. 1234.metadata) to the destination path (e.g. /path/to/remote/dir).
func mapExistingFiles(
	destinations []string,
	tableName string,
	connection *util.SSHSession) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract the file name from the path.

// Push the data in a single table to the remote location(s).
func pushTable(
	logger *slog.Logger,
	tableName string,
	sources []string,
	destinations []string,
	connection *util.SSHSession,
	deleteAfterTransfer bool,
	fsync bool,
	throttleMB float64,
	threads uint64) error {
	_ = "STUB: not implemented"

	// Figure out where data currently exists at the destination(s). We don't want this operation to cause a file
	// to exist in multiple places.
	return nil
}

// Gather segment files to send.

// Special handling if we are transferring data from a snapshot.

// Ensure the remote segment directories exists.

// Used to limit rsync concurrency.

// Transfer the files.

// Wait for all rsyncs to complete.

// Check if there were any errors during the transfer.

// Now that we have transferred the files, we can delete them if requested.

// Deletes local segments after they have been successfully transferred to the remote destination(s).
func deleteLocalSegments(
	segments map[uint32]*segment.Segment,
	tableName string,
	isSnapshot bool,
	sources []string,
	highestSegmentIndex uint32) error {
	_ = "STUB: not implemented"

	// Delete the segments.
	return nil
}

// Wait for deletion to complete.

// If we are dealing with a snapshot, update the lower bound file.
