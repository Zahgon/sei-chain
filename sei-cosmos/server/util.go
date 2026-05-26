package server

import (
	"context"
	"errors"
	"io"
	"net"

	tmcfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	"github.com/sei-protocol/seilog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// DONTCOVER

// ServerContextKey defines the context key used to retrieve a server.Context from
// a command's Context.
const ServerContextKey = sdk.ContextKey("server.context")

var (
	ErrShouldRestart = errors.New("node should be restarted")
	logger           = seilog.NewLogger("cosmos", "server")
)

// server context
type Context struct {
	Viper  *viper.Viper
	Config *tmcfg.Config
}

// ErrorCode contains the exit code for server exit.
type ErrorCode struct {
	Code int
}

func (e ErrorCode) Error() string { _ = "STUB: not implemented"; return "" }

func NewDefaultContext() *Context { _ = "STUB: not implemented"; return nil }

func NewContext(v *viper.Viper, config *tmcfg.Config) *Context {
	_ = "STUB: not implemented"
	return nil
}

func bindFlags(basename string, cmd *cobra.Command, v *viper.Viper) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//handled by named return value.

// Environment variables can't have dashes in them, so bind them to their equivalent
// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR

// Apply the viper config value to the flag when the flag is not set and viper has a value

func InterceptConfigs(cmd *cobra.Command) (*tmcfg.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the executable name and configure the viper instance so that environmental
// variables are checked based off that name. The underscore character is used
// as a separator

// Configure the viper instance

// intercept configuration files, using both Viper instances separately

// InterceptConfigsPreRunHandler performs a pre-run function for the root daemon
// application command. It will create a Viper literal and a default server
// Context. The server Tendermint configuration will either be read and parsed
// or created and saved to disk, where the server Context is updated to reflect
// the Tendermint configuration. It takes custom app config template and config
// settings to create a custom Tendermint configuration. If the custom template
// is empty, it uses default-template provided by the server. The Viper literal
// is used to read and parse the application configuration. Command handlers can
// fetch the server Context to get the Tendermint configuration or to get access
// to Viper.
func InterceptConfigsPreRunHandler(cmd *cobra.Command, customAppConfigTemplate string, customAppConfig any) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the executable name and configure the viper instance so that environmental
// variables are checked based off that name. The underscore character is used
// as a separator

// Configure the viper instance

// intercept configuration files, using both Viper instances separately

// return value is a tendermint configuration object

// For logging efficiency, seilog takes the log format at the time of
// initialization and intentionally does not provide the ability to dynamically
// alter log format at runtime. This means seilog can operate with
// zero-allocations during logging operation.

// The order of extrapolation priority for log level is:
// 1. explicit flag
// 2. env var
// 3. seid config
//
// Here, we check if flag is explicitly set first, if so override all log levels to it.
// If not, check if env var is set in which case it is already picked up by seilog; nothing to do.
// Otherwise, make sure to set the log level to what's configured in the config files.

// CLI flag set; take presence.

// No CLI flag and no env var; fall back to config

// Do nothing, since flag was not set but env var was non-empty, which gets
// handled by seilog init.

// GetServerContextFromCmd returns a Context from a command or an empty Context
// if it has not been set.
func GetServerContextFromCmd(cmd *cobra.Command) *Context { _ = "STUB: not implemented"; return nil }

// SetCmdServerContext sets a command's Context value to the provided argument.
func SetCmdServerContext(cmd *cobra.Command, serverCtx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// interceptConfigs parses and updates a Tendermint configuration file or
// creates a new one and saves it. It also parses and saves the application
// configuration file. The Tendermint configuration file is parsed given a root
// Viper object, whereas the application is parsed with the private package-aware
// viperCfg object.
func interceptConfigs(rootViper *viper.Viper, customAppTemplate string, customConfig any) (*tmcfg.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read into the configuration whatever data the viper instance has for it.
// This may come from the configuration file above but also any of the other
// sources viper uses.

// add server commands
func AddCommands(
	rootCmd *cobra.Command,
	defaultNodeHome string,
	appCreator types.AppCreator,
	appExport types.AppExporter,
	addStartFlags types.ModuleInitFlags,
	tracerProviderOptions []trace.TracerProviderOption,
) {
	_ = "STUB: not implemented"
	return
}

// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
// TODO there must be a better way to get external IP
func ExternalIP() (string, error) { _ = "STUB: not implemented"; return "", nil }

// not an ipv4 address

// TrapSignal traps SIGINT and SIGTERM and terminates the server correctly.
func TrapSignal(cleanupFunc func()) { _ = "STUB: not implemented"; return }

// WaitForQuitSignals waits for SIGINT and SIGTERM and returns.
func WaitForQuitSignals(ctx context.Context, restartCh chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// blocks forever on a nil channel

func skipInterface(iface net.Interface) bool { _ = "STUB: not implemented"; return false }

// interface down

// loopback interface

func addrToIP(addr net.Addr) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func openTraceWriter(traceWriterFile string) (w io.Writer, err error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}
