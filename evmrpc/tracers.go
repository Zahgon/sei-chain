package evmrpc

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	_ "github.com/ethereum/go-ethereum/eth/tracers/js"     // run init()s to register JS tracers
	_ "github.com/ethereum/go-ethereum/eth/tracers/native" // run init()s to register native tracers
	"github.com/ethereum/go-ethereum/export"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/sei-protocol/sei-chain/app/legacyabci"
	evmrpcconfig "github.com/sei-protocol/sei-chain/evmrpc/config"
	"github.com/sei-protocol/sei-chain/sei-cosmos/baseapp"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

const (
	IsPanicCacheSize = 5000
	IsPanicCacheTTL  = 1 * time.Minute

	callTracerName     = "callTracer"
	prestateTracerName = "prestateTracer"
	flatCallTracerName = "flatCallTracer"
)

var errTraceConcurrencyLimit = errors.New("trace request rejected due to concurrency limit: server busy")

type DebugAPI struct {
	tracersAPI         *tracers.API
	tmClient           client.LocalClient
	keeper             *keeper.Keeper
	ctxProvider        func(int64) sdk.Context
	txConfigProvider   func(int64) client.TxConfig
	connectionType     ConnectionType
	isPanicCache       *expirable.LRU[common.Hash, bool] // hash to isPanic
	backend            *Backend
	traceCallSemaphore chan struct{} // Semaphore for limiting concurrent trace calls
	maxBlockLookback   int64
	traceTimeout       time.Duration
	profiledBlockTrace bool
}

// acquireTraceSemaphore attempts to acquire a slot from the traceCallSemaphore.
// It returns a function that must be called (typically with defer) to release the semaphore.
// If the semaphore is nil (unlimited concurrency), it does nothing and returns a no-op release function.
// The acquisition respects cancellation and fails fast if all trace slots are in use.
func (api *DebugAPI) acquireTraceSemaphore(ctx context.Context) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If cancellation won the race at the same time as semaphore acquisition,
// release the slot and surface the context error.

// No-op if semaphore is not active

