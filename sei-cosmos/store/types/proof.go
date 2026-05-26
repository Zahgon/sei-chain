package types

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/merkle"
	tmmerkle "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

const (
	ProofOpIAVLCommitment         = "ics23:iavl"
	ProofOpSimpleMerkleCommitment = "ics23:simple"
)

// CommitmentOp implements merkle.ProofOperator by wrapping an ics23 CommitmentProof
// It also contains a Key field to determine which key the proof is proving.
// NOTE: CommitmentProof currently can either be ExistenceProof or NonexistenceProof
//
// Type and Spec are classified by the kind of merkle proof it represents allowing
// the code to be reused by more types. Spec is never on the wire, but mapped from type in the code.
type CommitmentOp struct {
	Type  string
	Spec  *ics23.ProofSpec
	Key   []byte
	Proof *ics23.CommitmentProof
}

var _ merkle.ProofOperator = CommitmentOp{}

func NewIavlCommitmentOp(key []byte, proof *ics23.CommitmentProof) CommitmentOp {
	_ = "STUB: not implemented"
	return *new(CommitmentOp)
}

func NewSimpleMerkleCommitmentOp(key []byte, proof *ics23.CommitmentProof) CommitmentOp {
	_ = "STUB: not implemented"
	return *new(CommitmentOp)
}

// CommitmentOpDecoder takes a merkle.ProofOp and attempts to decode it into a CommitmentOp ProofOperator
// The proofOp.Data is just a marshalled CommitmentProof. The Key of the CommitmentOp is extracted
// from the unmarshalled proof.
func CommitmentOpDecoder(pop tmmerkle.ProofOp) (merkle.ProofOperator, error) {
	_ = "STUB: not implemented"
	return *new(merkle.ProofOperator), nil
}

func (op CommitmentOp) GetKey() []byte {
	_ = "STUB: not implemented"

	// Run takes in a list of arguments and attempts to run the proof op against these arguments.
	// Returns the root wrapped in [][]byte if the proof op succeeds with given args. If not,
	// it will return an error.
	//
	// CommitmentOp will accept args of length 1 or length 0
	// If length 1 args is passed in, then CommitmentOp will attempt to prove the existence of the key
	// with the value provided by args[0] using the embedded CommitmentProof and return the CommitmentRoot of the proof.
	// If length 0 args is passed in, then CommitmentOp will attempt to prove the absence of the key
	// in the CommitmentOp and return the CommitmentRoot of the proof.
	return nil
}

func (op CommitmentOp) Run(args [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	// calculate root from proof
	return nil, nil
}

// Only support an existence proof or nonexistence proof (batch proofs currently unsupported)

// Args are nil, so we verify the absence of the key.

// Args is length 1, verify existence of key with value args[0]

// ProofOp implements ProofOperator interface and converts a CommitmentOp
// into a merkle.ProofOp format that can later be decoded by CommitmentOpDecoder
// back into a CommitmentOp for proof verification
func (op CommitmentOp) ProofOp() tmmerkle.ProofOp {
	_ = "STUB: not implemented"
	return *new(tmmerkle.ProofOp)
}
