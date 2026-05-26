// Package wgtest holds the helpers shared by the wireguard schema and
// channel-wiring tests. It exists so tests across the proto packages and
// the reactor packages don't have to redefine the same Commit / Evidence
// builders.
package wgtest

import (
	"testing"

	gogoproto "github.com/gogo/protobuf/proto"

	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// MaxCommitSignatures mirrors the (wireguard.max_count) annotation on
// tmproto.Commit.signatures and the types.MaxVotesCount cap enforced by
// Commit.ValidateBasic. Tests use it to build at-cap and over-cap payloads;
// the live runtime cap is whatever the generated SchemaForCommit says.
const MaxCommitSignatures = types.MaxVotesCount

// Marshal wraps gogoproto.Marshal with a test-fatal error.
func Marshal(t *testing.T, m gogoproto.Message) []byte { _ = "STUB: not implemented"; return nil }

// CommitWith returns a Commit with n empty CommitSig entries.
func CommitWith(n int) *tmproto.Commit { _ = "STUB: not implemented"; return nil }

// EvidenceWithCommit wraps a Commit in a LightClientAttackEvidence whose
// ConflictingBlock's signed header carries the commit. This is the path
// the wireguard schemas descend through to reach the Commit cap.
func EvidenceWithCommit(c *tmproto.Commit) tmproto.Evidence {
	_ = "STUB: not implemented"
	return *new(tmproto.Evidence)
}

// EvidenceList wraps each Commit in a LightClientAttackEvidence and returns
// the slice ready to drop into tmproto.EvidenceList.Evidence.
func EvidenceList(commits ...*tmproto.Commit) []tmproto.Evidence {
	_ = "STUB: not implemented"
	return nil
}
