package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

// querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over q
type querier struct {
	Keeper
}

// NewQuerier returns an implementation of the oracle QueryServer interface
// for the provided Keeper.
func NewQuerier(keeper Keeper) types.QueryServer {
	_ = "STUB: not implemented"
	return *new(types.QueryServer)
}

var _ types.QueryServer = querier{}

// Params queries params of distribution module
func (q querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExchangeRate queries exchange rate of a denom
func (q querier) ExchangeRate(c context.Context, req *types.QueryExchangeRateRequest) (*types.QueryExchangeRateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExchangeRates queries exchange rates of all denoms
func (q querier) ExchangeRates(c context.Context, _ *types.QueryExchangeRatesRequest) (*types.QueryExchangeRatesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Actives queries all denoms for which exchange rates exist
func (q querier) Actives(c context.Context, _ *types.QueryActivesRequest) (*types.QueryActivesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VoteTargets queries the voting target list on current vote period
func (q querier) VoteTargets(c context.Context, _ *types.QueryVoteTargetsRequest) (*types.QueryVoteTargetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q querier) PriceSnapshotHistory(c context.Context, _ *types.QueryPriceSnapshotHistoryRequest) (*types.QueryPriceSnapshotHistoryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q querier) Twaps(c context.Context, req *types.QueryTwapsRequest) (*types.QueryTwapsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FeederDelegation queries the account address that the validator operator delegated oracle vote rights to
func (q querier) FeederDelegation(c context.Context, req *types.QueryFeederDelegationRequest) (*types.QueryFeederDelegationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MissCounter queries oracle miss counter of a validator
func (q querier) VotePenaltyCounter(c context.Context, req *types.QueryVotePenaltyCounterRequest) (*types.QueryVotePenaltyCounterResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q querier) SlashWindow(
	goCtx context.Context,
	_ *types.QuerySlashWindowRequest,
) (*types.QuerySlashWindowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The window progress is the number of vote periods that have been completed in the current slashing window. With a vote period of 1, this will be equivalent to the number of blocks that have progressed in the slash window.

// nolint:gosec
