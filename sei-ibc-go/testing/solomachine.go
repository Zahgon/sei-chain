package ibctesting

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"

	commitmenttypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/23-commitment/types"
	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
	solomachinetypes "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/light-clients/06-solomachine/types"
)

// Solomachine is a testing helper used to simulate a counterparty
// solo machine client.
type Solomachine struct {
	t *testing.T

	cdc         codec.BinaryCodec
	ClientID    string
	PrivateKeys []cryptotypes.PrivKey // keys used for signing
	PublicKeys  []cryptotypes.PubKey  // keys used for generating solo machine pub key
	PublicKey   cryptotypes.PubKey    // key used for verification
	Sequence    uint64
	Time        uint64
	Diversifier string
}

// NewSolomachine returns a new solomachine instance with an `nKeys` amount of
// generated private/public key pairs and a sequence starting at 1. If nKeys
// is greater than 1 then a multisig public key is used.
func NewSolomachine(t *testing.T, cdc codec.BinaryCodec, clientID, diversifier string, nKeys uint64) *Solomachine {
	_ = "STUB: not implemented"
	return nil
}

// GenerateKeys generates a new set of secp256k1 private keys and public keys.
// If the number of keys is greater than one then the public key returned represents
// a multisig public key. The private keys are used for signing, the public
// keys are used for generating the public key and the public key is used for
// solo machine verification. The usage of secp256k1 is entirely arbitrary.
// The key type can be swapped for any key type supported by the PublicKey
// interface, if needed. The same is true for the amino based Multisignature
// public key.
func GenerateKeys(t *testing.T, n uint64) ([]cryptotypes.PrivKey, []cryptotypes.PubKey, cryptotypes.PubKey) {
	_ = "STUB: not implemented"
	return nil, nil, *new(cryptotypes.PubKey)
}

// generate multi sig pk
// #nosec G115 --- checked on top

// ClientState returns a new solo machine ClientState instance. Default usage does not allow update
// after governance proposal
func (solo *Solomachine) ClientState() *solomachinetypes.ClientState {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusState returns a new solo machine ConsensusState instance
func (solo *Solomachine) ConsensusState() *solomachinetypes.ConsensusState {
	_ = "STUB: not implemented"
	return nil
}

// GetHeight returns an exported.Height with Sequence as RevisionHeight
func (solo *Solomachine) GetHeight() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// CreateHeader generates a new private/public key pair and creates the
// necessary signature to construct a valid solo machine header.
func (solo *Solomachine) CreateHeader() *solomachinetypes.Header {
	_ = "STUB: not implemented"
	// generate new private keys and signature for header
	return nil
}

// assumes successful header update

// CreateMisbehaviour constructs testing misbehaviour for the solo machine client
// by signing over two different data bytes at the same sequence.
func (solo *Solomachine) CreateMisbehaviour() *solomachinetypes.Misbehaviour {
	_ = "STUB: not implemented"
	return nil
}

// misbehaviour signaturess can have different timestamps

// GenerateSignature uses the stored private keys to generate a signature
// over the sign bytes with each key. If the amount of keys is greater than
// 1 then a multisig data type is returned.
func (solo *Solomachine) GenerateSignature(signBytes []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// single public key

// generate multi signature data

// GetClientStatePath returns the commitment path for the client state.
func (solo *Solomachine) GetClientStatePath(counterpartyClientIdentifier string) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}

// GetConsensusStatePath returns the commitment path for the consensus state.
func (solo *Solomachine) GetConsensusStatePath(counterpartyClientIdentifier string, consensusHeight exported.Height) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}

// GetConnectionStatePath returns the commitment path for the connection state.
func (solo *Solomachine) GetConnectionStatePath(connID string) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}

// GetChannelStatePath returns the commitment path for that channel state.
func (solo *Solomachine) GetChannelStatePath(portID, channelID string) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}

// GetPacketCommitmentPath returns the commitment path for a packet commitment.
func (solo *Solomachine) GetPacketCommitmentPath(portID, channelID string) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}

// GetPacketAcknowledgementPath returns the commitment path for a packet acknowledgement.
func (solo *Solomachine) GetPacketAcknowledgementPath(portID, channelID string) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}

// GetPacketReceiptPath returns the commitment path for a packet receipt
// and an absent receipts.
func (solo *Solomachine) GetPacketReceiptPath(portID, channelID string) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}

// GetNextSequenceRecvPath returns the commitment path for the next sequence recv counter.
func (solo *Solomachine) GetNextSequenceRecvPath(portID, channelID string) commitmenttypes.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypes.MerklePath)
}
