package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = QueryConnectionClientStateResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryConnectionConsensusStateResponse{}
)

// NewQueryConnectionResponse creates a new QueryConnectionResponse instance
func NewQueryConnectionResponse(
	connection ConnectionEnd, proof []byte, height clienttypes.Height,
) *QueryConnectionResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryClientConnectionsResponse creates a new ConnectionPaths instance
func NewQueryClientConnectionsResponse(
	connectionPaths []string, proof []byte, height clienttypes.Height,
) *QueryClientConnectionsResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryClientConnectionsRequest creates a new QueryClientConnectionsRequest instance
func NewQueryClientConnectionsRequest(clientID string) *QueryClientConnectionsRequest {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryConnectionClientStateResponse creates a newQueryConnectionClientStateResponse instance
func NewQueryConnectionClientStateResponse(identifiedClientState clienttypes.IdentifiedClientState, proof []byte, height clienttypes.Height) *QueryConnectionClientStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qccsr QueryConnectionClientStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryConnectionConsensusStateResponse creates a newQueryConnectionConsensusStateResponse instance
func NewQueryConnectionConsensusStateResponse(clientID string, anyConsensusState *codectypes.Any, consensusStateHeight exported.Height, proof []byte, height clienttypes.Height) *QueryConnectionConsensusStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qccsr QueryConnectionConsensusStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
