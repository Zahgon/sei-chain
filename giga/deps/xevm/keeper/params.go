package keeper

import (
	"math"
	"math/big"

	"github.com/sei-protocol/sei-chain/giga/deps/xevm/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const BaseDenom = "usei"

var MaxUint64BigInt = new(big.Int).SetUint64(math.MaxUint64)

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }

func (k *Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

func (k *Keeper) GetParamsPreV580(ctx sdk.Context) (params types.ParamsPreV580) {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV580)
}

func (k *Keeper) GetParamsPreV600(ctx sdk.Context) (params types.ParamsPreV600) {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV600)
}

func (k *Keeper) GetParamsPreV601(ctx sdk.Context) (params types.ParamsPreV601) {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV601)
}

func (k *Keeper) GetParamsPreV606(ctx sdk.Context) (params types.ParamsPreV606) {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV606)
}

func (k *Keeper) GetParamsIfExists(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

func (k *Keeper) GetParamsPreV580IfExists(ctx sdk.Context) types.ParamsPreV580 {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV580)
}

func (k *Keeper) GetParamsPreV600IfExists(ctx sdk.Context) types.ParamsPreV600 {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV600)
}

func (k *Keeper) GetParamsPreV601IfExists(ctx sdk.Context) types.ParamsPreV601 {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV601)
}

func (k *Keeper) GetParamsPreV606IfExists(ctx sdk.Context) types.ParamsPreV606 {
	_ = "STUB: not implemented"
	return *new(types.ParamsPreV606)
}

func (k *Keeper) GetBaseDenom(ctx sdk.Context) string { _ = "STUB: not implemented"; return "" }

func (k *Keeper) GetPriorityNormalizer(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

func (k *Keeper) GetBaseFeePerGas(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

func (k *Keeper) GetMaxDynamicBaseFeeUpwardAdjustment(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Not present in pre-6.0.0 params; use default

func (k *Keeper) GetMaxDynamicBaseFeeDownwardAdjustment(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Not present in pre-6.0.0 params; use default

func (k *Keeper) GetMinimumFeePerGas(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

func (k *Keeper) GetMaximumFeePerGas(ctx sdk.Context) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Not present in pre-6.0.1 params; use default

func (k *Keeper) GetTargetGasUsedPerBlock(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// Not present in pre-6.0.0 params; use default

func (k *Keeper) GetDeliverTxHookWasmGasLimit(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// Not present in pre-5.8.0 params; use default

func (k *Keeper) GetRegisterPointerDisabled(ctx sdk.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Not present in pre-5.8.0 params; use default

func (k *Keeper) ChainID(ctx sdk.Context) *big.Int {
	_ = "STUB: not implemented"
	// return mapped chain ID
	return nil
}

/*
*
sei gas = evm gas * multiplier
sei gas price = fee / sei gas = fee / (evm gas * multiplier) = evm gas / multiplier
*/
func (k *Keeper) GetEVMGasLimitFromCtx(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

func (k *Keeper) GetCosmosGasLimitFromEVMGas(ctx sdk.Context, evmGas uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (k *Keeper) getEvmGasLimitFromCtx(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }
