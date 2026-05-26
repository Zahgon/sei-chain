package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/types"
)

// WebSocket handler

const (
	defaultWSWriteChanCapacity = 100
	defaultWSWriteWait         = 10 * time.Second
	defaultWSReadWait          = 30 * time.Second
	defaultWSPingPeriod        = (defaultWSReadWait * 9) / 10
)

// WebsocketManager provides a WS handler for incoming connections and passes a
// map of functions along with any additional params to new connections.
// NOTE: The websocket path is defined externally, e.g. in node/node.go
type WebsocketManager struct {
	websocket.Upgrader

	funcMap       map[string]*RPCFunc
	wsConnOptions []func(*wsConnection)
}

// NewWebsocketManager returns a new WebsocketManager that passes a map of
// functions, connection options and logger to new WS connections.
func NewWebsocketManager(funcMap map[string]*RPCFunc, wsConnOptions ...func(*wsConnection)) *WebsocketManager {
	_ = "STUB: not implemented"
	return nil
}

// TODO ???
//
// The default behavior would be relevant to browser-based clients,
// afaik. I suppose having a pass-through is a workaround for allowing
// for more complex security schemes, shifting the burden of
// AuthN/AuthZ outside the Tendermint RPC.
// I can't think of other uses right now that would warrant a TODO
// though. The real backstory of this TODO shall remain shrouded in
// mystery

// WebsocketHandler upgrades the request/response (via http.Hijack) and starts
// the wsConnection.
func (wm *WebsocketManager) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// The upgrader has already reported an HTTP error to the client, so we
// need only log it.

// register connection

// starting the conn is blocking

// WebSocket connection

// A single websocket connection contains listener id, underlying ws
// connection, and the event switch for subscribing to events.
//
// In case of an error, the connection is stopped.
type wsConnection struct {
	remoteAddr string
	baseConn   *websocket.Conn
	// writeChan is never closed, to allow WriteRPCResponse() to fail.
	writeChan chan rpctypes.RPCResponse

	// chan, which is closed when/if readRoutine errors
	// used to abort writeRoutine
	readRoutineQuit chan struct{}

	funcMap map[string]*RPCFunc

	// Connection times out if we haven't received *anything* in this long, not even pings.
	readWait time.Duration

	// Send pings to server with this period. Must be less than readWait, but greater than zero.
	pingPeriod time.Duration

	// Maximum message size.
	readLimit int64

	// callback which is called upon disconnect
	onDisconnect func(remoteAddr string)

	ctx    context.Context
	cancel context.CancelFunc
}

// NewWSConnection wraps websocket.Conn.
//
// See the commentary on the func(*wsConnection) functions for a detailed
// description of how to configure ping period and pong wait time. NOTE: if the
// write buffer is full, pongs may be dropped, which may cause clients to
// disconnect. see https://github.com/gorilla/websocket/issues/97
func newWSConnection(baseConn *websocket.Conn, funcMap map[string]*RPCFunc, options ...func(*wsConnection)) *wsConnection {
	_ = "STUB: not implemented"
	return nil
}

// OnDisconnect sets a callback which is used upon disconnect - not
// Goroutine-safe. Nop by default.
func OnDisconnect(onDisconnect func(remoteAddr string)) func(*wsConnection) {
	_ = "STUB: not implemented"
	return nil
}

// ReadWait sets the amount of time to wait before a websocket read times out.
// It should only be used in the constructor - not Goroutine-safe.
func ReadWait(readWait time.Duration) func(*wsConnection) { _ = "STUB: not implemented"; return nil }

// PingPeriod sets the duration for sending websocket pings.
// It should only be used in the constructor - not Goroutine-safe.
func PingPeriod(pingPeriod time.Duration) func(*wsConnection) {
	_ = "STUB: not implemented"
	return nil
}

// ReadLimit sets the maximum size for reading message.
// It should only be used in the constructor - not Goroutine-safe.
func ReadLimit(readLimit int64) func(*wsConnection) { _ = "STUB: not implemented"; return nil }

// Start starts the client service routines and blocks until there is an error.
func (wsc *wsConnection) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Read subscriptions/unsubscriptions to events

// Write responses, BLOCKING.

// Stop unsubscribes the remote from all subscriptions.
func (wsc *wsConnection) Stop() error { _ = "STUB: not implemented"; return nil }

// GetRemoteAddr returns the remote address of the underlying connection.
// It implements WSRPCConnection
func (wsc *wsConnection) GetRemoteAddr() string { _ = "STUB: not implemented"; return "" }

// WriteRPCResponse pushes a response to the writeChan, and blocks until it is
// accepted.
// It implements WSRPCConnection. It is Goroutine-safe.
func (wsc *wsConnection) WriteRPCResponse(ctx context.Context, resp rpctypes.RPCResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// TryWriteRPCResponse attempts to push a response to the writeChan, but does
// not block.
// It implements WSRPCConnection. It is Goroutine-safe
func (wsc *wsConnection) TryWriteRPCResponse(ctx context.Context, resp rpctypes.RPCResponse) bool {
	_ = "STUB: not implemented"
	return false
}

// Context returns the connection's context.
// The context is canceled when the client's connection closes.
func (wsc *wsConnection) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Read from the socket and subscribe to or unsubscribe from events
func (wsc *wsConnection) readRoutine(ctx context.Context) {
	_ = "STUB: not implemented"
	// readRoutine will block until response is written or WS connection is closed
	return
}

// reset deadline for every type of message (control or data)

// A Notification is a Request object without an "id" member.
// The Server MUST NOT reply to a Notification, including those that are within a batch request.

// Now, fetch the RPCFunc and execute it.

// receives on a write channel and writes out on the socket
func (wsc *wsConnection) writeRoutine(ctx context.Context) { _ = "STUB: not implemented"; return }

// https://github.com/gorilla/websocket/issues/97

// error in readRoutine

// All writes to the websocket must (re)set the write deadline.
// If some writes don't set it while others do, they may timeout incorrectly
// (https://github.com/tendermint/tendermint/issues/553)
func (wsc *wsConnection) writeMessageWithDeadline(msgType int, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}
