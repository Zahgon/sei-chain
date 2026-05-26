package cw721

import (
	"embed"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const CurrentVersion uint16 = 6

//go:embed CW721ERC721Pointer.abi
//go:embed CW721ERC721Pointer.bin
//go:embed legacy.bin
var f embed.FS

var cachedBin []byte
var cachedLegacyBin []byte
var cachedABI *abi.ABI
var cacheMtx *sync.RWMutex = &sync.RWMutex{}

func GetABI() []byte { _ = "STUB: not implemented"; return nil }

func GetParsedABI() *abi.ABI { _ = "STUB: not implemented"; return nil }

func GetBin() []byte { _ = "STUB: not implemented"; return nil }

func GetLegacyBin() []byte { _ = "STUB: not implemented"; return nil }

func getCachedABI() *abi.ABI { _ = "STUB: not implemented"; return nil }

func setCachedABI(a *abi.ABI) { _ = "STUB: not implemented"; return }

func getCachedBin() []byte { _ = "STUB: not implemented"; return nil }

func setCachedBin(bin []byte) { _ = "STUB: not implemented"; return }

func getCachedLegacyBin() []byte { _ = "STUB: not implemented"; return nil }

func setCachedLegacyBin(bin []byte) { _ = "STUB: not implemented"; return }

func IsCodeFromBin(code []byte) bool { _ = "STUB: not implemented"; return false }

func isCodeFromBin(code []byte, bin []byte) bool { _ = "STUB: not implemented"; return false }
