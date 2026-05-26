package types

func DefaultGenesis() *GenesisState { _ = "STUB: not implemented"; return nil }

func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

func ValidateStream(gensisStateCh <-chan GenesisState) error { _ = "STUB: not implemented"; return nil }
