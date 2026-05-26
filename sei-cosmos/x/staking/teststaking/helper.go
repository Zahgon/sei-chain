package teststaking

import (
	"testing"
	"time"

	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
)

// Helper is a structure which wraps the staking handler
// and provides methods useful in tests
type Helper struct {
	t *testing.T
	h sdk.Handler
	k keeper.Keeper

	Ctx        sdk.Context
	Commission stakingtypes.CommissionRates
	// Coin Denomination
	Denom string
}

// NewHelper creates staking Handler wrapper for tests
func NewHelper(t *testing.T, ctx sdk.Context, k keeper.Keeper) *Helper {
	_ = "STUB: not implemented"
	return nil
}

// CreateValidator calls handler to create a new staking validator
func (sh *Helper) CreateValidator(addr sdk.ValAddress, pk cryptotypes.PubKey, stakeAmount sdk.Int, ok bool) {
	_ = "STUB: not implemented"
	return
}

// CreateValidatorWithValPower calls handler to create a new staking validator with zero
// commission
func (sh *Helper) CreateValidatorWithValPower(addr sdk.ValAddress, pk cryptotypes.PubKey, valPower int64, ok bool) sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// CreateValidatorMsg returns a message used to create validator in this service.
func (sh *Helper) CreateValidatorMsg(addr sdk.ValAddress, pk cryptotypes.PubKey, stakeAmount sdk.Int) *stakingtypes.MsgCreateValidator {
	_ = "STUB: not implemented"
	return nil
}

func (sh *Helper) createValidator(addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, ok bool) {
	_ = "STUB: not implemented"
	return
}

// Delegate calls handler to delegate stake for a validator
func (sh *Helper) Delegate(delegator sdk.AccAddress, val sdk.ValAddress, amount sdk.Int) {
	_ = "STUB: not implemented"
	return
}

// DelegateWithPower calls handler to delegate stake for a validator
func (sh *Helper) DelegateWithPower(delegator sdk.AccAddress, val sdk.ValAddress, power int64) {
	_ = "STUB: not implemented"
	return
}

// Undelegate calls handler to unbound some stake from a validator.
func (sh *Helper) Undelegate(delegator sdk.AccAddress, val sdk.ValAddress, amount sdk.Int, ok bool) *sdk.Result {
	_ = "STUB: not implemented"
	return nil
}

// Handle calls staking handler on a given message
func (sh *Helper) Handle(msg sdk.Msg, ok bool) *sdk.Result { _ = "STUB: not implemented"; return nil }

// CheckValidator asserts that a validor exists and has a given status (if status!="")
// and if has a right jailed flag.
func (sh *Helper) CheckValidator(addr sdk.ValAddress, status stakingtypes.BondStatus, jailed bool) stakingtypes.Validator {
	_ = "STUB: not implemented"
	return *new(stakingtypes.Validator)
}

// CheckDelegator asserts that a delegator exists
func (sh *Helper) CheckDelegator(delegator sdk.AccAddress, val sdk.ValAddress, found bool) {
	_ = "STUB: not implemented"
	return
}

// TurnBlock calls EndBlocker and updates the block time
func (sh *Helper) TurnBlock(newTime time.Time) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// TurnBlockTimeDiff calls EndBlocker and updates the block time by adding the
// duration to the current block time
func (sh *Helper) TurnBlockTimeDiff(diff time.Duration) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// DefaultCommission constructs a commission rates with all min commission (0.05).
func DefaultCommission() stakingtypes.CommissionRates {
	_ = "STUB: not implemented"
	return *new(stakingtypes.CommissionRates)
}
