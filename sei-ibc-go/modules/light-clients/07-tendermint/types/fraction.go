package types

import (
	tmmath "github.com/sei-protocol/sei-chain/sei-tendermint/libs/math"
	"github.com/sei-protocol/sei-chain/sei-tendermint/light"
)

// DefaultTrustLevel is the tendermint light client default trust level
var DefaultTrustLevel = NewFractionFromTm(light.DefaultTrustLevel)

// NewFractionFromTm returns a new Fraction instance from a tmmath.Fraction
func NewFractionFromTm(f tmmath.Fraction) Fraction {
	_ = "STUB: not implemented"
	return *new(Fraction)
}

// ToTendermint converts Fraction to tmmath.Fraction
func (f Fraction) ToTendermint() tmmath.Fraction {
	_ = "STUB: not implemented"
	return *new(tmmath.Fraction)
}
