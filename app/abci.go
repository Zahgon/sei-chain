package app

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
)

func (app *App) BeginBlock(
	ctx sdk.Context,
	height int64,
	votes []abci.VoteInfo,
	byzantineValidators []abci.Misbehavior,
	checkHeight bool,
) (res abci.ResponseBeginBlock) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseBeginBlock)
}

// TODO(PLT-327): remove once app_abci_begin_block_duration_seconds verified

// inline begin block

// TODO(PLT-327): remove once app_build_info observable gauge verified
// check if we've reached a target height, if so, execute any applicable handlers

func (app *App) MidBlock(ctx sdk.Context, height int64) []abci.Event {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) EndBlock(ctx sdk.Context, height int64, blockGasUsed int64) (res abci.ResponseEndBlock) {
	_ = "STUB: not implemented"
	return *new(abci.ResponseEndBlock)
}

// TODO(PLT-327): remove once app_abci_end_block_duration_seconds verified

// TODO(PLT-327): remove once app_abci_module_end_block_duration_seconds verified

func (app *App) CheckTx(ctx context.Context, req *abci.RequestCheckTxV2) *abci.ResponseCheckTxV2 {
	_ = "STUB: not implemented"
	return nil
}

// TODO(PLT-327): remove once app_abci_check_tx_duration_seconds verified

//nolint:gosec

//nolint:gosec

func (app *App) EvmNonce(evmAddr common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

func (app *App) EvmBalance(evmAddr common.Address, seiAddrBz []byte) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) DeliverTx(ctx sdk.Context, req abci.RequestDeliverTxV2, tx sdk.Tx, checksum [32]byte) abci.ResponseDeliverTx {
	_ = "STUB: not implemented"
	return *new(abci.ResponseDeliverTx)
}

// ensure we carry the initial context from tracer here

// update context with trace span new context

// TODO(PLT-327): remove once app_abci_deliver_tx_duration_seconds verified
// TODO(PLT-327): remove once app_abci_deliver_tx_duration_seconds verified

// TODO(PLT-327): remove once app_tx_count_total verified
// TODO(PLT-327): remove once app_tx_count_total verified
// TODO(PLT-327): remove once app_tx_gas_used verified
// TODO(PLT-327): remove once app_tx_gas_wanted verified

// if we have a result, use those events instead of just the anteEvents

//nolint:gosec
//nolint:gosec

// TODO: populate error data for EVM err

// DeliverTxBatch is not part of the ABCI specification, but this is here for code convention
func (app *App) DeliverTxBatch(ctx sdk.Context, req sdk.DeliverTxBatchRequest) (res sdk.DeliverTxBatchResponse) {
	_ = "STUB: not implemented"
	return *new(sdk.DeliverTxBatchResponse)
}

// TODO(PLT-327): remove once app_abci_deliver_batch_tx_duration_seconds verified

// update context with trace span new context

// avoid overhead for empty batches

func (app *App) Commit(ctx context.Context) (res *abci.ResponseCommit, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// legacy: telemetry.MeasureSince in sei-cosmos/baseapp/abci.go TODO(PLT-327)

// After a successful Commit, publish the pending eth_newHeads event
// stashed by FinalizeBlocker. Subscribers see only committed state.
// Header.AppHash is intentionally left unset (Tendermint convention:
// it holds the previous block's hash); stateRoot is sourced from
// response.AppHash by encodeCommittedBlock.
