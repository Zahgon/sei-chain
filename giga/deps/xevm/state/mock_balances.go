//go:build mock_balances

package state

/*
==============================================================================
============================= !!! WARNING !!! ================================
==============================================================================
== This file is ONLY for TESTING/BENCHMARKING.                              ==
== It enables automatic top-off of EVM accounts with insufficient funds.    ==
== DO NOT USE IN PRODUCTION OR MAINNET BUILDS.                              ==
== This is enabled only when the 'mock_balances' build tag is set.          ==
==============================================================================
*/

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// TopOffAmount is the amount to mint when an account needs more funds (100 ETH)
var TopOffAmount = new(big.Int).Mul(big.NewInt(100), big.NewInt(1_000_000_000_000_000_000))

// ensureMinimumBalance tops off the account if balance is low.
// Called from GetBalance to ensure preCheck passes in StateTransition.
func (s *DBImpl) ensureMinimumBalance(evmAddr common.Address) { _ = "STUB: not implemented"; return }

// ensureSufficientBalance tops off the account if it doesn't have enough for the operation.
// Called from SubBalance before actually subtracting.
func (s *DBImpl) ensureSufficientBalance(evmAddr common.Address, amt *big.Int) {
	_ = "STUB: not implemented"
	return
}

// topOffAccount mints funds to an account.
func (s *DBImpl) topOffAccount(seiAddr sdk.AccAddress, amt *big.Int) {
	_ = "STUB: not implemented"
	// Ensure account exists
	return
}

// Mint and send (use NopLogger to suppress log spam)
