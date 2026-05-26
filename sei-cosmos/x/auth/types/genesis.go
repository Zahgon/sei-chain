package types

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

var _ types.UnpackInterfacesMessage = GenesisState{}

// RandomGenesisAccountsFn defines the function required to generate custom account types
type RandomGenesisAccountsFn func(simState *module.SimulationState) GenesisAccounts

// NewGenesisState - Create a new genesis state
func NewGenesisState(params Params, accounts GenesisAccounts) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (g GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// GetGenesisStateFromAppState returns x/auth GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) GenesisState {
	_ = "STUB: not implemented"
	return *new(GenesisState)
}

// ValidateGenesis performs basic validation of auth genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { _ = "STUB: not implemented"; return nil }

// SanitizeGenesisAccounts sorts accounts and coin sets.
func SanitizeGenesisAccounts(genAccs GenesisAccounts) GenesisAccounts {
	_ = "STUB: not implemented"
	return *new(GenesisAccounts)
}

// ValidateGenAccounts validates an array of GenesisAccounts and checks for duplicates
func ValidateGenAccounts(accounts GenesisAccounts) error { _ = "STUB: not implemented"; return nil }

// check for duplicated accounts

// check account specific validation

// GenesisAccountIterator implements genesis account iteration.
type GenesisAccountIterator struct{}

// IterateGenesisAccounts iterates over all the genesis accounts found in
// appGenesis and invokes a callback on each genesis account. If any call
// returns true, iteration stops.
func (GenesisAccountIterator) IterateGenesisAccounts(
	cdc codec.Codec, appGenesis map[string]json.RawMessage, cb func(AccountI) (stop bool),
) {
	_ = "STUB: not implemented"
	return
}

// PackAccounts converts GenesisAccounts to Any slice
func PackAccounts(accounts GenesisAccounts) ([]*types.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackAccounts converts Any slice to GenesisAccounts
func UnpackAccounts(accountsAny []*types.Any) (GenesisAccounts, error) {
	_ = "STUB: not implemented"
	return *new(GenesisAccounts), nil
}
