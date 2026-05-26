package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

var _ authz.MsgServer = Keeper{}

// GrantAuthorization implements the MsgServer.Grant method to create a new grant.
func (k Keeper) Grant(goCtx context.Context, msg *authz.MsgGrant) (*authz.MsgGrantResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RevokeAuthorization implements the MsgServer.Revoke method.
func (k Keeper) Revoke(goCtx context.Context, msg *authz.MsgRevoke) (*authz.MsgRevokeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exec implements the MsgServer.Exec method.
func (k Keeper) Exec(goCtx context.Context, msg *authz.MsgExec) (*authz.MsgExecResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
