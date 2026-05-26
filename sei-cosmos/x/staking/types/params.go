package types

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Staking params default values
const (
	// DefaultUnbondingTime reflects three weeks in seconds as the default
	// unbonding time.
	// TODO: Justify our choice of default here.
	DefaultUnbondingTime time.Duration = time.Hour * 24 * 7 * 3

	// Default maximum number of bonded validators
	DefaultMaxValidators uint32 = 35

	// Default maximum entries in a UBD/RED pair
	DefaultMaxEntries uint32 = 7

	// DefaultHistorical entries is 10000. Apps that don't use IBC can ignore this
	// value by not adding the staking module to the application module manager's
	// SetOrderBeginBlockers.
	DefaultHistoricalEntries uint32 = 10000
)

var (
	// DefaultMinCommissionRate is set to 0%
	DefaultMinCommissionRate = sdk.NewDecWithPrec(5, 2)
)

var (
	KeyUnbondingTime                      = []byte("UnbondingTime")
	KeyMaxValidators                      = []byte("MaxValidators")
	KeyMaxEntries                         = []byte("MaxEntries")
	KeyMaxVotingPower                     = []byte("MaxVotingPower")
	KeyMaxVotingPowerEnforcementThreshold = []byte("MaxVotingPowerEnforcementThreshold")
	KeyBondDenom                          = []byte("BondDenom")
	KeyHistoricalEntries                  = []byte("HistoricalEntries")
	KeyPowerReduction                     = []byte("PowerReduction")
	KeyMinCommissionRate                  = []byte("MinCommissionRate")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamTable for staking module
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewParams creates a new Params instance
func NewParams(
	unbondingTime time.Duration,
	maxValidators, maxEntries, historicalEntries uint32,
	bondDenom string,
	minCommissionRate sdk.Dec,
	maxVotingPowerRatio sdk.Dec,
	maxVotingPowerEnforcementThreshold sdk.Int,
) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// String returns a human readable string representation of the parameters.
func (p Params) String() string { _ = "STUB: not implemented"; return "" }

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.LegacyAmino, value []byte) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.LegacyAmino, value []byte) (params Params, err error) {
	_ = "STUB: not implemented"
	return *new(Params), nil
}

// validate a set of params
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

func validateUnbondingTime(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMaxValidators(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMaxEntries(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMaxVotingPowerEnforcementThreshold(i interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMaxVotingPowerRatio(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateHistoricalEntries(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateBondDenom(i interface{}) error { _ = "STUB: not implemented"; return nil }

func ValidatePowerReduction(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMinCommissionRate(i interface{}) error { _ = "STUB: not implemented"; return nil }
