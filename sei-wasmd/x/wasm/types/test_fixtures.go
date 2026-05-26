package types

func GenesisFixture(mutators ...func(*GenesisState)) GenesisState {
	_ = "STUB: not implemented"
	return *new(GenesisState)
}

// #nosec G115 -- constant iterations.

func randBytes(n int) []byte { _ = "STUB: not implemented"; return nil }

func CodeFixture(mutators ...func(*Code)) Code { _ = "STUB: not implemented"; return *new(Code) }

func CodeInfoFixture(mutators ...func(*CodeInfo)) CodeInfo {
	_ = "STUB: not implemented"
	return *new(CodeInfo)
}

func ContractFixture(mutators ...func(*Contract)) Contract {
	_ = "STUB: not implemented"
	return *new(Contract)
}

func OnlyGenesisFields(info *ContractInfo) { _ = "STUB: not implemented"; return }

func ContractInfoFixture(mutators ...func(*ContractInfo)) ContractInfo {
	_ = "STUB: not implemented"
	return *new(ContractInfo)
}

func WithSHA256CodeHash(wasmCode []byte) func(info *CodeInfo) {
	_ = "STUB: not implemented"
	return nil
}

func MsgStoreCodeFixture(mutators ...func(*MsgStoreCode)) *MsgStoreCode {
	_ = "STUB: not implemented"
	return nil
}

func MsgInstantiateContractFixture(mutators ...func(*MsgInstantiateContract)) *MsgInstantiateContract {
	_ = "STUB: not implemented"
	return nil
}

func MsgExecuteContractFixture(mutators ...func(*MsgExecuteContract)) *MsgExecuteContract {
	_ = "STUB: not implemented"
	return nil
}

func StoreCodeProposalFixture(mutators ...func(*StoreCodeProposal)) *StoreCodeProposal {
	_ = "STUB: not implemented"
	return nil
}

func InstantiateContractProposalFixture(mutators ...func(p *InstantiateContractProposal)) *InstantiateContractProposal {
	_ = "STUB: not implemented"
	return nil
}

func MigrateContractProposalFixture(mutators ...func(p *MigrateContractProposal)) *MigrateContractProposal {
	_ = "STUB: not implemented"
	return nil
}

func SudoContractProposalFixture(mutators ...func(p *SudoContractProposal)) *SudoContractProposal {
	_ = "STUB: not implemented"
	return nil
}

func ExecuteContractProposalFixture(mutators ...func(p *ExecuteContractProposal)) *ExecuteContractProposal {
	_ = "STUB: not implemented"
	return nil
}

func UpdateAdminProposalFixture(mutators ...func(p *UpdateAdminProposal)) *UpdateAdminProposal {
	_ = "STUB: not implemented"
	return nil
}

func ClearAdminProposalFixture(mutators ...func(p *ClearAdminProposal)) *ClearAdminProposal {
	_ = "STUB: not implemented"
	return nil
}
