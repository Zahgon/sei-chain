package main

import (
	"sync"

	"golang.org/x/time/rate"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	typestx "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
)

type LoadTestClient struct {
	LoadTestConfig     Config
	TestConfig         EncodingConfig
	AccountKeys        []cryptotypes.PrivKey
	TxClients          []typestx.ServiceClient
	EvmTxClients       []*EvmTxClient
	SignerClient       *SignerClient
	ChainID            string
	GrpcConns          []*grpc.ClientConn
	StakingQueryClient stakingtypes.QueryClient
	// Staking specific variables
	Validators []stakingtypes.Validator
	// DelegationMap is a map of delegator -> validator -> delegated amount
	DelegationMap map[string]map[string]int
	// Tokenfactory specific variables
	TokenFactoryDenomOwner map[string]string
	// Only one admin message can go in per block
	generatedAdminMessageForBlock bool
	// Messages that has to be sent from the admin account
	isAdminMessageMapping map[string]bool
	mtx                   *sync.RWMutex
}

func NewLoadTestClient(config Config) *LoadTestClient { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// Fill message type maps with empty values

func (c *LoadTestClient) SetValidators() { _ = "STUB: not implemented"; return }

// BuildGrpcClients build a list of grpc clients
func BuildGrpcClients(config *Config) ([]typestx.ServiceClient, []*grpc.ClientConn) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // Use insecure skip verify.

// spin up goroutine for monitoring and reconnect purposes

// BuildEvmTxClients build a list of EvmTxClients with a list of go-ethereum client
func BuildEvmTxClients(config *Config, keys []cryptotypes.PrivKey) []*EvmTxClient {
	_ = "STUB: not implemented"
	return nil
}

// Get chainId

// Get gas price

// Build one client per key

func (c *LoadTestClient) Close() { _ = "STUB: not implemented"; return }

func (c *LoadTestClient) BuildTxs(
	txQueue chan SignedTx,
	keyIndex int,
	wg *sync.WaitGroup,
	done <-chan struct{},
	rateLimiter *rate.Limiter,
	producedCountPerMsgType map[string]*int64,
) {
	_ = "STUB: not implemented"
	return
}

// Generate a message type first

// Sign EVM and Cosmos TX differently

func (c *LoadTestClient) generateSignedEvmTx(keyIndex int, msgType string) *ethtypes.Transaction {
	_ = "STUB: not implemented"
	return nil
}

func (c *LoadTestClient) generateSignedCosmosTxs(keyIndex int, msgType string, msgTypeCount int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use random seqno to get around txs that might already be seen in mempool
//nolint:gosec // msgTypeCount is a small positive counter; no overflow risk

func (c *LoadTestClient) SendTxs(
	txQueue chan SignedTx,
	keyIndex int,
	done <-chan struct{},
	sentCountPerMsgType map[string]*int64,
	semaphore *semaphore.Weighted,
	wg *sync.WaitGroup,
) {
	_ = "STUB: not implemented"
	return
}

// Acquire a semaphore

// Send Cosmos Transactions

// Send EVM Transactions

// Release the semaphore

//nolint:staticcheck
func (c *LoadTestClient) GetTxClient() typestx.ServiceClient {
	_ = "STUB: not implemented"
	return *new(typestx.ServiceClient)
}
