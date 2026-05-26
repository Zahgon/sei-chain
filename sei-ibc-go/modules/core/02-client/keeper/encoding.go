package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// UnmarshalClientState attempts to decode and return an ClientState object from
// raw encoded bytes.
func (k Keeper) UnmarshalClientState(bz []byte) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// MustUnmarshalClientState attempts to decode and return an ClientState object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalClientState(bz []byte) exported.ClientState {
	_ = "STUB: not implemented"
	return *new(exported.ClientState)
}

// UnmarshalConsensusState attempts to decode and return an ConsensusState object from
// raw encoded bytes.
func (k Keeper) UnmarshalConsensusState(bz []byte) (exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), nil
}

// MustUnmarshalConsensusState attempts to decode and return an ConsensusState object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalConsensusState(bz []byte) exported.ConsensusState {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState)
}

// MustMarshalClientState attempts to encode an ClientState object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalClientState(clientState exported.ClientState) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MustMarshalConsensusState attempts to encode an ConsensusState object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalConsensusState(consensusState exported.ConsensusState) []byte {
	_ = "STUB: not implemented"
	return nil
}
