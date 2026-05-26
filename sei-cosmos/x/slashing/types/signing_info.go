package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewValidatorSigningInfo creates a new ValidatorSigningInfo instance
func NewValidatorSigningInfo(
	condAddr sdk.ConsAddress, startHeight, indexOffset int64,
	jailedUntil time.Time, tombstoned bool, missedBlocksCounter int64,
) ValidatorSigningInfo {
	_ = "STUB: not implemented"
	return *new(ValidatorSigningInfo)
}

// String implements the stringer interface for ValidatorSigningInfo
func (i ValidatorSigningInfo) String() string { _ = "STUB: not implemented"; return "" }

// unmarshal a validator signing info from a store value
func UnmarshalValSigningInfo(cdc codec.Codec, value []byte) (signingInfo ValidatorSigningInfo, err error) {
	_ = "STUB: not implemented"
	return *new(ValidatorSigningInfo), nil
}
