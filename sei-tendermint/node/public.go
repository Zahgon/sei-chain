// Package node provides a high level wrapper around tendermint services.
package node

import (
	"context"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/local"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/seilog"
	"go.opentelemetry.io/otel/sdk/trace"
)

var logger = seilog.NewLogger("tendermint", "node")

// New constructs a tendermint node. The provided app runs in the same
// process as the tendermint node and will be wrapped in a local ABCI client
// inside this function. The final option is a pointer to a Genesis document:
// if the value is nil, the genesis document is read from the file specified
// in the config, and otherwise the node uses value of the final argument.
func New(
	ctx context.Context,
	conf *config.Config,
	restartEvent func(),
	app abci.Application,
	gen *tmtypes.GenesisDoc,
	tracerProviderOptions []trace.TracerProviderOption,
	nodeMetrics *NodeMetrics,
	consensusPolicy tmtypes.ConsensusPolicy,
) (local.NodeService, error) {
	_ = "STUB: not implemented"
	return *new(local.NodeService), nil
}
