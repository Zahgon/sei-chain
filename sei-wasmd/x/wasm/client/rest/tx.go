package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router) { _ = "STUB: not implemented"; return }

type storeCodeReq struct {
	BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
	WasmBytes []byte       `json:"wasm_bytes"`
}

type instantiateContractReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Label   string       `json:"label" yaml:"label"`
	Deposit sdk.Coins    `json:"deposit" yaml:"deposit"`
	Admin   string       `json:"admin,omitempty" yaml:"admin"`
	Msg     []byte       `json:"msg" yaml:"msg"`
}

type executeContractReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	ExecMsg []byte       `json:"exec_msg" yaml:"exec_msg"`
	Amount  sdk.Coins    `json:"coins" yaml:"coins"`
}

func storeCodeHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// gzip the wasm file

// build and sign the transaction, then broadcast to Tendermint

func instantiateContractHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// get the id of the code to instantiate

func executeContractHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
