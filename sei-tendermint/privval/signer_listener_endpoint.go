package privval

import (
	"context"
	"net"
	"sync"
	"time"

	privvalproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/privval"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("tendermint", "privval")

// SignerListenerEndpointOption sets an optional parameter on the SignerListenerEndpoint.
type SignerListenerEndpointOption func(*SignerListenerEndpoint)

// SignerListenerEndpointTimeoutReadWrite sets the read and write timeout for
// connections from external signing processes.
//
// Default: 5s
func SignerListenerEndpointTimeoutReadWrite(timeout time.Duration) SignerListenerEndpointOption {
	_ = "STUB: not implemented"
	return *new(SignerListenerEndpointOption)
}

// SignerListenerEndpoint listens for an external process to dial in and keeps
// the connection alive by dropping and reconnecting.
//
// The process will send pings every ~3s (read/write timeout * 2/3) to keep the
// connection alive.
type SignerListenerEndpoint struct {
	signerEndpoint

	listener              net.Listener
	connectRequestCh      chan struct{}
	connectionAvailableCh chan net.Conn

	timeoutAccept time.Duration
	pingTimer     *time.Ticker
	pingInterval  time.Duration

	instanceMtx sync.Mutex // Ensures instance public methods access, i.e. SendRequest
}

// NewSignerListenerEndpoint returns an instance of SignerListenerEndpoint.
func NewSignerListenerEndpoint(
	listener net.Listener,
	options ...SignerListenerEndpointOption,
) *SignerListenerEndpoint {
	_ = "STUB: not implemented"
	return nil
}

// OnStart implements service.Service.
func (sl *SignerListenerEndpoint) OnStart(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: ping timeout must be less than read/write timeout

// OnStop implements service.Service
func (sl *SignerListenerEndpoint) OnStop() { _ = "STUB: not implemented"; return }

// Stop listening

// WaitForConnection waits maxWait for a connection or returns a timeout error
func (sl *SignerListenerEndpoint) WaitForConnection(ctx context.Context, maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// SendRequest ensures there is a connection, sends a request and waits for a response
func (sl *SignerListenerEndpoint) SendRequest(ctx context.Context, request privvalproto.Message) (*privvalproto.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset pingTimer to avoid sending unnecessary pings.

func (sl *SignerListenerEndpoint) ensureConnection(ctx context.Context, maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Is there a connection ready? then use it

// block until connected or timeout

func (sl *SignerListenerEndpoint) acceptNewConnection() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// wait for a new conn

func (sl *SignerListenerEndpoint) triggerConnect() { _ = "STUB: not implemented"; return }

func (sl *SignerListenerEndpoint) triggerReconnect() { _ = "STUB: not implemented"; return }

func (sl *SignerListenerEndpoint) serviceLoop(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// We have a good connection, wait for someone that needs one otherwise cancellation

func (sl *SignerListenerEndpoint) pingLoop(ctx context.Context) { _ = "STUB: not implemented"; return }
