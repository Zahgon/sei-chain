package types

import (
	"errors"
	"fmt"
	"sync"

	simplelru "github.com/hashicorp/golang-lru/v2/simplelru"
	yaml "gopkg.in/yaml.v2"

	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

const (
	// Constants defined here are the defaults value for address.
	// You can use the specific values for your project.
	// Add the follow lines to the `main()` of your server.
	//
	//	config := sdk.GetConfig()
	//	config.SetBech32PrefixForAccount(yourBech32PrefixAccAddr, yourBech32PrefixAccPub)
	//	config.SetBech32PrefixForValidator(yourBech32PrefixValAddr, yourBech32PrefixValPub)
	//	config.SetBech32PrefixForConsensusNode(yourBech32PrefixConsAddr, yourBech32PrefixConsPub)
	//	config.SetPurpose(yourPurpose)
	//	config.SetCoinType(yourCoinType)
	//	config.Seal()

	// Bech32MainPrefix defines the main SDK Bech32 prefix of an account's address
	Bech32MainPrefix = "sei"

	// Purpose is the ATOM purpose as defined in SLIP44 (https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
	Purpose = 44

	// CoinType is the ATOM coin type as defined in SLIP44 (https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
	CoinType = 118

	// FullFundraiserPath is the parts of the BIP44 HD path that are fixed by
	// what we used during the ATOM fundraiser.
	FullFundraiserPath = "m/44'/118'/0'/0/0"

	// PrefixAccount is the prefix for account keys
	PrefixAccount = "acc"
	// PrefixValidator is the prefix for validator keys
	PrefixValidator = "val"
	// PrefixConsensus is the prefix for consensus keys
	PrefixConsensus = "cons"
	// PrefixPublic is the prefix for public keys
	PrefixPublic = "pub"
	// PrefixOperator is the prefix for operator keys
	PrefixOperator = "oper"

	// PrefixAddress is the prefix for addresses
	PrefixAddress = "addr"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32MainPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic
)

// cache variables
var (
	// AccAddress.String() is expensive and if unoptimized dominantly showed up in profiles,
	// yet has no mechanisms to trivially cache the result given that AccAddress is a []byte type.
	accAddrMu     sync.Mutex
	accAddrCache  *simplelru.LRU[string, string]
	consAddrMu    sync.Mutex
	consAddrCache *simplelru.LRU[string, string]
	valAddrMu     sync.Mutex
	valAddrCache  *simplelru.LRU[string, string]
)

func init() {
	var err error
	// in total the cache size is 61k entries. Key is 32 bytes and value is around 50-70 bytes.
	// That will make around 92 * 61k * 2 (LRU) bytes ~ 11 MB
	if accAddrCache, err = simplelru.NewLRU[string, string](60000, nil); err != nil {
		panic(err)
	}
	if consAddrCache, err = simplelru.NewLRU[string, string](500, nil); err != nil {
		panic(err)
	}
	if valAddrCache, err = simplelru.NewLRU[string, string](500, nil); err != nil {
		panic(err)
	}
}

// Address is a common interface for different types of addresses used by the SDK
type Address interface {
	Equals(Address) bool
	Empty() bool
	Marshal() ([]byte, error)
	MarshalJSON() ([]byte, error)
	Bytes() []byte
	String() string
	Format(s fmt.State, verb rune)
}

// Ensure that different address types implement the interface
var _ Address = AccAddress{}
var _ Address = ValAddress{}
var _ Address = ConsAddress{}

var _ yaml.Marshaler = AccAddress{}
var _ yaml.Marshaler = ValAddress{}
var _ yaml.Marshaler = ConsAddress{}

// ----------------------------------------------------------------------------
// account
// ----------------------------------------------------------------------------

// AccAddress a wrapper around bytes meant to represent an account address.
// When marshaled to a string or JSON, it uses Bech32.
type AccAddress []byte

// AccAddressFromHex creates an AccAddress from a hex string.
func AccAddressFromHex(address string) (addr AccAddress, err error) {
	_ = "STUB: not implemented"
	return *new(AccAddress), nil
}

// VerifyAddressFormat verifies that the provided bytes form a valid address
// according to the default address rules or a custom address verifier set by
// GetConfig().SetAddressVerifier()
func VerifyAddressFormat(bz []byte) error { _ = "STUB: not implemented"; return nil }

// MustAccAddressFromBech32 calls AccAddressFromBech32 and panics on error.
func MustAccAddressFromBech32(address string) AccAddress {
	_ = "STUB: not implemented"
	return *new(AccAddress)
}

// AccAddressFromBech32 creates an AccAddress from a Bech32 string.
func AccAddressFromBech32(address string) (addr AccAddress, err error) {
	_ = "STUB: not implemented"
	return *new(AccAddress), nil
}

// Returns boolean for whether two AccAddresses are Equal
func (aa AccAddress) Equals(aa2 Address) bool { _ = "STUB: not implemented"; return false }

// Returns boolean for whether an AccAddress is empty
func (aa AccAddress) Empty() bool { _ = "STUB: not implemented"; return false }

// Marshal returns the raw address bytes. It is needed for protobuf
// compatibility.
func (aa AccAddress) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"

	// Unmarshal sets the address to the given data. It is needed for protobuf
	// compatibility.
	return nil, nil
}

func (aa *AccAddress) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshals to JSON using Bech32.
func (aa AccAddress) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML marshals to YAML using Bech32.
func (aa AccAddress) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalJSON unmarshals from JSON assuming Bech32 encoding.
		nil
}

func (aa *AccAddress) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML unmarshals from JSON assuming Bech32 encoding.
func (aa *AccAddress) UnmarshalYAML(data []byte) error { _ = "STUB: not implemented"; return nil }

