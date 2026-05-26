package tests

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/app"
	"github.com/sei-protocol/sei-chain/evmrpc"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type RpcResponse struct {
	Jsonrpc string                     `json:"jsonrpc"`
	Id      int                        `json:"id"`
	Result  evmrpc.StateAccessResponse `json:"result"`
}

func mockStatesFromBlockJson(ctx sdk.Context, blockNum uint64, a *app.App, client *MockClient) int64 {
	_ = "STUB: not implemented"
	return 0
}

func mockStatesFromTxJson(ctx sdk.Context, hash string, a *app.App, client *MockClient) int64 {
	_ = "STUB: not implemented"
	return 0
}

func mockStatesFromJsonFile(ctx sdk.Context, filepath string, a *app.App, client *MockClient) int64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec

func mockTendermintStateFromJson(tmStateRaw json.RawMessage, client *MockClient) int64 {
	_ = "STUB: not implemented"
	return 0
}

func mockStateFromJson(ctx sdk.Context, a *app.App, stateRaw json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

// initialize WASM code

func mockReceipt(ctx sdk.Context, a *app.App, receipts json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

//nolint:gosec
//nolint:gosec

func parseInt64(arg string) int64 { _ = "STUB: not implemented"; return 0 }

func check(e error) { _ = "STUB: not implemented"; return }
