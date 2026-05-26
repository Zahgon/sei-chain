package keyring

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/hd"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Info is the publicly exposed information about a keypair
type Info interface {
	// Human-readable type for key listing
	GetType() KeyType
	// Name of the key
	GetName() string
	// Public key
	GetPubKey() cryptotypes.PubKey
	// Address
	GetAddress() types.AccAddress
	// Bip44 Path
	GetPath() (*hd.BIP44Params, error)
	// Algo
	GetAlgo() hd.PubKeyType
}

var (
	_ Info = &LocalInfo{}
	_ Info = &ledgerInfo{}
	_ Info = &offlineInfo{}
	_ Info = &multiInfo{}
)

// LocalInfo is the public information about a locally stored key
// Note: Algo must be last field in struct for backwards amino compatibility
type LocalInfo struct {
	Name         string             `json:"name"`
	PubKey       cryptotypes.PubKey `json:"pubkey"`
	PrivKeyArmor string             `json:"privkey.armor"`
	Algo         hd.PubKeyType      `json:"algo"`
}

func newLocalInfo(name string, pub cryptotypes.PubKey, privArmor string, algo hd.PubKeyType) Info {
	_ = "STUB: not implemented"
	return *new(Info)
}

// GetType implements Info interface
func (i LocalInfo) GetType() KeyType {
	_ = "STUB: not implemented"

	// GetType implements Info interface
	return *new(KeyType)
}

func (i LocalInfo) GetName() string {
	_ = "STUB: not implemented"

	// GetType implements Info interface
	return ""
}

func (i LocalInfo) GetPubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"

	// GetType implements Info interface
	return *new(cryptotypes.PubKey)
}

func (i LocalInfo) GetAddress() types.AccAddress {
	_ = "STUB: not implemented"
	return *new(types.AccAddress)
}

// GetType implements Info interface
func (i LocalInfo) GetAlgo() hd.PubKeyType {
	_ = "STUB: not implemented"

	// GetType implements Info interface
	return *new(hd.PubKeyType)
}

func (i LocalInfo) GetPath() (*hd.BIP44Params, error) { _ = "STUB: not implemented"; return nil, nil }

// ledgerInfo is the public information about a Ledger key
// Note: Algo must be last field in struct for backwards amino compatibility
type ledgerInfo struct {
	Name   string             `json:"name"`
	PubKey cryptotypes.PubKey `json:"pubkey"`
	Path   hd.BIP44Params     `json:"path"`
	Algo   hd.PubKeyType      `json:"algo"`
}

func newLedgerInfo(name string, pub cryptotypes.PubKey, path hd.BIP44Params, algo hd.PubKeyType) Info {
	_ = "STUB: not implemented"
	return *new(Info)
}

// GetType implements Info interface
func (i ledgerInfo) GetType() KeyType {
	_ = "STUB: not implemented"

	// GetName implements Info interface
	return *new(KeyType)
}

func (i ledgerInfo) GetName() string {
	_ = "STUB: not implemented"

	// GetPubKey implements Info interface
	return ""
}

func (i ledgerInfo) GetPubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"

	// GetAddress implements Info interface
	return *new(cryptotypes.PubKey)
}

func (i ledgerInfo) GetAddress() types.AccAddress {
	_ = "STUB: not implemented"
	return *new(types.AccAddress)
}

// GetPath implements Info interface
func (i ledgerInfo) GetAlgo() hd.PubKeyType {
	_ = "STUB: not implemented"

	// GetPath implements Info interface
	return *new(hd.PubKeyType)
}

func (i ledgerInfo) GetPath() (*hd.BIP44Params, error) { _ = "STUB: not implemented"; return nil, nil }

// offlineInfo is the public information about an offline key
// Note: Algo must be last field in struct for backwards amino compatibility
type offlineInfo struct {
	Name   string             `json:"name"`
	PubKey cryptotypes.PubKey `json:"pubkey"`
	Algo   hd.PubKeyType      `json:"algo"`
}

func newOfflineInfo(name string, pub cryptotypes.PubKey, algo hd.PubKeyType) Info {
	_ = "STUB: not implemented"
	return *new(Info)
}

// GetType implements Info interface
func (i offlineInfo) GetType() KeyType {
	_ = "STUB: not implemented"

	// GetName implements Info interface
	return *new(KeyType)
}

func (i offlineInfo) GetName() string {
	_ = "STUB: not implemented"

	// GetPubKey implements Info interface
	return ""
}

func (i offlineInfo) GetPubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"

	// GetAlgo returns the signing algorithm for the key
	return *new(cryptotypes.PubKey)
}

func (i offlineInfo) GetAlgo() hd.PubKeyType {
	_ = "STUB: not implemented"

	// GetAddress implements Info interface
	return *new(hd.PubKeyType)
}

func (i offlineInfo) GetAddress() types.AccAddress {
	_ = "STUB: not implemented"
	return *new(types.AccAddress)
}

// GetPath implements Info interface
func (i offlineInfo) GetPath() (*hd.BIP44Params, error) { _ = "STUB: not implemented"; return nil, nil }

// Deprecated: this structure is not used anymore and it's here only to allow
// decoding old multiInfo records from keyring.
// The problem with legacy.Cdc.UnmarshalLengthPrefixed - the legacy codec doesn't
// tolerate extensibility.
type multisigPubKeyInfo struct {
	PubKey cryptotypes.PubKey `json:"pubkey"`
	Weight uint               `json:"weight"`
}

// multiInfo is the public information about a multisig key
type multiInfo struct {
	Name      string               `json:"name"`
	PubKey    cryptotypes.PubKey   `json:"pubkey"`
	Threshold uint                 `json:"threshold"`
	PubKeys   []multisigPubKeyInfo `json:"pubkeys"`
}

// NewMultiInfo creates a new multiInfo instance
func NewMultiInfo(name string, pub cryptotypes.PubKey) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

// GetType implements Info interface
func (i multiInfo) GetType() KeyType {
	_ = "STUB: not implemented"

	// GetName implements Info interface
	return *new(KeyType)
}

func (i multiInfo) GetName() string {
	_ = "STUB: not implemented"

	// GetPubKey implements Info interface
	return ""
}

func (i multiInfo) GetPubKey() cryptotypes.PubKey {
	_ = "STUB: not implemented"

	// GetAddress implements Info interface
	return *new(cryptotypes.PubKey)
}

func (i multiInfo) GetAddress() types.AccAddress {
	_ = "STUB: not implemented"
	return *new(types.AccAddress)
}

// GetPath implements Info interface
func (i multiInfo) GetAlgo() hd.PubKeyType {
	_ = "STUB: not implemented"
	return *

	// GetPath implements Info interface
	new(hd.PubKeyType)
}

func (i multiInfo) GetPath() (*hd.BIP44Params, error) { _ = "STUB: not implemented"; return nil, nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (i multiInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// encoding info
func marshalInfo(i Info) []byte { _ = "STUB: not implemented"; return nil }

// decoding info
func unmarshalInfo(bz []byte) (info Info, err error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}

// After unmarshalling into &info, if we notice that the info is a
// multiInfo, then we unmarshal again, explicitly in a multiInfo this time.
// Since multiInfo implements UnpackInterfacesMessage, this will correctly
// unpack the underlying anys inside the multiInfo.
//
// This is a workaround, as go cannot check that an interface (Info)
// implements another interface (UnpackInterfacesMessage).
