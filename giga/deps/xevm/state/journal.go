package state

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type journalEntry interface {
	// revert undoes the changes introduced by this journal entry.
	revert(*DBImpl)
}

// revision marks a snapshot point in the journal.
type revision struct {
	id           int
	journalIndex int
}

type (
	accountStatusChange struct {
		account common.Address
		prev    []byte
	}

	addLogChange struct{}

	refundChange struct {
		prev uint64
	}

	// Changes to the access list
	accessListAddAccountChange struct {
		address common.Address
	}
	accessListAddSlotChange struct {
		address common.Address
		slot    common.Hash
	}

	// Changes to transient storage
	transientStorageChange struct {
		account       common.Address
		key, prevalue common.Hash
	}

	surplusChange struct {
		delta sdk.Int
	}

	// storageChange records a KV storage mutation so it can be reverted.
	storageChange struct {
		addr common.Address
		key  common.Hash
		prev common.Hash
	}

	// codeChange records a code mutation so it can be reverted.
	codeChange struct {
		addr           common.Address
		prevCode       []byte
		prevCodeExists bool
		prevMapping    addressMappingState
	}

	// nonceChange records a nonce mutation so it can be reverted.
	nonceChange struct {
		addr       common.Address
		prev       uint64
		prevExists bool
	}

	// balanceChange records an Add or Sub balance so it can be reverted.
	balanceChange struct {
		evmAddr common.Address
		seiAddr sdk.AccAddress
		usei    sdk.Int
		wei     sdk.Int
		isAdd   bool // true if AddBalance was called
	}

	// createAccountChange records the previous state cleared by clearAccountStateJournaled.
	createAccountChange struct {
		addr            common.Address
		prevCode        []byte
		prevCodeExists  bool
		prevNonce       uint64
		prevNonceExists bool
		prevSlots       map[common.Hash]common.Hash
	}

	// deleteMappingChange records a DeleteAddressMapping so it can be reverted.
	deleteMappingChange struct {
		evmAddr common.Address
		seiAddr sdk.AccAddress
	}

	addressMappingState struct {
		exists              bool
		seiAddr             sdk.AccAddress
		accountCreated      bool
		globalAccountNumber []byte
	}
)

func (e *accessListAddAccountChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *accessListAddSlotChange) revert(s *DBImpl) {
	_ = "STUB: not implemented"
	// since slot change always comes after address change, and revert
	// happens in reverse order, the address access list hasn't been
	// cleared at this point.
	return
}

// If the address was already removed or has no slots (idx == -1),
// there is nothing to revert.

// Bounds check in case a prior revert already modified the slots slice.

func (e *surplusChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *addLogChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *refundChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *transientStorageChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

// If the per-account transient map was already removed by a later revert,
// there is nothing to delete.

// A prior revert may have deleted the per-account map when it became empty.
// Re-create it so we can restore a non-zero prevalue.

func (e *accountStatusChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *storageChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *codeChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *nonceChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *balanceChange) revert(s *DBImpl) {
	_ = "STUB: not implemented"
	// Suppress events on revert
	return
}

// Was AddBalance: reverse by subtracting

// Was SubBalance: reverse by adding

func (e *createAccountChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *deleteMappingChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func captureAddressMapping(s *DBImpl, addr common.Address) addressMappingState {
	_ = "STUB: not implemented"
	return *new(addressMappingState)
}

func (m addressMappingState) restore(s *DBImpl, addr common.Address) {
	_ = "STUB: not implemented"
	return
}

func restoreCode(s *DBImpl, addr common.Address, code []byte, exists bool) {
	_ = "STUB: not implemented"
	return
}

func restoreNonce(s *DBImpl, addr common.Address, nonce uint64, exists bool) {
	_ = "STUB: not implemented"
	return
}
