package main

import (
	"crypto/ecdsa"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

var (
	DefaultPriorityFee = big.NewInt(1000000000)    // 1gwei
	DefaultMaxFee      = big.NewInt(1000000000000) // 1000gwei
)

type EvmTxClient struct {
	accountAddress common.Address
	nonce          atomic.Uint64
	chainId        *big.Int
	gasPrice       *big.Int
	ethClients     []*ethclient.Client
	mtx            sync.RWMutex
	privateKey     *ecdsa.PrivateKey
	evmAddresses   *EVMAddresses
	useEip1559     bool
}

func NewEvmTxClient(
	key cryptotypes.PrivKey,
	chainId *big.Int,
	gasPrice *big.Int,
	ethClients []*ethclient.Client,
	evmAddresses *EVMAddresses,
	useEip1559 bool,
) *EvmTxClient {
	_ = "STUB: not implemented"
	return nil
}

// Set starting nonce

func (txClient *EvmTxClient) GetTxForMsgType(msgType string) *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

func randomValue() *big.Int { _ = "STUB: not implemented"; return nil }

// GenerateSendFundsTx returns a random send funds tx
//
//nolint:staticcheck
func (txClient *EvmTxClient) GenerateSendFundsTx() *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// GenerateERC20TransferTx returns a random ERC20 send
// the contract it interacts with needs no funding (infinite balances)
func (txClient *EvmTxClient) GenerateERC20TransferTx() *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// override gas limit for an ERC20 transfer

func (txClient *EvmTxClient) GenerateUniV2SwapTx() *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

func (txClient *EvmTxClient) GenerateERC721Mint() *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// override gas limit for an ERC20 transfer

func (txClient *EvmTxClient) getTransactOpts() *bind.TransactOpts {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func (txClient *EvmTxClient) sign(tx *ethtypes.Transaction) *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// this should not happen

func (txClient *EvmTxClient) nextNonce() uint64 { _ = "STUB: not implemented"; return 0 }

// SendEvmTx takes any signed evm tx and send it out
func (txClient *EvmTxClient) SendEvmTx(signedTx *ethtypes.Transaction, onSuccess func()) {
	_ = "STUB: not implemented"
	return
}

// GetNextEthClient return the next available eth client randomly
//
//nolint:staticcheck
func GetNextEthClient(clients []*ethclient.Client) *ethclient.Client {
	_ = "STUB: not implemented"
	return nil
}

// GetTxReceipt query the transaction receipt to check if the tx succeed or not
func (txClient *EvmTxClient) GetTxReceipt(txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// check receipt success
func (txClient *EvmTxClient) EnsureTxSuccess(txHash common.Hash) { _ = "STUB: not implemented"; return }

// ResetNonce need to be called when tx failed
func (txClient *EvmTxClient) ResetNonce() error { _ = "STUB: not implemented"; return nil }
