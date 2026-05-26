package params

import (
	evmrpcconfig "github.com/sei-protocol/sei-chain/evmrpc/config"
	srvconfig "github.com/sei-protocol/sei-chain/sei-cosmos/server/config"
	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
)

const (
	HumanCoinUnit = "sei"
	BaseCoinUnit  = "usei"
	UseiExponent  = 6

	DefaultBondDenom = BaseCoinUnit

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address.
	Bech32PrefixAccAddr = "sei"
)

// UnsafeBypassCommitTimeoutOverride commits block as soon as we reach consensus instead of waiting
// for timeout, this may cause validators to not get their votes in time
var UnsafeBypassCommitTimeoutOverride = false

var (
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key.
	Bech32PrefixAccPub = Bech32PrefixAccAddr + "pub"
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address.
	Bech32PrefixValAddr = Bech32PrefixAccAddr + "valoper"
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key.
	Bech32PrefixValPub = Bech32PrefixAccAddr + "valoperpub"
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address.
	Bech32PrefixConsAddr = Bech32PrefixAccAddr + "valcons"
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key.
	Bech32PrefixConsPub = Bech32PrefixAccAddr + "valconspub"
)

func init() {
	SetAddressPrefixes()
	RegisterDenoms()
}

func RegisterDenoms() { _ = "STUB: not implemented"; return }

func SetAddressPrefixes() { _ = "STUB: not implemented"; return }

// This is copied from the cosmos sdk v0.43.0-beta1
// source: https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-beta1/types/address.go#L141

// TODO: Do we want to allow addresses of lengths other than 20 and 32 bytes?

// NodeMode represents the type of node being run
// Extends Tendermint's Mode with additional archive mode
type NodeMode string

const (
	// Reuse Tendermint's mode constants
	NodeModeValidator NodeMode = tmcfg.ModeValidator // "validator"
	NodeModeFull      NodeMode = tmcfg.ModeFull      // "full"
	NodeModeSeed      NodeMode = tmcfg.ModeSeed      // "seed"
	// Additional mode specific to Sei Chain
	NodeModeArchive NodeMode = "archive"
)

// IsFullnodeType returns true if the node is a fullnode-like node (full or archive)
func (m NodeMode) IsFullnodeType() bool { _ = "STUB: not implemented"; return false }

// setValidatorTypeTendermintConfig sets common Tendermint config for validator-like nodes
func setValidatorTypeTendermintConfig(config *tmcfg.Config) { _ = "STUB: not implemented"; return }

// Validators don't need tx indexing

// setFullnodeTypeTendermintConfig sets common Tendermint config for fullnode-like nodes
func setFullnodeTypeTendermintConfig(config *tmcfg.Config) { _ = "STUB: not implemented"; return }

// Full nodes need tx indexing for queries

// SetTendermintConfigByMode sets Tendermint config values based on node mode
// Note: config.Mode should be set by the caller before calling this function
// Archive nodes should have config.Mode = "full" since Tendermint doesn't recognize "archive"
func SetTendermintConfigByMode(config *tmcfg.Config) { _ = "STUB: not implemented"; return }

// Seed nodes need more connections to serve peers

// Archive nodes use full node Tendermint config
// The difference is in app config (keeping all history)

// setValidatorTypeAppConfig sets common app config for validator-like nodes
func setValidatorTypeAppConfig(config *srvconfig.Config) {
	_ = "STUB: not implemented"
	// Services: validators should minimize exposed services for security
	return
}

// setFullnodeTypeAppConfig sets common app config for fullnode-like nodes
func setFullnodeTypeAppConfig(config *srvconfig.Config) {
	_ = "STUB: not implemented"
	// Services: full nodes provide query services
	return
}

// StateStore: full nodes keep recent history for queries

// MinRetainBlocks: prune Tendermint blocks older than 100k blocks

// Pruning uses defaults (nothing,0,0 - keep all state history)

// setArchiveTypeAppConfig configures archive node settings
// Archive nodes are like full nodes but keep all history
func setArchiveTypeAppConfig(config *srvconfig.Config) {
	_ = "STUB: not implemented"
	// Start with full node configuration
	return
}

// Archive nodes keep all history
// 0 = keep all state history
// 0 = keep all Tendermint blocks

// Pruning uses defaults (nothing,0,0 - keep all state history)

// SetAppConfigByMode sets app config values based on node mode
func SetAppConfigByMode(config *srvconfig.Config, mode NodeMode) { _ = "STUB: not implemented"; return }

// Validator and Seed nodes share the same app config

// Default to full node settings

// SetEVMConfigByMode sets EVM config based on node mode
// Validators and seeds have EVM disabled, full nodes and archives have it enabled
func SetEVMConfigByMode(config *evmrpcconfig.Config, mode NodeMode) {
	_ = "STUB: not implemented"
	return
}
