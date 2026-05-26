package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

var _ codectypes.UnpackInterfacesMessage = GenesisState{}

// DefaultGenesisState returns the ibc module's default genesis state.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (gs GenesisState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs *GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }
