package debug

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	rpchttp "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client/http"
)

// dumpStatus gets node status state dump from the Tendermint RPC and writes it
// to file. It returns an error upon failure.
func dumpStatus(ctx context.Context, rpc *rpchttp.HTTP, dir, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// dumpNetInfo gets network information state dump from the Tendermint RPC and
// writes it to file. It returns an error upon failure.
func dumpNetInfo(ctx context.Context, rpc *rpchttp.HTTP, dir, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// dumpConsensusState gets consensus state dump from the Tendermint RPC and
// writes it to file. It returns an error upon failure.
func dumpConsensusState(ctx context.Context, rpc *rpchttp.HTTP, dir, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// copyWAL copies the Tendermint node's WAL file. It returns an error if the
// WAL file cannot be read or copied.
func copyWAL(conf *config.Config, dir string) error { _ = "STUB: not implemented"; return nil }

// copyConfig copies the Tendermint node's config file. It returns an error if
// the config file cannot be read or copied.
func copyConfig(home, dir string) error { _ = "STUB: not implemented"; return nil }

func dumpProfile(dir, addr, profile string, debug int) error { _ = "STUB: not implemented"; return nil }

// nolint: gosec
