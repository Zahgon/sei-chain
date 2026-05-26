package keeper

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
)

type SlashInfo struct {
	height             int64
	power              int64
	distributionHeight int64
	minHeight          int64
	minSignedPerWindow int64
}

// This performs similar logic to the above HandleValidatorSignature, but only performs READs such that it can be performed in parallel for all validators.
// Instead of updating appropriate validator bit arrays / signing infos, this will return the pending values to be written in a consistent order
func (k Keeper) HandleValidatorSignatureConcurrent(ctx sdk.Context, addr cryptotypes.Address, power int64, signed bool) (consAddr sdk.ConsAddress, missedInfo types.ValidatorMissedBlockArray, signInfo types.ValidatorSigningInfo, shouldSlash bool, slashInfo SlashInfo) {
	_ = "STUB: not implemented"
	return *new(sdk.ConsAddress), *new(types.ValidatorMissedBlockArray), *new(types.ValidatorSigningInfo), false, *new(SlashInfo)
}

// fetch the validator public key

// fetch signing info

// Array value has changed from not missed to missed, increment counter

// Array value has changed from missed to not missed, decrement counter

// Array value at this index has not changed, no need to update counter

// bump index offset (and mod circular) after performing potential resizing

// if we are past the minimum height and the validator has missed too many blocks, punish them

// Downtime confirmed: slash and jail the validator
// We need to retrieve the stake distribution which signed the block, so we subtract ValidatorUpdateDelay from the evidence height,
// and subtract an additional 1 since this is the LastCommit.
// Note that this *can* result in a negative "distributionHeight" up to -ValidatorUpdateDelay-1,
// i.e. at the end of the pre-genesis block (none) = at the beginning of the genesis block.
// That's fine since this is just used to filter unbonding delegations & redelegations.

// This value is passed back and the validator is slashed and jailed appropriately

// validator was (a) not found or (b) already jailed so we do not slash

func (k Keeper) SlashJailAndUpdateSigningInfo(ctx sdk.Context, consAddr sdk.ConsAddress, slashInfo SlashInfo, signInfo types.ValidatorSigningInfo) types.ValidatorSigningInfo {
	_ = "STUB: not implemented"
	return *new(types.ValidatorSigningInfo)
}

// Slashed for missing too many block

func (k Keeper) ResizeMissedBlockArray(missedInfo types.ValidatorMissedBlockArray, signInfo types.ValidatorSigningInfo, window int64, index int64) (types.ValidatorMissedBlockArray, types.ValidatorSigningInfo, int64) {
	_ = "STUB: not implemented"
	// we need to resize the missed block array AND update the signing info accordingly
	return *new(types.ValidatorMissedBlockArray), *new(types.ValidatorSigningInfo), 0
}

// missed block array too short, lets expand it

// insert `0`s corresponding to the difference between the new window size and old window size

// if window size is reduced, we would like to make a clean state so that no validators are unexpectedly jailed due to more recent missed blocks