// Bytes returns the raw address bytes.
func (aa AccAddress) Bytes() []byte {
	_ = "STUB: not implemented"

	// String implements the Stringer interface.
	return nil
}

func (aa AccAddress) String() string { _ = "STUB: not implemented"; return "" }

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (aa AccAddress) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// ----------------------------------------------------------------------------
// validator operator
// ----------------------------------------------------------------------------

// ValAddress defines a wrapper around bytes meant to present a validator's
// operator. When marshaled to a string or JSON, it uses Bech32.
type ValAddress []byte

// ValAddressFromHex creates a ValAddress from a hex string.
func ValAddressFromHex(address string) (addr ValAddress, err error) {
	_ = "STUB: not implemented"
	return *new(ValAddress), nil
}

// ValAddressFromBech32 creates a ValAddress from a Bech32 string.
func ValAddressFromBech32(address string) (addr ValAddress, err error) {
	_ = "STUB: not implemented"
	return *new(ValAddress), nil
}

// Returns boolean for whether two ValAddresses are Equal
func (va ValAddress) Equals(va2 Address) bool { _ = "STUB: not implemented"; return false }

// Returns boolean for whether an AccAddress is empty
func (va ValAddress) Empty() bool { _ = "STUB: not implemented"; return false }

// Marshal returns the raw address bytes. It is needed for protobuf
// compatibility.
func (va ValAddress) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"

	// Unmarshal sets the address to the given data. It is needed for protobuf
	// compatibility.
	return nil, nil
}

func (va *ValAddress) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshals to JSON using Bech32.
func (va ValAddress) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML marshals to YAML using Bech32.
func (va ValAddress) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalJSON unmarshals from JSON assuming Bech32 encoding.
		nil
}

func (va *ValAddress) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML unmarshals from YAML assuming Bech32 encoding.
func (va *ValAddress) UnmarshalYAML(data []byte) error { _ = "STUB: not implemented"; return nil }

// Bytes returns the raw address bytes.
func (va ValAddress) Bytes() []byte {
	_ = "STUB: not implemented"

	// String implements the Stringer interface.
	return nil
}

func (va ValAddress) String() string { _ = "STUB: not implemented"; return "" }

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (va ValAddress) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// ----------------------------------------------------------------------------
// consensus node
// ----------------------------------------------------------------------------

// ConsAddress defines a wrapper around bytes meant to present a consensus node.
// When marshaled to a string or JSON, it uses Bech32.
type ConsAddress []byte

// ConsAddressFromHex creates a ConsAddress from a hex string.
func ConsAddressFromHex(address string) (addr ConsAddress, err error) {
	_ = "STUB: not implemented"
	return *new(ConsAddress), nil
}

// ConsAddressFromBech32 creates a ConsAddress from a Bech32 string.
func ConsAddressFromBech32(address string) (addr ConsAddress, err error) {
	_ = "STUB: not implemented"
	return *new(ConsAddress), nil
}

// get ConsAddress from pubkey
func GetConsAddress(pubkey cryptotypes.PubKey) ConsAddress {
	_ = "STUB: not implemented"
	return *new(ConsAddress)
}

// Returns boolean for whether two ConsAddress are Equal
func (ca ConsAddress) Equals(ca2 Address) bool { _ = "STUB: not implemented"; return false }

// Returns boolean for whether an ConsAddress is empty
func (ca ConsAddress) Empty() bool { _ = "STUB: not implemented"; return false }

// Marshal returns the raw address bytes. It is needed for protobuf
// compatibility.
func (ca ConsAddress) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"

	// Unmarshal sets the address to the given data. It is needed for protobuf
	// compatibility.
	return nil, nil
}

func (ca *ConsAddress) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshals to JSON using Bech32.
func (ca ConsAddress) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML marshals to YAML using Bech32.
func (ca ConsAddress) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalJSON unmarshals from JSON assuming Bech32 encoding.
		nil
}

func (ca *ConsAddress) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML unmarshals from YAML assuming Bech32 encoding.
func (ca *ConsAddress) UnmarshalYAML(data []byte) error { _ = "STUB: not implemented"; return nil }

// Bytes returns the raw address bytes.
func (ca ConsAddress) Bytes() []byte {
	_ = "STUB: not implemented"

	// String implements the Stringer interface.
	return nil
}

func (ca ConsAddress) String() string { _ = "STUB: not implemented"; return "" }

// Bech32ifyAddressBytes returns a bech32 representation of address bytes.
// Returns an empty sting if the byte slice is 0-length. Returns an error if the bech32 conversion
// fails or the prefix is empty.
func Bech32ifyAddressBytes(prefix string, bs []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustBech32ifyAddressBytes returns a bech32 representation of address bytes.
// Returns an empty sting if the byte slice is 0-length. It panics if the bech32 conversion
// fails or the prefix is empty.
func MustBech32ifyAddressBytes(prefix string, bs []byte) string {
	_ = "STUB: not implemented"
	return ""
}

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (ca ConsAddress) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// ----------------------------------------------------------------------------
// auxiliary
// ----------------------------------------------------------------------------

var errBech32EmptyAddress = errors.New("decoding Bech32 address failed: must provide a non empty address")

// GetFromBech32 decodes a bytestring from a Bech32 encoded string.
func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addressBytesFromHexString(address string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cacheBech32Addr is not concurrency safe. Concurrent access to cache causes race condition.
func cacheBech32Addr(prefix string, addr []byte, cache *simplelru.LRU[string, string], cacheKey string) string {
	_ = "STUB: not implemented"
	return ""
}
