package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

var (
	_ authz.Authorization = &SendAuthorization{}
)

// NewSendAuthorization creates a new SendAuthorization object.
func NewSendAuthorization(spendLimit sdk.Coins) *SendAuthorization {
	_ = "STUB: not implemented"
	return nil
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a SendAuthorization) MsgTypeURL() string { _ = "STUB: not implemented"; return "" }

// Accept implements Authorization.Accept.
func (a SendAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	_ = "STUB: not implemented"
	return *new(authz.AcceptResponse), nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a SendAuthorization) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
