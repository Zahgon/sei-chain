package ics23

// IsCompressed returns true if the proof was compressed
func IsCompressed(proof *CommitmentProof) bool { _ = "STUB: not implemented"; return false }

// Compress will return a CompressedBatchProof if the input is BatchProof
// Otherwise it will return the input.
// This is safe to call multiple times (idempotent)
func Compress(proof *CommitmentProof) *CommitmentProof { _ = "STUB: not implemented"; return nil }

// Decompress will return a BatchProof if the input is CompressedBatchProof
// Otherwise it will return the input.
// This is safe to call multiple times (idempotent)
func Decompress(proof *CommitmentProof) *CommitmentProof { _ = "STUB: not implemented"; return nil }

func compress(batch *BatchProof) *CompressedBatchProof { _ = "STUB: not implemented"; return nil }

func compressEntry(entry *BatchEntry, lookup *[]*InnerOp, registry map[string]int32) *CompressedBatchEntry {
	_ = "STUB: not implemented"
	return nil
}

func compressExist(exist *ExistenceProof, lookup *[]*InnerOp, registry map[string]int32) *CompressedExistenceProof {
	_ = "STUB: not implemented"
	return nil
}

func compressStep(step *InnerOp, lookup *[]*InnerOp, registry map[string]int32) int32 {
	_ = "STUB: not implemented"
	return 0
}

// load from cache if there

// create new step if not there

func decompress(comp *CompressedBatchProof) *BatchProof { _ = "STUB: not implemented"; return nil }

// TendermintSpec constrains the format from proofs-tendermint (crypto/merkle SimpleProof)
var TendermintSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		PrehashKey:   HashOp_NO_HASH,
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_VAR_PROTO,
	},
	InnerSpec: &InnerSpec{
		ChildOrder:      []int32{0, 1},
		MinPrefixLength: 1,
		MaxPrefixLength: 1,
		ChildSize:       32, // (no length byte)
		Hash:            HashOp_SHA256,
	},
}

func decompressExist(exist *CompressedExistenceProof, lookup []*InnerOp) *ExistenceProof {
	_ = "STUB: not implemented"
	return nil
}
