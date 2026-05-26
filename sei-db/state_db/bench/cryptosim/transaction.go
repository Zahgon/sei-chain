package cryptosim

import (
	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
)

// The data needed to execute a transaction.
type transaction struct {
	// The simulated ERC20 contract that will be interacted with. This value is read.
	erc20Contract []byte

	// The source account that will be interacted with. This value is read and written.
	srcAccount []byte
	// If true, the source account is new and needs to be created.
	isSrcNew bool

	// The destination account that will be interacted with. This value is read and written.
	dstAccount []byte
	// If true, the destination account is new and needs to be created.
	isDstNew bool

	// The source account's storage slot that will be interacted with. This value is read and written.
	srcAccountSlot []byte
	// The destination account's storage slot that will be interacted with. This value is read and written.
	dstAccountSlot []byte

	// Pre-generated random value for the source account's new native balance.
	newSrcBalance []byte
	// Pre-generated random value for the destination account's new native balance.
	newDstBalance []byte
	// Pre-generated random value for the fee collection account's new native balance.
	newFeeBalance []byte
	// Pre-generated random value for the source account's ERC20 storage slot.
	newSrcAccountSlot []byte
	// Pre-generated random value for the destination account's ERC20 storage slot.
	newDstAccountSlot []byte

	// If true, capture detailed (and potentially expensive) metrics about this transaction.
	// We may only sample a small percentage of transactions with this flag set to true.
	captureMetrics bool
}

// Generate all data needed to execute a transaction.
//
// This method is not thread safe to call concurrently with other calls to BuildTransaction().
func BuildTransaction(
	dataGenerator *DataGenerator,
) (*transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Execute the transaction.
//
// This method is thread safe with other calls to Execute(),
// but must not be called concurrently with CryptoSim.finalizeBlock().
func (txn *transaction) Execute(
	database *Database,
	feeCollectionAddress []byte,
	phaseTimer *metrics.PhaseTimer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the simulated ERC20 contract.

// Read the following:
// - the sender's native balance / nonce / codehash
// - the receiver's native balance
// - the sender's storage slot for the ERC20 contract
// - the receiver's storage slot for the ERC20 contract
// - the fee collection account's native balance

// Read the sender's native balance / nonce / codehash.
// Technically, we are just requesting to read the codehash, but internally the codehash is bundled with
// the nonce and balance, so all of this data will be read from low level storage, even if it isn't being
// returned to the caller.

// Read the receiver's native balance / nonce / codehash.

// Read the sender's storage slot for the ERC20 contract.
// We don't care if the value isn't in the DB yet, since we don't pre-populate the database with storage slots.

// Read the receiver's storage slot for the ERC20 contract.
// We don't care if the value isn't in the DB yet, since we don't pre-populate the database with storage slots.

// Read the fee collection account's native balance.

// Write the following:
// - the sender's native balance
// - the receiver's native balance
// - the sender's storage slot for the ERC20 contract
// - the receiver's storage slot for the ERC20 contract
// - the fee collection account's native balance

// Write the sender's account data.

// Write the receiver's account data.

// Write the sender's storage slot for the ERC20 contract.

// Write the receiver's storage slot for the ERC20 contract.

// Write the fee collection account's native balance.

// Returns true if metrics should be captured for this transaction.
func (txn *transaction) ShouldCaptureMetrics() bool { _ = "STUB: not implemented"; return false }
