package keys

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
)

const (
	flagInteractive = "interactive"
	flagRecover     = "recover"
	flagNoBackup    = "no-backup"
	flagCoinType    = "coin-type"
	flagAccount     = "account"
	flagIndex       = "index"
	flagMultisig    = "multisig"
	flagNoSort      = "nosort"
	flagHDPath      = "hd-path"

	// DefaultKeyPass contains the default key password for genesis transactions
	DefaultKeyPass = "12345678"
)

// AddKeyCommand defines a keys command to add a generated or recovered private key to keybase.
func AddKeyCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runAddCmdPrepare(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

/*
input
  - bip39 mnemonic
  - bip39 passphrase
  - bip44 path
  - local encryption password

output
  - armor encrypted private key (saved to file)
*/
func runAddCmd(ctx client.Context, cmd *cobra.Command, args []string, inBuf *bufio.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// use in memory keybase

// account exists, ask for user confirmation

// If we're using ledger, only thing we need is the path and the bech32 prefix.

// Get bip39 mnemonic

// read entropy seed straight from tmcrypto.Rand and convert to mnemonic

// override bip39 passphrase

// if they use one, make them re-enter it

// Recover key from seed passphrase

// Hide mnemonic from output

func printCreate(cmd *cobra.Command, info keyring.Info, showMnemonic bool, mnemonic string, outputFormat string) error {
	_ = "STUB: not implemented"
	return nil
}

// print mnemonic unless requested not to.
