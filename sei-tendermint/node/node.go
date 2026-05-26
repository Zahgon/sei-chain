package node

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof" // nolint: gosec // securely exposed on separate, optional port

	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/consensus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventlog"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/evidence"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	rpccore "github.com/sei-protocol/sei-chain/sei-tendermint/internal/rpc/core"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/statesync"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/store"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	"github.com/sei-protocol/sei-chain/sei-tendermint/privval"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/local"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"

	_ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"

	_ "github.com/lib/pq" // provide the psql db driver
)

// nodeImpl is the highest level interface to a full Tendermint node.
// It includes all configuration information and running services.
type nodeImpl struct {
	service.BaseService

	// config
	config          *config.Config
	genesisDoc      *types.GenesisDoc   // initial validator set
	privValidator   types.PrivValidator // local node's validator key
	shouldHandshake bool                // set during makeNode
	consensusPolicy types.ConsensusPolicy

	// network
	router           *p2p.Router
	ServiceRestartCh chan []string
	nodeInfo         types.NodeInfo
	nodeKey          types.NodeKey // our node privkey

	// services
	eventSinks     []indexer.EventSink
	initialState   sm.State
	stateStore     sm.Store
	blockStore     *store.BlockStore // store the blockchain to disk
	mempool        *mempool.TxMempool
	evPool         *evidence.Pool
	indexerService *indexer.Service
	services       []service.Service
	rpcListeners   []net.Listener // rpc servers
	shutdownOps    closer
	rpcEnv         *rpccore.Environment
	prometheusSrv  *http.Server
}

// makeNode returns a new, ready to go, Tendermint Node.
func makeNode(
	ctx context.Context,
	cfg *config.Config,
	restartEvent func(),
	filePrivval *privval.FilePV,
	nodeKey types.NodeKey,
	proxyApp *proxy.Proxy,
	genesisDocProvider genesisDocProvider,
	dbProvider config.DBProvider,
	tracerProviderOptions []trace.TracerProviderOption,
	nodeMetrics *NodeMetrics,
	consensusPolicy types.ConsensusPolicy,
) (_ local.NodeService, err error) {
	_ = "STUB: not implemented"
	return *new(local.NodeService), nil
}

// TODO construct node here:

// Autobahn requires a local validator key; remote signers are not supported.

// Mempool gossiping is not compatible with Giga,
// so we disable the mempool reactor.

// make block executor for consensus and blockchain reactors to execute blocks

// Determine whether we should attempt state sync.

// Determine whether we should do block sync. This must happen after the handshake, since the
// app may modify the validator set, specifying ourself as the only validator.

// TODO(autobahn-recovery): handles only restart with local disk intact.
// A node that lost its WAL + app CMS (new validator, disk wipe) needs both
// app state sync and an autobahn WAL sync to catch up. Not yet supported.

// Create the blockchain reactor. Note, we do not start block sync if we're
// doing a state sync first.

// Make ConsensusReactor. Don't enable fully if doing a state sync and/or block sync first.
// FIXME We need to update metrics here, since other reactors don't have access to them.

// TODO: Some form of orchestrator is needed here between the state
// advancing reactors to be able to control which one of the three
// is running
// FIXME Very ugly to have these metrics bleed through here.

// Set up state sync reactor, and schedule a sync if requested.
// FIXME The way we do phased startups (e.g. replay -> block sync -> consensus) is very messy,
// we should clean this whole thing up. See:
// https://github.com/tendermint/tendermint/issues/4644
// The CometBFT handshaker reconciles the block store and state store with the app
// by replaying blocks and calling InitChain at genesis. Autobahn (giga) maintains
// its own data WAL and does not update the CometBFT block/state stores, so on
// restart the handshaker would observe storeHeight=0 < appHeight=N and fail with
// ErrAppBlockHeightTooHigh. We skip the handshaker in giga mode; instead the
// giga router's runExecute owns InitChain on fresh start (appHeight==0) and
// relies on the app's committed CMS to rebuild deliverState on restart.

// the post-sync operation

