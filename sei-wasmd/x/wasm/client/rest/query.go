package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) { _ = "STUB: not implemented"; return }

func listCodesHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryCodeHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func listContractsByCodeHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryContractHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryContractStateAllHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// parse res

func queryContractStateRawHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// ensure this is base64 encoded

type smartResponse struct {
	Smart []byte `json:"smart"`
}

func queryContractStateSmartHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// return as raw bytes (to be base64-encoded)

func queryContractHistoryFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

type argumentDecoder struct {
	// dec is the default decoder
	dec      func(string) ([]byte, error)
	encoding string
}

func newArgDecoder(def func(string) ([]byte, error)) *argumentDecoder {
	_ = "STUB: not implemented"
	return nil
}

func (a *argumentDecoder) DecodeString(s string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
