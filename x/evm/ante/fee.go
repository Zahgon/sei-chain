package ante

import (
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/types/ethtx"
)

type EVMFeeCheckDecorator struct {
	evmKeeper     *evmkeeper.Keeper
	upgradeKeeper *upgradekeeper.Keeper
}

func NewEVMFeeCheckDecorator(evmKeeper *evmkeeper.Keeper, upgradeKeeper *upgradekeeper.Keeper) *EVMFeeCheckDecorator {
	_ = "STUB: not implemented"
	return nil
}

func (fc EVMFeeCheckDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// if EVM version is Cancun or later, and the transaction contains at least one blob, we need to
// make sure the transaction carries a non-zero blob fee cap.

// For now we are simply assuming excessive blob gas is 0. In the future we might change it to be
// dynamic based on prior block usage.
// nolint:gosec

// check if the sender has enough balance to cover fees

// run stateless checks before charging gas (mimicking Geth behavior)

// we don't want to run nonce check here for CheckTx because we have special
// logic for pending nonce during CheckTx in sig.go

// calculate the priority by dividing the total fee with the native gas limit (i.e. the effective native gas price)

// minimum fee per gas required for a tx to be processed
func (fc EVMFeeCheckDecorator) getBaseFee(ctx sdk.Context) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// lowest allowed fee per gas, base fee will not be lower than this
func (fc EVMFeeCheckDecorator) getMinimumFee(ctx sdk.Context) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// CalculatePriority returns a priority based on the effective gas price of the transaction
func (fc EVMFeeCheckDecorator) CalculatePriority(ctx sdk.Context, txData ethtx.TxData) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// TODO(PLT-330): remove once evm_effective_gas_price verified

// effectiveGasPriceHistogramSample converts wei-per-gas to float64 for OTel without calling
// Uint64() on values larger than uint64 (undefined in math/big). Clamps +Inf so histograms
// stay finite.
func effectiveGasPriceHistogramSample(gp *big.Int) float64 { _ = "STUB: not implemented"; return 0 }
