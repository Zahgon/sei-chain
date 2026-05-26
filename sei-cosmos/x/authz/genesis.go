package authz

import (
	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// NewGenesisState creates new GenesisState object
func NewGenesisState(entries []GrantAuthorization) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// ValidateGenesis check the given genesis state has no integrity issues
func ValidateGenesis(data GenesisState) error {
	_ = "STUB: not implemented"

	// DefaultGenesisState - Return a default genesis state
	return nil
}

func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

var _ cdctypes.UnpackInterfacesMessage = GenesisState{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg GrantAuthorization) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
