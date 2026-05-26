package verify

import (
	"testing"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	"github.com/sei-protocol/sei-chain/testutil/processblock"
)

// assuming only `usei` will get distributed
func Allocation(t *testing.T, app *processblock.App, f BlockRunnable, _ []signing.Tx) BlockRunnable {
	_ = "STUB: not implemented"
	return *

	// fees collected in T-1 are allocated in T's BeginBlock, so we can simply
	// query fee collector's balance since this function is called between T-1
	// and T.
	new(BlockRunnable)
}

// in test, every val always signs

func getOutstandingRewards(app *processblock.App) map[string]sdk.DecCoin {
	_ = "STUB: not implemented"
	return nil
}
