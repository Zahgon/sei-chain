package cli

import (
	"github.com/spf13/cobra"

	govclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client"
	mintrest "github.com/sei-protocol/sei-chain/x/mint/client/rest"
)

var UpdateMinterHandler = govclient.NewProposalHandler(MsgUpdateMinterProposalCmd, mintrest.UpdateResourceDependencyProposalRESTHandler)

func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func MsgUpdateMinterProposalCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
