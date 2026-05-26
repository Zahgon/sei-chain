package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for the minting module.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryParams implements a command to return the current minting
// parameters.
func GetCmdQueryParams() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryEpochProvisions implements a command to return the current minting
// epoch provisions value.
func GetCmdQueryEpochProvisions() *cobra.Command { _ = "STUB: not implemented"; return nil }
