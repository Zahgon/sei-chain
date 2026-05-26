package state

import (
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// UseiToSweiMultiplier Fields that were denominated in usei will be converted to swei (1usei = 10^12swei)
// for existing Ethereum application (which assumes 18 decimal points) to display properly.
var UseiToSweiMultiplier = big.NewInt(1_000_000_000_000)
var SdkUseiToSweiMultiplier = sdk.NewIntFromBigInt(UseiToSweiMultiplier)

var CoinbaseAddressPrefix = []byte("evm_coinbase")

func GetCoinbaseAddress(txIdx int) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

//nolint:gosec

func SplitUseiWeiAmount(amt *big.Int) (sdk.Int, sdk.Int) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), *new(sdk.Int)
}
