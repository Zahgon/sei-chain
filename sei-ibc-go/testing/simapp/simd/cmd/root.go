package cmd

import (
	"io"
	"os"

	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/spf13/cobra"
	dbm "github.com/tendermint/tm-db"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/testing/simapp/params"
)

// NewRootCmd creates a new root command for simd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	_ = "STUB: not implemented"
	return nil, *new(params.EncodingConfig)
}

// In simapp, we don't use any prefix for env variables.

// set the default command outputs

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func initAppConfig() (string, interface{}) {
	_ = "STUB: not implemented"
	// The following code snippet is just for reference.
	return "", nil
}

// WASMConfig defines configuration for the wasm module.

// This is the maximum sdk gas (wasm and storage) that we allow for any x/wasm "smart" queries

// Address defines the gRPC-web server to listen on

// Optionally allow the chain developer to overwrite the SDK's default
// server config.

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

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	_ = "STUB: not implemented"
	return
}

// add keybase, auxiliary RPC, query, and tx child commands

// add rosetta

func addModuleInitFlags(startCmd *cobra.Command) { _ = "STUB: not implemented"; return }

func queryCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func txCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

type appCreator struct {
	encCfg params.EncodingConfig
}

// newApp is an appCreator
func (a appCreator) newApp(db dbm.DB, traceStore io.Writer, tmConfig *tmcfg.Config, appOpts servertypes.AppOptions) servertypes.Application {
	_ = "STUB: not implemented"
	return *new(servertypes.Application)
}

// appExport creates a new simapp (optionally at a given height)
// and exports state.
func (a appCreator) appExport(db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	_ *os.File,
) (servertypes.ExportedApp, error) {
	_ = "STUB: not implemented"
	return *new(servertypes.ExportedApp), nil
}
