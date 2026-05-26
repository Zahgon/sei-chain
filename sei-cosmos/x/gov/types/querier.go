package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// DONTCOVER

// query endpoints supported by the governance Querier
const (
	QueryParams    = "params"
	QueryProposals = "proposals"
	QueryProposal  = "proposal"
	QueryDeposits  = "deposits"
	QueryDeposit   = "deposit"
	QueryVotes     = "votes"
	QueryVote      = "vote"
	QueryTally     = "tally"

	ParamDeposit  = "deposit"
	ParamVoting   = "voting"
	ParamTallying = "tallying"
)

// QueryProposalParams Params for queries:
// - 'custom/gov/proposal'
// - 'custom/gov/deposits'
// - 'custom/gov/tally'
type QueryProposalParams struct {
	ProposalID uint64
}

// NewQueryProposalParams creates a new instance of QueryProposalParams
func NewQueryProposalParams(proposalID uint64) QueryProposalParams {
	_ = "STUB: not implemented"
	return *new(QueryProposalParams)
}

// QueryProposalVotesParams used for queries to 'custom/gov/votes'.
type QueryProposalVotesParams struct {
	ProposalID uint64
	Page       int
	Limit      int
}

// NewQueryProposalVotesParams creates new instance of the QueryProposalVotesParams.
func NewQueryProposalVotesParams(proposalID uint64, page, limit int) QueryProposalVotesParams {
	_ = "STUB: not implemented"
	return *new(QueryProposalVotesParams)
}

// QueryDepositParams params for query 'custom/gov/deposit'
type QueryDepositParams struct {
	ProposalID uint64
	Depositor  sdk.AccAddress
}

// NewQueryDepositParams creates a new instance of QueryDepositParams
func NewQueryDepositParams(proposalID uint64, depositor sdk.AccAddress) QueryDepositParams {
	_ = "STUB: not implemented"
	return *new(QueryDepositParams)
}

// QueryVoteParams Params for query 'custom/gov/vote'
type QueryVoteParams struct {
	ProposalID uint64
	Voter      sdk.AccAddress
}

// NewQueryVoteParams creates a new instance of QueryVoteParams
func NewQueryVoteParams(proposalID uint64, voter sdk.AccAddress) QueryVoteParams {
	_ = "STUB: not implemented"
	return *new(QueryVoteParams)
}

// QueryProposalsParams Params for query 'custom/gov/proposals'
type QueryProposalsParams struct {
	Page           int
	Limit          int
	Voter          sdk.AccAddress
	Depositor      sdk.AccAddress
	ProposalStatus ProposalStatus
}

// NewQueryProposalsParams creates a new instance of QueryProposalsParams
func NewQueryProposalsParams(page, limit int, status ProposalStatus, voter, depositor sdk.AccAddress) QueryProposalsParams {
	_ = "STUB: not implemented"
	return *new(QueryProposalsParams)
}
