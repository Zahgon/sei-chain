package wasmtesting

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

type MockMessageHandler struct {
	DispatchMsgFn func(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) (events []sdk.Event, data [][]byte, err error)
}

func (m *MockMessageHandler) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) (events []sdk.Event, data [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewCapturingMessageHandler() (*MockMessageHandler, *[]wasmvmtypes.CosmosMsg) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return one data item so that this doesn't cause an error in submessage processing (it takes the first element from data)

func NewErroringMessageHandler() *MockMessageHandler { _ = "STUB: not implemented"; return nil }
