package types

import (
	"context"
	"errors"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

// PrivValidatorType defines the implemtation types.
type PrivValidatorType uint8

const (
	MockSignerClient      = PrivValidatorType(0x00) // mock signer
	FileSignerClient      = PrivValidatorType(0x01) // signer client via file
	RetrySignerClient     = PrivValidatorType(0x02) // signer client with retry via socket
	SignerSocketClient    = PrivValidatorType(0x03) // signer client via socket
	ErrorMockSignerClient = PrivValidatorType(0x04) // error mock signer
	SignerGRPCClient      = PrivValidatorType(0x05) // signer client via gRPC
)

// PrivValidator defines the functionality of a local Tendermint validator
// that signs votes and proposals, and never double signs.
type PrivValidator interface {
	GetPubKey(context.Context) (crypto.PubKey, error)

	SignVote(ctx context.Context, chainID string, vote *tmproto.Vote) error
	SignProposal(ctx context.Context, chainID string, proposal *tmproto.Proposal) error
}

type PrivValidatorsByAddress []PrivValidator

func (pvs PrivValidatorsByAddress) Len() int { _ = "STUB: not implemented"; return 0 }

func (pvs PrivValidatorsByAddress) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pvs PrivValidatorsByAddress) Swap(i, j int) { _ = "STUB: not implemented"; return }

//----------------------------------------
// MockPV

// MockPV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type MockPV struct {
	PrivKey              crypto.PrivKey
	breakProposalSigning bool
	breakVoteSigning     bool
}

func NewMockPV() MockPV { _ = "STUB: not implemented"; return *new(MockPV) }

// NewMockPVWithParams allows one to create a MockPV instance, but with finer
// grained control over the operation of the mock validator. This is useful for
// mocking test failures.
func NewMockPVWithParams(privKey crypto.PrivKey, breakProposalSigning, breakVoteSigning bool) MockPV {
	_ = "STUB: not implemented"
	return *new(MockPV)
}

// Implements PrivValidator.
func (pv MockPV) GetPubKey(ctx context.Context) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// Implements PrivValidator.
func (pv MockPV) SignVote(ctx context.Context, chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil
}

// Implements PrivValidator.
func (pv MockPV) SignProposal(ctx context.Context, chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil
}

func (pv MockPV) ExtractIntoValidator(ctx context.Context, votingPower int64) *Validator {
	_ = "STUB: not implemented"
	return nil
}

// String returns a string representation of the MockPV.
func (pv MockPV) String() string { _ = "STUB: not implemented"; return "" }

// mockPV will never return an error, ignored here

// XXX: Implement.
func (pv MockPV) DisableChecks() {
	_ = "STUB: not implemented"
	// Currently this does nothing,
	// as MockPV has no safety checks at all.
	return
}

type ErroringMockPV struct {
	MockPV
}

var ErroringMockPVErr = errors.New("erroringMockPV always returns an error")

// Implements PrivValidator.
func (pv *ErroringMockPV) GetPubKey(ctx context.Context) (crypto.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PubKey), nil
}

// Implements PrivValidator.
func (pv *ErroringMockPV) SignVote(ctx context.Context, chainID string, vote *tmproto.Vote) error {
	_ = "STUB: not implemented"
	return nil

	// Implements PrivValidator.
}

func (pv *ErroringMockPV) SignProposal(ctx context.Context, chainID string, proposal *tmproto.Proposal) error {
	_ = "STUB: not implemented"
	return nil

	// NewErroringMockPV returns a MockPV that fails on each signing request. Again, for testing only.
}

func NewErroringMockPV() *ErroringMockPV { _ = "STUB: not implemented"; return nil }
