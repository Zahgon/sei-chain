package types

import (
	"sync"
)

// Gas consumption descriptors.
const (
	GasIterNextCostFlatDesc = "IterNextFlat"
	GasValuePerByteDesc     = "ValuePerByte"
	GasWritePerByteDesc     = "WritePerByte"
	GasReadPerByteDesc      = "ReadPerByte"
	GasWriteCostFlatDesc    = "WriteFlat"
	GasReadCostFlatDesc     = "ReadFlat"
	GasHasDesc              = "Has"
	GasDeleteDesc           = "Delete"
)

// Gas measured by the SDK
type Gas = uint64

// ErrorNegativeGasConsumed defines an error thrown when the amount of gas refunded results in a
// negative gas consumed amount.
type ErrorNegativeGasConsumed struct {
	Descriptor string
}

// ErrorOutOfGas defines an error thrown when an action results in out of gas.
type ErrorOutOfGas struct {
	Descriptor string
}

// ErrorGasOverflow defines an error thrown when an action results gas consumption
// unsigned integer overflow.
type ErrorGasOverflow struct {
	Descriptor string
}

// GasMeter interface to track gas consumption
type GasMeter interface {
	GasConsumed() Gas
	GasConsumedToLimit() Gas
	Limit() Gas
	ConsumeGas(amount Gas, descriptor string)
	RefundGas(amount Gas, descriptor string)
	IsPastLimit() bool
	IsOutOfGas() bool
	String() string
	Multiplier() (numerator uint64, denominator uint64)
}

type basicGasMeter struct {
	limit    Gas
	consumed Gas
	lock     *sync.Mutex
}

func (g *basicGasMeter) GasConsumed() Gas { _ = "STUB: not implemented"; return *new(Gas) }

func (g *basicGasMeter) Limit() Gas { _ = "STUB: not implemented"; return *new(Gas) }

func (g *basicGasMeter) GasConsumedToLimit() Gas { _ = "STUB: not implemented"; return *new(Gas) }

// addUint64Overflow performs the addition operation on two uint64 integers and
// returns a boolean on whether or not the result overflows.
func addUint64Overflow(a, b uint64) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (g *basicGasMeter) ConsumeGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

// cosmos_tx_gas_exceeded
func (g *basicGasMeter) incrGasExceededCounter(errorType string, descriptor string) {
	_ = "STUB: not implemented"
	return
}

// descriptor is a label to distinguish between different gas meters (e.g block vs tx)

// RefundGas will deduct the given amount from the gas consumed. If the amount is greater than the
// gas consumed, the function will panic.
//
// Use case: This functionality enables refunding gas to the transaction or block gas pools so that
// EVM-compatible chains can fully support the go-ethereum StateDb interface.
// See https://github.com/cosmos/cosmos-sdk/pull/9403 for reference.
func (g *basicGasMeter) RefundGas(amount Gas, descriptor string) { _ = "STUB: not implemented"; return }

func (g *basicGasMeter) IsPastLimit() bool { _ = "STUB: not implemented"; return false }

func (g *basicGasMeter) IsOutOfGas() bool { _ = "STUB: not implemented"; return false }

func (g *basicGasMeter) String() string { _ = "STUB: not implemented"; return "" }

func (g *basicGasMeter) Multiplier() (numerator uint64, denominator uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

type multiplierGasMeter struct {
	basicGasMeter
	multiplierNumerator   uint64
	multiplierDenominator uint64
}

func NewMultiplierGasMeter(limit Gas, multiplierNumerator uint64, multiplierDenominator uint64) GasMeter {
	_ = "STUB: not implemented"
	return *new(GasMeter)
}

func (g *multiplierGasMeter) adjustGas(original Gas) Gas {
	_ = "STUB: not implemented"
	return *new(Gas)
}

func (g *multiplierGasMeter) ConsumeGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

func (g *multiplierGasMeter) RefundGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

func (g *multiplierGasMeter) Multiplier() (numerator uint64, denominator uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

type infiniteGasMeter struct {
	consumed Gas
	lock     *sync.Mutex
}

func (g *infiniteGasMeter) GasConsumed() Gas { _ = "STUB: not implemented"; return *new(Gas) }

func (g *infiniteGasMeter) GasConsumedToLimit() Gas { _ = "STUB: not implemented"; return *new(Gas) }

func (g *infiniteGasMeter) Limit() Gas { _ = "STUB: not implemented"; return *new(Gas) }

func (g *infiniteGasMeter) ConsumeGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

// TODO: Should we set the consumed field after overflow checking?

// RefundGas will deduct the given amount from the gas consumed. If the amount is greater than the
// gas consumed, the function will panic.
//
// Use case: This functionality enables refunding gas to the trasaction or block gas pools so that
// EVM-compatible chains can fully support the go-ethereum StateDb interface.
// See https://github.com/cosmos/cosmos-sdk/pull/9403 for reference.
func (g *infiniteGasMeter) RefundGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

func (g *infiniteGasMeter) IsPastLimit() bool { _ = "STUB: not implemented"; return false }

func (g *infiniteGasMeter) IsOutOfGas() bool { _ = "STUB: not implemented"; return false }

func (g *infiniteGasMeter) String() string { _ = "STUB: not implemented"; return "" }

func (g *infiniteGasMeter) Multiplier() (numerator uint64, denominator uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

type infiniteMultiplierGasMeter struct {
	infiniteGasMeter
	multiplierNumerator   uint64
	multiplierDenominator uint64
}

func NewInfiniteMultiplierGasMeter(multiplierNumerator uint64, multiplierDenominator uint64) GasMeter {
	_ = "STUB: not implemented"
	return *new(GasMeter)
}

func (g *infiniteMultiplierGasMeter) adjustGas(original Gas) Gas {
	_ = "STUB: not implemented"
	return *new(Gas)
}

func (g *infiniteMultiplierGasMeter) ConsumeGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

func (g *infiniteMultiplierGasMeter) RefundGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

func (g *infiniteMultiplierGasMeter) Multiplier() (numerator uint64, denominator uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

type noConsumptionInfiniteGasMeter struct {
	infiniteGasMeter
}

func NewNoConsumptionInfiniteGasMeter() GasMeter { _ = "STUB: not implemented"; return *new(GasMeter) }

func (g *noConsumptionInfiniteGasMeter) GasConsumed() Gas {
	_ = "STUB: not implemented"
	return *new(Gas)
}

func (g *noConsumptionInfiniteGasMeter) GasConsumedToLimit() Gas {
	_ = "STUB: not implemented"
	return *new(Gas)
}

func (g *noConsumptionInfiniteGasMeter) ConsumeGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"
	return
}

func (g *noConsumptionInfiniteGasMeter) RefundGas(amount Gas, descriptor string) {
	_ = "STUB: not implemented"

	// GasConfig defines gas cost for each operation on KVStores
	return
}

type GasConfig struct {
	HasCost          Gas
	DeleteCost       Gas
	ReadCostFlat     Gas
	ReadCostPerByte  Gas
	WriteCostFlat    Gas
	WriteCostPerByte Gas
	IterNextCostFlat Gas
}

// KVGasConfig returns a default gas config for KVStores.
func KVGasConfig() GasConfig { _ = "STUB: not implemented"; return *new(GasConfig) }

// TransientGasConfig returns a default gas config for TransientStores.
func TransientGasConfig() GasConfig { _ = "STUB: not implemented"; return *new(GasConfig) }
