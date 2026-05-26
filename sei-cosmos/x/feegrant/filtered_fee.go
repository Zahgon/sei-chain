package feegrant

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// TODO: Revisit this once we have propoer gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const (
	gasCostPerIteration = uint64(10)
)

var _ FeeAllowanceI = (*AllowedMsgAllowance)(nil)
var _ types.UnpackInterfacesMessage = (*AllowedMsgAllowance)(nil)

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (a *AllowedMsgAllowance) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewAllowedMsgFeeAllowance creates new filtered fee allowance.
func NewAllowedMsgAllowance(allowance FeeAllowanceI, allowedMsgs []string) (*AllowedMsgAllowance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAllowance returns allowed fee allowance.
func (a *AllowedMsgAllowance) GetAllowance() (FeeAllowanceI, error) {
	_ = "STUB: not implemented"
	return *new(FeeAllowanceI), nil
}

// SetAllowance sets allowed fee allowance.
func (a *AllowedMsgAllowance) SetAllowance(allowance FeeAllowanceI) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept method checks for the filtered messages has valid expiry
func (a *AllowedMsgAllowance) Accept(ctx sdk.Context, fee sdk.Coins, msgs []sdk.Msg) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a *AllowedMsgAllowance) allowedMsgsToMap(ctx sdk.Context) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func (a *AllowedMsgAllowance) allMsgTypesAllowed(ctx sdk.Context, msgs []sdk.Msg) bool {
	_ = "STUB: not implemented"
	return false
}

// ValidateBasic implements FeeAllowance and enforces basic sanity checks
func (a *AllowedMsgAllowance) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
