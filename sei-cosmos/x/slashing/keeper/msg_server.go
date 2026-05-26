package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the slashing MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

var _ types.MsgServer = msgServer{}

// Unjail implements MsgServer.Unjail method.
// Validators must submit a transaction to unjail itself after
// having been jailed (and thus unbonded) for downtime
func (k msgServer) Unjail(goCtx context.Context, msg *types.MsgUnjail) (*types.MsgUnjailResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
