package types

import (
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var (
	_ exported.ChannelI             = (*Channel)(nil)
	_ exported.CounterpartyChannelI = (*Counterparty)(nil)
)

// NewChannel creates a new Channel instance
func NewChannel(
	state State, ordering Order, counterparty Counterparty,
	hops []string, version string,
) Channel {
	_ = "STUB: not implemented"
	return *new(Channel)
}

// GetState implements Channel interface.
func (ch Channel) GetState() int32 { _ = "STUB: not implemented"; return 0 }

// GetOrdering implements Channel interface.
func (ch Channel) GetOrdering() int32 { _ = "STUB: not implemented"; return 0 }

// GetCounterparty implements Channel interface.
func (ch Channel) GetCounterparty() exported.CounterpartyChannelI {
	_ = "STUB: not implemented"
	return *

	// GetConnectionHops implements Channel interface.
	new(exported.CounterpartyChannelI)
}

func (ch Channel) GetConnectionHops() []string { _ = "STUB: not implemented"; return nil }

// GetVersion implements Channel interface.
func (ch Channel) GetVersion() string {
	_ = "STUB: not implemented"

	// ValidateBasic performs a basic validation of the channel fields
	return ""
}

func (ch Channel) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewCounterparty returns a new Counterparty instance
func NewCounterparty(portID, channelID string) Counterparty {
	_ = "STUB: not implemented"
	return *new(Counterparty)
}

// GetPortID implements CounterpartyChannelI interface
func (c Counterparty) GetPortID() string {
	_ = "STUB: not implemented"

	// GetChannelID implements CounterpartyChannelI interface
	return ""
}

func (c Counterparty) GetChannelID() string {
	_ = "STUB: not implemented"

	// ValidateBasic performs a basic validation check of the identifiers
	return ""
}

func (c Counterparty) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewIdentifiedChannel creates a new IdentifiedChannel instance
func NewIdentifiedChannel(portID, channelID string, ch Channel) IdentifiedChannel {
	_ = "STUB: not implemented"
	return *new(IdentifiedChannel)
}

// ValidateBasic performs a basic validation of the identifiers and channel fields.
func (ic IdentifiedChannel) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
