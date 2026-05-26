package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("x", "tokenfactory", "keeper")

func (k Keeper) mintTo(ctx sdk.Context, amount sdk.Coin, mintTo string) error {
	_ = "STUB: not implemented"
	// verify that denom is an x/tokenfactory denom
	return nil
}

func (k Keeper) burnFrom(ctx sdk.Context, amount sdk.Coin, burnFrom string) error {
	_ = "STUB: not implemented"
	// verify that denom is an x/tokenfactory denom
	return nil
}
