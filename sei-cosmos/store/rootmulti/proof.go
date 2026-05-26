package rootmulti

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/merkle"
)

// RequireProof returns whether proof is required for the subpath.
func RequireProof(subpath string) bool {
	_ = "STUB: not implemented"
	// XXX: create a better convention.
	// Currently, only when query subpath is "/key", will proof be included in
	// response. If there are some changes about proof building in iavlstore.go,
	// we must change code here to keep consistency with iavlStore#Query.
	return false
}

//-----------------------------------------------------------------------------

// XXX: This should be managed by the rootMultiStore which may want to register
// more proof ops?
func DefaultProofRuntime() (prt *merkle.ProofRuntime) { _ = "STUB: not implemented"; return nil }
