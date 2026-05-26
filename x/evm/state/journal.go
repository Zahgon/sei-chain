package state

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type journalEntry interface {
	// revert undoes the changes introduced by this journal entry.
	revert(*DBImpl)
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

	watermark struct {
		version int
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

func (e *watermark) revert(s *DBImpl) { _ = "STUB: not implemented"; return }

func (e *accountStatusChange) revert(s *DBImpl) { _ = "STUB: not implemented"; return }
