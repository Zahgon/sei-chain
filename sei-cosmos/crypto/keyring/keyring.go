package keyring

import (
	"io"

	"github.com/99designs/keyring"

	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/hd"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Backend options for Keyring
const (
	BackendFile    = "file"
	BackendOS      = "os"
	BackendKWallet = "kwallet"
	BackendPass    = "pass"
	BackendTest    = "test"
	BackendMemory  = "memory"
)

const (
	keyringFileDirName = "keyring-file"
	keyringTestDirName = "keyring-test"
	passKeyringPrefix  = "keyring-%s"
)

var (
	_                          Keyring = &keystore{}
	maxPassphraseEntryAttempts         = 3
)

// Keyring exposes operations over a backend supported by github.com/99designs/keyring.
type Keyring interface {
	// List all keys.
	List() ([]Info, error)

	// Supported signing algorithms for Keyring and Ledger respectively.
	SupportedAlgorithms() (SigningAlgoList, SigningAlgoList)

	// Key and KeyByAddress return keys by uid and address respectively.
	Key(uid string) (Info, error)
	KeyByAddress(address sdk.Address) (Info, error)

	// Delete and DeleteByAddress remove keys from the keyring.
	Delete(uid string) error
	DeleteByAddress(address sdk.Address) error

	// NewMnemonic generates a new mnemonic, derives a hierarchical deterministic key from it, and
	// persists the key to storage. Returns the generated mnemonic and the key Info.
	// It returns an error if it fails to generate a key for the given algo type, or if
	// another key is already stored under the same name or address.
	//
	// A passphrase set to the empty string will set the passphrase to the DefaultBIP39Passphrase value.
	NewMnemonic(uid string, language Language, hdPath, bip39Passphrase string, algo SignatureAlgo) (Info, string, error)

	// NewAccount converts a mnemonic to a private key and BIP-39 HD Path and persists it.
	// It fails if there is an existing key Info with the same address.
	NewAccount(uid, mnemonic, bip39Passphrase, hdPath string, algo SignatureAlgo) (Info, error)

	// SaveLedgerKey retrieves a public key reference from a Ledger device and persists it.
	SaveLedgerKey(uid string, algo SignatureAlgo, hrp string, coinType, account, index uint32) (Info, error)

	// SavePubKey stores a public key and returns the persisted Info structure.
	SavePubKey(uid string, pubkey types.PubKey, algo hd.PubKeyType) (Info, error)

	// SaveMultisig stores and returns a new multsig (offline) key reference.
	SaveMultisig(uid string, pubkey types.PubKey) (Info, error)

	Signer

	Importer
	Exporter
}

// UnsafeKeyring exposes unsafe operations such as unsafe unarmored export in
// addition to those that are made available by the Keyring interface.
type UnsafeKeyring interface {
	Keyring
	UnsafeExporter
}

// Signer is implemented by key stores that want to provide signing capabilities.
type Signer interface {
	// Sign sign byte messages with a user key.
	Sign(uid string, msg []byte) ([]byte, types.PubKey, error)

	// SignByAddress sign byte messages with a user key providing the address.
	SignByAddress(address sdk.Address, msg []byte) ([]byte, types.PubKey, error)
}

// Importer is implemented by key stores that support import of public and private keys.
type Importer interface {
	// ImportPrivKey imports ASCII armored passphrase-encrypted private keys.
	ImportPrivKey(uid, armor, passphrase string) error
}

// Exporter is implemented by key stores that support export of public and private keys.
type Exporter interface {

	// ExportPrivKeyArmor returns a private key in ASCII armored format.
	// It returns an error if the key does not exist or a wrong encryption passphrase is supplied.
	ExportPrivKeyArmor(uid, encryptPassphrase string) (armor string, err error)
	ExportPrivKeyArmorByAddress(address sdk.Address, encryptPassphrase string) (armor string, err error)
}

// UnsafeExporter is implemented by key stores that support unsafe export
// of private keys' material.
type UnsafeExporter interface {
	// UnsafeExportPrivKeyHex returns a private key in unarmored hex format
	UnsafeExportPrivKeyHex(uid string) (string, error)
}

// Option overrides keyring configuration options.
type Option func(options *Options)

// Options define the options of the Keyring.
type Options struct {
	// supported signing algorithms for keyring
	SupportedAlgos SigningAlgoList
	// supported signing algorithms for Ledger
	SupportedAlgosLedger SigningAlgoList
}

// NewInMemory creates a transient keyring useful for testing
// purposes and on-the-fly key generation.
// Keybase options can be applied when generating this new Keybase.
func NewInMemory(opts ...Option) Keyring { _ = "STUB: not implemented"; return *new(Keyring) }

// NewInMemoryWithKeyring returns an in memory keyring using the specified keyring.Keyring
// as the backing keyring.
func NewInMemoryWithKeyring(kr keyring.Keyring, opts ...Option) Keyring {
	_ = "STUB: not implemented"
	return *new(Keyring)
}

// New creates a new instance of a keyring.
// Keyring ptions can be applied when generating the new instance.
// Available backends are "os", "file", "kwallet", "memory", "pass", "test".
func New(
	appName, backend, rootDir string, userInput io.Reader, opts ...Option,
) (Keyring, error) {
	_ = "STUB: not implemented"
	return *new(Keyring), nil
}

type keystore struct {
	db      keyring.Keyring
	options Options
}

func newKeystore(kr keyring.Keyring, opts ...Option) keystore {
	_ = "STUB: not implemented"
	// Default options for keybase
	return *new(keystore)
}

func (ks keystore) ExportPrivKeyArmor(uid, encryptPassphrase string) (armor string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExportPrivateKeyObject exports an armored private key object.
func (ks keystore) ExportPrivateKeyObject(uid string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ks keystore) ExportPrivKeyArmorByAddress(address sdk.Address, encryptPassphrase string) (armor string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ks keystore) ImportPrivKey(uid, armor, passphrase string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ks keystore) Sign(uid string, msg []byte) ([]byte, types.PubKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.PubKey), nil
}

func (ks keystore) SignByAddress(address sdk.Address, msg []byte) ([]byte, types.PubKey, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.PubKey), nil
}

