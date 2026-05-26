package node

import (
	"context"
	"errors"
	_ "net/http/pprof" // nolint: gosec // securely exposed on separate, optional port

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	atypes "github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/evidence"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/store"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/privval"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/version"
	dbm "github.com/tendermint/tm-db"
)

// ErrGenesisMaxGasInvalid is returned by buildGigaConfig when the genesis
// consensus_params.block.max_gas is missing or non-positive. Producer.MaxGasPerBlock
// must be a positive integer derived from this value; tests assert via errors.Is.
var ErrGenesisMaxGasInvalid = errors.New("genesis consensus_params.block.max_gas must be > 0")

type closer func() error

func makeCloser(cs []closer) closer { _ = "STUB: not implemented"; return *new(closer) }

func convertCancelCloser(cancel context.CancelFunc) closer {
	_ = "STUB: not implemented"
	return *new(closer)
}

func combineCloseError(err error, cl closer) error { _ = "STUB: not implemented"; return nil }

func initDBs(
	cfg *config.Config,
	dbProvider config.DBProvider,
) (*store.BlockStore, dbm.DB, closer, error) {
	_ = "STUB: not implemented"
	return nil, *new(dbm.DB), *new(closer), nil
}

func logNodeStartupInfo(state sm.State, pubKey utils.Option[crypto.PubKey], mode string) {
	_ = "STUB: not implemented"
	// Log the version info.
	return
}

// If the state and software differ in block version, at least log it.

// Log whether this node is a validator or an observer

func onlyValidatorIsUs(state sm.State, pubKey utils.Option[crypto.PubKey]) bool {
	_ = "STUB: not implemented"
	return false
}

func createEvidenceReactor(
	cfg *config.Config,
	dbProvider config.DBProvider,
	store sm.Store,
	blockStore *store.BlockStore,
	router *p2p.Router,
	metrics *evidence.Metrics,
	eventBus *eventbus.EventBus,
) (*evidence.Reactor, *evidence.Pool, closer, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(closer), nil
}

func loadAutobahnFileConfig(path string) (*config.AutobahnFileConfig, error) {
	_ = "STUB: not implemented"
	return nil,
		//nolint:gosec // G304: path is from operator-controlled config
		nil
}

// buildGigaConfig constructs a GigaRouterConfig from the autobahn config file, node key, and genesis doc.
func buildGigaConfig(
	autobahnConfigFile string,
	nodeKey types.NodeKey,
	validatorKey atypes.SecretKey,
	txMempool *mempool.TxMempool,
	genDoc *types.GenesisDoc,
) (*p2p.GigaRouterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify self is in the validator set.

// The producer's max-gas-per-block is the chain's gas-limit consensus
// rule, which lives in genesis (consensus_params.block.max_gas) — the
// same number the EVM runtime reads via ctx.ConsensusParams().Block.MaxGas.

//nolint:gosec // validated > 0 above

func createRouter(
	p2pMetrics *p2p.Metrics,
	nodeInfoProducer func() *types.NodeInfo,
	nodeKey types.NodeKey,
	validatorKey utils.Option[atypes.SecretKey],
	cfg *config.Config,
	txMempool utils.Option[*mempool.TxMempool],
	genDoc *types.GenesisDoc,
	dbProvider config.DBProvider,
) (*p2p.Router, closer, error) {
	_ = "STUB: not implemented"
	return nil, *new(closer), nil
}

// MaxConnections defaults to 64

// MaxOutbound defaults to 20, unless MaxConnections<40,
// then it defaults to half of the maxConnections.

// MaxInbound is simply MaxConnections - MaxOutbound,
// because now we have totally separate inbound and outbound connection pools.
// TODO(gprusak): eventually we should migrate configs to specify
// MaxInbound and MaxOutbound explicitly, rather than doing the computation above.

// Wire up Autobahn (GigaRouter) if enabled.

// TODO: add support for autobahn non-validator (observer) nodes that don't need a signing key.

// Resolve a relative persistent_state_dir against the node's --home dir,
// matching how other paths in the tendermint config are handled
// (config.go's rootify). Absolute paths pass through unchanged. None
// means the operator opted into in-memory-only mode and stays None.

func makeNodeInfo(
	cfg *config.Config,
	nodeKey types.NodeKey,
	eventSinks []indexer.EventSink,
	genDoc *types.GenesisDoc,
	versionInfo version.Consensus,
) (types.NodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeInfo), nil
}

// global

func makeSeedNodeInfo(
	cfg *config.Config,
	nodeKey types.NodeKey,
	genDoc *types.GenesisDoc,
	state sm.State,
) (types.NodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(types.NodeInfo), nil
}

// global

func createAndStartPrivValidatorSocketClient(
	ctx context.Context,
	listenAddr, chainID string,
) (types.PrivValidator, error) {
	_ = "STUB: not implemented"
	return *new(types.PrivValidator), nil
}

// try to get a pubkey from private validate first time

func createAndStartPrivValidatorGRPCClient(
	ctx context.Context,
	cfg *config.Config,
	chainID string,
) (types.PrivValidator, error) {
	_ = "STUB: not implemented"
	return *new(types.PrivValidator), nil
}

// try to get a pubkey from private validate first time

func createPrivval(ctx context.Context, conf *config.Config, genDoc *types.GenesisDoc, defaultPV *privval.FilePV) (types.PrivValidator, error) {
	_ = "STUB: not implemented"
	return *new(types.PrivValidator), nil
}

// FIXME: we should return un-started services and
// then start them later.
