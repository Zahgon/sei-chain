package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/types"
)

// GetParams returns the total set of ibc core module parameters.
func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// SetParams sets the ibc core module parameters.
func (k *Keeper) SetParams(ctx sdk.Context, p types.Params) { _ = "STUB: not implemented"; return }

// IsInboundEnabled returns true if inbound IBC is enabled.
func (k *Keeper) IsInboundEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// IsOutboundEnabled returns true if outbound IBC is enabled.
func (k *Keeper) IsOutboundEnabled(ctx sdk.Context) bool { _ = "STUB: not implemented"; return false }

// SetInboundEnabled sets inbound enabled flag.
func (k *Keeper) SetInboundEnabled(ctx sdk.Context, enabled bool) {
	_ = "STUB: not implemented"
	return
}

// SetOutboundEnabled sets outbound enabled flag.
func (k *Keeper) SetOutboundEnabled(ctx sdk.Context, enabled bool) {
	_ = "STUB: not implemented"
	return
}

// GetParamSpace returns the keeper's paramSpace (for other packages if needed).
func (k *Keeper) GetParamSpace() paramtypes.Subspace {
	_ = "STUB: not implemented"
	return *new(paramtypes.Subspace)
}
