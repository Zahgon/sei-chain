package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"

	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
)

var logger = seilog.NewLogger("x", "evm", "ante")

type EVMSigVerifyDecorator struct {
	evmKeeper *evmkeeper.Keeper
}

func NewEVMSigVerifyDecorator(evmKeeper *evmkeeper.Keeper, _ func() sdk.Context) *EVMSigVerifyDecorator {
	_ = "STUB: not implemented"
	return nil
}

func (svd *EVMSigVerifyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// set EVM properties

// validate chain ID on the transaction

// legacy either can have a zero or correct chain ID

// after legacy, all transactions must have the correct chain ID

// TODO(PLT-330): remove once evm_nonce_mismatch_total verified
