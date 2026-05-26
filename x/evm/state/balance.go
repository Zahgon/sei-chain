package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/holiman/uint256"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var ZeroInt = uint256.NewInt(0)

func (s *DBImpl) SubBalance(evmAddr common.Address, amtUint256 *uint256.Int, reason tracing.BalanceChangeReason) uint256.Int {
	_ = "STUB: not implemented"
	return *new(uint256.Int)
}

// this avoids emitting cosmos events for ephemeral bookkeeping transfers like send_native

// Hook for mock balances (no-op in production builds)

func (s *DBImpl) AddBalance(evmAddr common.Address, amtUint256 *uint256.Int, reason tracing.BalanceChangeReason) uint256.Int {
	_ = "STUB: not implemented"
	return *new(uint256.Int)
}

// this avoids emitting cosmos events for ephemeral bookkeeping transfers like send_native

func (s *DBImpl) GetBalance(evmAddr common.Address) *uint256.Int {
	_ = "STUB: not implemented"
	return nil
}

// Hook for mock balances (no-op in production builds)

// should only be called during simulation
func (s *DBImpl) SetBalance(evmAddr common.Address, amtUint256 *uint256.Int, reason tracing.BalanceChangeReason) {
	_ = "STUB: not implemented"
	return
}

func (s *DBImpl) getSeiAddress(evmAddr common.Address) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (s *DBImpl) send(from sdk.AccAddress, to sdk.AccAddress, amt *big.Int) {
	_ = "STUB: not implemented"
	return
}
