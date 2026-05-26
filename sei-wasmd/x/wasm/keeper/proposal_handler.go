package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// NewWasmProposalHandler creates a new governance Handler for wasm proposals
func NewWasmProposalHandler(k decoratedKeeper, enabledProposalTypes []types.ProposalType) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

// NewWasmProposalHandlerX creates a new governance Handler for wasm proposals
func NewWasmProposalHandlerX(k types.ContractOpsKeeper, enabledProposalTypes []types.ProposalType) govtypes.Handler {
	_ = "STUB: not implemented"
	return *new(govtypes.Handler)
}

func handleStoreCodeProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.StoreCodeProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleInstantiateProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.InstantiateContractProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleMigrateProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.MigrateContractProposal) error {
	_ = "STUB: not implemented"
	return nil
}

// runAs is not used if this is permissioned, so just put any valid address there (second contractAddr)

func handleSudoProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.SudoContractProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleExecuteProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.ExecuteContractProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUpdateAdminProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.UpdateAdminProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleClearAdminProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.ClearAdminProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handlePinCodesProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.PinCodesProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUnpinCodesProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.UnpinCodesProposal) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUpdateInstantiateConfigProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.UpdateInstantiateConfigProposal) error {
	_ = "STUB: not implemented"
	return nil
}
