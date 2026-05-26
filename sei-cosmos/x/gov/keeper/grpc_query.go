package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

var _ types.QueryServer = Keeper{}

// Proposal returns proposal details based on ProposalID
func (q Keeper) Proposal(c context.Context, req *types.QueryProposalRequest) (*types.QueryProposalResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Proposals implements the Query/Proposals gRPC method
func (q Keeper) Proposals(c context.Context, req *types.QueryProposalsRequest) (*types.QueryProposalsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// match status (if supplied/valid)

// match voter address (if supplied)

// match depositor (if supplied)

// Vote returns Voted information based on proposalID, voterAddr
func (q Keeper) Vote(c context.Context, req *types.QueryVoteRequest) (*types.QueryVoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Votes returns single proposal's votes
func (q Keeper) Votes(c context.Context, req *types.QueryVotesRequest) (*types.QueryVotesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Params queries all params
func (q Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deposit queries single deposit information based proposalID, depositAddr
func (q Keeper) Deposit(c context.Context, req *types.QueryDepositRequest) (*types.QueryDepositResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deposits returns single proposal's all deposits
func (q Keeper) Deposits(c context.Context, req *types.QueryDepositsRequest) (*types.QueryDepositsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TallyResult queries the tally of a proposal vote
func (q Keeper) TallyResult(c context.Context, req *types.QueryTallyResultRequest) (*types.QueryTallyResultResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// proposal is in voting period
