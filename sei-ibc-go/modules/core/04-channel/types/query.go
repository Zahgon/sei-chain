package types

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = QueryChannelClientStateResponse{}
	_ codectypes.UnpackInterfacesMessage = QueryChannelConsensusStateResponse{}
)

// NewQueryChannelResponse creates a new QueryChannelResponse instance
func NewQueryChannelResponse(channel Channel, proof []byte, height clienttypes.Height) *QueryChannelResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryChannelClientStateResponse creates a newQueryChannelClientStateResponse instance
func NewQueryChannelClientStateResponse(identifiedClientState clienttypes.IdentifiedClientState, proof []byte, height clienttypes.Height) *QueryChannelClientStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qccsr QueryChannelClientStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryChannelConsensusStateResponse creates a newQueryChannelConsensusStateResponse instance
func NewQueryChannelConsensusStateResponse(clientID string, anyConsensusState *codectypes.Any, consensusStateHeight exported.Height, proof []byte, height clienttypes.Height) *QueryChannelConsensusStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (qccsr QueryChannelConsensusStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketCommitmentResponse creates a new QueryPacketCommitmentResponse instance
func NewQueryPacketCommitmentResponse(
	commitment []byte, proof []byte, height clienttypes.Height,
) *QueryPacketCommitmentResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketReceiptResponse creates a new QueryPacketReceiptResponse instance
func NewQueryPacketReceiptResponse(
	recvd bool, proof []byte, height clienttypes.Height,
) *QueryPacketReceiptResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketAcknowledgementResponse creates a new QueryPacketAcknowledgementResponse instance
func NewQueryPacketAcknowledgementResponse(
	acknowledgement []byte, proof []byte, height clienttypes.Height,
) *QueryPacketAcknowledgementResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryNextSequenceReceiveResponse creates a new QueryNextSequenceReceiveResponse instance
func NewQueryNextSequenceReceiveResponse(
	sequence uint64, proof []byte, height clienttypes.Height,
) *QueryNextSequenceReceiveResponse {
	_ = "STUB: not implemented"
	return nil
}
