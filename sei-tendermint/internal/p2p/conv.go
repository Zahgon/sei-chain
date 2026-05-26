package p2p

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/ed25519"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/conn"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/p2p/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

type NodeSecretKey ed25519.SecretKey
type NodePublicKey ed25519.PublicKey

func (k NodePublicKey) String() string { _ = "STUB: not implemented"; return "" }
func (k NodePublicKey) GoString() string {
	_ = "STUB: not implemented"

	// NodePublicKeyFromString parses a NodePublicKey from its string representation ("node:ed25519:public:hex").
	return ""
}

func NodePublicKeyFromString(s string) (NodePublicKey, error) {
	_ = "STUB: not implemented"
	return *new(NodePublicKey), nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (k NodePublicKey) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (k *NodePublicKey) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }

func (k NodeSecretKey) String() string   { _ = "STUB: not implemented"; return "" }
func (k NodeSecretKey) GoString() string { _ = "STUB: not implemented"; return "" }

type NodeChallengeSig struct {
	utils.ReadOnly
	key NodePublicKey
	sig ed25519.Signature
}

func (k NodePublicKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }
func (k NodeSecretKey) Public() NodePublicKey {
	_ = "STUB: not implemented"
	return *new(NodePublicKey)
}
func (k NodeSecretKey) SignChallenge(challenge conn.Challenge) NodeChallengeSig {
	_ = "STUB: not implemented"
	return *new(NodeChallengeSig)
}

func (s NodeChallengeSig) Key() NodePublicKey {
	_ = "STUB: not implemented"
	return *new(NodePublicKey)
}
func (s NodeChallengeSig) Verify(challenge conn.Challenge) error {
	_ = "STUB: not implemented"
	return nil
}

func (k NodePublicKey) NodeID() types.NodeID { _ = "STUB: not implemented"; return *new(types.NodeID) }

var nodePublicKeyConv = protoutils.Conv[NodePublicKey, *pb.NodePublicKey]{
	Encode: func(k NodePublicKey) *pb.NodePublicKey {
		return &pb.NodePublicKey{Ed25519: k.Bytes()}
	},
	Decode: func(p *pb.NodePublicKey) (NodePublicKey, error) {
		k, err := ed25519.PublicKeyFromBytes(p.Ed25519)
		if err != nil {
			return NodePublicKey{}, fmt.Errorf("Ed25519: %w", err)
		}
		return NodePublicKey(k), nil
	},
}

type handshakeSpec struct {
	SelfAddr          utils.Option[NodeAddress]
	PexAddrs          []NodeAddress
	SeiGigaConnection bool
}

type handshakeMsg struct {
	NodeAuth NodeChallengeSig
	handshakeSpec
}

var handshakeMsgConv = protoutils.Conv[*handshakeMsg, *pb.Handshake]{
	Encode: func(m *handshakeMsg) *pb.Handshake {
		var selfAddr *string
		if addr, ok := m.SelfAddr.Get(); ok {
			selfAddr = utils.Alloc(addr.String())
		}
		pexAddrs := make([]string, len(m.PexAddrs))
		for i, addr := range m.PexAddrs {
			pexAddrs[i] = addr.String()
		}

		return &pb.Handshake{
			NodeAuthKey:       nodePublicKeyConv.Encode(m.NodeAuth.Key()),
			NodeAuthSig:       m.NodeAuth.sig.Bytes(),
			SelfAddr:          selfAddr,
			PexAddrs:          pexAddrs,
			SeiGigaConnection: m.SeiGigaConnection,
		}
	},
	Decode: func(p *pb.Handshake) (*handshakeMsg, error) {
		nodeAuthKey, err := nodePublicKeyConv.DecodeReq(p.NodeAuthKey)
		if err != nil {
			return nil, fmt.Errorf("NodeAuthKey: %w", err)
		}
		nodeAuthSig, err := ed25519.SignatureFromBytes(p.NodeAuthSig)
		if err != nil {
			return nil, fmt.Errorf("NodeAuthSig: %w", err)
		}
		var selfAddr utils.Option[NodeAddress]
		if p.SelfAddr != nil {
			addr, err := ParseNodeAddress(*p.SelfAddr)
			if err != nil {
				return nil, fmt.Errorf("SelfAddr: %w", err)
			}
			selfAddr = utils.Some(addr)
		}
		pexAddrs := make([]NodeAddress, len(p.PexAddrs))
		for i, addrString := range p.PexAddrs {
			addr, err := ParseNodeAddress(addrString)
			if err != nil {
				return nil, fmt.Errorf("PexAddrs[%v]: %w", i, err)
			}
			pexAddrs[i] = addr
		}
		return &handshakeMsg{
			NodeAuth: NodeChallengeSig{key: nodeAuthKey, sig: nodeAuthSig},
			handshakeSpec: handshakeSpec{
				SelfAddr:          selfAddr,
				PexAddrs:          pexAddrs,
				SeiGigaConnection: p.SeiGigaConnection,
			},
		}, nil
	},
}
