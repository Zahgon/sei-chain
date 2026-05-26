package processblock

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (a *App) NewAccount() sdk.AccAddress { _ = "STUB: not implemented"; return *new(sdk.AccAddress) }

func (a *App) NewSignableAccount(name string) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}
