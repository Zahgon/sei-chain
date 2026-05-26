package helpers

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/secp256k1"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	big2  = big.NewInt(2)
	big8  = big.NewInt(8)
	big27 = big.NewInt(27)
)

// AdjustV adjusts the V value from a raw signature for pubkey recovery.
// For non-legacy transactions, V is bumped by 27.
// For legacy transactions, V is adjusted based on chainID per EIP-155.
// This function is used by both the EVM ante handler and the Giga executor.
func AdjustV(V *big.Int, txType uint8, chainID *big.Int) *big.Int {
	_ = "STUB: not implemented"
	// Non-legacy TX always needs to be bumped by 27
	return nil
}

// Legacy TX needs to be adjusted based on chainID
// V = V - 2*chainID - 8

func GetAddresses(V *big.Int, R *big.Int, S *big.Int, data common.Hash) (common.Address, sdk.AccAddress, cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(sdk.AccAddress), *new(cryptotypes.PubKey), nil
}

func GetAddressesFromPubkeyBytes(pubkey []byte) (common.Address, sdk.AccAddress, cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(sdk.AccAddress), *new(cryptotypes.PubKey), nil
}

// first half of go-ethereum/core/types/transaction_signing.go:recoverPlain
func RecoverPubkey(sighash common.Hash, R, S, Vb *big.Int, homestead bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encode the signature in uncompressed format

// recover the public key from the signature

// second half of go-ethereum/core/types/transaction_signing.go:recoverPlain
func PubkeyToEVMAddress(pub []byte) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func PubkeyBytesToSeiPubKey(pub []byte) secp256k1.PubKey {
	_ = "STUB: not implemented"
	return *new(secp256k1.PubKey)
}

// RecoverAddressesFromTx recovers the sender's EVM address, Sei address, and public key
// from a signed PROTECTED Ethereum transaction using the provided signer.
// This is the core recovery function used by both the EVM ante handler and Giga executor.
//
// IMPORTANT: This function calls AdjustV internally, which is only correct for protected
// (EIP-155) transactions. For unprotected legacy transactions (blocktest only), use
// GetAddresses directly with the raw V value.
//
// The caller must provide the appropriate signer for the context. Use evmante.SignerMap[version](chainID)
// where version is determined by evmante.GetVersion(ctx, ethCfg) to ensure consistent behavior
// with the EVM ante handler.
func RecoverAddressesFromTx(ethTx *ethtypes.Transaction, signer ethtypes.Signer, chainID *big.Int) (common.Address, sdk.AccAddress, cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), *new(sdk.AccAddress), *new(cryptotypes.PubKey), nil
}
