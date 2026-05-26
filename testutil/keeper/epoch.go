package keeper

import (
	"testing"

	"github.com/sei-protocol/sei-chain/app"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/epoch/keeper"
)

func TestApp(t *testing.T) *app.App { _ = "STUB: not implemented"; return nil }

func EpochKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	_ = "STUB: not implemented"
	return nil, *new(sdk.Context)
}

// Initialize params
