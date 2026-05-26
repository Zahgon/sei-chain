package grpc

import (
	context "context"

	privvalproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/privval"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// SignerServer implements PrivValidatorAPIServer 9generated via protobuf services)
// Handles remote validator connections that provide signing services
type SignerServer struct {
	chainID string
	privVal types.PrivValidator
}

func NewSignerServer(chainID string, privVal types.PrivValidator) *SignerServer {
	_ = "STUB: not implemented"
	return nil
}

var _ privvalproto.PrivValidatorAPIServer = (*SignerServer)(nil)

// PubKey receives a request for the pubkey
// returns the pubkey on success and error on failure
func (ss *SignerServer) GetPubKey(ctx context.Context, req *privvalproto.PubKeyRequest) (
	*privvalproto.PubKeyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignVote receives a vote sign requests, attempts to sign it
// returns SignedVoteResponse on success and error on failure
func (ss *SignerServer) SignVote(ctx context.Context, req *privvalproto.SignVoteRequest) (*privvalproto.SignedVoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignProposal receives a proposal sign requests, attempts to sign it
// returns SignedProposalResponse on success and error on failure
func (ss *SignerServer) SignProposal(ctx context.Context, req *privvalproto.SignProposalRequest) (*privvalproto.SignedProposalResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
