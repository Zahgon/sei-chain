package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) BeginBlock(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// TODO(PLT-330): remove once evm_abci_begin_blocker_duration_seconds verified

// clear tx/tx responses from last block

// mock beacon root if replaying

func (k *Keeper) EndBlock(ctx sdk.Context, height int64, blockGasUsed int64) {
	_ = "STUB: not implemented"
	return
}

// TODO(PLT-330): remove once evm_abci_end_blocker_duration_seconds verified

// Bake height-1: at EndBlock(N) the indexer's safe latest is N-1. When
// the snapshot store is wired, also Put a memiavl snapshot keyed by
// its committed version (= N-1, since Commit fires after EndBlock);
// the baker tracing block H looks up snapshot[H-1].

// TODO: remove after all TxHashes have been removed

// Migrate legacy EVM receipts to receipt.db in small batches every N blocks

// nolint:gosec

// TODO(PLT-330): remove once evm_block_base_fee verified

// nolint:gosec

// Only EVM receipts in this block that are not synthetic
//nolint:gosec

// Re-create a per-tx bloom from EVM-only logs (exclude synthetic receipts but not synthetic logs)
