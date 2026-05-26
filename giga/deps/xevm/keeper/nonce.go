package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) GetNonce(ctx sdk.Context, addr common.Address) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (k *Keeper) SetNonce(ctx sdk.Context, addr common.Address, nonce uint64) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) IterateAllNonces(ctx sdk.Context, cb func(addr common.Address, nonce uint64) bool) {
	_ = "STUB: not implemented"
	return
}
