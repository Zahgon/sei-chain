package cli

import (
	"github.com/spf13/cobra"
)

const TrueStr = "true"
const FalseStr = "false"

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	_ = "STUB: not implemented"
	// Group epoch queries under a subcommand
	return nil
}

func CmdQuerySeiAddress() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryEVMAddress() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryERC20() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryPayload() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryERC20Payload() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryERC721Payload() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryERC1155Payload() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryPointer() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryPointerVersion() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryPointee() *cobra.Command { _ = "STUB: not implemented"; return nil }

func CmdQueryTxByHash() *cobra.Command { _ = "STUB: not implemented"; return nil }
