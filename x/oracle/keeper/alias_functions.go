package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

// GetOracleAccount returns oracle ModuleAccount
func (k Keeper) GetOracleAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	_ = "STUB: not implemented"
	return *new(authtypes.ModuleAccountI)
}

// GetRewardPool retrieves the balance of the oracle module account
func (k Keeper) GetRewardPool(ctx sdk.Context, denom string) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

// GetRewardPool retrieves the balance of the oracle module account
func (k Keeper) GetRewardPoolLegacy(ctx sdk.Context) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}
