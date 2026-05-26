package migrations

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

func StoreCWPointerCode(ctx sdk.Context, k *keeper.Keeper, store20 bool, store721 bool, store1155 bool) error {
	_ = "STUB: not implemented"
	return nil
}
