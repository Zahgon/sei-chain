package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) SetAddressMapping(ctx sdk.Context, seiAddress sdk.AccAddress, evmAddress common.Address) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) DeleteAddressMapping(ctx sdk.Context, seiAddress sdk.AccAddress, evmAddress common.Address) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) GetEVMAddress(ctx sdk.Context, seiAddress sdk.AccAddress) (common.Address, bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), false
}

func (k *Keeper) GetEVMAddressOrDefault(ctx sdk.Context, seiAddress sdk.AccAddress) common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (k *Keeper) GetSeiAddress(ctx sdk.Context, evmAddress common.Address) (sdk.AccAddress, bool) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), false
}

func (k *Keeper) GetSeiAddressOrDefault(ctx sdk.Context, evmAddress common.Address) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (k *Keeper) IterateSeiAddressMapping(ctx sdk.Context, cb func(evmAddr common.Address, seiAddr sdk.AccAddress) bool) {
	_ = "STUB: not implemented"
	return
}

// A sdk.AccAddress may not receive funds from bank if it's the result of direct-casting
// from an EVM address AND the originating EVM address has already been associated with
// a true (i.e. derived from the same pubkey) sdk.AccAddress.
func (k *Keeper) CanAddressReceive(ctx sdk.Context, addr sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

// casting goes both directions since both address formats have 20 bytes

// if the associated address is the cast address itself, allow the address to receive (e.g. EVM contract addresses)
// this means it's either a cast address that's not associated yet, or not a cast address at all.

type EvmAddressHandler struct {
	evmKeeper *Keeper
}

func NewEvmAddressHandler(evmKeeper *Keeper) EvmAddressHandler {
	_ = "STUB: not implemented"
	return *new(EvmAddressHandler)
}

func (h EvmAddressHandler) GetSeiAddressFromString(ctx sdk.Context, address string) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}
