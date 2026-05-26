package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the feegrant MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(k Keeper) feegrant.MsgServer {
	_ = "STUB: not implemented"
	return *new(feegrant.MsgServer)
}

var _ feegrant.MsgServer = msgServer{}

// GrantAllowance grants an allowance from the granter's funds to be used by the grantee.
func (k msgServer) GrantAllowance(goCtx context.Context, msg *feegrant.MsgGrantAllowance) (*feegrant.MsgGrantAllowanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Checking for duplicate entry

// RevokeAllowance revokes a fee allowance between a granter and grantee.
func (k msgServer) RevokeAllowance(goCtx context.Context, msg *feegrant.MsgRevokeAllowance) (*feegrant.MsgRevokeAllowanceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
