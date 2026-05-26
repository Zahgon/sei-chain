package ibctesting

import (

	//	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/04-channel/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// Endpoint is a which represents a channel endpoint and its associated
// client and connections. It contains client, connection, and channel
// configuration parameters. Endpoint functions will utilize the parameters
// set in the configuration structs when executing IBC messages.
type Endpoint struct {
	Chain        *TestChain
	Counterparty *Endpoint
	ClientID     string
	ConnectionID string
	ChannelID    string

	ClientConfig     ClientConfig
	ConnectionConfig *ConnectionConfig
	ChannelConfig    *ChannelConfig
}

// NewEndpoint constructs a new endpoint without the counterparty.
// CONTRACT: the counterparty endpoint must be set by the caller.
func NewEndpoint(
	chain *TestChain, clientConfig ClientConfig,
	connectionConfig *ConnectionConfig, channelConfig *ChannelConfig,
) *Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultEndpoint constructs a new endpoint using default values.
// CONTRACT: the counterparty endpoitn must be set by the caller.
func NewDefaultEndpoint(chain *TestChain) *Endpoint { _ = "STUB: not implemented"; return nil }

// QueryProof queries proof associated with this endpoint using the lastest client state
// height on the counterparty chain.
func (endpoint *Endpoint) QueryProof(key []byte) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	// obtain the counterparty client representing the chain associated with the endpoint
	return nil, *new(clienttypes.Height)
}

// query proof on the counterparty using the latest height of the IBC client

// QueryProofAtHeight queries proof associated with this endpoint using the proof height
// providied
func (endpoint *Endpoint) QueryProofAtHeight(key []byte, height uint64) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	// query proof on the counterparty using the latest height of the IBC client
	return nil, *new(clienttypes.Height)
}

// #nosec G115 -- height is bounds checked above

// CreateClient creates an IBC client on the endpoint. It will update the
// clientID for the endpoint if the message is successfully executed.
// NOTE: a solo machine client will be created with an empty diversifier.
func (endpoint *Endpoint) CreateClient() (err error) {
	_ = "STUB: not implemented"
	// ensure counterparty has committed state
	return nil
}

// TODO
//		solo := NewSolomachine(chain.t, endpoint.Chain.Codec, clientID, "", 1)
//		clientState = solo.ClientState()
//		consensusState = solo.ConsensusState()

// UpdateClient updates the IBC client associated with the endpoint.
func (endpoint *Endpoint) UpdateClient() (err error) {
	_ = "STUB: not implemented"
	// ensure counterparty has committed state
	return nil
}

// ConnOpenInit will construct and execute a MsgConnectionOpenInit on the associated endpoint.
func (endpoint *Endpoint) ConnOpenInit() error { _ = "STUB: not implemented"; return nil }

// ConnOpenTry will construct and execute a MsgConnectionOpenTry on the associated endpoint.
func (endpoint *Endpoint) ConnOpenTry() error { _ = "STUB: not implemented"; return nil }

// does not support handshake continuation

// ConnOpenAck will construct and execute a MsgConnectionOpenAck on the associated endpoint.
func (endpoint *Endpoint) ConnOpenAck() error { _ = "STUB: not implemented"; return nil }

// testing doesn't use flexible selection

// ConnOpenConfirm will construct and execute a MsgConnectionOpenConfirm on the associated endpoint.
func (endpoint *Endpoint) ConnOpenConfirm() error { _ = "STUB: not implemented"; return nil }

// QueryConnectionHandshakeProof returns all the proofs necessary to execute OpenTry or Open Ack of
// the connection handshakes. It returns the counterparty client state, proof of the counterparty
// client state, proof of the counterparty consensus state, the consensus state height, proof of
// the counterparty connection, and the proof height for all the proofs returned.
func (endpoint *Endpoint) QueryConnectionHandshakeProof() (
	clientState exported.ClientState, proofClient,
	proofConsensus []byte, consensusHeight clienttypes.Height,
	proofConnection []byte, proofHeight clienttypes.Height,
) {
	_ = "STUB: not implemented"
	// obtain the client state on the counterparty chain
	return *new(exported.ClientState), nil, nil, *new(clienttypes.Height), nil, *new(clienttypes.Height)
}

// query proof for the client state on the counterparty

// query proof for the consensus state on the counterparty

// query proof for the connection on the counterparty

// ChanOpenInit will construct and execute a MsgChannelOpenInit on the associated endpoint.
func (endpoint *Endpoint) ChanOpenInit() error { _ = "STUB: not implemented"; return nil }