func (ks keystore) SaveLedgerKey(uid string, algo SignatureAlgo, hrp string, coinType, account, index uint32) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func (ks keystore) writeLedgerKey(name string, pub types.PubKey, path hd.BIP44Params, algo hd.PubKeyType) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func (ks keystore) SaveMultisig(uid string, pubkey types.PubKey) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func (ks keystore) SavePubKey(uid string, pubkey types.PubKey, algo hd.PubKeyType) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func (ks keystore) DeleteByAddress(address sdk.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (ks keystore) Delete(uid string) error { _ = "STUB: not implemented"; return nil }

func (ks keystore) KeyByAddress(address sdk.Address) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func wrapKeyNotFound(err error, msg string) error { _ = "STUB: not implemented"; return nil }

func (ks keystore) List() ([]Info, error) { _ = "STUB: not implemented"; return nil, nil }

// add the name of the key in case the user wants to retrieve it
// afterwards

// add the name of the key in case the user wants to retrieve it
// afterwards

// add the name of the key in case the user wants to retrieve it
// afterwards

func (ks keystore) NewMnemonic(uid string, language Language, hdPath, bip39Passphrase string, algo SignatureAlgo) (Info, string, error) {
	_ = "STUB: not implemented"
	return *new(Info), "", nil
}

// Default number of words (24): This generates a mnemonic directly from the
// number of words by reading system entropy.

func (ks keystore) NewAccount(name string, mnemonic string, bip39Passphrase string, hdPath string, algo SignatureAlgo) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

// create master key and derive first key for keyring

// check if the a key already exists with the same address and return an error
// if found

func (ks keystore) isSupportedSigningAlgo(algo SignatureAlgo) bool {
	_ = "STUB: not implemented"
	return false
}

func (ks keystore) key(infoKey string) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func (ks keystore) Key(uid string) (Info, error) { _ = "STUB: not implemented"; return *new(Info), nil }

// SupportedAlgorithms returns the keystore Options' supported signing algorithm.
// for the keyring and Ledger.
func (ks keystore) SupportedAlgorithms() (SigningAlgoList, SigningAlgoList) {
	_ = "STUB: not implemented"
	return *new(SigningAlgoList), *new(SigningAlgoList)
}

// SignWithLedger signs a binary message with the ledger device referenced by an Info object
// and returns the signed bytes and the public key. It returns an error if the device could
// not be queried or it returned an error.
func SignWithLedger(info Info, msg []byte) (sig []byte, pub types.PubKey, err error) {
	_ = "STUB: not implemented"
	return nil, *new(types.PubKey), nil
}

// Use single-connection signing to avoid race conditions from multiple device open/close cycles

// Validate that the public key from the device matches the cached key in the keyring

func newOSBackendKeyringConfig(appName, dir string, buf io.Reader) keyring.Config {
	_ = "STUB: not implemented"
	return *new(keyring.Config)
}

func newTestBackendKeyringConfig(appName, dir string) keyring.Config {
	_ = "STUB: not implemented"
	return *new(keyring.Config)
}

func newKWalletBackendKeyringConfig(appName, _ string, _ io.Reader) keyring.Config {
	_ = "STUB: not implemented"
	return *new(keyring.Config)
}

func newPassBackendKeyringConfig(appName, _ string, _ io.Reader) keyring.Config {
	_ = "STUB: not implemented"
	return *new(keyring.Config)
}

func newFileBackendKeyringConfig(name, dir string, buf io.Reader) keyring.Config {
	_ = "STUB: not implemented"
	return *new(keyring.Config)
}

func newRealPrompt(dir string, buf io.Reader) func(string) (string, error) {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: LGTM.io reports a false positive alert that states we are printing the password,
// but we only log the error.
//
// lgtm [go/clear-text-logging]

// NOTE: LGTM.io reports a false positive alert that states we are printing the password,
// but we only log the error.
//
// lgtm [go/clear-text-logging]

func (ks keystore) writeLocalKey(name string, priv types.PrivKey, algo hd.PubKeyType) (Info, error) {
	_ = "STUB: not implemented"
	// encrypt private key using keyring
	return *new(Info), nil
}

func (ks keystore) writeInfo(info Info) error { _ = "STUB: not implemented"; return nil }

// existsInDb returns true if key is in DB. Error is returned only when we have error
// different thant ErrKeyNotFound
func (ks keystore) existsInDb(info Info) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// address lookup succeeds - info exists

// received unexpected error - returns error

// uid lookup succeeds - info exists

// received unexpected error - returns

// both lookups failed, info does not exist

func (ks keystore) writeOfflineKey(name string, pub types.PubKey, algo hd.PubKeyType) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

func (ks keystore) writeMultisigKey(name string, pub types.PubKey) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

type unsafeKeystore struct {
	keystore
}

// NewUnsafe returns a new keyring that provides support for unsafe operations.
func NewUnsafe(kr Keyring) UnsafeKeyring {
	_ = "STUB: not implemented"
	// The type assertion is against the only keystore
	// implementation that is currently provided.
	return *new(UnsafeKeyring)
}

// UnsafeExportPrivKeyHex exports private keys in unarmored hexadecimal format.
func (ks unsafeKeystore) UnsafeExportPrivKeyHex(uid string) (privkey string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func addrHexKeyAsString(address sdk.Address) string { _ = "STUB: not implemented"; return "" }