// OnStart starts the Node. It implements service.Service.
func (n *nodeImpl) OnStart(ctx context.Context) error {
	_ = "STUB: not implemented"
	// EventBus and IndexerService must be started before the handshake because
	// we might need to index the txs of the replayed block as this might not have happened
	// when the node stopped last time (i.e. the node stopped or crashed after it saved the block
	// but before it indexed the txs)
	return nil
}

// state sync will cover initialization the chain. Also calling InitChain isn't safe
// when there is state sync as InitChain itself doesn't commit application state which
// would get mixed up with application state writes by state sync.

// Create the handshaker, which calls RequestInfo, sets the AppVersion on the state,
// and replays any blocks as necessary to sync tendermint with the app.

// Reload the state. It will have the Version.Consensus.App set by the
// Handshake, and may have other modifications as well (ie. depending on
// what happened during block replay).

// TODO: Fetch and provide real options and do proper p2p bootstrapping.
// TODO: Use a persistent peer database.

// Start Internal Services

// Start the transport.

// Start the RPC server before the P2P server
// so we can eg. receive txs for the first block

// OnStop stops the Node. It implements service.Service.
func (n *nodeImpl) OnStop() { _ = "STUB: not implemented"; return }

// stop the listeners / external services first

// Error from closing listeners, or context timeout:

// startPrometheusServer starts a Prometheus HTTP server, listening for metrics
// collectors on addr.
func (n *nodeImpl) startPrometheusServer(ctx context.Context, addr string) *http.Server {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G112: mitigate slowloris attacks

func (n *nodeImpl) NodeInfo() *types.NodeInfo {
	_ = "STUB: not implemented"

	// EventBus returns the Node's EventBus.
	return nil
}

func (n *nodeImpl) EventBus() *eventbus.EventBus { _ = "STUB: not implemented"; return nil }

// GenesisDoc returns the Node's GenesisDoc.
func (n *nodeImpl) GenesisDoc() *types.GenesisDoc { _ = "STUB: not implemented"; return nil }

// RPCEnvironment makes sure RPC has all the objects it needs to operate.
func (n *nodeImpl) RPCEnvironment() *rpccore.Environment {
	_ = "STUB: not implemented"

	// ------------------------------------------------------------------------------
	return nil
}

// genesisDocProvider returns a GenesisDoc.
// It allows the GenesisDoc to be pulled from sources other than the
// filesystem, for instance from a distributed key-value store cluster.
type genesisDocProvider func() (*types.GenesisDoc, error)

// defaultGenesisDocProviderFunc returns a GenesisDocProvider that loads
// the GenesisDoc from the config.GenesisFile() on the filesystem.
func defaultGenesisDocProviderFunc(cfg *config.Config) genesisDocProvider {
	_ = "STUB: not implemented"
	return *new(genesisDocProvider)
}

type NodeMetrics struct {
	consensus *consensus.Metrics
	eventlog  *eventlog.Metrics
	indexer   *indexer.Metrics
	mempool   *mempool.Metrics
	p2p       *p2p.Metrics
	proxy     *proxy.Metrics
	state     *sm.Metrics
	statesync *statesync.Metrics
	evidence  *evidence.Metrics
}

// metricsProvider returns consensus, p2p, mempool, state, statesync Metrics.
type metricsProvider func(chainID string) *NodeMetrics

func NoOpMetricsProvider() *NodeMetrics { _ = "STUB: not implemented"; return nil }

// defaultMetricsProvider returns Metrics build using Prometheus client library
// if Prometheus is enabled. Otherwise, it returns no-op Metrics.
func DefaultMetricsProvider(cfg *config.InstrumentationConfig) metricsProvider {
	_ = "STUB: not implemented"
	return *new(metricsProvider)
}

//------------------------------------------------------------------------------

// LoadStateFromDBOrGenesisDocProvider attempts to load the state from the
// database, or creates one using the given genesisDocProvider. On success this also
// returns the genesis doc loaded through the given provider.
func LoadStateFromDBOrGenesisDocProvider(stateStore sm.Store, genDoc *types.GenesisDoc) (sm.State, error) {
	_ = "STUB: not implemented"

	// 1. Attempt to load state form the database
	return *new(sm.State), nil
}

// 2. If it's not there, derive it from the genesis doc

// 3. save the gensis document to the state store so
// its fetchable by other callers.
