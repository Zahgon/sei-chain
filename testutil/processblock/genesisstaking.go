package processblock

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (a *App) NewValidator() sdk.ValAddress { _ = "STUB: not implemented"; return *new(sdk.ValAddress) }

func (a *App) NewDelegation(delegator sdk.AccAddress, validator sdk.ValAddress, amount int64) {
	_ = "STUB: not implemented"
	return
}
