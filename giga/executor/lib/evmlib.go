package lib

import (
	"github.com/ethereum/evmc/v12/bindings/go/evmc"
)

//go:generate go run ./gen/main.go .

// InitEvmoneVM initializes the EVMC VM by loading the platform-specific evmone library.
// It does not verify that the loaded version is compatible with evmc version.
func InitEvmoneVM() (*evmc.VM, error) { _ = "STUB: not implemented"; return nil, nil }
