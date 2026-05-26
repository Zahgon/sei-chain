package types

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"

	tmprotocrypto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

// OracleDecPrecision nolint
const OracleDecPrecision = 8

// GenerateRandomTestCase nolint
// nolint:staticcheck
func GenerateRandomTestCase() (rates []float64, valValAddrs []sdk.ValAddress, stakingKeeper DummyStakingKeeper) {
	_ = "STUB: not implemented"
	return nil, nil, *new(DummyStakingKeeper)
}

var _ StakingKeeper = DummyStakingKeeper{}

// DummyStakingKeeper dummy staking keeper to test ballot
type DummyStakingKeeper struct {
	validators []MockValidator
}

// NewDummyStakingKeeper returns new DummyStakingKeeper instance
func NewDummyStakingKeeper(validators []MockValidator) DummyStakingKeeper {
	_ = "STUB: not implemented"
	return *new(DummyStakingKeeper)
}

// Validators nolint
func (sk DummyStakingKeeper) Validators() []MockValidator { _ = "STUB: not implemented"; return nil }

// Validator nolint
func (sk DummyStakingKeeper) Validator(_ sdk.Context, address sdk.ValAddress) stakingtypes.ValidatorI {
	_ = "STUB: not implemented"
	return *new(stakingtypes.ValidatorI)
}

// TotalBondedTokens nolint
func (DummyStakingKeeper) TotalBondedTokens(_ sdk.Context) sdk.Int {
	_ = "STUB: not implemented"
	return *

	// Slash nolint
	new(sdk.Int)
}

func (DummyStakingKeeper) Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec) {
	_ = "STUB: not implemented"

	// ValidatorsPowerStoreIterator
	return
}

func (DummyStakingKeeper) ValidatorsPowerStoreIterator(_ sdk.Context) sdk.Iterator {
	_ = "STUB: not implemented"
	return *new(sdk.Iterator)
}

// Jail
func (DummyStakingKeeper) Jail(sdk.Context, sdk.ConsAddress) {
	_ = "STUB: not implemented"

	// GetLastValidatorPower
	return
}

func (sk DummyStakingKeeper) GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64) {
	_ = "STUB: not implemented"
	return 0
}

// MaxValidators returns the maximum amount of bonded validators
func (DummyStakingKeeper) MaxValidators(sdk.Context) uint32 {
	_ = "STUB: not implemented"

	// PowerReduction - is the amount of staking tokens required for 1 unit of consensus-engine power
	return 0
}

func (DummyStakingKeeper) PowerReduction(_ sdk.Context) (res sdk.Int) {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}

// MockValidator
type MockValidator struct {
	power    int64
	operator sdk.ValAddress
}

var _ stakingtypes.ValidatorI = MockValidator{}

func (MockValidator) IsJailed() bool     { _ = "STUB: not implemented"; return false }
func (MockValidator) GetMoniker() string { _ = "STUB: not implemented"; return "" }
func (MockValidator) GetStatus() stakingtypes.BondStatus {
	_ = "STUB: not implemented"
	return *new(stakingtypes.BondStatus)
}
func (MockValidator) IsBonded() bool    { _ = "STUB: not implemented"; return false }
func (MockValidator) IsUnbonded() bool  { _ = "STUB: not implemented"; return false }
func (MockValidator) IsUnbonding() bool { _ = "STUB: not implemented"; return false }
func (v MockValidator) GetOperator() sdk.ValAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ValAddress)
}
func (MockValidator) ConsPubKey() (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}
func (MockValidator) TmConsPublicKey() (tmprotocrypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(tmprotocrypto.PublicKey), nil
}

func (MockValidator) GetConsAddr() (sdk.ConsAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.ConsAddress), nil
}
func (v MockValidator) GetTokens() sdk.Int { _ = "STUB: not implemented"; return *new(sdk.Int) }

func (v MockValidator) GetBondedTokens() sdk.Int { _ = "STUB: not implemented"; return *new(sdk.Int) }

func (v MockValidator) GetConsensusPower(_ sdk.Int) int64 { _ = "STUB: not implemented"; return 0 }
func (v *MockValidator) SetConsensusPower(power int64)    { _ = "STUB: not implemented"; return }
func (v MockValidator) GetCommission() sdk.Dec            { _ = "STUB: not implemented"; return *new(sdk.Dec) }
func (v MockValidator) GetMinSelfDelegation() sdk.Int {
	_ = "STUB: not implemented"
	return *new(sdk.Int)
}
func (v MockValidator) GetDelegatorShares() sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}
func (v MockValidator) TokensFromShares(sdk.Dec) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}
func (v MockValidator) TokensFromSharesTruncated(sdk.Dec) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}
func (v MockValidator) TokensFromSharesRoundUp(sdk.Dec) sdk.Dec {
	_ = "STUB: not implemented"
	return *new(sdk.Dec)
}
func (v MockValidator) SharesFromTokens(_ sdk.Int) (sdk.Dec, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec), nil
}
func (v MockValidator) SharesFromTokensTruncated(_ sdk.Int) (sdk.Dec, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Dec), nil
}

func NewMockValidator(valAddr sdk.ValAddress, power int64) MockValidator {
	_ = "STUB: not implemented"
	return *new(MockValidator)
}
