package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
)

// GetMaxExpectedTimePerBlock retrieves the maximum expected time per block from the paramstore
func (k Keeper) GetMaxExpectedTimePerBlock(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// GetParams returns the total set of ibc-connection parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the total set of ibc-connection parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
