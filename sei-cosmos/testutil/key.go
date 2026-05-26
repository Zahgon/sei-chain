package testutil

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// GenerateCoinKey generates a new key mnemonic along with its addrress.
func GenerateCoinKey(algo keyring.SignatureAlgo) (sdk.AccAddress, string, error) {
	_ = "STUB: not implemented"
	// generate a private key, with mnemonic
	return *new(sdk.AccAddress), "", nil
}

// GenerateSaveCoinKey generates a new key mnemonic with its address.
// If mnemonic is provided then it's used for key generation.
// The key is saved in the keyring. The function returns error if overwrite=true and the key
// already exists.
func GenerateSaveCoinKey(
	keybase keyring.Keyring,
	keyName, mnemonic string,
	overwrite bool,
	algo keyring.SignatureAlgo,
) (sdk.AccAddress, string, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), "", nil
}

// ensure no overwrite
