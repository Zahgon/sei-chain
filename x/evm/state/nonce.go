package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
)

func (s *DBImpl) GetNonce(addr common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

func (s *DBImpl) SetNonce(addr common.Address, nonce uint64, reason tracing.NonceChangeReason) {
	_ = "STUB: not implemented"
	return
}

// The SetCode method could be modified to return the old code/hash directly.
