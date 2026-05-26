package processblock

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

type Preset struct {
	Admin            sdk.AccAddress // a signable account that's not supposed to run out of tokens
	SignableAccounts []sdk.AccAddress
	AllAccounts      []sdk.AccAddress
	AllValidators    []sdk.ValAddress
}

// 3 unsignable accounts
// 3 bonded validators
func CommonPreset(app *App) *Preset { _ = "STUB: not implemented"; return nil }

// always with enough fee
func (p *Preset) AdminSign(app *App, msgs ...sdk.Msg) signing.Tx {
	_ = "STUB: not implemented"
	return *new(signing.Tx)
}
