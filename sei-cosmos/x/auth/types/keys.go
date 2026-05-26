package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// module name
	ModuleName = "auth"

	// StoreKey is string representation of the store key for auth
	StoreKey = "acc"

	// FeeCollectorName the root string for the fee collector account address
	FeeCollectorName = "fee_collector"

	// QuerierRoute is the querier route for auth
	QuerierRoute = ModuleName
)

var (
	// AddressStoreKeyPrefix prefix for account-by-address store
	AddressStoreKeyPrefix = []byte{0x01}

	// param key for global account number
	GlobalAccountNumberKey = []byte("globalAccountNumber")
)

// AddressStoreKey turn an address to key used to get it from the account store
func AddressStoreKey(addr sdk.AccAddress) []byte { _ = "STUB: not implemented"; return nil }

func CreateAddressStoreKeyFromBech32(addr string) []byte { _ = "STUB: not implemented"; return nil }
