package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/ethereum/go-ethereum/core"
	"github.com/sei-protocol/sei-chain/x/evm/state"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

// Number of blocks between legacy receipt migration batches
const LegacyReceiptMigrationInterval int64 = 10

// Number of receipts to migrate per batch
const LegacyReceiptMigrationBatchSize int = 100

// SetTransientReceipt sets a data structure that stores EVM specific transaction metadata.
func (k *Keeper) SetTransientReceipt(ctx sdk.Context, txHash common.Hash, receipt *types.Receipt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) GetTransientReceipt(ctx sdk.Context, txHash common.Hash, txIndex uint64) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Keeper) DeleteTransientReceipt(ctx sdk.Context, txHash common.Hash, txIndex uint64) {
	_ = "STUB: not implemented"
	return
}

// GetReceipt returns a data structure that stores EVM specific transaction metadata.
// Many EVM applications (e.g. MetaMask) relies on being on able to query receipt
// by EVM transaction hash (not Sei transaction hash) to function properly.
func (k *Keeper) GetReceipt(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only used for testing
func (k *Keeper) GetReceiptFromReceiptStore(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetReceiptWithRetry attempts to get a receipt with retries to handle race conditions
// where the receipt might not be immediately available after the transaction.
func (k *Keeper) GetReceiptWithRetry(ctx sdk.Context, txHash common.Hash, maxRetries int) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If it's not a "not found" error, return immediately

// Wait before retrying, with increasing delay, 200ms, 400ms, 600ms, etc.

//	MockReceipt sets a data structure that stores EVM specific transaction metadata.
//
// this is currently used by a number of tests to set receipts at the moment
func (k *Keeper) MockReceipt(ctx sdk.Context, txHash common.Hash, rcpt *types.Receipt) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) FlushTransientReceipts(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) FlushTransientReceiptsAsync(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func isLegacyReceipt(ctx sdk.Context, receipt *types.Receipt) bool {
	_ = "STUB: not implemented"
	return false
}

func (k *Keeper) flushTransientReceipts(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// TransientReceiptStore is recreated on commit meaning it will only contain receipts for a single block at a time
// and will never flush a subset of block's receipts.
// However in our test suite it can happen that the transient store can contain receipts from different blocks
// and we need to account for that.

// MigrateLegacyReceiptsBatch moves up to batchSize receipts from the legacy KV store
// into the persistent receipt store and deletes them from the legacy store.
// It returns the number of receipts migrated.
func (k *Keeper) MigrateLegacyReceiptsBatch(ctx sdk.Context, batchSize int) (int, error) {
	_ = "STUB: not implemented"
	// Iterate over legacy receipt keys under prefix types.ReceiptKeyPrefix
	return 0, nil
}

// Early exit if nothing to migrate

// tx hash bytes (without prefix)
// serialized receipt bytes

// Derive tx hash directly from key suffix

// Save the suffix for deletion from legacy store after successful write

// Write to transient receipt store first; they'll be flushed to receipt.db at pre-commit

// After a successful write, delete from legacy store

func (k *Keeper) WriteReceipt(
	ctx sdk.Context,
	stateDB *state.DBImpl,
	msg *core.Message,
	txType uint32,
	txHash common.Hash,
	gasUsed uint64,
	vmError string,
) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint:gosec
//nolint:gosec

// append precompile error to VM error
