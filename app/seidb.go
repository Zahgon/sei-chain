package app

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	"github.com/sei-protocol/sei-chain/sei-db/config"
	seidb "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

const (
	// SC Store configs
	FlagSCEnable                     = "state-commit.sc-enable"
	FlagSCDirectory                  = "state-commit.sc-directory"
	FlagSCAsyncCommitBuffer          = "state-commit.sc-async-commit-buffer"
	FlagSCSnapshotKeepRecent         = "state-commit.sc-keep-recent"
	FlagSCSnapshotInterval           = "state-commit.sc-snapshot-interval"
	FlagSCSnapshotMinTimeInterval    = "state-commit.sc-snapshot-min-time-interval"
	FlagSCSnapshotWriterLimit        = "state-commit.sc-snapshot-writer-limit"
	FlagSCSnapshotPrefetchThreshold  = "state-commit.sc-snapshot-prefetch-threshold"
	FlagSCSnapshotWriteRateMBps      = "state-commit.sc-snapshot-write-rate-mbps"
	FlagSCHistoricalProofMaxInFlight = "state-commit.sc-historical-proof-max-inflight"
	FlagSCHistoricalProofRateLimit   = "state-commit.sc-historical-proof-rate-limit"
	FlagSCHistoricalProofBurst       = "state-commit.sc-historical-proof-burst"
	FlagSCWriteMode                  = "state-commit.sc-write-mode"

	// SS Store configs
	FlagSSEnable            = "state-store.ss-enable"
	FlagSSDirectory         = "state-store.ss-db-directory"
	FlagSSBackend           = "state-store.ss-backend"
	FlagSSAsyncWriterBuffer = "state-store.ss-async-write-buffer"
	FlagSSKeepRecent        = "state-store.ss-keep-recent"
	FlagSSPruneInterval     = "state-store.ss-prune-interval"
	FlagSSImportNumWorkers  = "state-store.ss-import-num-workers"

	// EVM SS optimization (embedded in SS config, controlled via write/read mode)
	FlagEVMSSDirectory   = "state-store.evm-ss-db-directory"
	FlagEVMSSSplit       = "state-store.evm-ss-split"
	FlagEVMSSSeparateDBs = "state-store.evm-ss-separate-dbs"

	// Other configs
	FlagSnapshotInterval = "state-sync.snapshot-interval"
)

var GigaKeys = []string{"evm", "bank"}

func SetupSeiDB(
	homePath string,
	appOpts servertypes.AppOptions,
	baseAppOptions []func(*baseapp.BaseApp),
) ([]func(*baseapp.BaseApp), seidb.StateStore) {
	_ = "STUB: not implemented"
	return nil, *new(seidb.StateStore)
}

// cms must be overridden before the other options, because they may use the cms,
// make sure the cms aren't be overridden by the other options later on.

func parseSCConfigs(appOpts servertypes.AppOptions) config.StateCommitConfig {
	_ = "STUB: not implemented"
	return *new(config.StateCommitConfig)
}

func parseSSConfigs(appOpts servertypes.AppOptions) config.StateStoreConfig {
	_ = "STUB: not implemented"
	return *new(config.StateStoreConfig)
}

// EVM optimization fields (embedded in SS config)

func validateConfigs(appOpts servertypes.AppOptions) { _ = "STUB: not implemented"; return }

// Make sure when snapshot is enabled, we should enable SS store
