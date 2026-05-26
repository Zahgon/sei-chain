package cli

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdLibVersion gets current libwasmvm version.
func GetCmdLibVersion() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdListCode lists all wasm code uploaded
func GetCmdListCode() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdListContractByCode lists all wasm code uploaded for given code id
func GetCmdListContractByCode() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryCode returns the bytecode for a given contract
func GetCmdQueryCode() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryCodeInfo returns the code info for a given code id
func GetCmdQueryCodeInfo() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdGetContractInfo gets details about a given contract
func GetCmdGetContractInfo() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdGetContractState dumps full internal state of a given contract
func GetCmdGetContractState() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetCmdGetContractStateAll() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetCmdGetContractStateRaw() *cobra.Command { _ = "STUB: not implemented"; return nil }

func GetCmdGetContractStateSmart() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdGetContractHistory prints the code history for a given contract
func GetCmdGetContractHistory() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdListPinnedCode lists all wasm code ids that are pinned
func GetCmdListPinnedCode() *cobra.Command { _ = "STUB: not implemented"; return nil }

type argumentDecoder struct {
	// dec is the default decoder
	dec                func(string) ([]byte, error)
	asciiF, hexF, b64F bool
}

func newArgDecoder(def func(string) ([]byte, error)) *argumentDecoder {
	_ = "STUB: not implemented"
	return nil
}

func (a *argumentDecoder) RegisterFlags(f *flag.FlagSet, argName string) {
	_ = "STUB: not implemented"
	return
}

func (a *argumentDecoder) DecodeString(s string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func asciiDecodeString(s string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// sdk ReadPageRequest expects binary but we encoded to base64 in our marshaller
		nil
}

func withPageKeyDecoded(flagSet *flag.FlagSet) *flag.FlagSet { _ = "STUB: not implemented"; return nil }
