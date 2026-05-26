package testutil

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil"
)

func TxSignExec(clientCtx client.Context, from fmt.Stringer, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func TxBroadcastExec(clientCtx client.Context, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func TxEncodeExec(clientCtx client.Context, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func TxValidateSignaturesExec(clientCtx client.Context, filename string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func TxMultiSignExec(clientCtx client.Context, from string, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func TxSignBatchExec(clientCtx client.Context, from fmt.Stringer, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func TxDecodeExec(clientCtx client.Context, encodedTx string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func QueryAccountExec(clientCtx client.Context, address fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func TxMultiSignBatchExec(clientCtx client.Context, filename string, from string, sigFile1 string, sigFile2 string, extraArgs ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

// DONTCOVER
