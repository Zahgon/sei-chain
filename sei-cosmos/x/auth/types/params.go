package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Default parameter values
const (
	DefaultMaxMemoCharacters      uint64 = 256
	DefaultTxSigLimit             uint64 = 7
	DefaultTxSizeCostPerByte      uint64 = 10
	DefaultSigVerifyCostED25519   uint64 = 590
	DefaultSigVerifyCostSecp256k1 uint64 = 1000
)

// Parameter keys
var (
	KeyMaxMemoCharacters      = []byte("MaxMemoCharacters")
	KeyTxSigLimit             = []byte("TxSigLimit")
	KeyTxSizeCostPerByte      = []byte("TxSizeCostPerByte")
	KeySigVerifyCostED25519   = []byte("SigVerifyCostED25519")
	KeySigVerifyCostSecp256k1 = []byte("SigVerifyCostSecp256k1")
	KeyDisableSeqnoCheck      = []byte("KeyDisableSeqnoCheck")
)

var _ paramtypes.ParamSet = &Params{}

// NewParams creates a new Params object
func NewParams(
	maxMemoCharacters, txSigLimit, txSizeCostPerByte, sigVerifyCostED25519, sigVerifyCostSecp256k1 uint64,
) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// ParamKeyTable for auth module
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// SigVerifyCostSecp256r1 returns gas fee of secp256r1 signature verification.
// Set by benchmarking current implementation:
//
//	BenchmarkSig/secp256k1     4334   277167 ns/op   4128 B/op   79 allocs/op
//	BenchmarkSig/secp256r1    10000   108769 ns/op   1672 B/op   33 allocs/op
//
// Based on the results above secp256k1 is 2.7x is slwer. However we propose to discount it
// because we are we don't compare the cgo implementation of secp256k1, which is faster.
func (p Params) SigVerifyCostSecp256r1() uint64 { _ = "STUB: not implemented"; return 0 }

// String implements the stringer interface.
func (p Params) String() string { _ = "STUB: not implemented"; return "" }

func validateTxSigLimit(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateSigVerifyCostED25519(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateSigVerifyCostSecp256k1(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMaxMemoCharacters(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateTxSizeCostPerByte(i interface{}) error { _ = "STUB: not implemented"; return nil }

// Validate checks that the parameters have valid values.
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

func (p *Params) GetSr25519VerifyCost() uint64 {
	_ = "STUB: not implemented"
	// TODO:: define param for sr25519 once its confirmed that it will be supported
	return 0
}
