package keeper

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/evm/types"
)

type EVMCallFunc func(caller common.Address, addr *common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)

var MaxUint64BigInt = new(big.Int).SetUint64(math.MaxUint64)

func (k *Keeper) HandleInternalEVMCall(ctx sdk.Context, req *types.MsgInternalEVMCall) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Keeper) HandleInternalEVMDelegateCall(ctx sdk.Context, req *types.MsgInternalEVMDelegateCall) (*sdk.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// delegatecall caller must be associated; otherwise any state change on EVM contract will be lost
// after they asssociate.

// TODO(PLT-330): remove once evm_association_error_total verified

func (k *Keeper) CallEVM(ctx sdk.Context, from common.Address, to *common.Address, val *sdk.Int, data []byte) (retdata []byte, reterr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This call was not part of an existing StateTransition, so it should trigger one

// replay attack is prevented by the AccountSequence number set on the CW transaction that triggered this call

// fees are already paid on the CW transaction

// should not increment nonce since this isn't a transaction

//nolint:gosec

func (k *Keeper) StaticCallEVM(ctx sdk.Context, from sdk.AccAddress, to *common.Address, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Keeper) callEVM(ctx sdk.Context, from common.Address, to *common.Address, val *sdk.Int, data []byte, f EVMCallFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only used for StaticCalls
func (k *Keeper) createReadOnlyEVM(ctx sdk.Context, from sdk.AccAddress) (*vm.EVM, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Keeper) getEvmGasLimitFromCtx(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

func (k *Keeper) consumeEvmGas(ctx sdk.Context, usedEvmGas uint64) {
	_ = "STUB: not implemented"
	return
}
