package wasmbinding

import (
	"encoding/json"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	epochwasm "github.com/sei-protocol/sei-chain/x/epoch/client/wasm"
	evmwasm "github.com/sei-protocol/sei-chain/x/evm/client/wasm"
	oraclewasm "github.com/sei-protocol/sei-chain/x/oracle/client/wasm"
	tokenfactorywasm "github.com/sei-protocol/sei-chain/x/tokenfactory/client/wasm"
)

type QueryPlugin struct {
	oracleHandler       oraclewasm.OracleWasmQueryHandler
	epochHandler        epochwasm.EpochWasmQueryHandler
	tokenfactoryHandler tokenfactorywasm.TokenFactoryWasmQueryHandler
	evmHandler          evmwasm.EVMQueryHandler
	stakingKeeper       stakingkeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(oh *oraclewasm.OracleWasmQueryHandler, eh *epochwasm.EpochWasmQueryHandler, th *tokenfactorywasm.TokenFactoryWasmQueryHandler, evmh *evmwasm.EVMQueryHandler, sk stakingkeeper.Keeper) *QueryPlugin {
	_ = "STUB: not implemented"
	return nil
}

func (qp QueryPlugin) HandleOracleQuery(ctx sdk.Context, queryData json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (qp QueryPlugin) HandleEpochQuery(ctx sdk.Context, queryData json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (qp QueryPlugin) HandleTokenFactoryQuery(ctx sdk.Context, queryData json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (qp QueryPlugin) HandleEVMQuery(ctx sdk.Context, queryData json.RawMessage) (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type StakingExtQueryType string

const (
	UnbondingDelegationsType StakingExtQueryType = "staking_ext_unbonding_delegations"
)

type StakingExtQuery struct {
	UnbondingDelegations *UnbondingDelegationsRequest `json:"unbonding_delegations,omitempty"`
}

func (seq *StakingExtQuery) GetQueryType() StakingExtQueryType {
	_ = "STUB: not implemented"
	return *new(StakingExtQueryType)
}

type UnbondingDelegationsRequest struct {
	Delegator string `json:"delegator,omitempty"`
}

type UnbondingDelegationsResponse struct {
	Entries []UnbondingDelegationEntry `json:"entries"`
}

type UnbondingDelegationEntry struct {
	CreationHeight int64   `json:"creation_height"`
	CompletionTime string  `json:"completion_time"`
	InitialBalance sdk.Int `json:"initial_balance"`
	Balance        sdk.Int `json:"balance"`
}

func (qp QueryPlugin) HandleStakingExtQuery(ctx sdk.Context, queryData json.RawMessage) (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
