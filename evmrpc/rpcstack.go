// Copyright 2020 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package evmrpc

import (
	"compress/gzip"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

// HTTPConfig is the JSON-RPC/HTTP configuration.
type HTTPConfig struct {
	Modules            []string
	CorsAllowedOrigins []string
	Vhosts             []string
	DenyList           []string
	// SeiLegacyAllowlist is BuildSeiLegacyEnabledSet(app.toml enabled_legacy_sei_apis); nil skips the HTTP gate
	// for gated sei_* and sei2_* methods.
	SeiLegacyAllowlist map[string]struct{}
	prefix             string // path prefix on which to mount http handler
	RPCEndpointConfig
}

// WsConfig is the JSON-RPC/Websocket configuration
type WsConfig struct {
	Origins []string
	Modules []string
	prefix  string // path prefix on which to mount ws handler
	RPCEndpointConfig
}

type RPCEndpointConfig struct {
	JwtSecret              []byte // optional JWT secret
	batchItemLimit         int
	batchResponseSizeLimit int
	readLimit              int64
}

type rpcHandler struct {
	http.Handler
	server *rpc.Server
}

type HTTPServer struct {
	timeouts rpc.HTTPTimeouts
	mux      http.ServeMux // registered handlers go here

	mu       sync.Mutex
	server   *http.Server
	listener net.Listener // non-nil when server is running

	// HTTP RPC handler things.

	HTTPConfig  HTTPConfig
	httpHandler atomic.Value // *rpcHandler

	// WebSocket handler things.
	WsConfig  WsConfig
	wsHandler atomic.Value // *rpcHandler

	// These are set by SetListenAddr.
	endpoint string
	host     string
	port     int

	handlerNames map[string]string
}

const (
	shutdownTimeout        = 5 * time.Second
	metricsPrinterInterval = 5 * time.Second
)

func NewHTTPServer(timeouts rpc.HTTPTimeouts) *HTTPServer { _ = "STUB: not implemented"; return nil }

// SetListenAddr configures the listening address of the server.
// The address can only be set while the server isn't running.
func (h *HTTPServer) SetListenAddr(host string, port int) error {
	_ = "STUB: not implemented"
	return nil
}

// ListenAddr returns the listening address of the server.
func (h *HTTPServer) ListenAddr() string { _ = "STUB: not implemented"; return "" }

// Start starts the HTTP server if it is enabled and not already running.
func (h *HTTPServer) Start() error { _ = "STUB: not implemented"; return nil }

// already running or not configured

// Initialize the server.

// Start the server.

// If the server fails to start, we need to clear out the RPC and WS
// configuration so they can be configured another time.

// if server is websocket only, return after logging

// Log http endpoint.

// Start metrics printer
// Prometheus metrics are always exported; stdout printing requires EVM_DEBUG_METRICS=true

// Log all handlers mounted on server.

func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// check if ws request and serve if ws enabled
	return
}

// if http-rpc is enabled, try to serve request

// First try to route in the mux.
// Requests to a path below root are handled by the mux,
// which has all the handlers registered via Node.RegisterHandler.
// These are made available when RPC is enabled.

// CheckPath checks whether a given request URL matches a given path prefix.
func CheckPath(r *http.Request, path string) bool {
	_ = "STUB: not implemented"
	// if no prefix has been specified, request URL must be on root
	return false
}

// otherwise, check to make sure prefix matches

// Stop shuts down the HTTP server.
func (h *HTTPServer) Stop() { _ = "STUB: not implemented"; return }

func (h *HTTPServer) doStop() { _ = "STUB: not implemented"; return }

// not running

// Stop metrics printer

// Shut down the server.

// Clear out everything to allow re-configuring it later.

// EnableRPC turns on JSON-RPC over HTTP on the server.
func (h *HTTPServer) EnableRPC(apis []rpc.API, config HTTPConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Create RPC server and handler.

// disableRPC stops the HTTP RPC handler. This is internal, the caller must hold h.mu.
func (h *HTTPServer) disableRPC() bool { _ = "STUB: not implemented"; return false }

// EnableWS turns on JSON-RPC over WebSocket on the server.
func (h *HTTPServer) EnableWS(apis []rpc.API, config WsConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Create RPC server and handler.

// stopWS disables JSON-RPC over WebSocket and also stops the server if it only serves WebSocket.
//
//lint:ignore U1000 lifecycle method retained for completeness
func (h *HTTPServer) stopWS() { _ = "STUB: not implemented"; return }

// disableWS disables the WebSocket handler. This is internal, the caller must hold h.mu.
func (h *HTTPServer) disableWS() bool { _ = "STUB: not implemented"; return false }

// rpcAllowed returns true when JSON-RPC over HTTP is enabled.
func (h *HTTPServer) rpcAllowed() bool { _ = "STUB: not implemented"; return false }

// wsAllowed returns true when JSON-RPC over WebSocket is enabled.
func (h *HTTPServer) wsAllowed() bool { _ = "STUB: not implemented"; return false }

// NewHTTPHandlerStack returns wrapped http-related handlers.
func NewHTTPHandlerStack(srv http.Handler, cors []string, vhosts []string, JwtSecret []byte) http.Handler {
	_ = "STUB: not implemented"
	// Wrap the CORS-handler within a host-handler
	return *new(http.Handler)
}

// NewWSHandlerStack returns a wrapped ws-related handler.
func NewWSHandlerStack(srv http.Handler, JwtSecret []byte) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func newCorsHandler(srv http.Handler, allowedOrigins []string) http.Handler {
	_ = "STUB: not implemented"
	// disable CORS support if user has not specified a custom CORS configuration
	return *new(http.Handler)
}

// virtualHostHandler is a handler which validates the Host-header of incoming requests.
// Using virtual hosts can help prevent DNS rebinding attacks, where a 'random' domain name points to
// the service ip address (but without CORS headers). By verifying the targeted virtual host, we can
// ensure that it's a destination that the node operator has defined.
type virtualHostHandler struct {
	vhosts map[string]struct{}
	next   http.Handler
}

func newVHostHandler(vhosts []string, next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// ServeHTTP serves JSON-RPC requests over HTTP, implements http.Handler
func (h *virtualHostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// if r.Host is not set, we can continue serving since a browser would set the Host header
	return
}

// Either invalid (too many colons) or no port specified

// It's an IP address, we can serve that

// Not an IP address, but a hostname. Need to validate

var gzPool = sync.Pool{
	New: func() interface{} {
		w := gzip.NewWriter(io.Discard)
		return w
	},
}

type gzipResponseWriter struct {
	resp http.ResponseWriter

	gz            *gzip.Writer
	contentLength uint64 // total length of the uncompressed response
	written       uint64 // amount of written bytes from the uncompressed response
	hasLength     bool   // true if uncompressed response had Content-Length
	inited        bool   // true after init was called for the first time
}

// init runs just before response headers are written. Among other things, this function
// also decides whether compression will be applied at all.
func (w *gzipResponseWriter) init() {
	if w.inited {
		return
	}
	w.inited = true

	hdr := w.resp.Header()
	length := hdr.Get("content-length")
	if len(length) > 0 {
		if n, err := strconv.ParseUint(length, 10, 64); err == nil {
			w.hasLength = true
			w.contentLength = n
		}
	}

	// Setting Transfer-Encoding to "identity" explicitly disables compression. net/http
	// also recognizes this header value and uses it to disable "chunked" transfer
	// encoding, trimming the header from the response. This means downstream handlers can
	// set this without harm, even if they aren't wrapped by NewGzipHandler.
	//
	// In go-ethereum, we use this signal to disable compression for certain error
	// responses which are flushed out close to the write deadline of the response. For
	// these cases, we want to avoid chunked transfer encoding and compression because
	// they require additional output that may not get written in time.
	passthrough := hdr.Get("transfer-encoding") == "identity"
	if !passthrough {
		w.gz = gzPool.Get().(*gzip.Writer)
		w.gz.Reset(w.resp)
		hdr.Del("content-length")
		hdr.Set("content-encoding", "gzip")
	}
}

func (w *gzipResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (w *gzipResponseWriter) WriteHeader(status int) { _ = "STUB: not implemented"; return }

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Compression is disabled.
		nil
}

//nolint:gosec

// The HTTP handler has finished writing the entire uncompressed response. Close
// the gzip stream to ensure the footer will be seen by the client in case the
// response is flushed after this call to write.

func (w *gzipResponseWriter) Flush() { _ = "STUB: not implemented"; return }

func (w *gzipResponseWriter) close() { _ = "STUB: not implemented"; return }

func NewGzipHandler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// RegisterApis checks the given modules' availability, generates an allowlist based on the allowed modules,
// and then registers all of the APIs exposed by the services.
func RegisterApis(apis []rpc.API, modules []string, srv *rpc.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate the allow list based on the allowed modules

// Register all the APIs exposed by the services
