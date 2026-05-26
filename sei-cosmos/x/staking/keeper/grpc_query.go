package keeper

import (
	"context"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/query"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Validators queries all validators that match the given status
func (k Querier) Validators(c context.Context, req *types.QueryValidatorsRequest) (*types.QueryValidatorsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate the provided status, return all the validators if the status is empty

// Validator queries validator info for given validator address
func (k Querier) Validator(c context.Context, req *types.QueryValidatorRequest) (*types.QueryValidatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidatorDelegations queries delegate info for given validator
func (k Querier) ValidatorDelegations(c context.Context, req *types.QueryValidatorDelegationsRequest) (*types.QueryValidatorDelegationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidatorUnbondingDelegations queries unbonding delegations of a validator
func (k Querier) ValidatorUnbondingDelegations(c context.Context, req *types.QueryValidatorUnbondingDelegationsRequest) (*types.QueryValidatorUnbondingDelegationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delegation queries delegate info for given validator delegator pair
func (k Querier) Delegation(c context.Context, req *types.QueryDelegationRequest) (*types.QueryDelegationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnbondingDelegation queries unbonding info for give validator delegator pair
func (k Querier) UnbondingDelegation(c context.Context, req *types.QueryUnbondingDelegationRequest) (*types.QueryUnbondingDelegationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegatorDelegations queries all delegations of a give delegator address
func (k Querier) DelegatorDelegations(c context.Context, req *types.QueryDelegatorDelegationsRequest) (*types.QueryDelegatorDelegationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegatorValidator queries validator info for given delegator validator pair
func (k Querier) DelegatorValidator(c context.Context, req *types.QueryDelegatorValidatorRequest) (*types.QueryDelegatorValidatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegatorUnbondingDelegations queries all unbonding delegations of a given delegator address
func (k Querier) DelegatorUnbondingDelegations(c context.Context, req *types.QueryDelegatorUnbondingDelegationsRequest) (*types.QueryDelegatorUnbondingDelegationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HistoricalInfo queries the historical info for given height
func (k Querier) HistoricalInfo(c context.Context, req *types.QueryHistoricalInfoRequest) (*types.QueryHistoricalInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Redelegations queries redelegations of given address
func (k Querier) Redelegations(c context.Context, req *types.QueryRedelegationsRequest) (*types.QueryRedelegationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegatorValidators queries all validators info for given delegator address
func (k Querier) DelegatorValidators(c context.Context, req *types.QueryDelegatorValidatorsRequest) (*types.QueryDelegatorValidatorsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pool queries the pool info
func (k Querier) Pool(c context.Context, _ *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Params queries the staking parameters
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryRedelegation(ctx sdk.Context, k Querier, req *types.QueryRedelegationsRequest) (redels types.Redelegations, err error) {
	_ = "STUB: not implemented"
	return *new(types.Redelegations), nil
}

func queryRedelegationsFromSrcValidator(store sdk.KVStore, k Querier, req *types.QueryRedelegationsRequest) (redels types.Redelegations, res *query.PageResponse, err error) {
	_ = "STUB: not implemented"
	return *new(types.Redelegations), nil, nil
}

func queryAllRedelegations(store sdk.KVStore, k Querier, req *types.QueryRedelegationsRequest) (redels types.Redelegations, res *query.PageResponse, err error) {
	_ = "STUB: not implemented"
	return *new(types.Redelegations), nil, nil
}
