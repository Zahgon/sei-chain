package types

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Periods stores all vesting periods passed as part of a PeriodicVestingAccount
type Periods []Period

// Duration is converts the period Length from seconds to a time.Duration
func (p Period) Duration() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// String implements the fmt.Stringer interface
func (p Period) String() string { _ = "STUB: not implemented"; return "" }

// TotalLength return the total length in seconds for a period
func (p Periods) TotalLength() int64 { _ = "STUB: not implemented"; return 0 }

// TotalDuration returns the total duration of the period
func (p Periods) TotalDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// TotalDuration returns the sum of coins for the period
func (p Periods) TotalAmount() sdk.Coins { _ = "STUB: not implemented"; return *new(sdk.Coins) }

// String implements the fmt.Stringer interface
func (p Periods) String() string { _ = "STUB: not implemented"; return "" }
