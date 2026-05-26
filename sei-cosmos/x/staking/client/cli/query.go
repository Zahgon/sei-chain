package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryValidator implements the validator query command.
func GetCmdQueryValidator() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryValidators implements the query all validators command.
func GetCmdQueryValidators() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Leaving status empty on purpose to query all validators.

// GetCmdQueryHexAddress returns the validator that matches the hex address
func GetCmdQueryHexAddress() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Based on type URL, do different decoding

// GetCmdQueryValidatorUnbondingDelegations implements the query all unbonding delegatations from a validator command.
func GetCmdQueryValidatorUnbondingDelegations() *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// GetCmdQueryValidatorRedelegations implements the query all redelegatations
// from a validator command.
func GetCmdQueryValidatorRedelegations() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryDelegation the query delegation command.
func GetCmdQueryDelegation() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryDelegations implements the command to query all the delegations
// made from one delegator.
func GetCmdQueryDelegations() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryValidatorDelegations implements the command to query all the
// delegations to a specific validator.
func GetCmdQueryValidatorDelegations() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryUnbondingDelegation implements the command to query a single
// unbonding-delegation record.
func GetCmdQueryUnbondingDelegation() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryUnbondingDelegations implements the command to query all the
// unbonding-delegation records for a delegator.
func GetCmdQueryUnbondingDelegations() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryRedelegation implements the command to query a single
// redelegation record.
func GetCmdQueryRedelegation() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryRedelegations implements the command to query all the
// redelegation records for a delegator.
func GetCmdQueryRedelegations() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryHistoricalInfo implements the historical info query command
func GetCmdQueryHistoricalInfo() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryPool implements the pool query command.
func GetCmdQueryPool() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command { _ = "STUB: not implemented"; return nil }
