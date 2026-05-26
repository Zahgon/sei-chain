package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/03-connection/types"
)

// EmitConnectionOpenInitEvent emits a connection open init event
func EmitConnectionOpenInitEvent(ctx sdk.Context, connectionID string, clientID string, counterparty types.Counterparty) {
	_ = "STUB: not implemented"
	return
}

// EmitConnectionOpenTryEvent emits a connection open try event
func EmitConnectionOpenTryEvent(ctx sdk.Context, connectionID string, clientID string, counterparty types.Counterparty) {
	_ = "STUB: not implemented"
	return
}

// EmitConnectionOpenAckEvent emits a connection open acknowledge event
func EmitConnectionOpenAckEvent(ctx sdk.Context, connectionID string, connectionEnd types.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}

// EmitConnectionOpenConfirmEvent emits a connection open confirm event
func EmitConnectionOpenConfirmEvent(ctx sdk.Context, connectionID string, connectionEnd types.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}
