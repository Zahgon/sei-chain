package legacytx

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// LegacyMsg defines the old interface a message must fulfill, containing
// Amino signing method and legacy router info.
// Deprecated: Please use `Msg` instead.
type LegacyMsg interface {
	sdk.Msg

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Return the message type.
	// Must be alphanumeric or empty.
	Route() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Type() string
}

// StdSignDoc is replay-prevention structure.
// It includes the result of msg.GetSignBytes(),
// as well as the ChainID (prevent cross chain replay)
// and the Sequence numbers for each signature (prevent
// inchain replay and enforce tx ordering per account).
type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number" yaml:"account_number"`
	Sequence      uint64            `json:"sequence" yaml:"sequence"`
	TimeoutHeight uint64            `json:"timeout_height,omitempty" yaml:"timeout_height"`
	ChainID       string            `json:"chain_id" yaml:"chain_id"`
	Memo          string            `json:"memo" yaml:"memo"`
	Fee           json.RawMessage   `json:"fee" yaml:"fee"`
	Msgs          []json.RawMessage `json:"msgs" yaml:"msgs"`
}

// StdSignBytes returns the bytes to sign for a transaction.
func StdSignBytes(chainID string, accnum, sequence, timeout uint64, fee StdFee, msgs []sdk.Msg, memo string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: StdSignature represents a sig
type StdSignature struct {
	cryptotypes.PubKey `json:"pub_key" yaml:"pub_key"` // optional
	Signature          []byte                          `json:"signature" yaml:"signature"`
}

// Deprecated
func NewStdSignature(pk cryptotypes.PubKey, sig []byte) StdSignature {
	_ = "STUB: not implemented"
	return *new(StdSignature)
}

// GetSignature returns the raw signature bytes.
func (ss StdSignature) GetSignature() []byte { _ = "STUB: not implemented"; return nil }

// GetPubKey returns the public key of a signature as a cryptotypes.PubKey using the
// Amino codec.
func (ss StdSignature) GetPubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"

	// MarshalYAML returns the YAML representation of the signature.
	return *new(cryptotypes.PubKey)
}

func (ss StdSignature) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ss StdSignature) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// StdSignatureToSignatureV2 converts a StdSignature to a SignatureV2
func StdSignatureToSignatureV2(cdc *codec.LegacyAmino, sig StdSignature) (signing.SignatureV2, error) {
	_ = "STUB: not implemented"
	return *new(signing.SignatureV2), nil
}

func pubKeySigToSigData(cdc *codec.LegacyAmino, key cryptotypes.PubKey, sig []byte) (signing.SignatureData, error) {
	_ = "STUB: not implemented"
	return *new(signing.SignatureData), nil
}
