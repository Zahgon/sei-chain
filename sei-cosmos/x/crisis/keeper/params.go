package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// GetConstantFee get's the constant fee from the paramSpace
func (k Keeper) GetConstantFee(ctx sdk.Context) (constantFee sdk.Coin) {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

// GetConstantFee set's the constant fee in the paramSpace
func (k Keeper) SetConstantFee(ctx sdk.Context, constantFee sdk.Coin) {
	_ = "STUB: not implemented"
	return
}
