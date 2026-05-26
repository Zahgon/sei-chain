package keeper

import (
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

const (
	// DefaultGasMultiplier is how many CosmWasm gas points = 1 Cosmos SDK gas point.
	//
	// CosmWasm gas strategy is documented in https://github.com/CosmWasm/cosmwasm/blob/v1.0.0-beta/docs/GAS.md.
	// Cosmos SDK reference costs can be found here: https://github.com/cosmos/cosmos-sdk/blob/v0.42.10/store/types/gas.go#L198-L209.
	//
	// The original multiplier of 100 up to CosmWasm 0.16 was based on
	//     "A write at ~3000 gas and ~200us = 10 gas per us (microsecond) cpu/io
	//     Rough timing have 88k gas at 90us, which is equal to 1k sdk gas... (one read)"
	// as well as manual Wasmer benchmarks from 2019. This was then multiplied by 150_000
	// in the 0.16 -> 1.0 upgrade (https://github.com/CosmWasm/cosmwasm/pull/1120).
	//
	// The multiplier deserves more reproducible benchmarking and a strategy that allows easy adjustments.
	// This is tracked in https://github.com/CosmWasm/wasmd/issues/566 and https://github.com/CosmWasm/wasmd/issues/631.
	// Gas adjustments are consensus breaking but may happen in any release marked as consensus breaking.
	// Do not make assumptions on how much gas an operation will consume in places that are hard to adjust,
	// such as hardcoding them in contracts.
	//
	// Please note that all gas prices returned to wasmvm should have this multiplied.
	// Benchmarks and numbers were discussed in: https://github.com/CosmWasm/wasmd/pull/634#issuecomment-938055852
	DefaultGasMultiplier uint64 = 140_000_000
	// DefaultInstanceCost is how much SDK gas we charge each time we load a WASM instance.
	// Creating a new instance is costly, and this helps put a recursion limit to contracts calling contracts.
	// Benchmarks and numbers were discussed in: https://github.com/CosmWasm/wasmd/pull/634#issuecomment-938056803
	DefaultInstanceCost uint64 = 60_000
	// DefaultCompileCost is how much SDK gas is charged *per byte* for compiling WASM code.
	// Benchmarks and numbers were discussed in: https://github.com/CosmWasm/wasmd/pull/634#issuecomment-938056803
	DefaultCompileCost uint64 = 3
	// DefaultEventAttributeDataCost is how much SDK gas is charged *per byte* for attribute data in events.
	// This is used with len(key) + len(value)
	DefaultEventAttributeDataCost uint64 = 1
	// DefaultContractMessageDataCost is how much SDK gas is charged *per byte* of the message that goes to the contract
	// This is used with len(msg). Note that the message is deserialized in the receiving contract and this is charged
	// with wasm gas already. The derserialization of results is also charged in wasmvm. I am unsure if we need to add
	// additional costs here.
	// Note: also used for error fields on reply, and data on reply. Maybe these should be pulled out to a different (non-zero) field
	DefaultContractMessageDataCost uint64 = 0
	// DefaultPerAttributeCost is how much SDK gas we charge per attribute count.
	DefaultPerAttributeCost uint64 = 10
	// DefaultPerCustomEventCost is how much SDK gas we charge per event count.
	DefaultPerCustomEventCost uint64 = 20
	// DefaultEventAttributeDataFreeTier number of bytes of total attribute data we do not charge.
	DefaultEventAttributeDataFreeTier = 100
)

// GasRegister abstract source for gas costs
type GasRegister interface {
	// NewContractInstanceCosts costs to crate a new contract instance from code
	NewContractInstanceCosts(pinned bool, msgLen int) sdk.Gas
	// CompileCosts costs to persist and "compile" a new wasm contract
	CompileCosts(byteLength int) sdk.Gas
	// InstantiateContractCosts costs when interacting with a wasm contract
	InstantiateContractCosts(pinned bool, msgLen int) sdk.Gas
	// ReplyCosts costs to to handle a message reply
	ReplyCosts(pinned bool, reply wasmvmtypes.Reply) sdk.Gas
	// EventCosts costs to persist an event
	EventCosts(attrs []wasmvmtypes.EventAttribute, events wasmvmtypes.Events) sdk.Gas
	// ToWasmVMGas converts from sdk gas to wasmvm gas
	ToWasmVMGas(source sdk.Gas) uint64
	// FromWasmVMGas converts from wasmvm gas to sdk gas
	FromWasmVMGas(source uint64) sdk.Gas
}

// WasmGasRegisterConfig config type
type WasmGasRegisterConfig struct {
	// InstanceCost costs when interacting with a wasm contract
	InstanceCost sdk.Gas
	// CompileCosts costs to persist and "compile" a new wasm contract
	CompileCost sdk.Gas
	// GasMultiplier is how many cosmwasm gas points = 1 sdk gas point
	// SDK reference costs can be found here: https://github.com/cosmos/cosmos-sdk/blob/02c6c9fafd58da88550ab4d7d494724a477c8a68/store/types/gas.go#L153-L164
	GasMultiplier sdk.Gas
	// EventPerAttributeCost is how much SDK gas is charged *per byte* for attribute data in events.
	// This is used with len(key) + len(value)
	EventPerAttributeCost sdk.Gas
	// EventAttributeDataCost is how much SDK gas is charged *per byte* for attribute data in events.
	// This is used with len(key) + len(value)
	EventAttributeDataCost sdk.Gas
	// EventAttributeDataFreeTier number of bytes of total attribute data that is free of charge
	EventAttributeDataFreeTier uint64
	// ContractMessageDataCost SDK gas charged *per byte* of the message that goes to the contract
	// This is used with len(msg)
	ContractMessageDataCost sdk.Gas
	// CustomEventCost cost per custom event
	CustomEventCost uint64
}

// DefaultGasRegisterConfig default values
func DefaultGasRegisterConfig() WasmGasRegisterConfig {
	_ = "STUB: not implemented"
	return *new(WasmGasRegisterConfig)
}

// WasmGasRegister implements GasRegister interface
type WasmGasRegister struct {
	c WasmGasRegisterConfig
}

// NewDefaultWasmGasRegister creates instance with default values
func NewDefaultWasmGasRegister() WasmGasRegister {
	_ = "STUB: not implemented"
	return *new(WasmGasRegister)
}

// NewWasmGasRegister constructor
func NewWasmGasRegister(c WasmGasRegisterConfig) WasmGasRegister {
	_ = "STUB: not implemented"
	return *new(WasmGasRegister)
}

// NewContractInstanceCosts costs to crate a new contract instance from code
func (g WasmGasRegister) NewContractInstanceCosts(pinned bool, msgLen int) storetypes.Gas {
	_ = "STUB: not implemented"
	return *new(storetypes.Gas)
}

// CompileCosts costs to persist and "compile" a new wasm contract
func (g WasmGasRegister) CompileCosts(byteLength int) storetypes.Gas {
	_ = "STUB: not implemented"
	return *new(storetypes.Gas)
}

// InstantiateContractCosts costs when interacting with a wasm contract
func (g WasmGasRegister) InstantiateContractCosts(pinned bool, msgLen int) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

// ReplyCosts costs to handle a message reply
func (g WasmGasRegister) ReplyCosts(pinned bool, reply wasmvmtypes.Reply) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

//nolint:prealloc // not worth it

// apply free tier on the whole set not per event

// EventCosts costs to persist an event
func (g WasmGasRegister) EventCosts(attrs []wasmvmtypes.EventAttribute, events wasmvmtypes.Events) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}

// no free tier with event type

func (g WasmGasRegister) eventAttributeCosts(attrs []wasmvmtypes.EventAttribute, freeTier uint64) (sdk.Gas, uint64) {
	_ = "STUB: not implemented"
	return *new(sdk.Gas), 0
}

// total Length * costs + attribute count * costs

// apply free tier
func calcWithFreeTier(storedBytes uint64, freeTier uint64) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// ToWasmVMGas convert to wasmVM contract runtime gas unit
func (g WasmGasRegister) ToWasmVMGas(source storetypes.Gas) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// FromWasmVMGas converts to SDK gas unit
func (g WasmGasRegister) FromWasmVMGas(source uint64) sdk.Gas {
	_ = "STUB: not implemented"
	return *new(sdk.Gas)
}
