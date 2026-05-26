package cli

import (
	"github.com/spf13/cobra"
)

// NewQueryCmd returns a root CLI command handler for all x/params query commands.
func NewQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewQuerySubspaceParamsCmd returns a CLI command handler for querying subspace
// parameters managed by the x/params module.
func NewQuerySubspaceParamsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewQueryFeeParamsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewQueryCosmosGasParamsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewQueryBlockParamsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
