package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// Group gov queries under a subcommand
	return nil
}

// GetCmdQueryProposal implements the query proposal command.
func GetCmdQueryProposal() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposal id is a uint

// Query the proposal

// GetCmdQueryProposals implements a query proposals command. Command to Get a
// Proposal Information.
func GetCmdQueryProposals() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryVote implements the query proposal vote command. Command to Get a
// Proposal Information.
func GetCmdQueryVote() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposal id is a uint

// check to see if the proposal is in the store

// GetCmdQueryVotes implements the command to query for proposal votes.
func GetCmdQueryVotes() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposal id is a uint

// check to see if the proposal is in the store

// TODO migrate to use JSONCodec (implement MarshalJSONArray
// or wrap lists of proto.Message in some other message)

// GetCmdQueryDeposit implements the query proposal deposit command. Command to
// get a specific Deposit Information
func GetCmdQueryDeposit() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposal id is a uint

// check to see if the proposal is in the store

// GetCmdQueryDeposits implements the command to query for proposal deposits.
func GetCmdQueryDeposits() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposal id is a uint

// check to see if the proposal is in the store

// TODO migrate to use JSONCodec (implement MarshalJSONArray
// or wrap lists of proto.Message in some other message)

// GetCmdQueryTally implements the command to query for proposal tally result.
func GetCmdQueryTally() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposal id is a uint

// check to see if the proposal is in the store

// Query store

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Query store for all 3 params

// GetCmdQueryParam implements the query param command.
func GetCmdQueryParam() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Query store

// GetCmdQueryProposer implements the query proposer command.
func GetCmdQueryProposer() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposalID is a uint
