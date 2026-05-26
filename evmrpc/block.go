package evmrpc

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

const (
	EthNamespace  = "eth"
	SeiNamespace  = "sei"
	Sei2Namespace = "sei2"
)

// genesisBlockHashHex is the block hash returned by GetBlockByNumber("0x0"). Hash-based lookups
// must recognize this so that count/block-by-hash stay consistent with block-by-number.
const genesisBlockHashHex = "0xF9D3845DF25B43B1C6926F3CEDA6845C17F5624E12212FD8847D0BA01DA1AB9E"

var genesisBlockHash = common.HexToHash(genesisBlockHashHex)

// genesisBlockTxCount is the transaction count for the synthetic genesis block (eth_getBlockTransactionCountByHash/ByNumber for genesis).
var genesisBlockTxCount = func() *hexutil.Uint { u := hexutil.Uint(0); return &u }()

func encodeGenesisBlock() map[string]any { _ = "STUB: not implemented"; return nil }

// inapplicable to Sei
// inapplicable to Sei
// inapplicable to Sei

// inapplicable to Sei
// inapplicable to Sei

// inapplicable to Sei

type BlockAPI struct {
	tmClient             client.LocalClient
	keeper               *keeper.Keeper
	ctxProvider          func(int64) sdk.Context
	txConfigProvider     func(int64) client.TxConfig
	connectionType       ConnectionType
	namespace            string
	includeShellReceipts bool
	includeBankTransfers bool
	watermarks           *WatermarkManager
	globalBlockCache     BlockCache
	cacheCreationMutex   *sync.Mutex
}

type SeiBlockAPI struct {
	*BlockAPI
}

func NewBlockAPI(tmClient client.LocalClient, k *keeper.Keeper, ctxProvider func(int64) sdk.Context, txConfigProvider func(int64) client.TxConfig, connectionType ConnectionType, watermarks *WatermarkManager, globalBlockCache BlockCache, cacheCreationMutex *sync.Mutex) *BlockAPI {
	_ = "STUB: not implemented"
	return nil
}

func NewSeiBlockAPI(
	tmClient client.LocalClient,
	k *keeper.Keeper,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	connectionType ConnectionType,
	watermarks *WatermarkManager,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
) *SeiBlockAPI {
	_ = "STUB: not implemented"
	return nil
}

func NewSei2BlockAPI(
	tmClient client.LocalClient,
	k *keeper.Keeper,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	connectionType ConnectionType,
	watermarks *WatermarkManager,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
) *SeiBlockAPI {
	_ = "STUB: not implemented"
	return nil
}

func (a *SeiBlockAPI) GetBlockByNumberExcludeTraceFail(ctx context.Context, number rpc.BlockNumber, fullTx bool) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	// Match eth_getBlockByNumber("0x0"): synthetic genesis, not the Tendermint block at height 0.
	return nil, nil
}

// Exclude synthetic txs (filterTransactions drops them) and ante-failure
// stub receipts (EncodeTmBlock drops them via excludeUntraceable).

func (a *SeiBlockAPI) GetBlockByHashExcludeTraceFail(ctx context.Context, blockHash common.Hash, fullTx bool) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	// See note on GetBlockByNumberExcludeTraceFail.
	return nil, nil
}

func (a *BlockAPI) GetBlockTransactionCountByNumber(ctx context.Context, number rpc.BlockNumber) (result *hexutil.Uint, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *BlockAPI) GetBlockTransactionCountByHash(ctx context.Context, blockHash common.Hash) (result *hexutil.Uint, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *BlockAPI) GetBlockByHash(ctx context.Context, blockHash common.Hash, fullTx bool) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	// used for both: eth_ and sei_ namespaces
	return nil, nil
}

func (a *BlockAPI) getBlockByHash(ctx context.Context, blockHash common.Hash, fullTx bool, includeSyntheticTxs bool, excludeUntraceable bool) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ethereum spec: empty or non-existent block hash returns result=null, not error.

