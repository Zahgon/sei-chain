package slashing

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

type SlashingWriteInfo struct {
	ConsAddr    sdk.ConsAddress
	MissedInfo  types.ValidatorMissedBlockArray
	SigningInfo types.ValidatorSigningInfo
	ShouldSlash bool
	SlashInfo   keeper.SlashInfo
}

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, votes []abci.VoteInfo, k keeper.Keeper) {
	_ = "STUB: not implemented"
	return
}

// Iterate over all the validators which *should* have signed this block
// store whether or not they have actually signed it and slash/unbond any
// which have missed too many blocks in a row (downtime slashing)

// this allows us to preserve the original ordering for writing purposes

// Update the validator missed block bit array by index if different from last value at the index
