package host

import (
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

const (
	// ModuleName is the name of the IBC module
	ModuleName = "ibc"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the IBC module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the IBC module
	RouterKey string = ModuleName
)

// KVStore key prefixes for IBC
var (
	KeyClientStorePrefix = []byte("clients")
)

// KVStore key prefixes for IBC
const (
	KeyClientState             = "clientState"
	KeyConsensusStatePrefix    = "consensusStates"
	KeyConnectionPrefix        = "connections"
	KeyChannelEndPrefix        = "channelEnds"
	KeyChannelPrefix           = "channels"
	KeyPortPrefix              = "ports"
	KeySequencePrefix          = "sequences"
	KeyChannelCapabilityPrefix = "capabilities"
	KeyNextSeqSendPrefix       = "nextSequenceSend"
	KeyNextSeqRecvPrefix       = "nextSequenceRecv"
	KeyNextSeqAckPrefix        = "nextSequenceAck"
	KeyPacketCommitmentPrefix  = "commitments"
	KeyPacketAckPrefix         = "acks"
	KeyPacketReceiptPrefix     = "receipts"
)

// FullClientPath returns the full path of a specific client path in the format:
// "clients/{clientID}/{path}" as a string.
func FullClientPath(clientID string, path string) string { _ = "STUB: not implemented"; return "" }

// FullClientKey returns the full path of specific client path in the format:
// "clients/{clientID}/{path}" as a byte array.
func FullClientKey(clientID string, path []byte) []byte { _ = "STUB: not implemented"; return nil }

// ICS02
// The following paths are the keys to the store as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-002-client-semantics#path-space

// FullClientStatePath takes a client identifier and returns a Path under which to store a
// particular client state
func FullClientStatePath(clientID string) string { _ = "STUB: not implemented"; return "" }

// FullClientStateKey takes a client identifier and returns a Key under which to store a
// particular client state.
func FullClientStateKey(clientID string) []byte { _ = "STUB: not implemented"; return nil }

// ClientStateKey returns a store key under which a particular client state is stored
// in a client prefixed store
func ClientStateKey() []byte { _ = "STUB: not implemented"; return nil }

// FullConsensusStatePath takes a client identifier and returns a Path under which to
// store the consensus state of a client.
func FullConsensusStatePath(clientID string, height exported.Height) string {
	_ = "STUB: not implemented"
	return ""
}

// FullConsensusStateKey returns the store key for the consensus state of a particular
// client.
func FullConsensusStateKey(clientID string, height exported.Height) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusStatePath returns the suffix store key for the consensus state at a
// particular height stored in a client prefixed store.
func ConsensusStatePath(height exported.Height) string { _ = "STUB: not implemented"; return "" }

// ConsensusStateKey returns the store key for a the consensus state of a particular
// client stored in a client prefixed store.
func ConsensusStateKey(height exported.Height) []byte { _ = "STUB: not implemented"; return nil }

// ICS03
// The following paths are the keys to the store as defined in https://github.com/cosmos/ibc/blob/master/spec/core/ics-003-connection-semantics#store-paths

// ClientConnectionsPath defines a reverse mapping from clients to a set of connections
func ClientConnectionsPath(clientID string) string { _ = "STUB: not implemented"; return "" }

// ClientConnectionsKey returns the store key for the connections of a given client
func ClientConnectionsKey(clientID string) []byte { _ = "STUB: not implemented"; return nil }

// ConnectionPath defines the path under which connection paths are stored
func ConnectionPath(connectionID string) string { _ = "STUB: not implemented"; return "" }

// ConnectionKey returns the store key for a particular connection
func ConnectionKey(connectionID string) []byte { _ = "STUB: not implemented"; return nil }

// ICS04
// The following paths are the keys to the store as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-004-channel-and-packet-semantics#store-paths

// ChannelPath defines the path under which channels are stored
func ChannelPath(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

// ChannelKey returns the store key for a particular channel
func ChannelKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

// ChannelCapabilityPath defines the path under which capability keys associated
// with a channel are stored
func ChannelCapabilityPath(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

// NextSequenceSendPath defines the next send sequence counter store path
func NextSequenceSendPath(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

// NextSequenceSendKey returns the store key for the send sequence of a particular
// channel binded to a specific port.
func NextSequenceSendKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

// NextSequenceRecvPath defines the next receive sequence counter store path.
func NextSequenceRecvPath(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

// NextSequenceRecvKey returns the store key for the receive sequence of a particular
// channel binded to a specific port
func NextSequenceRecvKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

// NextSequenceAckPath defines the next acknowledgement sequence counter store path
func NextSequenceAckPath(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

// NextSequenceAckKey returns the store key for the acknowledgement sequence of
// a particular channel binded to a specific port.
func NextSequenceAckKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

// PacketCommitmentPath defines the commitments to packet data fields store path
func PacketCommitmentPath(portID, channelID string, sequence uint64) string {
	_ = "STUB: not implemented"
	return ""
}

// PacketCommitmentKey returns the store key of under which a packet commitment
// is stored
func PacketCommitmentKey(portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketCommitmentPrefixPath defines the prefix for commitments to packet data fields store path.
func PacketCommitmentPrefixPath(portID, channelID string) string {
	_ = "STUB: not implemented"
	return ""
}

// PacketAcknowledgementPath defines the packet acknowledgement store path
func PacketAcknowledgementPath(portID, channelID string, sequence uint64) string {
	_ = "STUB: not implemented"
	return ""
}

// PacketAcknowledgementKey returns the store key of under which a packet
// acknowledgement is stored
func PacketAcknowledgementKey(portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketAcknowledgementPrefixPath defines the prefix for commitments to packet data fields store path.
func PacketAcknowledgementPrefixPath(portID, channelID string) string {
	_ = "STUB: not implemented"
	return ""
}

// PacketReceiptPath defines the packet receipt store path
func PacketReceiptPath(portID, channelID string, sequence uint64) string {
	_ = "STUB: not implemented"
	return ""
}

// PacketReceiptKey returns the store key of under which a packet
// receipt is stored
func PacketReceiptKey(portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func channelPath(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

func sequencePath(sequence uint64) string { _ = "STUB: not implemented"; return "" }

// ICS05
// The following paths are the keys to the store as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-005-port-allocation#store-paths

// PortPath defines the path under which ports paths are stored on the capability module
func PortPath(portID string) string { _ = "STUB: not implemented"; return "" }
