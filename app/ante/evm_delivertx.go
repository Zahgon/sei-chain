package ante

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/derived"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/sei-chain/x/evm/types/ethtx"
)

func EvmDeliverTxAnte(
	ctx sdk.Context,
	txConfig client.TxConfig,
	tx sdk.Tx,
	upgradeKeeper *upgradekeeper.Keeper,
	ek *evmkeeper.Keeper,
) (returnCtx sdk.Context, returnErr error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// cached and validated

func EvmDeliverHandleSignatures(ctx sdk.Context, ek *evmkeeper.Keeper, txData ethtx.TxData, chainID *big.Int, msg *evmtypes.MsgEVMTransaction) (common.Address, sdk.AccAddress, derived.SignerVersion, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(sdk.AccAddress), *new(derived.SignerVersion), nil
}

func EvmDeliverChargeFees(ctx sdk.Context, ek *evmkeeper.Keeper, upgradeKeeper *upgradekeeper.Keeper, txData ethtx.TxData, etx *ethtypes.Transaction, msg *evmtypes.MsgEVMTransaction, version derived.SignerVersion, evmAddr common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func DecorateNonceCallback(ctx sdk.Context, ek *evmkeeper.Keeper, evmAddr common.Address, txNonce uint64) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// bump nonce if it is for some reason not incremented (e.g. ante failure)
