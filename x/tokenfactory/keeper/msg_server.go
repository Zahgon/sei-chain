package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"
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

func (server msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (server msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (server msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pay some extra gas cost to give a better error here.

func (server msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (server msgServer) ChangeAdmin(goCtx context.Context, msg *types.MsgChangeAdmin) (*types.MsgChangeAdminResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate new admin we change to should be different from current admin

func (server msgServer) SetDenomMetadata(goCtx context.Context, msg *types.MsgSetDenomMetadata) (*types.MsgSetDenomMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Defense in depth validation of metadata
