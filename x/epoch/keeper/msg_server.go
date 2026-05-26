package keeper

import (
	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

var _ types.MsgServer = msgServer{}
