package types

import (
	"encoding"
	"errors"
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/ed25519"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/hashable"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

var autobahnTag = utils.OrPanic1(ed25519.NewTag("sei_giga_autobahn"))

// Msg is the interface for all messages signable by a stream node.
type Msg interface{ asMsg() *pb.Msg }

func (m *LaneProposal) asMsg() *pb.Msg { _ = "STUB: not implemented"; return nil }

func (m *LaneVote) asMsg() *pb.Msg { _ = "STUB: not implemented"; return nil }

func (m *AppVote) asMsg() *pb.Msg { _ = "STUB: not implemented"; return nil }

func (m *Proposal) asMsg() *pb.Msg { _ = "STUB: not implemented"; return nil }

func (m *PrepareVote) asMsg() *pb.Msg { _ = "STUB: not implemented"; return nil }

func (m *CommitVote) asMsg() *pb.Msg { _ = "STUB: not implemented"; return nil }

func (m *TimeoutVote) asMsg() *pb.Msg { _ = "STUB: not implemented"; return nil }

// Hash is the hash of a message.
type Hash[T Msg] hashable.Hash[*pb.Msg]

// Hashed is a message with its hash.
type Hashed[T Msg] struct {
	utils.ReadOnly
	msg  T
	hash Hash[T]
}

// Msg returns the message.
func (m *Hashed[T]) Msg() T {
	_ = "STUB: not implemented"

	// Hash returns the hash.
	return *new(T)
}

func (m *Hashed[T]) Hash() Hash[T] {
	_ = "STUB: not implemented"

	// NewHashed creates a new Hashed message.
	return nil
}

func NewHashed[T Msg](msg T) *Hashed[T] { _ = "STUB: not implemented"; return nil }

// SecretKey is the secret key of the validator.
type SecretKey struct{ key ed25519.SecretKey }

// SecretKeyFromED25519 constructs a SecretKey from an ed25519.SecretKey.
func SecretKeyFromED25519(key ed25519.SecretKey) SecretKey {
	_ = "STUB: not implemented"
	return *new(SecretKey)
}

// Public returns the public key corresponding to the secret key.
func (k SecretKey) Public() PublicKey { _ = "STUB: not implemented"; return *new(PublicKey) }

// PublicKey is the public key of the validator.
// nolint:recvcheck
type PublicKey struct {
	utils.ReadOnly
	key ed25519.PublicKey
}

// Compare implements Comparable.
func (k PublicKey) Compare(other PublicKey) int { _ = "STUB: not implemented"; return 0 }

// Bytes converts the public key to bytes.
func (k PublicKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// PublicKeyFromBytes constructs a public key from bytes.
func PublicKeyFromBytes(b []byte) (PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(PublicKey), nil
}

// String returns a string representation.
func (k PublicKey) String() string { _ = "STUB: not implemented"; return "" }

// PublicKeyFromString constructs a public key from a string representation.
func PublicKeyFromString(s string) (PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(PublicKey), nil
}

// GoString returns a strings representation.
func (k PublicKey) GoString() string {
	_ = "STUB: not implemented"

	// MarshalText implements the encoding.TextMarshaler interface.
	return ""
}

func (k PublicKey) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (k *PublicKey) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }

var _ encoding.TextMarshaler = PublicKey{}
var _ encoding.TextUnmarshaler = (*PublicKey)(nil)

// String returns a log-safe representation of the secret key.
func (k SecretKey) String() string { _ = "STUB: not implemented"; return "" }

// GoString returns a log-safe representation of the secret key.
func (k SecretKey) GoString() string {
	_ = "STUB: not implemented"

	// Sign signs a message.
	return ""
}

func Sign[T Msg](key SecretKey, msg T) *Signed[T] { _ = "STUB: not implemented"; return nil }

// Signature represents a signature on the consensus message.
type Signature struct {
	utils.ReadOnly
	key PublicKey
	sig ed25519.Signature
}

// Signed is a hashed message with its signature.
type Signed[T Msg] struct {
	utils.ReadOnly
	hashed *Hashed[T]
	sig    *Signature
}

// Msg returns the message.
func (m *Signed[T]) Msg() T {
	_ = "STUB: not implemented"

	// Hash returns the hash of the message.
	return *new(T)
}

func (m *Signed[T]) Hash() Hash[T] { _ = "STUB: not implemented"; return nil }

// Sig returns the signature of the message.
func (m *Signed[T]) Sig() *Signature {
	_ = "STUB: not implemented"

	// Key returns the key whish signed the message.
	return nil
}

func (m *Signed[T]) Key() PublicKey {
	_ = "STUB: not implemented"

	// VerifySig verifies the signature of the message.
	return *new(PublicKey)
}

func (m *Signed[T]) VerifySig(c *Committee) error { _ = "STUB: not implemented"; return nil }

// verifyQC verifies a slice of signatures and checks if they form a quorum.
func (m *Hashed[T]) verifyQC(c *Committee, quorum int, sigs []*Signature) error {
	_ = "STUB: not implemented"
	return nil
}

// PublicKeyConv is a protobuf converter for PublicKey.
var PublicKeyConv = protoutils.Conv[PublicKey, *pb.PublicKey]{
	Encode: func(k PublicKey) *pb.PublicKey {
		return &pb.PublicKey{
			Ed25519: k.Bytes(),
		}
	},
	Decode: func(p *pb.PublicKey) (PublicKey, error) {
		key, err := PublicKeyFromBytes(p.Ed25519)
		if err != nil {
			return PublicKey{}, err
		}
		return key, nil
	},
}

// SignatureConv is a protobuf converter for Signature.
var SignatureConv = protoutils.Conv[*Signature, *pb.Signature]{
	Encode: func(s *Signature) *pb.Signature {
		return &pb.Signature{
			Key: PublicKeyConv.Encode(s.key),
			Sig: s.sig.Bytes(),
		}
	},
	Decode: func(p *pb.Signature) (*Signature, error) {
		key, err := PublicKeyConv.DecodeReq(p.Key)
		if err != nil {
			return nil, fmt.Errorf("key: %w", err)
		}
		sig, err := ed25519.SignatureFromBytes(p.Sig)
		if err != nil {
			return nil, fmt.Errorf("sig: %w", err)
		}
		return &Signature{key: key, sig: sig}, nil
	},
}

// MsgConv is a protobuf converter for Msg.
var MsgConv = protoutils.Conv[Msg, *pb.Msg]{
	Encode: func(m Msg) *pb.Msg {
		return m.asMsg()
	},
	Decode: func(m *pb.Msg) (Msg, error) {
		if m.T == nil {
			return nil, errors.New("empty")
		}
		switch t := m.T.(type) {
		case *pb.Msg_LaneProposal:
			return LaneProposalConv.DecodeReq(t.LaneProposal)
		case *pb.Msg_LaneVote:
			return LaneVoteConv.DecodeReq(t.LaneVote)
		case *pb.Msg_Proposal:
			return ProposalConv.DecodeReq(t.Proposal)
		case *pb.Msg_PrepareVote:
			return PrepareVoteConv.DecodeReq(t.PrepareVote)
		case *pb.Msg_CommitVote:
			return CommitVoteConv.DecodeReq(t.CommitVote)
		case *pb.Msg_TimeoutVote:
			return TimeoutVoteConv.DecodeReq(t.TimeoutVote)
		case *pb.Msg_AppVote:
			return AppVoteConv.DecodeReq(t.AppVote)
		default:
			return nil, fmt.Errorf("unknown Msg type: %T", t)
		}
	},
}

// AsMsg casts a hashed message to Hashed[Msg].
func (m *Hashed[T]) AsMsg() *Hashed[Msg] { _ = "STUB: not implemented"; return nil }

// AsMsg casts a signed message to Signed[Msg].
func (m *Signed[T]) AsMsg() *Signed[Msg] { _ = "STUB: not implemented"; return nil }

// HashedCast PANICS if msg.Msg() is not of type T.
func HashedCastOrPanic[T Msg](msg *Hashed[Msg]) *Hashed[T] { _ = "STUB: not implemented"; return nil }

// SignedCast PANICS if msg.Msg() is not of type T.
func SignedCastOrPanic[T Msg](msg *Signed[Msg]) *Signed[T] { _ = "STUB: not implemented"; return nil }

// SignedMsgConv is a protobuf converter for Signed[Msg].
func SignedMsgConv[T Msg]() *protoutils.Conv[*Signed[T], *pb.SignedMsg] {
	_ = "STUB: not implemented"
	return nil
}
