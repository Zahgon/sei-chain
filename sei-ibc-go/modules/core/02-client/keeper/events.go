package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// EmitCreateClientEvent emits a create client event
func EmitCreateClientEvent(ctx sdk.Context, clientID string, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// EmitUpdateClientEvent emits an update client event
func EmitUpdateClientEvent(ctx sdk.Context, clientID string, clientState exported.ClientState, consensusHeight exported.Height, headerStr string) {
	_ = "STUB: not implemented"
	return
}

// EmitUpdateClientEvent emits an upgrade client event
func EmitUpgradeClientEvent(ctx sdk.Context, clientID string, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// EmitUpdateClientProposalEvent emits an update client proposal event
func EmitUpdateClientProposalEvent(ctx sdk.Context, clientID string, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// EmitUpgradeClientProposalEvent emits an upgrade client proposal event
func EmitUpgradeClientProposalEvent(ctx sdk.Context, title string, height int64) {
	_ = "STUB: not implemented"
	return
}

// EmitSubmitMisbehaviourEvent emits a client misbehaviour event
func EmitSubmitMisbehaviourEvent(ctx sdk.Context, clientID string, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// EmitSubmitMisbehaviourEventOnUpdate emits a client misbehaviour event on a client update event
func EmitSubmitMisbehaviourEventOnUpdate(ctx sdk.Context, clientID string, clientState exported.ClientState, consensusHeight exported.Height, headerStr string) {
	_ = "STUB: not implemented"
	return
}
