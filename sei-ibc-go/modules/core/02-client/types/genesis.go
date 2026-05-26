package types

import (
	"sort"

	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = IdentifiedClientState{}
	_ codectypes.UnpackInterfacesMessage = ClientsConsensusStates{}
	_ codectypes.UnpackInterfacesMessage = ClientConsensusStates{}
	_ codectypes.UnpackInterfacesMessage = GenesisState{}
)

var (
	_ sort.Interface           = ClientsConsensusStates{}
	_ exported.GenesisMetadata = GenesisMetadata{}
)

// ClientsConsensusStates defines a slice of ClientConsensusStates that supports the sort interface
type ClientsConsensusStates []ClientConsensusStates

// Len implements sort.Interface
func (ccs ClientsConsensusStates) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface
	return 0
}

func (ccs ClientsConsensusStates) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface
func (ccs ClientsConsensusStates) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Sort is a helper function to sort the set of ClientsConsensusStates in place
func (ccs ClientsConsensusStates) Sort() ClientsConsensusStates {
	_ = "STUB: not implemented"
	return *new(ClientsConsensusStates)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (ccs ClientsConsensusStates) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewClientConsensusStates creates a new ClientConsensusStates instance.
func NewClientConsensusStates(clientID string, consensusStates []ConsensusStateWithHeight) ClientConsensusStates {
	_ = "STUB: not implemented"
	return *new(ClientConsensusStates)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (ccs ClientConsensusStates) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewGenesisState creates a GenesisState instance.
func NewGenesisState(
	clients []IdentifiedClientState, clientsConsensus ClientsConsensusStates, clientsMetadata []IdentifiedGenesisMetadata,
	params Params, createLocalhost bool, nextClientSequence uint64,
) GenesisState {
	_ = "STUB: not implemented"
	return *new(GenesisState)
}

// DefaultGenesisState returns the ibc client submodule's default genesis state.
func DefaultGenesisState() GenesisState { _ = "STUB: not implemented"; return *new(GenesisState) }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (gs GenesisState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	_ = "STUB: not implemented"
	// keep track of the max sequence to ensure it is less than
	// the next sequence used in creating client identifers.
	return nil
}

// add client id to validClients map

// check that consensus state is for a client in the genesis clients list

// ensure consensus state type matches client state type

// check that metadata is for a client in the genesis clients list

// NewGenesisMetadata is a constructor for GenesisMetadata
func NewGenesisMetadata(key, val []byte) GenesisMetadata {
	_ = "STUB: not implemented"
	return *new(GenesisMetadata)
}

// GetKey returns the key of metadata. Implements exported.GenesisMetadata interface.
func (gm GenesisMetadata) GetKey() []byte {
	_ = "STUB: not implemented"

	// GetValue returns the value of metadata. Implements exported.GenesisMetadata interface.
	return nil
}

func (gm GenesisMetadata) GetValue() []byte {
	_ = "STUB: not implemented"

	// Validate ensures key and value of metadata are not empty
	return nil
}

func (gm GenesisMetadata) Validate() error { _ = "STUB: not implemented"; return nil }

// NewIdentifiedGenesisMetadata takes in a client ID and list of genesis metadata for that client
// and constructs a new IdentifiedGenesisMetadata.
func NewIdentifiedGenesisMetadata(clientID string, gms []GenesisMetadata) IdentifiedGenesisMetadata {
	_ = "STUB: not implemented"
	return *new(IdentifiedGenesisMetadata)
}
