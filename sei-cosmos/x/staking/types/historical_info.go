package types

import (
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewHistoricalInfo will create a historical information struct from header and valset
// it will first sort valset before inclusion into historical info
func NewHistoricalInfo(header tmproto.Header, valSet Validators, powerReduction sdk.Int) HistoricalInfo {
	_ = "STUB: not implemented"
	// Must sort in the same way that tendermint does
	return *new(HistoricalInfo)
}

// MustUnmarshalHistoricalInfo wll unmarshal historical info and panic on error
func MustUnmarshalHistoricalInfo(cdc codec.BinaryCodec, value []byte) HistoricalInfo {
	_ = "STUB: not implemented"
	return *new(HistoricalInfo)
}

// UnmarshalHistoricalInfo will unmarshal historical info and return any error
func UnmarshalHistoricalInfo(cdc codec.BinaryCodec, value []byte) (hi HistoricalInfo, err error) {
	_ = "STUB: not implemented"
	return *new(HistoricalInfo), nil
}

// ValidateBasic will ensure HistoricalInfo is not nil and sorted
func ValidateBasic(hi HistoricalInfo) error { _ = "STUB: not implemented"; return nil }

// Equal checks if receiver is equal to the parameter
func (hi *HistoricalInfo) Equal(hi2 *HistoricalInfo) bool { _ = "STUB: not implemented"; return false }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (hi HistoricalInfo) UnpackInterfaces(c codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
