package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// create a new DelegatorStartingInfo
func NewDelegatorStartingInfo(previousPeriod uint64, stake sdk.Dec, height uint64) DelegatorStartingInfo {
	_ = "STUB: not implemented"
	return *new(DelegatorStartingInfo)
}
