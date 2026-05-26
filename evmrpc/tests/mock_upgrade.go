package tests

import (
	"github.com/sei-protocol/sei-chain/app"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func mockUpgrade(version string, height int64) func(ctx sdk.Context, a *app.App) {
	_ = "STUB: not implemented"
	return nil
}
