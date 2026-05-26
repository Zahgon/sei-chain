package cryptosim

import (
	"iter"

	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
)

// A simulated block of transactions.
type block struct {
	config *CryptoSimConfig

	// The transactions in the block.
	transactions []*transaction

	// If receipt generation is enabled, this will contain the receipts for each transaction in the block.
	reciepts []*evmtypes.Receipt

	// The block number. This is not currently preserved across benchmark restarts, but otherwise monotonically
	// increases as you'd expect.
	blockNumber int64

	// The next account ID to be used when creating a new account, as of the end of this block.
	nextAccountID int64

	// The number of cold accounts, as of the end of this block.
	numberOfColdAccounts int64

	// The next ERC20 contract ID to be used when creating a new ERC20 contract, as of the end of this block.
	nextErc20ContractID int64

	metrics *CryptosimMetrics
}

// Creates a new block with the given capacity.
func NewBlock(
	config *CryptoSimConfig,
	metrics *CryptosimMetrics,
	blockNumber int64,
	capacity int,
) *block {
	_ = "STUB: not implemented"
	return nil
}

// Returns an iterator over the transactions in the block.
func (b *block) Iterator() iter.Seq[*transaction] { _ = "STUB: not implemented"; return nil }

// Adds a transaction to the block.
func (b *block) AddTransaction(txn *transaction) { _ = "STUB: not implemented"; return }

// Adds a receipt to the block.
func (b *block) AddReceipt(receipt *evmtypes.Receipt) { _ = "STUB: not implemented"; return }

// Returns the block number.
func (b *block) BlockNumber() int64 { _ = "STUB: not implemented"; return 0 }

// Sets information about account state as of the end of this block.
func (b *block) SetBlockAccountStats(
	nextAccountID int64,
	numberOfColdAccounts int64,
	nextErc20ContractID int64,
) {
	_ = "STUB: not implemented"
	return
}

// This method should be called after a block is finished executing and finalized.
// Reports metrics about the block.
func (b *block) ReportBlockMetrics() { _ = "STUB: not implemented"; return }

// Returns the next account ID to be used when creating a new account, as of the end of this block.
func (b *block) NextAccountID() int64 { _ = "STUB: not implemented"; return 0 }

// Returns the next ERC20 contract ID to be used when creating a new ERC20 contract, as of the end of this block.
func (b *block) NextErc20ContractID() int64 { _ = "STUB: not implemented"; return 0 }

// Returns the number of transactions in the block.
func (b *block) TransactionCount() int64 { _ = "STUB: not implemented"; return 0 }
