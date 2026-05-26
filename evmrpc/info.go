package evmrpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

const DefaultBlockGasLimit = 10000000
const defaultPriorityFeePerGas = 1000000000 // 1gwei
const defaultThresholdPercentage = 80       // 80%

type InfoAPI struct {
	tmClient         client.LocalClient
	keeper           *keeper.Keeper
	ctxProvider      func(int64) sdk.Context
	txConfigProvider func(int64) client.TxConfig
	homeDir          string
	connectionType   ConnectionType
	maxBlocks        int64
	txDecoder        sdk.TxDecoder
	watermarks       *WatermarkManager
}

func NewInfoAPI(tmClient client.LocalClient, k *keeper.Keeper, ctxProvider func(int64) sdk.Context, txConfigProvider func(int64) client.TxConfig, homeDir string, maxBlocks int64, connectionType ConnectionType, txDecoder sdk.TxDecoder, watermarks *WatermarkManager) *InfoAPI {
	_ = "STUB: not implemented"
	return nil
}

type FeeHistoryResult struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

func (i *InfoAPI) BlockNumber(ctx context.Context) hexutil.Uint64 {
	_ = "STUB: not implemented"
	return *new(hexutil.Uint64)
}

//nolint:gosec

//nolint:revive
func (i *InfoAPI) ChainId(ctx context.Context) *hexutil.Big { _ = "STUB: not implemented"; return nil }

func (i *InfoAPI) Coinbase(ctx context.Context) (addr common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (i *InfoAPI) Accounts(ctx context.Context) (result []common.Address, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *InfoAPI) GasPrice(ctx context.Context) (result *hexutil.Big, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Helper function useful for testing
func (i *InfoAPI) GasPriceHelper(ctx context.Context, baseFee *big.Int, totalGasUsedPrevBlock uint64, medianRewardPrevBlock *big.Int) (*hexutil.Big, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// chain is not congested, increase base fee by 10% to get the gas price to get a tx included in a timely manner

// chain is congested, return the 50%-tile reward as the priority fee per gas

// lastBlock is inclusive
func (i *InfoAPI) FeeHistory(ctx context.Context, blockCount gmath.HexOrDecimal64, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (result *FeeHistoryResult, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// logic consistent with go-ethereum's validation (block < 1 means no block)

// default go-ethereum max block history is 1024
// https://github.com/ethereum/go-ethereum/blob/master/eth/gasprice/feehistory.go#L235
//nolint:gosec

// if someone needs more than 100 reward percentiles, we can discuss, but it's not likely

// validate reward percentiles

// fall back to genesis height if earliest watermark unavailable

//nolint:gosec

//nolint:gosec

// True only after we append header base fee for lastBlockNumber (avoids redundant CheckVersion and
// avoids appending a child base fee when the last block had no header entry, e.g. pruned base fee).

// Potentially parallelize the following logic

// either height is pruned or before EVM is introduced
// For non-EVM blocks or pruned blocks, use 0.0 as gas used ratio

// Calculate actual gas used ratio for this block

// If we can't calculate the ratio, use 0.0 as fallback

// Only continue with other fields if EVM state exists

// the block has been pruned

// block pruned from tendermint store. Skipping

// execution-apis eth_feeHistory / go-ethereum: baseFeePerGas has one more element than gasUsedRatio,
// the projected base fee for the child of the newest block in the range.
// Note: len(baseFeePerGas) may still differ from len(gasUsedRatio)+1 when some heights skip header
// base fees (pruned / partial data) while gasUsedRatio rows exist — same class of partial history as before.

func (i *InfoAPI) MaxPriorityFeePerGas(ctx context.Context) (fee *hexutil.Big, returnErr error) {
	_ = "STUB: not implemented"
	// Checks the most recent block. If it has high gas used, it will return the reward of the 50% percentile.
	// Otherwise, since the previous block has low gas used, a user shouldn't need to tip a high amount to get included,
	// so a default value is returned.
	return nil, nil
}

// chain is not congested, return 1gwei as the default priority fee per gas

// chain is congested, return the 50%-tile reward as the priority fee per gas

// if there is no EVM tx in the most recent block, return 0

func (i *InfoAPI) BlobBaseFee(ctx context.Context) (result *hexutil.Big, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Syncing implements eth_syncing. It is intentionally registered (not removed): the RPC returns
// JSON-RPC error -32000 with a clear message instead of -32601 method not found. Ethereum returns
// false or a sync object; Sei does not expose sync semantics on this API.
func (i *InfoAPI) Syncing(ctx context.Context) (result any, returnErr error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// safeGetHeaderBaseFee returns the base fee per gas for txs in block blockNum (same as eth block header
// and encodeRPCTransaction: GetNextBaseFee at parent committed height).
func (i *InfoAPI) safeGetHeaderBaseFee(blockNum int64) (res *big.Int) {
	_ = "STUB: not implemented"
	return nil
}

// safeGetChildBaseFeeAfter returns the base fee for the block after parentBlockNum (GetNextBaseFee at end of parentBlockNum).
func (i *InfoAPI) safeGetChildBaseFeeAfter(parentBlockNum int64) (res *big.Int) {
	_ = "STUB: not implemented"
	return nil
}

type GasAndReward struct {
	GasUsed uint64
	Reward  *big.Int
}

func (i *InfoAPI) getRewards(block *coretypes.ResultBlock, baseFee *big.Int, rewardPercentiles []float64) ([]*hexutil.Big, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// not evm tx

// okay to get from latest since receipt is immutable

// tx doesn't have a receipt because of nonce mismatch

// if effective gas price is 0, it's expected behavior for txs that failed ante.
// if it's not zero but still smaller than baseFee then something is wrong.

func (i *InfoAPI) getCongestionData(ctx context.Context, height *int64) (blockGasUsed uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// block pruned from tendermint store. Skipping

// not evm tx

// okay to get from latest since receipt is immutable

// We've had issues where is included in a block and fails but then is retried and included in a later block, overwriting the receipt.
// This is a temporary fix to ensure we only consider receipts that are included in the block we're querying.
//nolint:gosec

// CalculateGasUsedRatio calculates the actual gas used ratio for a specific block
func (i *InfoAPI) CalculateGasUsedRatio(ctx context.Context, blockHeight int64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Get the gas limit from consensus params using the SDK context

//nolint:gosec

// Fallback: try current context

//nolint:gosec

// Default fallback
// Default block gas limit for Sei

// Avoid division by zero

// Calculate total gas used by EVM transactions in this block

// not evm tx

// okay to get from latest since receipt is immutable

// We've had issues where tx is included in a block and fails but then is retried and included in a later block, overwriting the receipt.
// This is a temporary fix to ensure we only consider receipts that are included in the block we're querying.
//nolint:gosec

// We want 4 decimal places, so multiply by 10000, do integer division, then divide by 10000
// This preserves more precision during the integer calculation

func (i *InfoAPI) latestHeight(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *InfoAPI) earliestHeight(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Following go-ethereum implementation
// Specifically, the reward value at a percentile of p% will be the reward value of the
// lowest-rewarded transaction such that the sum of its gasUsed value and gasUsed values
// of all lower-rewarded transactions is no less than (total gasUsed * p%).
func CalculatePercentiles(rewardPercentiles []float64, GasAndRewards []GasAndReward, totalEVMGasUsed uint64) []*hexutil.Big {
	_ = "STUB: not implemented"
	return nil
}

// Return array of zeros for each percentile when no transactions exist

func (i *InfoAPI) isChainCongested(totalGasUsed uint64) bool {
	_ = "STUB: not implemented"
	return false
}
