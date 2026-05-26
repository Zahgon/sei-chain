package client

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// ClientContextKey defines the context key used to retrieve a client.Context from
// a command's Context.
const ClientContextKey = sdk.ContextKey("client.context")

// SetCmdClientContextHandler is to be used in a command pre-hook execution to
// read flags that populate a Context and sets that to the command's Context.
func SetCmdClientContextHandler(clientCtx Context, cmd *cobra.Command) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ValidateCmd returns unknown command error or Help display if help flag set
func ValidateCmd(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// search for help flag

// check if the current arg is a flag

// the next arg should be skipped if the current arg is a
// flag and does not use "=" to assign the flag's value

// skip current arg

// unknown command found
// continue searching for help flag

// return the help screen if no unknown command is found

// build suggestions for unknown argument

// ReadPersistentCommandFlags returns a Context with fields set for "persistent"
// or common flags that do not necessarily change with context.
//
// Note, the provided clientCtx may have field pre-populated. The following order
// of precedence occurs:
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func ReadPersistentCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	_ = "STUB: not implemented"
	return *new(Context), nil
}

// The keyring directory is optional and falls back to the home directory
// if omitted.

// readQueryCommandFlags returns an updated Context with fields set based on flags
// defined in AddQueryFlagsToCmd. An error is returned if any flag query fails.
//
// Note, the provided clientCtx may have field pre-populated. The following order
// of precedence occurs:
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func readQueryCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	_ = "STUB: not implemented"
	return *new(Context), nil
}

// readTxCommandFlags returns an updated Context with fields set based on flags
// defined in AddTxFlagsToCmd. An error is returned if any flag query fails.
//
// Note, the provided clientCtx may have field pre-populated. The following order
// of precedence occurs:
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func readTxCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	_ = "STUB: not implemented"
	return *new(Context), nil
}

// If the `from` signer account is a ledger key, we need to use
// SIGN_MODE_AMINO_JSON, because ledger doesn't support proto yet.
// ref: https://github.com/cosmos/cosmos-sdk/issues/8109

// GetClientQueryContext returns a Context from a command with fields set based on flags
// defined in AddQueryFlagsToCmd. An error is returned if any flag query fails.
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func GetClientQueryContext(cmd *cobra.Command) (Context, error) {
	_ = "STUB: not implemented"
	return *new(Context), nil
}

// GetClientTxContext returns a Context from a command with fields set based on flags
// defined in AddTxFlagsToCmd. An error is returned if any flag query fails.
//
// - client.Context field not pre-populated & flag not set: uses default flag value
// - client.Context field not pre-populated & flag set: uses set flag value
// - client.Context field pre-populated & flag not set: uses pre-populated value
// - client.Context field pre-populated & flag set: uses set flag value
func GetClientTxContext(cmd *cobra.Command) (Context, error) {
	_ = "STUB: not implemented"
	return *new(Context), nil
}

// GetClientContextFromCmd returns a Context from a command or an empty Context
// if it has not been set.
func GetClientContextFromCmd(cmd *cobra.Command) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// SetCmdClientContext sets a command's Context value to the provided argument.
func SetCmdClientContext(cmd *cobra.Command, clientCtx Context) error {
	_ = "STUB: not implemented"
	return nil
}
