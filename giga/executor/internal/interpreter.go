package internal

import (
	"github.com/ethereum/go-ethereum/core/vm"
)

var _ vm.IEVMInterpreter = (*EVMInterpreter)(nil)

// EVMInterpreter is a custom interpreter that delegates execution to evmone via EVMC.
type EVMInterpreter struct {
	hostContext *HostContext
	evm         *vm.EVM
	readOnly    bool
}

func NewEVMInterpreter(hostContext *HostContext, evm *vm.EVM) *EVMInterpreter {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the contract code via evmone.
func (e *EVMInterpreter) Run(callOpCode vm.OpCode, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	_ = "STUB: not implemented"
	// Increment the call depth which is restricted to 1024
	return nil, nil
}

// For CREATE/CREATE2, the initcode is in contract.Code, not in input.
// For regular calls, input contains the call data.

// Make sure the readOnly is only set if we aren't in readOnly yet.
// This also makes sure that the readOnly flag isn't removed for child calls.

// todo(pdrobnjak): sender and recipient might not be correctly propagated in case of DELEGATECALL

// Reset SSTORE gas adjustment before execution

//nolint:gosec // gosec: safe gas conversion

// Apply SSTORE gas adjustment for Sei's custom SSTORE cost.
// evmone uses standard EIP-2200 gas (20k), but Sei may have a different cost.
// The adjustment is tracked during SetStorage calls and applied here.
// Adjustment can be positive (charge more) or negative (refund/reduce).

// If gas goes negative, execution would have failed with out of gas

// Update the contract's gas to reflect what evmone consumed
// This is critical for proper gas accounting!
//nolint:gosec // safe conversion - gasLeft is always <= contract.Gas

// Apply gas refund to the EVM's refund counter
//nolint:gosec // safe conversion

func (e *EVMInterpreter) ReadOnly() bool { _ = "STUB: not implemented"; return false }
