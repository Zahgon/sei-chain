package feegrant

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var _ FeeAllowanceI = (*PeriodicAllowance)(nil)

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
func (a *PeriodicAllowance) Accept(ctx sdk.Context, fee sdk.Coins, _ []sdk.Msg) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// deduct from both the current period and the max amount

// tryResetPeriod will check if the PeriodReset has been hit. If not, it is a no-op.
// If we hit the reset period, it will top up the PeriodCanSpend amount to
// min(PeriodSpendLimit, Basic.SpendLimit) so it is never more than the maximum allowed.
// It will also update the PeriodReset. If we are within one Period, it will update from the
// last PeriodReset (eg. if you always do one tx per day, it will always reset the same time)
// If we are more then one period out (eg. no activity in a week), reset is one Period from the execution of this method
func (a *PeriodicAllowance) tryResetPeriod(blockTime time.Time) { _ = "STUB: not implemented"; return }

// set PeriodCanSpend to the lesser of Basic.SpendLimit and PeriodSpendLimit

// If we are within the period, step from expiration (eg. if you always do one tx per day, it will always reset the same time)
// If we are more then one period out (eg. no activity in a week), reset is one period from this time

// ValidateBasic implements FeeAllowance and enforces basic sanity checks
func (a PeriodicAllowance) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// We allow 0 for CanSpend

// ensure PeriodSpendLimit can be subtracted from total (same coin types)

// check times
