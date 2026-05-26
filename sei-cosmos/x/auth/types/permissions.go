package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// permissions
const (
	Minter  = "minter"
	Burner  = "burner"
	Staking = "staking"
)

// PermissionsForAddress defines all the registered permissions for an address
type PermissionsForAddress struct {
	permissions []string
	address     sdk.AccAddress
}

// NewPermissionsForAddress creates a new PermissionsForAddress object
func NewPermissionsForAddress(name string, permissions []string) PermissionsForAddress {
	_ = "STUB: not implemented"
	return *new(PermissionsForAddress)
}

// HasPermission returns whether the PermissionsForAddress contains permission.
func (pa PermissionsForAddress) HasPermission(permission string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetAddress returns the address of the PermissionsForAddress object
func (pa PermissionsForAddress) GetAddress() sdk.AccAddress {
	_ = "STUB: not implemented"

	// GetPermissions returns the permissions granted to the address
	return *new(sdk.AccAddress)
}

func (pa PermissionsForAddress) GetPermissions() []string { _ = "STUB: not implemented"; return nil }

// performs basic permission validation
func validatePermissions(permissions ...string) error { _ = "STUB: not implemented"; return nil }
