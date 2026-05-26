package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"

	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// VerifySignature verifies if the the provided public key generated the signature
// over the given data. Single and Multi signature public keys are supported.
// The signature data type must correspond to the public key type. An error is
// returned if signature verification fails or an invalid SignatureData type is
// provided.
func VerifySignature(pubKey cryptotypes.PubKey, signBytes []byte, sigData signing.SignatureData) error {
	_ = "STUB: not implemented"
	return nil
}

// The function supplied fulfills the VerifyMultisignature interface. No special
// adjustments need to be made to the sign bytes based on the sign mode.

// MisbehaviourSignBytes returns the sign bytes for verification of misbehaviour.
func MisbehaviourSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	dataType DataType,
	data []byte,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HeaderSignBytes returns the sign bytes for verification of misbehaviour.
func HeaderSignBytes(
	cdc codec.BinaryCodec,
	header *Header,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientStateSignBytes returns the sign bytes for verification of the
// client state.
func ClientStateSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	clientState exported.ClientState,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientStateDataBytes returns the client state data bytes used in constructing
// SignBytes.
func ClientStateDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
	clientState exported.ClientState,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConsensusStateSignBytes returns the sign bytes for verification of the
// consensus state.
func ConsensusStateSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	consensusState exported.ConsensusState,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConsensusStateDataBytes returns the consensus state data bytes used in constructing
// SignBytes.
func ConsensusStateDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
	consensusState exported.ConsensusState,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionStateSignBytes returns the sign bytes for verification of the
// connection state.
func ConnectionStateSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	connectionEnd exported.ConnectionI,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionStateDataBytes returns the connection state data bytes used in constructing
// SignBytes.
func ConnectionStateDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
	connectionEnd exported.ConnectionI,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChannelStateSignBytes returns the sign bytes for verification of the
// channel state.
func ChannelStateSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	channelEnd exported.ChannelI,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChannelStateDataBytes returns the channel state data bytes used in constructing
// SignBytes.
func ChannelStateDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
	channelEnd exported.ChannelI,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketCommitmentSignBytes returns the sign bytes for verification of the
// packet commitment.
func PacketCommitmentSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	commitmentBytes []byte,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketCommitmentDataBytes returns the packet commitment data bytes used in constructing
// SignBytes.
func PacketCommitmentDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
	commitmentBytes []byte,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketAcknowledgementSignBytes returns the sign bytes for verification of
// the acknowledgement.
func PacketAcknowledgementSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	acknowledgement []byte,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketAcknowledgementDataBytes returns the packet acknowledgement data bytes used in constructing
// SignBytes.
func PacketAcknowledgementDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
	acknowledgement []byte,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketReceiptAbsenceSignBytes returns the sign bytes for verification
// of the absence of an receipt.
func PacketReceiptAbsenceSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketReceiptAbsenceDataBytes returns the packet receipt absence data bytes
// used in constructing SignBytes.
func PacketReceiptAbsenceDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextSequenceRecvSignBytes returns the sign bytes for verification of the next
// sequence to be received.
func NextSequenceRecvSignBytes(
	cdc codec.BinaryCodec,
	sequence, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	nextSequenceRecv uint64,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextSequenceRecvDataBytes returns the next sequence recv data bytes used in constructing
// SignBytes.
func NextSequenceRecvDataBytes(
	cdc codec.BinaryCodec,
	path commitmenttypes.MerklePath,
	nextSequenceRecv uint64,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
