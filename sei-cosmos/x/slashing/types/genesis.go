package types

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params, signingInfos []SigningInfo, missedBlocks []ValidatorMissedBlockArray,
) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// ValidateGenesis validates the slashing genesis parameters
func ValidateGenesis(data GenesisState) error { _ = "STUB: not implemented"; return nil }
