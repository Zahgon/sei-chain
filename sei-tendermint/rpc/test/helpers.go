package rpctest

import (
	"context"
	"testing"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
)

// Options helps with specifying some parameters for our RPC testing for greater
// control.
type Options struct {
	suppressStdout bool
}

// waitForRPC connects to the RPC service and blocks until a /status call succeeds.
func waitForRPC(ctx context.Context, conf *config.Config) { _ = "STUB: not implemented"; return }

func randPort() int { _ = "STUB: not implemented"; return 0 }

// makeAddrs constructs local listener addresses for node services.  This
// implementation uses random ports so test instances can run concurrently.
func makeAddrs() (p2pAddr, rpcAddr string) { _ = "STUB: not implemented"; return "", "" }

func CreateConfig(t *testing.T, testName string) (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ServiceCloser func(context.Context) error

func StartTendermint(
	ctx context.Context,
	conf *config.Config,
	app abci.Application,
	opts ...func(*Options),
) (service.Service, ServiceCloser, error) {
	_ = "STUB: not implemented"
	return *new(service.Service), *new(ServiceCloser), nil
}

// SuppressStdout is an option that tries to make sure the RPC test Tendermint
// node doesn't log anything to stdout.
func SuppressStdout(o *Options) { _ = "STUB: not implemented"; return }
