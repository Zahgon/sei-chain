package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

func TransferWithoutEvents(db vm.StateDB, sender, recipient common.Address, amount *uint256.Int) {
	_ = "STUB: not implemented"
	return
}
