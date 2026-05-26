package keeper

import (
	"math/big"
	"sync"

	"github.com/ethereum/evmc/v12/bindings/go/evmc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	ethstate "github.com/ethereum/go-ethereum/core/state"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	bankkeeper "github.com/sei-protocol/sei-chain/giga/deps/xbank/keeper"
	"github.com/sei-protocol/sei-chain/giga/deps/xevm/types"
	putils "github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	stakingkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/keeper"
	upgradekeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/upgrade/keeper"
	"github.com/sei-protocol/sei-chain/sei-db/ledger_db/receipt"
	ibctransferkeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/apps/transfer/keeper"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	wasmkeeper "github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/keeper"
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
	nonceMx                      *sync.RWMutex
	pendingTxs                   map[string][]*PendingTx
	hashToNonce                  map[tmtypes.TxHash]*AddressNoncePair

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

	// EvmoneVM holds the loaded evmone VM instance for the Giga executor
	EvmoneVM *evmc.VM

	// UseRegularStore when true causes PrefixStore to use ctx.KVStore instead of ctx.GigaKVStore.
	// This is for debugging/testing to isolate Giga executor logic from GigaKVStore layer.
	UseRegularStore bool
}

type AddressNoncePair struct {
	Address common.Address
	Nonce   uint64
}

type PendingTx struct {
	Hash     tmtypes.TxHash
	Nonce    uint64
	Priority int64
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
	storeKey sdk.StoreKey, transientStoreKey sdk.StoreKey, paramstore paramtypes.Subspace, receiptStateStore receipt.ReceiptStore,
	bankKeeper bankkeeper.Keeper, accountKeeper *authkeeper.AccountKeeper, stakingKeeper *stakingkeeper.Keeper,
	transferKeeper ibctransferkeeper.Keeper, wasmKeeper *wasmkeeper.PermissionedKeeper, wasmViewKeeper *wasmkeeper.Keeper, upgradeKeeper *upgradekeeper.Keeper) *Keeper {
	_ = "STUB: not implemented"
	return nil
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

// GetKVStore returns the appropriate KVStore based on the UseRegularStore flag.
// When UseRegularStore is true (for debugging/testing), returns regular KVStore.
// Otherwise returns GigaKVStore.
func (k *Keeper) GetKVStore(ctx sdk.Context) sdk.KVStore {
	_ = "STUB: not implemented"
	return *new(sdk.KVStore)
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

// CalculateNextNonce calculates the next nonce for an address
// If includePending is true, it will consider pending nonces
// If includePending is false, it will only return the next nonce from GetNonce
func (k *Keeper) CalculateNextNonce(ctx sdk.Context, addr common.Address, includePending bool) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// we only want the latest nonce if we're not including pending

// get the pending nonces (nil is fine)

// Check each nonce starting from latest until we find a gap
// That gap is the next nonce we should use.

// if it's not in pending, then it's the next nonce

// AddPendingNonce adds a pending nonce to the keeper
func (k *Keeper) AddPendingNonce(hash tmtypes.TxHash, addr common.Address, nonce uint64, priority int64) {
	_ = "STUB: not implemented"
	return
}

// we want to no-op whether it's a genuine duplicate or not

// replace existing tx

// we don't need to return error here if priority is lower.
// Tendermint will take care of rejecting the tx from mempool

// RemovePendingNonce removes a pending nonce from the keeper but leaves a hole
// so that a future transaction must use this nonce.
func (k *Keeper) RemovePendingNonce(hash tmtypes.TxHash) { _ = "STUB: not implemented"; return }

func (k *Keeper) SetTxResults(txResults []*abci.ExecTxResult) { _ = "STUB: not implemented"; return }

func (k *Keeper) SetMsgs(msgs []*types.MsgEVMTransaction) {
	_ = "STUB: not implemented"

	// Test use only
	return
}

func (k *Keeper) GetPendingTxs() map[string][]*PendingTx { _ = "STUB: not implemented"; return nil }

// Test use only
func (k *Keeper) GetHashesToNonces() map[tmtypes.TxHash]*AddressNoncePair {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) GetBaseFee(ctx sdk.Context) *big.Int { _ = "STUB: not implemented"; return nil }

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

func (k *Keeper) GetGasPool() core.GasPool { _ = "STUB: not implemented"; return *new(core.GasPool) }

func uint64Cmp(a, b uint64) int { _ = "STUB: not implemented"; return 0 }
