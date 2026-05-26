package harness

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// BuildTransaction creates an Ethereum transaction from state test data
// Note: Rebuilds the transaction with Sei's chain ID since fixtures use chain ID 1
func BuildTransaction(st *StateTestJSON, subtest StateTestPost) (*ethtypes.Transaction, common.Address, error) {
	_ = "STUB: not implemented"
	// Get the private key
	return nil, *new(common.Address), nil
}

// Use sender from transaction if available, otherwise derive from key

// Get indexed values

// Parse data

// Parse gas limit

// Parse value

// Parse to address

// Parse nonce

// Parse gas prices

// Determine transaction type and create accordingly

// EIP-1559 transaction

// Use Sei's chain ID

// EIP-2930 transaction

// Legacy transaction

// Sign the transaction with Sei's chain ID

// EncodeTxForApp encodes a signed transaction for the Sei app
func EncodeTxForApp(signedTx *ethtypes.Transaction) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseHexBig parses a hex string (with possible leading zeros) to *big.Int
func parseHexBig(s string) *big.Int { _ = "STUB: not implemented"; return nil }

// parseHexUint64 parses a hex string to uint64
func parseHexUint64(s string) uint64 { _ = "STUB: not implemented"; return 0 }
