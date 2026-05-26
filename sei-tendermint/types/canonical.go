package types

import (
	"time"

	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

// Canonical* wraps the structs in types for amino encoding them for use in SignBytes / the Signable interface.

// TimeFormat is used for generating the sigs
const TimeFormat = time.RFC3339Nano

//-----------------------------------
// Canonicalize the structs

func CanonicalizeBlockID(bid tmproto.BlockID) *tmproto.CanonicalBlockID {
	_ = "STUB: not implemented"
	return nil
}

// CanonicalizeVote transforms the given PartSetHeader to a CanonicalPartSetHeader.
func CanonicalizePartSetHeader(psh tmproto.PartSetHeader) tmproto.CanonicalPartSetHeader {
	_ = "STUB: not implemented"
	return *new(tmproto.CanonicalPartSetHeader)
}

// CanonicalizeVote transforms the given Proposal to a CanonicalProposal.
func CanonicalizeProposal(chainID string, proposal *tmproto.Proposal) tmproto.CanonicalProposal {
	_ = "STUB: not implemented"
	return *new(tmproto.CanonicalProposal)
}

// encoded as sfixed64
// encoded as sfixed64

// CanonicalizeVote transforms the given Vote to a CanonicalVote, which does
// not contain ValidatorIndex and ValidatorAddress fields, or any fields
// relating to vote extensions.
func CanonicalizeVote(chainID string, vote *tmproto.Vote) tmproto.CanonicalVote {
	_ = "STUB: not implemented"
	return *new(tmproto.CanonicalVote)
}

// encoded as sfixed64
// encoded as sfixed64

// CanonicalTime can be used to stringify time in a canonical way.
func CanonicalTime(t time.Time) string {
	_ = "STUB: not implemented"
	// Note that sending time over amino resets it to
	// local time, we need to force UTC here, so the
	// signatures match
	return ""
}
