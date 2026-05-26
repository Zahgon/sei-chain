package types

import (
	"time"

	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	vestexported "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/vesting/exported"
)

// Compile-time type assertions
var (
	_ authtypes.AccountI          = (*BaseVestingAccount)(nil)
	_ vestexported.VestingAccount = (*ContinuousVestingAccount)(nil)
	_ vestexported.VestingAccount = (*PeriodicVestingAccount)(nil)
	_ vestexported.VestingAccount = (*DelayedVestingAccount)(nil)
)

// Base Vesting Account

// NewBaseVestingAccount creates a new BaseVestingAccount object. It is the
// callers responsibility to ensure the base account has sufficient funds with
// regards to the original vesting amount.
func NewBaseVestingAccount(baseAccount *authtypes.BaseAccount, originalVesting sdk.Coins, endTime int64, admin sdk.AccAddress) *BaseVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// LockedCoinsFromVesting returns all the coins that are not spendable (i.e. locked)
// for a vesting account given the current vesting coins. If no coins are locked,
// an empty slice of Coins is returned.
//
// CONTRACT: Delegated vesting coins and vestingCoins must be sorted.
func (bva BaseVestingAccount) LockedCoinsFromVesting(vestingCoins sdk.Coins) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// TrackDelegation tracks a delegation amount for any given vesting account type
// given the amount of coins currently vesting and the current account balance
// of the delegation denominations.
//
// CONTRACT: The account's coins, delegation coins, vesting coins, and delegated
// vesting coins must be sorted.
func (bva *BaseVestingAccount) TrackDelegation(balance, vestingCoins, amount sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// Panic if the delegation amount is zero or if the base coins does not
// exceed the desired delegation amount.

// compute x and y per the specification, where:
// X := min(max(V - DV, 0), D)
// Y := D - X

// TrackUndelegation tracks an undelegation amount by setting the necessary
// values by which delegated vesting and delegated vesting need to decrease and
// by which amount the base coins need to increase.
//
// NOTE: The undelegation (bond refund) amount may exceed the delegated
// vesting (bond) amount due to the way undelegation truncates the bond refund,
// which can increase the validator's exchange rate (tokens/shares) slightly if
// the undelegated tokens are non-integral.
//
// CONTRACT: The account's coins and undelegation coins must be sorted.
func (bva *BaseVestingAccount) TrackUndelegation(amount sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// panic if the undelegation amount is zero

// compute x and y per the specification, where:
// X := min(DF, D)
// Y := min(DV, D - X)

// GetOriginalVesting returns a vesting account's original vesting amount
func (bva BaseVestingAccount) GetOriginalVesting() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// GetDelegatedFree returns a vesting account's delegation amount that is not
// vesting.
func (bva BaseVestingAccount) GetDelegatedFree() sdk.Coins {
	_ = "STUB: not implemented"
	return *

	// GetDelegatedVesting returns a vesting account's delegation amount that is
	// still vesting.
	new(sdk.Coins)
}

func (bva BaseVestingAccount) GetDelegatedVesting() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// GetEndTime returns a vesting account's end time
func (bva BaseVestingAccount) GetEndTime() int64 {
	_ = "STUB: not implemented"

	// Validate checks for errors on the account fields
	return 0
}

func (bva BaseVestingAccount) Validate() error { _ = "STUB: not implemented"; return nil }

type vestingAccountYAML struct {
	Address          sdk.AccAddress `json:"address" yaml:"address"`
	PubKey           string         `json:"public_key" yaml:"public_key"`
	AccountNumber    uint64         `json:"account_number" yaml:"account_number"`
	Sequence         uint64         `json:"sequence" yaml:"sequence"`
	OriginalVesting  sdk.Coins      `json:"original_vesting" yaml:"original_vesting"`
	DelegatedFree    sdk.Coins      `json:"delegated_free" yaml:"delegated_free"`
	DelegatedVesting sdk.Coins      `json:"delegated_vesting" yaml:"delegated_vesting"`
	EndTime          int64          `json:"end_time" yaml:"end_time"`
	Admin            string         `json:"admin" yaml:"admin"`

	// custom fields based on concrete vesting type which can be omitted
	StartTime      int64   `json:"start_time,omitempty" yaml:"start_time,omitempty"`
	VestingPeriods Periods `json:"vesting_periods,omitempty" yaml:"vesting_periods,omitempty"`
}

func (bva BaseVestingAccount) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML returns the YAML representation of a BaseVestingAccount.
func (bva BaseVestingAccount) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Continuous Vesting Account

var _ vestexported.VestingAccount = (*ContinuousVestingAccount)(nil)
var _ authtypes.GenesisAccount = (*ContinuousVestingAccount)(nil)

// NewContinuousVestingAccountRaw creates a new ContinuousVestingAccount object from BaseVestingAccount
func NewContinuousVestingAccountRaw(bva *BaseVestingAccount, startTime int64) *ContinuousVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// NewContinuousVestingAccount returns a new ContinuousVestingAccount
func NewContinuousVestingAccount(baseAcc *authtypes.BaseAccount, originalVesting sdk.Coins, startTime, endTime int64, admin sdk.AccAddress) *ContinuousVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// GetVestedCoins returns the total number of vested coins. If no coins are vested,
// nil is returned.
func (cva ContinuousVestingAccount) GetVestedCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *

	// We must handle the case where the start time for a vesting account has
	// been set into the future or when the start of the chain is not exactly
	// known.
	new(sdk.Coins)
}

// calculate the vesting scalar

// GetVestingCoins returns the total number of vesting coins. If no coins are
// vesting, nil is returned.
func (cva ContinuousVestingAccount) GetVestingCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// LockedCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (cva ContinuousVestingAccount) LockedCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (cva *ContinuousVestingAccount) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// GetStartTime returns the time when vesting starts for a continuous vesting
// account.
func (cva ContinuousVestingAccount) GetStartTime() int64 { _ = "STUB: not implemented"; return 0 }

// Validate checks for errors on the account fields
func (cva ContinuousVestingAccount) Validate() error { _ = "STUB: not implemented"; return nil }

func (cva ContinuousVestingAccount) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML returns the YAML representation of a ContinuousVestingAccount.
func (cva ContinuousVestingAccount) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Periodic Vesting Account

var _ vestexported.VestingAccount = (*PeriodicVestingAccount)(nil)
var _ authtypes.GenesisAccount = (*PeriodicVestingAccount)(nil)

// NewPeriodicVestingAccountRaw creates a new PeriodicVestingAccount object from BaseVestingAccount
func NewPeriodicVestingAccountRaw(bva *BaseVestingAccount, startTime int64, periods Periods) *PeriodicVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// NewPeriodicVestingAccount returns a new PeriodicVestingAccount
func NewPeriodicVestingAccount(baseAcc *authtypes.BaseAccount, originalVesting sdk.Coins, startTime int64, periods Periods, admin sdk.AccAddress) *PeriodicVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// GetVestedCoins returns the total number of vested coins. If no coins are vested,
// nil is returned.
func (pva PeriodicVestingAccount) GetVestedCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *

	// We must handle the case where the start time for a vesting account has
	// been set into the future or when the start of the chain is not exactly
	// known.
	new(sdk.Coins)
}

