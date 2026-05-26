package native

import (
	"embed"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const CurrentVersion uint16 = 1

//go:embed NativeSeiTokensERC20.abi
//go:embed NativeSeiTokensERC20.bin
var f embed.FS

var cachedBin []byte
var cachedABI *abi.ABI
var cacheMtx *sync.RWMutex = &sync.RWMutex{}

func GetABI() []byte { _ = "STUB: not implemented"; return nil }

func GetParsedABI() *abi.ABI { _ = "STUB: not implemented"; return nil }

func GetBin() []byte { _ = "STUB: not implemented"; return nil }

func getCachedABI() *abi.ABI { _ = "STUB: not implemented"; return nil }

func setCachedABI(a *abi.ABI) { _ = "STUB: not implemented"; return }

func getCachedBin() []byte { _ = "STUB: not implemented"; return nil }

func setCachedBin(bin []byte) { _ = "STUB: not implemented"; return }

func IsCodeFromBin(code []byte) bool { _ = "STUB: not implemented"; return false }
