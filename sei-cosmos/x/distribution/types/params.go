package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeyCommunityTax        = []byte("communitytax")
	ParamStoreKeyBaseProposerReward  = []byte("baseproposerreward")
	ParamStoreKeyBonusProposerReward = []byte("bonusproposerreward")
	ParamStoreKeyWithdrawAddrEnabled = []byte("withdrawaddrenabled")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// 0%
// 0%
// 0%

func (p Params) String() string { _ = "STUB: not implemented"; return "" }

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	_ = "STUB: not implemented"
	return *new(paramtypes.ParamSetPairs)
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func validateCommunityTax(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateBaseProposerReward(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateBonusProposerReward(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateWithdrawAddrEnabled(i interface{}) error { _ = "STUB: not implemented"; return nil }
