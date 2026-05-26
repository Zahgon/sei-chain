package cmd

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/app/params"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
)

const (
	// FlagOverwrite defines a flag to overwrite an existing genesis JSON file.
	FlagOverwrite = "overwrite"

	// FlagSeed defines a flag to initialize the private validator key from a specific seed.
	FlagRecover = "recover"

	// FlagMode defines the node mode flag.
	FlagMode = "mode"
)

// isValidMode checks if the given node mode is valid
func isValidMode(mode params.NodeMode) bool { _ = "STUB: not implemented"; return false }

type printInfo struct {
	Moniker    string          `json:"moniker" yaml:"moniker"`
	ChainID    string          `json:"chain_id" yaml:"chain_id"`
	NodeID     string          `json:"node_id" yaml:"node_id"`
	GenTxsDir  string          `json:"gentxs_dir" yaml:"gentxs_dir"`
	AppMessage json.RawMessage `json:"app_message" yaml:"app_message"`
}

func newPrintInfo(moniker, chainID, nodeID, genTxsDir string, appMessage json.RawMessage) printInfo {
	_ = "STUB: not implemented"
	return *new(printInfo)
}

func displayInfo(info printInfo) error { _ = "STUB: not implemented"; return nil }

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// Get node mode from flag

// Validate mode

// Create and configure Tendermint config (outputs to config.toml)

// Tendermint only supports "validator", "full", "seed" modes
// Archive nodes use "full" mode in Tendermint but have different app config

// Get bip39 mnemonic

// Write Tendermint config.toml

// Create and configure app config (outputs to app.toml)

// Configure EVM based on node mode

// Get custom template from root.go

// Build custom app config with mode-specific values

func checkConfigOverwrite(configPath string, overwrite bool) error {
	_ = "STUB: not implemented"
	return nil
}

// loadOrWriteGenesis loads existing genesis at genFile if present and !overwrite, else writes embedded (well-known) or default.
func loadOrWriteGenesis(genFile, chainID string, overwrite bool, mbm module.BasicManager, cdc codec.JSONCodec) (*types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ensureGenesisPathIsFile(genFile string) error { _ = "STUB: not implemented"; return nil }
