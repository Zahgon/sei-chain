package keys

import (
	"io"

	cryptokeyring "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
)

// available output formats.
const (
	OutputFormatText = "text"
	OutputFormatJSON = "json"
)

type bechKeyOutFn func(keyInfo cryptokeyring.Info) (cryptokeyring.KeyOutput, error)

func printKeyInfo(w io.Writer, keyInfo cryptokeyring.Info, bechKeyOut bechKeyOutFn, output string) {
	_ = "STUB: not implemented"
	return
}

func printInfos(w io.Writer, infos []cryptokeyring.Info, output string) {
	_ = "STUB: not implemented"
	return
}

func printTextInfos(w io.Writer, kos []cryptokeyring.KeyOutput) { _ = "STUB: not implemented"; return }
