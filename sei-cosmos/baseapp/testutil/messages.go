package testutil

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func RegisterInterfaces(registry types.InterfaceRegistry) { _ = "STUB: not implemented"; return }

var _ sdk.Msg = &MsgCounter{}

func (msg *MsgCounter) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }
func (msg *MsgCounter) ValidateBasic() error         { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgCounter2{}

func (msg *MsgCounter2) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }
func (msg *MsgCounter2) ValidateBasic() error         { _ = "STUB: not implemented"; return nil }

var _ sdk.Msg = &MsgKeyValue{}

func (msg *MsgKeyValue) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (msg *MsgKeyValue) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
