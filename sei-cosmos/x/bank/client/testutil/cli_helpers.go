package testutil

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil"
)

func MsgSendExec(clientCtx client.Context, from, to, amount fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func QueryBalancesExec(clientCtx client.Context, address fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}
