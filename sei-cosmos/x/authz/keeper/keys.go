package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
)

// Keys for store prefixes
var (
	GrantKey = []byte{0x01} // prefix for each key
)

// StoreKey is the store key string for authz
const StoreKey = authz.ModuleName

// grantStoreKey - return authorization store key
// Items are stored with the following key: values
//
// - 0x01<granterAddressLen (1 Byte)><granterAddress_Bytes><granteeAddressLen (1 Byte)><granteeAddress_Bytes><msgType_Bytes>: Grant
func grantStoreKey(grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) []byte {
	_ = "STUB: not implemented"
	return nil
}

//	fmt.Println(">>>> len", l, key)

// addressesFromGrantStoreKey - split granter & grantee address from the authorization key
func addressesFromGrantStoreKey(key []byte) (granterAddr, granteeAddr sdk.AccAddress) {
	_ = "STUB: not implemented"
	// key is of format:
	// 0x01<granterAddressLen (1 Byte)><granterAddress_Bytes><granteeAddressLen (1 Byte)><granteeAddress_Bytes><msgType_Bytes>
	return *new(sdk.AccAddress), *new(sdk.AccAddress)
}

// remove prefix key

// firstAddressFromGrantStoreKey parses the first address only
func firstAddressFromGrantStoreKey(key []byte) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}
