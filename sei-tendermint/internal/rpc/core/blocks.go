package core

import (
	"context"

	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
)

// BlockchainInfo gets block headers for minHeight <= height <= maxHeight.
//
// If maxHeight does not yet exist, blocks up to the current height will be
// returned. If minHeight does not exist (due to pruning), earliest existing
// height will be used.
//
// At most 20 items will be returned. Block headers are returned in descending
// order (highest first).
//
// More: https://docs.tendermint.com/master/rpc/#/Info/blockchain
func (env *Environment) BlockchainInfo(ctx context.Context, req *coretypes.RequestBlockchainInfo) (*coretypes.ResultBlockchainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// error if either min or max are negative or min > max
// if 0, use blockstore base for min, latest block height for max
// enforce limit.
func filterMinMax(base, height, min, max, limit int64) (int64, int64, error) {
	_ = "STUB: not implemented"
	// filter negatives
	return 0, 0, nil
}

// adjust for default values

// limit max to the height

// limit min to the base

// limit min to within `limit` of max
// so the total number of blocks returned will be `limit`

// Block gets block at a given height.
// If no height is provided, it will fetch the latest block.
// More: https://docs.tendermint.com/master/rpc/#/Info/block
//
// Under Autobahn the CometBFT BlockStore is not populated; route through
// GigaRouter, which reads the finalized global block from data.State and
// returns it in the same coretypes.ResultBlock shape. This keeps every
// downstream consumer (evmrpc, /block HTTP, seid q block) working without
// individually branching on consensus mode.
func (env *Environment) Block(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cast happens at the boundary so giga.BlockByNumber stays
// strongly typed on atypes.GlobalBlockNumber. autobahnCheckAndGetHeight
// has already validated height is positive and within chain head.

// autobahnCheckAndGetHeight resolves a caller-supplied height pointer to a
// concrete int64 and validates it against the current ABCI head. nil (or
// zero) means "latest". Returns the same error sentinels as env.getHeight
// (ErrZeroOrNegativeHeight, ErrHeightExceedsChainHead, ErrHeightNotAvailable)
// so downstream consumers like evmrpc — which translate
// ErrHeightExceedsChainHead-class errors into the Ethereum-spec `null`
// response for non-existent blocks — see consistent shapes under both
// consensus engines.
//
// TODO(autobahn): wire a real lower bound and pass it as `base` to
// env.getHeight. We currently pass env.BlockStore.Base() (always 0 under
// Autobahn), which means any positive height < chain head passes validation
// here and is rejected one layer down (data.GlobalBlock returns
// data.ErrPruned, which BlockByNumber maps to ErrHeightNotAvailable). With
// a real lower bound the rejection happens at this layer instead. The
// natural source becomes available once sei-db/ledger_db/block.BlockDB is
// wired into block execution: switch this and BlockByNumber to read from
// BlockDB, and source `base` from BlockDB.GetLowestBlockHeight.
func (env *Environment) autobahnCheckAndGetHeight(ctx context.Context, heightPtr *int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// BlockByHash gets block by hash.
// More: https://docs.tendermint.com/master/rpc/#/Info/block_by_hash
//
// Under Autobahn the CometBFT BlockStore is not populated, so we route
// through the GigaRouter's temporary in-memory hash index. Match CometBFT
// semantics: an unknown hash returns &ResultBlock{Block: nil} with no
// error, never an error response — external tools (block explorers,
// monitoring) treat that as "no such block" rather than a failure.
//
// The Tendermint RPC boundary (req.Hash is bytes.HexBytes — a []byte alias
// for wire-format flexibility) gets converted to the strongly-typed
// atypes.BlockHeaderHash here, before reaching GigaRouter.BlockByHash.
// Wrong-size inputs short-circuit to the same zero-result CometBFT
// returns for an unknown hash.
func (env *Environment) BlockByHash(ctx context.Context, req *coretypes.RequestBlockByHash) (*coretypes.ResultBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If block is not nil, then blockMeta can't be nil.

// Header gets block header at a given height.
// If no height is provided, it will fetch the latest header.
// More: https://docs.tendermint.com/master/rpc/#/Info/header
func (env *Environment) Header(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HeaderByHash gets header by hash.
// More: https://docs.tendermint.com/master/rpc/#/Info/header_by_hash
func (env *Environment) HeaderByHash(ctx context.Context, req *coretypes.RequestBlockByHash) (*coretypes.ResultHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Commit gets block commit at a given height.
// If no height is provided, it will fetch the commit for the latest block.
// More: https://docs.tendermint.com/master/rpc/#/Info/commit
func (env *Environment) Commit(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the next block has not been committed yet,
// use a non-canonical commit

// NOTE: we can't yet ensure atomicity of operations in asserting
// whether this is the latest height and retrieving the seen commit

// Return the canonical commit (comes from the block at height+1)

// BlockResults gets ABCIResults at a given height.
// If no height is provided, it will fetch results for the latest block.
//
// Results are for the height of the block containing the txs.
// More: https://docs.tendermint.com/master/rpc/#/Info/block_results
//
// Under Autobahn, FinalizeBlock responses are not persisted to StateStore
// (giga_router.executeBlock never calls SaveFinalizeBlockResponses), so this
// returns a valid-but-empty ResultBlockResults at the requested height. That
// lets downstream consumers (evmrpc's eth_getBlockByNumber, which looks up
// BlockResults to enrich the response) keep working — the block envelope
// renders correctly, just without per-tx ExecTxResult details. Properly
// populating these under Autobahn is a separate follow-up.
func (env *Environment) BlockResults(ctx context.Context, req *coretypes.RequestBlockInfo) (*coretypes.ResultBlockResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// evmrpc's EncodeTmBlock reads ConsensusParamUpdates.Block.MaxGas to
// populate the eth_getBlockByNumber gasLimit field. Populate it from
// Autobahn's producer config so that path doesn't nil-deref.

// BlockSearch searches for a paginated set of blocks matching the provided query.
func (env *Environment) BlockSearch(ctx context.Context, req *coretypes.RequestBlockSearch) (*coretypes.ResultBlockSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort results (must be done before pagination)

// paginate results
