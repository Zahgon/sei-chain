package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// TotalSupply checks that the total supply reflects all the coins held in accounts
func TotalSupply(k Keeper) sdk.Invariant { _ = "STUB: not implemented"; return *new(sdk.Invariant) }

// also iterate over deferred balances
