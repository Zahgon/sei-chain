package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
)

// UnmarshalDenomTrace attempts to decode and return an DenomTrace object from
// raw encoded bytes.
func (k Keeper) UnmarshalDenomTrace(bz []byte) (types.DenomTrace, error) {
	_ = "STUB: not implemented"
	return *new(types.DenomTrace), nil
}

// MustUnmarshalDenomTrace attempts to decode and return an DenomTrace object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalDenomTrace(bz []byte) types.DenomTrace {
	_ = "STUB: not implemented"
	return *new(types.DenomTrace)
}

// MarshalDenomTrace attempts to encode an DenomTrace object and returns the
// raw encoded bytes.
func (k Keeper) MarshalDenomTrace(denomTrace types.DenomTrace) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustMarshalDenomTrace attempts to encode an DenomTrace object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalDenomTrace(denomTrace types.DenomTrace) []byte {
	_ = "STUB: not implemented"
	return nil
}
