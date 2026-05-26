// nolint: gosec
package main

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/config"
	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

const (
	PrivvalAddressTCP     = "tcp://0.0.0.0:27559"
	PrivvalAddressGRPC    = "grpc://0.0.0.0:27559"
	PrivvalAddressUNIX    = "unix:///var/run/privval.sock"
	PrivvalKeyFile        = "config/priv_validator_key.json"
	PrivvalStateFile      = "data/priv_validator_state.json"
	PrivvalDummyKeyFile   = "config/dummy_validator_key.json"
	PrivvalDummyStateFile = "data/dummy_validator_state.json"
)

// Setup sets up the testnet configuration.
func Setup(testnet *e2e.Testnet) error { _ = "STUB: not implemented"; return nil }

// light clients don't need an app directory

// stop early if a light client

// Set up a dummy validator. Tendermint requires a file PV even when not used, so we
// give it a dummy such that it will fail if it actually tries to use it.

// MakeDockerCompose generates a Docker Compose config for a testnet.
func MakeDockerCompose(testnet *e2e.Testnet) ([]byte, error) {
	_ = "STUB: not implemented"
	// Must use version 2 Docker Compose format, to support IPv6.
	return nil, nil
}

// MakeGenesis generates a genesis document.
func MakeGenesis(testnet *e2e.Testnet) (types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return *new(types.GenesisDoc), nil
}

// The validator set will be sorted internally by Tendermint ranked by power,
// but we sort it here as well so that all genesis files are identical.

// MakeConfig generates a Tendermint config for a node.
func MakeConfig(node *e2e.Node) (*config.Config, error) { _ = "STUB: not implemented"; return nil, nil }

// Tendermint errors if it does not have a privval key set up, regardless of whether
// it's actually needed (e.g. for remote KMS or non-validators). We set up a dummy
// key here by default, and use the real key for actual validators that should use
// the file privval.

// Don't need to do anything, since we're using a dummy privval key by default.

// MakeAppConfig generates an ABCI application config for a node.
func MakeAppConfig(node *e2e.Node) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UpdateConfigStateSync updates the state sync config for a node.
func UpdateConfigStateSync(node *e2e.Node, height int64, hash []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME Apparently there's no function to simply load a config file without
// involving the entire Viper apparatus, so we'll just resort to regexps.
