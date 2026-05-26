package verify

import (
	"testing"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	"github.com/sei-protocol/sei-chain/testutil/processblock"
)

// Check balance changes as result of executing the provided transactions.
// Only works if all transactions are successful.
func Balance(t *testing.T, app *processblock.App, f BlockRunnable, txs []signing.Tx) BlockRunnable {
	_ = "STUB: not implemented"
	return *new(BlockRunnable)
}

// denom -> (account -> delta)

// TODO: add coverage for other balance-affecting messages to enable testing for those message types

func updateMultipleExpectedBalanceChange(changes map[string]map[string]int64, account string, coins sdk.Coins, positive bool) {
	_ = "STUB: not implemented"
	return
}

func updateExpectedBalanceChange(changes map[string]map[string]int64, account string, coin sdk.Coin, positive bool) {
	_ = "STUB: not implemented"
	return
}
