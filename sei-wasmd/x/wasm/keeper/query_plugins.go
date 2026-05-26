package keeper

import (
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

type QueryHandler struct {
	Ctx         sdk.Context
	Plugins     WasmVMQueryHandler
	Caller      sdk.AccAddress
	gasRegister GasRegister
}

func NewQueryHandler(ctx sdk.Context, vmQueryHandler WasmVMQueryHandler, caller sdk.AccAddress, gasRegister GasRegister) QueryHandler {
	_ = "STUB: not implemented"
	return *new(QueryHandler)
}

type GRPCQueryRouter interface {
	Route(path string) baseapp.GRPCQueryHandler
}

// -- end baseapp interfaces --

var _ wasmvmtypes.Querier = QueryHandler{}

func (q QueryHandler) Query(request wasmvmtypes.QueryRequest, gasLimit uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	// set a limit for a subCtx
	return nil, nil
}

// discard all changes/ events in subCtx by not committing the cached context

// make sure we charge the higher level context even on panic

// short-circuit, the rest is dealing with handling existing errors

// special mappings to system error (which are not redacted)

// Issue #759 - we don't return error string for worries of non-determinism

func (q QueryHandler) GasConsumed() uint64 { _ = "STUB: not implemented"; return 0 }

type CustomQuerier func(ctx sdk.Context, request json.RawMessage) ([]byte, error)

type QueryPlugins struct {
	Bank     func(ctx sdk.Context, request *wasmvmtypes.BankQuery) ([]byte, error)
	Custom   CustomQuerier
	IBC      func(ctx sdk.Context, caller sdk.AccAddress, request *wasmvmtypes.IBCQuery) ([]byte, error)
	Staking  func(ctx sdk.Context, request *wasmvmtypes.StakingQuery) ([]byte, error)
	Stargate func(ctx sdk.Context, request *wasmvmtypes.StargateQuery) ([]byte, error)
	Wasm     func(ctx sdk.Context, request *wasmvmtypes.WasmQuery) ([]byte, error)
}

type contractMetaDataSource interface {
	GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *types.ContractInfo
}

type wasmQueryKeeper interface {
	contractMetaDataSource
	QueryRaw(ctx sdk.Context, contractAddress sdk.AccAddress, key []byte) []byte
	QuerySmart(ctx sdk.Context, contractAddr sdk.AccAddress, req []byte) ([]byte, error)
	IsPinnedCode(ctx sdk.Context, codeID uint64) bool
}

func DefaultQueryPlugins(
	bank types.BankViewKeeper,
	staking types.StakingKeeper,
	distKeeper types.DistributionKeeper,
	channelKeeper types.ChannelKeeper,
	queryRouter GRPCQueryRouter,
	wasm wasmQueryKeeper,
) QueryPlugins {
	_ = "STUB: not implemented"
	return *new(QueryPlugins)
}

func (e QueryPlugins) Merge(o *QueryPlugins) QueryPlugins {
	_ = "STUB: not implemented"
	// only update if this is non-nil and then only set values
	return *new(QueryPlugins)
}

// HandleQuery executes the requested query
func (e QueryPlugins) HandleQuery(ctx sdk.Context, caller sdk.AccAddress, request wasmvmtypes.QueryRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	// do the query
	return nil, nil
}

func BankQuerier(bankKeeper types.BankViewKeeper) func(ctx sdk.Context, request *wasmvmtypes.BankQuery) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

func NoCustomQuerier(sdk.Context, json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IBCQuerier(wasm contractMetaDataSource, channelKeeper types.ChannelKeeper) func(ctx sdk.Context, caller sdk.AccAddress, request *wasmvmtypes.IBCQuery) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// it must match the port and be in open state

// it must be in open state

func StargateQuerier(queryRouter GRPCQueryRouter) func(ctx sdk.Context, request *wasmvmtypes.StargateQuery) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

func StakingQuerier(keeper types.StakingKeeper, distKeeper types.DistributionKeeper) func(ctx sdk.Context, request *wasmvmtypes.StakingQuery) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// validators := keeper.GetAllValidators(ctx)

func sdkToDelegations(ctx sdk.Context, keeper types.StakingKeeper, delegations []stakingtypes.Delegation) (wasmvmtypes.Delegations, error) {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.Delegations), nil
}

// shares to amount logic comes from here:
// https://github.com/cosmos/cosmos-sdk/blob/v0.38.3/x/staking/keeper/querier.go#L404

func sdkToFullDelegation(ctx sdk.Context, keeper types.StakingKeeper, distKeeper types.DistributionKeeper, delegation stakingtypes.Delegation) (*wasmvmtypes.FullDelegation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: this is very rough but better than nothing...
// https://github.com/CosmWasm/wasmd/issues/282
// if this (val, delegate) pair is receiving a redelegation, it cannot redelegate more
// otherwise, it can redelegate the full amount
// (there are cases of partial funds redelegated, but this is a start)

// FIXME: make a cleaner way to do this (modify the sdk)
// we need the info from `distKeeper.calculateDelegationRewards()`, but it is not public
// neither is `queryDelegationRewards(ctx sdk.Context, _ []string, req abci.RequestQuery, k Keeper)`
// so we go through the front door of the querier....

// FIXME: simplify this enormously when
// https://github.com/cosmos/cosmos-sdk/issues/7466 is merged
func getAccumulatedRewards(ctx sdk.Context, distKeeper types.DistributionKeeper, delegation stakingtypes.Delegation) ([]wasmvmtypes.Coin, error) {
	_ = "STUB: not implemented"
	// Try to get *delegator* reward info!
	return nil, nil
}

// now we have it, convert it into wasmvm types

func WasmQuerier(k wasmQueryKeeper) func(ctx sdk.Context, request *wasmvmtypes.WasmQuery) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// ConvertSdkCoinsToWasmCoins covert sdk type to wasmvm coins type
func ConvertSdkCoinsToWasmCoins(coins []sdk.Coin) wasmvmtypes.Coins {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.Coins)
}

// ConvertSdkCoinToWasmCoin covert sdk type to wasmvm coin type
func ConvertSdkCoinToWasmCoin(coin sdk.Coin) wasmvmtypes.Coin {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.Coin)
}

var _ WasmVMQueryHandler = WasmVMQueryHandlerFn(nil)

// WasmVMQueryHandlerFn is a helper to construct a function based query handler.
type WasmVMQueryHandlerFn func(ctx sdk.Context, caller sdk.AccAddress, request wasmvmtypes.QueryRequest) ([]byte, error)

// HandleQuery delegates call into wrapped WasmVMQueryHandlerFn
func (w WasmVMQueryHandlerFn) HandleQuery(ctx sdk.Context, caller sdk.AccAddress, request wasmvmtypes.QueryRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
