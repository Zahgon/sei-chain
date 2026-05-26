package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
	wasmvm "github.com/sei-protocol/sei-chain/sei-wasmvm"
	wasmvmtypes "github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

// contractMemoryLimit is the memory limit of each contract execution (in MiB)
// constant value so all nodes run with the same limit.
const contractMemoryLimit = 32

type contextKey int

const (
	// private type creates an interface key for Context that cannot be accessed by any other package
	contextKeyQueryStackSize contextKey = iota
	contextKeyCallDepth      contextKey = iota
)

// Option is an extension point to instantiate keeper with non default values
type Option interface {
	apply(*Keeper)
}

// WasmVMQueryHandler is an extension point for custom query handler implementations
type WasmVMQueryHandler interface {
	// HandleQuery executes the requested query
	HandleQuery(ctx sdk.Context, caller sdk.AccAddress, request wasmvmtypes.QueryRequest) ([]byte, error)
}

type CoinTransferrer interface {
	// TransferCoins sends the coin amounts from the source to the destination with rules applied.
	TransferCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// WasmVMResponseHandler is an extension point to handles the response data returned by a contract call.
type WasmVMResponseHandler interface {
	// Handle processes the data returned by a contract invocation.
	Handle(
		ctx sdk.Context,
		contractAddr sdk.AccAddress,
		ibcPort string,
		messages []wasmvmtypes.SubMsg,
		origRspData []byte,
		info wasmvmtypes.MessageInfo,
		codeInfo types.CodeInfo,
	) ([]byte, error)
}

// Keeper will have a reference to Wasmer with it's own data directory.
type Keeper struct {
	storeKey              sdk.StoreKey
	cdc                   codec.Codec
	accountKeeper         types.AccountKeeper
	bank                  CoinTransferrer
	portKeeper            types.PortKeeper
	capabilityKeeper      types.CapabilityKeeper
	paramsKeeper          types.ParamsKeeper
	upgradeKeeper         types.UpgradeKeeper
	wasmVM                types.WasmerEngine
	simulationWasmVM      types.WasmerEngine
	rpcWasmVM             types.WasmerEngine
	rpcWasmVM152          types.WasmerEngine
	rpcWasmVM155          types.WasmerEngine
	wasmVMQueryHandler    WasmVMQueryHandler
	wasmVMResponseHandler WasmVMResponseHandler
	messenger             Messenger
	// queryGasLimit is the max wasmvm gas that can be spent on executing a query with a contract
	queryGasLimit     uint64
	paramSpace        paramtypes.Subspace
	gasRegister       GasRegister
	maxQueryStackSize uint32
	maxCallDepth      uint32
}

// NewKeeper creates a new contract Keeper instance
// If customEncoders is non-nil, we can use this to override some of the message handler, especially custom
func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	paramsKeeper types.ParamsKeeper,
	paramSpace paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	distKeeper types.DistributionKeeper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	capabilityKeeper types.CapabilityKeeper,
	upgradeKeeper types.UpgradeKeeper,
	portSource types.ICS20TransferPortSource,
	router MessageRouter,
	queryRouter GRPCQueryRouter,
	homeDir string,
	wasmConfig types.WasmConfig,
	supportedFeatures string,
	opts ...Option,
) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// set KeyTable if it has not already been set

// not updateable, yet

func (k Keeper) getWasmer(ctx sdk.Context) types.WasmerEngine {
	_ = "STUB: not implemented"
	return *new(types.WasmerEngine)
}

func (k Keeper) getUploadAccessConfig(ctx sdk.Context) types.AccessConfig {
	_ = "STUB: not implemented"
	return *new(types.AccessConfig)
}

func (k Keeper) getInstantiateAccessConfig(ctx sdk.Context) types.AccessType {
	_ = "STUB: not implemented"
	return *new(types.AccessType)
}

// GetParams returns the total set of wasm parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

func (k Keeper) SetParams(ctx sdk.Context, ps types.Params) { _ = "STUB: not implemented"; return }

