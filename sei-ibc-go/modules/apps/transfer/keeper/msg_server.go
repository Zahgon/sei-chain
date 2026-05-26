package keeper

import (
	"context"

	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/types"
)

var logger = seilog.NewLogger("ibc-go", "modules", "apps", "transfer", "keeper")

var _ types.MsgServer = Keeper{}

// Transfer defines a rpc handler method for MsgTransfer.
func (k Keeper) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
