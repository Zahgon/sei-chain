package types

import (
	"sort"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = IdentifiedClientState{}
	_ codectypes.UnpackInterfacesMessage = ConsensusStateWithHeight{}
)

// NewIdentifiedClientState creates a new IdentifiedClientState instance
func NewIdentifiedClientState(clientID string, clientState exported.ClientState) IdentifiedClientState {
	_ = "STUB: not implemented"
	return *new(IdentifiedClientState)
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (ics IdentifiedClientState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

var _ sort.Interface = IdentifiedClientStates{}

// IdentifiedClientStates defines a slice of ClientConsensusStates that supports the sort interface
type IdentifiedClientStates []IdentifiedClientState

// Len implements sort.Interface
func (ics IdentifiedClientStates) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface
	return 0
}

func (ics IdentifiedClientStates) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface
func (ics IdentifiedClientStates) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Sort is a helper function to sort the set of IdentifiedClientStates in place
func (ics IdentifiedClientStates) Sort() IdentifiedClientStates {
	_ = "STUB: not implemented"
	return *new(IdentifiedClientStates)
}

// NewConsensusStateWithHeight creates a new ConsensusStateWithHeight instance
func NewConsensusStateWithHeight(height Height, consensusState exported.ConsensusState) ConsensusStateWithHeight {
	_ = "STUB: not implemented"
	return *new(ConsensusStateWithHeight)
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (cswh ConsensusStateWithHeight) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateClientType validates the client type. It cannot be blank or empty. It must be a valid
// client identifier when used with '0' or the maximum uint64 as the sequence.
func ValidateClientType(clientType string) error { _ = "STUB: not implemented"; return nil }

// IsValidClientID will check client type format and if the sequence is a uint64
