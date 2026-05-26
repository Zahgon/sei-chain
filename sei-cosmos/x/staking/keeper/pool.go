package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

// GetBondedPool returns the bonded tokens pool's module account
func (k Keeper) GetBondedPool(ctx sdk.Context) (bondedPool authtypes.ModuleAccountI) {
	_ = "STUB: not implemented"
	return *new(authtypes.ModuleAccountI)
}

// GetNotBondedPool returns the not bonded tokens pool's module account
func (k Keeper) GetNotBondedPool(ctx sdk.Context) (notBondedPool authtypes.ModuleAccountI) {
	_ = "STUB: not implemented"
	return *new(authtypes.ModuleAccountI)
}

// bondedTokensToNotBonded transfers coins from the bonded to the not bonded pool within staking
func (k Keeper) bondedTokensToNotBonded(ctx sdk.Context, tokens sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// notBondedTokensToBonded transfers coins from the not bonded to the bonded pool within staking
func (k Keeper) notBondedTokensToBonded(ctx sdk.Context, tokens sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// burnBondedTokens removes coins from the bonded pool module account
func (k Keeper) burnBondedTokens(ctx sdk.Context, amt sdk.Int) error {
	_ = "STUB: not implemented"
	return nil

	// skip as no coins need to be burned
}

// burnNotBondedTokens removes coins from the not bonded pool module account
func (k Keeper) burnNotBondedTokens(ctx sdk.Context, amt sdk.Int) error {
	_ = "STUB: not implemented"
	return nil

	// skip as no coins need to be burned
}

// TotalBondedTokens total staking tokens supply which is bonded
func (k Keeper) TotalBondedTokens(ctx sdk.Context) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// StakingTokenSupply staking tokens from the total supply
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// BondedRatio the fraction of the staking tokens which are currently bonded
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}
