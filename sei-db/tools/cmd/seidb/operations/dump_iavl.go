package operations

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/memiavl"
	"github.com/spf13/cobra"
)

func DumpIAVLCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executeDumpIAVL(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// DumpIAVLData print the raw keys and values for given module at given height for memIAVL tree
func DumpIAVLData(module string, db *memiavl.DB, outputDir string) error {
	_ = "STUB: not implemented"
	return nil
}
