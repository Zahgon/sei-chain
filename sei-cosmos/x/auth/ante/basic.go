package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// ValidateBasicDecorator will call tx.ValidateBasic and return any non-nil error.
// If ValidateBasic passes, decorator calls next AnteHandler in chain. Note,
// ValidateBasicDecorator decorator will not get executed on ReCheckTx since it
// is not dependent on application state.
type ValidateBasicDecorator struct{}

func NewValidateBasicDecorator() ValidateBasicDecorator {
	_ = "STUB: not implemented"
	return *new(ValidateBasicDecorator)
}

func (vbd ValidateBasicDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	// no need to validate basic on recheck tx, call next antehandler
	return *new(sdk.Context), nil
}

// ValidateMemoDecorator will validate memo given the parameters passed in
// If memo is too large decorator returns with error, otherwise call next AnteHandler
// CONTRACT: Tx must implement TxWithMemo interface
type ValidateMemoDecorator struct {
	ak AccountKeeper
}

func NewValidateMemoDecorator(ak AccountKeeper) ValidateMemoDecorator {
	_ = "STUB: not implemented"
	return *new(ValidateMemoDecorator)
}

func (vmd ValidateMemoDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// ConsumeTxSizeGasDecorator will take in parameters and consume gas proportional
// to the size of tx before calling next AnteHandler. Note, the gas costs will be
// slightly over estimated due to the fact that any given signing account may need
// to be retrieved from state.
//
// CONTRACT: If simulate=true, then signatures must either be completely filled
// in or empty.
// CONTRACT: To use this decorator, signatures of transaction must be represented
// as legacytx.StdSignature otherwise simulate mode will incorrectly estimate gas cost.
type ConsumeTxSizeGasDecorator struct {
	ak AccountKeeper
}

func NewConsumeGasForTxSizeDecorator(ak AccountKeeper) ConsumeTxSizeGasDecorator {
	_ = "STUB: not implemented"
	return *new(ConsumeTxSizeGasDecorator)
}

func (cgts ConsumeTxSizeGasDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

//nolint:gosec // len() is always non-negative

// simulate gas cost for signatures in simulate mode

// in simulate mode, each element should be a nil signature

// if signature is already filled in, no need to simulate gas cost

// use placeholder simSecp256k1Pubkey if sig is nil

// use stdsignature to mock the size of a full signature
//nolint:staticcheck // this will be removed when proto is ready

//nolint:gosec // len() + 6 is always non-negative and small

// If the pubkey is a multi-signature pubkey, then we estimate for the maximum
// number of signers.

// isIncompleteSignature tests whether SignatureData is fully filled in for simulation purposes
func isIncompleteSignature(data signing.SignatureData) bool {
	_ = "STUB: not implemented"
	return false
}

type (
	// TxTimeoutHeightDecorator defines an AnteHandler decorator that checks for a
	// tx height timeout.
	TxTimeoutHeightDecorator struct{}

	// TxWithTimeoutHeight defines the interface a tx must implement in order for
	// TxHeightTimeoutDecorator to process the tx.
	TxWithTimeoutHeight interface {
		sdk.Tx

		GetTimeoutHeight() uint64
	}
)

// TxTimeoutHeightDecorator defines an AnteHandler decorator that checks for a
// tx height timeout.
func NewTxTimeoutHeightDecorator() TxTimeoutHeightDecorator {
	_ = "STUB: not implemented"
	return *new(TxTimeoutHeightDecorator)
}

// AnteHandle implements an AnteHandler decorator for the TxHeightTimeoutDecorator
// type where the current block height is checked against the tx's height timeout.
// If a height timeout is provided (non-zero) and is less than the current block
// height, then an error is returned.
func (txh TxTimeoutHeightDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

//nolint:gosec // block height is always non-negative
