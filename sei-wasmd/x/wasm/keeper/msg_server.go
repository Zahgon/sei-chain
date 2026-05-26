package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper types.ContractOpsKeeper
}

func NewMsgServerImpl(k types.ContractOpsKeeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

func (m msgServer) StoreCode(goCtx context.Context, msg *types.MsgStoreCode) (*types.MsgStoreCodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m msgServer) InstantiateContract(goCtx context.Context, msg *types.MsgInstantiateContract) (*types.MsgInstantiateContractResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m msgServer) ExecuteContract(goCtx context.Context, msg *types.MsgExecuteContract) (*types.MsgExecuteContractResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m msgServer) MigrateContract(goCtx context.Context, msg *types.MsgMigrateContract) (*types.MsgMigrateContractResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m msgServer) UpdateAdmin(goCtx context.Context, msg *types.MsgUpdateAdmin) (*types.MsgUpdateAdminResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m msgServer) ClearAdmin(goCtx context.Context, msg *types.MsgClearAdmin) (*types.MsgClearAdminResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
