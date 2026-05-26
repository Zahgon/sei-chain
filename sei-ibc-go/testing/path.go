package ibctesting

import (
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
func (path *Path) RelayPacket(packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

// packet found, relay from A to B

// packet found, relay B to A
