package main

import (
	"log/slog"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/disktable/segment"
	"github.com/urfave/cli/v2"
)

// pruneCommand can be used to remove data from a LittDB instance/snapshot.
func pruneCommand(ctx *cli.Context) error { _ = "STUB: not implemented"; return nil }

// prune deletes data from a littDB database/snapshot.
func prune(logger *slog.Logger, sources []string, allowedTables []string, maxAgeSeconds uint64, fsync bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Forbid touching tables in active use.

// Determine which tables to prune.

// Prune each table.

// pruneTable performs offline garbage collection on a LittDB database/snapshot.
func pruneTable(
	logger *slog.Logger,
	sources []string,
	tableName string,
	maxAgeSeconds uint64,
	fsync bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Determine if we are working on the snapshot directory (i.e. the directory with symlinks to the segments).

// If we are dealing with a snapshot, respect the snapshot upper bound specified by LittDB.

// Delete old segments.

//nolint:gosec // CLI flag bounded
// We've pruned all segments that we can.

// Wait for deletion to complete.

// This is a snapshot. Write a lower bound file to tell the DB not to re-snapshot files than have been pruned.

// If we are doing GC on a table that isn't a snapshot, then we need to delete the snapshots/keymap
// for the table. The DB will automatically rebuild the snapshots directory & keymap on the next startup.

// Updates the lower bound file after segments have been deleted.
func writeLowerBoundFile(snapshotRoot string, tableName string, deletedSegments []*segment.Segment) error {
	_ = "STUB: not implemented"
	return nil
}

// No segments were deleted, no need to write a lower bound file.

// deletes the snapshot directories in all sources for the given table
func deleteSnapshots(sources []string, tableName string) error {
	_ = "STUB: not implemented"
	return nil
}
