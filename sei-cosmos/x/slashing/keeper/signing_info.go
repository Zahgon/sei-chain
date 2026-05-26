package keeper

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
)

const UINT_64_NUM_BITS = 64

// GetValidatorSigningInfo retruns the ValidatorSigningInfo for a specific validator
// ConsAddress
func (k Keeper) GetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress) (info types.ValidatorSigningInfo, found bool) {
	_ = "STUB: not implemented"
	return *new(types.ValidatorSigningInfo), false
}

// HasValidatorSigningInfo returns if a given validator has signing information
// persited.
func (k Keeper) HasValidatorSigningInfo(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
	_ = "STUB: not implemented"
	return false
}

// SetValidatorSigningInfo sets the validator signing info to a consensus address key
func (k Keeper) SetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress, info types.ValidatorSigningInfo) {
	_ = "STUB: not implemented"
	return
}

// IterateValidatorSigningInfos iterates over the stored ValidatorSigningInfo
func (k Keeper) IterateValidatorSigningInfos(ctx sdk.Context,
	handler func(address sdk.ConsAddress, info types.ValidatorSigningInfo) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// GetValidatorMissedBlockArray gets the missed blocks array
func (k Keeper) GetValidatorMissedBlocks(ctx sdk.Context, address sdk.ConsAddress) (missedInfo types.ValidatorMissedBlockArray, found bool) {
	_ = "STUB: not implemented"
	return *new(types.ValidatorMissedBlockArray), false
}

// SetValidatorMissedBlockArray sets the missed blocks array
func (k Keeper) SetValidatorMissedBlocks(ctx sdk.Context, address sdk.ConsAddress, missedInfo types.ValidatorMissedBlockArray) {
	_ = "STUB: not implemented"
	return
}

// Get a boolean representing whether a validator missed a block with a specific index offset
func (k Keeper) GetBooleanFromBitGroups(bitGroupArray []uint64, index int64) bool {
	_ = "STUB: not implemented"
	// convert the index into indexKey + indexShift
	return false
}

// shift 1 by the indexShift value to generate bit mask (to index into the bitGroup)

// apply the mask and if the value at that `indexShift` is 1 (indicating miss), then the value would be non-zero

// Set the missed value for whether a validator missed a block
func (k Keeper) SetBooleanInBitGroups(bitGroupArray []uint64, index int64, missed bool) []uint64 {
	_ = "STUB: not implemented"
	// convert the index into indexKey + indexShift
	return nil
}

// set bit to 1 by ORing with the specific position as 1

// set bit to 0 by AND NOTing with the specific position as 1

// set after updating the position

func (k Keeper) ParseBitGroupsToBoolArray(bitGroups []uint64, window int64) []bool {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) ParseBoolArrayToBitGroups(boolArray []bool) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

// JailUntil attempts to set a validator's JailedUntil attribute in its signing
// info. It will panic if the signing info does not exist for the validator.
func (k Keeper) JailUntil(ctx sdk.Context, consAddr sdk.ConsAddress, jailTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// Tombstone attempts to tombstone a validator. It will panic if signing info for
// the given validator does not exist.
func (k Keeper) Tombstone(ctx sdk.Context, consAddr sdk.ConsAddress) {
	_ = "STUB: not implemented"
	return
}

// IsTombstoned returns if a given validator by consensus address is tombstoned.
func (k Keeper) IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
	_ = "STUB: not implemented"
	return false
}

// clearValidatorMissedBlockBitArray deletes every instance of ValidatorMissedBlockBitArray in the store
func (k Keeper) ClearValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress) {
	_ = "STUB: not implemented"
	return
}
