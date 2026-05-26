package authz

import (
	"time"

	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// NewGrant returns new Grant
func NewGrant( /*blockTime time.Time, */ a Authorization, expiration time.Time) (Grant, error) {
	_ = "STUB: not implemented"
	// TODO: add this for 0.45
	// if !expiration.After(blockTime) {
	// 	return Grant{}, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be after the current block time (%v), got %v", blockTime.Format(time.RFC3339), expiration.Format(time.RFC3339))
	// }
	return *new(Grant), nil
}

var (
	_ cdctypes.UnpackInterfacesMessage = &Grant{}
)

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (g Grant) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAuthorization returns the cached value from the Grant.Authorization if present.
func (g Grant) GetAuthorization() Authorization {
	_ = "STUB: not implemented"
	return *new(Authorization)
}

func (g Grant) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
