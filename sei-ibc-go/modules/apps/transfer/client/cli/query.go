package cli

import (
	"github.com/spf13/cobra"
)

// GetCmdQueryDenomTrace defines the command to query a a denomination trace from a given trace hash or ibc denom.
func GetCmdQueryDenomTrace() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryDenomTraces defines the command to query all the denomination trace infos
// that this chain mantains.
func GetCmdQueryDenomTraces() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdParams returns the command handler for ibc-transfer parameter querying.
func GetCmdParams() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdParams returns the command handler for ibc-transfer parameter querying.
func GetCmdQueryEscrowAddress() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryDenomHash defines the command to query a denomination hash from a given trace.
func GetCmdQueryDenomHash() *cobra.Command { _ = "STUB: not implemented"; return nil }
