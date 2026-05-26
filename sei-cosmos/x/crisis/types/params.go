package types

import (
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

var (
	// key for constant fee parameter
	ParamStoreKeyConstantFee = []byte("ConstantFee")
)

// type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	_ = "STUB: not implemented"
	return *new(paramtypes.KeyTable)
}

func validateConstantFee(i interface{}) error { _ = "STUB: not implemented"; return nil }
