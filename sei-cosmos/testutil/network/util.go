package network

import (
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
)

func startInProcess(cfg Config, val *Validator) error { _ = "STUB: not implemented"; return nil }

// We'll need a RPC client if the validator exposes a gRPC or REST endpoint.

// assume server started successfully

func collectGenFiles(cfg Config, vals []*Validator, outputDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// overwrite each validator's genesis file to have a canonical genesis time

func initGenFiles(cfg Config, genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance, genFiles []string) error {
	_ = "STUB: not implemented"

	// set the accounts in the genesis state
	return nil
}

// set the balances in the genesis state

// generate empty genesis files for each validator and save

func writeFile(name string, dir string, contents []byte) error {
	_ = "STUB: not implemented"
	return nil
}
