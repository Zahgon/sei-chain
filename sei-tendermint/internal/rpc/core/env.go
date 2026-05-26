package core

import (
	"context"
	"net"
	"time"

	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/blocksync"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/consensus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventbus"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/eventlog"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/mempool"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/proxy"
	sm "github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/statesync"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const (
	// see README
	defaultPerPage = 30
	maxPerPage     = 100

	// SubscribeTimeout is the maximum time we wait to subscribe for an event.
	// must be less than the server's write timeout (see rpcserver.DefaultConfig)
	SubscribeTimeout = 5 * time.Second

	// genesisChunkSize is the maximum size, in bytes, of each
	// chunk in the genesis structure for the chunked API
	genesisChunkSize = 16 * 1024 * 1024 // 16
)

var logger = seilog.NewLogger("tendermint", "internal", "rpc", "core")

//----------------------------------------------
// These interfaces are used by RPC and must be thread safe

type consensusState interface {
	GetState() sm.State
	GetValidators() (int64, []*types.Validator)
	GetLastHeight() int64
	GetRoundStateJSON() ([]byte, error)
	GetRoundStateSimpleJSON() ([]byte, error)
}

// ----------------------------------------------
// Environment contains objects and interfaces used by the RPC. It is expected
// to be setup once during startup.
type Environment struct {
	// external, thread safe interfaces
	App *proxy.Proxy

	// interfaces defined in types and above
	StateStore       sm.Store
	BlockStore       sm.BlockStore
	EvidencePool     sm.EvidencePool
	ConsensusState   consensusState
	ConsensusReactor *consensus.Reactor
	BlockSyncReactor *blocksync.Reactor

	IsListening bool
	Listeners   []string
	NodeInfo    types.NodeInfo

	Router *p2p.Router

	// objects
	PubKey            utils.Option[crypto.PubKey]
	GenDoc            *types.GenesisDoc // cache the genesis structure
	EventSinks        []indexer.EventSink
	EventBus          *eventbus.EventBus // thread safe
	EventLog          *eventlog.Log
	Mempool           *mempool.TxMempool
	StateSyncMetricer statesync.Metricer

	Config config.RPCConfig

	// cache of chunked genesis data.
	genChunks []string
}

//----------------------------------------------

func validatePage(pagePtr *int, perPage, totalCount int) (int, error) {
	_ = "STUB: not implemented"
	// this can only happen if we haven't first run validatePerPage
	return 0, nil
}

// no page parameter

// one page (even if it's empty)

// gigaRouter returns the GigaRouter when one is wired into env.Router (which
// is the definition of "Autobahn is active for this Environment"). Returns
// None when running under CometBFT or when the Router itself isn't set.
// Handlers that produce different shapes under Autobahn just branch on the
// returned Option:
//
//	if r, ok := env.gigaRouter().Get(); ok {
//	    // Autobahn path, r is the router
//	}
func (env *Environment) gigaRouter() utils.Option[*p2p.GigaRouter] {
	_ = "STUB: not implemented"
	return nil
	// inspect mode
}

func (env *Environment) validatePerPage(perPagePtr *int) int { _ = "STUB: not implemented"; return 0 }

// no per_page parameter

// in unsafe mode there is no max on the page size but in safe mode
// we cap it to maxPerPage

// InitGenesisChunks configures the environment and should be called on service
// startup.
func (env *Environment) InitGenesisChunks() error { _ = "STUB: not implemented"; return nil }

func validateSkipCount(page, perPage int) int { _ = "STUB: not implemented"; return 0 }

// latestHeight can be either latest committed or uncommitted (+1) height.
func (env *Environment) getHeight(latestHeight int64, heightPtr *int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (env *Environment) latestUncommittedHeight() int64 { _ = "STUB: not implemented"; return 0 }

// consensus reactor can be nil in inspect mode.

// StartService constructs and starts listeners for the RPC service
// according to the config object, returning an error if the service
// cannot be constructed or started. The listeners, which provide
// access to the service, run until the context is canceled.
func (env *Environment) StartService(ctx context.Context, conf *config.Config) ([]net.Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435
// Note we don't need to adjust anything if the timeout is already unlimited.

// If the event log is enabled, subscribe to all events published to the
// event bus, and forward them to the event log.

// TODO(creachadair): This is kind of a hack, ideally we'd share the
// observer with the indexer, but it's tricky to plumb them together.
// For now, use a "normal" subscription with a big buffer allowance.
// The event log should always be able to keep up.

// essentially "no limit"

// N.B. Use background for unsubscribe, ctx is already terminated.

// We may expose the RPC over both TCP and a Unix-domain socket.
