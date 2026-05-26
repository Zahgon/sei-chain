package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryValidatorOutstandingRewards implements the query validator
// outstanding rewards command.
func GetCmdQueryValidatorOutstandingRewards() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryValidatorCommission implements the query validator commission command.
func GetCmdQueryValidatorCommission() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryValidatorSlashes implements the query validator slashes command.
func GetCmdQueryValidatorSlashes() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryDelegatorRewards implements the query delegator rewards command.
func GetCmdQueryDelegatorRewards() *cobra.Command { _ = "STUB: not implemented"; return nil }

// query for rewards from a particular delegation

// GetCmdQueryCommunityPool returns the command for fetching community pool info.
func GetCmdQueryCommunityPool() *cobra.Command { _ = "STUB: not implemented"; return nil }
