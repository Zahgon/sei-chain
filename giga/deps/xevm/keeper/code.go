package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) GetCode(ctx sdk.Context, addr common.Address) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) SetCode(ctx sdk.Context, addr common.Address, code []byte) {
	_ = "STUB: not implemented"
	return
}

// set association with direct cast Sei address for the contract address

func (k *Keeper) GetCodeHash(ctx sdk.Context, addr common.Address) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

// per Ethereum behavior, if an address has no code, balance, or nonce, return Hash(0)

// if an address has no code but has balance or nonce, return EmptyCodeHash

func (k *Keeper) GetCodeSize(ctx sdk.Context, addr common.Address) int {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec

func (k *Keeper) IterateAllCode(ctx sdk.Context, cb func(addr common.Address, code []byte) bool) {
	_ = "STUB: not implemented"
	return
}
