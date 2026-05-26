package wasmtesting

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// MockMessageRouter mock for testing
type MockMessageRouter struct {
	HandlerFn func(msg sdk.Msg) baseapp.MsgServiceHandler
}

// Handler is the entry point
func (m MockMessageRouter) Handler(msg sdk.Msg) baseapp.MsgServiceHandler {
	_ = "STUB: not implemented"
	return *new(baseapp.MsgServiceHandler)
}

// MessageRouterFunc convenient type to match the keeper.MessageRouter interface
type MessageRouterFunc func(msg sdk.Msg) baseapp.MsgServiceHandler

// Handler is the entry point
func (m MessageRouterFunc) Handler(msg sdk.Msg) baseapp.MsgServiceHandler {
	_ = "STUB: not implemented"
	return *new(baseapp.MsgServiceHandler)
}
