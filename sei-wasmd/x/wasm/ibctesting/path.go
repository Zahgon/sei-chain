package ibctesting

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
)

// Path contains two endpoints representing two chains connected over IBC
type Path struct {
	EndpointA *Endpoint
	EndpointB *Endpoint
}

// NewPath constructs an endpoint for each chain using the default values
// for the endpoints. Each endpoint is updated to have a pointer to the
// counterparty endpoint.
func NewPath(chainA, chainB *TestChain) *Path { _ = "STUB: not implemented"; return nil }

// SetChannelOrdered sets the channel order for both endpoints to ORDERED.
func (path *Path) SetChannelOrdered() { _ = "STUB: not implemented"; return }

// RelayPacket attempts to relay the packet first on EndpointA and then on EndpointB
// if EndpointA does not contain a packet commitment for that packet. An error is returned
// if a relay step fails or the packet commitment does not exist on either endpoint.
func (path *Path) RelayPacket(packet channeltypes.Packet, ack []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// packet found, relay from A to B

// packet found, relay B to A

// SendMsg delivers the provided messages to the chain. The counterparty
// client is updated with the new source consensus state.
func (path *Path) SendMsg(msgs ...sdk.Msg) error { _ = "STUB: not implemented"; return nil }

func (path *Path) Invert() *Path { _ = "STUB: not implemented"; return nil }
