package keys

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
)

const (
	flagUnarmoredHex = "unarmored-hex"
	flagUnsafe       = "unsafe"
)

// ExportKeyCommand exports private keys from the key store.
func ExportKeyCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func exportUnsafeUnarmored(cmd *cobra.Command, uid string, buf *bufio.Reader, kr keyring.Keyring) error {
	_ = "STUB: not implemented"
	// confirm deletion, unless -y is passed
	return nil
}
