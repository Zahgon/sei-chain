package legacyabci

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// RecoveryHandler handles recovery() object.
// Return a non-nil error if recoveryObj was processed.
// Return nil if recoveryObj was not processed.
type RecoveryHandler func(recoveryObj interface{}) error

// recoveryMiddleware is wrapper for RecoveryHandler to create chained recovery handling.
// returns (recoveryMiddleware, nil) if recoveryObj was not processed and should be passed to the next middleware in chain.
// returns (nil, error) if recoveryObj was processed and middleware chain processing should be stopped.
type recoveryMiddleware func(recoveryObj interface{}) (recoveryMiddleware, error)

// processRecovery processes recoveryMiddleware chain for recovery() object.
// Chain processing stops on non-nil error or when chain is processed.
func processRecovery(recoveryObj interface{}, middleware recoveryMiddleware) error {
	_ = "STUB: not implemented"
	return nil
}

// newRecoveryMiddleware creates a RecoveryHandler middleware.
func newRecoveryMiddleware(handler RecoveryHandler, next recoveryMiddleware) recoveryMiddleware {
	_ = "STUB: not implemented"
	return *new(recoveryMiddleware)
}

// newOutOfGasRecoveryMiddleware creates a standard OutOfGas recovery middleware for app.runTx method.
func newOutOfGasRecoveryMiddleware(gasWanted uint64, ctx sdk.Context, next recoveryMiddleware) recoveryMiddleware {
	_ = "STUB: not implemented"
	return *new(recoveryMiddleware)
}

// newOCCAbortRecoveryMiddleware creates a standard OCC Abort recovery middleware for app.runTx method.
func newOCCAbortRecoveryMiddleware(next recoveryMiddleware) recoveryMiddleware {
	_ = "STUB: not implemented"
	return *new(recoveryMiddleware)
}

// newDefaultRecoveryMiddleware creates a default (last in chain) recovery middleware for app.runTx method.
func newDefaultRecoveryMiddleware() recoveryMiddleware {
	_ = "STUB: not implemented"
	return *new(recoveryMiddleware)
}
