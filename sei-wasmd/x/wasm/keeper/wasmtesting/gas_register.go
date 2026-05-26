package wasmtesting

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

// MockGasRegister mock that implements keeper.GasRegister
type MockGasRegister struct {
	CompileCostFn             func(byteLength int) sdk.Gas
	NewContractInstanceCostFn func(pinned bool, msgLen int) sdk.Gas
	InstantiateContractCostFn func(pinned bool, msgLen int) sdk.Gas
	ReplyCostFn               func(pinned bool, reply wasmvmtypes.Reply) sdk.Gas
	EventCostsFn              func(evts []wasmvmtypes.EventAttribute) sdk.Gas
	ToWasmVMGasFn             func(source sdk.Gas) uint64
	FromWasmVMGasFn           func(source uint64) sdk.Gas
}

func (m MockGasRegister) NewContractInstanceCosts(pinned bool, msgLen int) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

func (m MockGasRegister) CompileCosts(byteLength int) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

func (m MockGasRegister) InstantiateContractCosts(pinned bool, msgLen int) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

func (m MockGasRegister) ReplyCosts(pinned bool, reply wasmvmtypes.Reply) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

func (m MockGasRegister) EventCosts(evts []wasmvmtypes.EventAttribute, events wasmvmtypes.Events) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

func (m MockGasRegister) ToWasmVMGas(source sdk.Gas) uint64 { _ = "STUB: not implemented"; return 0 }

func (m MockGasRegister) FromWasmVMGas(source uint64) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}
