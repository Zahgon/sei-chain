package light

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/light/provider"
	"github.com/sei-protocol/sei-chain/sei-tendermint/light/store"
)

// NewHTTPClient initiates an instance of a light client using HTTP addresses
// for both the primary provider and witnesses of the light client. A trusted
// header and hash must be passed to initialize the client.
//
// See all Option(s) for the additional configuration.
// See NewClient.
func NewHTTPClient(
	ctx context.Context,
	chainID string,
	trustOptions TrustOptions,
	primaryAddress string,
	witnessesAddresses []string,
	trustedStore store.Store,
	blacklistTTL time.Duration,
	options ...Option) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func providersFromAddresses(addrs []string, chainID string) ([]provider.Provider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
