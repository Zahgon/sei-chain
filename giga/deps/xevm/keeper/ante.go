package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (k *Keeper) AddAnteSurplus(ctx sdk.Context, txHash common.Hash, surplus sdk.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) GetAnteSurplusSum(ctx sdk.Context) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}
