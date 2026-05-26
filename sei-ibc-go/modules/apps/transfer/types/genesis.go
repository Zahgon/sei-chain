package types

// NewGenesisState creates a new ibc-transfer GenesisState instance.
func NewGenesisState(portID string, denomTraces Traces, params Params) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState returns a GenesisState with "transfer" as the default PortID.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }
