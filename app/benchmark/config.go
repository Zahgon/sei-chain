package benchmark

import (
	"github.com/sei-protocol/sei-load/config"
)

// LoadConfig reads a sei-load config file or returns a default config.
// The chainID and seiChainID are always overridden with actual values from the running chain.
func LoadConfig(configPath string, evmChainID int64, seiChainID string) (*config.LoadConfig, error) {
	_ = "STUB: not implemented"
	return nil,

		// Return default config (EVMTransfer scenario)
		nil
}

// We handle deployment in-process

//nolint:gosec // G304: configPath is from trusted env var

// Override chain IDs with actual values from the running chain

// Always use mock deploy since we handle deployment in-process
