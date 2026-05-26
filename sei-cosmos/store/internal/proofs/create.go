package proofs

import (
	"errors"

	ics23 "github.com/confio/ics23/go"
)

var (
	ErrEmptyKey       = errors.New("key is empty")
	ErrEmptyKeyInData = errors.New("data contains empty key")
)

// TendermintSpec constrains the format from ics23-tendermint (crypto/merkle SimpleProof)
var TendermintSpec = &ics23.ProofSpec{
	LeafSpec: &ics23.LeafOp{
		Prefix:       []byte{0},
		Hash:         ics23.HashOp_SHA256,
		PrehashValue: ics23.HashOp_SHA256,
		Length:       ics23.LengthOp_VAR_PROTO,
	},
	InnerSpec: &ics23.InnerSpec{
		ChildOrder:      []int32{0, 1},
		MinPrefixLength: 1,
		MaxPrefixLength: 1,  // fixed prefix + one child
		ChildSize:       32, // (no length byte)
		Hash:            ics23.HashOp_SHA256,
	},
}

/*
CreateMembershipProof will produce a CommitmentProof that the given key (and queries value) exists in the iavl tree.
If the key doesn't exist in the tree, this will return an error.
*/
func CreateMembershipProof(data map[string][]byte, key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
CreateNonMembershipProof will produce a CommitmentProof that the given key doesn't exist in the iavl tree.
If the key exists in the tree, this will return an error.
*/
func CreateNonMembershipProof(data map[string][]byte, key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ensure this key is not in the store

// include left proof unless key is left of entire map

// include right proof unless key is right of entire map

func createExistenceProof(data map[string][]byte, key []byte) (*ics23.ExistenceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
