package keeper

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	ethstate "github.com/ethereum/go-ethereum/core/state"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/tests"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	receipt "github.com/sei-protocol/sei-chain/sei-db/ledger_db/receipt"
	sctypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	ibctransferkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	wasmkeeper "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/keeper"

	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	"github.com/sei-protocol/sei-chain/x/evm/blocktest"
	"github.com/sei-protocol/sei-chain/x/evm/querier"
	"github.com/sei-protocol/sei-chain/x/evm/replay"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

const Pacific1ChainID = "pacific-1"
const DefaultBlockGasLimit = 10000000

type Keeper struct {
	storeKey          sdk.StoreKey
	transientStoreKey sdk.StoreKey

	Paramstore paramtypes.Subspace

	txResults []*abci.ExecTxResult
	msgs      []*types.MsgEVMTransaction

	bankKeeper     bankkeeper.Keeper
	accountKeeper  *authkeeper.AccountKeeper
	stakingKeeper  *stakingkeeper.Keeper
	transferKeeper ibctransferkeeper.Keeper
	wasmKeeper     *wasmkeeper.PermissionedKeeper
	wasmViewKeeper *wasmkeeper.Keeper
	upgradeKeeper  *upgradekeeper.Keeper

	cachedFeeCollectorAddressMtx *sync.RWMutex
	cachedFeeCollectorAddress    *common.Address

	QueryConfig *querier.Config

	// only used during ETH replay. Not used in chain critical path.
	EthClient       *ethclient.Client
	EthReplayConfig replay.Config

	// only used during blocktest. Not used in chain critical path.
	EthBlockTestConfig blocktest.Config
	BlockTest          *tests.BlockTest

	// used for both ETH replay and block tests. Not used in chain critical path.
	Trie        ethstate.Trie
	DB          ethstate.Database
	CachingDB   *ethstate.CachingDB
	Root        common.Hash
	ReplayBlock *ethtypes.Block

	receiptStore receipt.ReceiptStore

	customPrecompiles       map[common.Address]putils.VersionedPrecompiles
	latestCustomPrecompiles map[common.Address]vm.PrecompiledContract
	latestUpgrade           string

	// traceDB, when non-nil, serves cached debug_trace results and
	// forwards EndBlock heights to the registered baker. nil-safe.
	traceDB *TraceDB

	// traceSnapshotStore + traceSnapshotCapture, when set, capture an O(1)
	// memiavl snapshot of the SC tree at EndBlock so the baker replays
	// against in-memory state instead of SS-pebble. nil-safe.
	traceSnapshotStore   *TraceSnapshotStore
	traceSnapshotCapture func() sctypes.Committer
}

// only used during ETH replay
type ReplayChainContext struct {
	ethClient *ethclient.Client
	chainID   *big.Int
	params    types.Params
}

func (ctx *ReplayChainContext) Engine() consensus.Engine {
	_ = "STUB: not implemented"
	return *new(consensus.Engine)
}

func (ctx *ReplayChainContext) GetHeader(hash common.Hash, number uint64) *ethtypes.Header {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func (ctx *ReplayChainContext) Config() *params.ChainConfig { _ = "STUB: not implemented"; return nil }

func NewKeeper(
	storeKey sdk.StoreKey, transientStoreKey sdk.StoreKey, paramstore paramtypes.Subspace, receiptStore receipt.ReceiptStore,
	bankKeeper bankkeeper.Keeper, accountKeeper *authkeeper.AccountKeeper, stakingKeeper *stakingkeeper.Keeper,
	transferKeeper ibctransferkeeper.Keeper, wasmKeeper *wasmkeeper.PermissionedKeeper, wasmViewKeeper *wasmkeeper.Keeper, upgradeKeeper *upgradekeeper.Keeper) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) SetTraceDB(c *TraceDB) { _ = "STUB: not implemented"; return }
func (k *Keeper) TraceDB() *TraceDB     { _ = "STUB: not implemented"; return nil }

