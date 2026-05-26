package cli

import (
	"github.com/spf13/cobra"
)

const (
	flagEvents = "events"
	flagType   = "type"

	typeHash   = "hash"
	typeAccSeq = "acc_seq"
	typeSig    = "signature"

	eventFormat = "{eventType}.{eventAttribute}={value}"
)

// GetQueryCmd returns the transaction commands for this module
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// QueryParamsCmd returns the command handler for evidence parameter querying.
func QueryParamsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// QueryNextAccountNumberCmd returns the command handler for evidence parameter querying.
func QueryNextAccountNumberCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetAccountCmd returns a query account that will display the state of the
// account at a given address.
func GetAccountCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetAccountsCmd returns a query command that will display a list of accounts
func GetAccountsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// QueryTxsByEventsCmd returns a command to search through transactions by events.
func QueryTxsByEventsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// QueryTxCmd implements the default command for a tx query.
func QueryTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// If hash is given, then query the tx by hash.

// This case means there's a bug somewhere else in the code. Should not happen.

// This case means there's a bug somewhere else in the code. Should not happen.

// parseSigArgs parses comma-separated signatures from the CLI arguments.
func ParseSigArgs(args []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
