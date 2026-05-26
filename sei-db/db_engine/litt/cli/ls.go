package main

import (
	"log/slog"

	"github.com/urfave/cli/v2"
)

func lsCommand(ctx *cli.Context) error { _ = "STUB: not implemented"; return nil }

// Similar to ls, but searches for tables in multiple paths.
func lsPaths(logger *slog.Logger, rootPaths []string, lock bool, fsync bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of LittDB tables at the specified LittDB path. Tables are alphabetically sorted by their names.
// Returns an error if the path does not exist or if no tables are found.
func ls(logger *slog.Logger, rootPath string, lock bool, fsync bool) ([]string, error) {
	_ = "STUB: not implemented"

	// Forbid touching tables in active use.
	return nil, nil
}

// LittDB has one directory under the root directory per table, with the name
// of the table being the name of the directory.

// Each table directory will contain a "segments" directory. Infer that any directory containing this directory
// is a table. If we are looking at a real LittDB instance, there shouldn't be any other directories, but
// there is no need to enforce that here.

// Alphabetically sort the tables.
