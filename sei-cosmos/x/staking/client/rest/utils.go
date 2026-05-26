package rest

import (
	"context"
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// contains checks if the a given query contains one of the tx types
func contains(stringSlice []string, txType string) bool { _ = "STUB: not implemented"; return false }

// queries staking txs
func queryTxs(ctx context.Context, clientCtx client.Context, action string, delegatorAddr string) (*sdk.SearchTxsResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryBonds(clientCtx client.Context, endpoint string) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryDelegator(clientCtx client.Context, endpoint string) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func queryValidator(clientCtx client.Context, endpoint string) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
