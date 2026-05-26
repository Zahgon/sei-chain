package v630

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

// EmitEVMLog emits an EVM log from a precompile
func EmitEVMLog(evm *vm.EVM, address common.Address, topics []common.Hash, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// BlockNumber, BlockHash, TxHash, TxIndex, and Index are added later
// by the consensus engine when the block is being finalized.

// Event signatures for staking precompile
var (
	// Delegate(address indexed delegator, string validator, uint256 amount)
	DelegateEventSig = crypto.Keccak256Hash([]byte("Delegate(address,string,uint256)"))

	// Redelegate(address indexed delegator, string srcValidator, string dstValidator, uint256 amount)
	RedelegateEventSig = crypto.Keccak256Hash([]byte("Redelegate(address,string,string,uint256)"))

	// Undelegate(address indexed delegator, string validator, uint256 amount)
	UndelegateEventSig = crypto.Keccak256Hash([]byte("Undelegate(address,string,uint256)"))

	// ValidatorCreated(address indexed validator, string validatorAddress, string moniker)
	ValidatorCreatedEventSig = crypto.Keccak256Hash([]byte("ValidatorCreated(address,string,string)"))

	// ValidatorEdited(address indexed validator, string validatorAddress, string moniker)
	ValidatorEditedEventSig = crypto.Keccak256Hash([]byte("ValidatorEdited(address,string,string)"))
)

// Helper functions for common event patterns
func BuildDelegateEvent(delegator common.Address, validator string, amount *big.Int) ([]common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	// Pack the non-indexed data: validator string and amount
	// For strings in events, we need to encode: offset, length, and actual string data
	return nil, nil, nil
}

// 32 (offset) + 32 (amount) + 32 (string length) + paddedLen (string data)

// Offset for string (always 64 for first dynamic param when second param is uint256)

// Amount (uint256)

// String length
//nolint:gosec // validator names are short strings; no overflow risk

// String data (padded to 32 bytes)

// indexed

func EmitDelegateEvent(evm *vm.EVM, precompileAddr common.Address, delegator common.Address, validator string, amount *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func BuildRedelegateEvent(delegator common.Address, srcValidator, dstValidator string, amount *big.Int) ([]common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	// Pack the non-indexed data: srcValidator, dstValidator, amount
	return nil, nil, nil
}

// 3*32 (static) + 32 (srcLen) + paddedSrc + 32 (dstLen) + paddedDst

// offset for srcValidator. Static part is 3 * 32 = 96 bytes.

// placeholder offset for dstValidator, to be updated after we know the length of srcValidator

// amount

// srcValidator data part
// length of srcValidator
//nolint:gosec // validator names are short strings; no overflow risk
// data of srcValidator

// now calculate and update dstValidator offset

//nolint:gosec // offset is bounded by validator name lengths; no overflow risk

// dstValidator data part
// length of dstValidator
//nolint:gosec // validator names are short strings; no overflow risk
// data of dstValidator

// indexed

func EmitRedelegateEvent(evm *vm.EVM, precompileAddr common.Address, delegator common.Address, srcValidator, dstValidator string, amount *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func BuildUndelegateEvent(delegator common.Address, validator string, amount *big.Int) ([]common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	// Pack the non-indexed data: validator string and amount
	return nil, nil, nil
}

// 32 (offset) + 32 (amount) + 32 (string length) + paddedLen (string data)

// Offset for string

// Amount

// String length and data
//nolint:gosec // validator names are short strings; no overflow risk

// indexed

func EmitUndelegateEvent(evm *vm.EVM, precompileAddr common.Address, delegator common.Address, validator string, amount *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func BuildValidatorCreatedEvent(creator common.Address, validatorAddr string, moniker string) ([]common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	// Pack the non-indexed data: validatorAddr string and moniker string
	return nil, nil, nil
}

// 32 (offset1) + 32 (offset2) + 32 (addrLen) + paddedAddr + 32 (monikerLen) + paddedMoniker

// Offsets for two strings
// offset for validatorAddr
// placeholder offset for moniker

// validatorAddr string
//nolint:gosec // validator address is a short string; no overflow risk

// Adjust offset for moniker based on actual validatorAddr length

// Update the moniker offset in data
//nolint:gosec // offset is bounded by string lengths; no overflow risk

// moniker string
//nolint:gosec // moniker is a short string; no overflow risk

// indexed

func EmitValidatorCreatedEvent(evm *vm.EVM, precompileAddr common.Address, creator common.Address, validatorAddr string, moniker string) error {
	_ = "STUB: not implemented"
	return nil
}

func BuildValidatorEditedEvent(editor common.Address, validatorAddr string, moniker string) ([]common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	// Pack the non-indexed data: validatorAddr string and moniker string
	return nil, nil, nil
}

// 32 (offset1) + 32 (offset2) + 32 (addrLen) + paddedAddr + 32 (monikerLen) + paddedMoniker

// Offsets for two strings
// offset for validatorAddr
// placeholder offset for moniker

// validatorAddr string
//nolint:gosec // validator address is a short string; no overflow risk

// Adjust offset for moniker based on actual validatorAddr length

// Update the moniker offset in data
//nolint:gosec // offset is bounded by string lengths; no overflow risk

// moniker string
//nolint:gosec // moniker is a short string; no overflow risk

// indexed

func EmitValidatorEditedEvent(evm *vm.EVM, precompileAddr common.Address, editor common.Address, validatorAddr string, moniker string) error {
	_ = "STUB: not implemented"
	return nil
}
