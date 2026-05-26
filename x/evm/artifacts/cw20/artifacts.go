package cw20

import (
	"embed"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const currentVersion uint16 = 2

var versionOverride uint16

// SetVersionWithOffset allows for overriding the version for integration test scenarios
func SetVersionWithOffset(offset int16) {
	_ = "STUB: not implemented"
	// this allows for negative offsets to mock lower versions
	return
}

//nolint:gosec

func CurrentVersion(ctx sdk.Context) uint16 { _ = "STUB: not implemented"; return 0 }

//go:embed CW20ERC20Pointer.abi
//go:embed CW20ERC20Pointer.bin
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
