package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/exported"
)

var _ types.UnpackInterfacesMessage = GenesisState{}

// NewGenesisState creates a new genesis state for the evidence module.
func NewGenesisState(e []exported.Evidence) *GenesisState { _ = "STUB: not implemented"; return nil }

// DefaultGenesisState returns the evidence module's default genesis state.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// Validate performs basic gensis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (gs GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
