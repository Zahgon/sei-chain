package config

import (
	"github.com/spf13/viper"
)

const defaultConfigTemplate = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

###############################################################################
###                           Client Configuration                            ###
###############################################################################

# The network chain ID
chain-id = "{{ .ChainID }}"
# The keyring's backend, where the keys are stored (os|file|kwallet|pass|test|memory)
keyring-backend = "{{ .KeyringBackend }}"
# CLI output format (text|json)
output = "{{ .Output }}"
# <host>:<port> to Tendermint RPC interface for this chain
node = "{{ .Node }}"
# Transaction broadcasting mode (sync|async|block)
broadcast-mode = "{{ .BroadcastMode }}"
`

func SetClientConfig(key string, value string, configPath string, config *ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// writeConfigToFile parses defaultConfigTemplate, renders config using the template and writes it to
// configFilePath.
func writeConfigToFile(configFilePath string, config *ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureConfigPath creates a directory configPath if it does not exist
func ensureConfigPath(configPath string) error { _ = "STUB: not implemented"; return nil }

// getClientConfig reads values from client.toml file and unmarshalls them into ClientConfig
func GetClientConfig(configPath string, v *viper.Viper) (*ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
