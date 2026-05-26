package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epoch types.Epoch) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epoch types.Epoch) {
	_ = "STUB: not implemented"
	return
}
