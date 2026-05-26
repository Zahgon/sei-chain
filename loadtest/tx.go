package main

import (
	"context"

	typestx "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
)

func SendTx(
	ctx context.Context,
	txBytes []byte,
	mode typestx.BroadcastMode,
	loadtestClient LoadTestClient,
) bool {
	_ = "STUB: not implemented"
	return false
}
