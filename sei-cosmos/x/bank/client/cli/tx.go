package cli

import (
	"github.com/spf13/cobra"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewSendTxCmd returns a CLI command handler for creating a MsgSend transaction.
func NewSendTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
