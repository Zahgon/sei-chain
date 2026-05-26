package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// DistributeFromFeePool distributes funds from the distribution module account to
// a receiver address while updating the community pool
func (k Keeper) DistributeFromFeePool(ctx sdk.Context, amount sdk.Coins, receiveAddr sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE the community pool isn't a module account, however its coins
// are held in the distribution module account. Thus the community pool
// must be reduced separately from the SendCoinsFromModuleToAccount call
