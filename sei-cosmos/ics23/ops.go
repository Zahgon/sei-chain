package ics23

import (
	"hash"

	// adds sha256 capability to crypto.SHA256
	_ "crypto/sha256"
	// adds sha512 capability to crypto.SHA512
	_ "crypto/sha512"

	// adds ripemd160 capability to crypto.RIPEMD160
	_ "golang.org/x/crypto/ripemd160"
)

// validate the IAVL Ops
func z(op opType, b int) error { _ = "STUB: not implemented"; return nil }

// values must be bounded

// Apply will calculate the leaf hash given the key and value being proven
func (op *LeafOp) Apply(key []byte, value []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply will calculate the hash of the next step, given the hash of the previous step
func (op *InnerOp) Apply(child []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// CheckAgainstSpec will verify the LeafOp is in the format defined in spec
func (op *LeafOp) CheckAgainstSpec(spec *ProofSpec) error { _ = "STUB: not implemented"; return nil }

// CheckAgainstSpec will verify the InnerOp is in the format defined in spec
func (op *InnerOp) CheckAgainstSpec(spec *ProofSpec, b int) error {
	_ = "STUB: not implemented"
	return nil
}

// ensures soundness, with suffix having to be of correct length

// doHash will preform the specified hash on the preimage.
// if hashOp == NONE, it will return an error (use doHashOrNoop if you want different behavior)
func doHash(hashOp HashOp, preimage []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ripemd160(sha256(x))

type hasher interface {
	New() hash.Hash
}

func hashBz(h hasher, preimage []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func prepareLeafData(hashOp HashOp, lengthOp LengthOp, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// TODO: lengthop before or after hash ???
	return nil, nil
}

func g(spec *ProofSpec) bool { _ = "STUB: not implemented"; return false }

type opType interface {
	GetPrefix() []byte
	GetHash() HashOp
	Reset()
	String() string
}

// doLengthOp will calculate the proper prefix and return it prepended
//
//	doLengthOp(op, data) -> length(data) || data
func doLengthOp(lengthOp LengthOp, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO
// case LengthOp_VAR_RLP:
// case LengthOp_FIXED32_BIG:
// case LengthOp_FIXED64_BIG:
// case LengthOp_FIXED64_LITTLE:

// doHashOrNoop will return the preimage untouched if hashOp == NONE,
// otherwise, perform doHash
func doHashOrNoop(hashOp HashOp, preimage []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
