package proofs

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

// ConvertExistenceProof will convert the given proof into a valid
// existence proof, if that's what it is.
//
// This is the simplest case of the range proof and we will focus on
// demoing compatibility here
func ConvertExistenceProof(p *crypto.Proof, key, value []byte) (*ics23.ExistenceProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this is adapted from merkle/hash.go:leafHash()
// and merkle/simple_map.go:KVPair.Bytes()
func convertLeafOp() *ics23.LeafOp { _ = "STUB: not implemented"; return nil }

func convertInnerOps(p *crypto.Proof) ([]*ics23.InnerOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// combine with: 0x01 || lefthash || righthash

// buildPath returns a list of steps from leaf to root
// in each step, true means index is left side, false index is right side
// code adapted from merkle/simple_proof.go:computeHashFromAunts
func buildPath(idx, total int64) []bool { _ = "STUB: not implemented"; return nil }

// we put goLeft at the end of the array, as we recurse from top to bottom,
// and want the leaf to be first in array, root last

func getSplitPoint(length int64) int64 { _ = "STUB: not implemented"; return 0 }

//#nosec G115 -- length is validated positive above

//#nosec G115 -- bitlen derived from a positive int64, shift result fits in int64
