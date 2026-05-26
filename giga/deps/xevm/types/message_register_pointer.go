package types

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const TypeMsgRegisterPointer = "evm_register_pointer"

var (
	_ sdk.Msg = &MsgRegisterPointer{}
)

func NewMsgRegisterERC20Pointer(sender sdk.AccAddress, ercAddress common.Address) *MsgRegisterPointer {
	_ = "STUB: not implemented"
	return nil
}

func NewMsgRegisterERC721Pointer(sender sdk.AccAddress, ercAddress common.Address) *MsgRegisterPointer {
	_ = "STUB: not implemented"
	return nil
}

func NewMsgRegisterERC1155Pointer(sender sdk.AccAddress, ercAddress common.Address) *MsgRegisterPointer {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgRegisterPointer) Route() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgRegisterPointer) Type() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgRegisterPointer) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgRegisterPointer) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg *MsgRegisterPointer) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
