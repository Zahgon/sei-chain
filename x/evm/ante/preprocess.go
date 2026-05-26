package ante

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	accountkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"

	"github.com/sei-protocol/sei-chain/x/evm/derived"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/sei-protocol/sei-chain/x/evm/types/ethtx"
)

// Accounts need to have at least 1Sei to force association. Note that account won't be charged.
const BalanceThreshold uint64 = 1000000

var BigBalanceThreshold *big.Int = new(big.Int).SetUint64(BalanceThreshold)
var BigBalanceThresholdMinus1 *big.Int = new(big.Int).SetUint64(BalanceThreshold - 1)

var SignerMap = map[derived.SignerVersion]func(*big.Int) ethtypes.Signer{
	derived.London: ethtypes.NewLondonSigner,
	derived.Cancun: ethtypes.NewCancunSigner,
	derived.Prague: ethtypes.NewPragueSigner,
}
var AllowedTxTypes = map[derived.SignerVersion][]uint8{
	derived.London: {ethtypes.LegacyTxType, ethtypes.AccessListTxType, ethtypes.DynamicFeeTxType},
	derived.Cancun: {ethtypes.LegacyTxType, ethtypes.AccessListTxType, ethtypes.DynamicFeeTxType, ethtypes.BlobTxType},
	derived.Prague: {ethtypes.LegacyTxType, ethtypes.AccessListTxType, ethtypes.DynamicFeeTxType, ethtypes.BlobTxType, ethtypes.SetCodeTxType},
}

type EVMPreprocessDecorator struct {
	evmKeeper     *evmkeeper.Keeper
	accountKeeper *accountkeeper.AccountKeeper
}

func NewEVMPreprocessDecorator(evmKeeper *evmkeeper.Keeper, accountKeeper *accountkeeper.AccountKeeper) *EVMPreprocessDecorator {
	_ = "STUB: not implemented"
	return nil
}

//nolint:revive
func (p *EVMPreprocessDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// use infinite gas meter for EVM transaction because EVM handles gas checking from within

// check if the account has enough balance (without charging)

// TODO(PLT-330): remove once evm_association_error_total verified

// short-circuit without calling next

// noop; for readability

// not associatedTx and not already associated

func (p *EVMPreprocessDecorator) IsAccountBalancePositive(ctx sdk.Context, seiAddr sdk.AccAddress, evmAddr common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// stateless
func Preprocess(ctx sdk.Context, msgEVMTransaction *evmtypes.MsgEVMTransaction, chainID *big.Int, isBlockTest bool) error {
	_ = "STUB: not implemented"
	return nil
}

// PreprocessUnpacked does the same thing as Preprocess but accepts already unpacked txData to save computation
// if txData is nil, it will unpack from msgEVMTransaction.Data.
func PreprocessUnpacked(ctx sdk.Context, msgEVMTransaction *evmtypes.MsgEVMTransaction, chainID *big.Int, isBlockTest bool, txData ethtx.TxData) error {
	_ = "STUB: not implemented"
	return nil
}

// this means the message has `Derived` set from the outside, in which case we should reject

// already preprocessed

// TxData not passed in, unpack it.

// Hash custom message passed in

// need to allow unprotected legacy txs in blocktest
// to not lose coverage for other parts of the code

func IsTxTypeAllowed(version derived.SignerVersion, txType uint8) bool {
	_ = "STUB: not implemented"
	return false
}

// RecoverSenderFromEthTx recovers the sender's EVM address, Sei address, and public key
// from a signed Ethereum transaction. This encapsulates the version-based signer selection
// and signature recovery logic used by both the ante handler and Giga executor.
//
// This function rejects unprotected transactions (returns error).
// For protected transactions, it uses the appropriate signer based on the chain version.
func RecoverSenderFromEthTx(ctx sdk.Context, ethTx *ethtypes.Transaction, chainID *big.Int) (common.Address, sdk.AccAddress, cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(sdk.AccAddress), *new(cryptotypes.PubKey), nil
}

// Version-based signer selection (same logic as PreprocessUnpacked)

func GetVersion(ctx sdk.Context, ethCfg *params.ChainConfig) derived.SignerVersion {
	_ = "STUB: not implemented"
	return *new(derived.SignerVersion)
}

// nolint:gosec

type EVMAddressDecorator struct {
	evmKeeper     *evmkeeper.Keeper
	accountKeeper *accountkeeper.AccountKeeper
}

func NewEVMAddressDecorator(evmKeeper *evmkeeper.Keeper, accountKeeper *accountkeeper.AccountKeeper) *EVMAddressDecorator {
	_ = "STUB: not implemented"
	return nil
}

//nolint:revive
func (p *EVMAddressDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// check if there is non-zero balance
