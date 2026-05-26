package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
)

// TxFeeChecker check if the provided fee is enough and returns the effective fee and tx priority,
// the effective fee should be deducted later, and the priority should be returned in abci response.
type TxFeeChecker func(ctx sdk.Context, tx sdk.Tx, simulate bool, paramsKeeper paramskeeper.Keeper) (sdk.Coins, int64, error)

// DeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
type DeductFeeDecorator struct {
	accountKeeper  AccountKeeper
	bankKeeper     types.BankKeeper
	feegrantKeeper FeegrantKeeper
	paramsKeeper   paramskeeper.Keeper
	txFeeChecker   TxFeeChecker
}

func NewDeductFeeDecorator(
	ak AccountKeeper,
	bk types.BankKeeper,
	fk FeegrantKeeper,
	paramsKeeper paramskeeper.Keeper,
	tfc TxFeeChecker,
) DeductFeeDecorator {
	_ = "STUB: not implemented"
	return *new(DeductFeeDecorator)
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

func (dfd DeductFeeDecorator) checkDeductFee(ctx sdk.Context, sdkTx sdk.Tx, fee sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// if feegranter set deduct fee from feegranter account.
// this works with only when feegrant enabled.

// deduct the fees

// DeductFees deducts fees from the given account.
func DeductFees(bankKeeper types.BankKeeper, ctx sdk.Context, acc types.AccountI, fees sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}
