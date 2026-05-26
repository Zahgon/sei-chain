package wsei

import (
	"embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const CurrentVersion uint16 = 1

//go:embed WSEI.abi
//go:embed WSEI.bin
var f embed.FS

var cachedBin []byte
var cachedABI *abi.ABI

func GetABI() []byte { _ = "STUB: not implemented"; return nil }

func GetParsedABI() *abi.ABI { _ = "STUB: not implemented"; return nil }

func GetBin() []byte { _ = "STUB: not implemented"; return nil }
