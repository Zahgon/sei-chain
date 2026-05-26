package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	capabilitytypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/capability/types"
)

// bindIbcPort will reserve the port.
// returns a string name of the port or error if we cannot bind it.
// this will fail if call twice.
func (k Keeper) bindIbcPort(ctx sdk.Context, portID string) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureIbcPort is like registerIbcPort, but it checks if we already hold the port
// before calling register, so this is safe to call multiple times.
// Returns success if we already registered or just registered and error if we cannot
// (lack of permissions or someone else has it)
func (k Keeper) ensureIbcPort(ctx sdk.Context, contractAddr sdk.AccAddress) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

const portIDPrefix = "wasm."

func PortIDForContract(addr sdk.AccAddress) string { _ = "STUB: not implemented"; return "" }

func ContractFromPortID(portID string) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// ClaimCapability allows the transfer module to claim a capability
// that IBC module passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	_ = "STUB: not implemented"
	return nil
}
