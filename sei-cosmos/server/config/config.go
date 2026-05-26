package config

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/telemetry"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-db/config"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/spf13/viper"
)

const (
	// DefaultMinGasPrices defines the default minimum gas prices
	DefaultMinGasPrices = "0.01usei"

	// DefaultGRPCAddress defines the default address to bind the gRPC server to.
	DefaultGRPCAddress = "0.0.0.0:9090"

	// DefaultGRPCWebAddress defines the default address to bind the gRPC-web server to.
	DefaultGRPCWebAddress = "0.0.0.0:9091"

	// DefaultOccEanbled defines whether to use OCC for tx processing
	DefaultOccEnabled = true
)

var (
	// DefaultConcurrencyWorkers defines the default workers to use for concurrent transactions
	// Set to 2x CPU cores, capped between [10, 128]
	// - 2x CPU: In practice, goroutines often block on IO/network, so having more workers
	//   than CPU cores keeps CPUs busy. Load tests show only 60-70% CPU usage with 500 workers
	//   processing ~1500 txs/block, suggesting IO-bound workload benefits from oversubscription.
	// - Min 10: Minimum viable parallelism for transaction processing
	// - Max 128: Prevents unbounded goroutine creation on high-core machines. While 500 workers
	//   worked in tests, 128 provides sufficient parallelism without excessive scheduler overhead.
	DefaultConcurrencyWorkers = getConcurrencyWorkers()
)

// getConcurrencyWorkers returns the default number of concurrency workers
func getConcurrencyWorkers() int { _ = "STUB: not implemented"; return 0 }

// BaseConfig defines the server's basic configuration
type BaseConfig struct {
	// The minimum gas prices a validator is willing to accept for processing a
	// transaction. A transaction's fees must meet the minimum of any denomination
	// specified in this config (e.g. 0.25token1;0.0001token2).
	MinGasPrices string `mapstructure:"minimum-gas-prices"`

	Pruning           string `mapstructure:"pruning"`
	PruningKeepRecent string `mapstructure:"pruning-keep-recent"`
	PruningKeepEvery  string `mapstructure:"pruning-keep-every"`
	PruningInterval   string `mapstructure:"pruning-interval"`

	// HaltHeight contains a non-zero block height at which a node will gracefully
	// halt and shutdown that can be used to assist upgrades and testing.
	//
	// Note: Commitment of state will be attempted on the corresponding block.
	HaltHeight uint64 `mapstructure:"halt-height"`

	// HaltTime contains a non-zero minimum block time (in Unix seconds) at which
	// a node will gracefully halt and shutdown that can be used to assist
	// upgrades and testing.
	//
	// Note: Commitment of state will be attempted on the corresponding block.
	HaltTime uint64 `mapstructure:"halt-time"`

	// MinRetainBlocks defines the minimum block height offset from the current
	// block being committed, such that blocks past this offset may be pruned
	// from Tendermint. It is used as part of the process of determining the
	// ResponseCommit.RetainHeight value during ABCI Commit. A value of 0 indicates
	// that no blocks should be pruned.
	//
	// This configuration value is only responsible for pruning Tendermint blocks.
	// It has no bearing on application state pruning which is determined by the
	// "pruning-*" configurations.
	//
	// Note: Tendermint block pruning is dependant on this parameter in conunction
	// with the unbonding (safety threshold) period, state pruning and state sync
	// snapshot parameters to determine the correct minimum value of
	// ResponseCommit.RetainHeight.
	MinRetainBlocks uint64 `mapstructure:"min-retain-blocks"`

	// InterBlockCache enables inter-block caching.
	InterBlockCache bool `mapstructure:"inter-block-cache"`

	// IndexEvents defines the set of events in the form {eventType}.{attributeKey},
	// which informs Tendermint what to index. If empty, all events will be indexed.
	IndexEvents []string `mapstructure:"index-events"`

	// CompactionInterval sets (in seconds) the interval between forced levelDB
	// compaction. A value of 0 means no forced levelDB
	CompactionInterval uint64 `mapstructure:"compaction-interval"`

	// ConcurrencyWorkers defines the number of workers to use for concurrent
	// transaction execution. A value of -1 means unlimited workers.  Default value is 10.
	ConcurrencyWorkers int `mapstructure:"concurrency-workers"`
	// Whether to enable optimistic concurrency control for tx execution, default is true
	OccEnabled bool `mapstructure:"occ-enabled"`
}

// APIConfig defines the API listener configuration.
type APIConfig struct {
	// Enable defines if the API server should be enabled.
	Enable bool `mapstructure:"enable"`

	// Swagger defines if swagger documentation should automatically be registered.
	Swagger bool `mapstructure:"swagger"`

	// EnableUnsafeCORS defines if CORS should be enabled (unsafe - use it at your own risk)
	EnableUnsafeCORS bool `mapstructure:"enabled-unsafe-cors"`

	// Address defines the API server to listen on
	Address string `mapstructure:"address"`

	// MaxOpenConnections defines the number of maximum open connections
	MaxOpenConnections uint `mapstructure:"max-open-connections"`

	// RPCReadTimeout defines the Tendermint RPC read timeout (in seconds)
	RPCReadTimeout uint `mapstructure:"rpc-read-timeout"`

	// RPCWriteTimeout defines the Tendermint RPC write timeout (in seconds)
	RPCWriteTimeout uint `mapstructure:"rpc-write-timeout"`

	// RPCMaxBodyBytes defines the Tendermint maximum response body (in bytes)
	RPCMaxBodyBytes uint `mapstructure:"rpc-max-body-bytes"`

	// TODO: TLS/Proxy configuration.
	//
	// Ref: https://github.com/cosmos/cosmos-sdk/issues/6420
}

