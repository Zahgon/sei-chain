package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// Messenger is an extension point for custom wasmd message handling
type Messenger interface {
	// DispatchMsg encodes the wasmVM message and dispatches it.
	DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) (events []sdk.Event, data [][]byte, err error)
}

// replyer is a subset of keeper that can handle replies to submessages
type replyer interface {
	reply(ctx sdk.Context, contractAddress sdk.AccAddress, reply wasmvmtypes.Reply) ([]byte, error)
}

// MessageDispatcher coordinates message sending and submessage reply/ state commits
type MessageDispatcher struct {
	messenger Messenger
	keeper    replyer
}

// NewMessageDispatcher constructor
func NewMessageDispatcher(messenger Messenger, keeper replyer) *MessageDispatcher {
	_ = "STUB: not implemented"
	return nil
}

// DispatchMessages sends all messages.
func (d MessageDispatcher) DispatchMessages(ctx sdk.Context, contractAddr sdk.AccAddress, ibcPort string, msgs []wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// redispatch all events, (type sdk.EventTypeMessage will be filtered out in the handler)

// dispatchMsgWithGasLimit sends a message with gas limit applied
func (d MessageDispatcher) dispatchMsgWithGasLimit(ctx sdk.Context, contractAddr sdk.AccAddress, ibcPort string, msg wasmvmtypes.CosmosMsg, gasLimit uint64, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) (events []sdk.Event, data [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// catch out of gas panic and just charge the entire gas limit

// if it's not an OutOfGas error, raise it again

// log it to get the original stack trace somewhere (as panic(r) keeps message but stacktrace to here

// make sure we charge the parent what was spent

// DispatchSubmessages builds a sandbox to execute these messages and returns the execution result to the contract
// that dispatched them, both on success as well as failure
func (d MessageDispatcher) DispatchSubmessages(ctx sdk.Context, contractAddr sdk.AccAddress, ibcPort string, msgs []wasmvmtypes.SubMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// first, we build a sub-context which we can use inside the submessages

// check how much gas left locally, optionally wrap the gas meter

// if it succeeds, commit state changes from submessage, and pass on events to Event Manager

// on failure, revert state from sandbox, and ignore events (just skip doing the above)

// we only callback if requested. Short-circuit here the cases we don't want to

// otherwise, we create a SubMsgResult and pass it into the calling contract

// just take the first one for now if there are multiple sub-sdk messages
// and safely return nothing if no data

// Issue #759 - we don't return error string for worries of non-determinism

// now handle the reply, we use the parent context, and abort on error

// we can ignore any result returned as there is nothing to do with the data
// and the events are already in the ctx.EventManager()

// Issue #759 - we don't return error string for worries of non-determinism
func redactError(err error) error {
	_ = "STUB: not implemented"
	// Do not redact system errors
	// SystemErrors must be created in x/wasm and we can ensure determinism
	return nil
}

// FIXME: do we want to hardcode some constant string mappings here as well?
// Or better document them? (SDK error string may change on a patch release to fix wording)
// sdk/11 is out of gas
// sdk/5 is insufficient funds (on bank send)
// (we can theoretically redact less in the future, but this is a first step to safety)

func filterEvents(events []sdk.Event) []sdk.Event {
	_ = "STUB: not implemented"
	// pre-allocate space for efficiency
	return nil
}

func sdkEventsToWasmVMEvents(events []sdk.Event) []wasmvmtypes.Event {
	_ = "STUB: not implemented"
	return nil
}

func sdkAttributesToWasmVMAttributes(attrs []abci.EventAttribute) []wasmvmtypes.EventAttribute {
	_ = "STUB: not implemented"
	return nil
}
