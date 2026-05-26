package cryptosim

import (
	"hash"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sei-protocol/sei-chain/sei-db/common/keys"
	crand "github.com/sei-protocol/sei-chain/sei-db/common/rand"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
)

const (
	erc20TransferEventSignatureHex = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

	// These mirror immutable memiavl EVM key prefixes and are duplicated here to keep the hot path minimal.
	evmCodeKeyPrefixByte     = 0x07
	evmCodeHashKeyPrefixByte = 0x08
	evmStorageKeyPrefixByte  = 0x03

	hashLen            = 32
	indexedAddressBase = hashLen - keys.AddressLen

	syntheticReceiptMinBlockNumber uint64 = 1_000_000

	syntheticReceiptGasUsedBase     uint64 = 52_000
	syntheticReceiptGasUsedSpan     uint64 = 18_000
	syntheticReceiptPreviousGasBase uint64 = 21_000
	syntheticReceiptPreviousGasSpan uint64 = 35_000
	syntheticReceiptGasPriceBase    uint64 = 1_000_000_000
	syntheticReceiptGasPriceSpan    uint64 = 9_000_000_000
	syntheticReceiptTransferBase    uint64 = 1_000_000
	syntheticReceiptTransferSpan    uint64 = 10_000_000_000

	// Multiplied by blockNumber then added to txIndex to produce a unique seed per
	// transaction. Supports up to 1M txs per block before collisions. With int64,
	// block numbers up to ~9.2 trillion are safe before overflow (~290k years at 1 block/sec).
	syntheticTxIDBlockStride int64 = 1_000_000
)

var erc20TransferEventSignatureBytes = [hashLen]byte{
	0xdd, 0xf2, 0x52, 0xad, 0x1b, 0xe2, 0xc8, 0x9b,
	0x69, 0xc2, 0xb0, 0x68, 0xfc, 0x37, 0x8d, 0xaa,
	0x95, 0x2b, 0xa7, 0xf1, 0x63, 0xc4, 0xa1, 0x16,
	0x28, 0xf5, 0x5a, 0x4d, 0xf5, 0x23, 0xb3, 0xef,
}

// SyntheticTxHash returns a deterministic 32-byte tx hash for a given (blockNumber, txIndex) pair.
//
// It uses CannedRandom.SeededBytes, which is a pure read from the pre-generated buffer — no
// internal state is advanced, and the result depends only on the CannedRandom's seed/buffer
// and the inputs. This means any goroutine with a CannedRandom created from the same
// (seed, bufferSize) can reconstruct any tx hash from just the block number and tx index,
// without storing the hashes. Readers use this to compute query targets on the fly:
//
//	validRange  = [max(1, latestBlock - keepRecent + 1), latestBlock]
//	randomBlock = pick from validRange
//	randomTxIdx = pick from [0, txsPerBlock)
//	txHash      = SyntheticTxHash(crand, randomBlock, randomTxIdx)
//
// The hash automatically becomes invalid (returns no result) once the corresponding
// parquet file is pruned, so readers never need to track which hashes are live.
func SyntheticTxHash(crand *crand.CannedRandom, blockNumber uint64, txIndex uint32) []byte {
	_ = "STUB: not implemented"
	//nolint:gosec // block numbers and tx indices won't exceed int64 in benchmarks
	return nil
}

// BuildERC20TransferReceiptFromTxn produces a plausible successful ERC20 transfer receipt from a transaction.
func BuildERC20TransferReceiptFromTxn(
	crand *crand.CannedRandom,
	feeCollectionAccount []byte,
	blockNumber uint64,
	txIndex uint32,
	txn *transaction,
) (*evmtypes.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildERC20TransferReceipt produces a plausible successful ERC20 transfer receipt.
//
// The sender and receiver are derived from the address portion of the supplied storage keys, since cryptosim tracks
// ERC20 balances as storage slots rather than separate account references. The caller supplies the block number and tx
// index so the resulting receipt can line up with the simulated block being benchmarked.
func BuildERC20TransferReceipt(
	crand *crand.CannedRandom,
	feeCollectionAccount []byte,
	srcAccount []byte,
	dstAccount []byte,
	senderSlot []byte,
	receiverSlot []byte,
	erc20ContractCode []byte,
	blockNumber uint64,
	txIndex uint32,
) (*evmtypes.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // constants fit in int64

//nolint:gosec // constants fit in int64

//nolint:gosec // constants fit in int64

//nolint:gosec // constants fit in int64

func validateAccountKey(name string, key []byte) error { _ = "STUB: not implemented"; return nil }

// extractAccountKeyBytes accepts keys with either EVMKeyCode (0x07) or EVMKeyCodeHash (0x08) prefix,
// since cryptosim uses EVMKeyCodeHash for accounts while ERC20 contracts use EVMKeyCode.
func extractAccountKeyBytes(name string, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractCodeKeyBytes(name string, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractStorageKeyAddressBytes(name string, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addToBloom(hasher hash.Hash, digest *[hashLen]byte, bloom *ethtypes.Bloom, value []byte) {
	_ = "STUB: not implemented"
	return
}

func encodeUint256FromUint64(value uint64) []byte { _ = "STUB: not implemented"; return nil }
