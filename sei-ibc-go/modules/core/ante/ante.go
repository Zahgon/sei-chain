package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
)

type AnteDecorator struct {
	k *keeper.Keeper
}

func NewAnteDecorator(k *keeper.Keeper) AnteDecorator {
	_ = "STUB: not implemented"
	return *new(AnteDecorator)
}

// AnteDecorator returns an error if a multiMsg tx only contains packet messages (Recv, Ack, Timeout) and additional update messages
// and all packet messages are redundant. If the transaction is just a single UpdateClient message, or the multimsg transaction
// contains some other message type, then the antedecorator returns no error and continues processing to ensure these transactions
// are included. This will ensure that relayers do not waste fees on multiMsg transactions when another relayer has already submitted
// all packets, by rejecting the tx at the mempool layer.
func (ad AnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	// do not run redundancy check on DeliverTx or simulate
	return *new(sdk.Context), nil
}

// keep track of total packet messages and number of redundancies across `RecvPacket`, `AcknowledgePacket`, and `TimeoutPacket/OnClose`

// if the multiMsg tx has a msg that is not a packet msg or update msg, then we will not return error
// regardless of if all packet messages are redundant. This ensures that non-packet messages get processed
// even if they get batched with redundant packet messages.

// only return error if all packet messages are redundant