func (k Keeper) create(ctx sdk.Context, creator sdk.AccAddress, wasmCode []byte, instantiateAccess *types.AccessConfig, authZ AuthorizationPolicy) (codeID uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// figure out proper instantiate access

// we enforce this must be subset of default upload access

// #nosec G115 -- MaxWasmSize is checked above to be non-negative

func (k Keeper) storeCodeInfo(ctx sdk.Context, codeID uint64, codeInfo types.CodeInfo) {
	_ = "STUB: not implemented"
	return
}

// 0x01 | codeID (uint64) -> ContractInfo

func (k Keeper) ImportCode(ctx sdk.Context, codeID uint64, codeInfo types.CodeInfo, wasmCode []byte) error {
	_ = "STUB: not implemented"
	// #nosec G115 -- MaxWasmSize is a constant and always non-negative
	return nil
}

// 0x01 | codeID (uint64) -> ContractInfo

func (k Keeper) instantiate(ctx sdk.Context, codeID uint64, creator, admin sdk.AccAddress, initMsg []byte, label string, deposit sdk.Coins, authZ AuthorizationPolicy) (sdk.AccAddress, []byte, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil, nil
}

// create contract address

// deposit initial contract funds

// create an empty account (so we don't have issues later)
// TODO: can we remove this?

// get contact info

// prepare params for contract instantiate call

// create prefixed data store
// 0x03 | BuildContractAddress (sdk.AccAddress)

// prepare querier

// instantiate wasm contract

// persist instance first

// check for IBC flag

// register IBC port

// store contract before dispatch so that contract could be called back

// Execute executes the contract instance
func (k Keeper) execute(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add more funds

// prepare querier

func (k Keeper) migrate(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, newCodeID uint64, msg []byte, authZ AuthorizationPolicy) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check for IBC flag

// prevent update to non ibc contract

// add ibc port

// prepare querier

// delete old secondary index entry

// persist migration updates

// Sudo allows priviledged access to a contract. This can never be called by an external tx, but only by
// another native Go module directly, or on-chain governance (if sudo proposals are enabled). Thus, the keeper doesn't
// place any access controls on it, that is the responsibility or the app developer (who passes the wasm.Keeper in app.go)
func (k Keeper) Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// prepare querier

// reply is only called from keeper internal functions (dispatchSubmessages) after processing the submessage
func (k Keeper) reply(ctx sdk.Context, contractAddress sdk.AccAddress, reply wasmvmtypes.Reply) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// always consider this pinned

// prepare querier

// addToContractCodeSecondaryIndex adds element to the index for contracts-by-codeid queries
func (k Keeper) addToContractCodeSecondaryIndex(ctx sdk.Context, contractAddress sdk.AccAddress, entry types.ContractCodeHistoryEntry) {
	_ = "STUB: not implemented"
	return
}

// removeFromContractCodeSecondaryIndex removes element to the index for contracts-by-codeid queries
func (k Keeper) removeFromContractCodeSecondaryIndex(ctx sdk.Context, contractAddress sdk.AccAddress, entry types.ContractCodeHistoryEntry) {
	_ = "STUB: not implemented"
	return
}

// IterateContractsByCode iterates over all contracts with given codeID ASC on code update time.
func (k Keeper) IterateContractsByCode(ctx sdk.Context, codeID uint64, cb func(address sdk.AccAddress) bool) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) setContractAdmin(ctx sdk.Context, contractAddress, caller, newAdmin sdk.AccAddress, authZ AuthorizationPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) appendToContractHistory(ctx sdk.Context, contractAddr sdk.AccAddress, newEntries ...types.ContractCodeHistoryEntry) {
	_ = "STUB: not implemented"
	return
}

// find last element position

// then store with incrementing position

//nolint:gosec

func (k Keeper) GetContractHistory(ctx sdk.Context, contractAddr sdk.AccAddress) []types.ContractCodeHistoryEntry {
	_ = "STUB: not implemented"
	return nil
}

// getLastContractHistoryEntry returns the last element from history. To be used internally only as it panics when none exists
func (k Keeper) getLastContractHistoryEntry(ctx sdk.Context, contractAddr sdk.AccAddress) types.ContractCodeHistoryEntry {
	_ = "STUB: not implemented"
	return *new(types.ContractCodeHistoryEntry)
}

// all contracts have a history

// QuerySmartSafe queries the smart contract itself with a cached context
func (k Keeper) QuerySmartSafe(ctx sdk.Context, contractAddr sdk.AccAddress, req []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QuerySmart queries the smart contract itself.
func (k Keeper) QuerySmart(ctx sdk.Context, contractAddr sdk.AccAddress, req []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checks and increase query stack size

// prepare querier

// consume ALL remaining gas on error if the gasUsed is 0 due to error in consuming correct gas amount

func checkAndIncreaseQueryStackSize(ctx sdk.Context, maxQueryStackSize uint32) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *

	// read current value
	new(sdk.Context), nil
}

// increase

// did we go too far?

// set updated stack size

func checkAndIncreaseCallDepth(ctx sdk.Context, maxCallDepth uint32) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// increase

// did we go too far?

// set updated stack size

// QueryRaw returns the contract's state for give key. Returns `nil` when key is `nil`.
func (k Keeper) QueryRaw(ctx sdk.Context, contractAddress sdk.AccAddress, key []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) contractInstance(ctx sdk.Context, contractAddress sdk.AccAddress) (types.ContractInfo, types.CodeInfo, wasmvm.KVStore, error) {
	_ = "STUB: not implemented"
	return *new(types.ContractInfo), *new(types.CodeInfo), *new(wasmvm.KVStore), nil
}

func (k Keeper) GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *types.ContractInfo {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) HasContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

// storeContractInfo persists the ContractInfo. No secondary index updated here.
func (k Keeper) storeContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress, contract *types.ContractInfo) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) IterateContractInfo(ctx sdk.Context, cb func(sdk.AccAddress, types.ContractInfo) bool) {
	_ = "STUB: not implemented"
	return
}

