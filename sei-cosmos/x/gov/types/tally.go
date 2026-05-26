package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// ValidatorGovInfo used for tallying
type ValidatorGovInfo struct {
	Address             sdk.ValAddress      // address of the validator operator
	BondedTokens        sdk.Int             // Power of a Validator
	DelegatorShares     sdk.Dec             // Total outstanding delegator shares
	DelegatorDeductions sdk.Dec             // Delegator deductions from validator's delegators voting independently
	Vote                WeightedVoteOptions // Vote of the validator
}

// NewValidatorGovInfo creates a ValidatorGovInfo instance
func NewValidatorGovInfo(address sdk.ValAddress, bondedTokens sdk.Int, delegatorShares,
	delegatorDeductions sdk.Dec, options WeightedVoteOptions) ValidatorGovInfo {
	_ = "STUB: not implemented"
	return *new(ValidatorGovInfo)
}

// NewTallyResult creates a new TallyResult instance
func NewTallyResult(yes, abstain, no, noWithVeto sdk.Int) TallyResult {
	_ = "STUB: not implemented"
	return *new(TallyResult)
}

// NewTallyResultFromMap creates a new TallyResult instance from a Option -> Dec map
func NewTallyResultFromMap(results map[VoteOption]sdk.Dec) TallyResult {
	_ = "STUB: not implemented"
	return *new(TallyResult)
}

// EmptyTallyResult returns an empty TallyResult.
func EmptyTallyResult() TallyResult { _ = "STUB: not implemented"; return *new(TallyResult) }

// Equals returns if two proposals are equal.
func (tr TallyResult) Equals(comp TallyResult) bool { _ = "STUB: not implemented"; return false }

// String implements stringer interface
func (tr TallyResult) String() string { _ = "STUB: not implemented"; return "" }
