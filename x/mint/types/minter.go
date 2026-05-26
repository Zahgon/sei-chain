package types

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	epochTypes "github.com/sei-protocol/sei-chain/x/epoch/types"
)

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(
	startDate string,
	endDate string,
	denom string,
	totalMintAmount uint64,
) Minter {
	_ = "STUB: not implemented"
	return *new(Minter)
}

// InitialMinter returns an initial Minter object with default values with no previous mints
func InitialMinter() Minter { _ = "STUB: not implemented"; return *new(Minter) }

// DefaultInitialMinter returns a default initial Minter object for a new chain
// which uses an inflation rate of 0%.
func DefaultInitialMinter() Minter {
	_ = "STUB: not implemented"
	return *

	// validate minter
	new(Minter)
}

func ValidateMinter(minter Minter) error { _ = "STUB: not implemented"; return nil }

func (m *Minter) GetLastMintDateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// This should not happen as the date is validated when the minter is created

func (m *Minter) GetStartDateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// This should not happen as the date is validated when the minter is created

func (m *Minter) GetEndDateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// This should not happen as the date is validated when the minter is created

func (m Minter) GetLastMintAmountCoin() sdk.Coin { _ = "STUB: not implemented"; return *new(sdk.Coin) }

//nolint:gosec

func (m *Minter) GetReleaseAmountToday(currentTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

//nolint:gosec

func (m *Minter) RecordSuccessfulMint(ctx sdk.Context, epoch epochTypes.Epoch, mintedAmount uint64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

//nolint:gosec
// TODO(PLT-336): remove once mint_coins_minted verified

func (m *Minter) getReleaseAmountToday(currentTime time.Time) uint64 {
	_ = "STUB: not implemented"
	// Not yet started or already minted today
	return 0
}

// if it's already past the end date then release the remaining amount likely caused by outage

func (m *Minter) GetNumberOfDaysLeft(currentTime time.Time) uint64 {
	_ = "STUB: not implemented"
	// If the last mint date is after the start date then use the last mint date as there's an ongoing release
	return 0
}

func (m *Minter) OngoingRelease() bool { _ = "STUB: not implemented"; return false }

func DaysBetween(a, b time.Time) uint64 {
	_ = "STUB: not implemented"
	// Convert both times to UTC before comparing
	return 0
}

// Always return a positive value between the dates