// cb returns true to stop early

// IterateContractState iterates through all elements of the key value store for the given contract address and passes
// them to the provided callback function. The callback method can return true to abort early.
func (k Keeper) IterateContractState(ctx sdk.Context, contractAddress sdk.AccAddress, cb func(key, value []byte) bool) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) importContractState(ctx sdk.Context, contractAddress sdk.AccAddress, models []types.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) GetCodeInfo(ctx sdk.Context, codeID uint64) *types.CodeInfo {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) containsCodeInfo(ctx sdk.Context, codeID uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (k Keeper) IterateCodeInfos(ctx sdk.Context, cb func(uint64, types.CodeInfo) bool) {
	_ = "STUB: not implemented"
	return
}

// cb returns true to stop early

func (k Keeper) GetByteCode(ctx sdk.Context, codeID uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PinCode pins the wasm contract in wasmvm cache
func (k Keeper) pinCode(ctx sdk.Context, codeID uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// store 1 byte to not run into `nil` debugging issues

// UnpinCode removes the wasm contract from wasmvm cache
func (k Keeper) unpinCode(ctx sdk.Context, codeID uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// IsPinnedCode returns true when codeID is pinned in wasmvm cache
func (k Keeper) IsPinnedCode(ctx sdk.Context, codeID uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// InitializePinnedCodes updates wasmvm to pin to cache all contracts marked as pinned
func (k Keeper) InitializePinnedCodes(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

// setContractInfoExtension updates the extension point data that is stored with the contract info
func (k Keeper) setContractInfoExtension(ctx sdk.Context, contractAddr sdk.AccAddress, ext types.ContractInfoExtension) error {
	_ = "STUB: not implemented"
	return nil
}

// setAccessConfig updates the access config of a code id.
func (k Keeper) setAccessConfig(ctx sdk.Context, codeID uint64, config types.AccessConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// handleContractResponse processes the contract response data by emitting events and sending sub-/messages.
func (k *Keeper) handleContractResponse(
	ctx sdk.Context,
	contractAddr sdk.AccAddress,
	ibcPort string,
	msgs []wasmvmtypes.SubMsg,
	attrs []wasmvmtypes.EventAttribute,
	data []byte,
	evts wasmvmtypes.Events,
	info wasmvmtypes.MessageInfo,
	codeInfo types.CodeInfo,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// emit all events from this contract itself

// keep track of call depth

func (k Keeper) runtimeGasForContract(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

// infinite gas meter with limit=0 and not out of gas

func (k Keeper) consumeRuntimeGas(ctx sdk.Context, gas uint64) { _ = "STUB: not implemented"; return }

// throw OutOfGas error if we ran out (got exactly to zero due to better limit enforcing)

func (k Keeper) consumeRemainingGas(ctx sdk.Context) { _ = "STUB: not implemented"; return }

// generates a contract address from codeID + instanceID
func (k Keeper) generateContractAddress(ctx sdk.Context, codeID uint64) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// BuildContractAddress builds an sdk account address for a contract.
func BuildContractAddress(codeID, instanceID uint64) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

func (k Keeper) autoIncrementID(ctx sdk.Context, lastIDKey []byte) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// PeekAutoIncrementID reads the current value without incrementing it.
func (k Keeper) PeekAutoIncrementID(ctx sdk.Context, lastIDKey []byte) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (k Keeper) importAutoIncrementID(ctx sdk.Context, lastIDKey []byte, val uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) importContract(ctx sdk.Context, contractAddr sdk.AccAddress, c *types.ContractInfo, state []types.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) newQueryHandler(ctx sdk.Context, contractAddress sdk.AccAddress) QueryHandler {
	_ = "STUB: not implemented"
	return *new(QueryHandler)
}

func (k Keeper) emitCW721OwnerBeforeTransferIfApplicable(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) {
	_ = "STUB: not implemented"
	return
}

// MultipliedGasMeter wraps the GasMeter from context and multiplies all reads by out defined multiplier
type MultipliedGasMeter struct {
	originalMeter sdk.GasMeter
	GasRegister   GasRegister
}

func NewMultipliedGasMeter(originalMeter sdk.GasMeter, gr GasRegister) MultipliedGasMeter {
	_ = "STUB: not implemented"
	return *new(MultipliedGasMeter)
}

var _ wasmvm.GasMeter = MultipliedGasMeter{}

func (m MultipliedGasMeter) GasConsumed() sdk.Gas { _ = "STUB: not implemented"; return *new(sdk.Gas) }

func (k Keeper) gasMeter(ctx sdk.Context) MultipliedGasMeter {
	_ = "STUB: not implemented"
	return *new(MultipliedGasMeter)
}

// Querier creates a new grpc querier instance
func Querier(k *Keeper) *grpcQuerier {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

// QueryGasLimit returns the gas limit for smart queries.
func (k Keeper) QueryGasLimit() sdk.Gas {
	_ = "STUB: not implemented"
	return *

	// BankCoinTransferrer replicates the cosmos-sdk behaviour as in
	// https://github.com/cosmos/cosmos-sdk/blob/v0.41.4/x/bank/keeper/msg_server.go#L26
	new(sdk.Gas)
}

type BankCoinTransferrer struct {
	keeper types.BankKeeper
}

func NewBankCoinTransferrer(keeper types.BankKeeper) BankCoinTransferrer {
	_ = "STUB: not implemented"
	return *new(BankCoinTransferrer)
}

// TransferCoins transfers coins from source to destination account when coin send was enabled for them and the recipient
// is not in the blocked address list.
func (c BankCoinTransferrer) TransferCoins(parentCtx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amount sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// skip messages as we talk to the keeper directly

type msgDispatcher interface {
	DispatchSubmessages(ctx sdk.Context, contractAddr sdk.AccAddress, ibcPort string, msgs []wasmvmtypes.SubMsg, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]byte, error)
}

// DefaultWasmVMContractResponseHandler default implementation that first dispatches submessage then normal messages.
// The Submessage execution may include an success/failure response handling by the contract that can overwrite the
// original
type DefaultWasmVMContractResponseHandler struct {
	md msgDispatcher
}

func NewDefaultWasmVMContractResponseHandler(md msgDispatcher) *DefaultWasmVMContractResponseHandler {
	_ = "STUB: not implemented"
	return nil
}

// Handle processes the data returned by a contract invocation.
func (h DefaultWasmVMContractResponseHandler) Handle(ctx sdk.Context, contractAddr sdk.AccAddress, ibcPort string, messages []wasmvmtypes.SubMsg, origRspData []byte, info wasmvmtypes.MessageInfo, codeInfo types.CodeInfo) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
