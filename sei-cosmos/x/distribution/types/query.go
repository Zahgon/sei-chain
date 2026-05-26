package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// QueryDelegatorTotalRewardsResponse defines the properties of
// QueryDelegatorTotalRewards query's response.
type QueryDelegatorTotalRewardsResponse struct {
	Rewards []DelegationDelegatorReward `json:"rewards" yaml:"rewards"`
	Total   sdk.DecCoins                `json:"total" yaml:"total"`
}

// NewQueryDelegatorTotalRewardsResponse constructs a QueryDelegatorTotalRewardsResponse
func NewQueryDelegatorTotalRewardsResponse(rewards []DelegationDelegatorReward,
	total sdk.DecCoins) QueryDelegatorTotalRewardsResponse {
	_ = "STUB: not implemented"
	return *new(QueryDelegatorTotalRewardsResponse)
}

func (res QueryDelegatorTotalRewardsResponse) String() string { _ = "STUB: not implemented"; return "" }

// NewDelegationDelegatorReward constructs a DelegationDelegatorReward.
func NewDelegationDelegatorReward(valAddr sdk.ValAddress,
	reward sdk.DecCoins) DelegationDelegatorReward {
	_ = "STUB: not implemented"
	return *new(DelegationDelegatorReward)
}
