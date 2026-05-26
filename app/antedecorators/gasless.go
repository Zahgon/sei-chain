package antedecorators

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
	oracletypes "github.com/sei-protocol/sei-chain/x/oracle/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("app", "antedecorators")

type GaslessDecorator struct {
	wrapped      []sdk.AnteDecorator
	oracleKeeper oraclekeeper.Keeper
	evmKeeper    *evmkeeper.Keeper
}

func NewGaslessDecorator(wrapped []sdk.AnteDecorator, oracleKeeper oraclekeeper.Keeper, evmKeeper *evmkeeper.Keeper) GaslessDecorator {
	_ = "STUB: not implemented"
	return *new(GaslessDecorator)
}

func (gd GaslessDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// eagerly set infinite gas meter so that queries performed by IsTxGasless will not incur gas cost

// In the case of deliverTx, we want to deduct fees regardless of whether the tx is considered gasless or not, since
// gasless txs will be subject to application-specific fee requirements in later stage of ante, for which the payment
// of those app-specific fees happens here. Note that the minimum fee check in the wrapped deduct fee handler is only
// performed if the context is for CheckTx, so the check will be skipped for deliverTx and the deduct fee handler will
// only deduct fee without checking.
// Otherwise (i.e. in the case of checkTx), we only want to perform fee checks and fee deduction if the tx is not considered
// gasless, or if it specifies a non-zero gas limit even if it is considered gasless, so that the wrapped deduct fee
// handler will assign an appropriate priority to it.

func (gd GaslessDecorator) handleWrapped(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	// AnteHandle always takes a `next` so we need a no-op to execute only one handler at a time
	return *new(sdk.Context), nil
}

// iterating instead of recursing the handler for readability

// We need to replace with the new context returned by the handler otherwise we could be losing data

func IsTxGasless(tx sdk.Tx, ctx sdk.Context, oracleKeeper oraclekeeper.Keeper, evmKeeper *evmkeeper.Keeper) (isGasless bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// empty TX shouldn't be gasless

// ddos prevention

func oracleVoteIsGasless(msg *oracletypes.MsgAggregateExchangeRateVote, ctx sdk.Context, keeper oraclekeeper.Keeper) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// this returns an error IFF there is no vote present
// this also gets cleared out after every vote window, so if there is no vote present, we may want to allow gasless tx

// if there is no error that means there is a vote present, so we don't allow gasless tx

// otherwise we allow it

func evmAssociateIsGasless(msg *evmtypes.MsgAssociate, ctx sdk.Context, keeper *evmkeeper.Keeper) bool {
	_ = "STUB: not implemented"
	// not gasless if already associated
	return false
}
