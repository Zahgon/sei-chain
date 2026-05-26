package rpc

import (
	"context"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/merkle"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	service "github.com/sei-protocol/sei-chain/sei-tendermint/libs/service"
	rpcclient "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/client"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

// KeyPathFunc builds a merkle path out of the given path and key.
type KeyPathFunc func(path string, key []byte) (merkle.KeyPath, error)

// LightClient is an interface that contains functionality needed by Client from the light client.
//
//go:generate ../../scripts/mockery_generate.sh LightClient
type LightClient interface {
	ChainID() string
	Update(ctx context.Context, now time.Time) (*types.LightBlock, error)
	VerifyLightBlockAtHeight(ctx context.Context, height int64, now time.Time) (*types.LightBlock, error)
	TrustedLightBlock(height int64) (*types.LightBlock, error)
	Status(ctx context.Context) *types.LightClientInfo
}

var _ rpcclient.Client = (*Client)(nil)

// Client is an RPC client, which uses light#Client to verify data (if it can
// be proved). Note, merkle.DefaultProofRuntime is used to verify values
// returned by ABCI#Query.
type Client struct {
	service.BaseService

	next rpcclient.Client
	lc   LightClient

	// proof runtime used to verify values returned by ABCIQuery
	prt       *merkle.ProofRuntime
	keyPathFn KeyPathFunc

	closers []func()
}

// Option allow you to tweak Client.
type Option func(*Client)

// KeyPathFn option can be used to set a function, which parses a given path
// and builds the merkle path for the prover. It must be provided if you want
// to call ABCIQuery or ABCIQueryWithOptions.
func KeyPathFn(fn KeyPathFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultMerkleKeyPathFn creates a function used to generate merkle key paths
// from a path string and a key. This is the default used by the cosmos SDK.
// This merkle key paths are required when verifying /abci_query calls
func DefaultMerkleKeyPathFn() KeyPathFunc {
	_ = "STUB: not implemented"
	// regexp for extracting store name from /abci_query path
	return *new(KeyPathFunc)
}

// NewClient returns a new client.
func NewClient(next rpcclient.Client, lc LightClient, opts ...Option) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) OnStart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Client) OnStop() { _ = "STUB: not implemented"; return }

// Returns the status of the light client. Previously this was querying the primary connected to the client
// As a consequence of this change, running /status on the light client will return nil for SyncInfo, NodeInfo
// and ValdiatorInfo.
func (c *Client) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LagStatus  return 0 for lag status
func (c *Client) LagStatus(ctx context.Context) (*coretypes.ResultLagStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	_ = "STUB: not implemented"
	return nil,

		// ABCIQuery requests proof by default.
		nil
}

func (c *Client) ABCIQuery(ctx context.Context, path string, data tmbytes.HexBytes) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ABCIQueryWithOptions returns an error if opts.Prove is false.
// ABCIQueryWithOptions returns the result for the given height (opts.Height).
// If no height is provided, the results of the block preceding the latest are returned.
func (c *Client) ABCIQueryWithOptions(ctx context.Context, path string, data tmbytes.HexBytes,
	opts rpcclient.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	_ = "STUB: not implemented"

	// always request the proof
	return nil, nil
}

// Can't return the latest block results because we won't be able to
// prove them. Return the results for the previous block instead.

// Validate the response.

// Update the light client if we're behind.
// NOTE: AppHash for height H is in header H+1.

// Validate the value proof against the trusted header.

// build a Merkle key path from path and resp.Key

// verify value

// OR validate the absence proof against the trusted header.