// RosettaConfig defines the Rosetta API listener configuration.
type RosettaConfig struct {
	// Address defines the API server to listen on
	Address string `mapstructure:"address"`

	// Blockchain defines the blockchain name
	// defaults to DefaultBlockchain
	Blockchain string `mapstructure:"blockchain"`

	// Network defines the network name
	Network string `mapstructure:"network"`

	// Retries defines the maximum number of retries
	// rosetta will do before quitting
	Retries int `mapstructure:"retries"`

	// Enable defines if the API server should be enabled.
	Enable bool `mapstructure:"enable"`

	// Offline defines if the server must be run in offline mode
	Offline bool `mapstructure:"offline"`
}

// GRPCConfig defines configuration for the gRPC server.
type GRPCConfig struct {
	// Enable defines if the gRPC server should be enabled.
	Enable bool `mapstructure:"enable"`

	// Address defines the API server to listen on
	Address string `mapstructure:"address"`
}

// GRPCWebConfig defines configuration for the gRPC-web server.
type GRPCWebConfig struct {
	// Enable defines if the gRPC-web should be enabled.
	Enable bool `mapstructure:"enable"`

	// Address defines the gRPC-web server to listen on
	Address string `mapstructure:"address"`

	// EnableUnsafeCORS defines if CORS should be enabled (unsafe - use it at your own risk)
	EnableUnsafeCORS bool `mapstructure:"enable-unsafe-cors"`
}

// StateSyncConfig defines the state sync snapshot configuration.
type StateSyncConfig struct {
	// SnapshotInterval sets the interval at which state sync snapshots are taken.
	// 0 disables snapshots. Must be a multiple of PruningKeepEvery.
	SnapshotInterval uint64 `mapstructure:"snapshot-interval"`

	// SnapshotKeepRecent sets the number of recent state sync snapshots to keep.
	// 0 keeps all snapshots.
	SnapshotKeepRecent uint32 `mapstructure:"snapshot-keep-recent"`

	// SnapshotDirectory sets the parent directory for where state sync snapshots are persisted.
	// Default is empty which will then store under the app home directory.
	SnapshotDirectory string `mapstructure:"snapshot-directory"`
}

// GenesisConfig defines the genesis export, validation, and import configuration
type GenesisConfig struct {
	// StreamImport defines if the genesis.json is in stream form or not.
	StreamImport bool `mapstructure:"stream-import"`

	// GenesisStreamFile sets the genesis json file from which to stream from
	GenesisStreamFile string `mapstructure:"genesis-stream-file"`
}

// Config defines the server's top level configuration
type Config struct {
	BaseConfig `mapstructure:",squash"`

	// Telemetry defines the application telemetry configuration
	Telemetry   telemetry.Config         `mapstructure:"telemetry"`
	API         APIConfig                `mapstructure:"api"`
	GRPC        GRPCConfig               `mapstructure:"grpc"`
	Rosetta     RosettaConfig            `mapstructure:"rosetta"`
	GRPCWeb     GRPCWebConfig            `mapstructure:"grpc-web"`
	StateSync   StateSyncConfig          `mapstructure:"state-sync"`
	StateCommit config.StateCommitConfig `mapstructure:"state-commit"`
	StateStore  config.StateStoreConfig  `mapstructure:"state-store"`
	Genesis     GenesisConfig            `mapstructure:"genesis"`
}

// SetMinGasPrices sets the validator's minimum gas prices.
func (c *Config) SetMinGasPrices(gasPrices sdk.DecCoins) { _ = "STUB: not implemented"; return }

// GetMinGasPrices returns the validator's minimum gas prices based on the set
// configuration.
func (c *Config) GetMinGasPrices() sdk.DecCoins {
	_ = "STUB: not implemented"
	return *new(sdk.DecCoins)
}

// DefaultConfig returns server's default configuration.
func DefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// GetConfig returns a fully parsed Config object.
func GetConfig(v *viper.Viper) (Config, error) { _ = "STUB: not implemented"; return *new(Config), nil }

// Resolve sc-write-mode through ParseWriteMode so that misspellings fail
// at parse-time with a clear error rather than later from
// StateCommitConfig.Validate(). An empty value preserves the in-code
// default, matching the behavior of app/seidb.go.

// ValidateBasic returns an error if min-gas-prices field is empty in BaseConfig. Otherwise, it returns nil.
func (c Config) ValidateBasic(tendermintConfig *tmcfg.Config) error {
	_ = "STUB: not implemented"
	return nil
}
