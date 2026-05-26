package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "slashing"

	// StoreKey is the store key string for slashing
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute is the querier route for slashing
	QuerierRoute = ModuleName
)

// Keys for slashing store
// Items are stored with the following key: values
//
// - 0x01<consAddrLen (1 Byte)><consAddress_Bytes>: ValidatorSigningInfo
//
// - 0x02<consAddrLen (1 Byte)><consAddress_Bytes><period_Bytes>: bool
//
// - 0x03<accAddrLen (1 Byte)><accAddr_Bytes>: cryptotypes.PubKey
var (
	ValidatorSigningInfoKeyPrefix         = []byte{0x01} // Prefix for signing info
	ValidatorMissedBlockBitArrayKeyPrefix = []byte{0x02} // Prefix for missed block bit array
	AddrPubkeyRelationKeyPrefix           = []byte{0x03} // Prefix for address-pubkey relation
)

// ValidatorSigningInfoKey - stored by *Consensus* address (not operator address)
func ValidatorSigningInfoKey(v sdk.ConsAddress) []byte { _ = "STUB: not implemented"; return nil }

// ValidatorSigningInfoAddress - extract the address from a validator signing info key
func ValidatorSigningInfoAddress(key []byte) (v sdk.ConsAddress) {
	_ = "STUB: not implemented"
	// Remove prefix and address length.
	return *new(sdk.ConsAddress)
}

// ValidatorMissedBlockBitArrayKey - stored by *Consensus* address (not operator address)
func ValidatorMissedBlockBitArrayKey(v sdk.ConsAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// AddrPubkeyRelationKey gets pubkey relation key used to get the pubkey from the address
func AddrPubkeyRelationKey(addr []byte) []byte { _ = "STUB: not implemented"; return nil }
