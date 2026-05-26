package cli

import (
	"github.com/spf13/cobra"
)

const (
	HomeFlag   = "home"
	TraceFlag  = "trace"
	OutputFlag = "output" // used in the cli
)

type Executable interface {
	Execute() error
}

// PrepareBaseCmd is meant for tendermint and other servers
func PrepareBaseCmd(cmd *cobra.Command, envPrefix, defaultHome string) *cobra.Command {
	_ = "STUB: not implemented"
	// the primary caller of this command is in the SDK and
	// returning the cobra.Command object avoids breaking that
	// code. In the long term, the SDK could avoid this entirely.
	return nil
}

// InitEnv sets to use ENV variables if set.
func InitEnv(prefix string) {
	_ = "STUB: not implemented"
	// This copies all variables like TMROOT to TM_ROOT,
	// so we can support both formats for the user
	return
}

// env variables with TM prefix (eg. TM_ROOT)

type cobraCmdFunc func(cmd *cobra.Command, args []string) error

// Returns a single function that calls each argument function in sequence
// RunE, PreRunE, PersistentPreRunE, etc. all have this same signature
func concatCobraCmdFuncs(fs ...cobraCmdFunc) cobraCmdFunc {
	_ = "STUB: not implemented"
	return *new(cobraCmdFunc)
}

// Bind all flags and read the config into viper
func BindFlagsLoadViper(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// cmd.Flags() includes flags from this command and all persistent flags from the parent
	return nil
}

// name of config file (without extension)
// search root directory
// search root directory /config

// If a config file is found, read it in.

// stderr, so if we redirect output to json file, this doesn't appear
// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

// ignore not found error, return other errors

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

type ExitCoder interface {
	ExitCode() int
}

// execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (e Executor) Execute() error { _ = "STUB: not implemented"; return nil }

// return error code 1 by default, can override it with a special error type
