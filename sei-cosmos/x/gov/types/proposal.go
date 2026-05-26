package types

import (
	"fmt"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// DefaultStartingProposalID is 1
const DefaultStartingProposalID uint64 = 1

// NewProposal creates a new Proposal instance
func NewProposal(content Content, id uint64, submitTime, depositEndTime time.Time, isExpedited bool) (Proposal, error) {
	_ = "STUB: not implemented"
	return *new(Proposal), nil
}

// String implements stringer interface
func (p Proposal) String() string { _ = "STUB: not implemented"; return "" }

// GetContent returns the proposal Content
func (p Proposal) GetContent() Content { _ = "STUB: not implemented"; return *new(Content) }

func (p Proposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

func (p Proposal) ProposalRoute() string { _ = "STUB: not implemented"; return "" }

func (p Proposal) GetTitle() string { _ = "STUB: not implemented"; return "" }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// Proposals is an array of proposal
type Proposals []Proposal

var _ types.UnpackInterfacesMessage = Proposals{}

// Equal returns true if two slices (order-dependant) of proposals are equal.
func (p Proposals) Equal(other Proposals) bool { _ = "STUB: not implemented"; return false }

// String implements stringer interface
func (p Proposals) String() string { _ = "STUB: not implemented"; return "" }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposals) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

type (
	// ProposalQueue defines a queue for proposal ids
	ProposalQueue []uint64
)

// ProposalStatusFromString turns a string into a ProposalStatus
func ProposalStatusFromString(str string) (ProposalStatus, error) {
	_ = "STUB: not implemented"
	return *new(ProposalStatus), nil
}

// ValidProposalStatus returns true if the proposal status is valid and false
// otherwise.
func ValidProposalStatus(status ProposalStatus) bool { _ = "STUB: not implemented"; return false }

// Marshal needed for protobuf compatibility
func (status ProposalStatus) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal needed for protobuf compatibility
func (status *ProposalStatus) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (status ProposalStatus) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// TODO: Do this conversion more directly

// Proposal types
const (
	ProposalTypeText string = "Text"
)

// Implements Content Interface
var _ Content = &TextProposal{}

// NewTextProposal creates a text proposal Content
func NewTextProposal(title, description string, isExpedited bool) Content {
	_ = "STUB: not implemented"
	return *new(Content)
}

// GetTitle returns the proposal title
func (tp *TextProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the proposal description
	return ""
}

func (tp *TextProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the proposal router key
func (tp *TextProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// ProposalType is "Text"
	return ""
}

func (tp *TextProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic validates the content's title and description of the proposal
func (tp *TextProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements Stringer interface
func (tp TextProposal) String() string { _ = "STUB: not implemented"; return "" }

var validProposalTypes = map[string]struct{}{
	ProposalTypeText: {},
}

// RegisterProposalType registers a proposal type. It will panic if the type is
// already registered.
func RegisterProposalType(ty string) { _ = "STUB: not implemented"; return }

// ContentFromProposalType returns a Content object based on the proposal type.
func ContentFromProposalType(title, desc, ty string, isExpedited bool) Content {
	_ = "STUB: not implemented"
	return *new(Content)
}

// IsValidProposalType returns a boolean determining if the proposal type is
// valid.
//
// NOTE: Modules with their own proposal types must register them.
func IsValidProposalType(ty string) bool { _ = "STUB: not implemented"; return false }

// ProposalHandler implements the Handler interface for governance module-based
// proposals (ie. TextProposal ). Since these are
// merely signaling mechanisms at the moment and do not affect state, it
// performs a no-op.
func ProposalHandler(_ sdk.Context, c Content) error { _ = "STUB: not implemented"; return nil }

// both proposal types do not change state so this performs a no-op
