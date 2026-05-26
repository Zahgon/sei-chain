package privval

// IsConnTimeout returns a boolean indicating whether the error is known to
// report that a connection timeout occurred. This detects both fundamental
// network timeouts, as well as ErrConnTimeout errors.
func IsConnTimeout(err error) bool { _ = "STUB: not implemented"; return false }

// NewSignerListener creates a new SignerListenerEndpoint using the corresponding listen address
func NewSignerListener(listenAddr string) (*SignerListenerEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:goconst

// TODO: persist this key so external signer can actually authenticate us

// semantically unreachable
