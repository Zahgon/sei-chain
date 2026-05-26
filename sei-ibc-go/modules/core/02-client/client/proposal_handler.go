package client

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	govclient "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client"
	govrest "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/client/rest"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/client/cli"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal, emptyRestHandler)
	UpgradeProposalHandler      = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, emptyRestHandler)
)

func emptyRestHandler(client.Context) govrest.ProposalRESTHandler {
	_ = "STUB: not implemented"
	return *new(govrest.ProposalRESTHandler)
}
