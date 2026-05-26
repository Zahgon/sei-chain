package admin

import (
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
)

const (
	DefaultEnabled = false
	DefaultAddress = "127.0.0.1:9095"
)

// Config defines configuration for the admin gRPC server.
type Config struct {
	// Enabled controls whether the admin gRPC server starts.
	Enabled bool `mapstructure:"admin_enabled"`
	// Address is the listen address. Must be a loopback address.
	Address string `mapstructure:"admin_address"`
}

var DefaultConfig = Config{
	Enabled: DefaultEnabled,
	Address: DefaultAddress,
}

// ReadConfig reads admin config from app options (Viper-backed).
func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

// validateLoopback ensures the address is bound to a loopback interface.
func validateLoopback(address string) error { _ = "STUB: not implemented"; return nil }

// ConfigTemplate is the TOML template for the [admin] section of app.toml.
const ConfigTemplate = `
###############################################################################
###                       Admin Configuration (Auto-managed)                ###
###############################################################################
 
[admin_server]
 
# Enable the admin gRPC server for runtime log level control.
admin_enabled = {{ .Admin.Enabled }}
 
# Listen address for the admin gRPC server. Must be a loopback address.
admin_address = "{{ .Admin.Address }}"
`
