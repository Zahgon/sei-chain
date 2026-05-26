package cli

import (
	"github.com/spf13/cobra"
)

// GetEncodeCommand returns the encode command to take a JSONified transaction and turn it into
// Amino-serialized bytes
func GetEncodeCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// re-encode it

// base64 encode the encoded tx bytes
