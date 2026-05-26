package cli

import (
	"github.com/spf13/cobra"
)

// Proposal flags
const (
	FlagTitle        = "title"
	FlagDescription  = "description"
	FlagIsExpedited  = "is-expedited"
	FlagProposalType = "type"
	FlagDeposit      = "deposit"
	flagVoter        = "voter"
	flagDepositor    = "depositor"
	flagStatus       = "status"
	FlagProposal     = "proposal"
)

type proposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	IsExpedited bool   `json:"is_expedited,omitempty"`
	Deposit     string `json:"deposit,omitempty"`
}

// ProposalFlags defines the core required fields of a proposal. It is used to
// verify that these values are not provided in conjunction with a JSON proposal
// file.
var ProposalFlags = []string{
	FlagTitle,
	FlagDescription,
	FlagProposalType,
	FlagDeposit,
}

// NewTxCmd returns the transaction commands for this module
// governance ModuleClient is slightly different from other ModuleClients in that
// it contains a slice of "proposal" child commands. These commands are respective
// to proposal type handlers that are implemented in other modules but are mounted
// under the governance CLI (eg. parameter change proposals).
func NewTxCmd(propCmds []*cobra.Command) *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCmdSubmitProposal implements submitting a proposal transaction command.
func NewCmdSubmitProposal() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCmdDeposit implements depositing tokens for an active proposal.
func NewCmdDeposit() *cobra.Command { _ = "STUB: not implemented"; return nil }

// validate that the proposal id is a uint

// Get depositor address

// Get amount of coins

// NewCmdVote implements creating a new vote command.
func NewCmdVote() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Get voting address

// validate that the proposal id is a uint

// Find out which vote option user chose

// Build vote message and run basic validation

// NewCmdWeightedVote implements creating a new weighted vote command.
func NewCmdWeightedVote() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Get voter address

// validate that the proposal id is a uint

// Figure out which vote options user chose

// Build vote message and run basic validation
