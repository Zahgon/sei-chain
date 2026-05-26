package verify

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	"github.com/sei-protocol/sei-chain/testutil/processblock"
)

func Epoch(t *testing.T, app *processblock.App, f BlockRunnable, _ []signing.Tx) BlockRunnable {
	_ = "STUB: not implemented"
	return *new(BlockRunnable)
}
