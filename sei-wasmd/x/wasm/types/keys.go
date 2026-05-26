package types

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// ModuleName is the name of the contract module
	ModuleName = "wasm"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// TStoreKey is the string transient store representation
	TStoreKey = "transient_" + ModuleName

	// QuerierRoute is the querier route for the wasm module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the wasm module
	RouterKey = ModuleName
)

// nolint
var (
	CodeKeyPrefix                                  = []byte{0x01}
	ContractKeyPrefix                              = []byte{0x02}
	ContractStorePrefix                            = []byte{0x03}
	SequenceKeyPrefix                              = []byte{0x04}
	ContractCodeHistoryElementPrefix               = []byte{0x05}
	ContractByCodeIDAndCreatedSecondaryIndexPrefix = []byte{0x06}
	PinnedCodeIndexPrefix                          = []byte{0x07}
	TXCounterPrefix                                = []byte{0x08}

	KeyLastCodeID     = append(SequenceKeyPrefix, []byte("lastCodeId")...)
	KeyLastInstanceID = append(SequenceKeyPrefix, []byte("lastContractId")...)
)

// GetCodeKey constructs the key for retreiving the ID for the WASM code
func GetCodeKey(codeID uint64) []byte { _ = "STUB: not implemented"; return nil }

// GetContractAddressKey returns the key for the WASM contract instance
func GetContractAddressKey(addr sdk.AccAddress) []byte { _ = "STUB: not implemented"; return nil }

// GetContractStorePrefix returns the store prefix for the WASM contract instance
func GetContractStorePrefix(addr sdk.AccAddress) []byte { _ = "STUB: not implemented"; return nil }

// GetContractByCreatedSecondaryIndexKey returns the key for the secondary index:
// `<prefix><codeID><created/last-migrated><contractAddr>`
func GetContractByCreatedSecondaryIndexKey(contractAddr sdk.AccAddress, c ContractCodeHistoryEntry) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetContractByCodeIDSecondaryIndexPrefix returns the prefix for the second index: `<prefix><codeID>`
func GetContractByCodeIDSecondaryIndexPrefix(codeID uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetContractCodeHistoryElementKey returns the key a contract code history entry: `<prefix><contractAddr><position>`
func GetContractCodeHistoryElementKey(contractAddr sdk.AccAddress, pos uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetContractCodeHistoryElementPrefix returns the key prefix for a contract code history entry: `<prefix><contractAddr>`
func GetContractCodeHistoryElementPrefix(contractAddr sdk.AccAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetPinnedCodeIndexPrefix returns the key prefix for a code id pinned into the wasmvm cache
func GetPinnedCodeIndexPrefix(codeID uint64) []byte { _ = "STUB: not implemented"; return nil }

// ParsePinnedCodeIndex converts the serialized code ID back.
func ParsePinnedCodeIndex(s []byte) uint64 { _ = "STUB: not implemented"; return 0 }
