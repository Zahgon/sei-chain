package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all x/upgrade CLi query commands.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCurrentPlanCmd returns the query upgrade plan command.
func GetCurrentPlanCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetAppliedPlanCmd returns information about the block at which a completed
// upgrade was applied.
func GetAppliedPlanCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// we got the height, now let's return the headers

// always output json as Header is unreable in toml ([]byte is a long list of numbers)

// GetModuleVersionsCmd returns the module version list from state
func GetModuleVersionsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
