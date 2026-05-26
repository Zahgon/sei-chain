package authz

import (
	"time"

	cdctypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
)

var (
	_ sdk.Msg = &MsgGrant{}
	_ sdk.Msg = &MsgRevoke{}
	_ sdk.Msg = &MsgExec{}

	// For amino support.
	_ legacytx.LegacyMsg = &MsgGrant{}
	_ legacytx.LegacyMsg = &MsgRevoke{}
	_ legacytx.LegacyMsg = &MsgExec{}

	_ cdctypes.UnpackInterfacesMessage = &MsgGrant{}
	_ cdctypes.UnpackInterfacesMessage = &MsgExec{}
)

// NewMsgGrant creates a new MsgGrant
func NewMsgGrant(granter sdk.AccAddress, grantee sdk.AccAddress, a Authorization, expiration time.Time) (*MsgGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSigners implements Msg
func (msg MsgGrant) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements Msg
func (msg MsgGrant) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Type implements the LegacyMsg.Type method.
func (msg MsgGrant) Type() string { _ = "STUB: not implemented"; return "" }

// Route implements the LegacyMsg.Route method.
func (msg MsgGrant) Route() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgGrant) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetAuthorization returns the cache value from the MsgGrant.Authorization if present.
func (msg *MsgGrant) GetAuthorization() Authorization {
	_ = "STUB: not implemented"
	return *new(Authorization)
}

// SetAuthorization converts Authorization to any and adds it to MsgGrant.Authorization.
func (msg *MsgGrant) SetAuthorization(a Authorization) error { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgExec) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgGrant) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgRevoke creates a new MsgRevoke
func NewMsgRevoke(granter sdk.AccAddress, grantee sdk.AccAddress, msgTypeURL string) MsgRevoke {
	_ = "STUB: not implemented"
	return *new(MsgRevoke)
}

// GetSigners implements Msg
func (msg MsgRevoke) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements MsgRequest.ValidateBasic
func (msg MsgRevoke) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Type implements the LegacyMsg.Type method.
func (msg MsgRevoke) Type() string { _ = "STUB: not implemented"; return "" }

// Route implements the LegacyMsg.Route method.
func (msg MsgRevoke) Route() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgRevoke) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// NewMsgExec creates a new MsgExecAuthorized
func NewMsgExec(grantee sdk.AccAddress, msgs []sdk.Msg) MsgExec {
	_ = "STUB: not implemented"
	return *new(MsgExec)
}

// GetMessages returns the cache values from the MsgExecAuthorized.Msgs if present.
func (msg MsgExec) GetMessages() ([]sdk.Msg, error) { _ = "STUB: not implemented"; return nil, nil }

// GetSigners implements Msg
func (msg MsgExec) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements Msg
func (msg MsgExec) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Type implements the LegacyMsg.Type method.
func (msg MsgExec) Type() string { _ = "STUB: not implemented"; return "" }

// Route implements the LegacyMsg.Route method.
func (msg MsgExec) Route() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgExec) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }
