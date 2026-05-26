package testutil

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

func MsgWithdrawDelegatorRewardExec(clientCtx client.Context, valAddr fmt.Stringer, extraArgs ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
