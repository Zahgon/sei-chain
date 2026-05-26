package multisig

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// tmMultisig implements a K of N threshold multisig. It is used for
// Amino JSON marshaling of LegacyAminoPubKey (see below for details).
//
// This struct is copy-pasted from:
// https://github.com/tendermint/tendermint/blob/v0.33.9/crypto/multisig/threshold_pubkey.go
//
// This struct was used in the SDK <=0.39. In 0.40 and the switch to protobuf,
// it has been converted to LegacyAminoPubKey. However, there's one difference:
// the threshold field was an `uint` before, and an `uint32` after. This caused
// amino marshaling to be breaking: amino marshals `uint32` as a JSON number,
// and `uint` as a JSON string.
//
// In this file, we're overriding LegacyAminoPubKey's default JSON Amino
// marshaling by using this struct. Please note that we are NOT overriding the
// Amino binary marshaling, as that _might_ introduce breaking changes in the
// keyring, where multisigs are amino-binary-encoded.
//
// ref: https://github.com/cosmos/cosmos-sdk/issues/8776
type tmMultisig struct {
	K       uint                 `json:"threshold"`
	PubKeys []cryptotypes.PubKey `json:"pubkeys"`
}

// protoToTm converts a LegacyAminoPubKey into a tmMultisig.
func protoToTm(protoPk *LegacyAminoPubKey) (tmMultisig, error) {
	_ = "STUB: not implemented"
	return *new(tmMultisig), nil
}

// tmToProto converts a tmMultisig into a LegacyAminoPubKey.
func tmToProto(tmPk tmMultisig) (*LegacyAminoPubKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // bounds checked above

// MarshalAminoJSON overrides amino JSON unmarshaling.
func (m LegacyAminoPubKey) MarshalAminoJSON() (tmMultisig, error) {
	_ = "STUB: not implemented" //nolint:revive
	return *new(tmMultisig), nil
}

// UnmarshalAminoJSON overrides amino JSON unmarshaling.
func (m *LegacyAminoPubKey) UnmarshalAminoJSON(tmPk tmMultisig) error {
	_ = "STUB: not implemented"
	return nil
}

// Instead of just doing `*m = *protoPk`, we prefer to modify in-place the
// existing Anys inside `m` (instead of allocating new Anys), as so not to
// break the `.compat` fields in the existing Anys.

// create the compat jsonBz value

// UnmarshalJSON():
// just sets the compat.jsonBz value.
// always succeeds: err == nil
