package client

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

// Name returns the IBC client name
func Name() string { _ = "STUB: not implemented"; return "" }

// GetQueryCmd returns no root query command for the IBC client
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetTxCmd returns the root tx command for 02-client.
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// RegisterQueryService registers the gRPC query service for IBC client.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	_ = "STUB: not implemented"
	return
}
