package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryExchangeRates implements the query rate command.
func GetCmdQueryExchangeRates() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetCmdQueryPriceSnapshotHistory() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetCmdQueryTwaps() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryActives implements the query actives command.
func GetCmdQueryActives() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryFeederDelegation implements the query feeder delegation command
func GetCmdQueryFeederDelegation() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryVotePenaltyCounter implements the query vote penalty counter of the validator command
func GetCmdQueryVotePenaltyCounter() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryVoteTargets implements the query params command.
func GetCmdQueryVoteTargets() *cobra.Command { _ = "STUB: not implemented"; return nil }
