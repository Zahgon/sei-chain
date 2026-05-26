package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// TokensToConsensusPower - convert input tokens to potential consensus-engine power
func (k Keeper) TokensToConsensusPower(ctx sdk.Context, tokens sdk.Int) int64 {
	_ = "STUB: not implemented"
	return 0
}

// TokensFromConsensusPower - convert input power to tokens
func (k Keeper) TokensFromConsensusPower(ctx sdk.Context, power int64) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}
