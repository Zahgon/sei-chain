package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// RawContractMessage defines a json message that is sent or returned by a wasm contract.
// This type can hold any type of bytes. Until validateBasic is called there should not be
// any assumptions made that the data is valid syntax or semantic.
type RawContractMessage []byte

func (r RawContractMessage) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RawContractMessage) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r *RawContractMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Bytes returns raw bytes type
func (r RawContractMessage) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgStoreCode) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgStoreCode) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgStoreCode) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgStoreCode) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgStoreCode) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// should never happen as valid basic rejects invalid addresses

func (msg MsgInstantiateContract) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgInstantiateContract) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgInstantiateContract) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgInstantiateContract) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgInstantiateContract) GetSigners() []sdk.AccAddress {
	_ = "STUB: not implemented"
	return nil
}

// should never happen as valid basic rejects invalid addresses

func (msg MsgExecuteContract) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgExecuteContract) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgExecuteContract) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgExecuteContract) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgExecuteContract) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// should never happen as valid basic rejects invalid addresses

func (msg MsgMigrateContract) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgMigrateContract) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgMigrateContract) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgMigrateContract) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgMigrateContract) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// should never happen as valid basic rejects invalid addresses

func (msg MsgUpdateAdmin) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgUpdateAdmin) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgUpdateAdmin) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgUpdateAdmin) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgUpdateAdmin) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// should never happen as valid basic rejects invalid addresses

func (msg MsgClearAdmin) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgClearAdmin) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgClearAdmin) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgClearAdmin) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgClearAdmin) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// should never happen as valid basic rejects invalid addresses

func (msg MsgIBCSend) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgIBCSend) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgIBCSend) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgIBCSend) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgIBCSend) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg MsgIBCCloseChannel) Route() string { _ = "STUB: not implemented"; return "" }

func (msg MsgIBCCloseChannel) Type() string { _ = "STUB: not implemented"; return "" }

func (msg MsgIBCCloseChannel) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg MsgIBCCloseChannel) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg MsgIBCCloseChannel) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }
