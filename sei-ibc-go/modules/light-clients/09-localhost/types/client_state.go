package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	clienttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

// NewClientState creates a new ClientState instance
func NewClientState(chainID string, height clienttypes.Height) *ClientState {
	_ = "STUB: not implemented"
	return nil
}

// GetChainID returns an empty string
func (cs ClientState) GetChainID() string {
	_ = "STUB: not implemented"

	// ClientType is localhost.
	return ""
}

func (cs ClientState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetLatestHeight returns the latest height stored.
func (cs ClientState) GetLatestHeight() exported.Height {
	_ = "STUB: not implemented"

	// Status always returns Active. The localhost status cannot be changed.
	return *new(exported.Height)
}

func (cs ClientState) Status(_ sdk.Context, _ sdk.KVStore, _ codec.BinaryCodec,
) exported.Status {
	_ = "STUB: not implemented"
	return *

	// Validate performs a basic validation of the client state fields.
	new(exported.Status)
}

func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }

// ZeroCustomFields returns the same client state since there are no custom fields in localhost
func (cs ClientState) ZeroCustomFields() exported.ClientState {
	_ = "STUB: not implemented"

	// Initialize ensures that initial consensus state for localhost is nil
	return *new(exported.ClientState)
}

func (cs ClientState) Initialize(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, consState exported.ConsensusState) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportMetadata is a no-op for localhost client
func (cs ClientState) ExportMetadata(_ sdk.KVStore) []exported.GenesisMetadata {
	_ = "STUB: not implemented"

	// CheckHeaderAndUpdateState updates the localhost client. It only needs access to the context
	return nil
}

func (cs *ClientState) CheckHeaderAndUpdateState(
	ctx sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Header,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	// use the chain ID from context since the localhost client is from the running chain (i.e self).
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// #nosec G115 --- overflow checked above

// CheckMisbehaviourAndUpdateState implements ClientState
// Since localhost is the client of the running chain, misbehaviour cannot be submitted to it
// Thus, CheckMisbehaviourAndUpdateState returns an error for localhost
func (cs ClientState) CheckMisbehaviourAndUpdateState(
	_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Misbehaviour,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// CheckSubstituteAndUpdateState returns an error. The localhost cannot be modified by
// proposals.
func (cs ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context, _ codec.BinaryCodec, _, _ sdk.KVStore,
	_ exported.ClientState,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// VerifyUpgradeAndUpdateState returns an error since localhost cannot be upgraded
func (cs ClientState) VerifyUpgradeAndUpdateState(
	_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore,
	_ exported.ClientState, _ exported.ConsensusState, _, _ []byte,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// VerifyClientState verifies that the localhost client state is stored locally
func (cs ClientState) VerifyClientState(
	store sdk.KVStore, cdc codec.BinaryCodec,
	_ exported.Height, _ exported.Prefix, _ string, _ []byte, clientState exported.ClientState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyClientConsensusState returns nil since a local host client does not store consensus
// states.
func (cs ClientState) VerifyClientConsensusState(
	sdk.KVStore, codec.BinaryCodec,
	exported.Height, string, exported.Height, exported.Prefix,
	[]byte, exported.ConsensusState,
) error {
	_ = "STUB: not implemented"

	// VerifyConnectionState verifies a proof of the connection state of the
	// specified connection end stored locally.
	return nil
}

func (cs ClientState) VerifyConnectionState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	_ exported.Height,
	_ exported.Prefix,
	_ []byte,
	connectionID string,
	connectionEnd exported.ConnectionI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyChannelState verifies a proof of the channel state of the specified
// channel end, under the specified port, stored on the local machine.
func (cs ClientState) VerifyChannelState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	_ exported.Height,
	prefix exported.Prefix,
	_ []byte,
	portID,
	channelID string,
	channel exported.ChannelI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketCommitment verifies a proof of an outgoing packet commitment at
// the specified port, specified channel, and specified sequence.
func (cs ClientState) VerifyPacketCommitment(
	ctx sdk.Context,
	store sdk.KVStore,
	_ codec.BinaryCodec,
	_ exported.Height,
	_ uint64,
	_ uint64,
	_ exported.Prefix,
	_ []byte,
	portID,
	channelID string,
	sequence uint64,
	commitmentBytes []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketAcknowledgement verifies a proof of an incoming packet
// acknowledgement at the specified port, specified channel, and specified sequence.
func (cs ClientState) VerifyPacketAcknowledgement(
	ctx sdk.Context,
	store sdk.KVStore,
	_ codec.BinaryCodec,
	_ exported.Height,
	_ uint64,
	_ uint64,
	_ exported.Prefix,
	_ []byte,
	portID,
	channelID string,
	sequence uint64,
	acknowledgement []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketReceiptAbsence verifies a proof of the absence of an
// incoming packet receipt at the specified port, specified channel, and
// specified sequence.
func (cs ClientState) VerifyPacketReceiptAbsence(
	ctx sdk.Context,
	store sdk.KVStore,
	_ codec.BinaryCodec,
	_ exported.Height,
	_ uint64,
	_ uint64,
	_ exported.Prefix,
	_ []byte,
	portID,
	channelID string,
	sequence uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyNextSequenceRecv verifies a proof of the next sequence number to be
// received of the specified channel at the specified port.
func (cs ClientState) VerifyNextSequenceRecv(
	ctx sdk.Context,
	store sdk.KVStore,
	_ codec.BinaryCodec,
	_ exported.Height,
	_ uint64,
	_ uint64,
	_ exported.Prefix,
	_ []byte,
	portID,
	channelID string,
	nextSequenceRecv uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}
