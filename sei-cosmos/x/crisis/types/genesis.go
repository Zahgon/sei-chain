package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(constantFee sdk.Coin) *GenesisState { _ = "STUB: not implemented"; return nil }

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// ValidateGenesis - validate crisis genesis data
func ValidateGenesis(data *GenesisState) error { _ = "STUB: not implemented"; return nil }
