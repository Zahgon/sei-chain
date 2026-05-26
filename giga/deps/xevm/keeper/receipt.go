package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/ethereum/go-ethereum/core"
	"github.com/sei-protocol/sei-chain/giga/deps/xevm/state"
	"github.com/sei-protocol/sei-chain/giga/deps/xevm/types"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
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

// Only fall back to legacy store for a not found error.

// try legacy store for older receipts

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
func (k *Keeper) MockReceipt(ctx sdk.Context, txHash common.Hash, receipt *types.Receipt) error {
	_ = "STUB: not implemented"
	return nil
}

func convertReceipt(r *evmtypes.Receipt) *types.Receipt { _ = "STUB: not implemented"; return nil }

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
