package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
)

var logger = seilog.NewLogger("tendermint", "test", "e2e", "node")

const builtinProtocol = "builtin"

// main is the binary entrypoint.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <configfile>", os.Args[0])
		return
	}
	configFile := ""
	if len(os.Args) == 2 {
		configFile = os.Args[1]
	}

	if err := run(ctx, configFile); err != nil {
		os.Exit(1)
	}
	<-ctx.Done()
	logger.Info("Shutting down...")
}

// run runs the application - basically like main() with error handling.
func run(ctx context.Context, configFile string) error { _ = "STUB: not implemented"; return nil }

// Start remote signer (must start before node if running builtin).

// Start app server.

// startNode starts a Tendermint node running the application directly. It assumes the Tendermint
// configuration is in $TMHOME/config/tendermint.toml.
//
// FIXME There is no way to simply load the configuration from a file, so we need to pull in Viper.
func startNode(ctx context.Context, cfg *Config) error { _ = "STUB: not implemented"; return nil }

func startSeedNode(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func startLightNode(ctx context.Context, cfg *Config) error { _ = "STUB: not implemented"; return nil }

// If necessary adjust global WriteTimeout to ensure it's greater than
// TimeoutBroadcastTxCommit.
// See https://github.com/tendermint/tendermint/issues/3435
// Note we don't need to adjust anything if the timeout is already unlimited.

// Error starting or closing listener:

// startSigner starts a signer server connecting to the given endpoint.
func startSigner(ctx context.Context, cfg *Config) error { _ = "STUB: not implemented"; return nil }

// no need to clean up since we remove docker containers

func setupNode() (*config.Config, error) { _ = "STUB: not implemented"; return nil, nil }

// rpcEndpoints takes a list of persistent peers and splits them into a list of rpc endpoints
// using 26657 as the port number
func rpcEndpoints(peers string) []string { _ = "STUB: not implemented"; return nil }

// use RPC port instead

// for ipv6 addresses

// for ipv4 addresses
