package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

// NewGenesisState creates a new genesis state for the governance module
func NewGenesisState(startingProposalID uint64, dp DepositParams, vp VotingParams, tp TallyParams) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState defines the default governance genesis state
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

func (data GenesisState) Equal(other GenesisState) bool { _ = "STUB: not implemented"; return false }

// Empty returns true if a GenesisState is empty
func (data GenesisState) Empty() bool { _ = "STUB: not implemented"; return false }

// ValidateGenesis checks if parameters are within valid ranges
func ValidateGenesis(data *GenesisState) error { _ = "STUB: not implemented"; return nil }

var _ types.UnpackInterfacesMessage = GenesisState{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
