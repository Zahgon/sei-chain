package node

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	rpccore "github.com/sei-protocol/sei-chain/sei-tendermint/internal/rpc/core"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/local"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type seedNodeImpl struct {
	service.BaseService

	// config
	config     *config.Config
	genesisDoc *types.GenesisDoc // initial validator set

	nodeInfo types.NodeInfo

	// network
	router      *p2p.Router
	nodeKey     types.NodeKey // our node privkey
	isListening bool

	// services
	pexReactor  service.Service // for exchanging peer addresses
	shutdownOps closer
	rpcEnv      *rpccore.Environment
}

// makeSeedNode returns a new seed node, containing only p2p, pex reactor
func makeSeedNode(
	cfg *config.Config,
	dbProvider config.DBProvider,
	nodeKey types.NodeKey,
	genesisDocProvider genesisDocProvider,
	nodeMetrics *NodeMetrics,
) (_ local.NodeService, err error) {
	_ = "STUB: not implemented"
	return *new(local.NodeService), nil
}

// OnStart starts the Seed Node. It implements service.Service.
func (n *seedNodeImpl) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G112: mitigate slowloris attacks

// Start the transport.

// OnStop stops the Seed Node. It implements service.Service.
func (n *seedNodeImpl) OnStop() { _ = "STUB: not implemented"; return }

// EventBus returns the Node's EventBus.
func (n *seedNodeImpl) EventBus() *eventbus.EventBus { _ = "STUB: not implemented"; return nil }

// RPCEnvironment makes sure RPC has all the objects it needs to operate.
func (n *seedNodeImpl) RPCEnvironment() *rpccore.Environment { _ = "STUB: not implemented"; return nil }

func (n *seedNodeImpl) NodeInfo() *types.NodeInfo { _ = "STUB: not implemented"; return nil }
