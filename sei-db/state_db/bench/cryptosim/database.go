package cryptosim

import (
	"github.com/sei-protocol/sei-chain/sei-db/state_db/bench/wrappers"
)

// Encapsulates the database for the cryptosim benchmark.
type Database struct {
	// The configuration for the benchmark.
	config *CryptoSimConfig

	// The database implementation to use for the benchmark.
	db wrappers.DBWrapper

	// The total number of transactions executed by the benchmark since it last started.
	transactionCount int64

	// A count of the number of transactions in the current batch.
	transactionsInCurrentBlock int64

	// The number of blocks that have been executed since the last commit.
	uncommittedBlockCount int64

	// The next block number to be persisted. Tracked internally and incremented after each finalized block.
	nextBlockNumber uint64

	// The current batch of key-value pairs waiting to be committed. Represents changes we are accumulating
	// as part of a simulated "block". Stored as value []byte; converted to NamedChangeSet when applied to the DB.
	batch *SyncMap[string, []byte]

	// A method that flushes the executors.
	flushFunc func()

	// The metrics for the benchmark.
	metrics *CryptosimMetrics
}

// Creates a new database for the cryptosim benchmark.
func NewDatabase(
	config *CryptoSimConfig,
	db wrappers.DBWrapper,
	metrics *CryptosimMetrics,
	initialNextBlockNumber uint64,
) *Database {
	_ = "STUB: not implemented"
	return nil
}

// Insert a key-value pair into the database/cache.
//
// This method is safe to call concurrently with other calls to Put() and Get(). Is not thread
// safe with FinalizeBlock(). It is not thread safe to modify the returned value (make a copy first).
func (d *Database) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

// Retrieve a value from the database/cache.
//
// This method is safe to call concurrently with other calls to Put() and Get(). Is not thread
// safe with FinalizeBlock().
func (d *Database) Get(key []byte) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Signal that a transaction has been added to the current block.
func (d *Database) IncrementTransactionCount() { _ = "STUB: not implemented"; return }

// Reset the transaction count. Useful for when changing test phases.
func (d *Database) ResetTransactionCount() { _ = "STUB: not implemented"; return }

// Get the total number of transactions executed by the benchmark since it last started.
func (d *Database) TransactionCount() int64 { _ = "STUB: not implemented"; return 0 }

// Commit the current batch if it has reached the configured number of transactions.
// Returns true if the batch was finalized, false if not.
func (d *Database) MaybeFinalizeBlock(
	nextAccountID int64,
	nextErc20ContractID int64,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Push the current block out to the database.
func (d *Database) FinalizeBlock(
	nextAccountID int64,
	nextErc20ContractID int64,
	forceCommit bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for all transactions in the current block to be executed.

// Persist the account ID counter in every batch.

//nolint:gosec // G115 - nextAccountID is benchmark counter, overflow acceptable

// Persist the ERC20 contract ID counter in every batch.

//nolint:gosec // G115 - nextErc20ContractID is benchmark counter, overflow acceptable

// Persist the block number counter in every batch.

// Periodically commit the changes to the database.

// Close the database and release any resources.
func (d *Database) Close(nextAccountID int64, nextErc20ContractID int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Close the database and release any resources without finalizing the last batch.
func (d *Database) CloseWithoutFinalizing() error { _ = "STUB: not implemented"; return nil }

// Set the function that flushes the executors. This setter is required to break a circular dependency.
func (d *Database) SetFlushFunc(flushFunc func()) { _ = "STUB: not implemented"; return }
