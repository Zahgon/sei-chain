package cryptosim

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
)

type TransactionExecutor struct {
	ctx    context.Context
	cancel context.CancelFunc
	config *CryptoSimConfig

	// The database for the benchmark.
	database *Database

	// The address of the fee collection account.
	feeCollectionAddress []byte

	// The Incoming transactions to be executed.
	workChan chan any

	// Used to time the execution of transactions.
	phaseTimer *metrics.PhaseTimer
}

// A request to flush the transaction executor.
type flushRequest struct {
	doneChan chan struct{}
}

// A single threaded transaction executor.
func NewTransactionExecutor(
	ctx context.Context,
	cancel context.CancelFunc,
	config *CryptoSimConfig,
	database *Database,
	feeCollectionAddress []byte,
	queueSize int,
	metrics *CryptosimMetrics,
) *TransactionExecutor {
	_ = "STUB: not implemented"
	return nil
}

// Schedule a transaction for execution.
func (e *TransactionExecutor) ScheduleForExecution(txn *transaction) {
	_ = "STUB: not implemented"
	return
}

// Blocks until all currently queued transactions have been executed.
func (e *TransactionExecutor) Flush() { _ = "STUB: not implemented"; return }

func (e *TransactionExecutor) mainLoop() { _ = "STUB: not implemented"; return }
