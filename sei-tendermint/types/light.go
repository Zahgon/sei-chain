package types

import (
	"time"

	tbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

// LightClientInfo describes the status of the light client.
type LightClientInfo struct {
	PrimaryID         string          `json:"primaryID"`
	WitnessesID       []string        `json:"witnessesID"`
	NumPeers          int             `json:"number_of_peers,string"`
	LastTrustedHeight int64           `json:"last_trusted_height,string"`
	LastTrustedHash   tbytes.HexBytes `json:"last_trusted_hash"`
	LatestBlockTime   time.Time       `json:"latest_block_time"`
	TrustingPeriod    string          `json:"trusting_period"`
	// TrustedBlockExpired is true if LatestBlockTime + TrustingPeriod is before
	// the time /status was called.
	TrustedBlockExpired bool `json:"trusted_block_expired"`
}

// LightBlock pairs a SignedHeader with its ValidatorSet and forms the basis of
// the light client.
type LightBlock struct {
	*SignedHeader `json:"signed_header"`
	ValidatorSet  *ValidatorSet `json:"validator_set"`
}

// ValidateBasic checks that the LightBlock is internally consistent. It does
// not verify any cryptographic signatures.
func (lb LightBlock) ValidateBasic(chainID string) error { _ = "STUB: not implemented"; return nil }

// The validator set must match the hash committed in the header.

// String returns a string representation of the LightBlock.
func (lb LightBlock) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented string representation of the LightBlock,
// showing its SignedHeader and ValidatorSet.
func (lb LightBlock) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts the LightBlock to its protobuf representation.
func (lb *LightBlock) ToProto() (*tmproto.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LightBlockFromProto converts a protobuf LightBlock back into a LightBlock.
// It returns an error if the signed header or validator set is missing or
// invalid.
func LightBlockFromProto(pb *tmproto.LightBlock) (*LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//-----------------------------------------------------------------------------

// SignedHeader is a Header along with the Commit that proves it.
type SignedHeader struct {
	*Header `json:"header"`
	Commit  *Commit `json:"commit"`
}

// ValidateBasic checks that the header and commit are internally consistent
// and belong to the given chain. It does not verify cryptographic signatures;
// use a Verifier to establish that the commit actually proves the header.
func (sh SignedHeader) ValidateBasic(chainID string) error { _ = "STUB: not implemented"; return nil }

// The commit must reference the same block as the header.

// String returns a string representation of the SignedHeader.
func (sh SignedHeader) String() string { _ = "STUB: not implemented"; return "" }

// StringIndented returns an indented string representation of the
// SignedHeader, showing its Header and Commit.
func (sh SignedHeader) StringIndented(indent string) string { _ = "STUB: not implemented"; return "" }

// ToProto converts the SignedHeader to its protobuf representation.
func (sh *SignedHeader) ToProto() *tmproto.SignedHeader { _ = "STUB: not implemented"; return nil }

// SignedHeaderFromProto converts a protobuf SignedHeader back into a
// SignedHeader. It returns an error if the header or commit is invalid.
func SignedHeaderFromProto(shp *tmproto.SignedHeader) (*SignedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
