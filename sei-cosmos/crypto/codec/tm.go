package codec

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	pb "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"

	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// FromTmProtoPublicKey converts a TM's pb.PublicKey into our own PubKey.
func FromTmProtoPublicKey(protoPk pb.PublicKey) (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// ToTmProtoPublicKey converts our own PubKey to TM's pb.PublicKey.
func ToTmProtoPublicKey(pk cryptotypes.PubKey) (pb.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(pb.PublicKey), nil
}

// FromTmPubKeyInterface converts TM's crypto.PubKey to our own PubKey.
func FromTmPubKeyInterface(tmPk crypto.PubKey) (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// ToTmPubKeyInterface converts our own PubKey to TM's crypto.PubKey.
func ToTmPubKeyInterface(pk cryptotypes.PubKey) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}
