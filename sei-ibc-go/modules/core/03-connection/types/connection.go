package types

import (
	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.ConnectionI = (*ConnectionEnd)(nil)

// NewConnectionEnd creates a new ConnectionEnd instance.
func NewConnectionEnd(state State, clientID string, counterparty Counterparty, versions []*Version, delayPeriod uint64) ConnectionEnd {
	_ = "STUB: not implemented"
	return *new(ConnectionEnd)
}

// GetState implements the Connection interface
func (c ConnectionEnd) GetState() int32 { _ = "STUB: not implemented"; return 0 }

// GetClientID implements the Connection interface
func (c ConnectionEnd) GetClientID() string {
	_ = "STUB: not implemented"

	// GetCounterparty implements the Connection interface
	return ""
}

func (c ConnectionEnd) GetCounterparty() exported.CounterpartyConnectionI {
	_ = "STUB: not implemented"
	return *

	// GetVersions implements the Connection interface
	new(exported.CounterpartyConnectionI)
}

func (c ConnectionEnd) GetVersions() []exported.Version { _ = "STUB: not implemented"; return nil }

// GetDelayPeriod implements the Connection interface
func (c ConnectionEnd) GetDelayPeriod() uint64 { _ = "STUB: not implemented"; return 0 }

// ValidateBasic implements the Connection interface.
// NOTE: the protocol supports that the connection and client IDs match the
// counterparty's.
func (c ConnectionEnd) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

var _ exported.CounterpartyConnectionI = (*Counterparty)(nil)

// NewCounterparty creates a new Counterparty instance.
func NewCounterparty(clientID, connectionID string, prefix commitmenttypes.MerklePrefix) Counterparty {
	_ = "STUB: not implemented"
	return *new(Counterparty)
}

// GetClientID implements the CounterpartyConnectionI interface
func (c Counterparty) GetClientID() string {
	_ = "STUB: not implemented"

	// GetConnectionID implements the CounterpartyConnectionI interface
	return ""
}

func (c Counterparty) GetConnectionID() string { _ = "STUB: not implemented"; return "" }

// GetPrefix implements the CounterpartyConnectionI interface
func (c Counterparty) GetPrefix() exported.Prefix {
	_ = "STUB: not implemented"

	// ValidateBasic performs a basic validation check of the identifiers and prefix
	return *new(exported.Prefix)
}

func (c Counterparty) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewIdentifiedConnection creates a new IdentifiedConnection instance
func NewIdentifiedConnection(connectionID string, conn ConnectionEnd) IdentifiedConnection {
	_ = "STUB: not implemented"
	return *new(IdentifiedConnection)
}

// ValidateBasic performs a basic validation of the connection identifier and connection fields.
func (ic IdentifiedConnection) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
