package types

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// NewGenesisState creates a new GenesisState instanc e
func NewGenesisState(params Params, validators []Validator, delegations []Delegation) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// GetGenesisStateFromAppState returns x/staking GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (g GenesisState) UnpackInterfaces(c codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
