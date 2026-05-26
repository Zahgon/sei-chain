package authz

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	_ Authorization = &GenericAuthorization{}
)

// NewGenericAuthorization creates a new GenericAuthorization object.
func NewGenericAuthorization(msgTypeURL string) *GenericAuthorization {
	_ = "STUB: not implemented"
	return nil
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a GenericAuthorization) MsgTypeURL() string {
	_ = "STUB: not implemented"

	// Accept implements Authorization.Accept.
	return ""
}

func (a GenericAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (AcceptResponse, error) {
	_ = "STUB: not implemented"
	return *new(AcceptResponse), nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a GenericAuthorization) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
