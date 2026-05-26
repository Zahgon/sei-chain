package proxy

import (
	"context"
	"net"
	"net/http"

	"github.com/sei-protocol/sei-chain/sei-tendermint/light"
	lrpc "github.com/sei-protocol/sei-chain/sei-tendermint/light/rpc"
	rpcserver "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/server"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "light", "proxy")

// A Proxy defines parameters for running an HTTP server proxy.
type Proxy struct {
	Addr     string // TCP address to listen on, ":http" if empty
	Config   *rpcserver.Config
	Client   *lrpc.Client
	Listener net.Listener
}

// NewProxy creates the struct used to run an HTTP server for serving light
// client rpc requests.
func NewProxy(
	lightClient *light.Client,
	listenAddr, providerAddr string,
	config *rpcserver.Config,
	opts ...lrpc.Option,
) (*Proxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListenAndServe configures the rpcserver.WebsocketManager, sets up the RPC
// routes to proxy via Client, and starts up an HTTP server on the TCP network
// address p.Addr.
// See http#Server#ListenAndServe.
func (p *Proxy) ListenAndServe(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// ListenAndServeTLS acts identically to ListenAndServe, except that it expects
// HTTPS connections.
// See http#Server#ListenAndServeTLS.
func (p *Proxy) ListenAndServeTLS(ctx context.Context, certFile, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Proxy) listen(ctx context.Context) (net.Listener, *http.ServeMux, error) {
	_ = "STUB: not implemented"
	return *

	// 1) Register regular routes.
	new(net.Listener), nil, nil
}

// 2) Allow websocket connections.

// 3) Start a client.

// 4) Start listening for new connections.
