package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) { _ = "STUB: not implemented"; return }

// Unjail TX body
type UnjailReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
}

// NewUnjailRequestHandlerFn returns an HTTP REST handler for creating a MsgUnjail
// transaction.
func NewUnjailRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
