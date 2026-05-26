package cli

import (
	"github.com/spf13/cobra"
)

const flagHex = "hex"

// GetDecodeCommand returns the decode command to take serialized bytes and turn
// it into a JSON-encoded transaction.
func GetDecodeCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }
