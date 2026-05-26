package app

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethtests "github.com/ethereum/go-ethereum/tests"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
)

func Replay(a *App) { _ = "STUB: not implemented"; return }

//nolint:gosec

//nolint:gosec

func BlockTest(a *App, bt *ethtests.BlockTest) { _ = "STUB: not implemented"; return }

// Check post-state after all blocks are run

// Not checking compliance with EIP-4788

func encodeTx(tx *ethtypes.Transaction, txConfig client.TxConfig) []byte {
	_ = "STUB: not implemented"
	return nil
}

func IsWithdrawalAddress(addr common.Address, blocks []*ethtypes.Block) bool {
	_ = "STUB: not implemented"
	return false
}
