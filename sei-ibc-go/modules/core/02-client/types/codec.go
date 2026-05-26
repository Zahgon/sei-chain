package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// RegisterInterfaces registers the client interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

// PackClientState constructs a new Any packed with the given client state value. It returns
// an error if the client state can't be casted to a protobuf message or if the concrete
// implemention is not registered to the protobuf codec.
func PackClientState(clientState exported.ClientState) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackClientState unpacks an Any into a ClientState. It returns an error if the
// client state can't be unpacked into a ClientState.
func UnpackClientState(any *codectypes.Any) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// PackConsensusState constructs a new Any packed with the given consensus state value. It returns
// an error if the consensus state can't be casted to a protobuf message or if the concrete
// implemention is not registered to the protobuf codec.
func PackConsensusState(consensusState exported.ConsensusState) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustPackConsensusState calls PackConsensusState and panics on error.
func MustPackConsensusState(consensusState exported.ConsensusState) *codectypes.Any {
	_ = "STUB: not implemented"
	return nil
}

// UnpackConsensusState unpacks an Any into a ConsensusState. It returns an error if the
// consensus state can't be unpacked into a ConsensusState.
func UnpackConsensusState(any *codectypes.Any) (exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), nil
}

// PackHeader constructs a new Any packed with the given header value. It returns
// an error if the header can't be casted to a protobuf message or if the concrete
// implemention is not registered to the protobuf codec.
func PackHeader(header exported.Header) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackHeader unpacks an Any into a Header. It returns an error if the
// consensus state can't be unpacked into a Header.
func UnpackHeader(any *codectypes.Any) (exported.Header, error) {
	_ = "STUB: not implemented"
	return *new(exported.Header), nil
}

// PackMisbehaviour constructs a new Any packed with the given misbehaviour value. It returns
// an error if the misbehaviour can't be casted to a protobuf message or if the concrete
// implemention is not registered to the protobuf codec.
func PackMisbehaviour(misbehaviour exported.Misbehaviour) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackMisbehaviour unpacks an Any into a Misbehaviour. It returns an error if the
// Any can't be unpacked into a Misbehaviour.
func UnpackMisbehaviour(any *codectypes.Any) (exported.Misbehaviour, error) {
	_ = "STUB: not implemented"
	return *new(exported.Misbehaviour), nil
}
