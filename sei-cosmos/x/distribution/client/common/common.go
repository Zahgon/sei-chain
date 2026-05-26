package common

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// QueryDelegationRewards queries a delegation rewards between a delegator and a
// validator.
func QueryDelegationRewards(clientCtx client.Context, delAddr, valAddr string) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// QueryDelegatorValidators returns delegator's list of validators
// it submitted delegations to.
func QueryDelegatorValidators(clientCtx client.Context, delegatorAddr sdk.AccAddress) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryValidatorCommission returns a validator's commission.
func QueryValidatorCommission(clientCtx client.Context, validatorAddr sdk.ValAddress) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithdrawAllDelegatorRewards builds a multi-message slice to be used
// to withdraw all delegations rewards for the given delegator.
func WithdrawAllDelegatorRewards(clientCtx client.Context, delegatorAddr sdk.AccAddress) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	// retrieve the comprehensive list of all validators which the
	// delegator had submitted delegations to
	return nil, nil
}

// build multi-message transaction

// WithdrawValidatorRewardsAndCommission builds a two-message message slice to be
// used to withdraw both validation's commission and self-delegation reward.
func WithdrawValidatorRewardsAndCommission(validatorAddr sdk.ValAddress) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// build and validate MsgWithdrawDelegatorReward
