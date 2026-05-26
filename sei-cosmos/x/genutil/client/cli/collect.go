package cli

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil/types"
)

const flagGenTxDir = "gentx-dir"

// CollectGenTxsCmd - return the cobra command to collect genesis transactions
func CollectGenTxsCmd(genBalIterator types.GenesisBalancesIterator, defaultNodeHome string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
