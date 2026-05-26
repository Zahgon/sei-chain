package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/distribution/types"
)

var _ types.QueryServer = Keeper{}

// Params queries params of distribution module
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidatorOutstandingRewards queries rewards of a validator address
func (k Keeper) ValidatorOutstandingRewards(c context.Context, req *types.QueryValidatorOutstandingRewardsRequest) (*types.QueryValidatorOutstandingRewardsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidatorCommission queries accumulated commission for a validator
func (k Keeper) ValidatorCommission(c context.Context, req *types.QueryValidatorCommissionRequest) (*types.QueryValidatorCommissionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidatorSlashes queries slash events of a validator
func (k Keeper) ValidatorSlashes(c context.Context, req *types.QueryValidatorSlashesRequest) (*types.QueryValidatorSlashesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegationRewards the total rewards accrued by a delegation
func (k Keeper) DelegationRewards(c context.Context, req *types.QueryDelegationRewardsRequest) (*types.QueryDelegationRewardsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegationTotalRewards the total rewards accrued by a each validator
func (k Keeper) DelegationTotalRewards(c context.Context, req *types.QueryDelegationTotalRewardsRequest) (*types.QueryDelegationTotalRewardsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegatorValidators queries the validators list of a delegator
func (k Keeper) DelegatorValidators(c context.Context, req *types.QueryDelegatorValidatorsRequest) (*types.QueryDelegatorValidatorsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelegatorWithdrawAddress queries Query/delegatorWithdrawAddress
func (k Keeper) DelegatorWithdrawAddress(c context.Context, req *types.QueryDelegatorWithdrawAddressRequest) (*types.QueryDelegatorWithdrawAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CommunityPool queries the community pool coins
func (k Keeper) CommunityPool(c context.Context, req *types.QueryCommunityPoolRequest) (*types.QueryCommunityPoolResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
