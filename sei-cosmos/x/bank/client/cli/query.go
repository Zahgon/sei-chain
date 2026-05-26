package cli

import (
	"github.com/spf13/cobra"
)

const (
	FlagDenom = "denom"
)

// GetQueryCmd returns the parent command for all x/bank CLi query commands. The
// provided clientCtx should have, at a minimum, a verifier, Tendermint RPC client,
// and marshaler set.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetBalancesCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdDenomsMetadata defines the cobra command to query client denomination metadata.
func GetCmdDenomsMetadata() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetCmdQueryTotalSupply() *cobra.Command { _ = "STUB: not implemented"; return nil }
