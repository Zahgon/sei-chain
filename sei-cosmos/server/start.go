package server

// DONTCOVER

import (
	_ "net/http/pprof" //nolint:gosec

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/telemetry"
	"github.com/sei-protocol/sei-chain/sei-tendermint/node"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	// Tendermint full-node start flags
	flagAddress            = "address"
	flagTransport          = "transport"
	flagTraceStore         = "trace-store"
	flagCPUProfile         = "cpu-profile"
	FlagMinGasPrices       = "minimum-gas-prices"
	FlagHaltHeight         = "halt-height"
	FlagHaltTime           = "halt-time"
	FlagInterBlockCache    = "inter-block-cache"
	FlagUnsafeSkipUpgrades = "unsafe-skip-upgrades"
	FlagTrace              = "trace"
	FlagProfile            = "profile"
	FlagInvCheckPeriod     = "inv-check-period"

	// Legacy pruning flags (kept for backward compatibility)
	FlagPruning           = "pruning"
	FlagPruningKeepRecent = "pruning-keep-recent"
	FlagPruningKeepEvery  = "pruning-keep-every"
	FlagPruningInterval   = "pruning-interval"

	FlagIndexEvents        = "index-events"
	FlagMinRetainBlocks    = "min-retain-blocks"
	FlagCompactionInterval = "compaction-interval"
	FlagConcurrencyWorkers = "concurrency-workers"

	// state sync-related flags
	FlagStateSyncSnapshotInterval   = "state-sync.snapshot-interval"
	FlagStateSyncSnapshotKeepRecent = "state-sync.snapshot-keep-recent"
	FlagStateSyncSnapshotDir        = "state-sync.snapshot-directory"

	// gRPC-related flags
	flagGRPCOnly       = "grpc-only"
	flagGRPCEnable     = "grpc.enable"
	flagGRPCAddress    = "grpc.address"
	flagGRPCWebEnable  = "grpc-web.enable"
	flagGRPCWebAddress = "grpc-web.address"

	// archival related flags
	FlagArchivalVersion                = "archival-version"
	FlagArchivalDBType                 = "archival-db-type"
	FlagArchivalArweaveIndexDBFullPath = "archival-arweave-index-db-full-path"
	FlagArchivalArweaveNodeURL         = "archival-arweave-node-url"

	// chain info
	FlagChainID = "chain-id"
)

// StartCmd runs the service passed in with Tendermint in-process.
func StartCmd(appCreator types.AppCreator, defaultNodeHome string, tracerProviderOptions []trace.TracerProviderOption) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// Bind flags to the Context's Viper so the app construction can set
// options accordingly.

// TODO: Should this be bound to all interfaces?

//nolint:gosec // no read/write timeout to allow long running pprofs for debugging

func addStartNodeFlags(cmd *cobra.Command, defaultNodeHome string) {
	_ = "STUB: not implemented"
	return
}

// add support for all Tendermint-specific command line options

func mustMarkDeprecated(cmd *cobra.Command, name, message string) {
	_ = "STUB: not implemented"
	return
}

func startInProcess(
	ctx *Context,
	clientCtx client.Context,
	appCreator types.AppCreator,
	tracerProviderOptions []trace.TracerProviderOption,
	nodeMetricsProvider *node.NodeMetrics,
	apiMetrics *telemetry.Metrics,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the tx service to the gRPC router. We only need to register this
// service if API or gRPC is enabled, and avoid doing so in the general
// case, because it spawns a new local tendermint RPC client.

// assume server started successfully

// At this point it is safe to block the process if we're in gRPC only mode as
// we do not need to start Rosetta or handle any Tendermint related processes.

// wait for signal capture and gracefully return

// If GRPC is not enabled rosetta cannot work in online mode, so it works in
// offline mode.

// assume server started successfully

// Defer cancelling as the last so that it is called first during unwinding.

// wait for signal capture and gracefully return
