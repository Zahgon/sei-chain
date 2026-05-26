package types

// NewConnectionPaths creates a ConnectionPaths instance.
func NewConnectionPaths(id string, paths []string) ConnectionPaths {
	_ = "STUB: not implemented"
	return *new(ConnectionPaths)
}

// NewGenesisState creates a GenesisState instance.
func NewGenesisState(
	connections []IdentifiedConnection, connPaths []ConnectionPaths,
	nextConnectionSequence uint64, params Params,
) GenesisState {
	_ = "STUB: not implemented"
	return *new(GenesisState)
}

// DefaultGenesisState returns the ibc connection submodule's default genesis state.
func DefaultGenesisState() GenesisState { _ = "STUB: not implemented"; return *new(GenesisState) }

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	_ = "STUB: not implemented"
	// keep track of the max sequence to ensure it is less than
	// the next sequence used in creating connection identifers.
	return nil
}
