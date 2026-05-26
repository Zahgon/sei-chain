package wasm

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/epoch/keeper"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
)

type EpochWasmQueryHandler struct {
	epochKeeper keeper.Keeper
}

func NewEpochWasmQueryHandler(keeper *keeper.Keeper) *EpochWasmQueryHandler {
	_ = "STUB: not implemented"
	return nil
}

func (handler EpochWasmQueryHandler) GetEpoch(ctx sdk.Context, req *types.QueryEpochRequest) (*types.QueryEpochResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
