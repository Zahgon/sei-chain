package utils

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("utils")

const HardFailPrefix = "hard fail error occurred"

func PanicHandler(recoverCallback func(any)) func() { _ = "STUB: not implemented"; return nil }

// LogPanicCallback returns a callback function, given a context and a recovered
// error value, that logs the error and a stack trace.
func LogPanicCallback(ctx sdk.Context, r any) func(any) { _ = "STUB: not implemented"; return nil }

func DecorateHardFailError(err error) error { _ = "STUB: not implemented"; return nil }

func shouldErrorHardFail(err string) bool {
	_ = "STUB: not implemented"
	// use Contains instead of HasPrefix in case the error is further decorated
	return false
}
