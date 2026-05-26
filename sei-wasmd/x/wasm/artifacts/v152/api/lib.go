package api

// #include <stdlib.h>
// #include "bindings152.h"
import "C"

import (
	"github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

// Value types
type (
	cint   = C.int
	cbool  = C.bool
	cusize = C.size_t
	cu8    = C.uint8_t
	cu32   = C.uint32_t
	cu64   = C.uint64_t
	ci8    = C.int8_t
	ci32   = C.int32_t
	ci64   = C.int64_t
)

// Pointers
type (
	cu8_ptr = *C.uint8_t
)

type Cache struct {
	ptr *C.cache_t
}

type Querier = types.Querier

func InitCache(dataDir string, supportedCapabilities string, cacheSize uint32, instanceMemoryLimit uint32) (Cache, error) {
	_ = "STUB: not implemented"
	return *new(Cache), nil
}

func ReleaseCache(cache Cache) { _ = "STUB: not implemented"; return }

func StoreCode(cache Cache, wasm []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func StoreCodeUnchecked(cache Cache, wasm []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveCode(cache Cache, checksum []byte) error { _ = "STUB: not implemented"; return nil }

func GetCode(cache Cache, checksum []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Pin(cache Cache, checksum []byte) error { _ = "STUB: not implemented"; return nil }

func Unpin(cache Cache, checksum []byte) error { _ = "STUB: not implemented"; return nil }

func AnalyzeCode(cache Cache, checksum []byte) (*types.AnalysisReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetMetrics(cache Cache) (*types.Metrics, error) { _ = "STUB: not implemented"; return nil, nil }

func Instantiate(
	cache Cache,
	checksum []byte,
	env []byte,
	info []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func Execute(
	cache Cache,
	checksum []byte,
	env []byte,
	info []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func Migrate(
	cache Cache,
	checksum []byte,
	env []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func Sudo(
	cache Cache,
	checksum []byte,
	env []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func Reply(
	cache Cache,
	checksum []byte,
	env []byte,
	reply []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func Query(
	cache Cache,
	checksum []byte,
	env []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func IBCChannelOpen(
	cache Cache,
	checksum []byte,
	env []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func IBCChannelConnect(
	cache Cache,
	checksum []byte,
	env []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func IBCChannelClose(
	cache Cache,
	checksum []byte,
	env []byte,
	msg []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func IBCPacketReceive(
	cache Cache,
	checksum []byte,
	env []byte,
	packet []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func IBCPacketAck(
	cache Cache,
	checksum []byte,
	env []byte,
	ack []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func IBCPacketTimeout(
	cache Cache,
	checksum []byte,
	env []byte,
	packet []byte,
	gasMeter *types.GasMeter,
	store types.KVStore,
	api *types.GoAPI,
	querier *Querier,
	gasLimit uint64,
	printDebug bool,
) ([]byte, types.GasReport, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.GasReport), nil
}

// Depending on the nature of the error, `gasUsed` will either have a meaningful value, or just 0.

func convertGasReport(report C.GasReport) types.GasReport {
	_ = "STUB: not implemented"
	return *new(types.GasReport)
}

/**** To error module ***/

func errorWithMessage(err error, b C.UnmanagedVector) error {
	_ = "STUB: not implemented"
	// this checks for out of gas as a special case
	return nil
}
