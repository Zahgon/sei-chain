package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/utils"
)

func (k *Keeper) RunWithOneOffEVMInstance(
	ctx sdk.Context, runner func(*vm.EVM) error, logger func(string, string),
) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) UpsertERCNativePointer(
	ctx sdk.Context, evm *vm.EVM, token string, metadata utils.ERCMetadata,
) (contractAddr common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (k *Keeper) UpsertERCCW20Pointer(
	ctx sdk.Context, evm *vm.EVM, cw20Addr string, metadata utils.ERCMetadata,
) (contractAddr common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (k *Keeper) UpsertERCCW721Pointer(
	ctx sdk.Context, evm *vm.EVM, cw721Addr string, metadata utils.ERCMetadata,
) (contractAddr common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (k *Keeper) UpsertERCCW1155Pointer(
	ctx sdk.Context, evm *vm.EVM, cw1155Addr string, metadata utils.ERCMetadata,
) (contractAddr common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (k *Keeper) UpsertERCPointer(
	ctx sdk.Context, evm *vm.EVM, typ string, args []interface{}, getter PointerGetter, setter PointerSetter,
) (contractAddr common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}
