package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(_ sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *

	// SetParams set the params
	new(types.Params)
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
