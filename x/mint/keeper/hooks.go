package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	epochTypes "github.com/sei-protocol/sei-chain/x/epoch/types"
)

func (k Keeper) BeforeEpochStart(_ sdk.Context, _ epochTypes.Epoch) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epoch epochTypes.Epoch) {
	_ = "STUB: not implemented"
	return
}

// mint coins, update supply

// send the minted coins to the fee collector account

// Released Succssfully, decrement the remaining amount by the daily release amount and update minter

type Hooks struct {
	k Keeper
}

var _ epochTypes.EpochHooks = Hooks{}

// Return the wrapper struct.
func (k Keeper) Hooks() Hooks {
	_ = "STUB: not implemented"

	// epochs hooks.
	return *new(Hooks)
}

func (h Hooks) BeforeEpochStart(ctx sdk.Context, epoch epochTypes.Epoch) {
	_ = "STUB: not implemented"
	return
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epoch epochTypes.Epoch) {
	_ = "STUB: not implemented"
	return
}
