package cmd

import (
	"github.com/sei-protocol/sei-chain/admin/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func LogLevelCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func logLevelSubCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func logLevelSetCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func logLevelGetCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func logLevelListCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func adminClient(cmd *cobra.Command) (types.AdminServiceClient, *grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return *new(types.AdminServiceClient), nil, nil
}
