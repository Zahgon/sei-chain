package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// staking message types
const (
	TypeMsgUndelegate      = "begin_unbonding"
	TypeMsgEditValidator   = "edit_validator"
	TypeMsgCreateValidator = "create_validator"
	TypeMsgDelegate        = "delegate"
	TypeMsgBeginRedelegate = "begin_redelegate"
)

var (
	_ sdk.Msg                            = &MsgCreateValidator{}
	_ codectypes.UnpackInterfacesMessage = (*MsgCreateValidator)(nil)
	_ sdk.Msg                            = &MsgCreateValidator{}
	_ sdk.Msg                            = &MsgEditValidator{}
	_ sdk.Msg                            = &MsgDelegate{}
	_ sdk.Msg                            = &MsgUndelegate{}
	_ sdk.Msg                            = &MsgBeginRedelegate{}
)

// NewMsgCreateValidator creates a new MsgCreateValidator instance.
// Delegator address and validator address are the same.
func NewMsgCreateValidator(
	valAddr sdk.ValAddress, pubKey cryptotypes.PubKey,
	selfDelegation sdk.Coin, description Description, commission CommissionRates, minSelfDelegation sdk.Int,
) (*MsgCreateValidator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Route implements the sdk.Msg interface.
func (msg MsgCreateValidator) Route() string {
	_ = "STUB: not implemented"

	// Type implements the sdk.Msg interface.
	return ""
}

func (msg MsgCreateValidator) Type() string { _ = "STUB: not implemented"; return "" }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
// If the validator address is not same as delegator's, then the validator must
// sign the msg as well.
func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	// delegator is first signer so delegator pays fees
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCreateValidator) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreateValidator) ValidateBasic() error {
	_ = "STUB: not implemented"
	// note that unmarshaling from bech32 ensures either empty or valid
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgCreateValidator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgEditValidator creates a new MsgEditValidator instance
func NewMsgEditValidator(valAddr sdk.ValAddress, description Description, newRate *sdk.Dec, newMinSelfDelegation *sdk.Int) *MsgEditValidator {
	_ = "STUB: not implemented"
	return nil
}

// Route implements the sdk.Msg interface.
func (msg MsgEditValidator) Route() string {
	_ = "STUB: not implemented"

	// Type implements the sdk.Msg interface.
	return ""
}

func (msg MsgEditValidator) Type() string { _ = "STUB: not implemented"; return "" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgEditValidator) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgEditValidator) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgEditValidator) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgDelegate creates a new MsgDelegate instance.
func NewMsgDelegate(delAddr sdk.AccAddress, valAddr sdk.ValAddress, amount sdk.Coin) *MsgDelegate {
	_ = "STUB: not implemented"
	return nil
}

// Route implements the sdk.Msg interface.
func (msg MsgDelegate) Route() string {
	_ = "STUB: not implemented"

	// Type implements the sdk.Msg interface.
	return ""
}

func (msg MsgDelegate) Type() string { _ = "STUB: not implemented"; return "" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgDelegate) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgDelegate) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgDelegate) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgBeginRedelegate creates a new MsgBeginRedelegate instance.
func NewMsgBeginRedelegate(
	delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress, amount sdk.Coin,
) *MsgBeginRedelegate {
	_ = "STUB: not implemented"
	return nil
}

// Route implements the sdk.Msg interface.
func (msg MsgBeginRedelegate) Route() string {
	_ = "STUB: not implemented"

	// Type implements the sdk.Msg interface
	return ""
}

func (msg MsgBeginRedelegate) Type() string { _ = "STUB: not implemented"; return "" }

// GetSigners implements the sdk.Msg interface
func (msg MsgBeginRedelegate) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgBeginRedelegate) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgBeginRedelegate) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgUndelegate creates a new MsgUndelegate instance.
func NewMsgUndelegate(delAddr sdk.AccAddress, valAddr sdk.ValAddress, amount sdk.Coin) *MsgUndelegate {
	_ = "STUB: not implemented"
	return nil
}

// Route implements the sdk.Msg interface.
func (msg MsgUndelegate) Route() string {
	_ = "STUB: not implemented"

	// Type implements the sdk.Msg interface.
	return ""
}

func (msg MsgUndelegate) Type() string { _ = "STUB: not implemented"; return "" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgUndelegate) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgUndelegate) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgUndelegate) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