// prepareTraceContext creates the trace timeout context and acquires a trace slot if one
// is immediately available, returning a cleanup function for acquired resources.
func (api *DebugAPI) prepareTraceContext(ctx context.Context) (context.Context, func(), error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

type SeiDebugAPI struct {
	*DebugAPI
}

func NewDebugAPI(
	tmClient client.LocalClient,
	k *keeper.Keeper,
	beginBlockKeepers legacyabci.BeginBlockKeepers,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	config *SimulateConfig,
	app *baseapp.BaseApp,
	antehandler sdk.AnteHandler,
	connectionType ConnectionType,
	debugCfg evmrpcconfig.Config,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
	watermarks *WatermarkManager,
) *DebugAPI {
	_ = "STUB: not implemented"
	return nil
}

func NewSeiDebugAPI(
	tmClient client.LocalClient,
	k *keeper.Keeper,
	beginBlockKeepers legacyabci.BeginBlockKeepers,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	config *SimulateConfig,
	app *baseapp.BaseApp,
	antehandler sdk.AnteHandler,
	connectionType ConnectionType,
	debugCfg evmrpcconfig.Config,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
	watermarks *WatermarkManager,
) *SeiDebugAPI {
	_ = "STUB: not implemented"
	return nil
}

// Note: The embedded DebugAPI here does not get its own isPanicCache initialized
// This is consistent with the original code. If it needs one, it should be added.

// isPanicCache: nil, // Explicitly nil as per original structure for SeiDebugAPI's embedded DebugAPI

func (api *DebugAPI) TraceTransaction(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) tryTraceCache(hash common.Hash, config *tracers.TraceConfig) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec

// blockTraceCacheGet assembles a per-tx hit; returns (nil, false) if any miss.
func blockTraceCacheGet(cache *keeper.TraceDB, height int64, txHashes []common.Hash, config *tracers.TraceConfig) ([]*tracers.TxTraceResult, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// tryBlockResultCache reads the per-block JSON in one seek. Preferred over
// blockTraceCacheGet which assembles N per-tx rows.
func tryBlockResultCache(cache *keeper.TraceDB, height int64, config *tracers.TraceConfig) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (api *DebugAPI) tryBlockTraceCacheByNumber(ctx context.Context, number rpc.BlockNumber, config *tracers.TraceConfig) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec

func (api *DebugAPI) tryBlockTraceCacheByHash(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec

// tryExcludeFailBlockTraceCacheByNumber reads the per-block JSON row, parses it,
// and drops entries with Error set. Per-tx rows are skipped — they omit Error.
func (api *DebugAPI) tryExcludeFailBlockTraceCacheByNumber(ctx context.Context, number rpc.BlockNumber, config *tracers.TraceConfig) ([]*tracers.TxTraceResult, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec

func (api *DebugAPI) tryExcludeFailBlockTraceCacheByHash(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) ([]*tracers.TxTraceResult, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec

// filterExcludeFailFromBlockCache is the cached counterpart to the fresh
// trace path in TraceBlockBy{Number,Hash}ExcludeTraceFail. Any change to
// what *ExcludeTraceFail filters out must keep these two in sync —
// delegate the per-trace decision to dropUntraceableTraces here just
// as the fresh path does, so the discriminator stays single-sourced.
// (TestTraceBlockByNumberExcludeTraceFail_AnteStub exercises the fresh
// path with TraceBakeEnabled=false; the cache path's filtering is
// covered transitively via the shared helper.)
func filterExcludeFailFromBlockCache(cache *keeper.TraceDB, height int64, tracer string, k *keeper.Keeper, sdkctx sdk.Context) ([]*tracers.TxTraceResult, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// dropUntraceableTraces filters out trace results that shouldn't surface from
// the *ExcludeTraceFail trace endpoints. With a non-nil keeper, every
// surviving trace must have a receipt that isReceiptUntraceable rejects —
// matching the receipt-side semantic at isPanicOrSyntheticTx (missing
// receipt → exclude, untraceable receipt → exclude). Drops are:
//
//   - trace.Error != "": tracer-level failure (timeout, internal error, etc.).
//   - GetReceipt errors: no receipt for the tx hash. Unreachable in practice
//     (the trace path's pre-filter in Backend.BlockByNumber drops no-receipt
//     txs before they get traced, and the baker only caches traces for
//     txs with receipts), but excluded here for parity with the tx-side
//     check so isReceiptUntraceable's "shared discriminator" promise holds
//     across every site.
//   - isReceiptUntraceable(receipt): synthetic (TxType==ShellEVMTxType) or
//     ante-deferred stub (EffectiveGasPrice==0 && GasUsed==0). For
//     callTracer / flatCallTracer (and the default-tracer insufficient-funds
//     path) errorTrace embeds the underlying error in the JSON and leaves
//     TxTraceResult.Error empty, so the trace.Error check alone can't catch
//     this — the receipt-shape check does.
//
// A nil keeper disables the receipt branch (used by the cache-filter unit
// test in isolation); the trace.Error check still applies.
func dropUntraceableTraces(traces []*tracers.TxTraceResult, k *keeper.Keeper, sdkctx sdk.Context) []*tracers.TxTraceResult {
	_ = "STUB: not implemented"
	return nil
}

func txHashesOf(txs gethtypes.Transactions) []common.Hash { _ = "STUB: not implemented"; return nil }

// bakeableTracerName returns the tracer name iff config matches what the
// baker produces (no per-call TracerConfig); empty means "fall through".
func bakeableTracerName(config *tracers.TraceConfig) string { _ = "STUB: not implemented"; return "" }

func (api *DebugAPI) AsRawJSON(result interface{}) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (api *SeiDebugAPI) TraceBlockByNumberExcludeTraceFail(ctx context.Context, number rpc.BlockNumber, config *tracers.TraceConfig) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *SeiDebugAPI) TraceBlockByHashExcludeTraceFail(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isPanicOrSyntheticTx returns true if the tx isn't traceable — used by the
// *ExcludeTraceFail endpoints to filter out txs whose trace would be empty
// or meaningless.
//
// Per evmrpc/README.md ("Tracing Failure Management Endpoints"), the target
// is txs "included in blocks but not executed" — pre-state-check failures
// (nonce mismatch, insufficient funds, etc.) and chain-generated synthetic
// txs. Reverts, OOG, and other in-VM failures all ran and produced traces;
// they stay in.
//
// Discriminator: receipts are written in two paths. WriteReceipt
// (x/evm/keeper/receipt.go) covers executed txs and sets EffectiveGasPrice
// from msg.GasPrice (>0 on chains with positive min gas price) and GasUsed
// > 0 (intrinsic gas at minimum). The ante-deferred stub path
// (x/evm/keeper/abci.go) writes EffectiveGasPrice=0 and GasUsed=0 for any
// nonce-bumping ante failure — regardless of which check failed (insufficient
// funds, fee, mempool admission, etc.). Both fields zero is the signal that
// the tx never reached the VM.
//
// (This does NOT use filterTransactions's isReceiptFromAnteError. That
// helper's post-v5.8.0 branch is intentionally narrow — PR #2343's
// TestAnteFailureOthers explicitly requires insufficient-funds receipts to
// be *included* in regular eth_getBlockBy* responses. *ExcludeTraceFail has
// the opposite semantic per the README, so it needs its own check.)
//
//   - GetReceipt error                          → no receipt yet           → exclude
//   - TxType == ShellEVMTxType (math.MaxUint32) → chain-generated synthetic → exclude
//   - EffectiveGasPrice==0 && GasUsed==0        → ante-deferred stub        → exclude
//   - anything else (success / revert / OOG)    → executed, has a trace    → include
func (api *DebugAPI) isPanicOrSyntheticTx(ctx context.Context, hash common.Hash) (isPanic bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// No receipt: treat as panic/synthetic. Not cached — the receipt
// store can lag the RPC for a freshly committed tx, so this answer
// may flip to "include" once the write lands.

func (api *DebugAPI) TraceBlockByNumber(ctx context.Context, number rpc.BlockNumber, config *tracers.TraceConfig) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) TraceBlockByHash(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) TraceCall(ctx context.Context, args export.TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, config *tracers.TraceCallConfig) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) GetRawHeader(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (_ hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

func (api *DebugAPI) GetRawBlock(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (_ hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

func (api *DebugAPI) GetRawReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (_ []hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *DebugAPI) GetRawTransaction(ctx context.Context, hash common.Hash) (_ hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

type StateAccessResponse struct {
	AppState        json.RawMessage `json:"app"`
	TendermintState json.RawMessage `json:"tendermint"`
	Receipt         json.RawMessage `json:"receipt"`
}

func (api *DebugAPI) TraceStateAccess(ctx context.Context, hash common.Hash) (result interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only mined txes are supported

// It shouldn't happen in practice.

//nolint:gosec
