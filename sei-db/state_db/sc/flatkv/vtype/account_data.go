package vtype

type AccountDataVersion uint8

// DO NOT CHANGE VERSION VALUES!!! Adding new versions is ok, but historical versions should never be removed/changed.
const (
	// The version of the account data field when FlatKV was first launched.
	AccountDataVersion0 AccountDataVersion = 0
)

/*
Serialization schema for AccountData version 0:

Full form (81 bytes):

| Version | Block Height | Balance  | Nonce    | Code Hash |
|---------|--------------|----------|----------|-----------|
| 1 byte  | 8 bytes      | 32 bytes | 8 bytes  | 32 bytes  |

Compact form (49 bytes) — used when code hash is all zeros:

| Version | Block Height | Balance  | Nonce    |
|---------|--------------|----------|----------|
| 1 byte  | 8 bytes      | 32 bytes | 8 bytes  |

Data is stored in big-endian order. At deserialization time, the two forms
are distinguished by length. The compact form is preferred for serialization
since ~97% of accounts have no code hash.
*/

const (
	accountVersionStart     = 0
	accountBlockHeightStart = accountVersionStart + VersionLength
	accountBalanceStart     = accountBlockHeightStart + BlockHeightLength
	accountNonceStart       = accountBalanceStart + BalanceLength
	accountCodeHashStart    = accountNonceStart + NonceLength

	accountCompactLength = VersionLength + BlockHeightLength + BalanceLength + NonceLength
	accountDataLength    = VersionLength + BlockHeightLength + BalanceLength + NonceLength + CodeHashLength
)

var _ VType = (*AccountData)(nil)

// Used for encapsulating and serializating account data in the FlatKV accounts database.
//
// This data structure is not threadsafe. Values passed into and values received from this data structure
// are not safe to modify without first copying them.
type AccountData struct {
	data []byte
}

// Create a new AccountData initialized to all 0s.
func NewAccountData() *AccountData { _ = "STUB: not implemented"; return nil }

// Serialize the account data to a byte slice. If the code hash is all zeros,
// the compact form (49 bytes) is returned; otherwise the full form (81 bytes).
//
// The returned byte slice is not safe to modify without first copying it.
func (a *AccountData) Serialize() []byte { _ = "STUB: not implemented"; return nil }

// Deserialize the account data from the given byte slice. Accepts both the
// compact (49 byte) and full (81 byte) forms.
func DeserializeAccountData(data []byte) (*AccountData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the serialization version for this AccountData instance.
func (a *AccountData) GetSerializationVersion() AccountDataVersion {
	_ = "STUB: not implemented"
	return *new(AccountDataVersion)
}

// Get the account's block height.
func (a *AccountData) GetBlockHeight() int64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec

// Get the account's balance.
func (a *AccountData) GetBalance() *Balance { _ = "STUB: not implemented"; return nil }

// Get the account's nonce.
func (a *AccountData) GetNonce() uint64 { _ = "STUB: not implemented"; return 0 }

// Get the account's code hash.
func (a *AccountData) GetCodeHash() *CodeHash { _ = "STUB: not implemented"; return nil }

// Check if this account data signifies a deletion operation. A deletion operation is automatically
// performed when all account data fields are 0 (with the exception of the serialization version and block height).
func (a *AccountData) IsDelete() bool { _ = "STUB: not implemented"; return false }

// Copy returns a deep copy of this AccountData. The copy has its own backing byte slice.
func (a *AccountData) Copy() *AccountData { _ = "STUB: not implemented"; return nil }

// Set the account's block height when this account was last modified/touched. Returns self.
func (a *AccountData) SetBlockHeight(blockHeight int64) *AccountData {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

// Set the account's balance. Returns self (or a new AccountData if nil).
func (a *AccountData) SetBalance(balance *Balance) *AccountData {
	_ = "STUB: not implemented"
	return nil
}

// Set the account's nonce. Returns self (or a new AccountData if nil).
func (a *AccountData) SetNonce(nonce uint64) *AccountData { _ = "STUB: not implemented"; return nil }

// Set the account's code hash. Returns self (or a new AccountData if nil).
func (a *AccountData) SetCodeHash(codeHash *CodeHash) *AccountData {
	_ = "STUB: not implemented"
	return nil
}
