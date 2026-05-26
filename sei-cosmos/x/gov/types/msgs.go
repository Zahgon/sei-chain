package types

import (
	"fmt"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Governance message types and routes
const (
	TypeMsgDeposit        = "deposit"
	TypeMsgVote           = "vote"
	TypeMsgVoteWeighted   = "weighted_vote"
	TypeMsgSubmitProposal = "submit_proposal"
)

var (
	_, _, _, _ sdk.Msg                       = &MsgSubmitProposal{}, &MsgDeposit{}, &MsgVote{}, &MsgVoteWeighted{}
	_          types.UnpackInterfacesMessage = &MsgSubmitProposal{}
)

// NewMsgSubmitProposal creates a new MsgSubmitProposal.
func NewMsgSubmitProposal(content Content, initialDeposit sdk.Coins, proposer sdk.AccAddress) (*MsgSubmitProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMsgSubmitProposalWithExpedite creates a new MsgSubmitProposal with expedited or not.
func NewMsgSubmitProposalWithExpedite(content Content, initialDeposit sdk.Coins, proposer sdk.AccAddress, isExpedited bool) (*MsgSubmitProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MsgSubmitProposal) GetInitialDeposit() sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

func (m *MsgSubmitProposal) GetProposer() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (m *MsgSubmitProposal) GetContent() Content { _ = "STUB: not implemented"; return *new(Content) }

func (m *MsgSubmitProposal) SetInitialDeposit(coins sdk.Coins) { _ = "STUB: not implemented"; return }

func (m *MsgSubmitProposal) SetProposer(address fmt.Stringer) { _ = "STUB: not implemented"; return }

func (m *MsgSubmitProposal) SetContent(content Content) error {
	_ = "STUB: not implemented"
	return nil
}

// Route implements Msg
func (m MsgSubmitProposal) Route() string {
	_ = "STUB: not implemented"

	// Type implements Msg
	return ""
}

func (m MsgSubmitProposal) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic implements Msg
func (m MsgSubmitProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetSignBytes implements Msg
func (m MsgSubmitProposal) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners implements Msg
func (m MsgSubmitProposal) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface
func (m MsgSubmitProposal) String() string { _ = "STUB: not implemented"; return "" }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgSubmitProposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgDeposit creates a new MsgDeposit instance
func NewMsgDeposit(depositor sdk.AccAddress, proposalID uint64, amount sdk.Coins) *MsgDeposit {
	_ = "STUB: not implemented"
	return nil
}

// Route implements Msg
func (msg MsgDeposit) Route() string {
	_ = "STUB: not implemented"

	// Type implements Msg
	return ""
}

func (msg MsgDeposit) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic implements Msg
func (msg MsgDeposit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface
func (msg MsgDeposit) String() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements Msg
func (msg MsgDeposit) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners implements Msg
func (msg MsgDeposit) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// NewMsgVote creates a message to cast a vote on an active proposal
func NewMsgVote(voter sdk.AccAddress, proposalID uint64, option VoteOption) *MsgVote {
	_ = "STUB: not implemented"
	return nil
}

// Route implements Msg
func (msg MsgVote) Route() string {
	_ = "STUB: not implemented"

	// Type implements Msg
	return ""
}

func (msg MsgVote) Type() string {
	_ = "STUB: not implemented"

	// ValidateBasic implements Msg
	return ""
}

func (msg MsgVote) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface
func (msg MsgVote) String() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements Msg
func (msg MsgVote) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners implements Msg
func (msg MsgVote) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }

// NewMsgVoteWeighted creates a message to cast a vote on an active proposal
func NewMsgVoteWeighted(voter sdk.AccAddress, proposalID uint64, options WeightedVoteOptions) *MsgVoteWeighted {
	_ = "STUB: not implemented"
	return nil
}

// Route implements Msg
func (msg MsgVoteWeighted) Route() string {
	_ = "STUB: not implemented"

	// Type implements Msg
	return ""
}

func (msg MsgVoteWeighted) Type() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic implements Msg
func (msg MsgVoteWeighted) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface
func (msg MsgVoteWeighted) String() string { _ = "STUB: not implemented"; return "" }

// GetSignBytes implements Msg
func (msg MsgVoteWeighted) GetSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetSigners implements Msg
func (msg MsgVoteWeighted) GetSigners() []sdk.AccAddress { _ = "STUB: not implemented"; return nil }
