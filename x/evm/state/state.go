package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func (s *DBImpl) CreateAccount(acc common.Address) { _ = "STUB: not implemented"; return }

// clear any existing state but keep balance untouched

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

func (s *DBImpl) getState(ctx sdk.Context, addr common.Address, hash common.Hash) common.Hash {
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

func (s *DBImpl) Snapshot() int { _ = "STUB: not implemented"; return 0 }

func (s *DBImpl) RevertToSnapshot(rev int) {
	_ = "STUB: not implemented"
	// Add bounds checking
	return
}

// Find the watermark index to truncate the journal

// Truncate the journal to remove reverted entries

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

func (s *DBImpl) clearAccountState(acc common.Address) { _ = "STUB: not implemented"; return }

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
