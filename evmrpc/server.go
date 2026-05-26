package evmrpc

import (
	"github.com/sei-protocol/sei-chain/app/legacyabci"
	evmrpcconfig "github.com/sei-protocol/sei-chain/evmrpc/config"
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type ConnectionType string

var ConnectionTypeWS ConnectionType = "websocket"
var ConnectionTypeHTTP ConnectionType = "http"

const LocalAddress = "0.0.0.0"
const DefaultWebsocketMaxMessageSize = 10 * 1024 * 1024

type EVMServer interface {
	Start() error
	Stop()
}

func NewEVMHTTPServer(
	config evmrpcconfig.Config,
	tmClient client.LocalClient,
	k *keeper.Keeper,
	beginBlockKeepers legacyabci.BeginBlockKeepers,
	app *baseapp.BaseApp,
	antehandler sdk.AnteHandler,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	homeDir string,
	stateStore types.StateStore,
	traceCtxProviders ...TraceContextProvider,
) (EVMServer, error) {
	_ = "STUB: not implemented"

	// Initialize global worker pool with configuration (metrics are embedded in pool)
	return *new(EVMServer), nil
}

// Get pool for logging and DB semaphore setup

// Set DB semaphore capacity in metrics (aligned with worker count)
// Only set once to avoid races when multiple test servers start in parallel.
//nolint:gosec // G115: safe, max is 64

// Initialize RPC tracker

// DB semaphore aligned with worker count

//nolint:gosec

// Test API can only exist on non-live chain IDs.  These APIs instrument certain overrides.

func NewEVMWebSocketServer(
	config evmrpcconfig.Config,
	tmClient client.LocalClient,
	k *keeper.Keeper,
	beginBlockKeepers legacyabci.BeginBlockKeepers,
	app *baseapp.BaseApp,
	antehandler sdk.AnteHandler,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	homeDir string,
	stateStore types.StateStore,
	blockHeaderNotifier *BlockHeaderNotifier,
) (EVMServer, error) {
	_ = "STUB: not implemented"
	// Initialize global worker pool with configuration (metrics are embedded in pool)
	// This is idempotent - if HTTP server already initialized it, this is a no-op
	return *new(EVMServer), nil
}

// Initialize WebSocket tracker.

// DB semaphore aligned with worker count
