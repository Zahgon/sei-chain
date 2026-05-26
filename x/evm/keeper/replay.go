package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) VerifyBalance(ctx sdk.Context, addr common.Address) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) VerifyTxResult(ctx sdk.Context, hash common.Hash) {
	_ = "STUB: not implemented"
	return
}

// it's okay if remote also doesn't have receipt

//nolint:gosec

func (k *Keeper) VerifyAccount(ctx sdk.Context, addr common.Address, accountData ethtypes.Account) {
	_ = "STUB: not implemented"
	// we no longer check eth balance due to limiting EVM max refund to 150% of used gas (https://github.com/sei-protocol/go-ethereum/pull/32)
	return
}

func contains(slice []common.Address, element common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (k *Keeper) VerifyState(ctx sdk.Context, addr common.Address) {
	_ = "STUB: not implemented"
	return
}