// track the start time of the next period

// for each period, if the period is over, add those coins as vested and check the next period.

// update the start time of the next period

// GetVestingCoins returns the total number of vesting coins. If no coins are
// vesting, nil is returned.
func (pva PeriodicVestingAccount) GetVestingCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// LockedCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (pva PeriodicVestingAccount) LockedCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (pva *PeriodicVestingAccount) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// GetStartTime returns the time when vesting starts for a periodic vesting
// account.
func (pva PeriodicVestingAccount) GetStartTime() int64 { _ = "STUB: not implemented"; return 0 }

// GetVestingPeriods returns vesting periods associated with periodic vesting account.
func (pva PeriodicVestingAccount) GetVestingPeriods() Periods {
	_ = "STUB: not implemented"
	return *new(Periods)
}

// Validate checks for errors on the account fields
func (pva PeriodicVestingAccount) Validate() error { _ = "STUB: not implemented"; return nil }

func (pva PeriodicVestingAccount) String() string { _ = "STUB: not implemented"; return "" }

// MarshalYAML returns the YAML representation of a PeriodicVestingAccount.
func (pva PeriodicVestingAccount) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delayed Vesting Account

var _ vestexported.VestingAccount = (*DelayedVestingAccount)(nil)
var _ authtypes.GenesisAccount = (*DelayedVestingAccount)(nil)

