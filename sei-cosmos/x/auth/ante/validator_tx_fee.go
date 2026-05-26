package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
)

var BaseDenomGasPriceAmplfier = sdk.NewInt(1_000_000_000_000)

// checkTxFeeWithValidatorMinGasPrices implements the default fee logic, where the minimum price per
// unit of gas is fixed and set by each validator, can the tx priority is computed from the gas price.
func CheckTxFeeWithValidatorMinGasPrices(ctx sdk.Context, tx sdk.Tx, simulate bool, paramsKeeper paramskeeper.Keeper) (sdk.Coins, int64, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), 0, nil
}

// Ensure that the provided fees meet a minimum threshold for the validator,
// if this is a CheckTx. This is only for local mempool purposes, and thus
// is only ran on check tx.

// Determine the required fees by multiplying each required minimum gas
// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
//nolint:gosec // bounds checked above

// this is the lowest priority, and will be used specifically if gas limit is set to 0
// realistically, if the gas limit IS set to 0, the tx will run out of gas anyways.

//nolint:gosec // bounds checked above

func GetMinimumGasPricesWantedSorted(globalMinimumGasPrices, validatorMinimumGasPrices sdk.DecCoins) sdk.DecCoins {
	_ = "STUB: not implemented"
	return *new(sdk.DecCoins)
}

// GetTxPriority returns a naive tx priority based on the amount of the smallest denomination of the gas price
// provided in a transaction.
// If base denom is used as fee, the calculated gas price will be amplified by 10^12 to capture higher precision
// in priority differences.
// NOTE: This implementation should be used with a great consideration as it opens potential attack vectors
// where txs with multiple coins could not be prioritize as expected.
func GetTxPriority(fee sdk.Coins, gas int64) int64 { _ = "STUB: not implemented"; return 0 }
