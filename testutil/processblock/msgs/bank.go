package msgs

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
)

func Send(from sdk.AccAddress, to sdk.AccAddress, amount int64) *banktypes.MsgSend {
	_ = "STUB: not implemented"
	return nil
}
