package ethtx

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type AuthList []SetCodeAuthorization

func NewAuthList(ethAuthList *[]ethtypes.SetCodeAuthorization) AuthList {
	_ = "STUB: not implemented"
	return *new(AuthList)
}

func (al AuthList) ToEthAuthList() *[]ethtypes.SetCodeAuthorization {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec
