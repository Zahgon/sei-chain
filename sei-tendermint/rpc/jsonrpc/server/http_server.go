// Commons for HTTP handling
package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/sei-protocol/seilog"

	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/types"
)

var logger = seilog.NewLogger("tendermint", "rpc", "jsonrpc", "server")

// Config is a RPC server configuration.
type Config struct {
	// The maximum number of connections that will be accepted by the listener.
	// See https://godoc.org/golang.org/x/net/netutil#LimitListener
	MaxOpenConnections int

	// Used to set the HTTP server's per-request read timeout.
	// See https://godoc.org/net/http#Server.ReadTimeout
	ReadTimeout time.Duration

	// Used to set the HTTP server's per-request write timeout.  Note that this
	// affects ALL methods on the server, so it should not be set too low. This
	// should be used as a safety valve, not a resource-control timeout.
	//
	// See https://godoc.org/net/http#Server.WriteTimeout
	WriteTimeout time.Duration

	// Controls the maximum number of bytes the server will read parsing the
	// request body.
	MaxBodyBytes int64

	// Controls the maximum size of a request header.
	// See https://godoc.org/net/http#Server.MaxHeaderBytes
	MaxHeaderBytes int
}

// DefaultConfig returns a default configuration.
func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// unlimited

// no default timeout
// 1MB
// same as the net/http default

// Serve creates a http.Server and calls Serve with the given listener. It
// wraps handler to recover panics and limit the request body size.
func Serve(ctx context.Context, listener net.Listener, handler http.Handler, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

// Serve creates a http.Server and calls ServeTLS with the given listener,
// certFile and keyFile. It wraps handler to recover panics and limit the
// request body size.
func ServeTLS(ctx context.Context, listener net.Listener, handler http.Handler, certFile, keyFile string, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

// writeError writes an internal server error (500) to w with the text
// of err in the body. This is a fallback used when a handler is unable to
// write the expected response.
func writeError(w http.ResponseWriter, statusCode int, err error) {
	_ = "STUB: not implemented"
	return
}

// writeHTTPResponse writes a JSON-RPC response to w. If rsp encodes an error,
// the response body is its error object; otherwise its responses is the result.
//
// Unless there is an error encoding the response, the status is 200 OK.
func writeHTTPResponse(w http.ResponseWriter, rsp rpctypes.RPCResponse) {
	_ = "STUB: not implemented"
	return
}

// If there's any error for lag is high, override the status code and response body

// writeRPCResponse writes one or more JSON-RPC responses to w. A single
// response is encoded as an object, otherwise the response is sent as a batch
// (array) of response objects.
//
// Unless there is an error encoding the responses, the status is 200 OK.
func writeRPCResponse(w http.ResponseWriter, rsps ...rpctypes.RPCResponse) {
	_ = "STUB: not implemented"
	return
}

// If there's any error for lag is high, override the status code

//-----------------------------------------------------------------------------

// recoverAndLogHandler wraps an HTTP handler, adding error logging.  If the
// inner handler panics, the wrapper recovers, logs, sends an HTTP 500 error
// response to the client.
func recoverAndLogHandler(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Capture the HTTP status written by the handler.

// Recover panics from inside handler and try to send the client
// 500 Internal server error. If the handler panicked after already
// sending a (partial) response, this is a no-op.

// Log timing and response information from the handler.

// MaxBytesHandler wraps h in a handler that limits the size of the request
// body to at most maxBytes. If maxBytes <= 0, the request body is not limited.
func MaxBytesHandler(h http.Handler, maxBytes int64) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type maxBytesHandler struct {
	handler  http.Handler
	maxBytes int64
}

func (h maxBytesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// newStatusWriter wraps an http.ResponseWriter to capture the HTTP status code
// in *code.
func newStatusWriter(w http.ResponseWriter, code *int) statusWriter {
	_ = "STUB: not implemented"
	return *new(statusWriter)
}

type statusWriter struct {
	http.ResponseWriter
	http.Hijacker // to support websocket upgrade

	code *int
}

// WriteHeader implements part of http.ResponseWriter. It delegates to the
// wrapped writer, and as a side effect captures the written code.
//
// Note that if a request does not explicitly call WriteHeader, the code will
// not be updated.
func (w statusWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

// Listen starts a new net.Listener on the given address.
// It returns an error if the address is invalid or the call to Listen() fails.
func Listen(addr string, maxOpenConnections int) (listener net.Listener, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}
