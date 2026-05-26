package rpcutils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/derived"
)

var signerMap = map[derived.SignerVersion]func(*big.Int) ethtypes.Signer{
	derived.London: ethtypes.NewLondonSigner,
	derived.Cancun: ethtypes.NewCancunSigner,
	derived.Prague: ethtypes.NewPragueSigner,
}

// RecoverEVMSender recovers the sender address from an Ethereum transaction
// using the same logic as the preprocess ante handler.
// This ensures consistency between transaction preprocessing and RPC queries.
func RecoverEVMSender(ethTx *ethtypes.Transaction, blockHeight int64, blockTime int64) (common.Address, error) {
	_ = "STUB: not implemented"
	// Get the chain ID from the transaction
	return *new(common.Address), nil
}

// Get the chain config and determine the signer version

// Create the signer with the transaction's chain ID

// Get raw signature values

// Compute the transaction hash based on whether it's protected

// For protected transactions, adjust V and use signer hash

// For unprotected transactions, use Frontier signer

// Recover the sender address

// adjustV adjusts the V value for signature recovery based on transaction type and chain ID
func adjustV(V *big.Int, txType uint8, chainID *big.Int) *big.Int {
	_ = "STUB: not implemented"
	// Non-legacy TX always needs to be bumped by 27
	return nil
}

// Legacy TX needs to be adjusted based on chainID
// Formula: V = V - (chainID * 2) - 8

// getSignerVersion determines which signer version to use based on block height and time
func getSignerVersion(blockHeight int64, blockTime uint64, ethCfg *params.ChainConfig) derived.SignerVersion {
	_ = "STUB: not implemented"
	return *new(derived.SignerVersion)
}

// RecoverEVMSenderWithContext is a convenience wrapper that extracts block info from context
func RecoverEVMSenderWithContext(ctx sdk.Context, ethTx *ethtypes.Transaction) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}
