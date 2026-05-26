package processblock

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (a *App) FundAccount(acc sdk.AccAddress, amount int64) { _ = "STUB: not implemented"; return }

func (a *App) FundModule(moduleName string, amount int64) { _ = "STUB: not implemented"; return }

func (a *App) FundAccountWithDenom(acc sdk.AccAddress, amount int64, denom string) {
	_ = "STUB: not implemented"
	return
}

func (a *App) FundModuleWithDenom(moduleName string, amount int64, denom string) {
	_ = "STUB: not implemented"
	return
}
