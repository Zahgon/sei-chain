package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/rest"
)

func registerNewTxRoutes(cliCtx client.Context, r *mux.Router) { _ = "STUB: not implemented"; return }

type migrateContractReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Admin   string       `json:"admin,omitempty" yaml:"admin"`
	CodeID  uint64       `json:"code_id" yaml:"code_id"`
	Msg     []byte       `json:"msg,omitempty" yaml:"msg"`
}

type updateContractAdministrateReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Admin   string       `json:"admin,omitempty" yaml:"admin"`
}

func setContractAdminHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func migrateContractHandlerFn(cliCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