// NewDelayedVestingAccountRaw creates a new DelayedVestingAccount object from BaseVestingAccount
func NewDelayedVestingAccountRaw(bva *BaseVestingAccount) *DelayedVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// NewDelayedVestingAccount returns a DelayedVestingAccount
func NewDelayedVestingAccount(baseAcc *authtypes.BaseAccount, originalVesting sdk.Coins, endTime int64, admin sdk.AccAddress) *DelayedVestingAccount {
	_ = "STUB: not implemented"
	return nil
}

// GetVestedCoins returns the total amount of vested coins for a delayed vesting
// account. All coins are only vested once the schedule has elapsed.
func (dva DelayedVestingAccount) GetVestedCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// GetVestingCoins returns the total number of vesting coins for a delayed
// vesting account.
func (dva DelayedVestingAccount) GetVestingCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// LockedCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (dva DelayedVestingAccount) LockedCoins(blockTime time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (dva *DelayedVestingAccount) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// GetStartTime returns zero since a delayed vesting account has no start time.
func (dva DelayedVestingAccount) GetStartTime() int64 {
	_ = "STUB: not implemented"

	// Validate checks for errors on the account fields
	return 0
}

func (dva DelayedVestingAccount) Validate() error { _ = "STUB: not implemented"; return nil }

func (dva DelayedVestingAccount) String() string { _ = "STUB: not implemented"; return "" }

//-----------------------------------------------------------------------------
// Permanent Locked Vesting Account

var _ vestexported.VestingAccount = (*PermanentLockedAccount)(nil)
var _ authtypes.GenesisAccount = (*PermanentLockedAccount)(nil)

// NewPermanentLockedAccount returns a PermanentLockedAccount
func NewPermanentLockedAccount(baseAcc *authtypes.BaseAccount, coins sdk.Coins, admin sdk.AccAddress) *PermanentLockedAccount {
	_ = "STUB: not implemented"
	return nil
}

// ensure EndTime is set to 0, as PermanentLockedAccount's do not have an EndTime

// GetVestedCoins returns the total amount of vested coins for a permanent locked vesting
// account. All coins are only vested once the schedule has elapsed.
func (plva PermanentLockedAccount) GetVestedCoins(_ time.Time) sdk.Coins {
	_ = "STUB: not implemented"

	// GetVestingCoins returns the total number of vesting coins for a permanent locked
	// vesting account.
	return *new(sdk.Coins)
}

func (plva PermanentLockedAccount) GetVestingCoins(_ time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// LockedCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (plva PermanentLockedAccount) LockedCoins(_ time.Time) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (plva *PermanentLockedAccount) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	_ = "STUB: not implemented"
	return
}

// GetStartTime returns zero since a permanent locked vesting account has no start time.
func (plva PermanentLockedAccount) GetStartTime() int64 {
	_ = "STUB: not implemented"

	// GetEndTime returns a vesting account's end time, we return 0 to denote that
	// a permanently locked vesting account has no end time.
	return 0
}

func (plva PermanentLockedAccount) GetEndTime() int64 {
	_ = "STUB: not implemented"

	// Validate checks for errors on the account fields
	return 0
}

func (plva PermanentLockedAccount) Validate() error { _ = "STUB: not implemented"; return nil }

func (plva PermanentLockedAccount) String() string { _ = "STUB: not implemented"; return "" }

type getPK interface {
	GetPubKey() cryptotypes.PubKey
}

func getPKString(g getPK) string { _ = "STUB: not implemented"; return "" }

func marshalYaml(i interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
