package cli

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

// MigrateContractCmd will migrate a contract to a new code version
func MigrateContractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func parseMigrateContractArgs(args []string, cliCtx client.Context) (types.MsgMigrateContract, error) {
	_ = "STUB: not implemented"
	// get the id of the code to instantiate
	return *new(types.MsgMigrateContract), nil
}

// UpdateContractAdminCmd sets an new admin for a contract
func UpdateContractAdminCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func parseUpdateContractAdminArgs(args []string, cliCtx client.Context) (types.MsgUpdateAdmin, error) {
	_ = "STUB: not implemented"
	return *new(types.MsgUpdateAdmin), nil
}

// ClearContractAdminCmd clears an admin for a contract
func ClearContractAdminCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