// ChanOpenTry will construct and execute a MsgChannelOpenTry on the associated endpoint.
func (endpoint *Endpoint) ChanOpenTry() error { _ = "STUB: not implemented"; return nil }

// does not support handshake continuation

// ChanOpenAck will construct and execute a MsgChannelOpenAck on the associated endpoint.
func (endpoint *Endpoint) ChanOpenAck() error { _ = "STUB: not implemented"; return nil }

// testing doesn't use flexible selection

// ChanOpenConfirm will construct and execute a MsgChannelOpenConfirm on the associated endpoint.
func (endpoint *Endpoint) ChanOpenConfirm() error { _ = "STUB: not implemented"; return nil }

// ChanCloseInit will construct and execute a MsgChannelCloseInit on the associated endpoint.
//
// NOTE: does not work with ibc-transfer module
func (endpoint *Endpoint) ChanCloseInit() error { _ = "STUB: not implemented"; return nil }

// ChanCloseConfirm will construct and execute a NewMsgChannelCloseConfirm on the associated endpoint.
func (endpoint *Endpoint) ChanCloseConfirm() error { _ = "STUB: not implemented"; return nil }

// SendPacket sends a packet through the channel keeper using the associated endpoint
// The counterparty client is updated so proofs can be sent to the counterparty chain.
func (endpoint *Endpoint) SendPacket(packet exported.PacketI) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to send message, acting as a module

// commit changes since no message was sent

// RecvPacket receives a packet on the associated endpoint.
// The counterparty client is updated.
func (endpoint *Endpoint) RecvPacket(packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	// get proof of packet commitment on source
	return nil
}

// receive on counterparty and update source client

// WriteAcknowledgement writes an acknowledgement on the channel associated with the endpoint.
// The counterparty client is updated.
func (endpoint *Endpoint) WriteAcknowledgement(ack exported.Acknowledgement, packet exported.PacketI) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to send message, acting as a handler

// commit changes since no message was sent

// AcknowledgePacket sends a MsgAcknowledgement to the channel associated with the endpoint.
func (endpoint *Endpoint) AcknowledgePacket(packet channeltypes.Packet, ack []byte) error {
	_ = "STUB: not implemented"
	// get proof of acknowledgement on counterparty
	return nil
}

// TimeoutPacket sends a MsgTimeout to the channel associated with the endpoint.
func (endpoint *Endpoint) TimeoutPacket(packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	// get proof for timeout based on channel order
	return nil
}

// SetChannelClosed sets a channel state to CLOSED.
func (endpoint *Endpoint) SetChannelClosed() error { _ = "STUB: not implemented"; return nil }

// GetClientState retrieves the Client State for this endpoint. The
// client state is expected to exist otherwise testing will fail.
func (endpoint *Endpoint) GetClientState() exported.ClientState {
	_ = "STUB: not implemented"
	return *new(exported.ClientState)
}

// SetClientState sets the client state for this endpoint.
func (endpoint *Endpoint) SetClientState(clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// GetConsensusState retrieves the Consensus State for this endpoint at the provided height.
// The consensus state is expected to exist otherwise testing will fail.
func (endpoint *Endpoint) GetConsensusState(height exported.Height) exported.ConsensusState {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState)
}

// SetConsensusState sets the consensus state for this endpoint.
func (endpoint *Endpoint) SetConsensusState(consensusState exported.ConsensusState, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// GetConnection retrieves an IBC Connection for the endpoint. The
// connection is expected to exist otherwise testing will fail.
func (endpoint *Endpoint) GetConnection() connectiontypes.ConnectionEnd {
	_ = "STUB: not implemented"
	return *new(connectiontypes.ConnectionEnd)
}

// SetConnection sets the connection for this endpoint.
func (endpoint *Endpoint) SetConnection(connection connectiontypes.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}

// GetChannel retrieves an IBC Channel for the endpoint. The channel
// is expected to exist otherwise testing will fail.
func (endpoint *Endpoint) GetChannel() channeltypes.Channel {
	_ = "STUB: not implemented"
	return *new(channeltypes.Channel)
}

// SetChannel sets the channel for this endpoint.
func (endpoint *Endpoint) SetChannel(channel channeltypes.Channel) {
	_ = "STUB: not implemented"
	return
}

// QueryClientStateProof performs and abci query for a client stat associated
// with this endpoint and returns the ClientState along with the proof.
func (endpoint *Endpoint) QueryClientStateProof() (exported.ClientState, []byte) {
	_ = "STUB: not implemented"
	// retrieve client state to provide proof for
	return *new(exported.ClientState), nil
}
