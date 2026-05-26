package client

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sei-protocol/seilog"

	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/types"
)

var logger = seilog.NewLogger("tendermint", "rpc", "jsonrpc", "client")

// wsOptions carries optional settings for a websocket connection.
type wsOptions struct {
	MaxReconnectAttempts uint          // maximum attempts to reconnect
	ReadWait             time.Duration // deadline for any read op
	WriteWait            time.Duration // deadline for any write op
	PingPeriod           time.Duration // frequency with which pings are sent
}

// defaultWSOptions are the default websocket connection settings.
var defaultWSOptions = wsOptions{
	MaxReconnectAttempts: 10, // first: 2 sec, last: 17 min.
	WriteWait:            10 * time.Second,
	ReadWait:             0,
	PingPeriod:           0,
}

// WSClient is a JSON-RPC client, which uses WebSocket for communication with
// the remote server.
//
// WSClient is safe for concurrent use by multiple goroutines.
type WSClient struct {
	conn *websocket.Conn

	Address  string // IP:PORT or /path/to/socket
	Endpoint string // /websocket/url/endpoint
	Dialer   func(string, string) (net.Conn, error)

	// Single user facing channel to read RPCResponses from, closed only when the
	// client is being stopped.
	ResponsesCh chan rpctypes.RPCResponse

	// Callback, which will be called each time after successful reconnect.
	onReconnect func()

	// internal channels
	send            chan rpctypes.RPCRequest // user requests
	backlog         chan rpctypes.RPCRequest // stores a single user request received during a conn failure
	reconnectAfter  chan error               // reconnect requests
	readRoutineQuit chan struct{}            // a way for readRoutine to close writeRoutine

	// Maximum reconnect attempts (0 or greater; default: 25).
	maxReconnectAttempts uint

	// Support both ws and wss protocols
	protocol string

	wg sync.WaitGroup

	mtx          sync.RWMutex
	reconnecting bool
	nextReqID    int
	// sentIDs        map[types.JSONRPCIntID]bool // IDs of the requests currently in flight

	// Time allowed to write a message to the server. 0 means block until operation succeeds.
	writeWait time.Duration

	// Time allowed to read the next message from the server. 0 means block until operation succeeds.
	readWait time.Duration

	// Send pings to server with this period. Must be less than readWait. If 0, no pings will be sent.
	pingPeriod time.Duration
}

// NewWS returns a new client with default options. The endpoint argument must
// begin with a `/`. An error is returned on invalid remote.
func NewWS(remoteAddr, endpoint string) (*WSClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default to ws protocol, unless wss or https is specified

// sentIDs: make(map[types.JSONRPCIntID]bool),

// OnReconnect sets the callback, which will be called every time after
// successful reconnect.
// Could only be set before Start.
func (c *WSClient) OnReconnect(cb func()) {
	_ = "STUB: not implemented"

	// String returns WS client full address.
	return
}

func (c *WSClient) String() string { _ = "STUB: not implemented"; return "" }

// Start dials the specified service address and starts the I/O routines.  The
// service routines run until ctx terminates. To wait for the client to exit
// after ctx ends, call Stop.
func (c *WSClient) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// 1 additional error may come from the read/write
// goroutine depending on which failed first.

// capacity for 1 request. a user won't be able to send more because the send
// channel is unbuffered.

// Stop blocks until the client is shut down and returns nil.
//
// TODO(creachadair): This method exists for compatibility with the original
// service plumbing. Give it a better name (e.g., Wait).
func (c *WSClient) Stop() error {
	_ = "STUB: not implemented"
	// only close user-facing channels when we can't write to them
	return nil
}

// IsReconnecting returns true if the client is reconnecting right now.
func (c *WSClient) IsReconnecting() bool { _ = "STUB: not implemented"; return false }

// Send the given RPC request to the server. Results will be available on
// ResponsesCh, errors, if any, on ErrorsCh. Will block until send succeeds or
// ctx.Done is closed.
func (c *WSClient) Send(ctx context.Context, request rpctypes.RPCRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// c.mtx.Lock()
// c.sentIDs[request.ID.(types.JSONRPCIntID)] = true
// c.mtx.Unlock()

// Call enqueues a call request onto the Send queue. Requests are JSON encoded.
func (c *WSClient) Call(ctx context.Context, method string, params map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Private methods

func (c *WSClient) nextRequestID() int { _ = "STUB: not implemented"; return 0 }

func (c *WSClient) dial() error { _ = "STUB: not implemented"; return nil }

// nolint:bodyclose

// reconnect tries to redial up to maxReconnectAttempts with exponential
// backoff.
func (c *WSClient) reconnect(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// nolint:gosec // G404: Use of weak random number generator
// 1s == (1e9 ns)

func (c *WSClient) startReadWriteRoutines(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *WSClient) processBacklog() error { _ = "STUB: not implemented"; return nil }

// requeue request

func (c *WSClient) reconnectRoutine(ctx context.Context) { _ = "STUB: not implemented"; return }

// wait until writeRoutine and readRoutine finish

// drain reconnectAfter

// The client ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *WSClient) writeRoutine(ctx context.Context) { _ = "STUB: not implemented"; return }

// ticker with a predefined period

// ticker that never fires

// add request to the backlog, so we don't lose it

// The client ensures that there is at most one reader to a connection by
// executing all reads from this goroutine.
func (c *WSClient) readRoutine(ctx context.Context) { _ = "STUB: not implemented"; return }

// reset deadline for every message type (control or data)

// TODO: events resulting from /subscribe do not work with ->
// because they are implemented as responses with the subscribe request's
// ID. According to the spec, they should be notifications (requests
// without IDs).
// https://github.com/tendermint/tendermint/issues/2949
//
// Combine a non-blocking read on BaseService.Quit with a non-blocking write on ResponsesCh to avoid blocking
// c.wg.Wait() in c.Stop(). Note we rely on Quit being closed so that it sends unlimited Quit signals to stop
// both readRoutine and writeRoutine

// Predefined methods

// Subscribe to a query. Note the server must have a "subscribe" route
// defined.
func (c *WSClient) Subscribe(ctx context.Context, query string) error {
	_ = "STUB: not implemented"
	return nil
}

// Unsubscribe from a query. Note the server must have a "unsubscribe" route
// defined.
func (c *WSClient) Unsubscribe(ctx context.Context, query string) error {
	_ = "STUB: not implemented"
	return nil
}

// UnsubscribeAll from all. Note the server must have a "unsubscribe_all" route
// defined.
func (c *WSClient) UnsubscribeAll(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
