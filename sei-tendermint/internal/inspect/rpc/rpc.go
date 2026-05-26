package rpc

import (
	"context"
	"net/http"

	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/rpc/core"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/state/indexer"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/server"
)

var logger = seilog.NewLogger("tendermint", "internal", "inspect", "rpc")

// Server defines parameters for running an Inspector rpc server.
type Server struct {
	Addr    string // TCP address to listen on, ":http" if empty
	Handler http.Handler
	Config  *config.RPCConfig
}

type eventBusUnsubscriber interface {
	UnsubscribeAll(ctx context.Context, subscriber string) error
}

// Routes returns the set of routes used by the Inspector server.
func Routes(cfg config.RPCConfig, s state.Store, bs state.BlockStore, es []indexer.EventSink) core.RoutesMap {
	_ = "STUB: not implemented"
	return *new(core.RoutesMap)
}

// Handler returns the http.Handler configured for use with an Inspector server. Handler
// registers the routes on the http.Handler and also registers the websocket handler
// and the CORS handler if specified by the configuration options.
func Handler(rpcConfig *config.RPCConfig, routes core.RoutesMap) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func addCORSHandler(rpcConfig *config.RPCConfig, h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// ListenAndServe listens on the address specified in srv.Addr and handles any
// incoming requests over HTTP using the Inspector rpc handler specified on the server.
func (srv *Server) ListenAndServe(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// ListenAndServeTLS listens on the address specified in srv.Addr. ListenAndServeTLS handles
// incoming requests over HTTPS using the Inspector rpc handler specified on the server.
func (srv *Server) ListenAndServeTLS(ctx context.Context, certFile, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func serverRPCConfig(r *config.RPCConfig) *server.Config { _ = "STUB: not implemented"; return nil }

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435
// Note we don't need to adjust anything if the timeout is already unlimited.
