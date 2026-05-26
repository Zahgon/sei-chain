package ante

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"

	"github.com/sei-protocol/sei-chain/x/evm/derived"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/state"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/sei-chain/x/evm/types/ethtx"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("app", "ante")

func EvmCheckTxAnte(
	ctx sdk.Context,
	tx sdk.Tx,
	upgradeKeeper *upgradekeeper.Keeper,
	ek *evmkeeper.Keeper,
) (returnCtx sdk.Context, returnErr error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// cached and validated

func EvmStatelessChecks(ctx sdk.Context, tx sdk.Tx, chainID *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// this means the message has `Derived` set from the outside, in which case we should reject

// Check if gas exceed the limit

// If there exists a maximum block gas limit, we must ensure that the tx
// does not exceed it.
//nolint:gosec

// validate chain ID on the transaction

// legacy either can have a zero or correct chain ID

// after legacy, all transactions must have the correct chain ID

func DecorateContext(ctx sdk.Context, ek *evmkeeper.Keeper, tx sdk.Tx, txData ethtx.TxData, etx *ethtypes.Transaction, sender common.Address, seiSender sdk.AccAddress) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

// set EVM properties

//nolint:gosec

func HandleAssociateTx(ctx sdk.Context, ek *evmkeeper.Keeper, atx *ethtx.AssociateTx, readOnly bool) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// Hash custom message passed in

func CheckAndDecodeSignature(ctx sdk.Context, txData ethtx.TxData, chainID *big.Int, isBlockTest bool) (common.Address, sdk.AccAddress, cryptotypes.PubKey, derived.SignerVersion, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(sdk.AccAddress), *new(cryptotypes.PubKey), *new(derived.SignerVersion), nil
}

// need to allow unprotected legacy txs in blocktest
// to not lose coverage for other parts of the code

func AssociateAddress(ctx sdk.Context, ek *evmkeeper.Keeper, evmAddr common.Address, seiAddr sdk.AccAddress, seiPubkey cryptotypes.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

func EvmCheckAndChargeFees(ctx sdk.Context, sender common.Address, ek *evmkeeper.Keeper, upgradeKeeper *upgradekeeper.Keeper, txData ethtx.TxData, etx *ethtypes.Transaction, msg *evmtypes.MsgEVMTransaction, version derived.SignerVersion, statelessChecks bool) (*state.DBImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For now we are simply assuming excessive blob gas is 0. In the future we might change it to be
// dynamic based on prior block usage.
// nolint:gosec

func CheckNonce(ctx sdk.Context, ek *evmkeeper.Keeper, etx *ethtypes.Transaction, evmAddr common.Address) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

func IsAccountBalancePositive(ctx sdk.Context, evmKeeper *evmkeeper.Keeper, seiAddr sdk.AccAddress, evmAddr common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// minimum fee per gas required for a tx to be processed
func GetBaseFee(ctx sdk.Context, evmKeeper *evmkeeper.Keeper, upgradeKeeper *upgradekeeper.Keeper) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// lowest allowed fee per gas, base fee will not be lower than this
func GetMinimumFee(ctx sdk.Context, evmKeeper *evmkeeper.Keeper) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func CalculatePriority(ctx sdk.Context, txData ethtx.TxData, evmKeeper *evmkeeper.Keeper) *big.Int {
	_ = "STUB: not implemented"
	return nil
}
