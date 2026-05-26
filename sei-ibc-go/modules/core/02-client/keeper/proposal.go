package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/02-client/types"
)

// ClientUpdateProposal will retrieve the subject and substitute client.
// A callback will occur to the subject client state with the client
// prefixed store being provided for both the subject and the substitute client.
// The localhost client is not allowed to be modified with a proposal. The IBC
// client implementations are responsible for validating the parameters of the
// subtitute (enusring they match the subject's parameters) as well as copying
// the necessary consensus states from the subtitute to the subject client
// store. The substitute must be Active and the subject must not be Active.
func (k Keeper) ClientUpdateProposal(ctx sdk.Context, p *types.ClientUpdateProposal) error {
	_ = "STUB: not implemented"
	return nil
}

// emitting events in the keeper for proposal updates to clients

// HandleUpgradeProposal sets the upgraded client state in the upgrade store. It clears
// an IBC client state and consensus state if a previous plan was set. Then  it
// will schedule an upgrade and finally set the upgraded client state in upgrade
// store.
func (k Keeper) HandleUpgradeProposal(ctx sdk.Context, p *types.UpgradeProposal) error {
	_ = "STUB: not implemented"
	return nil
}

// zero out any custom fields before setting

// sets the new upgraded client in last height committed on this chain is at plan.Height,
// since the chain will panic at plan.Height and new chain will resume at plan.Height

// emitting an event for handling client upgrade proposal
