package evmrpc

import (
	"context"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/export"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

var ErrPanicTx = errors.New("transaction is panic tx")

const UnconfirmedTxQueryMaxPage = 20
const UnconfirmedTxQueryPerPage = 30

type TransactionAPI struct {
	tmClient           client.LocalClient
	keeper             *keeper.Keeper
	ctxProvider        func(int64) sdk.Context
	txConfigProvider   func(int64) client.TxConfig
	homeDir            string
	connectionType     ConnectionType
	includeSynthetic   bool
	watermarks         *WatermarkManager
	globalBlockCache   BlockCache
	cacheCreationMutex *sync.Mutex
}

type SeiTransactionAPI struct {
	*TransactionAPI
	isPanicTx func(ctx context.Context, hash common.Hash) (bool, error)
}

func NewTransactionAPI(
	tmClient client.LocalClient,
	k *keeper.Keeper,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	homeDir string,
	connectionType ConnectionType,
	watermarks *WatermarkManager,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
) *TransactionAPI {
	_ = "STUB: not implemented"
	return nil
}

func NewSeiTransactionAPI(
	tmClient client.LocalClient,
	k *keeper.Keeper,
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	homeDir string,
	connectionType ConnectionType,
	isPanicTx func(ctx context.Context, hash common.Hash) (bool, error),
	watermarks *WatermarkManager,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
) *SeiTransactionAPI {
	_ = "STUB: not implemented"
	return nil
}

func (t *SeiTransactionAPI) GetTransactionReceiptExcludeTraceFail(ctx context.Context, hash common.Hash) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TransactionAPI) GetTransactionReceipt(ctx context.Context, hash common.Hash) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTransactionReceipt(
	ctx context.Context,
	t *TransactionAPI,
	hash common.Hash,
	excludePanicTxs bool,
	isPanicTx func(ctx context.Context, hash common.Hash) (bool, error),
	includeSynthetic bool,
) (result map[string]interface{}, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When the transaction doesn't exist, the RPC method should return JSON null
// as per specification.

// Fetch block once — used both for ante-failure receipt population and encoding.
//nolint:gosec

// Ethereum JSON-RPC: receipt for a block above safe latest => null, not an error.

// Fill in the receipt if the transaction has failed and used 0 gas
// This case is for when a tx fails before it makes it to the VM

// Find the transaction in the block

// codecov:ignore - defensive error handling for invalid signatures
// codecov:ignore

// Update receipt with correct information

// For contract creation transactions, calculate the contract address

func (t *TransactionAPI) GetVMError(ctx context.Context, hash common.Hash) (result string, returnErr error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *TransactionAPI) GetTransactionByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, txIndex hexutil.Uint) (result *export.RPCTransaction, _err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//not returning error for invalid tx index for complying with Ethereum JSON-RPC spec

func (t *TransactionAPI) getTransactionByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, txIndex uint32) (result *export.RPCTransaction, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TransactionAPI) GetTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, txIndex hexutil.Uint) (result *export.RPCTransaction, _err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//not returning error for invalid tx index for complying with Ethereum JSON-RPC spec

func (t *TransactionAPI) GetTransactionByHash(ctx context.Context, hash common.Hash) (result *export.RPCTransaction, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// first try get from mempool

// codecov:ignore - defensive error handling for invalid signatures
// codecov:ignore
// codecov:ignore

// then try get from committed

//nolint:gosec

//nolint:gosec

func (t *TransactionAPI) GetTransactionErrorByHash(ctx context.Context, hash common.Hash) (result string, returnErr error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *TransactionAPI) GetTransactionCount(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (result *hexutil.Uint64, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HTTP transport pooling already happens globally underneath net/http, so
// creating a fresh RPC client per proxied request is fine here. If we
// start proxying over WebSocket, we'll need explicit custom pooling since
// the underlying TCP connection lifecycle is strictly bound to Dial -> Close calls.

func (t *TransactionAPI) getTransactionWithBlock(block *coretypes.ResultBlock, txIndex uint32, includeSynthetic bool) (*export.RPCTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec
// Ethereum JSON-RPC: eth_getTransactionByBlock*AndIndex returns null when the index has no transaction.

func (t *TransactionAPI) encodeRPCTransaction(ethtx *ethtypes.Transaction, block *coretypes.ResultBlock, txIndex uint32) (*export.RPCTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint:gosec

//nolint:gosec

// replaceFrom updates the From field of the transaction if it is not already set, for edge cases for Legacy Txs.
func replaceFrom(tx *export.RPCTransaction, receipt *types.Receipt) {
	_ = "STUB: not implemented"
	return
}

func (t *TransactionAPI) Sign(ctx context.Context, addr common.Address, data hexutil.Bytes) (result hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

func (t *TransactionAPI) getFilteredMsgs(block *coretypes.ResultBlock) []indexedMsg {
	_ = "STUB: not implemented"
	return nil
}

func getEthTxForTxBz(tx tmtypes.Tx, decoder sdk.TxDecoder) *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// not EVM tx since EVM tx will have exactly one msg

// receipt.TransactionIndex represents the index of the transaction among ALL transactions in the block.
// This function returns the index if irrelevant transactions are excluded.
// Specifically, if includeSynthetic is false, all Cosmos transactions are excluded. If includeSynthetic is true,
// Cosmos transactions without a receipt (i.e. Cosmos transactions that don't touch CW20/721/1155) are excluded.
// It also returns the log index offset, which always includes all logs of relevant transactions, regardless of
// whether logs themselves are synthetic or not.
func GetEvmTxIndex(ctx sdk.Context, block *coretypes.ResultBlock, msgs []indexedMsg, txIndex uint32, k *keeper.Keeper, cacheCreationMutex *sync.Mutex, globalBlockCache BlockCache) (index int, found bool, etx *ethtypes.Transaction, logIndexOffset int) {
	_ = "STUB: not implemented"
	return 0, false, nil, 0
}

func encodeReceipt(
	ctxProvider func(int64) sdk.Context,
	txConfigProvider func(int64) client.TxConfig,
	receipt *types.Receipt,
	k *keeper.Keeper,
	block *coretypes.ResultBlock,
	includeSynthetic bool,
	globalBlockCache BlockCache,
	cacheCreationMutex *sync.Mutex,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert tx index including cosmos txs to tx index excluding cosmos txs

//nolint:gosec
//nolint:gosec

//nolint:gosec

// nolint:gosec

type txUint32OverflowError struct {
	txIndex hexutil.Uint
}

func (e txUint32OverflowError) Error() string { _ = "STUB: not implemented"; return "" }

func txIndexToUint32(txIndex hexutil.Uint) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gosec

func cloneReceiptForMutation(receipt *types.Receipt) *types.Receipt {
	_ = "STUB: not implemented"
	return nil
}
