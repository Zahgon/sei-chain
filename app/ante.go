package app

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/utils/tracing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/ante"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	wasm "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm"
	wasmtypes "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper         *ibckeeper.Keeper
	WasmConfig        *wasmtypes.WasmConfig
	WasmKeeper        *wasm.Keeper
	OracleKeeper      *oraclekeeper.Keeper
	EVMKeeper         *evmkeeper.Keeper
	UpgradeKeeper     *upgradekeeper.Keeper
	TXCounterStoreKey sdk.StoreKey
	LatestCtxGetter   func() sdk.Context

	TracingInfo *tracing.Info
}

func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, sdk.AnteHandler, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AnteHandler), *new(sdk.AnteHandler), nil
}

// outermost AnteDecorator. SetUpContext must be called first

// after setup context to enforce limits early

// PriorityDecorator must be called after DeductFeeDecorator which sets tx priority based on tx fees

// SetPubKeyDecorator must be called before all signature verification decorators

// NOTE: NewEVMNoCosmosFieldsDecorator must come first to prevent writing state to chain without being charged.
// E.g. EVMPreprocessDecorator may short-circuit all the later ante handlers if AssociateTx and ignore NewEVMNoCosmosFieldsDecorator.

// NOTE: NewEVMNoCosmosFieldsDecorator must come first to prevent writing state to chain without being charged.
// E.g. EVMPreprocessDecorator may short-circuit all the later ante handlers if AssociateTx and ignore NewEVMNoCosmosFieldsDecorator.
