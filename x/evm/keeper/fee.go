package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// modified eip-1559 adjustment using target gas used
func (k *Keeper) AdjustDynamicBaseFeePerGas(ctx sdk.Context, blockGasUsed uint64) *sdk.Dec {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec
// avoid division by zero
// return the previous base fee as is

//nolint:gosec

// cap block gas used to block gas limit

// upward adjustment

// downward adjustment

// Ensure the new base fee is not lower than the minimum fee

// Ensure the new base fee is not higher than the maximum fee

// Set the new base fee for the next height

// NOTE: this is only used in migrate_base_fee_off_by_one migration. This is deprecated.
// dont have height be a prefix, just store the current base fee directly
func (k *Keeper) GetCurrBaseFeePerGas(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// NOTE: this is only used in migrate_base_fee_off_by_one migration. This is deprecated.
func (k *Keeper) SetCurrBaseFeePerGas(ctx sdk.Context, baseFeePerGas sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) SetNextBaseFeePerGas(ctx sdk.Context, baseFeePerGas sdk.Dec) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) GetNextBaseFeePerGas(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}
