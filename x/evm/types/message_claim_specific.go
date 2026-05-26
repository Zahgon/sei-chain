package types

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/utils"
)

const TypeMsgClaimSpecific = "evm_claim_specific"

var (
	_ sdk.Msg = &MsgClaimSpecific{}
)

func NewMsgClaimSpecific(sender sdk.AccAddress, claimer common.Address, assets ...*Asset) *MsgClaimSpecific {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgClaimSpecific) Route() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgClaimSpecific) Type() string { _ = "STUB: not implemented"; return "" }

func (msg *MsgClaimSpecific) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgClaimSpecific) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

func (msg *MsgClaimSpecific) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (msg *MsgClaimSpecific) GetIAssets() (res []utils.IAsset) {
	_ = "STUB: not implemented"
	return nil
}

func (a *Asset) IsCW20() bool { _ = "STUB: not implemented"; return false }

func (a *Asset) IsCW721() bool { _ = "STUB: not implemented"; return false }

func (a *Asset) IsNative() bool { _ = "STUB: not implemented"; return false }
