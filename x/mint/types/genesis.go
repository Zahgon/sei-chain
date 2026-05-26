package types

// NewGenesisState creates a new GenesisState object.
func NewGenesisState(minter Minter, params Params) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState creates a default GenesisState object.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error { _ = "STUB: not implemented"; return nil }
