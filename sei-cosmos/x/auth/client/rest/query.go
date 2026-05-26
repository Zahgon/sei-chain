package rest

import (
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// QueryAccountRequestHandlerFn is the query accountREST Handler.
func QueryAccountRequestHandlerFn(storeName string, clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// TODO: Handle more appropriately based on the error type.
// Ref: https://github.com/cosmos/cosmos-sdk/issues/4923

// QueryTxsRequestHandlerFn implements a REST handler that searches for transactions.
// Genesis transactions are returned if the height parameter is set to zero,
// otherwise the transactions are searched for by events.
func QueryTxsRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// if the height query param is set to zero, query for genesis transactions

// QueryTxRequestHandlerFn implements a REST handler that queries a transaction
// by hash in a committed block.
func QueryTxRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Error is already returned by packStdTxResponse.

func queryParamsHandler(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// packStdTxResponse takes a sdk.TxResponse, converts the Tx into a StdTx, and
// packs the StdTx again into the sdk.TxResponse Any. Amino then takes care of
// seamlessly JSON-outputting the Any.
func packStdTxResponse(w http.ResponseWriter, clientCtx client.Context, txRes *sdk.TxResponse) error {
	_ = "STUB: not implemented"
	// We just unmarshalled from Tendermint, we take the proto Tx's raw
	// bytes, and convert them into a StdTx to be displayed.
	return nil
}

// Pack the amino stdTx into the TxResponse's Any.

// checkAminoMarshalError checks if there are errors with marshalling non-amino
// txs with amino.
func checkAminoMarshalError(ctx client.Context, resp any, grpcEndPoint string) error {
	_ = "STUB: not implemented"
	// LegacyAmino used intentionally here to handle the SignMode errors
	return nil
}

// If there's an unmarshalling error, we assume that it's because we're
// using amino to unmarshal a non-amino tx.
