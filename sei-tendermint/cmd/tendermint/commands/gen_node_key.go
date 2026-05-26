package commands

import (
	"github.com/spf13/cobra"
)

// GenNodeKeyCmd allows the generation of a node key. It prints JSON-encoded
// NodeKey to the standard output.
var GenNodeKeyCmd = &cobra.Command{
	Use:   "gen-node-key",
	Short: "Generate a new node key",
	RunE:  genNodeKey,
}

func genNodeKey(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
