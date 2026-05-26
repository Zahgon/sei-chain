package fuzzing

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func FuzzDec(i int64, isNil bool) sdk.Dec { _ = "STUB: not implemented"; return *new(sdk.Dec) }

func FuzzCoin(denom string, isNil bool, i int64) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}
