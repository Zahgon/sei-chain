package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

// newWasmModuleEvent creates with wasm module event for interacting with the given contract. Adds custom attributes
// to this event.
func newWasmModuleEvent(customAttributes []wasmvmtypes.EventAttribute, contractAddr sdk.AccAddress) (sdk.Events, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Events), nil
}

// each wasm invocation always returns one sdk.Event

const eventTypeMinLength = 2

// newCustomEvents converts wasmvm events from a contract response to sdk type events
func newCustomEvents(evts wasmvmtypes.Events, contractAddr sdk.AccAddress) (sdk.Events, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Events), nil
}

// convert and add contract address issuing this event
func contractSDKEventAttributes(customAttributes []wasmvmtypes.EventAttribute, contractAddr sdk.AccAddress) ([]sdk.Attribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// append attributes from wasm to the sdk.Event

// ensure key and value are non-empty (and trim what is there)

// TODO: check if this is legal in the SDK - if it is, we can remove this check

// and reserve all _* keys for our use (not contract)
