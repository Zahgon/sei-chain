package keeper

import (
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec
	router   *baseapp.MsgServiceRouter
}

// NewKeeper constructs a message authorization Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.BinaryCodec, router *baseapp.MsgServiceRouter) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// getGrant returns grant stored at skey.
func (k Keeper) getGrant(ctx sdk.Context, skey []byte) (grant authz.Grant, found bool) {
	_ = "STUB: not implemented"
	return *new(authz.Grant), false
}

func (k Keeper) update(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, updated authz.Authorization) error {
	_ = "STUB: not implemented"
	return nil
}

// DispatchActions attempts to execute the provided messages via authorization
// grants from the message signer to the grantee.
func (k Keeper) DispatchActions(ctx sdk.Context, grantee sdk.AccAddress, msgs []sdk.Msg) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If granter != grantee then check authorization.Accept, otherwise we
// implicitly accept.

// emit the events from the dispatched actions

// SaveGrant method grants the provided authorization to the grantee on the granter's account
// with the provided expiration time. If there is an existing authorization grant for the
// same `sdk.Msg` type, this grant overwrites that.
func (k Keeper) SaveGrant(ctx sdk.Context, grantee, granter sdk.AccAddress, authorization authz.Authorization, expiration time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteGrant revokes any authorization for the provided message type granted to the grantee
// by the granter.
func (k Keeper) DeleteGrant(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAuthorizations Returns list of `Authorizations` granted to the grantee by the granter.
func (k Keeper) GetAuthorizations(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress) (authorizations []authz.Authorization) {
	_ = "STUB: not implemented"
	return nil
}

// GetCleanAuthorization returns an `Authorization` and it's expiration time for
// (grantee, granter, message name) grant. If there is no grant `nil` is returned.
// If the grant is expired, the grant is revoked, removed from the storage, and `nil` is returned.
func (k Keeper) GetCleanAuthorization(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) (cap authz.Authorization, expiration time.Time) {
	_ = "STUB: not implemented"
	return *new(authz.Authorization), *new(time.Time)
}

// IterateGrants iterates over all authorization grants
// This function should be used with caution because it can involve significant IO operations.
// It should not be used in query or msg services without charging additional gas.
func (k Keeper) IterateGrants(ctx sdk.Context,
	handler func(granterAddr sdk.AccAddress, granteeAddr sdk.AccAddress, grant authz.Grant) bool,
) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns a GenesisState for a given context.
func (k Keeper) ExportGenesis(ctx sdk.Context) *authz.GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// InitGenesis new authz genesis
func (k Keeper) InitGenesis(ctx sdk.Context, data *authz.GenesisState) {
	_ = "STUB: not implemented"
	return
}
