package types

import (
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
)

const (
	ModuleDenomPrefix = "factory"
	// See the TokenFactory readme for a derivation of these.
	// TL;DR, MaxSubdenomLength + MaxHrpLength = 60 comes from SDK max denom length = 128
	// and the structure of tokenfactory denoms.
	MaxSubdenomLength = 44
	MaxHrpLength      = 16
	// MaxCreatorLength = 59 + MaxHrpLength
	MaxCreatorLength = 59 + MaxHrpLength
)

// GetTokenDenom constructs a denom string for tokens created by tokenfactory
// based on an input creator address and a subdenom
// The denom constructed is factory/{creator}/{subdenom}
func GetTokenDenom(creator, subdenom string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DeconstructDenom takes a token denom string and verifies that it is a valid
// denom of the tokenfactory module, and is of the form `factory/{creator}/{subdenom}`
// If valid, it returns the creator address and subdenom
func DeconstructDenom(denom string) (creator string, subdenom string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Handle the case where a denom has a slash in its subdenom. For example,
// when we did the split, we'd turn factory/accaddr/atomderivative/sikka into ["factory", "accaddr", "atomderivative", "sikka"]
// So we have to join [2:] with a "/" as the delimiter to get back the correct subdenom which should be "atomderivative/sikka"

// NewTokenFactoryDenomMintCoinsRestriction creates and returns a BankMintingRestrictionFn that only allows minting of
// valid tokenfactory denoms
func NewTokenFactoryDenomMintCoinsRestriction() bankkeeper.MintingRestrictionFn {
	_ = "STUB: not implemented"
	return *new(bankkeeper.MintingRestrictionFn)
}
