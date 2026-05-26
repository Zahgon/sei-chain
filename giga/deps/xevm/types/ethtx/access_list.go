package ethtx

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type AccessList []AccessTuple

func NewAccessList(ethAccessList *ethtypes.AccessList) AccessList {
	_ = "STUB: not implemented"
	return *new(AccessList)
}

func (al AccessList) ToEthAccessList() *ethtypes.AccessList { _ = "STUB: not implemented"; return nil }
