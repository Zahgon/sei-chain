package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// Group tokenfactory queries under a subcommand
	return nil
}

// GetParams returns the params for the module
func GetParams() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdDenomAuthorityMetadata returns the authority metadata for a queried denom
func GetCmdDenomAuthorityMetadata() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdDenomsFromCreator a command to get a list of all tokens created by a specific creator address
func GetCmdDenomsFromCreator() *cobra.Command { _ = "STUB: not implemented"; return nil }
