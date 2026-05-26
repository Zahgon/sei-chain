package channel

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// Name returns the IBC channel ICS name.
func Name() string { _ = "STUB: not implemented"; return "" }

// GetTxCmd returns the root tx command for IBC channels.
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd returns the root query command for IBC channels.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// RegisterQueryService registers the gRPC query service for IBC channels.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	_ = "STUB: not implemented"
	return
}
