package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k msgServer) MultiSend(goCtx context.Context, msg *types.MsgMultiSend) (*types.MsgMultiSendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: totalIn == totalOut should already have been checked
