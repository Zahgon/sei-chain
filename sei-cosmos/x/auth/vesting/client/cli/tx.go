package cli

import (
	"github.com/spf13/cobra"
)

// Transaction command flags
const (
	FlagDelayed = "delayed"
	FlagAdmin   = "admin"
)

// GetTxCmd returns vesting module's transaction commands.
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewMsgCreateVestingAccountCmd returns a CLI command handler for creating a
// MsgCreateVestingAccount transaction.
func NewMsgCreateVestingAccountCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
