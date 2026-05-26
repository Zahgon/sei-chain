package types

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const TypeMsgClaim = "evm_claim"

var (
	_ sdk.Msg = &MsgClaim{}
)

func NewMsgClaim(sender sdk.AccAddress, claimer common.Address) *MsgClaim {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgClaim) Route() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgClaim) Type() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgClaim) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgClaim) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg *MsgClaim) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
