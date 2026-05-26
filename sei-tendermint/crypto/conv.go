package crypto

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/jsontypes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

func init() {
	jsontypes.MustRegister((*pb.PublicKey)(nil))
	jsontypes.MustRegister((*pb.PublicKey_Ed25519)(nil))
}

var PubKeyConv = utils.ProtoConv[PubKey, *pb.PublicKey]{
	Encode: func(k PubKey) *pb.PublicKey { return utils.Alloc(PubKeyToProto(k)) },
	Decode: func(x *pb.PublicKey) (PubKey, error) { return PubKeyFromProto(*x) },
}

// PubKeyToProto takes crypto.PubKey and transforms it to a protobuf Pubkey
func PubKeyToProto(k PubKey) pb.PublicKey { _ = "STUB: not implemented"; return *new(pb.PublicKey) }

// PubKeyFromProto takes a protobuf Pubkey and transforms it to a crypto.Pubkey
func PubKeyFromProto(k pb.PublicKey) (PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}
