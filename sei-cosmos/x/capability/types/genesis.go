package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState { _ = "STUB: not implemented"; return nil }

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	_ = "STUB: not implemented"
	// NOTE: index must be greater than 0
	return nil
}

// all exported existing indices must be between [1, gs.Index)
