package inspect

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	rpccore "github.com/sei-protocol/sei-chain/sei-tendermint/internal/rpc/core"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "internal", "inspect")

// Inspector manages an RPC service that exports methods to debug a failed node.
// After a node shuts down due to a consensus failure, it will no longer start
// up its state cannot easily be inspected. An Inspector value provides a similar interface
// to the node, using the underlying Tendermint data stores, without bringing up
// any other components. A caller can query the Inspector service to inspect the
// persisted state and debug the failure.
type Inspector struct {
	routes rpccore.RoutesMap

	config *config.RPCConfig

	indexerService *indexer.Service
	eventBus       *eventbus.EventBus
}

// New returns an Inspector that serves RPC on the specified BlockStore and StateStore.
// The Inspector type does not modify the state or block stores.
// The sinks are used to enable block and transaction querying via the RPC server.
// The caller is responsible for starting and stopping the Inspector service.
func New(cfg *config.RPCConfig, bs state.BlockStore, ss state.Store, es []indexer.EventSink) *Inspector {
	_ = "STUB: not implemented"
	return nil
}

// NewFromConfig constructs an Inspector using the values defined in the passed in config.
func NewFromConfig(cfg *config.Config) (*Inspector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run starts the Inspector servers and blocks until the servers shut down. The passed
// in context is used to control the lifecycle of the servers.
func (ins *Inspector) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func startRPCServers(ctx context.Context, cfg *config.RPCConfig, routes rpccore.RoutesMap) error {
	_ = "STUB: not implemented"
	return nil
}
