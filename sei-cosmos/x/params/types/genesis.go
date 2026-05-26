package types

func DefaultFeesParams() *FeesParams { _ = "STUB: not implemented"; return nil }

// 0.01 by default on a chain level

func DefaultCosmosGasParams() *CosmosGasParams { _ = "STUB: not implemented"; return nil }

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState { _ = "STUB: not implemented"; return nil }

func NewGenesisState(feesParams FeesParams, cosmosGasParams CosmosGasParams) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }
