package legacyabci

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/utils/tracing"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	feegrantkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/keeper"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

type DeliverTxKeepers struct {
	AccountKeeper  authkeeper.AccountKeeper
	BankKeeper     bankkeeper.Keeper
	FeeGrantKeeper *feegrantkeeper.Keeper
	OracleKeeper   oraclekeeper.Keeper
	EvmKeeper      *evmkeeper.Keeper
	ParamsKeeper   paramskeeper.Keeper
	UpgradeKeeper  *upgradekeeper.Keeper
}

func DeliverTx(
	ctx sdk.Context,
	tx sdk.Tx,
	txConfig client.TxConfig,
	keepers *DeliverTxKeepers,
	checksum [32]byte,
	contextCacher func(sdk.Context) (sdk.Context, sdk.CacheMultiStore),
	msgRunner func(ctx sdk.Context, msgs []sdk.Msg) (*sdk.Result, error), //TODO: remove
	tracingInfo *tracing.Info,
	evmHook func(ctx sdk.Context, tx sdk.Tx, checksum [32]byte, response sdk.DeliverTxHookInput),
) (
	gInfo sdk.GasInfo,
	result *sdk.Result,
	anteEvents []abci.Event,
	txCtx sdk.Context,
	err error,
) {
	_ = "STUB: not implemented"
	return *new(sdk.GasInfo), nil, nil, *new(sdk.Context), nil
}

// TODO(PLT-343): remove once tx_duration verified

// check for existing parent tracer, and if applicable, use it

// TODO: do we have to wrap with occ enabled check?

// trace AnteHandler

// TODO: simplify

// we do this since we will only be looking at result in DeliverTx

// append the events in the order of occurrence

// only apply hooks if no error
