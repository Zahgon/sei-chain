//go:build !mock_balances

package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ensureMinimumBalance is a no-op in production builds.
func (s *DBImpl) ensureMinimumBalance(evmAddr common.Address) {
	_ = "STUB: not implemented"

	// ensureSufficientBalance is a no-op in production builds.
	return
}

func (s *DBImpl) ensureSufficientBalance(evmAddr common.Address, amt *big.Int) {
	_ = "STUB: not implemented"
	return
}
