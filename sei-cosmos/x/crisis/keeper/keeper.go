package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("cosmos", "x", "crisis", "keeper")

// Keeper - crisis keeper
type Keeper struct {
	routes         []types.InvarRoute
	paramSpace     paramtypes.Subspace
	invCheckPeriod uint

	supplyKeeper types.SupplyKeeper

	feeCollectorName string // name of the FeeCollector ModuleAccount
}

// NewKeeper creates a new Keeper object
func NewKeeper(
	paramSpace paramtypes.Subspace, invCheckPeriod uint, supplyKeeper types.SupplyKeeper,
	feeCollectorName string,
) Keeper {
	_ = "STUB: not implemented"

	// set KeyTable if it has not already been set
	return *new(Keeper)
}

// RegisterRoute register the routes for each of the invariants
func (k *Keeper) RegisterRoute(moduleName, route string, invar sdk.Invariant) {
	_ = "STUB: not implemented"
	return
}

// Routes - return the keeper's invariant routes
func (k Keeper) Routes() []types.InvarRoute {
	_ = "STUB: not implemented"

	// Invariants returns a copy of all registered Crisis keeper invariants.
	return nil
}

func (k Keeper) Invariants() []sdk.Invariant { _ = "STUB: not implemented"; return nil }

// AssertInvariants asserts all registered invariants. If any invariant fails,
// the method panics.
func (k Keeper) AssertInvariants(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// TODO: Include app name as part of context to allow for this to be
// variable.

// InvCheckPeriod returns the invariant checks period.
func (k Keeper) InvCheckPeriod() uint { _ = "STUB: not implemented"; return 0 }

// SendCoinsFromAccountToFeeCollector transfers amt to the fee collector account.
func (k Keeper) SendCoinsFromAccountToFeeCollector(ctx sdk.Context, senderAddr sdk.AccAddress, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}
