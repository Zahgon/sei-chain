package debug

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// Cmd creates a main CLI command
func Cmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// getPubKeyFromString decodes SDK PubKey using JSON marshaler.
func getPubKeyFromString(ctx client.Context, pkstr string) (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

func PubkeyCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func AddrCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// try hex, then bech32

func RawBytesCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
