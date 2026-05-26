package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) GetState(ctx sdk.Context, addr common.Address, hash common.Hash) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (k *Keeper) SetState(ctx sdk.Context, addr common.Address, key common.Hash, val common.Hash) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) IterateState(ctx sdk.Context, cb func(addr common.Address, key common.Hash, val common.Hash) bool) {
	_ = "STUB: not implemented"
	return
}
