package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
)

// constants
const (
	TypeMsgCreateDenom      = "create_denom"
	TypeMsgUpdateDenom      = "update_denom"
	TypeMsgMint             = "mint"
	TypeMsgBurn             = "burn"
	TypeMsgChangeAdmin      = "change_admin"
	TypeMsgSetDenomMetadata = "set_denom_metadata"
)

var _ sdk.Msg = &MsgCreateDenom{}

// NewMsgCreateDenom creates a msg to create a new denom
func NewMsgCreateDenom(sender, subdenom string) *MsgCreateDenom {
	_ = "STUB: not implemented"
	return nil
}

func (m MsgCreateDenom) Route() string        { _ = "STUB: not implemented"; return "" }
func (m MsgCreateDenom) Type() string         { _ = "STUB: not implemented"; return "" }
func (m MsgCreateDenom) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (m MsgCreateDenom) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (m MsgCreateDenom) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgUpdateDenom{}

// NewMsgUpdateDenom creates a msg to update denom
func NewMsgUpdateDenom(sender, denom string, allowList *banktypes.AllowList) *MsgUpdateDenom {
	_ = "STUB: not implemented"
	return nil
}

func (m MsgUpdateDenom) Route() string        { _ = "STUB: not implemented"; return "" }
func (m MsgUpdateDenom) Type() string         { _ = "STUB: not implemented"; return "" }
func (m MsgUpdateDenom) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (m MsgUpdateDenom) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (m MsgUpdateDenom) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgMint{}

// NewMsgMint creates a message to mint tokens
func NewMsgMint(sender string, amount sdk.Coin) *MsgMint { _ = "STUB: not implemented"; return nil }

func (m MsgMint) Route() string        { _ = "STUB: not implemented"; return "" }
func (m MsgMint) Type() string         { _ = "STUB: not implemented"; return "" }
func (m MsgMint) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (m MsgMint) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (m MsgMint) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgBurn{}

// NewMsgBurn creates a message to burn tokens
func NewMsgBurn(sender string, amount sdk.Coin) *MsgBurn { _ = "STUB: not implemented"; return nil }

func (m MsgBurn) Route() string        { _ = "STUB: not implemented"; return "" }
func (m MsgBurn) Type() string         { _ = "STUB: not implemented"; return "" }
func (m MsgBurn) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (m MsgBurn) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (m MsgBurn) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgChangeAdmin{}

// NewMsgChangeAdmin creates a message to change admin for a denom
func NewMsgChangeAdmin(sender, denom, newAdmin string) *MsgChangeAdmin {
	_ = "STUB: not implemented"
	return nil
}

func (m MsgChangeAdmin) Route() string        { _ = "STUB: not implemented"; return "" }
func (m MsgChangeAdmin) Type() string         { _ = "STUB: not implemented"; return "" }
func (m MsgChangeAdmin) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (m MsgChangeAdmin) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (m MsgChangeAdmin) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgSetDenomMetadata{}

// NewMsgChangeAdmin creates a message to burn tokens
func NewMsgSetDenomMetadata(sender string, metadata banktypes.Metadata) *MsgSetDenomMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (m MsgSetDenomMetadata) Route() string        { _ = "STUB: not implemented"; return "" }
func (m MsgSetDenomMetadata) Type() string         { _ = "STUB: not implemented"; return "" }
func (m MsgSetDenomMetadata) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (m MsgSetDenomMetadata) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (m MsgSetDenomMetadata) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }
