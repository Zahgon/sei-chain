package privval

import (
	"context"
	"time"
)

const (
	defaultMaxDialRetries        = 10
	defaultRetryWaitMilliseconds = 100
)

// SignerServiceEndpointOption sets an optional parameter on the SignerDialerEndpoint.
type SignerServiceEndpointOption func(*SignerDialerEndpoint)

// SignerDialerEndpointTimeoutReadWrite sets the read and write timeout for
// connections from client processes.
func SignerDialerEndpointTimeoutReadWrite(timeout time.Duration) SignerServiceEndpointOption {
	_ = "STUB: not implemented"
	return *new(SignerServiceEndpointOption)
}

// SignerDialerEndpointConnRetries sets the amount of attempted retries to
// acceptNewConnection.
func SignerDialerEndpointConnRetries(retries int) SignerServiceEndpointOption {
	_ = "STUB: not implemented"
	return *new(SignerServiceEndpointOption)
}

// SignerDialerEndpointRetryWaitInterval sets the retry wait interval to a
// custom value.
func SignerDialerEndpointRetryWaitInterval(interval time.Duration) SignerServiceEndpointOption {
	_ = "STUB: not implemented"
	return *new(SignerServiceEndpointOption)
}

// SignerDialerEndpoint dials using its dialer and responds to any signature
// requests using its privVal.
type SignerDialerEndpoint struct {
	signerEndpoint

	dialer SocketDialer

	retryWait      time.Duration
	maxConnRetries int
}

// NewSignerDialerEndpoint returns a SignerDialerEndpoint that will dial using the given
// dialer and respond to any signature requests over the connection
// using the given privVal.
func NewSignerDialerEndpoint(
	dialer SocketDialer,
	options ...SignerServiceEndpointOption,
) *SignerDialerEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (sd *SignerDialerEndpoint) OnStart(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
func (sd *SignerDialerEndpoint) OnStop() { _ = "STUB: not implemented"; return }

func (sd *SignerDialerEndpoint) ensureConnection(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait between retries
