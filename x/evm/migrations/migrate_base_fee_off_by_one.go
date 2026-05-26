package migrations

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

func MigrateBaseFeeOffByOne(ctx sdk.Context, k *keeper.Keeper) error {
	_ = "STUB: not implemented"
	return nil
}
