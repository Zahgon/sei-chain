package cli

import (
	"github.com/spf13/cobra"

	gov "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

const (
	FlagUpgradeHeight = "upgrade-height"
	FlagUpgradeInfo   = "upgrade-info"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdSubmitUpgradeProposal() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCmdSubmitCancelUpgradeProposal implements a command handler for submitting a software upgrade cancel proposal transaction.
func NewCmdSubmitCancelUpgradeProposal() *cobra.Command { _ = "STUB: not implemented"; return nil }

func parseArgsToContent(cmd *cobra.Command, name string) (gov.Content, error) {
	_ = "STUB: not implemented"
	return *new(gov.Content), nil
}
