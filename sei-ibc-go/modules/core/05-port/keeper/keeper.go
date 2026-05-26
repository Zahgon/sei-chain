package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitykeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/keeper"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
)

var logger = seilog.NewLogger("ibc-go", "modules", "core", "05-port", "keeper")

// Keeper defines the IBC connection keeper
type Keeper struct {
	Router *types.Router

	scopedKeeper capabilitykeeper.ScopedKeeper
}

// NewKeeper creates a new IBC connection Keeper instance
func NewKeeper(sck capabilitykeeper.ScopedKeeper) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// IsBound checks a given port ID is already bounded.
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_ = "STUB: not implemented"
	return false
}

// BindPort binds to a port and returns the associated capability.
// Ports must be bound statically when the chain starts in `app.go`.
// The capability must then be passed to a module which will need to pass
// it as an extra parameter when calling functions on the IBC module.
func (k *Keeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	_ = "STUB: not implemented"
	return nil
}

// Authenticate authenticates a capability key against a port ID
// by checking if the memory address of the capability was previously
// generated and bound to the port (provided as a parameter) which the capability
// is being authenticated against.
func (k Keeper) Authenticate(ctx sdk.Context, key *capabilitytypes.Capability, portID string) bool {
	_ = "STUB: not implemented"
	return false
}

// LookupModuleByPort will return the IBCModule along with the capability associated with a given portID
func (k Keeper) LookupModuleByPort(ctx sdk.Context, portID string) (string, *capabilitytypes.Capability, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
