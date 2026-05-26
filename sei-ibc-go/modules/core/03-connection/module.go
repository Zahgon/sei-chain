package connection

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
)

// Name returns the IBC connection ICS name.
func Name() string { _ = "STUB: not implemented"; return "" }

// GetQueryCmd returns the root query command for the IBC connections.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// RegisterQueryService registers the gRPC query service for IBC connections.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	_ = "STUB: not implemented"
	return
}
