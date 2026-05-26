package debug

import (
	"context"

	"github.com/sei-protocol/seilog"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	rpchttp "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/http"
)

var logger = seilog.NewLogger("tendermint", "cmd", "tendermint", "commands", "debug")

func getDumpCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

//nolint: gosec

type dumpDebugDataArgs struct {
	conf     *config.Config
	outDir   string
	profAddr string
}

func dumpDebugData(ctx context.Context, rpc *rpchttp.HTTP, args dumpDebugDataArgs) {
	_ = "STUB: not implemented"
	return
}
