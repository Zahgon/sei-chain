package state

import (
	"github.com/ethereum/go-ethereum/common"
)

func (s *DBImpl) GetCodeHash(addr common.Address) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *DBImpl) GetCode(addr common.Address) []byte { _ = "STUB: not implemented"; return nil }

func (s *DBImpl) SetCode(addr common.Address, code []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// The SetCode method could be modified to return the old code/hash directly.

func (s *DBImpl) GetCodeSize(addr common.Address) int { _ = "STUB: not implemented"; return 0 }
