package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
)

var (
	_ GasTx = (*legacytx.StdTx)(nil) // assert StdTx implements GasTx
)

// GasTx defines a Tx with a GetGas() method which is needed to use SetUpContextDecorator
type GasTx interface {
	sdk.Tx
	GetGas() uint64
}

// SetUpContextDecorator sets the GasMeter in the Context and wraps the next AnteHandler with a defer clause
// to recover from any downstream OutOfGas panics in the AnteHandler chain to return an error with information
// on gas provided and gas used.
// CONTRACT: Must be first decorator in the chain
// CONTRACT: Tx must implement GasTx interface
type SetUpContextDecorator struct {
	gasMeterSetter func(bool, sdk.Context, uint64, sdk.Tx) sdk.Context
}

func NewDefaultSetUpContextDecorator() SetUpContextDecorator {
	_ = "STUB: not implemented"
	return *new(SetUpContextDecorator)
}

func NewSetUpContextDecorator(gasMeterSetter func(bool, sdk.Context, uint64, sdk.Tx) sdk.Context) SetUpContextDecorator {
	_ = "STUB: not implemented"
	return *new(SetUpContextDecorator)
}

func (sud SetUpContextDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_ = "STUB: not implemented"
	// all transactions must implement GasTx
	return *new(sdk.Context), nil
}

// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
// during runTx.

// If there exists a maximum block gas limit, we must ensure that the tx
// does not exceed it.
//nolint:gosec // MaxGas is validated positive by the condition

// Decorator will catch an OutOfGasPanic caused in the next antehandler
// AnteHandlers must have their own defer/recover in order for the BaseApp
// to know how much gas was used! This is because the GasMeter is created in
// the AnteHandler, but if it panics the context won't be set properly in
// runTx's recover call.

// SetGasMeter returns a new context with a gas meter set from a given context.
func SetGasMeter(simulate bool, ctx sdk.Context, gasLimit uint64, _ sdk.Tx) sdk.Context {
	_ = "STUB: not implemented"
	// In various cases such as simulation and during the genesis block, we do not
	// meter any gas utilization.
	return *new(sdk.Context)
}
