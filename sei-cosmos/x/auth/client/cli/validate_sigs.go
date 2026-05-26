package cli

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client/tx"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func GetValidateSignaturesCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func makeValidateSignaturesCmd() func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// printAndValidateSigs will validate the signatures of a given transaction over its
// expected signers. In addition, if offline has not been supplied, the signature is
// verified over the transaction sign bytes. Returns false if the validation fails.
func printAndValidateSigs(
	cmd *cobra.Command, clientCtx client.Context, chainID string, tx sdk.Tx, offline bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// validate the actual signature over the transaction bytes since we can
// reach out to a full node to query accounts.

func readTxAndInitContexts(clientCtx client.Context, cmd *cobra.Command, filename string) (client.Context, tx.Factory, sdk.Tx, error) {
	_ = "STUB: not implemented"
	return *new(client.Context), *new(tx.Factory), *new(sdk.Tx), nil
}
