package main

import (
	"log/slog"
	"time"

	"github.com/urfave/cli/v2"
)

// TableInfo contains high level information about a table in LittDB.
type TableInfo struct {
	// The number of key-value pairs in the table.
	KeyCount uint64
	// The size of the table in bytes.
	Size uint64
	// If true, the table at the specified path is a snapshot of another table.
	IsSnapshot bool
	// The time when the oldest segment was sealed.
	OldestSegmentSealTime time.Time
	// The time when the newest segment was sealed.
	NewestSegmentSealTime time.Time
	// The index of the oldest segment in the table.
	LowestSegmentIndex uint32
	// The index of the newest segment in the table.
	HighestSegmentIndex uint32
	// The type of the keymap used by the table. If "", then this table doesn't have a keymap (i.e. it will rebuild
	// a keymap the next time it is loaded).
	KeymapType string
}

// tableInfoCommand is the CLI command handler for the "table-info" command.
func tableInfoCommand(ctx *cli.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // duration non-negative
//nolint:gosec // duration non-negative

// tableInfo retrieves information about a table at the specified path.
func tableInfo(logger *slog.Logger, tableName string, paths []string, fsync bool) (*TableInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Forbid touching tables in active use.

// This should be impossible since we aren't doing anything on background threads that report to the
// error monitor, but it doesn't hurt to check.

// Do not attempt to read segments outside the limit set by the boundary file.
