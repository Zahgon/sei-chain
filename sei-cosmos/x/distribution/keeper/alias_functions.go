package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// get outstanding rewards
func (k Keeper) GetValidatorOutstandingRewardsCoins(ctx sdk.Context, val sdk.ValAddress) sdk.DecCoins {
	_ = "STUB: not implemented"
	return *new(sdk.DecCoins)
}

// get the community coins
func (k Keeper) GetFeePoolCommunityCoins(ctx sdk.Context) sdk.DecCoins {
	_ = "STUB: not implemented"
	return *new(sdk.DecCoins)
}

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	_ = "STUB: not implemented"
	return *new(authtypes.ModuleAccountI)
}

func (k Keeper) GetAllValidators(ctx sdk.Context) []stakingtypes.Validator {
	_ = "STUB: not implemented"
	return nil
}
