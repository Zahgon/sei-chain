package erc20

import (
	"embed"
	"sync"
)

const CurrentVersion uint16 = 2

//go:embed cwerc20.wasm
var f embed.FS

var cachedBin []byte
var cacheMtx *sync.RWMutex = &sync.RWMutex{}

func GetBin() []byte { _ = "STUB: not implemented"; return nil }

func getCachedBin() []byte { _ = "STUB: not implemented"; return nil }

func setCachedBin(bin []byte) { _ = "STUB: not implemented"; return }
