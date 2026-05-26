package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

// NewClientState creates a new ClientState instance.
func NewClientState(latestSequence uint64, consensusState *ConsensusState, allowUpdateAfterProposal bool) *ClientState {
	_ = "STUB: not implemented"
	return nil
}

// ClientType is Solo Machine.
func (cs ClientState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetLatestHeight returns the latest sequence number.
// Return exported.Height to satisfy ClientState interface
// Revision number is always 0 for a solo-machine.
func (cs ClientState) GetLatestHeight() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// Status returns the status of the solo machine client.
// The client may be:
// - Active: if frozen sequence is 0
// - Frozen: otherwise solo machine is frozen
func (cs ClientState) Status(_ sdk.Context, _ sdk.KVStore, _ codec.BinaryCodec) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// Validate performs basic validation of the client state fields.
func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }

// ZeroCustomFields returns solomachine client state with client-specific fields FrozenSequence,
// and AllowUpdateAfterProposal zeroed out
func (cs ClientState) ZeroCustomFields() exported.ClientState {
	_ = "STUB: not implemented"
	return *new(exported.ClientState)
}

// Initialize will check that initial consensus state is equal to the latest consensus state of the initial client.
func (cs ClientState) Initialize(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, consState exported.ConsensusState) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportMetadata is a no-op since solomachine does not store any metadata in client store
func (cs ClientState) ExportMetadata(_ sdk.KVStore) []exported.GenesisMetadata {
	_ = "STUB: not implemented"

	// VerifyUpgradeAndUpdateState returns an error since solomachine client does not support upgrades
	return nil
}

func (cs ClientState) VerifyUpgradeAndUpdateState(
	_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore,
	_ exported.ClientState, _ exported.ConsensusState, _, _ []byte,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// VerifyClientState verifies a proof of the client state of the running chain
// stored on the solo machine.
func (cs *ClientState) VerifyClientState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	counterpartyClientIdentifier string,
	proof []byte,
	clientState exported.ClientState,
) error {
	_ = "STUB: not implemented"
	// NOTE: the proof height sequence is incremented by one due to the connection handshake verification ordering
	return nil
}

// VerifyClientConsensusState verifies a proof of the consensus state of the
// running chain stored on the solo machine.
func (cs *ClientState) VerifyClientConsensusState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	counterpartyClientIdentifier string,
	consensusHeight exported.Height,
	prefix exported.Prefix,
	proof []byte,
	consensusState exported.ConsensusState,
) error {
	_ = "STUB: not implemented"
	// NOTE: the proof height sequence is incremented by two due to the connection handshake verification ordering
	return nil
}

// VerifyConnectionState verifies a proof of the connection state of the
// specified connection end stored on the target machine.
func (cs *ClientState) VerifyConnectionState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
	connectionID string,
	connectionEnd exported.ConnectionI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyChannelState verifies a proof of the channel state of the specified
// channel end, under the specified port, stored on the target machine.
func (cs *ClientState) VerifyChannelState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	channel exported.ChannelI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketCommitment verifies a proof of an outgoing packet commitment at
// the specified port, specified channel, and specified sequence.
func (cs *ClientState) VerifyPacketCommitment(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	_ uint64,
	_ uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	packetSequence uint64,
	commitmentBytes []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketAcknowledgement verifies a proof of an incoming packet
// acknowledgement at the specified port, specified channel, and specified sequence.
func (cs *ClientState) VerifyPacketAcknowledgement(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	_ uint64,
	_ uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	packetSequence uint64,
	acknowledgement []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketReceiptAbsence verifies a proof of the absence of an
// incoming packet receipt at the specified port, specified channel, and
// specified sequence.
func (cs *ClientState) VerifyPacketReceiptAbsence(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	_ uint64,
	_ uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	packetSequence uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyNextSequenceRecv verifies a proof of the next sequence number to be
// received of the specified channel at the specified port.
func (cs *ClientState) VerifyNextSequenceRecv(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	_ uint64,
	_ uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	nextSequenceRecv uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// produceVerificationArgs perfoms the basic checks on the arguments that are
// shared between the verification functions and returns the public key of the
// consensus state, the unmarshalled proof representing the signature and timestamp
// along with the solo-machine sequence encoded in the proofHeight.
func produceVerificationArgs(
	cdc codec.BinaryCodec,
	cs *ClientState,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
) (cryptotypes.PubKey, signing.SignatureData, uint64, uint64, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), *new(signing.SignatureData), 0, 0, nil
}

// sequence is encoded in the revision height of height struct

// sets the client state to the store
func setClientState(store sdk.KVStore, cdc codec.BinaryCodec, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}
