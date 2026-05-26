package types

import (
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

type ProposalType string

const (
	ProposalTypeStoreCode               ProposalType = "StoreCode"
	ProposalTypeInstantiateContract     ProposalType = "InstantiateContract"
	ProposalTypeMigrateContract         ProposalType = "MigrateContract"
	ProposalTypeSudoContract            ProposalType = "SudoContract"
	ProposalTypeExecuteContract         ProposalType = "ExecuteContract"
	ProposalTypeUpdateAdmin             ProposalType = "UpdateAdmin"
	ProposalTypeClearAdmin              ProposalType = "ClearAdmin"
	ProposalTypePinCodes                ProposalType = "PinCodes"
	ProposalTypeUnpinCodes              ProposalType = "UnpinCodes"
	ProposalTypeUpdateInstantiateConfig ProposalType = "UpdateInstantiateConfig"
)

// DisableAllProposals contains no wasm gov types.
var DisableAllProposals []ProposalType

// EnableAllProposals contains all wasm gov types as keys.
var EnableAllProposals = []ProposalType{
	ProposalTypeStoreCode,
	ProposalTypeInstantiateContract,
	ProposalTypeMigrateContract,
	ProposalTypeSudoContract,
	ProposalTypeExecuteContract,
	ProposalTypeUpdateAdmin,
	ProposalTypeClearAdmin,
	ProposalTypePinCodes,
	ProposalTypeUnpinCodes,
	ProposalTypeUpdateInstantiateConfig,
}

// ConvertToProposals maps each key to a ProposalType and returns a typed list.
// If any string is not a valid type (in this file), then return an error
func ConvertToProposals(keys []string) ([]ProposalType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func init() { // register new content types with the sdk
	govtypes.RegisterProposalType(string(ProposalTypeStoreCode))
	govtypes.RegisterProposalType(string(ProposalTypeInstantiateContract))
	govtypes.RegisterProposalType(string(ProposalTypeMigrateContract))
	govtypes.RegisterProposalType(string(ProposalTypeSudoContract))
	govtypes.RegisterProposalType(string(ProposalTypeExecuteContract))
	govtypes.RegisterProposalType(string(ProposalTypeUpdateAdmin))
	govtypes.RegisterProposalType(string(ProposalTypeClearAdmin))
	govtypes.RegisterProposalType(string(ProposalTypePinCodes))
	govtypes.RegisterProposalType(string(ProposalTypeUnpinCodes))
	govtypes.RegisterProposalType(string(ProposalTypeUpdateInstantiateConfig))
	govtypes.RegisterProposalTypeCodec(&StoreCodeProposal{}, "wasm/StoreCodeProposal")
	govtypes.RegisterProposalTypeCodec(&InstantiateContractProposal{}, "wasm/InstantiateContractProposal")
	govtypes.RegisterProposalTypeCodec(&MigrateContractProposal{}, "wasm/MigrateContractProposal")
	govtypes.RegisterProposalTypeCodec(&SudoContractProposal{}, "wasm/SudoContractProposal")
	govtypes.RegisterProposalTypeCodec(&ExecuteContractProposal{}, "wasm/ExecuteContractProposal")
	govtypes.RegisterProposalTypeCodec(&UpdateAdminProposal{}, "wasm/UpdateAdminProposal")
	govtypes.RegisterProposalTypeCodec(&ClearAdminProposal{}, "wasm/ClearAdminProposal")
	govtypes.RegisterProposalTypeCodec(&PinCodesProposal{}, "wasm/PinCodesProposal")
	govtypes.RegisterProposalTypeCodec(&UnpinCodesProposal{}, "wasm/UnpinCodesProposal")
	govtypes.RegisterProposalTypeCodec(&UpdateInstantiateConfigProposal{}, "wasm/UpdateInstantiateConfigProposal")
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p StoreCodeProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *StoreCodeProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p StoreCodeProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p StoreCodeProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p StoreCodeProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p StoreCodeProposal) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML pretty prints the wasm byte code
func (p StoreCodeProposal) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p InstantiateContractProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *InstantiateContractProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p InstantiateContractProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p InstantiateContractProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p InstantiateContractProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p InstantiateContractProposal) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML pretty prints the init message
func (p InstantiateContractProposal) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p MigrateContractProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *MigrateContractProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p MigrateContractProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p MigrateContractProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p MigrateContractProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p MigrateContractProposal) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML pretty prints the migrate message
func (p MigrateContractProposal) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p SudoContractProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *SudoContractProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p SudoContractProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p SudoContractProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p SudoContractProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p SudoContractProposal) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML pretty prints the migrate message
func (p SudoContractProposal) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p ExecuteContractProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *ExecuteContractProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p ExecuteContractProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p ExecuteContractProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p ExecuteContractProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p ExecuteContractProposal) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML pretty prints the migrate message
func (p ExecuteContractProposal) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p UpdateAdminProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *UpdateAdminProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p UpdateAdminProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p UpdateAdminProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p UpdateAdminProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p UpdateAdminProposal) String() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a parameter change proposal.
func (p ClearAdminProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *ClearAdminProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p ClearAdminProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p ClearAdminProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p ClearAdminProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p ClearAdminProposal) String() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a parameter change proposal.
func (p PinCodesProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *PinCodesProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p PinCodesProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p PinCodesProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p PinCodesProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p PinCodesProposal) String() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a parameter change proposal.
func (p UnpinCodesProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *UnpinCodesProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p UnpinCodesProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalType returns the type
func (p UnpinCodesProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the proposal
func (p UnpinCodesProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface.
func (p UnpinCodesProposal) String() string { _ = "STUB: not implemented"; return "" }

func validateProposalCommons(title, description string) error {
	_ = "STUB: not implemented"
	return nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p UpdateInstantiateConfigProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// GetTitle returns the title of the proposal
	return ""
}

func (p *UpdateInstantiateConfigProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the human readable description of the proposal
	return ""
}

func (p UpdateInstantiateConfigProposal) GetDescription() string {
	_ = "STUB: not implemented"
	return ""

	// ProposalType returns the type
}

func (p UpdateInstantiateConfigProposal) ProposalType() string {
	_ = "STUB: not implemented"
	return ""
}

// ValidateBasic validates the proposal
func (p UpdateInstantiateConfigProposal) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

// String implements the Stringer interface.
func (p UpdateInstantiateConfigProposal) String() string { _ = "STUB: not implemented"; return "" }

// String implements the Stringer interface.
func (c AccessConfigUpdate) String() string { _ = "STUB: not implemented"; return "" }
