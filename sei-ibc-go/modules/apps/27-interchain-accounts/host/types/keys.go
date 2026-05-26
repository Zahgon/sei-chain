package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// SubModuleName defines the interchain accounts host module name
	SubModuleName = "icahost"

	// StoreKey is the store key string for the interchain accounts host module
	StoreKey = SubModuleName
)

// ContainsMsgType returns true if the sdk.Msg TypeURL is present in allowMsgs, otherwise false
func ContainsMsgType(allowMsgs []string, msg sdk.Msg) bool {
	_ = "STUB: not implemented"
	// check that wildcard * option for allowing all message types is the only string in the array, if so, return true
	return false
}
