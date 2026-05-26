package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
)

func (s *DBImpl) CreateAccount(acc common.Address) {
	_ = "STUB: not implemented"
	// clear any existing state but keep balance untouched, journaled for revert
	return
}

// too slow on historical DB so not doing it for tracing for now.
// could cause tracing to be incorrect in theory.

func (s *DBImpl) GetCommittedState(addr common.Address, hash common.Hash) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *DBImpl) GetState(addr common.Address, hash common.Hash) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *DBImpl) SetState(addr common.Address, key common.Hash, val common.Hash) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *DBImpl) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *DBImpl) SetTransientState(addr common.Address, key, val common.Hash) {
	_ = "STUB: not implemented"
	return
}

// debits account's balance. The corresponding credit happens here:
// https://github.com/sei-protocol/go-ethereum/blob/master/core/vm/instructions.go#L825
// clear account's state except the transient state (in Ethereum transient states are
// still available even after self destruction in the same tx)
func (s *DBImpl) SelfDestruct(acc common.Address) uint256.Int {
	_ = "STUB: not implemented"
	return *new(uint256.Int)
}

// remove the association

// mark account as self-destructed

func (s *DBImpl) SelfDestruct6780(acc common.Address) (uint256.Int, bool) {
	_ = "STUB: not implemented"
	// only self-destruct if acc is newly created in the same block
	return *new(uint256.Int), false
}

// the Ethereum semantics of HasSelfDestructed checks if the account is self destructed in the
// **CURRENT** block
func (s *DBImpl) HasSelfDestructed(acc common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// Snapshot records the current journal length as a revision and pushes the current
// EventManager onto the stack, creating a fresh one for subsequent events.
func (s *DBImpl) Snapshot() int { _ = "STUB: not implemented"; return 0 }

// Push current EM and create a fresh one so reverted events are discarded.

// RevertToSnapshot reverts all journal entries back to the snapshot identified by rev,
// restores the EventManager, and truncates the revision list.
func (s *DBImpl) RevertToSnapshot(rev int) {
	_ = "STUB: not implemented"
	// Binary-search for the revision with the given id (like go-ethereum).
	return
}

// Revert journal entries in reverse order down to the snapshot point.

// Restore the EventManager that was active when the snapshot was taken.
// snapshottedEventManagers has one entry per snapshot; idx corresponds to this snapshot.

// Truncate the revision list (removing this snapshot and any taken after it).

func (s *DBImpl) handleResidualFundsInDestructedAccounts(st *TemporaryState) {
	_ = "STUB: not implemented"
	return
}

// we don't want to really "burn" the token since it will mess up
// total supply calculation, so we send them to fee collector instead

func (s *DBImpl) clearAccountStateIfDestructed(st *TemporaryState) {
	_ = "STUB: not implemented"
	return
}

// clearAccountState unconditionally wipes code and storage for acc.
// Used by Finalize (self-destruct cleanup) and SetStorage. NOT journaled.
func (s *DBImpl) clearAccountState(acc common.Address) { _ = "STUB: not implemented"; return }

// clearAccountStateJournaled wipes code, nonce, and storage for acc, recording
// the previous values in the journal so a RevertToSnapshot can restore them.
// Called from CreateAccount (when not tracing).
func (s *DBImpl) clearAccountStateJournaled(acc common.Address) {
	_ = "STUB: not implemented"
	// Only clear if a code hash exists (mirrors clearAccountState logic).
	return
}

// Save previous state for potential revert.

// Collect all storage slots for this account using GetAllKeyStrsInRange.
// The prefix store's GetAllKeyStrsInRange returns raw parent-store keys,
// so we strip the per-address state prefix to obtain each slot hash.

// Append journal entry before making changes.

// Clear the account state.

func (s *DBImpl) MarkAccount(acc common.Address, status []byte) { _ = "STUB: not implemented"; return }

func (s *DBImpl) Created(acc common.Address) bool { _ = "STUB: not implemented"; return false }

func (s *DBImpl) SetStorage(addr common.Address, states map[common.Hash]common.Hash) {
	_ = "STUB: not implemented"
	return
}

func (s *DBImpl) getTransientAccount(acc common.Address) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *DBImpl) getTransientModule(key []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *DBImpl) getTransientState(acc common.Address, key common.Hash) (common.Hash, bool) {
	_ = "STUB: not implemented"
	return *new(common.Hash), false
}

func deleteIfExists(store storetypes.KVStore, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}
