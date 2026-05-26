package types

import (
	"context"
	"sync"
)

// DefaultKeyringServiceName defines a default service name for the keyring.
const DefaultKeyringServiceName = "cosmos"

// Config is the structure that holds the SDK configuration parameters.
// This could be used to initialize certain configuration parameters for the SDK.
type Config struct {
	fullFundraiserPath  string
	bech32AddressPrefix map[string]string
	txEncoder           TxEncoder
	addressVerifier     func([]byte) error
	mtx                 sync.RWMutex

	// SLIP-44 related
	purpose  uint32
	coinType uint32

	sealed   bool
	sealedch chan struct{}
}

// cosmos-sdk wide global singleton
var (
	sdkConfig  *Config
	initConfig sync.Once
)

// New returns a new Config with default values.
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config { _ = "STUB: not implemented"; return nil }

// GetSealedConfig returns the config instance for the SDK if/once it is sealed.
func GetSealedConfig(ctx context.Context) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assertNotSealed panics if the config is sealed. Caller must hold the write lock.
func (config *Config) assertNotSealed() { _ = "STUB: not implemented"; return }

// SetBech32PrefixForAccount builds the Config with Bech32 addressPrefix and publKeyPrefix for accounts
// and returns the config instance
func (config *Config) SetBech32PrefixForAccount(addressPrefix, pubKeyPrefix string) {
	_ = "STUB: not implemented"
	return
}

// SetBech32PrefixForValidator builds the Config with Bech32 addressPrefix and publKeyPrefix for validators
//
//	and returns the config instance
func (config *Config) SetBech32PrefixForValidator(addressPrefix, pubKeyPrefix string) {
	_ = "STUB: not implemented"
	return
}

// SetBech32PrefixForConsensusNode builds the Config with Bech32 addressPrefix and publKeyPrefix for consensus nodes
// and returns the config instance
func (config *Config) SetBech32PrefixForConsensusNode(addressPrefix, pubKeyPrefix string) {
	_ = "STUB: not implemented"
	return
}

// SetTxEncoder builds the Config with TxEncoder used to marshal StdTx to bytes
func (config *Config) SetTxEncoder(encoder TxEncoder) { _ = "STUB: not implemented"; return }

// SetAddressVerifier builds the Config with the provided function for verifying that addresses
// have the correct format
func (config *Config) SetAddressVerifier(addressVerifier func([]byte) error) {
	_ = "STUB: not implemented"
	return
}

// Set the FullFundraiserPath (BIP44Prefix) on the config.
//
// Deprecated: This method is supported for backward compatibility only and will be removed in a future release. Use SetPurpose and SetCoinType instead.
func (config *Config) SetFullFundraiserPath(fullFundraiserPath string) {
	_ = "STUB: not implemented"
	return
}

// Set the BIP-0044 Purpose code on the config
func (config *Config) SetPurpose(purpose uint32) { _ = "STUB: not implemented"; return }

// Set the BIP-0044 CoinType code on the config
func (config *Config) SetCoinType(coinType uint32) { _ = "STUB: not implemented"; return }

// Seal seals the config such that the config state could not be modified further
func (config *Config) Seal() *Config { _ = "STUB: not implemented"; return nil }

// signal sealed after state exposed/unlocked

// GetBech32AccountAddrPrefix returns the Bech32 prefix for account address
func (config *Config) GetBech32AccountAddrPrefix() string { _ = "STUB: not implemented"; return "" }

// GetBech32ValidatorAddrPrefix returns the Bech32 prefix for validator address
func (config *Config) GetBech32ValidatorAddrPrefix() string { _ = "STUB: not implemented"; return "" }

// GetBech32ConsensusAddrPrefix returns the Bech32 prefix for consensus node address
func (config *Config) GetBech32ConsensusAddrPrefix() string { _ = "STUB: not implemented"; return "" }

// GetBech32AccountPubPrefix returns the Bech32 prefix for account public key
func (config *Config) GetBech32AccountPubPrefix() string { _ = "STUB: not implemented"; return "" }

// GetBech32ValidatorPubPrefix returns the Bech32 prefix for validator public key
func (config *Config) GetBech32ValidatorPubPrefix() string { _ = "STUB: not implemented"; return "" }

// GetBech32ConsensusPubPrefix returns the Bech32 prefix for consensus node public key
func (config *Config) GetBech32ConsensusPubPrefix() string { _ = "STUB: not implemented"; return "" }

// GetTxEncoder return function to encode transactions
func (config *Config) GetTxEncoder() TxEncoder { _ = "STUB: not implemented"; return *new(TxEncoder) }

// GetAddressVerifier returns the function to verify that addresses have the correct format
func (config *Config) GetAddressVerifier() func([]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPurpose returns the BIP-0044 Purpose code on the config.
func (config *Config) GetPurpose() uint32 { _ = "STUB: not implemented"; return 0 }

// GetCoinType returns the BIP-0044 CoinType code on the config.
func (config *Config) GetCoinType() uint32 { _ = "STUB: not implemented"; return 0 }

// GetFullFundraiserPath returns the BIP44Prefix.
//
// Deprecated: This method is supported for backward compatibility only and will be removed in a future release. Use GetFullBIP44Path instead.
func (config *Config) GetFullFundraiserPath() string { _ = "STUB: not implemented"; return "" }

// GetFullBIP44Path returns the BIP44Prefix.
func (config *Config) GetFullBIP44Path() string { _ = "STUB: not implemented"; return "" }

func KeyringServiceName() string { _ = "STUB: not implemented"; return "" }
