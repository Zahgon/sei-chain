package keeper

import (
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// LockedCoins doesn't use iterators
