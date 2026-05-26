package internal

import (
	"sync/atomic"

	"github.com/ethereum/evmc/v12/bindings/go/evmc"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

var _ evmc.HostContext = (*HostContext)(nil)

// StandardSstoreSetGasEIP2200 is the standard EIP-2200 SSTORE gas cost for setting
// storage from zero to non-zero. evmone uses this internally.
const StandardSstoreSetGasEIP2200 = uint64(20000)

// HostContextConfig holds configuration for the HostContext.
// Values are pre-computed from ChainConfig at construction time for efficiency.
type HostContextConfig struct {
	// SstoreGasDelta is the per-SSTORE gas delta (Sei custom cost - standard 20k).
	// This is added to gas consumption for each StorageAdded operation.
	// Can be negative if Sei cost is below standard (results in gas reduction).
	SstoreGasDelta int64
}

// NewHostContextConfig creates a HostContextConfig from the chain config.
// This extracts and pre-computes values needed by HostContext.
func NewHostContextConfig(chainConfig *params.ChainConfig) HostContextConfig {
	_ = "STUB: not implemented"
	return *new(HostContextConfig)
}

// Guard against overflow: seiSstoreGas is uint64, and casting to int64 would
// overflow if value > math.MaxInt64. This is a critical misconfiguration.

// Delta = Sei cost - standard cost (can be positive or negative)
// Safe to cast now that we've verified seiSstoreGas <= math.MaxInt64

type HostContext struct {
	vm     *evmc.VM
	evm    *vm.EVM
	config HostContextConfig
	// sstoreGasAdjustment accumulates the total extra gas to charge for SSTORE operations.
	// This is applied after evmone execution completes.
	sstoreGasAdjustment atomic.Int64
}

// NewHostContext creates a new HostContext with the given configuration.
func NewHostContext(vm *evmc.VM, evm *vm.EVM, config HostContextConfig) *HostContext {
	_ = "STUB: not implemented"
	return nil
}

// GetSstoreGasAdjustment returns the total accumulated SSTORE gas adjustment.
// This should be called after execution to apply the extra gas charge.
func (h *HostContext) GetSstoreGasAdjustment() int64 { _ = "STUB: not implemented"; return 0 }

// ResetSstoreGasAdjustment resets the accumulated SSTORE gas adjustment to zero.
func (h *HostContext) ResetSstoreGasAdjustment() { _ = "STUB: not implemented"; return }

func (h *HostContext) AccountExists(addr evmc.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *HostContext) GetStorage(addr evmc.Address, key evmc.Hash) evmc.Hash {
	_ = "STUB: not implemented"
	return *new(evmc.Hash)
}

func (h *HostContext) SetStorage(addr evmc.Address, key evmc.Hash, value evmc.Hash) evmc.StorageStatus {
	_ = "STUB: not implemented"
	return *new(evmc.StorageStatus)
}

// Accumulate SSTORE gas adjustment for StorageAdded operations.
// evmone uses standard EIP-2200 gas (20k), but Sei may have a different cost.
// We track the delta here and apply it after execution.
// Delta can be positive (higher cost) or negative (lower cost).

func (h *HostContext) GetBalance(addr evmc.Address) evmc.Hash {
	_ = "STUB: not implemented"
	return *new(evmc.Hash)
}

func (h *HostContext) GetCodeSize(addr evmc.Address) int { _ = "STUB: not implemented"; return 0 }

func (h *HostContext) GetCodeHash(addr evmc.Address) evmc.Hash {
	_ = "STUB: not implemented"
	return *new(evmc.Hash)
}

func (h *HostContext) GetCode(addr evmc.Address) []byte { _ = "STUB: not implemented"; return nil }

// todo(pdrobnjak): support historical selfdestruct logic as well
func (h *HostContext) Selfdestruct(addr evmc.Address, beneficiary evmc.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *HostContext) GetTxContext() evmc.TxContext {
	_ = "STUB: not implemented"
	return *new(evmc.TxContext)
}

//nolint:gosec // G115: safe integer conversions for Time and GasLimit

func (h *HostContext) GetBlockHash(number int64) evmc.Hash {
	_ = "STUB: not implemented"
	//nolint:gosec // G115: safe, block numbers are always positive
	return *new(evmc.Hash)
}

func (h *HostContext) EmitLog(addr evmc.Address, topics []evmc.Hash, data []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *HostContext) Execute(kind evmc.CallKind, recipient evmc.Address, sender evmc.Address, value evmc.Hash, input []byte, gas int64,
	depth int, static bool) ([]byte, int64, int64, evmc.Address, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, *new(evmc.Address), nil
}

// For CREATE/CREATE2, the input contains the initcode (constructor bytecode)
// For regular calls, fetch the code from the target address

// initcode is passed as input for contract creation

// todo(pdrobnjak): calculate/propagate created address

// The created address should be set in the execution result
// For now, return empty - this needs to be populated from evmone's result

func (h *HostContext) Call(
	kind evmc.CallKind, recipient evmc.Address, sender evmc.Address, value evmc.Hash, input []byte, gas int64,
	_ int, static bool, salt evmc.Hash, _ evmc.Address,
) ([]byte, int64, int64, evmc.Address, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, *new(evmc.Address), nil
}

//nolint:gosec // G115: safe integer conversions for gas values

// todo(pdrobnjak): sender and recipient might not be correctly propagated in case of DELEGATECALL

//nolint:gosec // G115: safe, leftoverGas won't exceed int64 max

func (h *HostContext) AccessAccount(addr evmc.Address) evmc.AccessStatus {
	_ = "STUB: not implemented"
	return *new(evmc.AccessStatus)
}

// After a cold access, add address to access list so subsequent accesses are warm

func (h *HostContext) AccessStorage(addr evmc.Address, key evmc.Hash) evmc.AccessStatus {
	_ = "STUB: not implemented"
	return *new(evmc.AccessStatus)
}

// After a cold access, add slot to access list so subsequent accesses are warm

func (h *HostContext) GetTransientStorage(addr evmc.Address, key evmc.Hash) evmc.Hash {
	_ = "STUB: not implemented"
	return *new(evmc.Hash)
}

func (h *HostContext) SetTransientStorage(addr evmc.Address, key evmc.Hash, value evmc.Hash) {
	_ = "STUB: not implemented"
	return
}

// getEVMRevision determines the EVM revision based on the current chain configuration
func (h *HostContext) getEVMRevision() evmc.Revision {
	_ = "STUB: not implemented"
	return *new(evmc.Revision)
}

// Get the rules for the current block

// Check from newest to oldest using rules
// NOTE: Prague support in evmone 0.12.0 may have incomplete gas rules,
// so we cap at Cancun for now until evmone is updated

// toEvmcError converts a Go error to an evmc.Error.
// The EVMC bindings expect Call() to return evmc.Error type, not standard Go errors.
//
// Note: The evmc Go bindings currently only expose evmc.Failure and evmc.Revert constants.
// Additional error codes like EVMC_OUT_OF_GAS (3), EVMC_INVALID_INSTRUCTION (4), etc.
// are defined in the C header (evmc/evmc.h) but not exported in the Go bindings.
// To add proper mapping for vm.ErrOutOfGas -> evmc.OutOfGas, the Go bindings in
// github.com/ethereum/evmc would need to be extended first.
func toEvmcError(err error) error { _ = "STUB: not implemented"; return nil }

// Already an evmc.Error, return as-is

// All other errors map to generic failure
// TODO: Add evmc.OutOfGas mapping once the Go bindings expose it

// To be called by an exported EVM create function which knows how to instantiate params like statedb.
func createEVMWithFailFastPrecompile(blockContext vm.BlockContext, statedb vm.StateDB, chainConfig *params.ChainConfig, vmConfig vm.Config) *vm.EVM {
	_ = "STUB: not implemented"
	return nil
}
