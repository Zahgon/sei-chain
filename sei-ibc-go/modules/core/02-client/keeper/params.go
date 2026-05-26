package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

// GetAllowedClients retrieves the allowed clients from the paramstore
func (k Keeper) GetAllowedClients(ctx sdk.Context) []string { _ = "STUB: not implemented"; return nil }

// GetParams returns the total set of ibc-client parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the total set of ibc-client parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
