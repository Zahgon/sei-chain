package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
)

// GetSendEnabled retrieves the send enabled boolean from the paramstore
func (k Keeper) GetSendEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// GetReceiveEnabled retrieves the receive enabled boolean from the paramstore
func (k Keeper) GetReceiveEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// GetParams returns the total set of ibc-transfer parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the total set of ibc-transfer parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
