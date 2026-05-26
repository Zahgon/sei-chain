package app

import (
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
)

var _ sdk.TxPrioritizer = (*SeiTxPrioritizer)(nil).GetTxPriorityHint

type SeiTxPrioritizer struct {
	evmKeeper     *evmkeeper.Keeper
	upgradeKeeper *upgradekeeper.Keeper
	paramsKeeper  *paramskeeper.Keeper
}

func NewSeiTxPrioritizer(ek *evmkeeper.Keeper, uk *upgradekeeper.Keeper, pk *paramskeeper.Keeper) *SeiTxPrioritizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SeiTxPrioritizer) GetTxPriorityHint(ctx sdk.Context, tx sdk.Tx) (_priorityHint int64, _err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Fall back to no-op priority if we panic for any reason. This is to avoid DoS
// vectors where a malicious actor crafts a transaction that panics the
// prioritizer. Since the prioritizer is used as a hint only, it's safe to fall
// back to zero priority in this case and log the panic for monitoring purposes.

// The context already has a priority set, return it.

// This should never happen since IsEVMMessage returned true. But we defensively
// return zero priority to be safe.

func (s *SeiTxPrioritizer) getEvmTxPriority(ctx sdk.Context, evmTx *evmtypes.MsgEVMTransaction) (int64, error) {
	_ = "STUB: not implemented"

	// Unpack the transaction data first to avoid double unpacking as part of preprocessing.
	return 0, nil
}

// Unassociated associate transactions have the second-highest priority.
// This is to ensure that associate transactions are processed before
// regular transactions, but after oracle transactions.
//
// Note that we are not checking if sufficient funds are present here to keep the
// priority calculation fast. CheckTx should fully check the transaction.

// Check txData for sanity.

// Check blob hashes for sanity. If EVM version is Cancun or later, and the
// transaction contains at least one blob, we need to make sure the transaction
// carries a non-zero blob fee cap.

// For now we are simply assuming excessive blob gas is 0. In the future we might change it to be
// dynamic based on prior block usage.

//nolint:gosec

func (s *SeiTxPrioritizer) getEvmBaseFee(ctx sdk.Context) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (s *SeiTxPrioritizer) getCosmosTxPriority(ctx sdk.Context, feeTx sdk.FeeTx) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gosec

func isOracleTx(tx sdk.FeeTx) bool { _ = "STUB: not implemented"; return false }
