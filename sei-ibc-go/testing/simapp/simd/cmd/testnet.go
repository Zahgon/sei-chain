package cmd

// DONTCOVER

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	tmconfig "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/spf13/cobra"
)

var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagNumValidators     = "v"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagStartingIPAddress = "starting-ip-address"
)

// get cmd to initialize all files for tendermint testnet and application
func testnetCmd(mbm module.BasicManager, genBalIterator banktypes.GenesisBalancesIterator) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

const nodeDirPerm = 0o755

// Initialize the testnet
func InitTestnet(
	clientCtx client.Context,
	cmd *cobra.Command,
	nodeConfig *tmconfig.Config,
	mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
	outputDir,
	chainID,
	minGasPrices,
	nodeDirPrefix,
	nodeDaemonHome,
	startingIPAddress,
	keyringBackend,
	algoStr string,
	numValidators int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// generate private keys, node IDs, and initial transactions

// save private key seed words

func initGenFiles(
	clientCtx client.Context, mbm module.BasicManager, chainID string,
	genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance,
	genFiles []string, numValidators int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// set the accounts in the genesis state

// set the balances in the genesis state

// generate empty genesis files for each validator and save

func collectGenFiles(
	clientCtx client.Context, nodeConfig *tmconfig.Config, chainID string,
	nodeIDs []string, valPubKeys []cryptotypes.PubKey, numValidators int,
	outputDir, nodeDirPrefix, nodeDaemonHome string, genBalIterator banktypes.GenesisBalancesIterator,
) error {
	_ = "STUB: not implemented"
	return nil
}

// no monotonic component

// set the canonical application state (they should not differ)

// overwrite each validator's genesis file to have a canonical genesis time

func getIP(i int, startingIPAddr string) (ip string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func calculateIP(ip string, i int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func writeFile(name string, dir string, contents []byte) error {
	_ = "STUB: not implemented"
	return nil
}
