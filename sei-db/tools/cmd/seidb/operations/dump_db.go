package operations

import (
	"github.com/spf13/cobra"
)

const outputFileName = "db_dump.kv"

func DumpDbCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// TODO: Accept multiple modules. Can pass empty to iterate over all stores

func executeDumpDB(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// Outputs the raw keys and values for all modules at a height to a file
func DumpDbData(dbBackend string, module string, outputDir string, dbDir string) {
	_ = "STUB: not implemented"
	// Create output directory
	return
}

// Create output file

// TODO: Defer Close Db

// Callback to write db entries to file
