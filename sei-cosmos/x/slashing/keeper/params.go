package keeper

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
)

// SignedBlocksWindow - sliding window for downtime slashing
func (k Keeper) SignedBlocksWindow(ctx sdk.Context) (res int64) {
	_ = "STUB: not implemented"
	return 0
}

// MinSignedPerWindow - minimum blocks signed per window
func (k Keeper) MinSignedPerWindow(ctx sdk.Context) int64 { _ = "STUB: not implemented"; return 0 }

// NOTE: RoundInt64 will never panic as minSignedPerWindow is
//       less than 1.

// DowntimeJailDuration - Downtime unbond duration
func (k Keeper) DowntimeJailDuration(ctx sdk.Context) (res time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// SlashFractionDoubleSign - fraction of power slashed in case of double sign
func (k Keeper) SlashFractionDoubleSign(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// SlashFractionDowntime - fraction of power slashed for downtime
func (k Keeper) SlashFractionDowntime(ctx sdk.Context) (res sdk.Dec) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GetParams returns the total set of slashing parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the slashing parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
