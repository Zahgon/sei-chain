package config

// AppOptions is a minimal interface for reading config (e.g. from Viper).
// Implemented by sei-cosmos server/types.AppOptions; defined here to avoid import cycles.
type AppOptions interface {
	Get(string) interface{}
}

const (
	flagRSDBDirectory          = "receipt-store.db-directory"
	flagRSBackend              = "receipt-store.rs-backend"
	flagRSMisnamedBackend      = "receipt-store.backend"
	flagRSAsyncWriteBuffer     = "receipt-store.async-write-buffer"
	flagRSPruneIntervalSeconds = "receipt-store.prune-interval-seconds"
	flagRSTxIndexBackend       = "receipt-store.tx-index-backend"

	ReceiptTxIndexBackendNone   = ""
	ReceiptTxIndexBackendPebble = "pebbledb"
)

func NormalizeReceiptTxIndexBackend(backend string) string { _ = "STUB: not implemented"; return "" }

// ReceiptStoreConfig defines configuration for the receipt store database.
type ReceiptStoreConfig struct {
	// DBDirectory defines the directory to store the receipt store db files
	// If not explicitly set, default to application home directory
	// default to empty
	DBDirectory string `mapstructure:"db-directory"`

	// Backend defines the backend database used for receipt-store.
	// Supported backends: pebbledb (aka pebble), parquet
	// defaults to pebbledb
	Backend string `mapstructure:"rs-backend"`

	// AsyncWriteBuffer defines the async queue length for commits to be applied to receipt store
	// Applies only to the pebbledb backend.
	// Set <= 0 for synchronous writes.
	// defaults to 100
	AsyncWriteBuffer int `mapstructure:"async-write-buffer"`

	// KeepRecent defines the number of versions to keep in receipt store.
	// Setting it to 0 means keep everything (no pruning).
	// This is NOT read from receipt-store config; it is always derived from
	// the global min-retain-blocks flag at the app layer.
	KeepRecent int `mapstructure:"-"`

	// PruneIntervalSeconds defines the interval in seconds to trigger pruning
	// default to every 600 seconds
	PruneIntervalSeconds int `mapstructure:"prune-interval-seconds"`

	// TxIndexBackend selects the tx-hash index implementation used by the
	// parquet receipt store. Set to "pebbledb" (the default) to maintain a
	// Pebble-backed tx_hash -> block_number index alongside parquet files so
	// receipt-by-hash lookups can target a single file instead of scanning all
	// files. Set to "" to disable the index; receipt-by-hash lookups that miss
	// the in-memory cache then fail (no full-parquet scan). Ignored when the
	// receipt backend is not parquet.
	TxIndexBackend string `mapstructure:"tx-index-backend"`
}

// DefaultReceiptStoreConfig returns the default ReceiptStoreConfig.
// KeepRecent defaults to 0 (no pruning). The app layer is responsible
// for setting KeepRecent from the global min-retain-blocks flag.
func DefaultReceiptStoreConfig() ReceiptStoreConfig {
	_ = "STUB: not implemented"
	return *new(ReceiptStoreConfig)
}

// ReadReceiptConfig reads receipt store config from app options (e.g. TOML / Viper).
func ReadReceiptConfig(opts AppOptions) (ReceiptStoreConfig, error) {
	_ = "STUB: not implemented"
	return *new(ReceiptStoreConfig), nil
}
