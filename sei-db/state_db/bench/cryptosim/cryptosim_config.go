package cryptosim

import (
	"github.com/sei-protocol/sei-chain/sei-db/config"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/bench/wrappers"
	flatkvConfig "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/config"
)

const (
	minPaddedAccountSize        = 8
	minErc20StorageSlotSize     = 32
	minErc20InteractionsPerAcct = 1
	receiptReadModeCache        = "cache"
	receiptReadModeDuckDB       = "duckdb"
)

// Defines the configuration for the cryptosim benchmark.
type CryptoSimConfig struct {

	// The number of hot accounts. Hot accounts are very frequently used. The number of hot accounts does
	// not change after the benchmark starts.
	//
	// Future work: add different distributions of hot account access. Currently, distribution is flat.
	NumberOfHotAccounts int

	// The minimum number of cold accounts that should be in the DB prior to the start of the benchmark.
	// Cold accounts are occasionally used, but not frequently.
	MinimumNumberOfColdAccounts int

	// The minimum number of dormant accounts that should be in the DB prior to the start of the benchmark.
	// Dormant accounts are not used after they are created.
	MinimumNumberOfDormantAccounts int

	// When creating a new account, this is the probability that the number of dormant accounts will be increased
	// by one. Should be a value between 0.0 and 1.0. A value of 1.0 means that all new account creation will increase
	// the number of dormant accounts. A value of 0.0 means all new account creation will increase the number of
	// cold accounts.
	NewAccountDormancyProbability float64

	// When selecting an account for a transaction, select a hot account with this probability. Should be
	// a value between 0.0 and 1.0.
	HotAccountProbability float64

	// When selecting a non-hot account for a transaction, the benchmark will create a new account with this
	// probability. Should be a value between 0.0 and 1.0.
	NewAccountProbability float64

	// Each account contains an integer value used to track a balance, plus a bunch of random
	// bytes for padding. This is the total size of the account after padding is added.
	PaddedAccountSize int

	// The minimum number of ERC20 contracts that should be in the DB prior to the start of the benchmark.
	// If there are fewer than this number of contracts, the benchmark will first create the necessary
	// contracts before starting its regular operations.
	MinimumNumberOfErc20Contracts int

	// When selecting an ERC20 contract for a transaction, select a hot ERC20 contract with this probability.
	// Should be a value between 0.0 and 1.0.
	HotErc20ContractProbability float64

	// The number of hot ERC20 contracts.
	HotErc20ContractSetSize int

	// The size of the a simulated ERC20 contract, in bytes.
	Erc20ContractSize int

	// The size of a simulated ERC20 storage slot, in bytes.
	Erc20StorageSlotSize int

	// The size of a simulated account balance, in bytes.
	AccountBalanceSize int

	// The number of ERC20 tokens that each account will interact with.
	// Each account will have an eth storage slot for tracking the balance of each ERC20 token it owns.
	// It is not legal to modify this value after the benchmark has started.
	Erc20InteractionsPerAccount int

	// The number of transactions that will be processed in each "block".
	TransactionsPerBlock int

	// Commit is called on the database after this many blocks have been processed.
	BlocksPerCommit int

	// The directory to store the benchmark data.
	DataDir string

	// The seed to use for the random number generator. Altering this seed for a pre-existing DB will result
	// in undefined behavior, don't change the seed unless you are starting a new run from scratch.
	Seed int64

	// The size of the CannedRandom buffer. Similar to the seed, altering this size for a pre-existing DB will result
	// in undefined behavior, don't change the size unless you are starting a new run from scratch.
	CannedRandomSize int

	// The backend to use for the benchmark database.
	Backend wrappers.DBType

	// StateStoreConfig controls SS-backed benchmark backends such as SSComposite.
	// The default preserves the benchmark SS defaults: pebbledb, async buffer 100.
	StateStoreConfig *config.StateStoreConfig

	// HistoricalOffload configures the transport used by the
	// SSHistoricalOffload backend.
	HistoricalOffload *wrappers.HistoricalOffloadConfig

	// This field is ignored, but allows for a comment to be added to the config file.
	// Something, something, why in the name of all things holy doesn't json support comments?
	Comment string

	// If this many seconds go by without a console update, the benchmark will print a report to the console.
	ConsoleUpdateIntervalSeconds float64

	// If this many transactions are executed without a console update, the benchmark will print a report to the console.
	ConsoleUpdateIntervalTransactions float64

	// When setting up the benchmark, print a console update after adding this many accounts to the DB.
	SetupUpdateIntervalCount int64

	// Run a number of threads equal to the number of cores on the host machine, multiplied by this value.
	ThreadsPerCore float64

	// Increase or decrease the thread count by this many threads. Total thread count is a function of
	// ThreadsPerCore and ConstantThreadCount.
	ConstantThreadCount int

	// The size of the queue for each transaction executor.
	ExecutorQueueSize int

	// The amount of time to run the benchmark for. If 0, the benchmark will run until it is stopped.
	MaxRuntimeSeconds int

	// Address for the Prometheus metrics HTTP server (e.g. ":9090"). If empty, metrics are disabled.
	MetricsAddr string

	// The probability of capturing detailed metrics about a transaction. Should be a value between 0.0 and 1.0.
	TransactionMetricsSampleRate float64

	// How often (in seconds) to scrape background metrics (data dir size, process I/O).
	// If 0, background metrics are disabled.
	BackgroundMetricsScrapeInterval int

	// If true, pressing Enter in the terminal will toggle suspend/resume of the benchmark.
	// If false, Enter has no effect.
	EnableSuspension bool

	// If true, the data directory will be deleted on startup if it exists.
	DeleteDataDirOnStartup bool

	// If true, the log directory will be deleted on startup if it exists.
	DeleteLogDirOnStartup bool

	// If true, the data directory will be deleted on a clean shutdown.
	DeleteDataDirOnShutdown bool

	// If true, the log directory will be deleted on a clean shutdown.
	DeleteLogDirOnShutdown bool

	// Configures the FlatKV database. Ignored if Backend is not "FlatKV".
	FlatKVConfig *flatkvConfig.Config

	// The capacity of the channel that holds blocks awaiting execution.
	BlockChannelCapacity int

	// If true, the benchmark will generate receipts for each transaction in each block and
	// feed those receipts into the receipt store.
	GenerateReceipts bool

	// The capacity of the channel that holds blocks sent to the receipt store.
	RecieptChannelCapacity int

	// If true, disables simulation of transaction execution, and writes very little to the database. This is
	// potentially useful when benchmarking things other than state storage (e.g. the receipt store).
	//
	// Note that switching execution on after previously running with execution disabled may result in buggy behavior,
	// as the benchmark will not be properly maintaining DB state when transaction execution is disabled. In order
	// to switch transaction execution back on, it is necessary to delete the on-disk database and start over.
	DisableTransactionExecution bool

	// If true, skip transaction-time database reads and only issue writes. Useful
	// when benchmarking write-path overhead for backends that are not consulted by
	// execution-time reads in production.
	DisableTransactionReads bool

	// If greater than 0, the benchmark will throttle the transaction rate to this value, in hertz.
	MaxTPS float64

	// Number of concurrent reader goroutines issuing receipt lookups. 0 disables reads.
	ReceiptReadConcurrency int

	// Target total receipt reads per second across all reader goroutines.
	// Reads are distributed evenly across readers.
	ReceiptReadsPerSecond int

	// Controls which block range receipt-by-hash reads target.
	// "cache" = only read receipts in the cache window (guaranteed cache hit).
	// "duckdb" = only read receipts older than the cache window (guaranteed cache miss, DuckDB fallback).
	// Required when ReceiptReadConcurrency > 0.
	ReceiptReadMode string

	// ReceiptTxIndexBackend selects the tx-hash index implementation for the
	// parquet receipt store. Set to "pebbledb" (the default) to maintain a
	// Pebble-backed tx_hash -> block_number index so receipt-by-hash lookups
	// target a single parquet file instead of scanning all files. Set to ""
	// to disable the index and fall back to full DuckDB scans.
	ReceiptTxIndexBackend string

	// Number of concurrent goroutines issuing log filter (eth_getLogs) queries. 0 disables log filter reads.
	// These goroutines are independent from the receipt reader goroutines.
	ReceiptLogFilterReadConcurrency int

	// Target total log filter reads per second across all log filter goroutines.
	ReceiptLogFilterReadsPerSecond int

	// Controls which block range log filter reads target.
	// "cache" = only query blocks in the cache window (DuckDB skipped).
	// "duckdb" = only query blocks older than the cache window (cache returns nothing).
	// Required when ReceiptLogFilterReadConcurrency > 0.
	ReceiptLogFilterReadMode string

	// Minimum number of blocks in a log filter query range. Default 1.
	ReceiptLogFilterMinBlockRange int

	// Maximum number of blocks in a log filter query range. Default 10.
	ReceiptLogFilterMaxBlockRange int

	// Number of recent blocks to keep before pruning parquet files. 0 disables pruning.
	ReceiptKeepRecent int64

	// Interval in seconds between prune checks. 0 disables pruning.
	ReceiptPruneIntervalSeconds int64

	// Directory for seilog output files. Independent of DataDir so logs and data
	// live in separate trees. Supports ~ expansion and relative paths (resolved
	// from cwd). Must be set, there is no default.
	LogDir string

	// Log level for seilog output. Valid values: debug, info, warn, error.
	LogLevel string
}

// Returns the default configuration for the cryptosim benchmark.
func DefaultCryptoSimConfig() *CryptoSimConfig {
	_ = "STUB: not implemented"

	// Note: if you add new fields or modify default values, be sure to keep config/basic-config.json in sync.
	// That file should contain every available config set to its default value, as a reference.
	return nil
}

// 2kb

// 1GB

// StringifiedConfig returns the config as human-readable, multi-line JSON.
func (c *CryptoSimConfig) StringifiedConfig() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Validate checks that the configuration is sane and returns an error if not.
func (c *CryptoSimConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// LoadConfigFromFile parses a JSON config file at the given path.
// Returns defaults with file values overlaid. Fails if the file contains
// unrecognized configuration keys.
func LoadConfigFromFile(path string) (*CryptoSimConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil

	//nolint:gosec // G304 - path comes from config file, filepath.Clean used to mitigate traversal
}
