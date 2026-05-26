package port

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/05-port/types"
)

// Name returns the IBC port ICS name.
func Name() string { _ = "STUB: not implemented"; return "" }

// GetQueryCmd returns the root query command for IBC ports.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// RegisterQueryService registers the gRPC query service for IBC ports.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	_ = "STUB: not implemented"
	return
}
