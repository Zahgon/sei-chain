package feegrant

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

var _ types.UnpackInterfacesMessage = GenesisState{}

// NewGenesisState creates new GenesisState object
func NewGenesisState(entries []Grant) *GenesisState { _ = "STUB: not implemented"; return nil }

// ValidateGenesis ensures all grants in the genesis state are valid
func ValidateGenesis(data GenesisState) error { _ = "STUB: not implemented"; return nil }

// DefaultGenesisState returns default state for feegrant module.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
