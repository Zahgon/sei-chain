package erc1155

import (
	"embed"
	"sync"
)

const CurrentVersion uint16 = 1

//go:embed cwerc1155.wasm
var f embed.FS

var cachedBin []byte
var cacheMtx = &sync.RWMutex{}

func GetBin() []byte { _ = "STUB: not implemented"; return nil }

func getCachedBin() []byte { _ = "STUB: not implemented"; return nil }

func setCachedBin(bin []byte) { _ = "STUB: not implemented"; return }
