package v600

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keys/secp256k1"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

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
