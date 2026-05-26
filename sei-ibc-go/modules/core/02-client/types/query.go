package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
)

var (
	_ codectypes.UnpackInterfacesMessage = QueryClientStateResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryClientStatesResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryConsensusStateResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryConsensusStatesResponse{}
)

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qcsr QueryClientStatesResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryClientStateResponse creates a new QueryClientStateResponse instance.
func NewQueryClientStateResponse(
	clientStateAny *codectypes.Any, proof []byte, height Height,
) *QueryClientStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qcsr QueryClientStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qcsr QueryConsensusStatesResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryConsensusStateResponse creates a new QueryConsensusStateResponse instance.
func NewQueryConsensusStateResponse(
	consensusStateAny *codectypes.Any, proof []byte, height Height,
) *QueryConsensusStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qcsr QueryConsensusStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
