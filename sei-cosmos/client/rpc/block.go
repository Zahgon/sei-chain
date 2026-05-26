package rpc

import (
	"context"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

// BlockCommand returns the verified block data for a given heights
func BlockCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// optional height

func getBlock(clientCtx client.Context, height *int64) ([]byte, error) {
	_ = "STUB: not implemented"
	// get the node
	return nil, nil
}

// header -> BlockchainInfo
// header, tx -> Block
// results -> BlockResults

// get the current blockchain height
func GetChainHeight(ctx context.Context, node client.Client) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// REST handler to get a block
func BlockRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// REST handler to get the latest block
func LatestBlockRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
