package types

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/exported"
)

// Message types for the evidence module
const (
	TypeMsgSubmitEvidence = "submit_evidence"
)

var (
	_ sdk.Msg                       = &MsgSubmitEvidence{}
	_ types.UnpackInterfacesMessage = MsgSubmitEvidence{}
	_ exported.MsgSubmitEvidenceI   = &MsgSubmitEvidence{}
)

// NewMsgSubmitEvidence returns a new MsgSubmitEvidence with a signer/submitter.
func NewMsgSubmitEvidence(s sdk.AccAddress, evi exported.Evidence) (*MsgSubmitEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Route returns the MsgSubmitEvidence's route.
func (m MsgSubmitEvidence) Route() string {
	_ = "STUB: not implemented"

	// Type returns the MsgSubmitEvidence's type.
	return ""
}

func (m MsgSubmitEvidence) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic performs basic (non-state-dependant) validation on a MsgSubmitEvidence.
func (m MsgSubmitEvidence) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSignBytes returns the raw bytes a signer is expected to sign when submitting
// a MsgSubmitEvidence message.
func (m MsgSubmitEvidence) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners returns the single expected signer for a MsgSubmitEvidence.
func (m MsgSubmitEvidence) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

func (m MsgSubmitEvidence) GetEvidence() exported.Evidence {
	_ = "STUB: not implemented"
	return *new(exported.Evidence)
}

func (m MsgSubmitEvidence) GetSubmitter() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (m MsgSubmitEvidence) UnpackInterfaces(ctx types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
