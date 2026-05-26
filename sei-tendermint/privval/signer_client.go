package privval

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// SignerClient implements PrivValidator.
// Handles remote validator connections that provide signing services
type SignerClient struct {
	endpoint *SignerListenerEndpoint
	chainID  string
}

var _ types.PrivValidator = (*SignerClient)(nil)

// NewSignerClient returns an instance of SignerClient.
// it will start the endpoint (if not already started)
func NewSignerClient(ctx context.Context, endpoint *SignerListenerEndpoint, chainID string) (*SignerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the underlying connection
func (sc *SignerClient) Close() error { _ = "STUB: not implemented"; return nil }

// IsConnected indicates with the signer is connected to a remote signing service
func (sc *SignerClient) IsConnected() bool { _ = "STUB: not implemented"; return false }

// WaitForConnection waits maxWait for a connection or returns a timeout error
func (sc *SignerClient) WaitForConnection(ctx context.Context, maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//--------------------------------------------------------
// Implement PrivValidator

// Ping sends a ping request to the remote signer
func (sc *SignerClient) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// GetPubKey retrieves a public key from a remote signer
// returns an error if client is not able to provide the key
func (sc *SignerClient) GetPubKey(ctx context.Context) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// SignVote requests a remote signer to sign a vote
func (sc *SignerClient) SignVote(ctx context.Context, chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// SignProposal requests a remote signer to sign a proposal
func (sc *SignerClient) SignProposal(ctx context.Context, chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}
