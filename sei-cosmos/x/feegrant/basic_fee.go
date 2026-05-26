package feegrant

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var _ FeeAllowanceI = (*BasicAllowance)(nil)

// Accept can use fee payment requested as well as timestamp of the current block
// to determine whether or not to process this. This is checked in
// Keeper.UseGrantedFees and the return values should match how it is handled there.
//
// If it returns an error, the fee payment is rejected, otherwise it is accepted.
// The FeeAllowance implementation is expected to update it's internal state
// and will be saved again after an acceptance.
//
// If remove is true (regardless of the error), the FeeAllowance will be deleted from storage
// (eg. when it is used up). (See call to RevokeAllowance in Keeper.UseGrantedFees)
func (a *BasicAllowance) Accept(ctx sdk.Context, fee sdk.Coins, _ []sdk.Msg) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ValidateBasic implements FeeAllowance and enforces basic sanity checks
func (a BasicAllowance) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
