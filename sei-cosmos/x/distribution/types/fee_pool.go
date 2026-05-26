package types

// zero fee pool
func InitialFeePool() FeePool { _ = "STUB: not implemented"; return *new(FeePool) }

// ValidateGenesis validates the fee pool for a genesis state
func (f FeePool) ValidateGenesis() error { _ = "STUB: not implemented"; return nil }
