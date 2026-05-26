package types

import (
	"github.com/gogo/protobuf/proto"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

const (
	defaultMemoryCacheSize    uint32 = 100 // in MiB
	defaultSmartQueryGasLimit uint64 = 3_000_000
	defaultContractDebugMode         = false

	// ContractAddrLen defines a valid address length for contracts
	ContractAddrLen = 32
	// SDKAddrLen defines a valid address length that was used in sdk address generation
	SDKAddrLen = 20
)

func (m Model) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func (c CodeInfo) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewCodeInfo fills a new CodeInfo struct
func NewCodeInfo(codeHash []byte, creator sdk.AccAddress, instantiatePermission AccessConfig) CodeInfo {
	_ = "STUB: not implemented"
	return *new(CodeInfo)
}

var AllCodeHistoryTypes = []ContractCodeHistoryOperationType{ContractCodeHistoryOperationTypeGenesis, ContractCodeHistoryOperationTypeInit, ContractCodeHistoryOperationTypeMigrate}

// NewContractInfo creates a new instance of a given WASM contract info
func NewContractInfo(codeID uint64, creator, admin sdk.AccAddress, label string, createdAt *AbsoluteTxPosition) ContractInfo {
	_ = "STUB: not implemented"
	return *new(ContractInfo)
}

// validatable is an optional interface that can be implemented by an ContractInfoExtension to enable validation
type validatable interface {
	ValidateBasic() error
}

// ValidateBasic does syntax checks on the data. If an extension is set and has the `ValidateBasic() error` method, then
// the method is called as well. It is recommend to implement `ValidateBasic` so that the data is verified in the setter
// but also in the genesis import process.
func (c *ContractInfo) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// SetExtension set new extension data. Calls `ValidateBasic() error` on non nil values when method is implemented by
// the extension.
func (c *ContractInfo) SetExtension(ext ContractInfoExtension) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadExtension copies the extension value to the pointer passed as argument so that there is no need to cast
// For example with a custom extension of type `MyContractDetails` it will look as following:
//
//	var d MyContractDetails
//	if err := info.ReadExtension(&d); err != nil {
//		return nil, sdkerrors.Wrap(err, "extension")
//	}
func (c *ContractInfo) ReadExtension(e ContractInfoExtension) error {
	_ = "STUB: not implemented"
	return nil
}

func (c ContractInfo) InitialHistory(initMsg []byte) ContractCodeHistoryEntry {
	_ = "STUB: not implemented"
	return *new(ContractCodeHistoryEntry)
}

func (c *ContractInfo) AddMigration(ctx sdk.Context, codeID uint64, msg []byte) ContractCodeHistoryEntry {
	_ = "STUB: not implemented"
	return *new(ContractCodeHistoryEntry)
}

// ResetFromGenesis resets contracts timestamp and history.
func (c *ContractInfo) ResetFromGenesis(ctx sdk.Context) ContractCodeHistoryEntry {
	_ = "STUB: not implemented"
	return *new(ContractCodeHistoryEntry)
}

// AdminAddr convert into sdk.AccAddress or nil when not set
func (c *ContractInfo) AdminAddr() sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// should never happen

// ContractInfoExtension defines the extension point for custom data to be stored with a contract info
type ContractInfoExtension interface {
	proto.Message
	String() string
}

var _ codectypes.UnpackInterfacesMessage = &ContractInfo{}

// UnpackInterfaces implements codectypes.UnpackInterfaces
func (c *ContractInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewAbsoluteTxPosition gets a block position from the context
func NewAbsoluteTxPosition(ctx sdk.Context) *AbsoluteTxPosition {
	_ = "STUB: not implemented"
	// we must safely handle nil gas meters
	return nil
}

// #nosec G115 -- checked above.
// #nosec G115 -- checked above.

// LessThan can be used to sort
func (a *AbsoluteTxPosition) LessThan(b *AbsoluteTxPosition) bool {
	_ = "STUB: not implemented"
	return false
}

// AbsoluteTxPositionLen number of elements in byte representation
const AbsoluteTxPositionLen = 16

// Bytes encodes the object into a 16 byte representation with big endian block height and tx index.
func (a *AbsoluteTxPosition) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// NewEnv initializes the environment for a contract instance
func NewEnv(ctx sdk.Context, contractAddr sdk.AccAddress) wasmvmtypes.Env {
	_ = "STUB: not implemented"
	// safety checks before casting below
	return *new(wasmvmtypes.Env)
}

// #nosec G115 -- checked above.

// NewInfo initializes the MessageInfo for a contract instance
func NewInfo(creator sdk.AccAddress, deposit sdk.Coins) wasmvmtypes.MessageInfo {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.MessageInfo)
}

// NewWasmCoins translates between Cosmos SDK coins and Wasm coins
func NewWasmCoins(cosmosCoins sdk.Coins) (wasmCoins []wasmvmtypes.Coin) {
	_ = "STUB: not implemented"
	return nil
}

// WasmConfig is the extra config required for wasm
type WasmConfig struct {
	// SimulationGasLimit is the max gas to be used in a tx simulation call.
	// When not set the consensus max block gas is used instead
	SimulationGasLimit *uint64
	// SimulationGasLimit is the max gas to be used in a smart query contract call
	SmartQueryGasLimit uint64
	// MemoryCacheSize in MiB not bytes
	MemoryCacheSize uint32
	// ContractDebugMode log what contract print
	ContractDebugMode bool
}

// DefaultWasmConfig returns the default settings for WasmConfig
func DefaultWasmConfig() WasmConfig { _ = "STUB: not implemented"; return *new(WasmConfig) }

// VerifyAddressLen ensures that the address matches the expected length
func VerifyAddressLen() func(addr []byte) error { _ = "STUB: not implemented"; return nil }

// IsSubset will return true if the caller is the same as the superset,
// or if the caller is more restrictive than the superset.
func (a AccessConfig) IsSubset(superSet AccessConfig) bool { _ = "STUB: not implemented"; return false }

// Everything is a subset of this

// Only an exact match is a subset of this

// An exact match or nobody
