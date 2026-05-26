package simulation

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	simtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/simulation"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/simulation"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// Simulation operation weights constants
//
//nolint:gosec
const (
	OpWeightMsgStoreCode           = "op_weight_msg_store_code"
	OpWeightMsgInstantiateContract = "op_weight_msg_instantiate_contract"
	OpReflectContractPath          = "op_reflect_contract_path"
)

// WasmKeeper is a subset of the wasm keeper used by simulations
type WasmKeeper interface {
	GetParams(ctx sdk.Context) types.Params
	IterateCodeInfos(ctx sdk.Context, cb func(uint64, types.CodeInfo) bool)
}

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	simstate *module.SimulationState,
	ak types.AccountKeeper,
	bk simulation.BankKeeper,
	wasmKeeper WasmKeeper,
) simulation.WeightedOperations {
	_ = "STUB: not implemented"
	return *new(simulation.WeightedOperations)
}

// simulations are run from the `app` folder

// SimulateMsgStoreCode generates a MsgStoreCode with random values
func SimulateMsgStoreCode(ak types.AccountKeeper, bk simulation.BankKeeper, wasmKeeper WasmKeeper, wasmBz []byte) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}

// SimulateMsgInstantiateContract generates a MsgInstantiateContract with random values
func SimulateMsgInstantiateContract(ak types.AccountKeeper, bk simulation.BankKeeper, wasmKeeper WasmKeeper) simtypes.Operation {
	_ = "STUB: not implemented"
	return *new(simtypes.Operation)
}
