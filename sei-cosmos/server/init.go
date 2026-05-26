package server

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Deprecated: GenerateCoinKey generates a new key mnemonic along with its addrress.
// Please use testutils.GenerateCoinKey instead.
func GenerateCoinKey(algo keyring.SignatureAlgo) (sdk.AccAddress, string, error) {
	_ = "STUB: not implemented"
	// generate a private key, with mnemonic
	return *new(sdk.AccAddress), "", nil
}

// Deprecated: GenerateSaveCoinKey generates a new key mnemonic with its addrress.
// If mnemonic is provided then it's used for key generation.
// The key is saved in the keyring. The function returns error if overwrite=true and the key
// already exists.
// Please use testutils.GenerateSaveCoinKey instead.
func GenerateSaveCoinKey(
	keybase keyring.Keyring,
	keyName string,
	overwrite bool,
	algo keyring.SignatureAlgo,
) (sdk.AccAddress, string, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), "", nil
}

// ensure no overwrite

// remove the old key by name if it exists
