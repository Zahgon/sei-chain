package wasm

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	tokenfactorykeeper "github.com/sei-protocol/sei-chain/x/tokenfactory/keeper"
	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"
)

type TokenFactoryWasmQueryHandler struct {
	tokenfactoryKeeper tokenfactorykeeper.Keeper
}

func NewTokenFactoryWasmQueryHandler(keeper *tokenfactorykeeper.Keeper) *TokenFactoryWasmQueryHandler {
	_ = "STUB: not implemented"
	return nil
}

func (handler TokenFactoryWasmQueryHandler) GetDenomAuthorityMetadata(ctx sdk.Context, req *types.QueryDenomAuthorityMetadataRequest) (*types.QueryDenomAuthorityMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (handler TokenFactoryWasmQueryHandler) GetDenomsFromCreator(ctx sdk.Context, req *types.QueryDenomsFromCreatorRequest) (*types.QueryDenomsFromCreatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
