package state

import (
	"github.com/ethereum/go-ethereum/common"
)

// Exist reports whether the given account exists in state.
// Notably this should also return true for self-destructed accounts.
func (s *DBImpl) Exist(addr common.Address) bool { _ = "STUB: not implemented"; return false }

// check if the address exists as a contract

// check if the address exists as an EOA

// check if account has a balance

// go-ethereum impl considers just-deleted accounts as "exist" as well

// Empty returns whether the given account is empty. Empty
// is defined according to EIP161 (balance = nonce = code = 0).
func (s *DBImpl) Empty(addr common.Address) bool { _ = "STUB: not implemented"; return false }
