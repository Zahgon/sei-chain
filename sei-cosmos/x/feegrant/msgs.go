package feegrant

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
)

var (
	_, _ sdk.Msg            = &MsgGrantAllowance{}, &MsgRevokeAllowance{}
	_, _ legacytx.LegacyMsg = &MsgGrantAllowance{}, &MsgRevokeAllowance{} // For amino support.

	_ types.UnpackInterfacesMessage = &MsgGrantAllowance{}
)

// NewMsgGrantAllowance creates a new MsgGrantAllowance.
func NewMsgGrantAllowance(feeAllowance FeeAllowanceI, granter, grantee sdk.AccAddress) (*MsgGrantAllowance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgGrantAllowance) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners gets the granter account associated with an allowance
func (msg MsgGrantAllowance) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// Type implements the LegacyMsg.Type method.
func (msg MsgGrantAllowance) Type() string { _ = "STUB: not implemented"; return "" }

// Route implements the LegacyMsg.Route method.
func (msg MsgGrantAllowance) Route() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgGrantAllowance) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetFeeAllowanceI returns unpacked FeeAllowance
func (msg MsgGrantAllowance) GetFeeAllowanceI() (FeeAllowanceI, error) {
	_ = "STUB: not implemented"
	return *new(FeeAllowanceI), nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgGrantAllowance) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgRevokeAllowance returns a message to revoke a fee allowance for a given
// granter and grantee
func NewMsgRevokeAllowance(granter sdk.AccAddress, grantee sdk.AccAddress) MsgRevokeAllowance {
	_ = "STUB: not implemented"
	return *new(MsgRevokeAllowance)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgRevokeAllowance) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSigners gets the granter address associated with an Allowance
// to revoke.
func (msg MsgRevokeAllowance) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// Type implements the LegacyMsg.Type method.
func (msg MsgRevokeAllowance) Type() string { _ = "STUB: not implemented"; return "" }

// Route implements the LegacyMsg.Route method.
func (msg MsgRevokeAllowance) Route() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (msg MsgRevokeAllowance) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }
