package feegrant

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "feegrant"

	// StoreKey is the store key string for supply
	StoreKey = ModuleName

	// RouterKey is the message route for supply
	RouterKey = ModuleName

	// QuerierRoute is the querier route for supply
	QuerierRoute = ModuleName
)

var (
	// FeeAllowanceKeyPrefix is the set of the kvstore for fee allowance data
	FeeAllowanceKeyPrefix = []byte{0x00}
)

// FeeAllowanceKey is the canonical key to store a grant from granter to grantee
// We store by grantee first to allow searching by everyone who granted to you
func FeeAllowanceKey(granter sdk.AccAddress, grantee sdk.AccAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// FeeAllowancePrefixByGrantee returns a prefix to scan for all grants to this given address.
func FeeAllowancePrefixByGrantee(grantee sdk.AccAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ParseAddressesFromFeeAllowanceKey(key []byte) (granter, grantee sdk.AccAddress) {
	_ = "STUB: not implemented"
	// key is of format:
	// 0x00<granteeAddressLen (1 Byte)><granteeAddress_Bytes><granterAddressLen (1 Byte)><granterAddress_Bytes><msgType_Bytes>
	return *new(sdk.AccAddress), *new(sdk.AccAddress)
}

// remove prefix key
