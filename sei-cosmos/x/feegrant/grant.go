package feegrant

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	_ types.UnpackInterfacesMessage = &Grant{}
)

// NewGrant creates a new FeeAllowanceGrant.
func NewGrant(granter, grantee sdk.AccAddress, feeAllowance FeeAllowanceI) (Grant, error) {
	_ = "STUB: not implemented"
	return *new(Grant), nil
}

// ValidateBasic performs basic validation on
// FeeAllowanceGrant
func (a Grant) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetGrant unpacks allowance
func (a Grant) GetGrant() (FeeAllowanceI, error) {
	_ = "STUB: not implemented"
	return *new(FeeAllowanceI), nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (a Grant) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
