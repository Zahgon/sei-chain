package vesting

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/vesting/types"
)

// NewHandler returns a handler for x/auth message types.
func NewHandler(ak keeper.AccountKeeper, bk types.BankKeeper) sdk.Handler {
	_ = "STUB: not implemented"
	return *new(sdk.Handler)
}
