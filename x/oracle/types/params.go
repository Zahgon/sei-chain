package types

import (
	"github.com/sei-protocol/sei-chain/x/oracle/utils"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramstypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Parameter keys
var (
	KeyVotePeriod        = []byte("VotePeriod")
	KeyVoteThreshold     = []byte("VoteThreshold")
	KeyRewardBand        = []byte("RewardBand")
	KeyWhitelist         = []byte("Whitelist")
	KeySlashFraction     = []byte("SlashFraction")
	KeySlashWindow       = []byte("SlashWindow")
	KeyMinValidPerWindow = []byte("MinValidPerWindow")
	KeyLookbackDuration  = []byte("LookbackDuration")
)

// Default parameter values
const (
	DefaultVotePeriod  = 2                      // Voting every other block
	DefaultSlashWindow = utils.BlocksPerDay * 2 // 2 days for oracle slashing
)

// Default parameter values
var (
	DefaultVoteThreshold = sdk.NewDecWithPrec(667, 3) // 66.7%
	DefaultRewardBand    = sdk.NewDecWithPrec(2, 2)   // 2% (-1, 1)
	DefaultWhitelist     = DenomList{
		{Name: utils.MicroAtomDenom},
		// 		{Name: utils.MicroUsdcDenom},
		// 		{Name: utils.MicroSeiDenom},
		{Name: utils.MicroEthDenom},
	}
	DefaultSlashFraction     = sdk.NewDecWithPrec(0, 4) // 0.00%
	DefaultMinValidPerWindow = sdk.ZeroDec()            // 0% now that Oracle Price Feeder is retired.
	DefaultLookbackDuration  = uint64(3600)             // in seconds
)

var _ paramstypes.ParamSet = &Params{}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramstypes.KeyTable)
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of oracle module's parameters.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramstypes.ParamSetPairs)
}

// String implements fmt.Stringer interface
func (p Params) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs basic validation on oracle parameters.
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

func validateVotePeriod(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateVoteThreshold(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateRewardBand(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateWhitelist(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateSlashFraction(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateSlashWindow(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateMinValidPerWindow(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateLookbackDuration(i interface{}) error { _ = "STUB: not implemented"; return nil }
