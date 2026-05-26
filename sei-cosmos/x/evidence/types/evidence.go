package types

import (
	"time"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/exported"
)

// Evidence type constants
const (
	RouteEquivocation = "equivocation"
	TypeEquivocation  = "equivocation"
)

var _ exported.Evidence = &Equivocation{}

// Route returns the Evidence Handler route for an Equivocation type.
func (e *Equivocation) Route() string { _ = "STUB: not implemented"; return "" }

// Type returns the Evidence Handler type for an Equivocation type.
func (e *Equivocation) Type() string { _ = "STUB: not implemented"; return "" }

func (e *Equivocation) String() string { _ = "STUB: not implemented"; return "" }

// Hash returns the hash of an Equivocation object.
func (e *Equivocation) Hash() tmbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(tmbytes.HexBytes)
}

// ValidateBasic performs basic stateless validation checks on an Equivocation object.
func (e *Equivocation) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetConsensusAddress returns the validator's consensus address at time of the
// Equivocation infraction.
func (e Equivocation) GetConsensusAddress() sdk.ConsAddress {
	_ = "STUB: not implemented"
	return *new(sdk.ConsAddress)
}

// GetHeight returns the height at time of the Equivocation infraction.
func (e Equivocation) GetHeight() int64 {
	_ = "STUB: not implemented"

	// GetTime returns the time at time of the Equivocation infraction.
	return 0
}

func (e Equivocation) GetTime() time.Time {
	_ = "STUB: not implemented"

	// GetValidatorPower returns the validator's power at time of the Equivocation
	// infraction.
	return *new(time.Time)
}

func (e Equivocation) GetValidatorPower() int64 {
	_ = "STUB: not implemented"

	// GetTotalPower is a no-op for the Equivocation type.
	return 0
}

func (e Equivocation) GetTotalPower() int64 {
	_ = "STUB: not implemented"

	// FromABCIEvidence converts a Tendermint concrete Evidence type to
	// SDK Evidence using Equivocation as the concrete type.
	return 0
}

func FromABCIEvidence(e abci.Evidence) exported.Evidence {
	_ = "STUB: not implemented"
	return *new(exported.Evidence)
}
