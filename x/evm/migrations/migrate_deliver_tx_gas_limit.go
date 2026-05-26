package migrations

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

func MigrateDeliverTxHookWasmGasLimitParam(ctx sdk.Context, k *keeper.Keeper) error {
	_ = "STUB: not implemented"
	// Fetch the v11 parameters
	return nil
}

// Add DeliverTxHookWasmGasLimit to with default value

// Set the updated parameters back in the keeper
