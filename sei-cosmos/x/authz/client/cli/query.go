package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryGrants implements the query authorization command.
func GetCmdQueryGrants() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetQueryGranterGrants() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetQueryGranteeGrants() *cobra.Command { _ = "STUB: not implemented"; return nil }
