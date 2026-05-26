package keys

import (
	"github.com/spf13/cobra"
)

const (
	flagYes   = "yes"
	flagForce = "force"
)

// DeleteKeyCommand deletes a key from the key store.
func DeleteKeyCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// confirm deletion, unless -y is passed
