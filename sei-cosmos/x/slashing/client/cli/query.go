package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// Group slashing queries under a subcommand
	return nil
}

// GetCmdQuerySigningInfo implements the command to query signing info.
func GetCmdQuerySigningInfo() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQuerySigningInfos implements the command to query signing infos.
func GetCmdQuerySigningInfos() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryParams implements a command to fetch slashing parameters.
func GetCmdQueryParams() *cobra.Command { _ = "STUB: not implemented"; return nil }
