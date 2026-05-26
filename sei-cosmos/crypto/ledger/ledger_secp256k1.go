package ledger

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/hd"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

var (
	// discoverLedger defines a function to be invoked at runtime for discovering
	// a connected Ledger device.
	discoverLedger discoverLedgerFn
)

type (
	// discoverLedgerFn defines a Ledger discovery function that returns a
	// connected device or an error upon failure. Its allows a method to avoid CGO
	// dependencies when Ledger support is potentially not enabled.
	discoverLedgerFn func() (SECP256K1, error)

	// SECP256K1 reflects an interface a Ledger API must implement for SECP256K1
	SECP256K1 interface {
		Close() error
		// Returns an uncompressed pubkey
		GetPublicKeySECP256K1([]uint32) ([]byte, error)
		// Returns a compressed pubkey and bech32 address (requires user confirmation)
		GetAddressPubKeySECP256K1([]uint32, string) ([]byte, string, error)
		// Signs a message (requires user confirmation)
		// signMode: 0 = SIGN_MODE_LEGACY_AMINO, 1 = SIGN_MODE_TEXTUAL
		SignSECP256K1([]uint32, []byte, byte) ([]byte, error)
	}

	// PrivKeyLedgerSecp256k1 implements PrivKey, calling the ledger nano we
	// cache the PubKey from the first call to use it later.
	PrivKeyLedgerSecp256k1 struct {
		// CachedPubKey should be private, but we want to encode it via
		// go-amino so we can view the address later, even without having the
		// ledger attached.
		CachedPubKey types.PubKey
		Path         hd.BIP44Params
	}
)

// NewPrivKeySecp256k1Unsafe will generate a new key and store the public key for later use.
//
// This function is marked as unsafe as it will retrieve a pubkey without user verification.
// It can only be used to verify a pubkey but never to create new accounts/keys. In that case,
// please refer to NewPrivKeySecp256k1
func NewPrivKeySecp256k1Unsafe(path hd.BIP44Params) (types.LedgerPrivKey, error) {
	_ = "STUB: not implemented"
	return *new(types.LedgerPrivKey), nil
}

// NewPrivKeySecp256k1 will generate a new key and store the public key for later use.
// The request will require user confirmation and will show account and index in the device
func NewPrivKeySecp256k1(path hd.BIP44Params, hrp string) (types.LedgerPrivKey, string, error) {
	_ = "STUB: not implemented"
	return *new(types.LedgerPrivKey), "", nil
}

// PubKey returns the cached public key.
func (pkl PrivKeyLedgerSecp256k1) PubKey() types.PubKey {
	_ = "STUB: not implemented"
	return *

	// Sign returns a secp256k1 signature for the corresponding message
	new(types.PubKey)
}

func (pkl PrivKeyLedgerSecp256k1) Sign(message []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignWithPath signs a message using a single device connection.
func SignWithPath(path hd.BIP44Params, message []byte) ([]byte, types.PubKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.PubKey), nil
}

// ShowAddress triggers a ledger device to show the corresponding address.
func ShowAddress(path hd.BIP44Params, expectedPubKey types.PubKey,
	accountAddressPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateKey allows us to verify the sanity of a public key after loading it
// from disk.
func (pkl PrivKeyLedgerSecp256k1) ValidateKey() error { _ = "STUB: not implemented"; return nil }

// AssertIsPrivKeyInner implements the PrivKey interface. It performs a no-op.
func (pkl *PrivKeyLedgerSecp256k1) AssertIsPrivKeyInner() {
	_ = "STUB: not implemented"

	// Bytes implements the PrivKey interface. It stores the cached public key so
	// we can verify the same key when we reconnect to a ledger.
	return
}

func (pkl PrivKeyLedgerSecp256k1) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Equals implements the PrivKey interface. It makes sure two private keys
// refer to the same public key.
func (pkl PrivKeyLedgerSecp256k1) Equals(other types.LedgerPrivKey) bool {
	_ = "STUB: not implemented"
	return false
}

func (pkl PrivKeyLedgerSecp256k1) Type() string { _ = "STUB: not implemented"; return "" }

func warnIfErrors(f func() error) {
	_ = "STUB: not implemented"

	// convertDERtoBER converts a DER-encoded signature to 64-byte R||S format.
	return
}

func convertDERtoBER(sig []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func getDevice() (SECP256K1, error) { _ = "STUB: not implemented"; return *new(SECP256K1), nil }

func validateKey(device SECP256K1, pkl PrivKeyLedgerSecp256k1) error {
	_ = "STUB: not implemented"
	return nil
}

func sign(device SECP256K1, pkl PrivKeyLedgerSecp256k1, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getPubKeyUnsafe retrieves pubkey without user verification (use getPubKeyAddrSafe for new keys).
func getPubKeyUnsafe(device SECP256K1, path hd.BIP44Params) (types.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(types.PubKey), nil
}

// getPubKeyAddrSafe retrieves pubkey with user confirmation (for creating new keys).
func getPubKeyAddrSafe(device SECP256K1, path hd.BIP44Params, hrp string) (types.PubKey, string, error) {
	_ = "STUB: not implemented"
	return *new(types.PubKey), "", nil
}
