package cli

import (
	"github.com/spf13/cobra"
)

//nolint:unused
const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// this line is used by starport scaffolding # 1
