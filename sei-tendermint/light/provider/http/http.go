package http

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/light/provider"
	rpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
	rpctypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/jsonrpc/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var defaultOptions = Options{
	MaxRetryAttempts:    5,
	Timeout:             5 * time.Second,
	NoBlockThreshold:    5,
	NoResponseThreshold: 5,
}

// http provider uses an RPC client to obtain the necessary information.
type http struct {
	chainID string
	client  rpcclient.RemoteClient

	// httt provider heuristics

	// The provider tracks the amount of times that the
	// client doesn't respond. If this exceeds the threshold
	// then the provider will return an unreliable provider error
	noResponseThreshold uint16
	noResponseCount     uint16

	// The provider tracks the amount of time the client
	// doesn't have a block. If this exceeds the threshold
	// then the provider will return an unreliable provider error
	noBlockThreshold uint16
	noBlockCount     uint16

	// In a single request, the provider attempts multiple times
	// with exponential backoff to reach the client. If this
	// exceeds the maxRetry attempts, this result in a ErrNoResponse
	maxRetryAttempts uint16
}

type Options struct {
	// 0 means no retries
	MaxRetryAttempts uint16
	// 0 means no timeout.
	Timeout time.Duration
	// The amount of requests that a client doesn't have the block
	// for before the provider deems the client unreliable
	NoBlockThreshold uint16
	// The amount of requests that a client doesn't respond to
	// before the provider deems the client unreliable
	NoResponseThreshold uint16
}

// New creates a HTTP provider, which is using the rpchttp.HTTP client under
// the hood. If no scheme is provided in the remote URL, http will be used by
// default. The 5s timeout is used for all requests.
func New(chainID, remote string) (provider.Provider, error) {
	_ = "STUB: not implemented"
	return *new(provider.Provider), nil
}

// NewWithOptions is an extension to creating a new http provider that allows the addition
// of a specified timeout and maxRetryAttempts
func NewWithOptions(chainID, remote string, options Options) (provider.Provider, error) {
	_ = "STUB: not implemented"
	// Ensure URL scheme is set (default HTTP) when not provided.
	return *new(provider.Provider), nil
}

func NewWithClient(chainID string, client rpcclient.RemoteClient) provider.Provider {
	_ = "STUB: not implemented"
	return *new(provider.Provider)
}

// NewWithClient allows you to provide a custom client.
func NewWithClientAndOptions(chainID string, client rpcclient.RemoteClient, options Options) provider.Provider {
	_ = "STUB: not implemented"
	return *new(provider.Provider)
}

// Identifies the provider with an IP in string format
func (p *http) ID() string { _ = "STUB: not implemented"; return "" }

// LightBlock fetches a LightBlock at the given height and checks the
// chainID matches.
func (p *http) LightBlock(ctx context.Context, height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReportEvidence calls `/broadcast_evidence` endpoint.
func (p *http) ReportEvidence(ctx context.Context, ev types.Evidence) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *http) validatorSet(ctx context.Context, height *int64) (*types.ValidatorSet, error) {
	_ = "STUB: not implemented"
	// Since the malicious node could report a massive number of pages, making us
	// spend a considerable time iterating, we restrict the number of pages here.
	// => 10000 validators max
	return nil, nil
}

// create another for loop to control retries. If p.maxRetryAttempts
// is negative we will keep repeating.

// if we have exceeded retry attempts then return a no response error

// request timed out: we wait and try again with exponential backoff

// process the rpc error and return the corresponding error to the light client

// check if the error stems from the context

// If we don't know the error then by default we return an unreliable provider error and
// terminate the connection with the peer.

// update the total and increment the page index so we can fetch the
// next page of validators if need be

func (p *http) signedHeader(ctx context.Context, height *int64) (*types.SignedHeader, error) {
	_ = "STUB: not implemented"
	// create a for loop to control retries. If p.maxRetryAttempts
	// is negative we will keep repeating.
	return nil, nil
}

// success!!

// check if the request timed out

// we wait and try again with exponential backoff

// check if the connection was refused or dropped

// else, as a catch all, we return the error as a bad light block response

// process the rpc error and return the corresponding error to the light client

// check if the error stems from the context

// If we don't know the error then by default we return an unreliable provider error and
// terminate the connection with the peer.

func (p *http) noResponse() error { _ = "STUB: not implemented"; return nil }

func (p *http) noBlock(e error) error { _ = "STUB: not implemented"; return nil }

// parseRPCError process the error and return the corresponding error to the light clent
// NOTE: When an error is sent over the wire it gets "flattened" hence we are unable to use error
// checking functions like errors.Is() to unwrap the error.
func (p *http) parseRPCError(e *rpctypes.RPCError) error {
	_ = "STUB: not implemented"

	// 1) check if the error indicates that the peer doesn't have the block
	return nil
}

// 2) check if the height requested is too high

// 3) check if the provider closed the connection

// 4) else return a generic error

func validateHeight(height int64) (*int64, error) { _ = "STUB: not implemented"; return nil, nil }

// exponential backoff (with jitter)
// 0.5s -> 2s -> 4.5s -> 8s -> 12.5 with 1s variation
func backoffTimeout(attempt uint16) time.Duration {
	_ = "STUB: not implemented"
	// nolint:gosec // G404: Use of weak random number generator
	return *new(time.Duration)
}
