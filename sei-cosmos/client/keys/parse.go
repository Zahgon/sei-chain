package keys

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func bech32Prefixes(config *sdk.Config) []string { _ = "STUB: not implemented"; return nil }

type hexOutput struct {
	Human string `json:"human"`
	Bytes string `json:"bytes"`
}

func (ho hexOutput) String() string { _ = "STUB: not implemented"; return "" }

func newHexOutput(human string, bs []byte) hexOutput {
	_ = "STUB: not implemented"
	return *new(hexOutput)
}

type bech32Output struct {
	Formats []string `json:"formats"`
}

func newBech32Output(config *sdk.Config, bs []byte) bech32Output {
	_ = "STUB: not implemented"
	return *new(bech32Output)
}

func (bo bech32Output) String() string { _ = "STUB: not implemented"; return "" }

// ParseKeyStringCommand parses an address from hex to bech32 and vice versa.
func ParseKeyStringCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func parseKey(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func doParseKey(cmd *cobra.Command, config *sdk.Config, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// print info from bech32
func runFromBech32(w io.Writer, bech32str, output string) bool {
	_ = "STUB: not implemented"
	return false
}

// print info from hex
func runFromHex(config *sdk.Config, w io.Writer, hexstr, output string) bool {
	_ = "STUB: not implemented"
	return false
}

func displayParseKeyInfo(w io.Writer, stringer fmt.Stringer, output string) {
	_ = "STUB: not implemented"
	return
}
