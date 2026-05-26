package rosetta

import (
	"context"
	"time"

	rosettatypes "github.com/coinbase/rosetta-sdk-go/types"

	crgtypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/rosetta/lib/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	auth "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	bank "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"

	tmrpc "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
)

// interface assertion
var _ crgtypes.Client = (*Client)(nil)

const defaultNodeTimeout = 15 * time.Second

// Client implements a single network client to interact with cosmos based chains
type Client struct {
	supportedOperations []string

	config *Config

	auth  auth.QueryClient
	bank  bank.QueryClient
	tmRPC tmrpc.Client

	version string

	converter Converter
}

// NewClient instantiates a new online servicer
func NewClient(cfg *Config) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// ---------- cosmos-rosetta-gateway.types.Client implementation ------------ //

// Bootstrap is gonna connect the client to the endpoints
func (c *Client) Bootstrap() error { _ = "STUB: not implemented"; return nil }

// Ready performs a health check and returns an error if the client is not ready.
func (c *Client) Ready() error { _ = "STUB: not implemented"; return nil }

func (c *Client) accountInfo(ctx context.Context, addr string, height *int64) (*SignerData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Balances(ctx context.Context, addr string, height *int64) ([]*rosettatypes.Amount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BlockByHash(ctx context.Context, hash string) (crgtypes.BlockResponse, error) {
	_ = "STUB: not implemented"
	return *new(crgtypes.BlockResponse), nil
}

func (c *Client) BlockByHeight(ctx context.Context, height *int64) (crgtypes.BlockResponse, error) {
	_ = "STUB: not implemented"
	return *new(crgtypes.BlockResponse), nil
}

func (c *Client) BlockTransactionsByHash(ctx context.Context, hash string) (crgtypes.BlockTransactionsResponse, error) {
	_ = "STUB: not implemented"
	// TODO(fdymylja): use a faster path, by searching the block by hash, instead of doing a double query operation
	return *new(crgtypes.BlockTransactionsResponse), nil
}

func (c *Client) BlockTransactionsByHeight(ctx context.Context, height *int64) (crgtypes.BlockTransactionsResponse, error) {
	_ = "STUB: not implemented"
	return *new(crgtypes.BlockTransactionsResponse), nil
}

// Coins fetches the existing coins in the application
func (c *Client) coins(ctx context.Context) (sdk.Coins, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

func (c *Client) TxOperationsAndSignersAccountIdentifiers(signed bool, txBytes []byte) (ops []*rosettatypes.Operation, signers []*rosettatypes.AccountIdentifier, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetTx returns a transaction given its hash. For Rosetta we  make a synthetic transaction for BeginBlock
//
//	and EndBlock to adhere to balance tracking rules.
func (c *Client) GetTx(ctx context.Context, hash string) (*rosettatypes.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get tx type and hash

// construct rosetta tx

// handle begin block hash

// get block height by hash

// get block txs

// handle end block hash

// get block height by hash

// get block txs

// get last tx

// unrecognized tx

// GetUnconfirmedTx gets an unconfirmed transaction given its hash
func (c *Client) GetUnconfirmedTx(ctx context.Context, hash string) (*rosettatypes.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assert that correct tx length is provided

// iterate over unconfirmed txs to find the one with matching hash

// Mempool returns the unconfirmed transactions in the mempool
func (c *Client) Mempool(ctx context.Context) ([]*rosettatypes.TransactionIdentifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Peers gets the number of peers
func (c *Client) Peers(ctx context.Context) ([]*rosettatypes.Peer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Status(ctx context.Context) (*rosettatypes.SyncStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) PostTx(txBytes []byte) (*rosettatypes.TransactionIdentifier, map[string]interface{}, error) {
	_ = "STUB: not implemented"
	// sync ensures it will go through checkTx
	return nil, nil, nil
}

// check if tx was broadcast successfully

// construction endpoints

// ConstructionMetadataFromOptions builds the metadata given the options
func (c *Client) ConstructionMetadataFromOptions(ctx context.Context, options map[string]interface{}) (meta map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) blockTxs(ctx context.Context, height *int64) (crgtypes.BlockTransactionsResponse, error) {
	_ = "STUB: not implemented"
	// get block info
	return *new(crgtypes.BlockTransactionsResponse), nil
}

// get block events

// wtf?

// process begin and end block txs

// process normal txs

func (c *Client) getHeight(ctx context.Context, height *int64) (realHeight *int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractInitialHeightFromGenesisChunk(genesisChunk string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gocritic