// Validate EVM block height for pacific-1 chain

func (a *BlockAPI) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for compatibility with the graph, always return genesis block

func (a *BlockAPI) getBlockByNumber(
	ctx context.Context,
	number rpc.BlockNumber,
	fullTx bool,
	includeSyntheticTxs bool,
	excludeUntraceable bool,
) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate EVM block height for pacific-1 chain

// Ethereum JSON-RPC: non-existent / future numeric block => null, not an error.

func (a *BlockAPI) GetBlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (result []map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ethereum spec: empty or non-existent block hash returns result=null, not error.

// Synthetic genesis (eth_getBlockByNumber("0x0")): empty receipts without TM/watermarks.
// Callers may pass the genesis hash or the literal block number 0x0 (parsed as number, not hash).

// Get height from params

// Get all tx hashes for the block

// Get tx receipts for all hashes in parallel

//nolint:gosec

// EncodeTmBlock renders a tendermint block as an eth_getBlockBy* response.
//
// excludeUntraceable, when true, drops EVM txs whose receipt is an
// ante-deferred stub (EffectiveGasPrice==0 && GasUsed==0). x/evm/keeper/abci.go
// writes such stubs for txs that passed the nonce check but failed a later
// ante step (insufficient funds, insufficient fee, etc.); they never reached
// the VM and have no meaningful trace. Used by the *ExcludeTraceFail block
// endpoints to satisfy evmrpc/README.md's "included in blocks but not
// executed" filter; the regular eth_getBlockBy* endpoints pass false so
// these txs still surface in normal block responses (per PR #2343's
// TestAnteFailureOthers — users want to see them).
func EncodeTmBlock(
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	block *coretypes.ResultBlock,
	k *keeper.Keeper,
	fullTx bool,
	includeBankTransfers bool,
	includeSyntheticTxs bool,
	excludeUntraceable bool,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Untraceable receipt — tx never reached the VM (ante-deferred
// stub) or is chain-generated synthetic. filterTransactions's
// isReceiptFromAnteError only catches the nonce-error subset
// post-v5.8.0 (per PR #2343, which keeps insufficient-funds
// receipts visible to the regular eth_getBlockBy* endpoints);
// *ExcludeTraceFail needs the broader discriminator. See
// isReceiptUntraceable for the shared definition used at every
// *ExcludeTraceFail site.

// derive gas used from receipt as TxResult.GasUsed may not be accurate
// for ante-failing EVM txs.
//nolint:gosec

//nolint:gosec

//nolint:gosec

// Source block.gasLimit from the active ConsensusParams in the SDK
// context — same place the EVM runtime reads block.gaslimit from
// (x/evm/keeper/keeper.go's BlockContext.GasLimit), so
// eth_getBlockByNumber.gasLimit and the GASLIMIT opcode return the
// same number.

// inapplicable to Sei
// inapplicable to Sei
// inapplicable to Sei

// inapplicable to Sei
// inapplicable to Sei
//nolint:gosec
//nolint:gosec
//nolint:gosec

//nolint:gosec
// inapplicable to Sei

// inapplicable to Sei

func FullBloom() ethtypes.Bloom { _ = "STUB: not implemented"; return *new(ethtypes.Bloom) }

// getEvmTxCount returns the same transaction count as EncodeTmBlock exposes: filterTransactions
// plus the same per-msg rules as EncodeTmBlock (EVM messages need GetReceipt to succeed).
func (a *BlockAPI) getEvmTxCount(block *coretypes.ResultBlock) *hexutil.Uint {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func countBlockTxsLikeEncodeTmBlock(
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	block *coretypes.ResultBlock,
	k *keeper.Keeper,
	includeShellReceipts bool,
	includeBankTransfers bool,
	cacheCreationMutex *sync.Mutex,
	globalBlockCache BlockCache,
) int {
	_ = "STUB: not implemented"
	return 0
}
