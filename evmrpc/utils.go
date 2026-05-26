package evmrpc

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

const LatestCtxHeight int64 = -1

// EVM launch block heights for different chains
const Pacific1EVMLaunchHeight int64 = 79123881

// ErrBlockNotFoundByHash is returned when no block exists for the given hash (e.g. empty or unknown hash).
// Ethereum-compatible RPCs should return result: null for this case instead of an error.
var ErrBlockNotFoundByHash = errors.New("block not found by hash")

// GetBlockNumberByNrOrHash returns the height of the block with the given number or hash.
func GetBlockNumberByNrOrHash(ctx context.Context, tmClient client.LocalClient, wm *WatermarkManager, blockNrOrHash rpc.BlockNumberOrHash) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Synthetic genesis from eth_getBlockByNumber("0x0") is not stored under this hash in Tendermint.

func getBlockNumber(ctx context.Context, tmClient client.LocalClient, number rpc.BlockNumber) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// requesting Block with nil means the latest block

func getHeightFromBigIntBlockNumber(latest int64, blockNumber *big.Int) int64 {
	_ = "STUB: not implemented"
	return 0
}

// this avoids a gosec lint error rather than just casting
func toUint64(value int64) uint64 { _ = "STUB: not implemented"; return 0 }

func getTestKeyring(homeDir string) (keyring.Keyring, error) {
	_ = "STUB: not implemented"
	return *new(keyring.Keyring), nil
}

func getAddressPrivKeyMap(kb keyring.Keyring) map[string]*ecdsa.PrivateKey {
	_ = "STUB: not implemented"
	return nil
}

// will only show local key

func blockByNumberWithRetry(ctx context.Context, client client.LocalClient, height *int64, maxRetries int) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// retry once, since application DB and block DB are not committed atomically so it's possible for
// receipt to exist while block results aren't committed yet

func blockByHash(ctx context.Context, client client.LocalClient, hash bytes.HexBytes) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blockByHashWithRetry(ctx context.Context, client client.LocalClient, hash bytes.HexBytes, maxRetries int) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// retry once, since application DB and block DB are not committed atomically so it's possible for
// receipt to exist while block results aren't committed yet

// ValidateEVMBlockHeight checks if the requested block height is valid for EVM queries
func ValidateEVMBlockHeight(chainID string, blockHeight int64) error {
	_ = "STUB: not implemented"
	// Only validate for pacific-1 chain
	return nil
}

type indexedMsg struct {
	msg   sdk.Msg
	index int
}

func filterTransactions(
	k *keeper.Keeper,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	block *coretypes.ResultBlock,
	includeSyntheticTxs bool,
	includeBankTransfers bool,
	cacheCreationMutex *sync.Mutex,
	globalBlockCache BlockCache,
) []indexedMsg {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

// check if the transaction bumped nonce. If not, exclude it

func recordMetrics(ctx context.Context, apiMethod string, connectionType ConnectionType, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func recordMetricsWithError(ctx context.Context, apiMethod string, connectionType ConnectionType, startTime time.Time, err error, panicValue any) {
	_ = "STUB: not implemented"
	return
}

// these are only metrics that are specifically typed errors for tracking.

// TODO(PLT-326): remove legacy dual-emit once dashboards are migrated to evmrpc_* OTEL metrics. Use metrics.requestLatencySeconds histogram instead.

func CheckVersion(ctx sdk.Context, k *keeper.Keeper) error { _ = "STUB: not implemented"; return nil }

func bankExists(ctx sdk.Context, k *keeper.Keeper) bool { _ = "STUB: not implemented"; return false }

func evmExists(ctx sdk.Context, k *keeper.Keeper) bool { _ = "STUB: not implemented"; return false }

func shouldIncludeSynthetic(namespace string) bool { _ = "STUB: not implemented"; return false }

type typedTxHash struct {
	hash  common.Hash
	isEvm bool
}

func getTxHashesFromBlock(
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	k *keeper.Keeper,
	block *coretypes.ResultBlock,
	shouldIncludeSynthetic bool,
	cacheCreationMutex *sync.Mutex,
	globalBlockCache BlockCache,
) []typedTxHash {
	_ = "STUB: not implemented"
	return nil
}

func isReceiptFromAnteError(ctx sdk.Context, receipt *types.Receipt) bool {
	_ = "STUB: not implemented"
	// hacky heuristic
	return false
}

// isReceiptUntraceable returns true if the receipt represents a tx whose
// trace would be empty or meaningless. Shared discriminator used by every
// *ExcludeTraceFail site (tx, block, trace) so they filter the same set.
//
//   - TxType == ShellEVMTxType: chain-generated synthetic, no real EVM
//     execution. app/receipt.go writes these for wasm txs to surface CW20
//     events on the EVM side; they have no trace.
//   - EffectiveGasPrice == 0 && GasUsed == 0: ante-deferred stub receipt
//     from x/evm/keeper/abci.go — the tx bumped its nonce in ante but
//     never reached the VM. WriteReceipt for any executed tx sets both
//     fields > 0 (intrinsic gas at minimum, msg.GasPrice for the fee on
//     a chain with positive min fee), so reverts and OOG pass through.
//
// This is intentionally narrower than isReceiptFromAnteError's
// post-v5.8.0 branch: that helper is tuned to keep insufficient-funds
// receipts visible to the regular eth_getBlockBy* endpoints (per
// PR #2343). *ExcludeTraceFail wants the opposite per evmrpc/README.md.
func isReceiptUntraceable(receipt *types.Receipt) bool { _ = "STUB: not implemented"; return false }

type ParallelRunner struct {
	Done  sync.WaitGroup
	Queue chan func()
}

var panicHook atomic.Value

func SetPanicHook(h func(interface{})) { _ = "STUB: not implemented"; return }

func NewParallelRunner(cnt int, capacity int) *ParallelRunner {
	_ = "STUB: not implemented"
	return nil
}

func runWithRecovery(f func()) { _ = "STUB: not implemented"; return }

func recoverAndLog() { _ = "STUB: not implemented"; return }

func must[V any](v V, err error) V { _ = "STUB: not implemented"; return *new(V) }
