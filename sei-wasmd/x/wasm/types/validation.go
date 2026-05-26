package types

var (
	// MaxLabelSize is the longest label that can be used when Instantiating a contract
	MaxLabelSize = 128 // extension point for chains to customize via compile flag.

	// MaxWasmSize is the largest a compiled contract code can be when storing code on chain
	MaxWasmSize = 800 * 1024 // extension point for chains to customize via compile flag.
)

func validateWasmCode(s []byte) error { _ = "STUB: not implemented"; return nil }

func validateLabel(label string) error { _ = "STUB: not implemented"; return nil }
