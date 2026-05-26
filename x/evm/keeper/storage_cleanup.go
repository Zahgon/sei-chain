package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const ZeroStorageCleanupBatchSize = 100

func (k *Keeper) GetZeroStorageCleanupCheckpoint(ctx sdk.Context) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) setZeroStorageCleanupCheckpoint(ctx sdk.Context, key []byte) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) PruneZeroStorageSlots(ctx sdk.Context, limit int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// No keys were iterated, so the saved checkpoint already points to the end
// of the iterator. Clear it so the next run restarts from the beginning
// rather than resuming from an exhausted position.

// TODO(PLT-330): remove once evm_zero_storage_processed_keys_total verified
//nolint:gosec

// TODO(PLT-330): remove once evm_zero_storage_pruned_keys_total verified
// TODO(PLT-330): remove once evm_zero_storage_pruned_bytes_total verified
//nolint:gosec
//nolint:gosec

func isZeroStorageValue(val []byte) bool { _ = "STUB: not implemented"; return false }