func (c *Client) BroadcastTxCommit(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BroadcastTxAsync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BroadcastTxSync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BroadcastTx(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) UnconfirmedTxs(ctx context.Context, page, perPage *int) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) NumUnconfirmedTxs(ctx context.Context) (*coretypes.ResultUnconfirmedTxs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) CheckTx(ctx context.Context, tx types.Tx) (*coretypes.ResultCheckTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) NetInfo(ctx context.Context) (*coretypes.ResultNetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) DumpConsensusState(ctx context.Context) (*coretypes.ResultDumpConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ConsensusState(ctx context.Context) (*coretypes.ResultConsensusState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ConsensusParams(ctx context.Context, height *int64) (*coretypes.ResultConsensusParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify hash.

func (c *Client) Events(ctx context.Context, req *coretypes.RequestEvents) (*coretypes.ResultEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Health(ctx context.Context) (*coretypes.ResultHealth, error) {
	_ = "STUB: not implemented"
	return nil,

		// BlockchainInfo calls rpcclient#BlockchainInfo and then verifies every header
		// returned.
		nil
}

func (c *Client) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify each of the BlockMetas.

func (c *Client) Genesis(ctx context.Context) (*coretypes.ResultGenesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GenesisChunked(ctx context.Context, id uint) (*coretypes.ResultGenesisChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Block calls rpcclient#Block and then verifies the result.
func (c *Client) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify block.

// BlockByHash calls rpcclient#BlockByHash and then verifies the result.
func (c *Client) BlockByHash(ctx context.Context, hash tmbytes.HexBytes) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Verify block.

// BlockResults returns the block results for the given height. If no height is
// provided, the results of the block preceding the latest are returned.
func (c *Client) BlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Can't return the latest block results here because we won't be able to
// prove them. Return the results for the previous block instead.

// Validate res.

// Update the light client if we're behind.

// proto-encode FinalizeBlock events

// Build a Merkle tree out of the slice.

// Verify block results.

// Header fetches and verifies the header directly via the light client
func (c *Client) Header(ctx context.Context, height *int64) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HeaderByHash calls rpcclient#HeaderByHash and updates the client if it's falling behind.
func (c *Client) HeaderByHash(ctx context.Context, hash tmbytes.HexBytes) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Commit(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	// Update the light client if we're behind and retrieve the light block at the requested height
	// or at the latest height if no height is provided.
	return nil, nil
}

// Tx calls rpcclient#Tx method and then verifies the proof if such was
// requested.
func (c *Client) Tx(ctx context.Context, hash tmbytes.HexBytes, prove bool) (*coretypes.ResultTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate res.

// Update the light client if we're behind.

// Validate the proof.

func (c *Client) TxSearch(
	ctx context.Context,
	query string,
	prove bool,
	page, perPage *int,
	orderBy string,
) (*coretypes.ResultTxSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BlockSearch(
	ctx context.Context,
	query string,
	page, perPage *int,
	orderBy string,
) (*coretypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validators fetches and verifies validators.
func (c *Client) Validators(
	ctx context.Context,
	height *int64,
	pagePtr, perPagePtr *int,
) (*coretypes.ResultValidators, error) {
	_ = "STUB: not implemented"

	// Update the light client if we're behind and retrieve the light block at the
	// requested height or at the latest height if no height is provided.
	return nil, nil
}

//nolint:gosec // perPage is bounded to maxPerPage (100); no overflow risk

func (c *Client) BroadcastEvidence(ctx context.Context, ev types.Evidence) (*coretypes.ResultBroadcastEvidence, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Subscribe(ctx context.Context, subscriber, query string,
	outCapacity ...int) (out <-chan coretypes.ResultEvent, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

func (c *Client) Unsubscribe(ctx context.Context, subscriber, query string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func (c *Client) UnsubscribeAll(ctx context.Context, subscriber string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func (c *Client) updateLightClientIfNeededTo(ctx context.Context, height *int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) RegisterOpDecoder(typ string, dec merkle.OpDecoder) {
	_ = "STUB: not implemented"
	return
}

// SubscribeWS subscribes for events using the given query and remote address as
// a subscriber, but does not verify responses (UNSAFE)!
// TODO: verify data
func (c *Client) SubscribeWS(ctx context.Context, query string) (*coretypes.ResultSubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

// We should have a switch here that performs a validation
// depending on the event's type.

// UnsubscribeWS calls original client's Unsubscribe using remote address as a
// subscriber.
func (c *Client) UnsubscribeWS(ctx context.Context, query string) (*coretypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

// UnsubscribeAllWS calls original client's UnsubscribeAll using remote address
// as a subscriber.
func (c *Client) UnsubscribeAllWS(ctx context.Context) (*coretypes.ResultUnsubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

// XXX: Copied from rpc/core/env.go
const (
	// see README
	defaultPerPage = 30
	maxPerPage     = 100
)

func validatePage(pagePtr *int, perPage uint, totalCount int) (int, error) {
	_ = "STUB: not implemented"

	// no page parameter
	return 0, nil
}

//nolint:gosec // perPage is bounded to maxPerPage (100); no overflow risk

// one page (even if it's empty)

func validatePerPage(perPagePtr *int) uint { _ = "STUB: not implemented"; return 0 }

// no per_page parameter

//nolint:gosec // perPage is bounds-checked above to [1, maxPerPage]

func validateSkipCount(page int, perPage uint) int { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // perPage is bounded to maxPerPage (100) and page is validated; no overflow risk
