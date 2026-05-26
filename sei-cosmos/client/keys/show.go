package keys

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
)

const (
	// FlagAddress is the flag for the user's address on the command line.
	FlagAddress = "address"
	// FlagPublicKey represents the user's public key on the command line.
	FlagPublicKey = "pubkey"
	// FlagBechPrefix defines a desired Bech32 prefix encoding for a key.
	FlagBechPrefix = "bech"
	// FlagDevice indicates that the information should be shown in the device
	FlagDevice = "device"

	flagMultiSigThreshold = "multisig-threshold"

	defaultMultiSigKeyName = "multi"
)

// ShowKeysCmd shows key information for a given key name.
func ShowKeysCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runShowCmd(cmd *cobra.Command, args []string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Override and show in the device

func fetchKey(kb keyring.Keyring, keyref string) (keyring.Info, error) {
	_ = "STUB: not implemented"
	// firstly check if the keyref is a key name of a key registered in a keyring.
	return *new(keyring.Info), nil
}

// if the key is not there or if we have a problem with a keyring itself then we move to a
// fallback: searching for key by address.

func validateMultisigThreshold(k, nKeys int) error { _ = "STUB: not implemented"; return nil }

func getBechKeyOut(bechPrefix string) (bechKeyOutFn, error) {
	_ = "STUB: not implemented"
	return *new(bechKeyOutFn), nil
}
