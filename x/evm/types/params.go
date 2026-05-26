package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

var (
	KeyPriorityNormalizer                  = []byte("KeyPriorityNormalizer")
	KeyMinFeePerGas                        = []byte("KeyMinFeePerGas")
	KeyMaxFeePerGas                        = []byte("KeyMaximumFeePerGas")
	KeyDeliverTxHookWasmGasLimit           = []byte("KeyDeliverTxHookWasmGasLimit")
	KeyMaxDynamicBaseFeeUpwardAdjustment   = []byte("KeyMaxDynamicBaseFeeUpwardAdjustment")
	KeyMaxDynamicBaseFeeDownwardAdjustment = []byte("KeyMaxDynamicBaseFeeDownwardAdjustment")
	KeyTargetGasUsedPerBlock               = []byte("KeyTargetGasUsedPerBlock")
	KeySeiSstoreSetGasEIP2200              = []byte("KeySeiSstoreSetGasEIP2200")
	// deprecated
	KeyBaseFeePerGas                          = []byte("KeyBaseFeePerGas")
	KeyWhitelistedCwCodeHashesForDelegateCall = []byte("KeyWhitelistedCwCodeHashesForDelegateCall")
	KeyRegisterPointerDisabled                = []byte("KeyRegisterPointerDisabled")
)

var DefaultPriorityNormalizer = sdk.NewDec(1)

// DefaultBaseFeePerGas determines how much usei per gas spent is
// burnt rather than go to validators (similar to base fee on
// Ethereum).
var DefaultBaseFeePerGas = sdk.NewDec(0)         // used for static base fee, deprecated in favor of dynamic base fee
var DefaultMinFeePerGas = sdk.NewDec(1000000000) // 1gwei
var DefaultDeliverTxHookWasmGasLimit = uint64(300000)

var DefaultWhitelistedCwCodeHashesForDelegateCall = generateDefaultWhitelistedCwCodeHashesForDelegateCall()

var DefaultMaxDynamicBaseFeeUpwardAdjustment = sdk.NewDecWithPrec(189, 4)  // 1.89%
var DefaultMaxDynamicBaseFeeDownwardAdjustment = sdk.NewDecWithPrec(39, 4) // .39%
var DefaultTargetGasUsedPerBlock = uint64(250000)                          // 250k
var DefaultMaxFeePerGas = sdk.NewDec(1000000000000)                        // 1,000gwei
var DefaultRegisterPointerDisabled = false
var DefaultSeiSstoreSetGasEIP2200 = uint64(20000) // 20k

var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func (ppre580 *ParamsPreV580) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func (ppre600 *ParamsPreV600) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func (ppre601 *ParamsPreV601) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func (ppre606 *ParamsPreV606) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

func validateBaseFeeAdjustment(i interface{}) error { _ = "STUB: not implemented"; return nil }

func (p Params) String() string { _ = "STUB: not implemented"; return "" }

func (ppre580 ParamsPreV580) String() string { _ = "STUB: not implemented"; return "" }

func (ppre600 ParamsPreV600) String() string { _ = "STUB: not implemented"; return "" }

func (ppre601 ParamsPreV601) String() string { _ = "STUB: not implemented"; return "" }

func (ppre606 ParamsPreV606) String() string { _ = "STUB: not implemented"; return "" }

func validatePriorityNormalizer(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateBaseFeePerGas(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMinFeePerGas(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMaxFeePerGas(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateDeliverTxHookWasmGasLimit(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateWhitelistedCwHashesForDelegateCall(i interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRegisterPointerDisabled(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateSeiSstoreSetGasEIP2200(i interface{}) error { _ = "STUB: not implemented"; return nil }

func generateDefaultWhitelistedCwCodeHashesForDelegateCall() [][]byte {
	_ = "STUB: not implemented"
	return nil
}
