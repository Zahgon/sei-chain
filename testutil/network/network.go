package network

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/network"
)

type (
	Network = network.Network
	Config  = network.Config
)

type TestAppOptions struct{}

func (t TestAppOptions) Get(s string) any { _ = "STUB: not implemented"; return *new(any) }

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func New(t *testing.T, configs ...network.Config) *network.Network {
	_ = "STUB: not implemented"
	return nil
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfig() network.Config { _ = "STUB: not implemented"; return *new(network.Config) }
