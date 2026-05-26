package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

const EpochKey = "epoch"

func (k Keeper) SetEpoch(ctx sdk.Context, epoch types.Epoch) { _ = "STUB: not implemented"; return }

func (k Keeper) GetEpoch(ctx sdk.Context) (epoch types.Epoch) {
	_ = "STUB: not implemented"
	return *new(types.Epoch)
}
