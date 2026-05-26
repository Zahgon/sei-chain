package keys

import (
	"github.com/spf13/cobra"
)

const (
	flagUserEntropy = "unsafe-entropy"

	mnemonicEntropySize = 256
)

// MnemonicKeyCommand computes the bip39 memonic for input entropy.
func MnemonicKeyCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// prompt the user to enter some entropy

// hash input entropy to get entropy seed

// read entropy seed straight from crypto.Rand
