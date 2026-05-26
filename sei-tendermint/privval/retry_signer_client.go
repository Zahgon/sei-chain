package privval

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// RetrySignerClient wraps SignerClient adding retry for each operation (except
// Ping) w/ a timeout.
type RetrySignerClient struct {
	next    *SignerClient
	retries int
	timeout time.Duration
}

// NewRetrySignerClient returns RetrySignerClient. If +retries+ is 0, the
// client will be retrying each operation indefinitely.
func NewRetrySignerClient(sc *SignerClient, retries int, timeout time.Duration) *RetrySignerClient {
	_ = "STUB: not implemented"
	return nil
}

var _ types.PrivValidator = (*RetrySignerClient)(nil)

func (sc *RetrySignerClient) Close() error { _ = "STUB: not implemented"; return nil }

func (sc *RetrySignerClient) IsConnected() bool { _ = "STUB: not implemented"; return false }

func (sc *RetrySignerClient) WaitForConnection(ctx context.Context, maxWait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//--------------------------------------------------------
// Implement PrivValidator

func (sc *RetrySignerClient) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (sc *RetrySignerClient) GetPubKey(ctx context.Context) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// If remote signer errors, we don't retry.

func (sc *RetrySignerClient) SignVote(ctx context.Context, chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// If remote signer errors, we don't retry.

func (sc *RetrySignerClient) SignProposal(ctx context.Context, chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

// If remote signer errors, we don't retry.
