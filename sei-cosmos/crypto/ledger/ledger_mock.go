//go:build ledger && test_ledger_mock
// +build ledger,test_ledger_mock

package ledger

// If ledger support (build tag) has been enabled, which implies a CGO dependency,
// set the discoverLedger function which is responsible for loading the Ledger
// device at runtime or returning an error.
func init() {
	discoverLedger = func() (SECP256K1, error) {
		return LedgerSECP256K1Mock{}, nil
	}
}

type LedgerSECP256K1Mock struct {
}

func (mock LedgerSECP256K1Mock) Close() error {
	_ = "STUB: not implemented"

	// GetPublicKeySECP256K1 mocks a ledger device
	// as per the original API, it returns an uncompressed key
	return nil
}

func (mock LedgerSECP256K1Mock) GetPublicKeySECP256K1(derivationPath []uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use btcec v2 API

// GetAddressPubKeySECP256K1 mocks a ledger device
// as per the original API, it returns a compressed key and a bech32 address
func (mock LedgerSECP256K1Mock) GetAddressPubKeySECP256K1(derivationPath []uint32, hrp string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// re-serialize in the 33-byte compressed format

// Generate the bech32 addr using existing tmcrypto/etc.

func (mock LedgerSECP256K1Mock) SignSECP256K1(derivationPath []uint32, message []byte, _ byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use btcec v2 API

// Use single SHA256 to match the secp256k1.VerifySignature expectations

// true=compressed pubkey

// Return 64-byte R||S format (remove recovery id)

// ShowAddressSECP256K1 shows the address for the corresponding bip32 derivation path
func (mock LedgerSECP256K1Mock) ShowAddressSECP256K1(bip32Path []uint32, hrp string) error {
	_ = "STUB: not implemented"
	return nil
}
