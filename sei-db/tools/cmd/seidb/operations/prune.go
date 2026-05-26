package operations

import (
	"github.com/spf13/cobra"
)

func PruneCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executePrune(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// Prunes DB at given height
func PruneDB(dbBackend string, dbDir string, version int64) {
	_ = "STUB: not implemented"
	// TODO: Defer Close Db
	return
}

// Callback to write db entries to file