func (k *Keeper) SetTraceSnapshotStore(s *TraceSnapshotStore) { _ = "STUB: not implemented"; return }
func (k *Keeper) TraceSnapshotStore() *TraceSnapshotStore     { _ = "STUB: not implemented"; return nil }
func (k *Keeper) SetTraceSnapshotCapture(f func() sctypes.Committer) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) SetCustomPrecompiles(cp map[common.Address]putils.VersionedPrecompiles, latestUpgrade string) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) CustomPrecompiles(ctx sdk.Context) map[common.Address]vm.PrecompiledContract {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) GetCustomPrecompilesVersions(ctx sdk.Context) map[common.Address]string {
	_ = "STUB: not implemented"
	return nil
}

// requested height hasn't seen this upgrade version yet.

func (k *Keeper) AccountKeeper() *authkeeper.AccountKeeper { _ = "STUB: not implemented"; return nil }

func (k *Keeper) BankKeeper() bankkeeper.Keeper {
	_ = "STUB: not implemented"
	return *new(bankkeeper.Keeper)
}

func (k *Keeper) ReceiptStore() receipt.ReceiptStore {
	_ = "STUB: not implemented"
	return *new(receipt.ReceiptStore)
}

func (k *Keeper) WasmKeeper() *wasmkeeper.PermissionedKeeper { _ = "STUB: not implemented"; return nil }

func (k *Keeper) UpgradeKeeper() *upgradekeeper.Keeper { _ = "STUB: not implemented"; return nil }

func (k *Keeper) GetStoreKey() sdk.StoreKey { _ = "STUB: not implemented"; return *new(sdk.StoreKey) }

func (k *Keeper) IterateAll(ctx sdk.Context, pref []byte, cb func(key, val []byte) bool) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) PrefixStore(ctx sdk.Context, pref []byte) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
}

func (k *Keeper) PurgePrefix(ctx sdk.Context, pref []byte) { _ = "STUB: not implemented"; return }

func (k *Keeper) GetVMBlockContext(ctx sdk.Context, gp core.GasPool) (*vm.BlockContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use hash of block timestamp as info for PREVRANDAO

//nolint:gosec

//nolint:gosec
// only needed for PoW

// Cancun not enabled

// returns a function that provides block header hash based on block number
func (k *Keeper) GetHashFn(ctx sdk.Context) vm.GetHashFunc {
	_ = "STUB: not implemented"
	return *new(vm.GetHashFunc)
}

// current header hash is in the context already

// future block doesn't have a hash yet

// fetch historical hash from historical info

func (k *Keeper) getHistoricalHash(ctx sdk.Context, h int64) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

// too old, already pruned

func (k *Keeper) SetTxResults(txResults []*abci.ExecTxResult) { _ = "STUB: not implemented"; return }

func (k *Keeper) SetMsgs(msgs []*types.MsgEVMTransaction) {
	_ = "STUB: not implemented"

	// Only used in ETH replay
	return
}

func (k *Keeper) PrepareReplayedAddr(ctx sdk.Context, addr common.Address) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) GetBaseFee(ctx sdk.Context) *big.Int { _ = "STUB: not implemented"; return nil }

func (k *Keeper) GetReplayedHeight(ctx sdk.Context) int64 { _ = "STUB: not implemented"; return 0 }

func (k *Keeper) SetReplayedHeight(ctx sdk.Context) { _ = "STUB: not implemented"; return }

func (k *Keeper) GetReplayInitialHeight(ctx sdk.Context) int64 { _ = "STUB: not implemented"; return 0 }

func (k *Keeper) SetReplayInitialHeight(ctx sdk.Context, h int64) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) setInt64State(ctx sdk.Context, key []byte, val int64) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

func (k *Keeper) getInt64State(ctx sdk.Context, key []byte) int64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:gosec

func (k *Keeper) getBlockTestBlockCtx(ctx sdk.Context) (*vm.BlockContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Keeper) getReplayBlockCtx(ctx sdk.Context) (*vm.BlockContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func uint64Cmp(a, b uint64) int { _ = "STUB: not implemented"; return 0 }
