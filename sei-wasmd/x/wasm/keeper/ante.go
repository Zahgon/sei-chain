package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// CountTXDecorator ante handler to count the tx position in a block.
type CountTXDecorator struct {
	storeKey sdk.StoreKey
}

// NewCountTXDecorator constructor
func NewCountTXDecorator(storeKey sdk.StoreKey) *CountTXDecorator {
	_ = "STUB: not implemented"
	return nil
}

// AnteHandle handler stores a tx counter with current height encoded in the store to let the app handle
// global rollback behavior instead of keeping state in the handler itself.
// The ante handler passes the counter value via sdk.Context upstream. See `types.TXCounter(ctx)` to read the value.
// Simulations don't get a tx counter value assigned.
func (a CountTXDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// start with 0
// load counter when exists

// then use stored counter

// else use `0` from above to start with

// store next counter value for current height

func encodeHeightCounter(height int64, counter uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G115 -- height is checked above to be non-negative

func decodeHeightCounter(bz []byte) (int64, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// LimitSimulationGasDecorator ante decorator to limit gas in simulation calls
type LimitSimulationGasDecorator struct {
	gasLimit *sdk.Gas
	// inputs: simulate, ctx, gas limit, tx
	gasMeterSetter func(bool, sdk.Context, uint64, sdk.Tx) sdk.Context
}

// NewLimitSimulationGasDecorator constructor accepts nil value to fallback to block gas limit.
func NewLimitSimulationGasDecorator(gasLimit *sdk.Gas, gasMeterSetter func(bool, sdk.Context, uint64, sdk.Tx) sdk.Context) *LimitSimulationGasDecorator {
	_ = "STUB: not implemented"
	return nil
}

func DefaultGasMeterSetter() func(bool, sdk.Context, uint64, sdk.Tx) sdk.Context {
	_ = "STUB: not implemented"
	return nil
}

// AnteHandle that limits the maximum gas available in simulations only.
// A custom max value can be configured and will be applied when set. The value should not
// exceed the max block gas limit.
// Different values on nodes are not consensus breaking as they affect only
// simulations but may have effect on client user experience.
//
// When no custom value is set then the max block gas is used as default limit.
func (d LimitSimulationGasDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"

	// Wasm code is not executed in checkTX so that we don't need to limit it further.
	// Tendermint rejects the TX afterwards when the tx.gas > max block gas.
	// On deliverTX we rely on the tendermint/sdk mechanics that ensure
	// tx has gas set and gas < max block gas
	return *new(sdk.Context), nil
}

// apply custom node gas limit

// default to max block gas when set, to be on the safe side
