package wasm

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

type OracleWasmQueryHandler struct {
	oracleKeeper oraclekeeper.Keeper
}

func NewOracleWasmQueryHandler(keeper *oraclekeeper.Keeper) *OracleWasmQueryHandler {
	_ = "STUB: not implemented"
	return nil
}

func (handler OracleWasmQueryHandler) GetExchangeRates(ctx sdk.Context) (*types.QueryExchangeRatesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (handler OracleWasmQueryHandler) GetOracleTwaps(ctx sdk.Context, req *types.QueryTwapsRequest) (*types.QueryTwapsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
