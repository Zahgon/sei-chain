package state

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// all custom precompiles have an address greater than or equal to this address
var CustomPrecompileStartingAddr = common.HexToAddress("0x0000000000000000000000000000000000001001")

// Forked from go-ethereum, except journaling logic which is unnecessary with cacheKV

type accessList struct {
	Addresses map[common.Address]int
	Slots     []map[common.Hash]struct{}
}

func (s *DBImpl) AddressInAccessList(addr common.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *DBImpl) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (s *DBImpl) AddAddressToAccessList(addr common.Address) { _ = "STUB: not implemented"; return }

func (s *DBImpl) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	_ = "STUB: not implemented"
	return
}

// Address not present, or addr present but no slots there

// Slot already present, nothing to do (no journal entry needed)

func (s *DBImpl) Prepare(_ params.Rules, sender, coinbase common.Address, dest *common.Address, precompiles []common.Address, txAccesses ethtypes.AccessList) {
	_ = "STUB: not implemented"
	return
}

// If it's a create-tx, the destination will be added inside evm.create

// skip any custom precompile

func (s *DBImpl) getCurrentAccessList() *accessList { _ = "STUB: not implemented"; return nil }
