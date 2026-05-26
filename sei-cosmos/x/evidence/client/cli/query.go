package cli

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/query"
)

// GetQueryCmd returns the CLI command with all evidence module query commands
// mounted.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// QueryEvidenceCmd returns the command handler for evidence querying. Evidence
// can be queried for by hash or paginated evidence can be returned.
func QueryEvidenceCmd() func(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }

func queryEvidence(clientCtx client.Context, hash string) error {
	_ = "STUB: not implemented"
	return nil
}

func queryAllEvidence(clientCtx client.Context, pageReq *query.PageRequest) error {
	_ = "STUB: not implemented"
	return nil
}
