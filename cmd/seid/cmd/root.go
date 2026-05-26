package cmd

import (
	"io"
	"os"

	"github.com/sei-protocol/sei-chain/app"
	"github.com/sei-protocol/sei-chain/app/params"

	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/seilog"
	"github.com/spf13/cobra"
	dbm "github.com/tendermint/tm-db"
)

var logger = seilog.NewLogger("cmd", "seid", "cmd")

// Option configures root command option.
type Option func(*rootOptions)

// scaffoldingOptions keeps set of options to apply scaffolding.
//
//nolint:unused // preserving this becase don't know if it is needed.
type rootOptions struct{}

// NewRootCmd creates a new root command for a Cosmos SDK application
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	_ = "STUB: not implemented"
	return nil, *new(params.EncodingConfig)
}

// set the default command outputs

// Skip creating config.toml/app.toml when running "init"; init creates them itself.
// Otherwise the PreRun would create them in the init home, and init would then error

func initRootCmd(
	rootCmd *cobra.Command,
	encodingConfig params.EncodingConfig,
) {
	_ = "STUB: not implemented"
	return
}

// extend debug command

// add server commands

// add keybase, auxiliary RPC, query, and tx child commands

// queryCommand returns the sub-command to send queries to the app
func queryCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// txCommand returns the sub-command to send transactions to the app
func txCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func addModuleInitFlags(startCmd *cobra.Command) { _ = "STUB: not implemented"; return }

// newApp creates a new Cosmos SDK app
func newApp(
	db dbm.DB,
	traceStore io.Writer,
	tmConfig *tmcfg.Config,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	_ = "STUB: not implemented"
	return *new(servertypes.Application)
}

// This varies from the default value of 140_000_000 because we would like to appropriately represent the
// compute time required as a proportion of block gas used for a wasm contract that performs a lot of compute
// This makes it such that the wasm VM gas converts to sdk gas at a 6.66x rate vs that of the previous multiplier

// appExport creates a new simapp (optionally at a given height)
func appExport(
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	file *os.File,
) (servertypes.ExportedApp, error) {
	_ = "STUB: not implemented"
	return *new(servertypes.ExportedApp), nil
}

func getExportableApp(
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	appOpts servertypes.AppOptions,
) (*app.App, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPrimeNums(lo int, hi int) []int { _ = "STUB: not implemented"; return nil }

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
// nolint: staticcheck
func initAppConfig() (string, interface{}) {
	_ = "STUB: not implemented"
	// Optionally allow the chain developer to overwrite the SDK's default
	// server config.
	return "", nil
}

// The SDK's default minimum gas price is set to "" (empty value) inside
// app.toml. If left empty by validators, the node will halt on startup.
// However, the chain developer can set a default app.toml value for their
// validators here.
//
// In summary:
// - if you leave srvCfg.MinGasPrices = "", all validators MUST tweak their
//   own app.toml config,
// - if you set srvCfg.MinGasPrices non-empty, validators CAN tweak their
//   own app.toml to override, or use this default value.
//
// In simapp, we set the min gas prices to 0.

// Pruning configs

// Randomly generate pruning interval. Note this only takes affect if using custom pruning. We want the following properties:
//   - random: if everyone has the same value, the block that everyone prunes will be slow
//   - prime: no overlap

// Metrics

// Use shared CustomAppConfig from app_config.go
