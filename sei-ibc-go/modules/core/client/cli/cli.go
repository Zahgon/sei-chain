package cli

import (
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// Group ibc queries under a subcommand
	return nil
}
