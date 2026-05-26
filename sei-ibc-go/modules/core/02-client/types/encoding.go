package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// MustUnmarshalClientState attempts to decode and return an ClientState object from
// raw encoded bytes. It panics on error.
func MustUnmarshalClientState(cdc codec.BinaryCodec, bz []byte) exported.ClientState {
	_ = "STUB: not implemented"
	return *new(exported.ClientState)
}

// MustMarshalClientState attempts to encode an ClientState object and returns the
// raw encoded bytes. It panics on error.
func MustMarshalClientState(cdc codec.BinaryCodec, clientState exported.ClientState) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MarshalClientState protobuf serializes an ClientState interface
func MarshalClientState(cdc codec.BinaryCodec, clientStateI exported.ClientState) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalClientState returns an ClientState interface from raw encoded clientState
// bytes of a Proto-based ClientState type. An error is returned upon decoding
// failure.
func UnmarshalClientState(cdc codec.BinaryCodec, bz []byte) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// MustUnmarshalConsensusState attempts to decode and return an ConsensusState object from
// raw encoded bytes. It panics on error.
func MustUnmarshalConsensusState(cdc codec.BinaryCodec, bz []byte) exported.ConsensusState {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState)
}

// MustMarshalConsensusState attempts to encode a ConsensusState object and returns the
// raw encoded bytes. It panics on error.
func MustMarshalConsensusState(cdc codec.BinaryCodec, consensusState exported.ConsensusState) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MarshalConsensusState protobuf serializes a ConsensusState interface
func MarshalConsensusState(cdc codec.BinaryCodec, cs exported.ConsensusState) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalConsensusState returns a ConsensusState interface from raw encoded consensus state
// bytes of a Proto-based ConsensusState type. An error is returned upon decoding
// failure.
func UnmarshalConsensusState(cdc codec.BinaryCodec, bz []byte) (exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), nil
}

// MarshalHeader protobuf serializes a Header interface
func MarshalHeader(cdc codec.BinaryCodec, h exported.Header) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// MustMarshalHeader attempts to encode a Header object and returns the
	// raw encoded bytes. It panics on error.
}

func MustMarshalHeader(cdc codec.BinaryCodec, header exported.Header) []byte {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalHeader returns a Header interface from raw proto encoded header bytes.
// An error is returned upon decoding failure.
func UnmarshalHeader(cdc codec.BinaryCodec, bz []byte) (exported.Header, error) {
	_ = "STUB: not implemented"
	return *new(exported.Header), nil
}
