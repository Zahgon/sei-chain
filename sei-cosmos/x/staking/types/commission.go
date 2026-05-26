package types

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// NewCommissionRates returns an initialized validator commission rates.
func NewCommissionRates(rate, maxRate, maxChangeRate sdk.Dec) CommissionRates {
	_ = "STUB: not implemented"
	return *new(CommissionRates)
}

// NewCommission returns an initialized validator commission.
func NewCommission(rate, maxRate, maxChangeRate sdk.Dec) Commission {
	_ = "STUB: not implemented"
	return *new(Commission)
}

// NewCommissionWithTime returns an initialized validator commission with a specified
// update time which should be the current block BFT time.
func NewCommissionWithTime(rate, maxRate, maxChangeRate sdk.Dec, updatedAt time.Time) Commission {
	_ = "STUB: not implemented"
	return *new(Commission)
}

// String implements the Stringer interface for a Commission object.
func (c Commission) String() string { _ = "STUB: not implemented"; return "" }

// String implements the Stringer interface for a CommissionRates object.
func (cr CommissionRates) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs basic sanity validation checks of initial commission
// parameters. If validation fails, an SDK error is returned.
func (cr CommissionRates) Validate() error { _ = "STUB: not implemented"; return nil }

// max rate cannot be negative

// max rate cannot be greater than 1

// rate cannot be negative

// rate cannot be greater than the max rate

// change rate cannot be negative

// change rate cannot be greater than the max rate

// ValidateNewRate performs basic sanity validation checks of a new commission
// rate. If validation fails, an SDK error is returned.
func (c Commission) ValidateNewRate(newRate sdk.Dec, blockTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// new rate cannot be changed more than once within 24 hours

// new rate cannot be negative

// new rate cannot be greater than the max rate

// new rate % points change cannot be greater than the max change rate
