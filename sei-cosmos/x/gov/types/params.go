package types

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Default period for deposits & voting
const (
	DefaultPeriod          time.Duration = time.Hour * 24 * 2 // 2 days
	DefaultExpeditedPeriod time.Duration = time.Hour * 24     // 1 day
)

// Default governance params
var (
	DefaultMinDepositTokens          = sdk.NewInt(10000000)
	DefaultMinExpeditedDepositTokens = sdk.NewInt(20000000)
	DefaultQuorum                    = sdk.NewDecWithPrec(334, 3)
	DefaultExpeditedQuorum           = sdk.NewDecWithPrec(667, 3)
	DefaultThreshold                 = sdk.NewDecWithPrec(5, 1)
	DefaultExpeditedThreshold        = sdk.NewDecWithPrec(667, 3)
	DefaultVetoThreshold             = sdk.NewDecWithPrec(334, 3)
)

// Parameter store key
var (
	ParamStoreKeyDepositParams = []byte("depositparams")
	ParamStoreKeyVotingParams  = []byte("votingparams")
	ParamStoreKeyTallyParams   = []byte("tallyparams")
)

// ParamKeyTable - Key declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// NewDepositParams creates a new DepositParams object
func NewDepositParams(minDeposit sdk.Coins, minExpeditedDeposit sdk.Coins, maxDepositPeriod time.Duration) DepositParams {
	_ = "STUB: not implemented"
	return *new(DepositParams)
}

// DefaultDepositParams default parameters for deposits
func DefaultDepositParams() DepositParams { _ = "STUB: not implemented"; return *new(DepositParams) }

// String implements stringer insterface
func (dp DepositParams) String() string { _ = "STUB: not implemented"; return "" }

// GetMinimumDeposit returns minimum deposit based on the value isExpedited
func (dp DepositParams) GetMinimumDeposit(isExpedited bool) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// Equal checks equality of DepositParams
func (dp DepositParams) Equal(dp2 DepositParams) bool { _ = "STUB: not implemented"; return false }

func validateDepositParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

// NewTallyParams creates a new TallyParams object
func NewTallyParams(quorum, expeditedQuorum, threshold, expeditedThreshold, vetoThreshold sdk.Dec) TallyParams {
	_ = "STUB: not implemented"
	return *new(TallyParams)
}

// DefaultTallyParams default parameters for tallying
func DefaultTallyParams() TallyParams { _ = "STUB: not implemented"; return *new(TallyParams) }

// GetThreshold returns threshold based on the value isExpedited
func (tp TallyParams) GetThreshold(isExpedited bool) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// GetQuorum returns quorum based on the value isExpedited
func (tp TallyParams) GetQuorum(isExpedited bool) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}

// Equal checks equality of TallyParams
func (tp TallyParams) Equal(other TallyParams) bool { _ = "STUB: not implemented"; return false }

// String implements stringer insterface
func (tp TallyParams) String() string { _ = "STUB: not implemented"; return "" }

func validateTallyParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

// NewVotingParams creates a new VotingParams object
func NewVotingParams(votingPeriod time.Duration, expeditedPeriod time.Duration) VotingParams {
	_ = "STUB: not implemented"
	return *new(VotingParams)
}

// GetVotingPeriod returns voting period based on whether isExpedited is requested.
func (vp VotingParams) GetVotingPeriod(isExpedited bool) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// DefaultVotingParams default parameters for voting
func DefaultVotingParams() VotingParams { _ = "STUB: not implemented"; return *new(VotingParams) }

// Equal checks equality of TallyParams
func (vp VotingParams) Equal(other VotingParams) bool { _ = "STUB: not implemented"; return false }

// String implements stringer interface
func (vp VotingParams) String() string { _ = "STUB: not implemented"; return "" }

func validateVotingParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

// Params returns all of the governance params
type Params struct {
	VotingParams  VotingParams  `json:"voting_params" yaml:"voting_params"`
	TallyParams   TallyParams   `json:"tally_params" yaml:"tally_params"`
	DepositParams DepositParams `json:"deposit_params" yaml:"deposit_params"`
}

func (gp Params) String() string { _ = "STUB: not implemented"; return "" }

// NewParams creates a new gov Params instance
func NewParams(vp VotingParams, tp TallyParams, dp DepositParams) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// DefaultParams default governance params
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }
