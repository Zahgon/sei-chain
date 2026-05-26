package types

// NewPacketState creates a new PacketState instance.
func NewPacketState(portID, channelID string, seq uint64, data []byte) PacketState {
	_ = "STUB: not implemented"
	return *new(PacketState)
}

// Validate performs basic validation of fields returning an error upon any
// failure.
func (pa PacketState) Validate() error { _ = "STUB: not implemented"; return nil }

// NewPacketSequence creates a new PacketSequences instance.
func NewPacketSequence(portID, channelID string, seq uint64) PacketSequence {
	_ = "STUB: not implemented"
	return *new(PacketSequence)
}

// Validate performs basic validation of fields returning an error upon any
// failure.
func (ps PacketSequence) Validate() error { _ = "STUB: not implemented"; return nil }

// NewGenesisState creates a GenesisState instance.
func NewGenesisState(
	channels []IdentifiedChannel, acks, receipts, commitments []PacketState,
	sendSeqs, recvSeqs, ackSeqs []PacketSequence, nextChannelSequence uint64,
) GenesisState {
	_ = "STUB: not implemented"
	return *new(GenesisState)
}

// DefaultGenesisState returns the ibc channel submodule's default genesis state.
func DefaultGenesisState() GenesisState { _ = "STUB: not implemented"; return *new(GenesisState) }

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	_ = "STUB: not implemented"
	// keep track of the max sequence to ensure it is less than
	// the next sequence used in creating connection identifers.
	return nil
}

func validateGenFields(portID, channelID string, sequence uint64) error {
	_ = "STUB: not implemented"
	return nil
}
