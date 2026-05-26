package types

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const TypeMsgSend = "evm_send"

var (
	_ sdk.Msg = &MsgSend{}
)

func NewMsgSend(fromAddr sdk.AccAddress, toAddress common.Address, amount sdk.Coins) *MsgSend {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgSend) Route() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgSend) Type() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgSend) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgSend) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg *MsgSend) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
