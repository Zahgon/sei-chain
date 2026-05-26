package rest

import (
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
)

type (
	// DecodeReq defines a tx decoding request.
	DecodeReq struct {
		Tx string `json:"tx"`
	}

	// DecodeResp defines a tx decoding response.
	DecodeResp legacytx.StdTx
)

// DecodeTxRequestHandlerFn returns the decode tx REST handler. In particular,
// it takes base64-decoded bytes, decodes it from the Amino wire protocol,
// and responds with a json-formatted transaction.
func DecodeTxRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// NOTE: amino is used intentionally here, don't migrate it

// Error is already returned by convertToStdTx.

// convertToStdTx converts tx proto binary bytes retrieved from Tendermint into
// a StdTx. Returns the StdTx, as well as a flag denoting if the function
// successfully converted or not.
func convertToStdTx(w http.ResponseWriter, clientCtx client.Context, txBytes []byte) (legacytx.StdTx, error) {
	_ = "STUB: not implemented"
	return *new(legacytx.StdTx), nil
}
