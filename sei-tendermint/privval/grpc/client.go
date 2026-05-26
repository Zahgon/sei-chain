package grpc

import (
	"context"

	grpc "google.golang.org/grpc"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	privvalproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/privval"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// SignerClient implements PrivValidator.
// Handles remote validator connections that provide signing services
type SignerClient struct {
	client  privvalproto.PrivValidatorAPIClient
	conn    *grpc.ClientConn
	chainID string
}

var _ types.PrivValidator = (*SignerClient)(nil)

// NewSignerClient returns an instance of SignerClient.
// it will start the endpoint (if not already started)
func NewSignerClient(conn *grpc.ClientConn,
	chainID string) (*SignerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the Private Validator Client

// Close closes the underlying connection
func (sc *SignerClient) Close() error { _ = "STUB: not implemented"; return nil }

//--------------------------------------------------------
// Implement PrivValidator

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
