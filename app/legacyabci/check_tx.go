package legacyabci

import (
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/utils/tracing"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	feegrantkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/keeper"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

var defaultRecoveryMiddleware = newDefaultRecoveryMiddleware()

type CheckTxKeepers struct {
	AccountKeeper  authkeeper.AccountKeeper
	BankKeeper     bankkeeper.Keeper
	FeeGrantKeeper *feegrantkeeper.Keeper
	IBCKeeper      *ibckeeper.Keeper
	OracleKeeper   oraclekeeper.Keeper
	EvmKeeper      *evmkeeper.Keeper
	ParamsKeeper   paramskeeper.Keeper
	UpgradeKeeper  *upgradekeeper.Keeper
}

func CheckTx(
	ctx sdk.Context,
	tx sdk.Tx,
	txConfig client.TxConfig,
	keepers *CheckTxKeepers,
	checksum [32]byte,
	contextCacher func(sdk.Context) (sdk.Context, sdk.CacheMultiStore),
	latestCtxGetter func() sdk.Context,
	tracingInfo *tracing.Info,
) (
	gInfo sdk.GasInfo,
	result *sdk.Result,
	txCtx sdk.Context,
	err error,
) {
	_ = "STUB: not implemented"
	return *new(sdk.GasInfo), nil, *new(sdk.Context), nil
}

// TODO(PLT-343): remove once tx_duration verified

// trace AnteHandler

// GasMeter expected to be set in AnteHandler
