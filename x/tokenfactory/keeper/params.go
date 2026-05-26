package keeper

import (
	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// GetParams returns the total set params.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the total set of params.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
